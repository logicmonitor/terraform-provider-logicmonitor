---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_collectorgroup"
sidebar_current: "docs-logicmonitor-resource-collectorgroup"
description: |-
  Provides a LogicMonitor collector group resource. This can be used to create and manage LogicMonitor users
---

# logicmonitor_collectorgroup

Provides a LogicMonitor collector group resource. This can be used to create and manage LogicMonitor users

## Example Usage

```hcl
# Create a new LogicMonitor collector group
resource "logicmonitor_collectorgroup" "group1" {
  name  = "collector_group_1"
  description = "a new test group"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Set description of collector group
* `description` - (Required) Name of collector group
