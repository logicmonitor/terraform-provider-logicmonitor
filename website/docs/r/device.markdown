---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_device"
sidebar_current: "docs-logicmonitor-resource-device"
description: |-
  Provides a LogicMonitor device resource. This can be used to create and manage LogicMonitor devices.
---

# logicmonitor_device

Provides a LogicMonitor device resource. This can be used to create and manage LogicMonitor devices.

## Example Usage
```hcl
# Create a LogicMonitor device
resource "logicmonitor_device" "my_device" {
  auto_balanced_collector_group_id = 0
  contains_multi_value = 
  current_collector_id = 1
  current_log_collector_id = 1
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
  description = "This is a Cisco Router"
  device_type = 0
  disable_alerting = true
  display_name = "Cisco Router"
  enable_netflow = true
  host_group_ids = "16,4,3"
  is_preferred_log_collector_configured = true
  link = "www.ciscorouter.com"
  log_collector_id = 1
  name = "collector.localhost"
  need_stc_grp_and_sorted_c_p = true
  netflow_collector_id = 1
  op = ""
  preferred_collector_id = 2
  related_device_id = -1
}
```

## Argument Reference

The following arguments are **required**:
* `display_name` - The display name of the device
   (string)
* `name` - The host name or IP address of the device
   (string)
* `preferred_collector_id` - The Id of the preferred collector assigned to monitor the device
   (int32)

The following arguments are **optional**:
* `auto_balanced_collector_group_id` - The Auto Balanced Collector Group id. 0 means not monitored by ABCG (int32)
* `contains_multi_value` - request contains multi value field (bool)
* `current_collector_id` - The id of the collector currently monitoring the device and discovering instances (int32)
* `current_log_collector_id` - The id of the Log collector currently collecting logs. (int32)
* `custom_properties` - Any non-system properties (aside from system.categories) defined for this device ([]*NameAndValue)
    Provide custom properties in alphabetical order based on their property name.
  + `name` - The name of a property (required)
  + `value` - The value of a property (required)
* `description` - The device description (string)
* `device_type` - The type of device: 0 indicates a regular device, 2 indicates an AWS device, 4 indicates an Azure device (int32)
* `disable_alerting` - Indicates whether alerting is disabled (true) or enabled (false) for this device (bool)
* `enable_netflow` - Indicates whether Netflow is enabled (true) or disabled (false) for the device (bool)
* `host_group_ids` - The Id(s) of the groups the device is in, where multiple group ids are comma separated (string)
* `is_preferred_log_collector_configured` - Indicates whether Preferred Log Collector is configured  (true) or not (false) for the device (bool)
* `link` - The URL link associated with the device (string)
* `log_collector_id` - The Id of the netflow collector associated with the device (int32)
* `need_stc_grp_and_sorted_c_p` - Indicates whether Static group and sorted Custom properties are needed (bool)
* `netflow_collector_id` - The Id of the netflow collector associated with the device (int32)
* `op` - whether to use AND or OR for device matching (string)
* `preferred_collector_group_id` - The id of the Collector Group associated with the device's preferred collector (int32), It can be 0 for auto balanced collector group .
* `related_device_id` - The Id of the AWS EC2 instance related to this device, if one exists in the LogicMonitor account. This value defaults to -1, which indicates that there are no related devices (int32)
* `resource_ids` - Any non-system properties (aside from system.categories) defined for this device ([]*NameAndValue)
* `synthetics_collector_ids` - The list of ids of the collectors currently monitoring the resource and discovering instances ([]int32)

## Import

devices can be imported using their device ID or name
```
$ terraform import logicmonitor_device.my_device 66
$ terraform import logicmonitor_device.my_device collector.localhost
```