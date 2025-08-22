---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_dashboard_group"
sidebar_current: "docs-logicmonitor-datasources-dashboard-group"
description: |-
  Get information on a LogicMonitor dashboard group resource
---

# logicmonitor_dashboard_group

This can be used to get information on a LogicMonitor dashboard group resource given a filter value from argument list

## Example Usage    
### DashboardGroup
```hcl
# Datasource to get information of LogicMonitor dashboard group
data "logicmonitor_DashboardGroup" "my_DashboardGroup" {
        filter = "description~\"LM dashboard group testing\""
        depends_on = [
            logicmonitor_dashboard_group.myDashboardGroup
        ]
}
```

## Argument Reference

The following arguments are supported:
* `filter` - (Optional) Filters the response according to the operator and value specified.More Info: https://www.logicmonitor.com/support/adding-dashboard-groups. Please refer the filter arguments from resources tab.

* `depends_on` - (Optional) meta-argument within data blocks defers reading of the data source until after all changes to the dependencies have been applied.

