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

### LogicMonitor Provider Credentials

```hcl
provider "logicmonitor" {
  api_id = var.logicmonitor_api_id
  api_key = var.logicmonitor_api_key
  company = var.logicmonitor_company
}
```

### AlertRule

```hcl
# create a new LogicMonitor alert rule
resource "logicmonitor_alert_rule" "my_alert_rule" {
        datapoint = "*"
        instance = "*"
        devices = [
        "Cisco Router"
        ]
        escalating_chain_id = 9
        priority = 100
        suppress_alert_ack_sdt = true
        datasource = "Port-"
        suppress_alert_clear = true
        name = "Alert Rule Test"
        level_str = "Critical"
        device_groups = [
        "Devices by Type"
        ]
        escalation_interval = 15
}
```

### Collector

```hcl
# create a new LogicMonitor collector
resource "logicmonitor_collector" "my_collector" {
      company = "lm"
      description: "Linux Collector"
}
```

### CollectorGroup

```hcl
# create a new LogicMonitor collector group
resource "logicmonitor_collector_group" "my_collector_group" {
      custom_properties = [
        {
          name = "addr"
              value = "127.0.0.1"
        },
        {
          name = "host"
              value = "localhost"
        }
      ]
      description = "Group for collectors dedicated to Network Devices"
      name = "Collector_Group_(Network Devices)"
}
```

### Dashboard

```hcl
# create a new LogicMonitor dashboard
resource "logicmonitor_dashboard" "my_dashboard" {
      	description = "my dashboard"
        name = "test_dashboard"
        sharable = true
}
```

### DashboardGroup

```hcl
# create a new LogicMonitor dashboard group
resource "logicmonitor_dashboard_group" "my_dashboard_group" {
        name = "LogicMonitor Dashboard Group"
        description  = "LM dashboard group testing"
}
```

### DataResourceAwsExternalID

```hcl
# create a new LogicMonitor data resource aws external id
resource "logicmonitor_data_resource_aws_external_id" "my_data_resource_aws_external_id" {
}
```

### Device

```hcl
# create a new LogicMonitor device
resource "logicmonitor_device" "my_device" {
      description = "This is a Cisco Router"
      display_name = "Cisco Router"
      name = "collector.localhost"
      preferred_collector_id = 2
}
```

### DeviceGroup

```hcl
# create a new LogicMonitor device group
resource "logicmonitor_device_group" "my_device_group" {
      description = "my device group"
      disable_alerting = false
      enable_netflow = false
      default_collector_id = 2
      name = "my device group"
      parent_id = 1
      custom_properties { 
          name = "addr"      
          value = "127.0.0.1" 
        }
      group_type = "Normal"
}
```

### EscalationChain

```hcl
# create a new LogicMonitor escalation chain
resource "logicmonitor_escalation_chain" "my_escalation_chain" {
        name = "LogicMonitor Escalation Chain"
        destinations = [
         {
           period = [{
            week_days = [2,3]
            timezone = "UTC"
            start_minutes = 10
            end_minutes   = 15
         }],
          stages = [{
            type    = "Admin"
            addr    = "unicornsparkles@rainbow.io"
            method  = "EMAIL"
            contact = "78362637"
         }]
          type = "timebased"
         }
        ]
        description  = "LM Escalation Chain testing"
        throttling_alerts = 40
        enable_throttling = true
        throttling_period = 30
}
```

### Website

```hcl
# create a new LogicMonitor website
resource "logicmonitor_website" "my_website" {
          type = "pingcheck"
          host = "google.com"
          overall_alert_level = "warn"
          polling_interval = 5
          description = "website test"
          disable_alerting = true
          stop_monitoring = true
          user_permission = "string"
          individual_sm_alert_enable = false
          steps = [
          {
          schema = "https"
          resp_type = "config"
          timeout = 0
          match_type = "plain"
          description = "string"
          use_default_root = true
          http_method = "GET"
          enable = true
          http_version = "1.1"
          follow_redirection = true
          post_data_edit_type = "raw"
          name = "string"
          req_type = "config"
          fullpage_load = false
          }
          ]
         transition = 1
         global_sm_alert_cond = 0
         is_internal = false
         domain = "www.ebay.com"
         name = "Ebay"
         use_default_location_setting = false
         use_default_alert_setting = true
         individual_alert_level = "warn"
}
```

### WebsiteGroup

```hcl
# create a new LogicMonitor website group
resource "logicmonitor_website_group" "my_website_group" {   
         stop_monitoring = false
         description = "website group test001"
         disable_alerting = false
         parent_id = 1
         name = "Amazon Website Checks"
         properties = [
         {
         name : "addr",
         value : "127.0.0.1"
         }  
         ]
}
```


## Argument Reference

The following arguments are supported:
* `api_id` - (Required) LogicMonitor API id. This can also be set via the `LM_API_ID` environment variable.
* `api_key` - (Required) LogicMonitor API key. This can also be set via the `LM_API_KEY` environment variable.
* `company` - (Required) LogicMonitor company name. This can also be set via the `LM_COMPANY` environment variable.
* `bulk_resource` - (Optional) True if going for bulk resources that can exceed the rate limit. The default value is false.