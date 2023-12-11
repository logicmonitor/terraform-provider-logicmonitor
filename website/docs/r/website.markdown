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
  description = "Monitor Ebay site response times"
  disable_alerting = 
  domain = "www.ebay.com"
  global_sm_alert_cond = 0
  group_id = 1
  host = ""
  individual_alert_level = "warn"
  individual_sm_alert_enable = false
  is_internal = false
  name = "Ebay"
  overall_alert_level = "warn"
  polling_interval = 5
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
    require_auth = false
    path = "string"
    auth = [{ 
      password = "string"
      type = "basic"
      domain = "string"
      user_name = "string"
    }]
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
  transition = 1
  type = "webcheck"
  use_default_alert_setting = true
  use_default_location_setting = false
  user_permission = ""
}
```

## Argument Reference

The following arguments are **required**:
* `name` - The name of the website
* `type` - The type of the website. Acceptable values are: pingcheck, webcheck

The following arguments are **optional**:
* `description` - The description of the website
* `disable_alerting` - true: alerting is disabled for the website
false: alerting is enabled for the website
If stopMonitoring=true, then alerting will also by default be disabled for the website
* `domain` - Required for type=webcheck , The domain of the service. This is the base URL of the service
* `global_sm_alert_cond` - The number of test locations that checks must fail at to trigger an alert, where the alert triggered will be consistent with the value of overallAlertLevel. Possible values and corresponding number of Site Monitor locations are
0 : all
1 : half
2 : more than one
3 : any
* `group_id` - The id of the group the website is in
* `host` - The URL to check, without the scheme or protocol (e.g http or https)
E.g. if the URL is "http://www.google.com, then the host="www.google.com"
* `individual_alert_level` - warn | error | critical
The level of alert to trigger if the website fails a check from an individual test location
* `individual_sm_alert_enable` - true: an alert will be triggered if a check fails from an individual test location
false: an alert will not be triggered if a check fails from an individual test location
* `is_internal` - Whether or not the website is internal
* `overall_alert_level` - warn | error | critical
The level of alert to trigger if the website fails the number of checks specified by transition from the test locations specified by globalSmAlertCond
* `polling_interval` - 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10
The polling interval for the website, in units of minutes. This value indicates how often the website is checked. The minimum is 1 minute, and the maximum is 10 minutes
* `steps` - Required for type=webcheck , An object comprising one or more steps, see the table below for the properties included in each step
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
If stopMonitoring=true, then alerting will also by default be disabled for the website
* `template` - The website template
* `transition` - 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 | 30 | 60
The number of checks that must fail before an alert is triggered
* `use_default_alert_setting` - true: The alert settings configured in the website Default Settings will be used
false: Service Default Settings will not be used, and you will need to specify individualSMAlertEnable, individualAlertLevel, globalSmAlertConf, overallAlertLevel and pollingInterval
* `use_default_location_setting` - true: The checkpoint locations configured in the website Default Settings will be used
false: The checkpoint locations specified in the testLocation will be used
* `user_permission` - write | read | ack. The permission level of the user that made the API request

## Import

websites can be imported using their website ID or name
```
$ terraform import logicmonitor_website.my_website 66
$ terraform import logicmonitor_website.my_website Ebay
```