package apimutex

import (
	"net/http"
	"testing"
)

func TestHasCapacity(t *testing.T) {
	amu, err := NewAPIMutex()
	if err != nil {
		t.Fatalf("api mutex constructor had error %+v", err)
	}

	endPoint := "/santaba/rest/dashboard/dashboards"
	amu.Update(200, 10, http.MethodGet, endPoint)
	if !amu.HasCapacity(http.MethodGet, endPoint) {
		t.Fatalf("ami mutex should have capacity, 90 limit, 46 remaining")
	}

	amu.Update(90, 10, http.MethodGet, endPoint)
	if !amu.HasCapacity(http.MethodGet, endPoint) {
		t.Fatalf("ami mutex should have capacity, 200 limit, 110 remaining")
	}

	amu.Update(0, 10, http.MethodGet, endPoint)
	if amu.HasCapacity(http.MethodGet, endPoint) {
		t.Fatalf("ami mutex shouldn't have capacity, 200 limit, 0 remaining")
	}
}

func TestGet(t *testing.T) {
	amu, err := NewAPIMutex()
	if err != nil {
		t.Fatalf("api mutex constructor had error %+v", err)
	}
	if len(amu.status) != 7 {
		t.Fatalf("amu status map should sized 7 but was sized %d", len(amu.status))
	}
	keys := []string{
		"collectors",
		"collector_group",
		"device",
		"device_group",
		"dashboard",
		"dashboard_group",
		"other",
	}
	for _, key := range keys {
		if _, found := amu.status[key]; !found {
			t.Fatalf("amu should have status for key %q", key)
		}
	}
}

func TestNormalizeKey(t *testing.T) {
	// Attempts to cover the rules listed at:s
	// https://developer.okta.com/docs/reference/rl-global-mgmt/
	tests := []struct {
		method   string
		endPoint string
		expected string
	}{
		//  1. [GET|POST|PUT|DELETE] /santaba/rest/dashboard/dashboards/
		{method: http.MethodGet, endPoint: "/santaba/rest/dashboard/dashboards", expected: "dashboard"},
		{method: http.MethodPost, endPoint: "/santaba/rest/dashboard/dashboards", expected: "dashboard"},
		{method: http.MethodPut, endPoint: "/santaba/rest/dashboard/dashboards/1", expected: "dashboard"},
		{method: http.MethodDelete, endPoint: "/santaba/rest/dashboard/dashboards/1", expected: "dashboard"},
		{method: http.MethodGet, endPoint: "/santaba/rest/dashboard/dashboards/1", expected: "dashboard"},
		{method: http.MethodPost, endPoint: "/santaba/rest/dashboard/dashboards/1", expected: "dashboard"},

		//  2. starts with /santaba/rest/dashboard/groups
		{method: http.MethodGet, endPoint: "/santaba/rest/dashboard/groups", expected: "dashboard_group"},
		{method: http.MethodPost, endPoint: "/santaba/rest/dashboard/groups", expected: "dashboard_group"},
		{method: http.MethodDelete, endPoint: "/santaba/rest/dashboard/groups/1", expected: "dashboard_group"},
		{method: http.MethodPut, endPoint: "/santaba/rest/dashboard/groups/2", expected: "dashboard_group"},
		{method: http.MethodGet, endPoint: "/santaba/rest/dashboard/groups/3", expected: "dashboard_group"},

		//  3. [GET|PUT|DELETE][POST] /santaba/rest/setting/collector/collectors
		{method: http.MethodGet, endPoint: "/santaba/rest/setting/collector/collectors", expected: "collectors"},
		{method: http.MethodPost, endPoint: "/santaba/rest/setting/collector/collectors", expected: "collectors"},
		{method: http.MethodGet, endPoint: "/santaba/rest/setting/collector/collectors/2", expected: "collectors"},
		{method: http.MethodDelete, endPoint: "/santaba/rest/setting/collector/collectors/4", expected: "collectors"},
		{method: http.MethodPut, endPoint: "/santaba/rest/setting/collector/collectors/4", expected: "collectors"},

		// // 14. GET /santaba/rest/
		{method: http.MethodGet, endPoint: "/santaba/rest/authorizationServers", expected: "other"},
		{method: http.MethodGet, endPoint: "/santaba/rest/authorizationServers/TESTID", expected: "other"},
		{method: http.MethodGet, endPoint: "/santaba/rest/behaviors", expected: "other"},
		{method: http.MethodGet, endPoint: "/santaba/rest/behaviors/TESTID", expected: "other"},
		{method: http.MethodGet, endPoint: "/santaba/rest/domains", expected: "other"},
		{method: http.MethodGet, endPoint: "/santaba/rest/domains/TESTID", expected: "other"},
		{method: http.MethodGet, endPoint: "/santaba/rest/idps", expected: "other"},
		{method: http.MethodGet, endPoint: "/santaba/rest/idps/TESTID", expected: "other"},
		{method: http.MethodGet, endPoint: "/santaba/rest/internal", expected: "other"},
		{method: http.MethodGet, endPoint: "/santaba/rest/internal/TESTID", expected: "other"},
		{method: http.MethodGet, endPoint: "/santaba/rest/mappings", expected: "other"},
		{method: http.MethodGet, endPoint: "/santaba/rest/mappings/TESTID", expected: "other"},
		{method: http.MethodGet, endPoint: "/santaba/rest/meta", expected: "other"},
		{method: http.MethodGet, endPoint: "/santaba/rest/meta/TESTID", expected: "other"},
		{method: http.MethodGet, endPoint: "/santaba/rest/org", expected: "other"},
		{method: http.MethodGet, endPoint: "/santaba/rest/org/TESTID", expected: "other"},
		{method: http.MethodGet, endPoint: "/santaba/rest/policies", expected: "other"},
		{method: http.MethodGet, endPoint: "/santaba/rest/policies/TESTID", expected: "other"},
		{method: http.MethodGet, endPoint: "/santaba/rest/templates", expected: "other"},
		{method: http.MethodGet, endPoint: "/santaba/rest/templates/TESTID", expected: "other"},
	}

	amu, err := NewAPIMutex()
	if err != nil {
		t.Fatalf("api mutex constructor had error %+v", err)
	}
	for _, tc := range tests {
		// test that private normalizedKey function is operating correctly
		key := amu.normalizeKey(tc.method, tc.endPoint)
		if key != tc.expected {
			t.Fatalf("got %q, expected %q for method: %q, endPoint: %q", key, tc.expected, tc.method, tc.endPoint)
		}
	}
}
