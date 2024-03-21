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

