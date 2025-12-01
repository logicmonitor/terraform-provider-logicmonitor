---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_role"
sidebar_current: "docs-logicmonitor-datasources-role"
description: |-
  Get information on a LogicMonitor role resource
---

# logicmonitor_role

This can be used to get information on a LogicMonitor role resource given a filter value from argument list

## Example Usage    
### Role
```hcl
# Datasource to get information of LogicMonitor role
data "logicmonitor_Role" "my_Role" {
       filter = "description~\"Role test\""
 	   depends_on = [
	   logicmonitor_role.my_role
 	   ]
}
```

## Argument Reference

The following arguments are supported:
* `filter` - (Optional) Filters the response according to the operator and value specified.More Info: https://www.logicmonitor.com/support/adding-roles

* `depends_on` - (Optional) meta-argument within data blocks defers reading of the data source until after all changes to the dependencies have been applied.

