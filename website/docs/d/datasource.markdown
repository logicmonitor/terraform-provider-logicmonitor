---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_datasource"
sidebar_current: "docs-logicmonitor-datasources-datasource"
description: |-
  Get information on a LogicMonitor datasource resource
---

# logicmonitor_datasource

This can be used to get information on a LogicMonitor datasource resource given a filter value from argument list

## Example Usage    
### Datasource
```hcl
# create a new LogicMonitor datasource
data "logicmonitor_Datasource" "my_Datasource" {
       filter = "description~\"datasource test\""
 	   depends_on = [
	   logicmonitor_datasource.my_datasource
 	   ]
}
```

## Argument Reference

The following arguments are supported:

* `filter` - (Optional) Filters the response according to the operator and value specified.More Info: https://www.logicmonitor.com/support/rest-api-developers-guide/v1/device-groups/. Please refer the filter arguments from resources tab.
* `depends_on` - (Optional) meta-argument within data blocks defers reading of the data source until after all changes to the dependencies have been applied.

