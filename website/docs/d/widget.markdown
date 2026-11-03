---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_widget"
sidebar_current: "docs-logicmonitor-datasources-widget"
description: |-
  Get information on a LogicMonitor widget resource
---

# logicmonitor_widget

This can be used to get information on a LogicMonitor widget resource given a filter value from argument list

## Example Usage    
### Widget
```hcl
# Datasource to get information of LogicMonitor widget
data "logicmonitor_Widget" "my_Widget" {
        filter = "name~\"Memory - log-ingest:pods\""
        depends_on = [
            logicmonitor_widget.custom_graph
        ]
}
```

## Argument Reference

The following arguments are supported:
* `filter` - (Optional) Filters the response according to the operator and value specified. More Info: https://www.logicmonitor.com/support/widgets-overview. Please refer the filter arguments from resources tab. Supported widget types include `cgraph`, `bigNumber`, and `pieChart`.

* `depends_on` - (Optional) meta-argument within data blocks defers reading of the data source until after all changes to the dependencies have been applied.


## Retrieving multiple matches

`logicmonitor_widgets` returns every widget that matches `filter` (or all of them, if `filter` is omitted), instead of requiring a filter that narrows the result down to exactly one record.

```hcl
data "logicmonitor_widgets" "all_matches" {
  filter    = "name~\"Memory - log-ingest:pods\""
  page_size = 300
}

output "widget_count" {
  value = data.logicmonitor_widgets.all_matches.total
}
```

### Argument Reference

* `filter` - (Optional) Same filter syntax as `logicmonitor_widget` above. Omit it to retrieve every widget.
* `page_size` - (Optional) Number of records requested per underlying API page while the data source transparently pages through all matches. Defaults to `300`.

### Attribute Reference

* `total` - Total number of widget records matching `filter`.
* `items` - List of matching widget records. Each entry exposes the same read-only attributes as `logicmonitor_widget` above, so e.g. `data.logicmonitor_widgets.all_matches.items[0].id` is available. This is typically consumed with a `for_each`:

```hcl
resource "logicmonitor_widget" "imported" {
  for_each = {
    for item in data.logicmonitor_widgets.all_matches.items :
    tostring(item.id) => item
  }
  # ... map the remaining arguments from each.value as needed
}
```