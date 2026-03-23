---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_device"
sidebar_current: "docs-logicmonitor-resource-device"
description: |-
  Provides a LogicMonitor device resource. This can be used to create and manage LogicMonitor devices.
---

# logicmonitor_device

Provides a LogicMonitor device resource. This can be used to create and manage LogicMonitor devices.

## Example Usage
```hcl
# Create a LogicMonitor device
resource "logicmonitor_device" "my_device" {
  alert_expr = "\u003c 200 100 50"
  auto_balanced_collector_group_id = 0
  contains_multi_value = 
  current_collector_id = 1
  current_log_collector_id = 1
  custom_properties = [
		{
			name = "addr"
      		value = "127.0.0.1"
		},
		{
			name = "host"
      		value = "localhost"
		},
    {
      name  = "system.categories"
      value = "" 
    }
	]
  description = "This is a Cisco Router"
  device_type = 0
  disable_alerting = true
  display_name = "Cisco Router"
  domain = "www.ebay.com"
  enable_netflow = true
  global_sm_alert_cond = 0
  host = ""
  host_group_ids = "16,4,3"
  ignore_s_s_l = true
  individual_alert_level = "warn"
  individual_sm_alert_enable = false
  is_internal = false
  is_preferred_log_collector_configured = true
  link = "www.ciscorouter.com"
  log_collector_id = 1
  model = "websiteDevice"
  name = "collector.localhost"
  need_stc_grp_and_sorted_c_p = true
  netflow_collector_id = 1
  op = ""
  overall_alert_level = "warn"
  page_load_alert_time_in_m_s = 30000
  percent_pkts_not_receive_in_time = 
  polling_interval = 5
  preferred_collector_id = 2
  properties = []
  related_device_id = -1
  schema = "https"
  steps = [
   {
    schema = "https"
    resp_type = "config"
    timeout = 0
    match_type = "plain"
    description = "string"
    use_default_root = true
    http_method = "GET"
    enable = true
    http_version = "1.1"
    http_headers = "X-Version:3"
    follow_redirection = true
    post_data_edit_type = "raw"
    name = "string"
    req_type = "config"
    fullpage_load = false
    require_auth = var.require_auth
    path = "string"
    auth = var.require_auth ? [{
    password  = "string"
    type      = "basic"
    domain    = "string"
    user_name = "string"
    }] : []
    keyword = "DEVWRT-SANRT-JOB1-9127"
    http_body = "string"
    resp_script = "string"
    req_script = "string"
    label = "string"
    url = "/rest/version"
    type = "string"
    invert_match = false
    status_code = "200"
   }
  ]
  stop_monitoring = 
  test_location = [
    {
      all = false
      collector_ids = [1, 2, 3]
      smg_ids = [85, 90]
    }
  ]
  timeout_in_m_s_pkts_not_receive = 
  transition = 1
  trigger_s_s_l_expiration_alert = false
  trigger_s_s_l_status_alert = false
  type = "webcheck"
  use_default_alert_setting = true
  use_default_location_setting = false
}
```

## Argument Reference

The following arguments are **required**:
* `display_name` - The display name of the device
   (string)
* `name` - The host name or IP address of the device
   (string)
* `preferred_collector_id` - The Id of the preferred collector assigned to monitor the device
   (int32)

The following arguments are **optional**:
* `alert_expr` - The threshold (in days) for triggering SSL certification alerts (string)
* `auto_balanced_collector_group_id` - The Auto Balanced Collector Group id. 0 means not monitored by ABCG (int32)
* `checkpoints` - The checkpoints from which the website device device is monitored. This object should reference each location specified in testLocation in addition to an 'Overall' checkpoint ([]*WebsiteCheckPoint)
* `contains_multi_value` - request contains multi value field (bool)
* `current_collector_id` - The id of the collector currently monitoring the device and discovering instances (int32)
* `current_log_collector_id` - The id of the Log collector currently collecting logs. (int32)
* `custom_properties` - Any non-system properties (aside from system.categories) defined for this device ([]*NameAndValue)
    Provide custom properties in alphabetical order based on their property name.
  + `name` - The name of a property (required)
  + `value` - The value of a property (required)
* `description` - The device description (string)
* `device_type` - The type of device: 0 indicates a regular device, 2 indicates an AWS device, 4 indicates an Azure device,18 indicates an uptime webcheck,19 indicates an uptime pingcheck (int32)
* `disable_alerting` - Indicates whether alerting is disabled (true) or enabled (false) for this device (bool)
* `domain` - Required for type=webcheck, the domain of the service. This is the base URL of the service (string)
* `enable_netflow` - Indicates whether Netflow is enabled (true) or disabled (false) for the device (bool)
* `global_sm_alert_cond` - The number of test locations that checks must fail at to trigger an alert, where the alert triggered will be consistent with the value of overallAlertLevel. Possible values and corresponding number of Site Monitor locations are
0 : all
1 : half
2 : more than one
3 : any (int32)
* `group_ids` - The ids of the groups the website device is in ([]int32)
* `host` - The URL to check, without the scheme or protocol (e.g http or https)
E.g. if the URL is "http://www.google.com", then the host="www.google.com" (string)
* `host_group_ids` - The Id(s) of the groups the device is in, where multiple group ids are comma separated (string)
* `ignore_s_s_l` - Whether or not SSL should be ignored, the default value is true (bool)
* `individual_alert_level` - The values can be warn|error|critical
The level of alert to trigger if the website device device fails a check from an individual test location (string)
* `individual_sm_alert_enable` - The values can be true|false where
true: an alert will be triggered if a check fails from an individual test location
false: an alert will not be triggered if a check fails from an individual test location (bool)
* `is_internal` - Whether or not the website device is internal (bool)
* `is_preferred_log_collector_configured` - Indicates whether Preferred Log Collector is configured  (true) or not (false) for the device (bool)
* `link` - The URL link associated with the device (string)
* `log_collector_id` - The Id of the netflow collector associated with the device (int32)
* `model` - The model of the website device, which is determined by the type (string)
* `need_stc_grp_and_sorted_c_p` - Indicates whether Static group and sorted Custom properties are needed (bool)
* `netflow_collector_id` - The Id of the netflow collector associated with the device (int32)
* `op` - whether to use AND or OR for device matching (string)
* `overall_alert_level` - The values can be warn|error|critical
The level of alert to trigger if the website device device fails the number of checks specified by transition from the test locations specified by globalSmAlertCond (string)
* `page_load_alert_time_in_m_s` - The time in milliseconds that the page must load within for each step to avoid triggering an alert (int32)
* `percent_pkts_not_receive_in_time` - The percentage of packets that should be returned in the time period specified by timeoutInMSPktsNotReceive for each ping check (int32), the default value is 80
* `polling_interval` - The values can be 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10
The polling interval for the website device device, in units of minutes. This value indicates how often the website device device is checked. The minimum is 1 minute, and the maximum is 10 minutes (int32)
* `preferred_collector_group_id` - The id of the Collector Group associated with the device's preferred collector (int32), It can be 0 for auto balanced collector group .
* `properties` - The properties associated with the website device ([]*NameAndValue)
* `related_device_id` - The Id of the AWS EC2 instance related to this device, if one exists in the LogicMonitor account. This value defaults to -1, which indicates that there are no related devices (int32)
* `resource_ids` - Any non-system properties (aside from system.categories) defined for this device ([]*NameAndValue)
* `schema` - The scheme or protocol associated with the URL to check. Acceptable values are: http, https (string)
* `steps` - Required for type=webcheck, an object comprising one or more steps, see the table below for the properties included in each step ([]*UptimeWebCheckStep)
  + `schema` -  HTTP schema
  + `matchType` - Body match type(plain | JSON | XML | Glob Expression | Multi-line key-value pairs)
  + `description` - The description of the Step
  + `httpVersion` - 1.1 | 1\nSpecifies HTTP version
  + `respType` - Plain Text/String | Glob expression | JSON | XML | Multi line key value pair\nStep Response Type
  + `reqType` - script | config\nStep Request Type
  + `followRedirection` - true | false\nSpecifies whether to follow redirection or not
  + `httpMehod` - GET | HEAD | POST\nSpecifies the type of HTTP method
  + `enable` - true | false\nSpecifies whether to enable step or not
  + `name` - The name of the Step
  + `timeout` - Request timeout measured in seconds
  + `useDefaultRoot` - true | falseCheck if using the default root
  + `postDataEditType` - Raw | Formatted Data\nSpecifies POST data type
  + `fullpageLoad` - true | false\nChecks if full page should be loaded or not
  + `requireAuth` - true | false\nChecks if authorization required or not
     define this variable in .tf configuration -
     variable "require_auth" {
     description = "Whether authentication is required"
     type        = bool
     default     = true
    }
  + `path` - Path for JSON, XPATH
  + `auth`- Authorization Information
    + `password` - NTLM authentication password
    + `type` - Authentication type
    + `userName` - NTLM  authentication userName
    + `domain` - NTLM domain (if auth type is NTLM)
  + `httpHeaders` - HTTP header
  + `httpBody` - HTTP Body
  + `keyword` - Keyword that matches the body
  + `respScript` - The Step Response Script
  + `reqScript` - The Request Script
  + `label` - The Label of the Step
  + `url` - The URL of service step
  + `invertMatch` - true | false\nChecks if invert matches or not
  + `type` - script | config\nThe type of service step
  + `statusCode` - The expected status code
* `stop_monitoring` - The values can be true|false where
true: monitoring is disabled for the website device device
false: monitoring is enabled for the website device device
If stopMonitoring=true, then alerting will also be disabled by default for the website device device (bool)
* `synthetics_collector_ids` - The list of ids of the collectors currently monitoring the resource and discovering instances ([]int32)
* `template` -  (JSONObject)
* `test_location` -  (WebsiteLocation)
  + `all` - (true | false) Indicates that the service will be monitored from all checkpoint locations
  + `collectorIds` - indicates that the service will be monitored from checkpoint locations 1, 2 and 3
  + `collectors` - Need to pass 'null' value
  + `smgIds` - indicates that the service will be monitored by Collectors 85 and 90
* `timeout_in_m_s_pkts_not_receive` - The time period that the percentage of packets specified by percentPktsNotReceiveInTime must be returned in for each ping check (int32), the default value is 500
* `transition` - The values can be 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 | 30 | 60
The number of checks that must fail before an alert is triggered (int32)
* `trigger_s_s_l_expiration_alert` - Whether or not SSL expiration alerts should be triggered (bool)
* `trigger_s_s_l_status_alert` - Whether or not SSL status alerts should be triggered (bool)
* `type` - Specifies the v3 Uptime Device request type. Supported values:
 webcheck - for Uptime Web Check devices
 pingcheck - for Uptime Ping Check devices blank/empty value- for regular devices (string)
* `use_default_alert_setting` - The values can be true|false where
true: The alert settings configured in the website device device Default Settings will be used
false: Service Default Settings will not be used, and you will need to specify individualSMAlertEnable, individualAlertLevel, globalSmAlertConf, overallAlertLevel and pollingInterval (bool)
* `use_default_location_setting` - The values can be true|false where
true: The checkpoint locations configured in the website device device Default Settings will be used
false: The checkpoint locations specified in the testLocation will be used (bool)

## Import

devices can be imported using their device ID or name
```
$ terraform import logicmonitor_device.my_device 66
$ terraform import logicmonitor_device.my_device collector.localhost
```