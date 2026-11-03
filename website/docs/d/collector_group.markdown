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


## Retrieving multiple matches

`logicmonitor_collector_groups` returns every collector group that matches `filter` (or all of them, if `filter` is omitted), instead of requiring a filter that narrows the result down to exactly one record.

```hcl
data "logicmonitor_collector_groups" "all_matches" {
  filter    = "description~\"Group for collectors dedicated to testing purpose\""
  page_size = 300
}

output "collector_group_count" {
  value = data.logicmonitor_collector_groups.all_matches.total
}
```

### Argument Reference

* `filter` - (Optional) Same filter syntax as `logicmonitor_collector_group` above. Omit it to retrieve every collector group.
* `page_size` - (Optional) Number of records requested per underlying API page while the data source transparently pages through all matches. Defaults to `300`.

### Attribute Reference

* `total` - Total number of collector group records matching `filter`.
* `items` - List of matching collector group records. Each entry exposes the same read-only attributes as `logicmonitor_collector_group` above, so e.g. `data.logicmonitor_collector_groups.all_matches.items[0].id` is available. This is typically consumed with a `for_each`:

```hcl
resource "logicmonitor_collector_group" "imported" {
  for_each = {
    for item in data.logicmonitor_collector_groups.all_matches.items :
    tostring(item.id) => item
  }
  # ... map the remaining arguments from each.value as needed
}
```