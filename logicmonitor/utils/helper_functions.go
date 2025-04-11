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
	// Convert *schema.Set to []interface{} first
	var list []interface{}

	if set, ok := r.(*schema.Set); ok {
		list = set.List() // Converts *schema.Set to []interface{}
	} else if arr, ok := r.([]interface{}); ok {
		list = arr // Already in expected format
	} else {
		return // Return empty if type is unknown
	}

	// Iterate over the list
	for _, i := range list {
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

// retrieve resource widget tokens from resource structure
func GetPropFromWTMap(d *schema.ResourceData, key string) (t []*models.WidgetToken) {
	if r, ok := d.GetOk(key); ok {
		return getPropFromWTInterface(r)
	}
	return
}

func getPropFromWTInterface(r interface{}) (t []*models.WidgetToken) {
	for _, i := range r.([]interface{}) {
		if m, ok := i.(map[string]interface{}); ok {
			var name = m["name"].(string)
			var value = m["value"].(string)
			model := &models.WidgetToken{
				Name:  name,
				Value: value,
			}
			t = append(t, model)
		}
	}
	return
}

func ConvertSetToInt32Slice(set interface{}) (slice []int32) {
    if set == nil {
        return
    }
    setList := set.([]interface{})
    slice = make([]int32, len(setList))
    for i, v := range setList {
        slice[i] = int32(v.(int))
    }
    return slice
}

func GetPropFromCCMap(d *schema.ResourceData, key string) (t []*models.Recipient) {
	if r, ok := d.GetOk(key); ok {
		return getPropFromCCInterface(r)
	}
	return
}
func GetPropFromDesMap(d *schema.ResourceData, key string) (t []*models.Chain) {
	if r, ok := d.GetOk(key); ok {
		return getPropFromDesInterface(r)
	}
	return
}

func getPropFromCCInterface(r interface{}) (t []*models.Recipient) {
	for _, i := range r.([]interface{}) {
		if m, ok := i.(map[string]interface{}); ok {
			var addr = m["addr"].(string)
			var contact = m["contact"].(string)
			var method = m["method"].(string)
			var Type = m["type"].(string)

			model := &models.Recipient{
				Addr:  addr,
				Contact: contact,
				Method: &method,
                Type: &Type,
			}
			t = append(t, model)
		}
	}
	return
}

func getPropFromDesInterface(r interface{}) (t []*models.Chain) {
    

    for _, i := range r.([]interface{}) {
    if m, ok := i.(map[string]interface{}); ok {
      var period = GetPeriod(m["period"].([]interface{}))
		var stages =  GetRecipient(m["stages"].([]interface{}))
		var Type = m["type"].(string)
            
            model := &models.Chain{
                Period: &period,
                Stages: stages,
				Type: &Type,
            }
            t = append(t, model)
        }
    }

    return 
}
func GetRecipient(d []interface{}) (t [][]*models.Recipient) {
	for _, stageData := range d {
		stage, ok := stageData.([]interface{})
		if !ok {
			continue
		}
		var stageRecipients []*models.Recipient
		for _, recipientData := range stage {
			recipient, ok := recipientData.(map[string]interface{})
			if !ok {
				continue
			}
			var addr = recipient["addr"].(string)
			var contact = recipient["contact"].(string)
			var method = recipient["method"].(string)
			var Type = recipient["type"].(string)

			model := &models.Recipient{
				Addr:    addr,
				Contact: contact,
				Method:  &method,
				Type:    &Type,
			}
			stageRecipients = append(stageRecipients, model)
		}
		t = append(t, stageRecipients)
	}
	return
}
func GetPeriod(d []interface{}) (t models.Period) {
	for _, i := range d {
		if m, ok := i.(map[string]interface{}); ok {
			var endMinutes = int32(m["end_minutes"].(int))
			var startMinutes = int32(m["start_minutes"].(int))
			var timezone = m["timezone"].(string)
            var weekDays = ConvertSetToInt32Slice(m["week_days"])
			
			model := models.Period{
				EndMinutes: &endMinutes,
				StartMinutes: &startMinutes,
				Timezone: &timezone,
			    WeekDays: weekDays,
			}
			t = model
		}
	}
	return
}

func GetPropFromSteps(d *schema.ResourceData, key string) (t []*models.WebCheckStep) {
	if r, ok := d.GetOk(key); ok {
		return getPropFromStepsInterface(r)
	}
	return
}

func getPropFromStepsInterface(r interface{}) (t []*models.WebCheckStep) {
	for _, i := range r.([]interface{}) {
		if m, ok := i.(map[string]interface{}); ok {
			var schema = m["schema"].(string)
			var matchType = m["match_type"].(string)
            var description = m["description"].(string)
			var httpVersion = m["http_version"].(string)
			var respType = m["resp_type"].(string)
			var reqType = m["req_type"].(string)
			var followRedirection = m["follow_redirection"].(bool)
			var httpMehod = m["http_method"].(string)
			var enable = m["enable"].(bool)
			var name = m["name"].(string)
			var timeout = int32(m["timeout"].(int))
			var useDefaultRoot = m["use_default_root"].(bool)
			var postDataEditType = m["post_data_edit_type"].(string)
			var fullpageLoad = m["fullpage_load"].(bool)
			var requireAuth = m["require_auth"].(bool)
			var path = m["path"].(string)
            var auth *models.Authentication
            if requireAuth {
                auth = GetAuth(m["auth"].([]interface{}))
            }
			var httpHeaders = m["http_headers"].(string)
			var httpBody = m["http_body"].(string)
            var keyword = m["keyword"].(string)
			var respScript = m["resp_script"].(string)
			var label = m["label"].(string)
			var URL = m["url"].(string)
			var invertMatch = m["invert_match"].(bool)
			var reqScript = m["req_script"].(string)
			var typee = m["type"].(string)
			var statusCode = m["status_code"].(string)
            
			model := &models.WebCheckStep{
				Schema:  schema,
				MatchType: matchType,
                Description: description,
                HTTPVersion: httpVersion,
                RespType: respType,
				ReqType: reqType,
				FollowRedirection: followRedirection,
				HTTPMethod: httpMehod,
				Enable: enable,
				Name: name,
				Timeout: timeout,
				UseDefaultRoot: &useDefaultRoot,
				PostDataEditType: postDataEditType,
				FullpageLoad: fullpageLoad,
				RequireAuth: requireAuth,
				Path: path,
				HTTPHeaders: httpHeaders,
				HTTPBody: httpBody, 
				Keyword: keyword,
                RespScript: respScript,
				Label: label,
				URL: URL,
				InvertMatch: invertMatch,
                ReqScript: reqScript,
				Type: typee,
				StatusCode: statusCode,
               }
			   if requireAuth {
                model.Auth = auth
            }
			t = append(t, model)
		}
	}
	return
}
func GetAuth(d []interface{}) (t *models.Authentication) {
	for _, i := range d {
		if m, ok := i.(map[string]interface{}); ok {
			var userName = m["user_name"].(string)
			var password = m["password"].(string)
			var typee = m["type"].(string)
			var domain = m["domain"].(string)
            model := &models.Authentication{
				Password: &password,
				UserName: &userName,
				Type: &typee,
				Domain: domain,
			}
			t = model
		}
	}
	return
}
func GetCollectorAttribute(d []interface{}) (t models.CollectorAttribute) {
	for _, i := range d {
		if m, ok := i.(map[string]interface{}); ok {
           var name = m["name"].(string)
           model := models.CollectorAttribute{
              Name: &name,
			}
			t = model
		}
	}
	return
}
func GetPropFromDatapoint(d *schema.ResourceData, key string) (t []*models.DataPoint) {
	if r, ok := d.GetOk(key); ok {
		return getPropFromDPInterface(r)
	}
	return
}
func getPropFromDPInterface(r interface{}) (t []*models.DataPoint ) {
	for _, i := range r.([]interface{}) {
		if m, ok := i.(map[string]interface{}); ok {
		   var name = m["name"].(string)
		   var postProcessorMethod = m["post_processor_method"].(string)
		   var alertForNoData = int32(m["alert_for_no_data"].(int))
           var postProcessorParam = m["post_processor_param"].(string)
		   var maxDigits = int32(m["max_digits"].(int))
		   var description = m["description"].(string)
		   var alertClearTransitionInterval = int32(m["alert_clear_transition_interval"].(int))
		   var alertTransitionInterval = int32(m["alert_transition_interval"].(int))
		   var dataType = int32(m["data_type"].(int))
		   var maxValue = m["max_value"].(string)
		   var minValue = m["min_value"].(string)
		   var alertBody = m["alert_body"].(string)
		   var alertSubject = m["alert_subject"].(string)
		   var alertExpr = m["alert_expr"].(string)
		   var alertExprNote = m["alert_expr_note"].(string)
		   var typee = int32(m["type"].(int))
		   var rawDataFieldName = m["raw_data_field_name"].(string)
         model := &models.DataPoint{
				Name: &name,
				Description: description,
				AlertForNoData: alertForNoData,
				PostProcessorParam: postProcessorParam,
				PostProcessorMethod: postProcessorMethod,
				MaxDigits: maxDigits,
				AlertClearTransitionInterval: alertClearTransitionInterval,
				AlertTransitionInterval: alertTransitionInterval,
				DataType: dataType,
				MaxValue: maxValue,
                MinValue: minValue,
				AlertBody: alertBody,
				AlertSubject: alertSubject,
				AlertExpr: alertExpr,
				AlertExprNote: alertExprNote,
				Type: typee,
				RawDataFieldName: rawDataFieldName,
             }
			t = append(t, model)
		}
	}
	return
}
func GetPropFromLocationMap(d *schema.ResourceData, key string) (t *models.WebsiteLocation) {
	if r, ok := d.GetOk(key); ok {
		return getPropFromLocInterface(r)
	}
	return
}
func getPropFromLocInterface(r interface{}) (t *models.WebsiteLocation) {
    for _, i := range r.([]interface{}) {
    if m, ok := i.(map[string]interface{}); ok {
        var smgIds = ConvertSetToInt32Slice(m["smg_ids"])
		var collectorIds = ConvertSetToInt32Slice(m["collector_ids"])
		var all = m["all"].(bool)
          model := &models.WebsiteLocation{
                CollectorIds: collectorIds,
                SmgIds: smgIds,
				All: all,
				}
            t = model
        }
    }
 return 
}
func GetFilters(d []interface{}) []*models.AutoDiscoveryFilter {
	var filters []*models.AutoDiscoveryFilter
	for _, i := range d {
		if m, ok := i.(map[string]interface{}); ok {
			attribute := m["attribute"].(string)
			comment := m["comment"].(string)
			operation := m["operation"].(string)
			value := m["value"].(string)

			model := &models.AutoDiscoveryFilter{
				Attribute: &attribute,
				Comment:   comment,
				Operation: &operation,
				Value:     value,
			}
			filters = append(filters, model)
		}
	}
	return filters
}

