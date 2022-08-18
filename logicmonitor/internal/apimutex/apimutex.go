package apimutex

import (
	"log"
	"net/http"
	"regexp"
	"sync"
)

type APIMutex struct {
	lock     sync.Mutex
	status   map[string]*APIStatus
	capacity int
}

// APIStatus is used to hold rate limit information from LogicMonitor API
type APIStatus struct {
	limitRemaining  int
	remainingWindow int
	class           string
}

const (
	COLLECTOR_KEY       = "collectors"
	COLLECTOR_GROUP_KEY = "collector_group"
	DEVICE_KEY          = "device"
	DEVICE_GROUP_KEY    = "device_group"
	DASHBOARD_KEY       = "dashboard"
	DASHBOARD_GROUP_KEY = "dashboard_group"
	OTHER_KEY           = "other"
)

// NewAPIMutex returns a new api mutex object that represents untilized capacity
func NewAPIMutex() (*APIMutex, error) {
	status := map[string]*APIStatus{
		COLLECTOR_KEY:       {class: COLLECTOR_KEY},
		COLLECTOR_GROUP_KEY: {class: COLLECTOR_GROUP_KEY},
		DEVICE_KEY:          {class: DEVICE_KEY},
		DEVICE_GROUP_KEY:    {class: DEVICE_GROUP_KEY},
		DASHBOARD_KEY:       {class: DASHBOARD_KEY},
		DASHBOARD_GROUP_KEY: {class: DASHBOARD_GROUP_KEY},
		OTHER_KEY:           {class: OTHER_KEY},
	}
	return &APIMutex{
		capacity: 500,
		status:   status,
	}, nil
}

var (
	reDashboard      = regexp.MustCompile(`/santaba/rest/dashboard/dashboards/*[/\w]*$`)
	reDashboardGroup = regexp.MustCompile(`/santaba/rest/dashboard/groups/*[/\w]*$`)
	reCollector      = regexp.MustCompile(`/santaba/rest/setting/collector/collectors/*[/\w]*$`)
	reCollectorGroup = regexp.MustCompile(`/santaba/rest/setting/collector/groups/*[/\w]*$`)
	reDevice         = regexp.MustCompile(`/santaba/rest/device/devices/*[/\w]*$`)
	reDeviceGroup    = regexp.MustCompile(`/santaba/rest/device/groups/*[/\w]*$`)
)

// Update updates the known status for the given API endpoint. It is synchronous
// and intelligently accounts for new values regardless of parallelism.
func (m *APIMutex) Update(limitRemaining, remainingWindow int, method, endPoint string) {
	m.lock.Lock()
	defer m.lock.Unlock()

	key := m.normalizeKey(method, endPoint)
	status := m.status[key]
	status.remainingWindow = remainingWindow
	status.limitRemaining = limitRemaining
	log.Println("[INFO] Post Update: - STATUS ", status)
	return
}

// Remaining returns the current remaining value of the api status object.
func (s *APIStatus) GetRemainingWindow() int {
	return s.remainingWindow
}

func (s *APIStatus) GetLimitRemaining() int {
	return s.limitRemaining
}

// HasCapacity approximates if there is capacity below the api mutex's maximum
// capacity threshold.
func (m *APIMutex) HasCapacity(method, endPoint string) bool {
	status := m.get(method, endPoint)
	if status.remainingWindow != 0 {
		m.capacity = status.limitRemaining
	}
	return m.capacity > 0
}

func (m *APIMutex) get(method, endPoint string) *APIStatus {
	key := m.normalizeKey(method, endPoint)
	return m.status[key]
}

func (m *APIMutex) normalizeKey(method, endPoint string) string {
	getPostPutDelete := (http.MethodGet == method) ||
		(http.MethodPost == method) ||
		(http.MethodPut == method) ||
		(http.MethodDelete == method)
	var result string

	switch {
	case reCollector.MatchString(endPoint) && getPostPutDelete:
		result = COLLECTOR_KEY
	case reCollectorGroup.MatchString(endPoint) && getPostPutDelete:
		result = COLLECTOR_GROUP_KEY
	case reDevice.MatchString(endPoint) && getPostPutDelete:
		result = DEVICE_KEY
	case reDeviceGroup.MatchString(endPoint) && getPostPutDelete:
		result = DEVICE_GROUP_KEY
	case reDashboard.MatchString(endPoint) && getPostPutDelete:
		result = DASHBOARD_KEY
	case reDashboardGroup.MatchString(endPoint) && getPostPutDelete:
		result = DASHBOARD_GROUP_KEY
	default:
		result = "other"
	}
	return result
}

// Status return the APIStatus for the given class of endpoint.
func (m *APIMutex) Status(method, endPoint string) *APIStatus {
	return m.get(method, endPoint)
}

// Class returns the api endpoint class for this status.
func (s *APIStatus) GetClass() string {
	return s.class
}
