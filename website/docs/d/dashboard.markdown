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
# Datasource to get information of LogicMonitor dashboard
data "logicmonitor_Dashboard" "my_Dashboard" {
        filter = "description~\"my dashboard\""
        depends_on = [
            logicmonitor_dashboard.mydashboard
        ]
}
```

## Argument Reference

The following arguments are supported:
* `filter` - (Optional) Filters the response according to the operator and value specified.More Info: https://www.logicmonitor.com/support/adding-dashboards. Please refer the filter arguments from resources tab.

* `depends_on` - (Optional) meta-argument within data blocks defers reading of the data source until after all changes to the dependencies have been applied.


## Retrieving multiple matches

`logicmonitor_dashboards` returns every dashboard that matches `filter` (or all of them, if `filter` is omitted), instead of requiring a filter that narrows the result down to exactly one record.

```hcl
data "logicmonitor_dashboards" "all_matches" {
  filter    = "description~\"my dashboard\""
  page_size = 300
}

output "dashboard_count" {
  value = data.logicmonitor_dashboards.all_matches.total
}
```

### Argument Reference

* `filter` - (Optional) Same filter syntax as `logicmonitor_dashboard` above. Omit it to retrieve every dashboard.
* `page_size` - (Optional) Number of records requested per underlying API page while the data source transparently pages through all matches. Defaults to `300`.

### Attribute Reference

* `total` - Total number of dashboard records matching `filter`.
* `items` - List of matching dashboard records. Each entry exposes the same read-only attributes as `logicmonitor_dashboard` above, so e.g. `data.logicmonitor_dashboards.all_matches.items[0].id` is available. This is typically consumed with a `for_each`:

```hcl
resource "logicmonitor_dashboard" "imported" {
  for_each = {
    for item in data.logicmonitor_dashboards.all_matches.items :
    tostring(item.id) => item
  }
  # ... map the remaining arguments from each.value as needed
}
```