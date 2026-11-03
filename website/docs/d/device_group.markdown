---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_device_group"
sidebar_current: "docs-logicmonitor-datasources-device-group"
description: |-
  Get information on a LogicMonitor device group resource
---

# logicmonitor_device_group

This can be used to get information on a LogicMonitor device group resource given a filter value from argument list

## Example Usage    
### DeviceGroup
```hcl
# Datasource to get information of LogicMonitor device group
data "logicmonitor_DeviceGroup" "my_DeviceGroup" {
        filter = "description~\"normal device group test\""
        depends_on = [
            logicmonitor_device_group.my_device_group
        ]
}
```

## Argument Reference

The following arguments are supported:
* `filter` - (Optional) Filters the response according to the operator and value specified.More Info: https://www.logicmonitor.com/support/rest-api-developers-guide/v1/device-groups/about-the-device-group-resource. Please refer the filter arguments from resources tab.

* `depends_on` - (Optional) meta-argument within data blocks defers reading of the data source until after all changes to the dependencies have been applied.


## Retrieving multiple matches

`logicmonitor_device_groups` returns every device group that matches `filter` (or all of them, if `filter` is omitted), instead of requiring a filter that narrows the result down to exactly one record.

```hcl
data "logicmonitor_device_groups" "all_matches" {
  filter    = "description~\"normal device group test\""
  page_size = 300
}

output "device_group_count" {
  value = data.logicmonitor_device_groups.all_matches.total
}
```

### Argument Reference

* `filter` - (Optional) Same filter syntax as `logicmonitor_device_group` above. Omit it to retrieve every device group.
* `page_size` - (Optional) Number of records requested per underlying API page while the data source transparently pages through all matches. Defaults to `300`.

### Attribute Reference

* `total` - Total number of device group records matching `filter`.
* `items` - List of matching device group records. Each entry exposes the same read-only attributes as `logicmonitor_device_group` above, so e.g. `data.logicmonitor_device_groups.all_matches.items[0].id` is available. This is typically consumed with a `for_each`:

```hcl
resource "logicmonitor_device_group" "imported" {
  for_each = {
    for item in data.logicmonitor_device_groups.all_matches.items :
    tostring(item.id) => item
  }
  # ... map the remaining arguments from each.value as needed
}
```