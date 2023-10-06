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
}
```

## Argument Reference

The following arguments are **required**:
* `name` - The name of the group

The following arguments are **optional**:
* `description` - The description of the group
* `disable_alerting` - true: alerting is disabled for the websites in the group
false: alerting is enabled for the websites in the group
If stopMonitoring=true, then alerting will also by default be disabled for the websites in the group
* `parent_id` - The Id of the parent group. If parentId=1 then the group exists under the root  group
* `properties` - 
  + `name` - The name of a property (required)
  + `value` - The value of a property (required)
* `stop_monitoring` - true: monitoring is disabled for the websites in the group
false: monitoring is enabled for the websites in the group
If stopMonitoring=true, then alerting will also by default be disabled for the websites in the group

## Import

website groups can be imported using their website group ID or name
```
$ terraform import logicmonitor_website_group.my_website_group 66
$ terraform import logicmonitor_website_group.my_website_group Amazon Website Checks
```