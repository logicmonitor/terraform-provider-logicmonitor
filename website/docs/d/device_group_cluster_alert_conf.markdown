---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_device_group_cluster_alert_conf"
sidebar_current: "docs-logicmonitor-datasources-device-group-cluster-alert-conf"
description: |-
  Get information on a LogicMonitor device group cluster alert conf resource
---

# logicmonitor_device_group_cluster_alert_conf

This can be used to get information on a LogicMonitor device group cluster alert conf resource given a filter value from argument list

## Example Usage    
### DeviceGroupClusterAlertConf
```hcl
# Datasource to get information of LogicMonitor device group cluster alert conf
data "logicmonitor_DeviceGroupClusterAlertConf" "my_DeviceGroupClusterAlertConf" {
}
```

## Argument Reference

The following arguments are supported:

* `depends_on` - (Optional) meta-argument within data blocks defers reading of the data source until after all changes to the dependencies have been applied.

