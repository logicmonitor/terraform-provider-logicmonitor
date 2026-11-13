---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_action_rule"
sidebar_current: "docs-logicmonitor-resource-action-rule"
description: |-
  Provides a LogicMonitor action rule resource. This can be used to create and manage LogicMonitor action rules.
---

# logicmonitor_action_rule
Provides a LogicMonitor action rule resource. This can be used to create and manage LogicMonitor action rules.

## Example Usage
```hcl
# Create a LogicMonitor action rule
resource "logicmonitor_action_rule" "my_action_rule" {
  name            = "parag_action_rule pravesh"
  action_chain_id = 216
  level_str       = "All"
  datasource      = "Pingds05"
  instance        = "*"
  datapoint       = "*"
  device_groups   = ["Parag_device_group"]
  devices         = ["*"]
}
```

## Argument Reference

The following arguments are **required**:
* `action_chain_id` - The action chain id for the action rule
   (int32)
* `device_groups` - The device groups for which the action rule is configured to match
   ([]string)
* `level_str` - The severity levels for which the action rule is configured to match.
   (string)
* `name` - The name of the action rule
   (string)

The following arguments are **optional**:
* `datapoint` - The datapoint for which the action rule is configured to match (string)
* `datasource` - The datasource for which the action rule is configured to match (string)
* `devices` - The devices for which the action rule is configured to match ([]string)
* `enabled` - Whether action Rule is enabled or not (bool)
* `instance` - The instance for which the action rule is configured to match (string)
* `resource_properties` - The resource property filters list ([]*DeviceProperty)

## Import

action rules can be imported using their action rule ID or name
```
$ terraform import logicmonitor_action_rule.my_action_rule 66
$ terraform import logicmonitor_action_rule.my_action_rule action Rule1
```