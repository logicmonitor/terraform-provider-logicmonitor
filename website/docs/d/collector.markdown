---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_collector"
sidebar_current: "docs-logicmonitor-datasources-collector"
description: |-
  Get information on a LogicMonitor collector resource
---

# logicmonitor_collector

This can be used to get information on a LogicMonitor collector resource given a filter value from argument list

## Example Usage    
### Collector
```hcl
# Datasource to get information of LogicMonitor collector
data "logicmonitor_Collector" "my_Collector" {
        filter = "description~\"Linux Collector\""
        depends_on = [
            logicmonitor_collector.my_collector
        ]
}
```

## Argument Reference

The following arguments are supported:
* `filter` - (Optional) Filters the response according to the operator and value specified.More Info: https://www.logicmonitor.com/support/rest-api-developers-guide/v1/collectors/about-the-collectors-resource. Please refer the filter arguments from resources tab.

* `depends_on` - (Optional) meta-argument within data blocks defers reading of the data source until after all changes to the dependencies have been applied.


## Retrieving multiple matches

`logicmonitor_collectors` returns every collector that matches `filter` (or all of them, if `filter` is omitted), instead of requiring a filter that narrows the result down to exactly one record.

```hcl
data "logicmonitor_collectors" "all_matches" {
  filter    = "description~\"Linux Collector\""
  page_size = 300
}

output "collector_count" {
  value = data.logicmonitor_collectors.all_matches.total
}
```

### Argument Reference

* `filter` - (Optional) Same filter syntax as `logicmonitor_collector` above. Omit it to retrieve every collector.
* `page_size` - (Optional) Number of records requested per underlying API page while the data source transparently pages through all matches. Defaults to `300`.

### Attribute Reference

* `total` - Total number of collector records matching `filter`.
* `items` - List of matching collector records. Each entry exposes the same read-only attributes as `logicmonitor_collector` above, so e.g. `data.logicmonitor_collectors.all_matches.items[0].id` is available. This is typically consumed with a `for_each`:

```hcl
resource "logicmonitor_collector" "imported" {
  for_each = {
    for item in data.logicmonitor_collectors.all_matches.items :
    tostring(item.id) => item
  }
  # ... map the remaining arguments from each.value as needed
}
```