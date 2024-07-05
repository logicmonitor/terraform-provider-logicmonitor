---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_alert_rule"
sidebar_current: "docs-logicmonitor-resource-alert-rule"
description: |-
  Provides a LogicMonitor alert rule resource. This can be used to create and manage LogicMonitor alert rules.
---

# logicmonitor_alert_rule

Provides a LogicMonitor alert rule resource. This can be used to create and manage LogicMonitor alert rules.

## Example Usage
```hcl
# Create a LogicMonitor alert rule
resource "logicmonitor_alert_rule" "my_alert_rule" {
  datapoint = "*"
  datasource = "Port-"
  devices = [
    "Cisco Router"
  ]
  device_groups = [
  "Devices by Type"
  ]
  escalating_chain_id = 12
  escalation_interval = 15
  instance = "*"
  level_str = "Warn"
  name = "Warning"
  priority = 100
  suppress_alert_ack_sdt = true
  suppress_alert_clear = true
}
```

## Argument Reference

The following arguments are **required**:
* `datapoint` - The datapoint the alert rule is configured to match
   (string)
* `datasource` - The datasource the alert rule is configured to match
   (string)
* `device_groups` - The device groups and service groups the alert rule is configured to match
   ([]string)
* `devices` - The device names and service names the alert rule is configured to match
   ([]string)
* `escalating_chain_id` - The id of the escalation chain associated with the alert rule
   (int32)
* `instance` - The instance the alert rule is configured to match
   (string)
* `name` - The name of the alert rule
   (string)
* `priority` - The priority associated with the alert rule
   (int32)

The following arguments are **optional**:
* `escalation_interval` - The escalation interval associated with the alert rule, in minutes (int32)
* `level_str` - The alert severity levels the alert rule is configured to match. Acceptable values are: All, Warn, Error, Critical (string)
* `suppress_alert_ack_sdt` - Whether or not status notifications for acknowledgements and SDTs should be sent to the alert rule (bool)
* `suppress_alert_clear` - Whether or not alert clear notifications should be sent to the alert rule (bool)

## Import

alert rules can be imported using their alert rule ID or name
```
$ terraform import logicmonitor_alert_rule.my_alert_rule 66
$ terraform import logicmonitor_alert_rule.my_alert_rule Warning
```