package logicmonitor

import (
	"testing"
)

func TestGetDeviceTypeID(t *testing.T) {
	cases := []struct {
		name       string
		deviceType string
		expected   int32
	}{
		{"host", "host", 0},
		{"hostcaps", "HOST", 0},
		{"service", "service", 6},
		{"servicecaps", "SERVICE", 6},
		{"defaulttozero", "default", 0},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := getDeviceTypeID(tc.deviceType)
			if actual != tc.expected {
				t.Fatalf("expected: %d, got %d", tc.expected, actual)
			}
		})
	}
}

func TestGetDeviceType(t *testing.T) {
	cases := []struct {
		name         string
		deviceTypeID int32
		expected     string
	}{
		{"host", 0, "host"},
		{"service", 6, "service"},
		{"defaulttohost", 1, "host"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := getDeviceType(tc.deviceTypeID)
			if actual != tc.expected {
				t.Fatalf("expected: %s, got %s", tc.expected, actual)
			}
		})
	}
}
