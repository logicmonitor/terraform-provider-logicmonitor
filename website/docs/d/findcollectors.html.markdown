---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_findcollectors"
sidebar_current: "docs-logicmonitor-datasource-findcollectors"
description: |-
  Use this datasource to get the ID of an available collector.
---

# logicmonitor_findcollectors

Use this datasource to get the ID of an available collector.

## Example Usage

```hcl
# Look up a LogicMonitor collector id
data "logicmonitor_findcollectors" "collectors" {
  filters {
    "property" = "hostname"
    "operator" = "~"
    "value" = "test"
  },
"most_recent" = true
}
```

## Argument Reference

The following arguments are supported:

* `size` - (Optional) The number of results to display. Max is 1000. Default is 50
* `offset` - (Optional) The number of results to offset the displayed results by. Default is 0
* `most_recent` - (Optional) The most recent collector installed that is online
* `filters` - (Optional) Filters the response according to the operator and value specified. Note that you can use * to match on more than one character. More Info: https://www.logicmonitor.com/support/rest-api-developers-guide/device-groups/get-device-groups/

### Nested `filters` blocks

Nested `filters` blocks have the following structure: `property{operator}value`
* `property` - (Required if using filters) The name of filtered property. Currently the properties supported are `hostname` and `description`
* `operator` - (Required if using filters) The type of operator. Currently the operators supported are `:` `~` `!:` `!~`
* `value` - (Required if using filters) The value of the filtered property.
