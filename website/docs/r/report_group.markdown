---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_report_group"
sidebar_current: "docs-logicmonitor-resource-report-group"
description: |-
  Provides a LogicMonitor report group resource. This can be used to create and manage LogicMonitor report groups.
---

# logicmonitor_report_group

Provides a LogicMonitor report group resource. This can be used to create and manage LogicMonitor report groups.

## Example Usage
```hcl
# Create a LogicMonitor report group
resource "logicmonitor_report_group" "my_report_group" {
  description = "This is daily firewall report"
  name = "Firewall Reports"
}
```

## Argument Reference

The following arguments are **required**:
* `name` - The report group name
   (string)

The following arguments are **optional**:
* `description` - The report group description (string)

## Import

report groups can be imported using their report group ID or name
```
$ terraform import logicmonitor_report_group.my_report_group 66
$ terraform import logicmonitor_report_group.my_report_group Firewall Reports
```