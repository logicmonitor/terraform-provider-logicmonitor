---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_sdt"
sidebar_current: "docs-logicmonitor-datasources-sdt"
description: |-
  Get information on a LogicMonitor sdt resource
---

# logicmonitor_sdt

This can be used to get information on a LogicMonitor sdt resource given a filter value from argument list

## Example Usage    
### Sdt
```hcl
# Datasource to get information of LogicMonitor sdt
data "logicmonitor_Sdt" "my_Sdt" {
       filter = "description~\"Sdt test\""
 	   depends_on = [
	   logicmonitor_sdt.my_sdt
 	   ]
}
```

## Argument Reference

The following arguments are supported:
* `filter` - (Optional) Filters the response according to the operator and value specified.More Info: https://www.logicmonitor.com/support/rest-api-developers-guide/v1/sdts/about-the-sdt-resource. Please refer the filter arguments from resources tab.

* `depends_on` - (Optional) meta-argument within data blocks defers reading of the data source until after all changes to the dependencies have been applied.

