---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_device"
sidebar_current: "docs-logicmonitor-resource-device-x"
description: |-
  Provides a LogicMonitor device resource. This can be used to create and manage LogicMonitor devices
---

# logicmonitor_device

Provides a LogicMonitor device resource. This can be used to create and manage LogicMonitor devices

## Example Usage

```hcl
# Create a new LogicMonitor device
resource "logicmonitor_device" "host" {
  ip_addr = "10.32.12.18"
  disable_alerting = true
  collector = "2"
  properties = {
   "app" = "haproxy"
   "system.categories" = "a,b,c,d"
  }
}
```

```hcl
# Create a new LogicMonitor device and device group with some data source lookups and computed attributes.
resource "logicmonitor_device" "host" {
  ip_addr = "10.32.12.18"
  disable_alerting = true
  collector = "${data.logicmonitor_collectors.collectors.id}"
  hostgroup_id = "${logicmonitor_device_group.group1.id}"
  properties = {
   "app" = "haproxy"
   "system.categories" = "a,b,c,d"
  }
}

resource "logicmonitor_device_group" "group1" {
    name = "newgroup"
    properties = {
     "system.categories" = "ec2"
     "jmx.port" = "3008"
     "snmp.version" = "v2c"
    }
}

data "logicmonitor_collectors" "collectors" {
  most_recent = true
}
```

## Argument Reference

The following arguments are supported:

* `ip_addr` - (Required) Ip Address/Hostname of device
* `collector` - (required) The id of the collector that will monitoring the device
* `display_name` - (Optional) Display name of device, (default is ip_addr)
* `disable_alerting` - (Optional) The host is created with alerting disabled (default is true)
* `hostgroup_id` - (Optional) The host group id that specifies which group the device belongs to (multiple host group ids can be added, represented by a comma separated string)
* `properties` - (Optional) The properties associated with this device group. Any string value pair will work (see example).

## Import

Devices can be imported using their device id or ip address/dns name

```
$ terraform import logicmonitor_device.host 751
$ terraform import logicmonitor_device.host server01.us-east-1.logicmonitor.net
```
