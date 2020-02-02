---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_dashboard"
sidebar_current: "docs-logicmonitor-dashboard"
description: |-
  Provides a LogicMonitor dashboard resource. This can be used to create and manage LogicMonitor dashboards
---

# logicmonitor_dashboard

Provides a LogicMonitor dashboard resource. This can be used to create and manage LogicMonitor dashboards
Currently only creating a dashboard based on an existing JSON template is supported

## Example Usage

```hcl
# Create a new LogicMonitor dashboard
resource "logicmonitor_dashboard" "dash" {
  name = "FiveRings"
  description = "a new dashboard"
  widget_tokens {
    "defaultResourceGroup" = "Lakers/Championships"
  }
}
```

```hcl
# Create a new LogicMonitor dashboard based on a template stored in s3
resource "logicmonitor_dashboard" "dash" {
  name = "FiveRings"
  description = "a new dashboard"
  widget_tokens {
    "defaultResourceGroup" = "Lakers/Championships"
  }
  template = "${data.aws_s3_bucket_object.template.body}"
}

data "aws_s3_bucket_object" "template" {
  bucket = "basketball-dashboards"
  key = "Ballers.json"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of dashboard
* `description` - (Optional) Description of dashboard
* `group_id` - (Optional) The id of the parent group for this dashboard
* `public` - (Optional) Defines if it is a public or private dashboard.
* `template` - (Optional) Defines if an existing exported JSON template is used to create dashboard
* `widget_tokens` - (Optional) Dashboard tokens allow users to apply a single dashboard template to different device or website groups simply by changing the tokens’ values.

## Import

Device Groups can be imported using their group id or full path

```
$ terraform import logicmonitor_dashboard.dash 12
$ terraform import logicmonitor_dashboard.dash LakersDash
```
