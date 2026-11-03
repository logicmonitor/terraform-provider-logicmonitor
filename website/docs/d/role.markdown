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


## Retrieving multiple matches

`logicmonitor_roles` returns every role that matches `filter` (or all of them, if `filter` is omitted), instead of requiring a filter that narrows the result down to exactly one record.

```hcl
data "logicmonitor_roles" "all_matches" {
  filter    = "description~\"Role test\""
  page_size = 300
}

output "role_count" {
  value = data.logicmonitor_roles.all_matches.total
}
```

### Argument Reference

* `filter` - (Optional) Same filter syntax as `logicmonitor_role` above. Omit it to retrieve every role.
* `page_size` - (Optional) Number of records requested per underlying API page while the data source transparently pages through all matches. Defaults to `300`.

### Attribute Reference

* `total` - Total number of role records matching `filter`.
* `items` - List of matching role records. Each entry exposes the same read-only attributes as `logicmonitor_role` above, so e.g. `data.logicmonitor_roles.all_matches.items[0].id` is available. This is typically consumed with a `for_each`:

```hcl
resource "logicmonitor_role" "imported" {
  for_each = {
    for item in data.logicmonitor_roles.all_matches.items :
    tostring(item.id) => item
  }
  # ... map the remaining arguments from each.value as needed
}
```