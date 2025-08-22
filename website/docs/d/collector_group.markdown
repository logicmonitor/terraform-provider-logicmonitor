---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_collector_group"
sidebar_current: "docs-logicmonitor-datasources-collector-group"
description: |-
  Get information on a LogicMonitor collector group resource
---

# logicmonitor_collector_group

This can be used to get information on a LogicMonitor collector group resource given a filter value from argument list

## Example Usage    
### CollectorGroup
```hcl
# Datasource to get information of LogicMonitor collector group
data "logicmonitor_CollectorGroup" "my_CollectorGroup" {
        filter = "description~\"Group for collectors dedicated to testing purpose\""
        depends_on = [
            logicmonitor_collector_group.my_collector_group
        ]
}
```

## Argument Reference

The following arguments are supported:
* `filter` - (Optional) Filters the response according to the operator and value specified.More Info: https://www.logicmonitor.com/support/adding-collector-groups. Please refer the filter arguments from resources tab.

* `depends_on` - (Optional) meta-argument within data blocks defers reading of the data source until after all changes to the dependencies have been applied.

