---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_device_group_cluster_alert_conf"
sidebar_current: "docs-logicmonitor-resource-device-group-cluster-alert-conf"
description: |-
  Provides a LogicMonitor device group cluster alert conf resource. This can be used to create and manage LogicMonitor device group cluster alert confs.
---

# logicmonitor_device_group_cluster_alert_conf

Provides a LogicMonitor device group cluster alert conf resource. This can be used to create and manage LogicMonitor device group cluster alert confs.

## Example Usage
```hcl
# Create a LogicMonitor device group cluster alert conf
resource "logicmonitor_device_group_cluster_alert_conf" "my_device_group_cluster_alert_conf" {
  alert_expr = "\u003e=2"
  count_by = "host"
  data_point_id = 2016
  data_source_display_name = "Interfaces (64 bit)-"
  data_source_id = 396
  device_group_id = 
  disable_alerting = false
  min_alert_level = 4
  suppress_ind_alert = true
  threshold_type = "absolute"
}
```

## Argument Reference

The following arguments are **required**:

The following arguments are **optional**:
* `alert_expr` - The expression that indicates the number of objects (devices or instances) that need to be in alert to trigger the cluster alert. E.g. > 5 (string)
* `count_by` - host | instance - Whether the cluster alert is based on an alert count across devices or instances (string)
* `data_point_id` - The id of the dataPoint you want to base the cluster alert on (int32)
* `data_source_display_name` - The display name of the dataSource you want to base the cluster alert on (string)
* `data_source_id` - The id of the dataSource you want to base the cluster alert on (int32)
* `device_group_id` - The id of the device group that the SDT will be associated with (int32)
* `disable_alerting` - Whether or not alerting will be disabled (bool)
* `min_alert_level` - The alert level that must be present for the devices or instances to trigger the cluster alert. Acceptable values are: 2, 3, 4 (int32)
* `suppress_ind_alert` - Whether or not alerting will be suppressed for individual alerts, the default value is true (bool)
* `threshold_type` - whether the alert expression should be evaluated as a total number of devices or instances (absolute) or as a percentage of devices or instances (percentage). Acceptable values are: absolute, percentage (string)

## Import

device group cluster alert confs can be imported using their device group cluster alert conf ID or name
```
$ terraform import logicmonitor_device_group_cluster_alert_conf.my_device_group_cluster_alert_conf 66
$ terraform import logicmonitor_device_group_cluster_alert_conf.my_device_group_cluster_alert_conf
```