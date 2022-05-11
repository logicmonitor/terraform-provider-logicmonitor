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
# create a new LogicMonitor device group
data "logicmonitor_DeviceGroup" "my_DeviceGroup" {
        filter = "description~\"normal device group test\""
        depends_on = [
            logicmonitor_device_group.my_device_group
        ]
}
```

## Argument Reference

The following arguments are supported:

* `filter` - (Optional) Filters the response according to the operator and value specified.More Info: https://www.logicmonitor.com/support/rest-api-developers-guide/v1/device-groups/. Please refer the filter arguments from resources tab.
* `depends_on` - (Optional) meta-argument within data blocks defers reading of the data source until after all changes to the dependencies have been applied.

