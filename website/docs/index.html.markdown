---
layout: "logicmonitor"
page_title: "Provider: LogicMonitor"
sidebar_current: "docs-logicmonitor-index"
description: |-
  The LogicMonitor provider is used to interact with the resources supported by LogicMonitor. The provider needs to be configured with the proper credentials before it can be used.
---

# LogicMonitor Provider

The LogicMonitor provider is used to interact with the resources supported by LogicMonitor. The provider needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
provider "logicmonitor" {
  api_id = "${var.logicmonitor_api_id}"
  api_key = "${var.logicmonitor_api_key}"
  company = "${var.logicmonitor_company}"
}

#
resource "logicmonitor_device" "host" {
  ip_addr = "10.32.12.18"
  disable_alerting = true
  collector = "${data.logicmonitor_collectors.collectors.id}"
  hostgroup_id = "${logicmonitor_device_group.group1.id}"
  properties {
   "app" = "haproxy"
   "system.categories" = "a,b,c,d"
  }
}

resource "logicmonitor_device_group" "group1" {
    name = "NewGroup"
    properties {
     "jmx.port" = "9003"
     "system.categories" = "ec2"
    }
}

data "logicmonitor_collectors" "collectors" {
  most_recent = true
}

```

## Argument Reference

The following arguments are supported:

* `api_id` - (Required) LogicMonitor API id. This can also be set via the `LM_API_ID` environment variable.
* `api_key` - (Required) LogicMonitor API key. This can also be set via the `LM_API_KEY` environment variable.
* `company` - (Required) LogicMonitor company name. This can also be set via the `LM_COMPANY` environment variable.
