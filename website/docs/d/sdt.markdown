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


## Retrieving multiple matches

`logicmonitor_sdts` returns every sdt that matches `filter` (or all of them, if `filter` is omitted), instead of requiring a filter that narrows the result down to exactly one record.

```hcl
data "logicmonitor_sdts" "all_matches" {
  filter    = "description~\"Sdt test\""
  page_size = 300
}

output "sdt_count" {
  value = data.logicmonitor_sdts.all_matches.total
}
```

### Argument Reference

* `filter` - (Optional) Same filter syntax as `logicmonitor_sdt` above. Omit it to retrieve every sdt.
* `page_size` - (Optional) Number of records requested per underlying API page while the data source transparently pages through all matches. Defaults to `300`.

### Attribute Reference

* `total` - Total number of sdt records matching `filter`.
* `items` - List of matching sdt records. Each entry exposes the same read-only attributes as `logicmonitor_sdt` above, so e.g. `data.logicmonitor_sdts.all_matches.items[0].id` is available. This is typically consumed with a `for_each`:

```hcl
resource "logicmonitor_sdt" "imported" {
  for_each = {
    for item in data.logicmonitor_sdts.all_matches.items :
    tostring(item.id) => item
  }
  # ... map the remaining arguments from each.value as needed
}
```