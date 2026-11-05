package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
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

// SuppressSensitiveIfRedactedOrEmpty prevents perpetual drift when APIs return
// redacted or empty values for secret fields during refresh.
func SuppressSensitiveIfRedactedOrEmpty(_ string, oldValue, newValue string, _ *schema.ResourceData) bool {
	// Keep create-time diffs visible when there is no prior state.
	if oldValue == "" && strings.TrimSpace(newValue) != "" {
		return false
	}

	oldMaskedOrEmpty := isSensitiveMaskedOrEmpty(oldValue)
	newMaskedOrEmpty := isSensitiveMaskedOrEmpty(newValue)

	// Suppress when either side is masked/empty because API refresh may redact secrets.
	if oldMaskedOrEmpty || newMaskedOrEmpty {
		return true
	}

	return false
}

func isSensitiveMaskedOrEmpty(value string) bool {
	normalized := strings.TrimSpace(strings.ToLower(value))
	if normalized == "" {
		return true
	}

	if strings.Contains(normalized, "redacted") || strings.Contains(normalized, "masked") {
		return true
	}

	allMaskChars := true
	for _, r := range normalized {
		if r != '*' && r != 'x' && r != 'X' && r != '-' {
			allMaskChars = false
			break
		}
	}
	if allMaskChars {
		return true
	}

	return false
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
	r := d.Get(key)
	return getPropertiesFromInterface(r)
}

func getPropertiesFromInterface(r interface{}) (t []*models.NameAndValue) {
	// Initialize empty slice to ensure we return [] instead of nil
	t = make([]*models.NameAndValue, 0)

	// Convert *schema.Set to []interface{} first
	var list []interface{}

	if set, ok := r.(*schema.Set); ok {
		list = set.List() // Converts *schema.Set to []interface{}
	} else if arr, ok := r.([]interface{}); ok {
		list = arr // Already in expected format
	} else {
		return t // Return empty slice if type is unknown
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
	return t
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

// IsNotFoundError checks whether an API error indicates a missing resource.
func IsNotFoundError(err error) bool {
	if err == nil {
		return false
	}
	errText := strings.ToLower(err.Error())
	return strings.Contains(errText, "[404]") ||
		strings.Contains(errText, " status 404") ||
		strings.Contains(errText, "not found")
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
func GetResourceProperties(d []interface{}) (t []*models.DeviceProperty ) {
	for _, i := range d {
		if m, ok := i.(map[string]interface{}); ok {
			var name = m["name"].(string)
			var value = m["value"].(string)

			model := &models.DeviceProperty{
				Name:      &name,
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

// ExpandJSONString parses a JSON document supplied as a string (e.g. via
// jsonencode(...) or file(...)) into a generic Go value so it can be sent to the
// API with its original structure and value types (objects, arrays, numbers,
// booleans) preserved. An empty string yields nil so the field is omitted, and
// an unparseable string is passed through unchanged.
func ExpandJSONString(raw string) interface{} {
	if strings.TrimSpace(raw) == "" {
		return nil
	}
	var parsed interface{}
	if err := json.Unmarshal([]byte(raw), &parsed); err != nil {
		log.Printf("[WARN] value is not valid JSON, sending as a raw string: %v", err)
		return raw
	}
	return parsed
}

// FlattenJSONToString serializes a generic value returned by the API back into a
// JSON string for storage in a TypeString schema field.
func FlattenJSONToString(v interface{}) string {
	if v == nil {
		return ""
	}
	if s, ok := v.(string); ok {
		return s
	}
	b, err := json.Marshal(v)
	if err != nil {
		log.Printf("[WARN] failed to serialize value to JSON: %v", err)
		return ""
	}
	return string(b)
}

// SuppressEquivalentJSON suppresses diffs between two JSON documents that are
// semantically equal but differ only in key ordering or whitespace, preventing
// perpetual drift when the API echoes back a normalized document.
func SuppressEquivalentJSON(_ string, oldValue, newValue string, _ *schema.ResourceData) bool {
	if oldValue == newValue {
		return true
	}

	oldTrimmed := strings.TrimSpace(oldValue)
	newTrimmed := strings.TrimSpace(newValue)
	// Keep create-time and cleared-value diffs visible.
	if oldTrimmed == "" || newTrimmed == "" {
		return false
	}

	var oldJSON, newJSON interface{}
	if err := json.Unmarshal([]byte(oldTrimmed), &oldJSON); err != nil {
		return false
	}
	if err := json.Unmarshal([]byte(newTrimmed), &newJSON); err != nil {
		return false
	}
	return reflect.DeepEqual(oldJSON, newJSON)
}

// canonicalPropertyValue returns a serialization-independent form of a
// name/value property value, so that two encodings of the same JSON document
// compare equal.
//
// Only JSON objects and arrays are canonicalized. Those are the values whose
// serialization is ambiguous (key order, indentation, escaping) and which the
// API re-serializes when it stores them, the Service Insight membership document
// held in "predef.bizservice.members" being the common case. Anything else,
// including JSON-looking scalars such as "1.50", is returned untouched so that a
// genuine change to a plain property is never hidden.
func canonicalPropertyValue(value string) string {
	trimmed := strings.TrimSpace(value)
	if !strings.HasPrefix(trimmed, "{") && !strings.HasPrefix(trimmed, "[") {
		return value
	}

	var parsed interface{}
	if err := json.Unmarshal([]byte(trimmed), &parsed); err != nil {
		return value
	}
	canonical, err := json.Marshal(parsed)
	if err != nil {
		return value
	}
	return string(canonical)
}

// HashNameAndValue is the TypeSet element hash for every name/value property
// collection (custom_properties, system_properties, resource_ids, ...).
//
// A set element is identified by its hash, so hashing the value verbatim makes a
// property whose JSON the API re-serialized (same document, different key order
// or whitespace) look like a different element, which Terraform can only report
// as "remove this element, add that element" on every plan and apply. Hashing
// the canonical form keeps such a property as a single stable element.
func HashNameAndValue(v interface{}) int {
	property, ok := v.(map[string]interface{})
	if !ok {
		return 0
	}

	name, _ := property["name"].(string)
	value, _ := property["value"].(string)
	canonicalValue := canonicalPropertyValue(value)

	// Length prefixes keep ("ab", "c") and ("a", "bc") from hashing alike.
	return schema.HashString(fmt.Sprintf("%d:%s%d:%s", len(name), name, len(canonicalValue), canonicalValue))
}

// PreserveEquivalentProperties returns apiProperties with every value the API
// only re-serialized replaced by the value already held under the same property
// name, which during a read is the practitioner's own configuration.
//
// Storing the API's serialization instead would leave state and configuration
// textually different for the rest of the resource's life, so the value would be
// rewritten on every apply even though nothing about it changed. Values that
// genuinely differ, and properties absent from the configuration, are returned
// exactly as the API reported them.
func PreserveEquivalentProperties(d *schema.ResourceData, key string, apiProperties []*models.NameAndValue) []*models.NameAndValue {
	if d == nil || len(apiProperties) == 0 {
		return apiProperties
	}

	knownValues := propertyValuesByName(d.Get(key))
	if len(knownValues) == 0 {
		return apiProperties
	}

	preserved := make([]*models.NameAndValue, 0, len(apiProperties))
	for _, apiProperty := range apiProperties {
		if apiProperty == nil || apiProperty.Name == nil || apiProperty.Value == nil {
			preserved = append(preserved, apiProperty)
			continue
		}

		name := *apiProperty.Name
		value := *apiProperty.Value
		for _, knownValue := range knownValues[name] {
			if knownValue != value && canonicalPropertyValue(knownValue) == canonicalPropertyValue(value) {
				value = knownValue
				break
			}
		}

		if value == *apiProperty.Value {
			preserved = append(preserved, apiProperty)
			continue
		}
		preserved = append(preserved, &models.NameAndValue{Name: &name, Value: &value})
	}
	return preserved
}

// propertyValuesByName groups the values currently held for a name/value
// property collection by property name, accepting both the set form used by
// resources and the list form used by data sources. A property name can
// legitimately repeat, so each name maps to every value recorded for it.
func propertyValuesByName(raw interface{}) map[string][]string {
	var list []interface{}
	switch typed := raw.(type) {
	case *schema.Set:
		list = typed.List()
	case []interface{}:
		list = typed
	default:
		return nil
	}

	values := make(map[string][]string, len(list))
	for _, element := range list {
		property, ok := element.(map[string]interface{})
		if !ok {
			continue
		}
		name, _ := property["name"].(string)
		value, _ := property["value"].(string)
		values[name] = append(values[name], value)
	}
	return values
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

// retrieve resource widget tokens from resource structure
func GetPropFromPriviligeMap(d *schema.ResourceData, key string) (t []*models.Privilege) {
	if r, ok := d.GetOk(key); ok {
		return getPropFromPrivilegeInterface(r)
	}
	return
}

func getPropFromPrivilegeInterface(r interface{}) (t []*models.Privilege) {
	for _, i := range r.([]interface{}) {
		if m, ok := i.(map[string]interface{}); ok {
			var objectId = m["object_id"].(string)
			var objectType = m["object_type"].(string)
			var operation = m["operation"].(string)
			model := &models.Privilege{
				ObjectID:  &objectId,
				ObjectType: &objectType,
				Operation: &operation,
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

func GetPropFromUptimeSteps(d *schema.ResourceData, key string) (t []*models.UptimeWebCheckStep) {
	if r, ok := d.GetOk(key); ok {
		return getPropFromUptimeStepsInterface(r)
	}
	return
}

func getPropFromUptimeStepsInterface(r interface{}) (t []*models.UptimeWebCheckStep) {
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
            
			model := &models.UptimeWebCheckStep{
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
				UseDefaultRoot: useDefaultRoot,
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

// ConformToSchema converts flattened sub-resource data (a slice of
// map[string]interface{} or *map[string]interface{}, as returned by the generated
// Set<Model>SubResourceData functions) into a value that ResourceData.Set accepts
// for a list attribute whose element is a *schema.Resource with the given attributes.
//
// ResourceData.Set rejects the entire list when any nested value has an unexpected
// type, so loosely typed API fields are coerced:
//   - TypeMap attributes with string elements keep string values and JSON-encode
//     any other non-nil value (numbers, booleans, nested objects).
//   - Nested list/set attributes whose element is a *schema.Resource are converted
//     recursively.
//   - Keys without a matching attribute schema are dropped.
//
// All other values are passed through unchanged.
func ConformToSchema(items interface{}, attrs map[string]*schema.Schema) []interface{} {
	out := []interface{}{}
	rv := reflect.ValueOf(items)
	if !rv.IsValid() || (rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array) {
		return out
	}
	for i := 0; i < rv.Len(); i++ {
		item, ok := asStringKeyedMap(rv.Index(i).Interface())
		if !ok {
			continue
		}
		conformed := make(map[string]interface{}, len(item))
		for key, value := range item {
			if attr, known := attrs[key]; known {
				conformed[key] = conformValueToSchema(value, attr)
			}
		}
		out = append(out, conformed)
	}
	return out
}

// conformValueToSchema coerces a single attribute value for ConformToSchema.
func conformValueToSchema(value interface{}, attr *schema.Schema) interface{} {
	switch attr.Type {
	case schema.TypeMap:
		if elem, ok := attr.Elem.(*schema.Schema); ok && elem.Type != schema.TypeString {
			return value
		}
		m, ok := asStringKeyedMap(value)
		if !ok {
			return nil
		}
		result := make(map[string]interface{}, len(m))
		for k, v := range m {
			switch typed := v.(type) {
			case nil:
			case string:
				result[k] = typed
			default:
				if encoded, err := json.Marshal(typed); err == nil {
					result[k] = string(encoded)
				} else {
					result[k] = fmt.Sprintf("%v", typed)
				}
			}
		}
		return result
	case schema.TypeList, schema.TypeSet:
		if nested, ok := attr.Elem.(*schema.Resource); ok {
			return ConformToSchema(value, nested.SchemaMap())
		}
	}
	return value
}

// asStringKeyedMap returns v as a map[string]interface{} when v is a non-nil map
// (or pointer to a map) with string keys.
func asStringKeyedMap(v interface{}) (map[string]interface{}, bool) {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return nil, false
		}
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Map || rv.IsNil() || rv.Type().Key().Kind() != reflect.String {
		return nil, false
	}
	out := make(map[string]interface{}, rv.Len())
	iter := rv.MapRange()
	for iter.Next() {
		out[iter.Key().String()] = iter.Value().Interface()
	}
	return out, true
}

// FetchAllPages drives the generic offset/size pagination loop shared by
// every plural ("list") data source reader, typed and polymorphic alike.
//
// fetchPage is called with increasing offsets; it must fetch exactly one
// page and report how many items that page contained plus the API's
// reported running total. FetchAllPages itself never touches the fetched
// items directly -- callers accumulate them (e.g. into a slice) from within
// fetchPage's closure -- so this helper has no model-specific type.
//
// Paging stops as soon as:
//   - fetchPage returns an error (propagated to the caller), or
//   - a page reports zero items (nothing left to fetch), or
//   - the running count of fetched items reaches total, or
//   - maxIterations pages have been fetched, a safety cap against a
//     misbehaving total/items response that would otherwise loop forever.
//
// The offset passed to fetchPage advances by the number of items the
// previous page actually contained rather than by pageSize, so a server
// that caps the effective page size below what was requested is still
// paged through correctly instead of skipping or re-fetching records.
func FetchAllPages(pageSize int32, fetchPage func(offset, size int32) (items int, total int32, err error)) error {
	const maxIterations = 1000

	var fetched int32
	offset := int32(0)
	for i := 0; i < maxIterations; i++ {
		items, total, err := fetchPage(offset, pageSize)
		if err != nil {
			return err
		}

		fetched += int32(items)
		if items == 0 || fetched >= total {
			return nil
		}

		offset += int32(items)
	}

	return nil
}

