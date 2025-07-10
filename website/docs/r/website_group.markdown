---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_website_group"
sidebar_current: "docs-logicmonitor-resource-website-group"
description: |-
  Provides a LogicMonitor website group resource. This can be used to create and manage LogicMonitor website groups.
---

# logicmonitor_website_group

Provides a LogicMonitor website group resource. This can be used to create and manage LogicMonitor website groups.

## Example Usage
```hcl
# Create a LogicMonitor website group
resource "logicmonitor_website_group" "my_website_group" {
  description = "Amazon web and ping checks"
  disable_alerting = false
  name = "Amazon Website Checks"
  parent_id = 1
  properties = [
   {
    name : "addr",
    value : "127.0.0.1"
   }  
  ]
  stop_monitoring = false
  test_location = [
    {
      all = false
      collector_ids = [1, 2, 3]
      collectors = null
      smg_ids = [85, 90]
    }
  ]
}
```

## Argument Reference

The following arguments are **required**:
* `name` - The name of the group
   (string)

The following arguments are **optional**:
* `description` - The description of the group (string)
* `disable_alerting` - true: alerting is disabled for the websites in the group
false: alerting is enabled for the websites in the group
If stopMonitoring=true, then alerting will also by default be disabled for the websites in the group (bool)
* `parent_id` - The Id of the parent group. If parentId=1 then the group exists under the root  group (int32)
* `properties` -  ([]*NameAndValue)
    Provide custom properties in alphabetical order based on their property name.
  + `name` - The name of a property (required)
  + `value` - The value of a property (required)
* `stop_monitoring` - true: monitoring is disabled for the websites in the group
false: monitoring is enabled for the websites in the group
If stopMonitoring=true, then alerting will also by default be disabled for the websites in the group (bool)
* `test_location` - An object that indicates the websites locations.
eg. {'all': false, smgId:[1,2,3], collectorIds:[14,16]} (WebsiteLocation)
  + `all` - (true | false) Indicates that the service will be monitored from all checkpoint locations
  + `collectorIds` - indicates that the service will be monitored from checkpoint locations 1, 2 and 3
  + `collectors` - Need to pass 'null' value
  + `smgIds` - indicates that the service will be monitored by Collectors 85 and 90

## Import

website groups can be imported using their website group ID or name
```
$ terraform import logicmonitor_website_group.my_website_group 66
$ terraform import logicmonitor_website_group.my_website_group Amazon Website Checks
```