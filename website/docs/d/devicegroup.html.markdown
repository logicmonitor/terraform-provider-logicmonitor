---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_device_group"
sidebar_current: "docs-logicmonitor-datasource-device-group"
description: |-
  Use this datasource to get the ID of an available device group.
---

# logicmonitor_device_group

Use this datasource to get the ID of an available device group.

## Example Usage

```hcl
# Look up a LogicMonitor device group id
data "logicmonitor_device_group" "devicegroups" {
  filters {
    property = "name"
    operator = ":"
    value = "Production"
  },

  filters {
    custom_property_name = "app.user"
    operator = ":"
    custom_property_value = "api"
  }
}
```

## Argument Reference

The following arguments are supported:

* `size` - (Optional) The number of results to display. Max is 1000. Default is 50
* `offset` - (Optional) The number of results to offset the displayed results by. Default is 0
* `filters` - (Optional) Filters the response according to the operator and value specified. Note that you can use * to match on more than one character. More Info: https://www.logicmonitor.com/support/rest-api-developers-guide/device-groups/get-device-groups/

## Nested filters blocks

Nested `filters` blocks have the following structure: `property{operator}value`
* `property` - (Required if using filters) The name of filtered property.
* `operator` - (Required if using filters) The type of operator.
* `value` - (Required if using filters) The value of the filtered property.

You can also do custom properties
* `custom_property_name` - (Required if using filters and custom properties) The name of filtered custom property.
* `operator` - (Required if using filters) The type of operator.
* `custom_property_value` - (Required if using filters and custom properties) The value of the filtered custom property.
