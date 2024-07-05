---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_dashboard_group"
sidebar_current: "docs-logicmonitor-resource-dashboard-group"
description: |-
  Provides a LogicMonitor dashboard group resource. This can be used to create and manage LogicMonitor dashboard groups.
---

# logicmonitor_dashboard_group

Provides a LogicMonitor dashboard group resource. This can be used to create and manage LogicMonitor dashboard groups.

## Example Usage
```hcl
# Create a LogicMonitor dashboard group
resource "logicmonitor_dashboard_group" "my_dashboard_group" {
  description = "Servers in LA DataCenter"
  name = "LogicMonitor Dashboards"
  parent_id = 1
  template = 
}
```

## Argument Reference

The following arguments are **required**:
* `name` - The name of the dashboard group
   (string)

The following arguments are **optional**:
* `description` - This is a description of the dashboard group (string)
* `parent_id` - The Id of the parent dashboard group (int32)
* `template` - The template which is used for import dashboard group (interface{})
* `widget_tokens` - The tokens assigned at the group level ([]*WidgetToken)
  + `name` (required)
  + `value` (required)
  + `type` (required) - Need to pass 'null' value
  + `inherit_list` (required)  - Need to pass 'null' value

## Import

dashboard groups can be imported using their dashboard group ID or name
```
$ terraform import logicmonitor_dashboard_group.my_dashboard_group 66
$ terraform import logicmonitor_dashboard_group.my_dashboard_group LogicMonitor Dashboards
```