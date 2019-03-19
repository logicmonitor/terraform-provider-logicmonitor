package logicmonitor

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/logicmonitor/lm-sdk-go/models"
)

// add device helper function
func getProperties(d *schema.ResourceData) (t []*models.NameAndValue) {
	// interate through hashmap to get custom/system properties
	if r, ok := d.GetOk("properties"); ok {
		for k, v := range r.(map[string]interface{}) {
			key := k
			value, _ := v.(string)
			t = append(t, &models.NameAndValue{Name: &key, Value: &value})
		}
	}
	return
}

// add widget token helper function
func getWidgetTokens(d *schema.ResourceData) (t []*models.WidgetToken) {
	// interate through hashmap to get custom/system properties
	if r, ok := d.GetOk("widget_tokens"); ok {
		for k, v := range r.(map[string]interface{}) {
			v, _ := v.(string)
			t = append(t, &models.WidgetToken{Name: k, Value: v})
		}
	}
	return
}

// get the filters for hostgroup lookup
func getFilters(d *schema.ResourceData) (t string) {
	m := d.Get("filters").(*schema.Set)
	var groupIds []string
	for _, v := range m.List() {
		p := v.(map[string]interface{})

		if p["property"].(string) != "" {
			groupIds = append(groupIds, fmt.Sprintf("%s%s\"%s\"", p["property"].(string), p["operator"].(string), p["value"].(string)))
		}

		if p["custom_property_name"].(string) != "" {
			groupIds = append(groupIds, fmt.Sprintf("customProperties.name%s\"%s\",customProperties.value%s\"%s\"", p["operator"].(string), p["custom_property_name"].(string), p["operator"].(string), p["custom_property_value"].(string)))
		}
	}

	t = strings.Join(groupIds, ",")
	return
}

// get the filters for collector lookup
func getCollectorFilters(d *schema.ResourceData) (t string) {
	m := d.Get("filters").(*schema.Set)
	var filters []string
	for _, v := range m.List() {
		p := v.(map[string]interface{})
		if p["property"].(string) != "" {
			filters = append(filters, fmt.Sprintf("%s%s\"%s\"", p["property"].(string), p["operator"].(string), p["value"].(string)))
		}

		if p["custom_property_name"].(string) != "" {
			filters = append(filters, fmt.Sprintf("customProperties.name%s\"%s\",customProperties.value%s\"%s\"", p["operator"].(string), p["custom_property_name"].(string), p["operator"].(string), p["custom_property_value"].(string)))
		}
	}
	t = strings.Join(filters, ",")
	return
}

// builds the device object with device properties
func makeDeviceObject(d *schema.ResourceData) (output models.Device) {
	var cid = int32(d.Get("collector").(int))
	var hostgroupID = d.Get("hostgroup_id").(string)
	var name = d.Get("ip_addr").(string)

	// if displayname is not there, we can automatically add ipaddr
	var displayname = d.Get("display_name").(string)
	if displayname == "" {
		displayname = name
	}

	output = models.Device{
		Name:                 &name,
		DisplayName:          &displayname,
		DisableAlerting:      d.Get("disable_alerting").(bool),
		HostGroupIds:         &hostgroupID,
		PreferredCollectorID: &cid,
		CustomProperties:     getProperties(d),
	}

	return
}

func makeDashboardGroupObject(d *schema.ResourceData) (output models.DashboardGroup) {
	var name = d.Get("name").(string)
	template := []byte(d.Get("template").(string))
	var jsonOutput map[string]interface{}
	json.Unmarshal(template, &jsonOutput)

	output = models.DashboardGroup{
		Name:         &name,
		Description:  d.Get("description").(string),
		ParentID:     int32(d.Get("parent_id").(int)),
		WidgetTokens: getWidgetTokens(d),
		Template:     jsonOutput,
	}

	return
}

func makeDashboardObject(d *schema.ResourceData) (output models.Dashboard) {
	var name = d.Get("name").(string)
	template := []byte(d.Get("template").(string))
	var jsonOutput map[string]interface{}
	json.Unmarshal(template, &jsonOutput)

	output = models.Dashboard{
		Name:         &name,
		Description:  d.Get("description").(string),
		GroupID:      int32(d.Get("group_id").(int)),
		WidgetTokens: getWidgetTokens(d),
		Sharable:     d.Get("public").(bool),
		Template:     jsonOutput,
	}
	return
}

// builds the device group object with host group properties
func makeDeviceGroupObject(d *schema.ResourceData) (output models.DeviceGroup) {
	var name = d.Get("name").(string)

	output = models.DeviceGroup{
		Name:             &name,
		DisableAlerting:  d.Get("disable_alerting").(bool),
		Description:      d.Get("description").(string),
		AppliesTo:        d.Get("applies_to").(string),
		ParentID:         int32(d.Get("parent_id").(int)),
		CustomProperties: getProperties(d),
	}

	return
}

// add collector group helper functions
func makeDeviceCollectorGroupObject(d *schema.ResourceData) (output models.CollectorGroup) {
	var name = d.Get("name").(string)
	output = models.CollectorGroup{
		Name:             &name,
		Description:      d.Get("description").(string),
		CustomProperties: getProperties(d),
	}

	return
}

// add collector helper functions
func makeDeviceCollectorObject(d *schema.ResourceData) models.Collector {
	output := models.Collector{
		BackupAgentID:                   int32(d.Get("backup_collector_id").(int)),
		CollectorGroupID:                int32(d.Get("collector_group_id").(int)),
		Description:                     d.Get("description").(string),
		EnableFailBack:                  d.Get("enable_failback").(bool),
		EnableFailOverOnCollectorDevice: d.Get("enable_collector_device_failover").(bool),
		EscalatingChainID:               int32(d.Get("escalation_chain_id").(int)),
		ResendIval:                      int32(d.Get("resend_interval").(int)),
		SuppressAlertClear:              d.Get("suppress_alert_clear").(bool),
		CustomProperties:                getProperties(d),
	}
	return output
}

// function to remove an item from array
func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

// check if entered string is numeric
func checkID(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
