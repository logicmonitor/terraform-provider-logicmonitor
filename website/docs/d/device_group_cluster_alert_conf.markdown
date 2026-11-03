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


## Retrieving multiple matches

`logicmonitor_device_group_cluster_alert_confs` returns every device group cluster alert conf that matches `filter` (or all of them, if `filter` is omitted), instead of requiring a filter that narrows the result down to exactly one record.

```hcl
data "logicmonitor_device_group_cluster_alert_confs" "all_matches" {
  filter    = ""
  page_size = 300
}

output "device_group_cluster_alert_conf_count" {
  value = data.logicmonitor_device_group_cluster_alert_confs.all_matches.total
}
```

### Argument Reference

* `filter` - (Optional) Same filter syntax as `logicmonitor_device_group_cluster_alert_conf` above. Omit it to retrieve every device group cluster alert conf.
* `page_size` - (Optional) Number of records requested per underlying API page while the data source transparently pages through all matches. Defaults to `300`.

### Attribute Reference

* `total` - Total number of device group cluster alert conf records matching `filter`.
* `items` - List of matching device group cluster alert conf records. Each entry exposes the same read-only attributes as `logicmonitor_device_group_cluster_alert_conf` above, so e.g. `data.logicmonitor_device_group_cluster_alert_confs.all_matches.items[0].id` is available. This is typically consumed with a `for_each`:

```hcl
resource "logicmonitor_device_group_cluster_alert_conf" "imported" {
  for_each = {
    for item in data.logicmonitor_device_group_cluster_alert_confs.all_matches.items :
    tostring(item.id) => item
  }
  # ... map the remaining arguments from each.value as needed
}
```