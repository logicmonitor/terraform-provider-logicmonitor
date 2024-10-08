---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_website"
sidebar_current: "docs-logicmonitor-resource-website"
description: |-
  Provides a LogicMonitor website resource. This can be used to create and manage LogicMonitor websites.
---

# logicmonitor_website

Provides a LogicMonitor website resource. This can be used to create and manage LogicMonitor websites.

## Example Usage
```hcl
# Create a LogicMonitor website
resource "logicmonitor_website" "my_website" {
  alert_expr = "\u003c 200 100 50"
  description = "Monitor Ebay site response times"
  disable_alerting = 
  domain = "www.ebay.com"
  global_sm_alert_cond = 0
  group_id = 1
  host = ""
  ignore_s_s_l = true
  individual_alert_level = "warn"
  individual_sm_alert_enable = false
  is_internal = false
  name = "Ebay"
  overall_alert_level = "warn"
  polling_interval = 5
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
  template = 
  test_location = [
    {
      all = false
      collector_ids = [1, 2, 3]
      collectors = null
      smg_ids = [85, 90]
    }
  ]
  transition = 1
  trigger_s_s_l_expiration_alert = false
  trigger_s_s_l_status_alert = false
  type = "webcheck"
  use_default_alert_setting = true
  use_default_location_setting = false
  user_permission = ""
}
```

## Argument Reference

The following arguments are **required**:
* `name` - The name of the website
   (string)
* `type` - The type of the website. Acceptable values are: pingcheck, webcheck
   (string)

The following arguments are **optional**:
* `alert_expr` - The threshold (in days) for triggering SSL certification alerts (string)
* `description` - The description of the website (string)
* `disable_alerting` - true: alerting is disabled for the website
false: alerting is enabled for the website
If stopMonitoring=true, then alerting will also by default be disabled for the website (bool)
* `domain` - Required for type=webcheck , The domain of the service. This is the base URL of the service (string)
* `global_sm_alert_cond` - The number of test locations that checks must fail at to trigger an alert, where the alert triggered will be consistent with the value of overallAlertLevel. Possible values and corresponding number of Site Monitor locations are
0 : all
1 : half
2 : more than one
3 : any (int32)
* `group_id` - The id of the group the website is in (int32)
* `host` - The URL to check, without the scheme or protocol (e.g http or https)
E.g. if the URL is "http://www.google.com, then the host="www.google.com" (string)
* `ignore_s_s_l` - Whether or not SSL should be ignored, the default value is true (bool)
* `individual_alert_level` - warn | error | critical
The level of alert to trigger if the website fails a check from an individual test location (string)
* `individual_sm_alert_enable` - true: an alert will be triggered if a check fails from an individual test location
false: an alert will not be triggered if a check fails from an individual test location (bool)
* `is_internal` - Whether or not the website is internal (bool)
* `overall_alert_level` - warn | error | critical
The level of alert to trigger if the website fails the number of checks specified by transition from the test locations specified by globalSmAlertCond (string)
* `polling_interval` - 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10
The polling interval for the website, in units of minutes. This value indicates how often the website is checked. The minimum is 1 minute, and the maximum is 10 minutes (int32)
* `schema` - The scheme or protocol associated with the URL to check. Acceptable values are: http, https (string)
* `steps` - Required for type=webcheck , An object comprising one or more steps, see the table below for the properties included in each step ([]*WebCheckStep)
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
* `stop_monitoring` - true: monitoring is disabled for the website
false: monitoring is enabled for the website
If stopMonitoring=true, then alerting will also by default be disabled for the website (bool)
* `template` - The website template (interface{})
* `test_location` - The locations from which the website is monitored. If the website is internal, this field should include Collectors. If Non-Internal, possible test locations are:
1 : US - LA
2 : US - DC
3 : US - SF
4 : Europe - Dublin
5 : Asia - Singapore
6 : Australia - Sydney (WebsiteLocation)
  + `all` - (true | false) Indicates that the service will be monitored from all checkpoint locations
  + `collectorIds` - indicates that the service will be monitored from checkpoint locations 1, 2 and 3
  + `collectors` - Need to pass 'null' value
  + `smgIds` - indicates that the service will be monitored by Collectors 85 and 90
* `transition` - 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 | 30 | 60
The number of checks that must fail before an alert is triggered (int32)
* `trigger_s_s_l_expiration_alert` - Whether or not SSL expiration alerts should be triggered (bool)
* `trigger_s_s_l_status_alert` - Whether or not SSL status alerts should be triggered (bool)
* `use_default_alert_setting` - true: The alert settings configured in the website Default Settings will be used
false: Service Default Settings will not be used, and you will need to specify individualSMAlertEnable, individualAlertLevel, globalSmAlertConf, overallAlertLevel and pollingInterval (bool)
* `use_default_location_setting` - true: The checkpoint locations configured in the website Default Settings will be used
false: The checkpoint locations specified in the testLocation will be used (bool)
* `user_permission` - write | read | ack. The permission level of the user that made the API request (string)

## Import

websites can be imported using their website ID or name
```
$ terraform import logicmonitor_website.my_website 66
$ terraform import logicmonitor_website.my_website Ebay
```