---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_collector"
sidebar_current: "docs-logicmonitor-resource-collector"
description: |-
  Provides a LogicMonitor collector resource. This can be used to create and manage LogicMonitor collectors
---

# logicmonitor_collector

Provides a LogicMonitor collector resource. This can be used to create and manage LogicMonitor collectors.

*Note:* This resource will only create the collector device in your account. See [Downloading a Collector Installer](https://www.logicmonitor.com/support/rest-api-developers-guide/collectors/downloading-a-collector-installer/) for
information on how to download and install an existing collector.

## Example Usage

```hcl
# Create a new LogicMonitor collector
resource "logicmonitor_collector" "collector1" {
  description     = "my terraformed collector"
  enable_failback = true
}
```

## Argument Reference

The following arguments are supported:

* `backup_collector_id` - (Optional) The Id of the failover Collector configured for this Collector
* `collector_group_id` - (Optional) The Id of the group the Collector is in
* `description` - (Optional) The Collector's description
* `enable_failback` - (Optional) Whether or not automatic failback is enabled for the Collector
* `enable_collector_device_failover` - (Optional) Whether or not the device the Collector is installed on is enabled for fail over
* `escalation_chain_id` - (Optional) The Id of the escalation chain associated with this Collector
* `resend_interval` - (Optional) The interval, in minutes, after which alert notifications for the Collector will be resent
* `suppress_alert_clear` - (Optional) Whether alert clear notifications are suppressed for the Collector
