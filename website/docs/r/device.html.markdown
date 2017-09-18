---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_device"
sidebar_current: "docs-logicmonitor-resource-device"
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
  collector = "123"
  hostgroup_id = "2"
  properties {
   "app" = "haproxy"
   "system.categories" = "a,b,c,d"
  }
}
```

```hcl
# Create a new LogicMonitor device with some data source lookups
resource "logicmonitor_device" "host" {
  ip_addr = "10.32.12.18"
  disable_alerting = true
  collector = "${data.logicmonitor_collectors.collectors.id}"
  hostgroup_id = "${logicmonitor_device_group.devicegroups.id}"
  properties {
   "app" = "haproxy"
   "system.categories" = "a,b,c,d"
  }
}

data "logicmonitor_device_group" "devicegroups" {
  filters {
    "property" = "name"
    "operator" = "~"
    "value" = "Mesos"
  },
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
