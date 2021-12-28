---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_collector"
sidebar_current: "docs-logicmonitor-resource-collector"
description: |-
  Provides a LogicMonitor collector resource. This can be used to create and manage LogicMonitor collectors.
---

# logicmonitor_collector

Provides a LogicMonitor collector resource. This can be used to create and manage LogicMonitor collectors.

## Example Usage
```hcl
# Create a LogicMonitor collector
resource "logicmonitor_collector" "my_collector" {
  arch = "linux64"
  automatic_upgrade_info {
    day_of_week = "MON"
    description = "regular MGD updates"
    hour = 15
    minute = 0
    occurrence = "Any"
    timezone = "America/Los_Angeles"
    version = "MGD"
  }
  backup_agent_id = 75
  build = "30001"
  collector_group_id = 10
  collector_size = "small"
  company = "lm"
  custom_properties = [
		{
			name = "addr"
      		value = "127.0.0.1"
		},
		{
			name = "host"
      		value = "localhost"
		}
	]
  description = "Linux Collector"
  ea = false
  enable_fail_back = true
  enable_fail_over_on_collector_device = true
  escalating_chain_id = 80
  monitor_others = false
  need_auto_create_collector_device = true
  number_of_instances = 0
  one_time_downgrade_info {
    description = "regular MGD updates"
    majorVersion = 27
    minorVersion = 1
    startEpoch = 1534888740
    timezone" = "America/Los_Angeles"
  }
  "one_time_upgrade_info {
    description = "regular MGD updates"
    majorVersion = 27
    minorVersion = 1
    startEpoch = 1534888740
    timezone" = "America/Los_Angeles"
  }
  resend_ival = 15
  specified_collector_device_group_id = 0
  suppress_alert_clear = true
}

# You can view the collector installation URL & commands via the installer_url_cmds variable
output "myInstallerURL" {
  value = logicmonitor_collector.myCollector.installer_url_cmds
}
```

## Argument Reference

**Note the following**:
* Successfully adding a collector returns the output variable `installer_url_cmds` that contains the installation URL, as well as the associated cURL commands. You can visit [Terraform's documentation](https://www.terraform.io/docs/language/values/outputs.html) to learn how to access output values.
* The argument `company` (i.e. the user's portal name) is required if the environment variable `LM_COMPANY` is not set.

The following arguments are **optional**:
* `arch` - The collector architecture (Windows | Linux platform followed by 32 | 64 bit), the default value is linux64
* `automatic_upgrade_info` - The details of the Collector's automatic upgrade schedule, if one exists
  + `dayOfWeek` (required) - Options include `SUN`, `MON`, `TUE`, `WED`, `THU`, `FRI`, `SAT`
  + `description`
  + `hour` (required)
  + `minute` (required)
  + `occurrence` (required) - Options inlcude `FIRST`, `SECOND`, `THIRD`, `FOURTH`, `ANY`
  + `timezone`
  + `version` (required) - Options include `ED`, `GD`, `MGD`
* `backup_agent_id` - The Id of the backup Collector assigned to the Collector
* `build` - The Collector version
* `collector_group_id` - The Id of the group the Collector is in
* `collector_size` - The size of the collector, the default value is medium
* `company` - The user's company (portal) name, this field is required if the environment variable LM_COMPANY is not set
* `custom_properties` - The custom properties defined for the Collector
  + `name` - The name of a property (required)
  + `value` - The value of a property (required)
* `description` - The Collector's description
* `ea` - Whether the collector is in EA version
* `enable_fail_back` - Whether or not automatic failback is enabled for the Collector, the default value is true
* `enable_fail_over_on_collector_device` - Whether or not the device the Collector is installed on is enabled for fail over
* `escalating_chain_id` - The Id of the escalation chain associated with this Collector
* `monitor_others` - Check if we shall monitor using local account (for windows), the default value is true
* `need_auto_create_collector_device` - Whether to create a collector device when instance collector, the default value is true
* `number_of_instances` - The number of instances that are monitored by this collector
* `onetime_downgrade_info` - The details of the Collector's automatic downgrade schedule, if one exists
  + `description`
  + `major_version` (required)
  + `minor_version` (required)
  + `start_epoch` (required)
  + `timezone`
* `onetime_upgrade_info` - The details of the Collector's one time upgrade, if one has been scheduled
  + `description`
  + `major_version` (required)
  + `minor_version` (required)
  + `start_epoch` (required)
  + `timezone`
* `resend_ival` - The interval, in minutes, after which alert notifications for the Collector will be resent
* `specified_collector_device_group_id` - The collector device group id assigned when creating a new collector device
* `suppress_alert_clear` - Whether alert clear notifications are suppressed for the Collector

## Import

collectors can be imported using their collector ID
```
$ terraform import logicmonitor_collector.my_collector 66
```