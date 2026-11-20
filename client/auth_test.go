package client

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"terraform-provider-logicmonitor/client/datasource"
)

// authServer answers every request with an empty datasource and records what came in. For LMv1 it checks the
// signature the way LogicMonitor does: recomputed from the request it received (method, epoch, body, resource path).
func authServer(t *testing.T, accessKey string) (*httptest.Server, *[]string) {
	var seen []string
	srv := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		seen = append(seen, auth)
		if r.Header.Get("X-Version") != "3" {
			t.Errorf("%s %s: X-Version %q", r.Method, r.URL.Path, r.Header.Get("X-Version"))
		}
		if strings.HasPrefix(auth, "LMv1 ") {
			parts := strings.Split(strings.TrimPrefix(auth, "LMv1 "), ":")
			body, _ := io.ReadAll(r.Body)
			path := strings.TrimPrefix(r.URL.Path, DefaultBasePath)
			if want := lmComputeSignature(accessKey, r.Method, parts[2], [][]byte{body}, path); len(parts) != 3 || parts[1] != want {
				t.Errorf("%s %s: LMv1 signature does not verify: %q", r.Method, path, auth)
			}
		}
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"id":5,"name":"x","collectInterval":60,"collectMethod":"script","collectorAttribute":{"name":"script"}}`)
	}))
	return srv, &seen
}

func newTestClient(srv *httptest.Server, id, key, bearer string) *LogicMonitorRESTAPI {
	u, _ := url.Parse(srv.URL)
	cfg := NewConfig()
	cfg.SetAccessID(&id)
	cfg.SetAccessKey(&key)
	cfg.SetBearerToken(&bearer)
	cfg.SetAccountDomain(&u.Host)
	return New(cfg, srv.Client())
}

// Both request paths (the generated clients and PolymorphicREST) must use the same credential: LMv1 when only
// api_id/api_key are set (unchanged behaviour), Bearer when a bearer token is set.
func TestAuthHeaders(t *testing.T) {
	for _, tc := range []struct{ name, id, key, bearer, prefix string }{
		{"lmv1", "myid", "mykey", "", "LMv1 myid:"},
		{"bearer", "", "", "tok123", "Bearer tok123"},
		{"bearer wins over lmv1", "myid", "mykey", "tok123", "Bearer tok123"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			srv, seen := authServer(t, tc.key)
			defer srv.Close()
			c := newTestClient(srv, tc.id, tc.key, tc.bearer)
			if _, err := c.Datasource.GetDatasourceByID(datasource.NewGetDatasourceByIDParams().WithID(5)); err != nil {
				t.Fatalf("generated client: %v", err)
			}
			if _, err := c.PolymorphicREST(context.Background(), http.MethodPut, "/setting/datasources/5/graphs/7", map[string]interface{}{"name": "g"}); err != nil {
				t.Fatalf("PolymorphicREST: %v", err)
			}
			if len(*seen) != 2 {
				t.Fatalf("want 2 requests, got %d", len(*seen))
			}
			for _, h := range *seen {
				if !strings.HasPrefix(h, tc.prefix) {
					t.Errorf("Authorization %q, want prefix %q", h, tc.prefix)
				}
			}
		})
	}
}
