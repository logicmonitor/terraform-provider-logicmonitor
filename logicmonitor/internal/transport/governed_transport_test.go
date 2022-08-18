package transport

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"terraform-provider-logicmonitor/logicmonitor/internal/apimutex"
	"testing"
)

func TestPreRequestHook(t *testing.T) {
	limit := 25
	remaining := 10
	path := "/santaba/rest/dashboard/dashboards"

	client := &http.Client{}
	apiMutex, _ := apimutex.NewAPIMutex()
	transport := NewGovernedTransport(client.Transport, apiMutex)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	apiMutex.Update(limit, remaining, http.MethodGet, path)
	if err := transport.preRequestHook(ctx, http.MethodGet, path); err != nil {
		t.Errorf("Didn't expect error, got %+v", err)
	}

	remaining--
	apiMutex.Update(0, 10, http.MethodGet, path)
	if err := transport.preRequestHook(ctx, http.MethodGet, path); err != context.Canceled {
		t.Errorf("Expected %v error, got %+v", context.Canceled, err)
	}
}

func TestPostRequestHook(t *testing.T) {
	client := &http.Client{}
	apiMutex, _ := apimutex.NewAPIMutex()
	transport := NewGovernedTransport(client.Transport, apiMutex)

	path := "/santaba/rest/dashboard/dashboards"
	request := http.Request{
		URL: &url.URL{
			Path: path,
		},
	}
	limit := 25
	remaining := 10
	remainingWindow := 5
	headers := http.Header{}
	headers.Add("X-Rate-Limit-Limit", fmt.Sprintf("%v", limit))
	headers.Add("X-Rate-Limit-Remaining", fmt.Sprintf("%v", remaining))
	headers.Add("X-Rate-Limit-Window", fmt.Sprintf("%v", remainingWindow))
	response := http.Response{
		Request: &request,
		Header:  headers,
	}

	transport.postRequestHook(http.MethodGet, path, &response)
	status := apiMutex.Status(http.MethodGet, path)
	if status.GetLimitRemaining() != remaining || status.GetRemainingWindow() != remainingWindow {
		t.Fatalf("expected %q api mutex status %+v to have limit %d, and remaining %d values", path, status, limit, remaining)
	}
}
