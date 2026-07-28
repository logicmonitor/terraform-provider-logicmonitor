package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// polymorphicBaseURL builds the REST base URL from transport configuration.
func (c *LogicMonitorRESTAPI) polymorphicBaseURL() string {
	cfg := c.transportCfg
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}
	scheme := "https"
	if len(cfg.Schemes) > 0 {
		scheme = cfg.Schemes[0]
	}
	return scheme + "://" + cfg.Host + cfg.BasePath
}

// polymorphicAuthPath returns the path segment used in LMv1 signatures (no query string).
func polymorphicAuthPath(requestPath string) string {
	if i := strings.Index(requestPath, "?"); i >= 0 {
		return requestPath[:i]
	}
	return requestPath
}

// PolymorphicEncodeBody serializes a JSON body for LMv1 signing.
// A nil map must not be passed through EncodeLMv1JSONBody: a typed-nil
// interface{} is not == nil, so encoding would produce "null\n" and break
// GET signatures (LM expects empty data for requests with no body).
func PolymorphicEncodeBody(body map[string]interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}
	return EncodeLMv1JSONBody(body)
}

// PolymorphicREST performs a raw JSON API request for polymorphic resources (e.g. Widget).
func (c *LogicMonitorRESTAPI) PolymorphicREST(ctx context.Context, method, path string, body map[string]interface{}) (map[string]interface{}, error) {
	bodyBytes, err := PolymorphicEncodeBody(body)
	if err != nil {
		return nil, err
	}
	authPath := polymorphicAuthPath(path)
	var bodyReader io.Reader
	if len(bodyBytes) > 0 {
		bodyReader = bytes.NewReader(bodyBytes)
	}
	req, err := http.NewRequestWithContext(ctx, method, c.polymorphicBaseURL()+path, bodyReader)
	if err != nil {
		return nil, err
	}
	if len(bodyBytes) > 0 {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Authorization", SignLMv1(c.accessID, c.accessKey, method, bodyBytes, authPath))
	req.Header.Set("X-version", "3")

	httpClient := c.httpClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("polymorphic API %s %s failed: %s", method, path, string(respBody))
	}
	if len(respBody) == 0 {
		return map[string]interface{}{}, nil
	}
	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// PolymorphicRESTList fetches a list endpoint with an optional filter query parameter.
func (c *LogicMonitorRESTAPI) PolymorphicRESTList(ctx context.Context, path string, filter *string) (map[string]interface{}, error) {
	if filter != nil && *filter != "" {
		path = path + "?filter=" + url.QueryEscape(*filter)
	}
	return c.PolymorphicREST(ctx, http.MethodGet, path, nil)
}
