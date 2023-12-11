---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_dashboard"
sidebar_current: "docs-logicmonitor-datasources-dashboard"
description: |-
  Get information on a LogicMonitor dashboard resource
---

# logicmonitor_dashboard

This can be used to get information on a LogicMonitor dashboard resource given a filter value from argument list

## Example Usage    
### Dashboard
```hcl
# create a new LogicMonitor dashboard
data "logicmonitor_Dashboard" "my_Dashboard" {
        filter = "description~\"my dashboard\""
        depends_on = [
            logicmonitor_dashboard.mydashboard
        ]
}
```

## Argument Reference

The following arguments are supported:
* `filter` - (Optional) Filters the response according to the operator and value specified.More Info: https://www.logicmonitor.com/support/rest-api-developers-guide/v1/dashboards/about-the-dashboards-resource. Please refer the filter arguments from resources tab.

* `depends_on` - (Optional) meta-argument within data blocks defers reading of the data source until after all changes to the dependencies have been applied.

