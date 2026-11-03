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
# Datasource to get information of LogicMonitor datasource
data "logicmonitor_Datasource" "my_Datasource" {
       filter = "description~\"datasource test\""
 	   depends_on = [
	   logicmonitor_datasource.my_datasource
 	   ]
}
```

## Argument Reference

The following arguments are supported:
* `filter` - (Optional) Filters the response according to the operator and value specified.More Info: https://www.logicmonitor.com/support/logicmodules/datasources/creating-managing-datasources/creating-datasources. Please refer the filter arguments from resources tab.

* `depends_on` - (Optional) meta-argument within data blocks defers reading of the data source until after all changes to the dependencies have been applied.


## Retrieving multiple matches

`logicmonitor_datasources` returns every datasource that matches `filter` (or all of them, if `filter` is omitted), instead of requiring a filter that narrows the result down to exactly one record.

```hcl
data "logicmonitor_datasources" "all_matches" {
  filter    = "description~\"datasource test\""
  page_size = 300
}

output "datasource_count" {
  value = data.logicmonitor_datasources.all_matches.total
}
```

### Argument Reference

* `filter` - (Optional) Same filter syntax as `logicmonitor_datasource` above. Omit it to retrieve every datasource.
* `page_size` - (Optional) Number of records requested per underlying API page while the data source transparently pages through all matches. Defaults to `300`.

### Attribute Reference

* `total` - Total number of datasource records matching `filter`.
* `items` - List of matching datasource records. Each entry exposes the same read-only attributes as `logicmonitor_datasource` above, so e.g. `data.logicmonitor_datasources.all_matches.items[0].id` is available. This is typically consumed with a `for_each`:

```hcl
resource "logicmonitor_datasource" "imported" {
  for_each = {
    for item in data.logicmonitor_datasources.all_matches.items :
    tostring(item.id) => item
  }
  # ... map the remaining arguments from each.value as needed
}
```