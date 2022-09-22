package transport

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"terraform-provider-logicmonitor/logicmonitor/internal/apimutex"
	"time"

	"github.com/hashicorp/go-hclog"
)

const (
	X_RATE_LIMIT_LIMIT     = "X-Rate-Limit-Limit"
	X_RATE_LIMIT_REMAINING = "X-Rate-Limit-Remaining"
	X_RATE_LIMIT_WINDOW    = "X-Rate-Limit-Window"
)

type GovernedTransport struct {
	base     http.RoundTripper
	apiMutex *apimutex.APIMutex
	logger   hclog.Logger
	lock     sync.Mutex
}

// NewGovernedTransport returns a governed transport that relies on pre- and post-
// requests from the http round tripper. The pre request consults the api mutex
// to determine if sleeping for the Logicmonitor API bucket is called for.
// The post request updates the information it is holding about the current api
// rate limits.
func NewGovernedTransport(base http.RoundTripper, apiMutex *apimutex.APIMutex) *GovernedTransport {
	return &GovernedTransport{
		base:     base,
		apiMutex: apiMutex,
	}
}

// RoundTrip returns the final http response after it has managed the api rate
// limit accounting in the pre and post request hooks.
func (t *GovernedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// avoid parallelisms
	t.lock.Lock()
	defer t.lock.Unlock()

	path := req.URL.Path
	log.Println("[INFO]Request Path: ", path)
	if err := t.preRequestHook(req.Context(), req.Method, path); err != nil {
		return nil, err
	}
	resp, err := t.base.RoundTrip(req)
	// always attempt to save x-headers
	t.postRequestHook(req.Method, path, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (t *GovernedTransport) preRequestHook(ctx context.Context, method, path string) error {

	log.Println("[INFO] PREREQUEST HOOK")
	if t.apiMutex.HasCapacity(method, path) {
		log.Println("[INFO] CAPACITY AVAILABLE")
		return nil
	}

	status := t.apiMutex.Status(method, path)
	timeToSleep := status.GetRemainingWindow()

	line := fmt.Sprintf("Throttling API requests; sleeping for %d seconds until rate limit reset (path class %q: %d remaining of %d total); current request \"%s %s\"",
		timeToSleep,
		status.GetClass(),
		status.GetRemainingWindow(),
		status.GetLimitRemaining(),
		method,
		path,
	)
	log.Println("[INFO] : ", line)
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.NewTimer(time.Second * time.Duration(timeToSleep)).C:
		return nil
	}
}

func (t *GovernedTransport) postRequestHook(method, path string, resp *http.Response) {
	log.Printf("[INFO] POSTREQUEST HOOK")
	if resp == nil {
		log.Println("[ERROR] Response is nil so returning")
		return
	}
	window, err := strconv.ParseInt(resp.Header.Get(X_RATE_LIMIT_WINDOW), 10, 64)
	if err != nil {
		log.Printf("[ERROR] %q response header is missing or invalid, skipping postRequestHook: %+v", X_RATE_LIMIT_WINDOW, err)
		return
	}
	remaining, err := strconv.Atoi(resp.Header.Get(X_RATE_LIMIT_REMAINING))
	if err != nil {
		log.Printf("[ERROR] %q response header is missing or invalid, skipping postRequestHook: %+v", X_RATE_LIMIT_REMAINING, err)
		return
	}
	t.apiMutex.Update(remaining, int(window), method, path)
}
