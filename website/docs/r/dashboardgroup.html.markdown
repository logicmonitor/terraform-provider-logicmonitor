---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_dashboard_group"
sidebar_current: "docs-logicmonitor-dashboard-group"
description: |-
  Provides a LogicMonitor dashboard group resource. This can be used to create and manage LogicMonitor dashboard groups
---

# logicmonitor_dashboard_group

Provides a LogicMonitor dashboard group resource. This can be used to create and manage LogicMonitor dashboard groups.
Currently only creating a dashboard group based on an existing JSON template is supported

## Example Usage

```hcl
# Create a new LogicMonitor dashboard group
resource "logicmonitor_dashboard_group" "dashgrp" {
  name = "Fadeaway"
  description = "a new dashboard group"
  widget_tokens {
    "defaultResourceGroup" = "Lakers/SignatureMoves"
  }
}
```

```hcl
# Create a new LogicMonitor dashboard group based on a template stored in s3
resource "logicmonitor_dashboard_group" "dashgrp" {
  name = "Fadeaway"
  description = "a new dashboard group"
  widget_tokens {
    "defaultResourceGroup" = "Lakers/SignatureMoves"
  }
  template = "${data.aws_s3_bucket_object.template.body}"
}

data "aws_s3_bucket_object" "template" {
  bucket = "production-dashboards"
  key = "Ballers.json"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of dashboard
* `description` - (Optional) Description of dashboard
* `parent_id` - (Optional) The id of the parent group for this dashboard group
* `force_delete` - (Optional) Force delete the dashboard group
* `template` - (Optional) Defines if an existing exported JSON template is used to create dashboard group
* `widget_tokens` - (Optional) Dashboard tokens allow users to apply a single dashboard group template to different device or website groups simply by changing the tokens’ values.

## Import

Device Groups can be imported using their group id or full path

```
$ terraform import logicmonitor_dashboard_group.dashgrp 12
$ terraform import logicmonitor_dashboard_group.dashgrp LakersDash/Players
```
