---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_collector_group"
sidebar_current: "docs-logicmonitor-resource-collector-group"
description: |-
  Provides a LogicMonitor collector group resource. This can be used to create and manage LogicMonitor collector groups
---

# logicmonitor_collector_group

Provides a LogicMonitor collector group resource. This can be used to create and manage LogicMonitor collector groups

## Example Usage

```hcl
# Create a new LogicMonitor collector group
resource "logicmonitor_collector_group" "group1" {
  name  = "collector_group_1"
  description = "a new test group"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of collector group
* `description` - (Optional) Set description of collector group
