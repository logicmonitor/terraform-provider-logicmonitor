package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"terraform-provider-logicmonitor/client"
	"terraform-provider-logicmonitor/client/collector"
	"terraform-provider-logicmonitor/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// convert string ID to floating point number
func IsID(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// retrieve installer url
func SetInstallerUrl(d *schema.ResourceData, m *models.Collector, client *client.LogicMonitorRESTAPI) diag.Diagnostics {
	var diags diag.Diagnostics

	params := collector.NewMiscGetCollectorDownloadTokenParams().WithID(m.ID)
	resp, err := client.Collector.MiscGetCollectorDownloadToken(params)
	if err != nil {
		diags = append(diags, diag.Errorf("Failed to retrieve colelctor download token: %s", err)...)
		return diags
	}

	respMap := resp.GetPayload().(map[string]interface{})
	token, tokenRetrieved := respMap["token"].(string)
	if !tokenRetrieved {
		diags = append(diags, diag.Errorf("unexpected %s", err)...)
		return diags
	}

	company := strings.ToLower(strings.TrimSpace(os.Getenv("LM_COMPANY")))
	if company == "" {
		company = strings.ToLower(strings.TrimSpace(d.Get("company").(string)))
		if company == "" {
			diags = append(diags, diag.Errorf("No 'company' field provided")...)
		}
	}

	arch := strings.ToLower(strings.TrimSpace(d.Get("arch").(string)))
	if arch != "linux" && arch != "linux32" && arch != "linux64" && arch != "win32" && arch != "win64" {
		diags = append(diags, diag.Errorf("'arch' must be formated as 'win' or 'linux' followed by '32' or '64' bit (e.g. 'linux64')")...)
	}

	url := fmt.Sprintf("https://%s.logicmonitor.com/santaba/rest/setting/collector/collectors/%d/bootstraps/%s?token=%s&useEA=%t&monitorOthers=%t&v=3",
		company, m.ID, arch, token, d.Get("ea").(bool), d.Get("monitor_others").(bool))

	collectorSize := strings.ToLower(strings.TrimSpace(d.Get("collector_size").(string)))
	if collectorSize != "nano" && collectorSize != "small" && collectorSize != "medium" && collectorSize != "large" {
		diags = append(diags, diag.Errorf("'collector_size' must be one of nano, small, medium, or large")...)
	} else {
		url = url + "&collectorSize=" + collectorSize
	}

	collectorVersion := strings.ReplaceAll(strings.TrimSpace(d.Get("build").(string)), ".", "")
	if collectorVersion != "" {
		url = url + "&collectorVersion=" + collectorVersion
	}

	if diags != nil {
		return diags
	}

	var urlMap map[string]string
	if arch == "win32" || arch == "win64" {
		urlMap = map[string]string{
			"URL": url,
		}
	} else {
		urlMap = map[string]string{
			"URL":      url,
			"cURL cmd": fmt.Sprintf("curl -o LogicmonitorBootstrap64_%d.bin '%s'", m.ID, url),
			"Wget cmd": fmt.Sprintf("wget -O LogicmonitorBootstrap64_%d.bin '%s'", m.ID, url),
		}
	}

	m.InstallerURLCmds = urlMap
	log.Printf("installer URL map: %v", m.InstallerURLCmds)
	return nil
}

// retrieve resource custom properties from map structure
func GetPropertiesFromMap(d map[string]interface{}, key string) (t []*models.NameAndValue) {
	if r, ok := d[key]; ok {
		return getPropertiesFromInterface(r)
	}
	return
}

// retrieve resource custom properties from resource structure
func GetPropertiesFromResource(d *schema.ResourceData, key string) (t []*models.NameAndValue) {
	if r, ok := d.GetOk(key); ok {
		return getPropertiesFromInterface(r)
	}
	return
}

func getPropertiesFromInterface(r interface{}) (t []*models.NameAndValue) {
	for _, i := range r.([]interface{}) {
		if m, ok := i.(map[string]interface{}); ok {
			var name = m["name"].(string)
			var value = m["value"].(string)
			model := &models.NameAndValue{
				Name:  &name,
				Value: &value,
			}
			t = append(t, model)
		}
	}
	return
}

// retrieve resource custom properties (techops version - json object input)
func GetPropertiesTechops(d *schema.ResourceData) (t []*models.NameAndValue) {
	// interate through hashmap to get custom/system properties
	if r, ok := d.GetOk("custom_properties"); ok {
		for k, v := range r.(map[string]interface{}) {
			key := k
			value, _ := v.(string)
			t = append(t, &models.NameAndValue{Name: &key, Value: &value})
		}
	}
	return
}

// remove an item from array
func Remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

// retrieve CloudCollectorConfig Objects
func GetCloudCollectorConfigs(d []interface{}) (t []*models.CloudCollectorConfig) {
	for _, i := range d {
		if m, ok := i.(map[string]interface{}); ok {
			var appliesTo = m["applies_to"].(string)
			var autoBalancedCollectorGroupId = int32(m["auto_balanced_collector_group_id"].(int))
			var collectorId = int32(m["collector_id"].(int))
			var priority = int32(m["priority"].(int))
			var usePublicIP = m["use_public_ip"].(bool)

			model := &models.CloudCollectorConfig{
				AppliesTo:                    &appliesTo,
				AutoBalancedCollectorGroupID: autoBalancedCollectorGroupId,
				CollectorID:                  &collectorId,
				Priority:                     &priority,
				UsePublicIP:                  usePublicIP,
			}
			t = append(t, model)
		}
	}
	return
}

// retrieve CloudTagFilter Objects
func GetCloudTagFilters(d []interface{}) (t []*models.CloudTagFilter) {
	for _, i := range d {
		if m, ok := i.(map[string]interface{}); ok {
			var name = m["name"].(string)
			var operation = m["operation"].(string)
			var value = m["value"].(string)

			model := &models.CloudTagFilter{
				Name:      &name,
				Operation: &operation,
				Value:     &value,
			}
			t = append(t, model)
		}
	}
	return
}

func ConvertSetToStringSlice(set *schema.Set) (slice []string) {
	if set == nil {
		return
	}
	setList := set.List()
	slice = make([]string, len(setList))
	for i, v := range setList {
		slice[i] = v.(string)
	}
	return
}
