---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_collector"
sidebar_current: "docs-logicmonitor-datasources-collector"
description: |-
  Get information on a LogicMonitor collector resource
---

# logicmonitor_collector

This can be used to get information on a LogicMonitor collector resource given a filter value from argument list

## Example Usage    
### Collector
```hcl
# create a new LogicMonitor collector
data "logicmonitor_Collector" "my_Collector" {
        filter = "description~\"Linux Collector\""
        depends_on = [
            logicmonitor_collector.my_collector
        ]
}
```

## Argument Reference

The following arguments are supported:

* `filter` - (Optional) Filters the response according to the operator and value specified.More Info: https://www.logicmonitor.com/support/rest-api-developers-guide/v1/device-groups/. Please refer the filter arguments from resources tab.
* `depends_on` - (Optional) meta-argument within data blocks defers reading of the data source until after all changes to the dependencies have been applied.

