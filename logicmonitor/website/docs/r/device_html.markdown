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
  current_collector_id = 1
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
  description = "This is a Cisco Router"
  device_type = 0
  disable_alerting = true
  display_name = "Cisco Router"
  enable_netflow = true
  host_group_ids = "16,4,3"
  link = "www.ciscorouter.com"
  name = "collector.localhost"
  netflow_collector_id = 1
  preferred_collector_id = 2
  related_device_id = -1
}
```

## Argument Reference

The following arguments are **required**:
* `display_name` - The display name of the device
* `name` - The host name or IP address of the device
* `preferred_collector_id` - The Id of the preferred collector assigned to monitor the device

The following arguments are **optional**:
* `auto_balanced_collector_group_id` - The Auto Balanced Collector Group id. 0 means not monitored by ABCG
* `current_collector_id` - The id of the collector currently monitoring the device and discovering instances
* `custom_properties` - Any non-system properties (aside from system.categories) defined for this device
  + `name` - The name of a property (required)
  + `value` - The value of a property (required)
* `description` - The device description
* `device_type` - The type of device: 0 indicates a regular device, 2 indicates an AWS device, 4 indicates an Azure device
* `disable_alerting` - Indicates whether alerting is disabled (true) or enabled (false) for this device
* `enable_netflow` - Indicates whether Netflow is enabled (true) or disabled (false) for the device
* `host_group_ids` - The Id(s) of the groups the device is in, where multiple group ids are comma separated
* `link` - The URL link associated with the device
* `netflow_collector_id` - The Id of the netflow collector associated with the device
* `related_device_id` - The Id of the AWS EC2 instance related to this device, if one exists in the LogicMonitor account. This value defaults to -1, which indicates that there are no related devices

## Import

devices can be imported using their device ID or name
```
$ terraform import logicmonitor_device.my_device 66
$ terraform import logicmonitor_device.my_device collector.localhost
```