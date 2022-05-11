---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_dashboard"
sidebar_current: "docs-logicmonitor-resource-dashboard"
description: |-
  Provides a LogicMonitor dashboard resource. This can be used to create and manage LogicMonitor dashboards.
---

# logicmonitor_dashboard

Provides a LogicMonitor dashboard resource. This can be used to create and manage LogicMonitor dashboards.

## Example Usage
```hcl
# Create a LogicMonitor dashboard
resource "logicmonitor_dashboard" "my_dashboard" {
  description = "Windows Servers Performance"
  group_id = 1
  group_name = "Server Dashboard"
  name = "Default Device Group"
  owner = ""
  sharable = true
  template = 
  widgets_config = 
}
```

## Argument Reference

The following arguments are **required**:
* `name` - The name of the dashboard

The following arguments are **optional**:
* `description` - The description of the dashboard
* `group_id` - The id of the group the dashboard belongs to
* `group_name` - The name of group where created dashboard will reside
* `owner` - This field will be empty unless the dashboard is a private dashboard, in which case the owner will be listed
* `sharable` - Whether or not the dashboard is sharable. This value will always be true unless the dashboard is a private dashboard
* `template` - The template which is used for import dashboard
* `widget_tokens` - If useDynamicWidget=true, this field must at least contain tokens defaultDeviceGroup and defaultServiceGroup
  + `name` (required)
  + `value` (required)
  + `type` (optional) - Need to pass 'null' value
  + `value` (optional)  - Need to pass 'null' value
* `widgets_config` - Information about widget configuration used by the UI

## Import

dashboards can be imported using their dashboard ID or name
```
$ terraform import logicmonitor_dashboard.my_dashboard 66
$ terraform import logicmonitor_dashboard.my_dashboard Default Device Group
```