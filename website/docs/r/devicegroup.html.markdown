---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_devicegroup"
sidebar_current: "docs-logicmonitor-resource-devicegroup"
description: |-
  Provides a LogicMonitor device group resource. This can be used to create and manage LogicMonitor device groups
---

# logicmonitor_devicegroup

Provides a LogicMonitor device group resource. This can be used to create and manage LogicMonitor device groups

## Example Usage

```hcl
# Create a new LogicMonitor device group
resource "logicmonitor_devicegroup" "group" {
    name = "NewGroup24"
    description = "new test group"
    properties {
     "group" = "test"
     "system.categories" = "a,b,c,d"
    }
}
```

```hcl
# Create a new LogicMonitor dynamic device group
resource "logicmonitor_devicegroup" "group1" {
    name = "NewGroup25"
    applies_to = "system.displayname =~ \"Prod\""
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of device group
* `description` - (Optional) Description of device group
* `parent_id` - (Optional) The id of the parent group for this device group (the root device group has an Id of 1)
* `applies_to` - (Optional) The Applies to custom query for this group. Setting this field will make this a dynamic group.
* `disable_alerting` - (Optional) Indicates whether alerting is disabled (true) or enabled (false) for this device group
* `properties` - (Optional) The properties associated with this device group. Any string value pair will work (see example).
