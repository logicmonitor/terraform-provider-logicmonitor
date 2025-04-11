---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_collector_group"
sidebar_current: "docs-logicmonitor-resource-collector-group"
description: |-
  Provides a LogicMonitor collector group resource. This can be used to create and manage LogicMonitor collector groups.
---

# logicmonitor_collector_group

Provides a LogicMonitor collector group resource. This can be used to create and manage LogicMonitor collector groups.

## Example Usage
```hcl
# Create a LogicMonitor collector group
resource "logicmonitor_collector_group" "my_collector_group" {
  auto_balance = true
  auto_balance_instance_count_threshold = 10000
  auto_balance_strategy = "none"
  custom_properties = [
		{
			name = "addr"
      		value = "127.0.0.1"
		},
		{
			name = "host"
      		value = "localhost"
		},
    {
      name  = "system.categories"
      value = "" 
    }
	]
  description = "Group for collectors dedicated to Network Devices"
  name = "Collector_Group_(Network Devices)"
}
```

## Argument Reference

The following arguments are **required**:
* `name` - The name of the Collector Group
   (string)

The following arguments are **optional**:
* `auto_balance` - Denotes whether or not the collector group should be auto balanced (bool)
* `auto_balance_instance_count_threshold` - Threshold for instance count strategy to check if a collector has high load (int32)
* `auto_balance_strategy` - The auto balance strategy. Typically left blank or set to 'none'. (string)
* `custom_properties` - The custom properties defined for the Collector group ([]*NameAndValue)
    Provide custom properties in alphabetical order based on their property name.
  + `name` - The name of a property (required)
  + `value` - The value of a property (required)
* `description` - The description of the Collector Group (string)

## Import

collector groups can be imported using their collector group ID or name
```
$ terraform import logicmonitor_collector_group.my_collector_group 66
$ terraform import logicmonitor_collector_group.my_collector_group Collector_Group_(Network Devices)
```