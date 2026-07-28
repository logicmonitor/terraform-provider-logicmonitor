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