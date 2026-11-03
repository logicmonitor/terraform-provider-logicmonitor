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


## Retrieving multiple matches

`logicmonitor_dashboard_groups` returns every dashboard group that matches `filter` (or all of them, if `filter` is omitted), instead of requiring a filter that narrows the result down to exactly one record.

```hcl
data "logicmonitor_dashboard_groups" "all_matches" {
  filter    = "description~\"LM dashboard group testing\""
  page_size = 300
}

output "dashboard_group_count" {
  value = data.logicmonitor_dashboard_groups.all_matches.total
}
```

### Argument Reference

* `filter` - (Optional) Same filter syntax as `logicmonitor_dashboard_group` above. Omit it to retrieve every dashboard group.
* `page_size` - (Optional) Number of records requested per underlying API page while the data source transparently pages through all matches. Defaults to `300`.

### Attribute Reference

* `total` - Total number of dashboard group records matching `filter`.
* `items` - List of matching dashboard group records. Each entry exposes the same read-only attributes as `logicmonitor_dashboard_group` above, so e.g. `data.logicmonitor_dashboard_groups.all_matches.items[0].id` is available. This is typically consumed with a `for_each`:

```hcl
resource "logicmonitor_dashboard_group" "imported" {
  for_each = {
    for item in data.logicmonitor_dashboard_groups.all_matches.items :
    tostring(item.id) => item
  }
  # ... map the remaining arguments from each.value as needed
}
```