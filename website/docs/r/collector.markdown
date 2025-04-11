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
		},
    {
      name  = "system.categories"
      value = "" 
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
* `arch` - The collector architecture (Windows | Linux platform followed by 32 | 64 bit) (string), the default value is linux64
* `automatic_upgrade_info` - The details of the Collector's automatic upgrade schedule, if one exists (AutomaticUpgradeInfo)
  + `dayOfWeek` (required) - Options include `SUN`, `MON`, `TUE`, `WED`, `THU`, `FRI`, `SAT`
  + `description`
  + `hour` (required)
  + `minute` (required)
  + `occurrence` (required) - Options inlcude `FIRST`, `SECOND`, `THIRD`, `FOURTH`, `ANY`
  + `timezone`
  + `version` (required) - Options include `EA`, `GD`, `MGD`
    Early Access (EA) – EA releases are often the first to debut new functionality. We sometimes release a major feature in batches through EA release. So, EA is not recommended for your entire infrastructure. They occur 9-10 times per year. If there are major bug fixes, we patch EA, and it is referred as EA patch release. A stable EA version is designated as an optional general release (GD).
    Optional General Releases (GD) – GD releases are stable collector updates that may have new features. However, it is not mandatory to update collectors with GD releases. They occur twice a year. If there are major bug fixes, we patch GD, and it is referred as GD patch release. A stable GD version is designated as Required General Release (MGD).
    Required General Releases (MGD) – An MGD is released once a year. When we designate a GD as an MGD, we schedule and announce a date to auto-upgrade collectors to the MGD version. To let customers upgrade collectors as per their convenience, we send communication at least 30 days before the scheduled auto-upgrade date. On the auto-upgrade date, we upgrade only those collectors which are still below the MGD version. Thus, going forward, the MGD becomes the minimum required version. If there are major bug fixes, we patch MGD, and it is referred as MGD patch release.
    For more details please refer: https://www.logicmonitor.com/support/collectors/collector-overview/collector-versions
* `backup_agent_id` - The Id of the backup Collector assigned to the Collector (int32)
* `build` - The Collector version (string)
* `collector_group_id` - The Id of the group the Collector is in (int32)
* `collector_size` - The size of the collector (string), the default value is medium
* `company` - The user's company (portal) name, this field is required if the environment variable LM_COMPANY is not set (string)
* `custom_properties` - The custom properties defined for the Collector ([]*NameAndValue)
    Provide custom properties in alphabetical order based on their property name.
  + `name` - The name of a property (required)
  + `value` - The value of a property (required)
* `description` - The Collector's description (string)
* `ea` - Whether the collector is in EA version (bool)
* `enable_fail_back` - Whether or not automatic failback is enabled for the Collector, the default value is true (bool)
* `enable_fail_over_on_collector_device` - Whether or not the device the Collector is installed on is enabled for fail over (bool)
* `escalating_chain_id` - The Id of the escalation chain associated with this Collector (int32)
* `monitor_others` - Check if we shall monitor using local account (for windows) (bool), the default value is true
* `need_auto_create_collector_device` - Whether to create a collector device when instance collector, the default value is true (bool)
* `number_of_instances` - The number of instances that are monitored by this collector (int32)
* `onetime_downgrade_info` - The details of the Collector's automatic downgrade schedule, if one exists (OnetimeUpgradeInfo)
  + `description`
  + `major_version` (required)
  + `minor_version` (required)
  + `start_epoch` (required)
  + `timezone`
* `onetime_upgrade_info` - The details of the Collector's one time upgrade, if one has been scheduled (OnetimeUpgradeInfo)
  + `description`
  + `major_version` (required)
  + `minor_version` (required)
  + `start_epoch` (required)
  + `timezone`
* `resend_ival` - The interval, in minutes, after which alert notifications for the Collector will be resent (int32)
* `specified_collector_device_group_id` - The collector device group id assigned when creating a new collector device (int32)
* `suppress_alert_clear` - Whether alert clear notifications are suppressed for the Collector (bool)

## Import

collectors can be imported using their collector ID
```
$ terraform import logicmonitor_collector.my_collector 66
```