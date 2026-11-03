---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_device"
sidebar_current: "docs-logicmonitor-datasources-device"
description: |-
  Get information on a LogicMonitor device resource
---

# logicmonitor_device

This can be used to get information on a LogicMonitor device resource given a filter value from argument list

## Example Usage    
### Device
```hcl
# Datasource to get information of LogicMonitor device
data "logicmonitor_Device" "my_Device" {
        filter = "displayName~\"Cisco Router Test\""
        depends_on = [
            logicmonitor_device.my_device
        ]
}
```

## Argument Reference

The following arguments are supported:
* `filter` - (Optional) Filters the response according to the operator and value specified.More Info: https://www.logicmonitor.com/support/rest-api-developers-guide/v1/devices/about-the-device-resource. Please refer the filter arguments from resources tab.

* `depends_on` - (Optional) meta-argument within data blocks defers reading of the data source until after all changes to the dependencies have been applied.


## Retrieving multiple matches

`logicmonitor_devices` returns every device that matches `filter` (or all of them, if `filter` is omitted), instead of requiring a filter that narrows the result down to exactly one record.

```hcl
data "logicmonitor_devices" "all_matches" {
  filter    = "displayName~\"Cisco Router Test\""
  page_size = 300
}

output "device_count" {
  value = data.logicmonitor_devices.all_matches.total
}
```

### Argument Reference

* `filter` - (Optional) Same filter syntax as `logicmonitor_device` above. Omit it to retrieve every device.
* `page_size` - (Optional) Number of records requested per underlying API page while the data source transparently pages through all matches. Defaults to `300`.

### Attribute Reference

* `total` - Total number of device records matching `filter`.
* `items` - List of matching device records. Each entry exposes the same read-only attributes as `logicmonitor_device` above, so e.g. `data.logicmonitor_devices.all_matches.items[0].id` is available. This is typically consumed with a `for_each`:

```hcl
resource "logicmonitor_device" "imported" {
  for_each = {
    for item in data.logicmonitor_devices.all_matches.items :
    tostring(item.id) => item
  }
  # ... map the remaining arguments from each.value as needed
}
```