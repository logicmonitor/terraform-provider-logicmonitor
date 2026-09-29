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
  template = jsonencode({
    name                 = "test dashboard"
    description          = "Created from a dashboard template"
    widgetsConfigVersion = 2
    type                 = "dashboard"
    version              = 2
    widgetTokens = [
      {
        name  = "defaultDeviceGroup"
        value = "*"
      },
      {
        name  = "defaultServiceGroup"
        value = "*"
      }
    ]
    widgets = [
      {
        position = {
          col   = 1
          row   = 1
          sizex = 6
          sizey = 4
        }
        config = {
          name        = "How to use dashboards"
          description = "HTML widget from template"
          type        = "html"
          theme       = "borderPurple"
          interval    = 5
          timescale   = "day"
          isCustom    = false
          resources = [
            {
              type = "html"
              URL  = "https://www.youtube.com/embed/-WNU4ffumk0"
            }
          ]
          version = 2
        }
      }
    ]
  })  
  widget_tokens = [
    {
      name  = "defaultDeviceGroup"
      value = "*"
      inherit_list = null
      type = null
    },
    {
      name  = "defaultServiceGroup"
      value = "*"
      inherit_list = null
      type = null
    }
  ]
  widgets_config = 
}
```

## Argument Reference

The following arguments are **required**:
* `name` - The name of the dashboard
   (string)

The following arguments are **optional**:
* `description` - The description of the dashboard (string)
* `group_id` - The id of the group the dashboard belongs to (int32)
* `group_name` - The name of group where created dashboard will reside (string)
* `owner` - This field will be empty unless the dashboard is a private dashboard, in which case the owner will be listed (string)
* `sharable` - Whether or not the dashboard is sharable. This value will always be true unless the dashboard is a private dashboard (bool)
* `template` - The template which is used for import dashboard (interface{})
* `widget_tokens` - If useDynamicWidget=true, this field must at least contain tokens defaultDeviceGroup and defaultServiceGroup ([]*WidgetToken)
  + `name` (required)
  + `value` (required)
  + `type` (required) - Need to pass 'null' value
  + `inherit_list` (required)  - Need to pass 'null' value
* `widgets_config` - Information about widget configuration used by the UI (interface{}), this field can remain empty for terraform.

## Import

dashboards can be imported using their dashboard ID or name
```
$ terraform import logicmonitor_dashboard.my_dashboard 66
$ terraform import logicmonitor_dashboard.my_dashboard Default Device Group
```