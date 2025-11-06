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
        resource_properties = [
         {
           name  = "test.pass"
           value = "string"
          }
         ]
        send_anomaly_suppressed_alert = true
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
        widget_tokens = [
        {
         name  = "defaultDeviceGroup"
         value = "*"
         inherit_list = null
         type = null
        }
      ]
}
```

### DashboardGroup

```hcl
# create a new LogicMonitor dashboard group
resource "logicmonitor_dashboard_group" "my_dashboard_group" {
        name = "LogicMonitor Dashboard Group"
        description  = "LM dashboard group testing"
        widget_tokens = [
        {
         name  = "defaultDeviceGroup"
         value = "*"
         inherit_list = null
         type = null
        }
       ]
}
```

### DataResourceAwsExternalID

```hcl
# create a new LogicMonitor data resource aws external id
resource "logicmonitor_data_resource_aws_external_id" "my_data_resource_aws_external_id" {
}
```

### Datasource

```hcl
# create a new LogicMonitor datasource
resource "logicmonitor_datasource" "my_datasource" { 
        collect_interval = 100
        has_multi_instances = true
        applies_to = "system.deviceId == \"22\""
        description = "test"
        collect_method = "script"
        eri_discovery_interval = 15
        enable_auto_discovery = true
        enable_eri_discovery = true
        eri_discovery_config {
        name = "ad_script"
        win_script = "string"
        groovy_script = "string"
        type = "string"
        linux_cmdline = "string"
        linux_script = "string"
        win_cmdline = "string"
        }
        name = "Amazon Website test"
        auto_discovery_config {
        persistent_instance = false
        schedule_interval = 0
        delete_inactive_instance = true
        method {
        name = "ad_script"
        }
       instance_auto_group_method = "none"
       instance_auto_group_method_params = ""
       filters = [
       {
        comment = "test"
        value = "test"
        operation = "string"
        attribute = "string"
       }]
      disable_instance = true
      }
      data_points = [{
          name = "CallCountTotal_mean8"
          description = "test"
          alert_for_no_data = 1
          max_digits = 0
          alert_clear_transition_interval = 0
          alert_transition_interval = 0
          data_type = 7
          max_value = ""
          min_value = ""
          alert_body = "string"
          alert_subject = "string"
          alert_expr = "string"
          alert_expr_note = "string"
          type = 2
          raw_data_field_name = "string"
          post_processor_method = "aggregation"
          post_processor_param = "{\"version\":\"1.0\",\"expression\":{\"funcName\":\"mean\",\"dataSourceName\":\"AWS_Cognito_GlobalAPICallStats\",\"dataPointName\":\"CallCountTotal\"},\"dataLack\":\"ignore\"}"
      }]
    display_name = "Testdemo"
    collector_attribute {
    name = "script"
    }
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
          stages = [
          [{
          type    = "Admin"
          addr    = "unicornsparkles@rainbow.io"
          method  = "EMAIL"
          contact = "78362637"
          },
          {
          type    = "Admin"
          addr    = "unicornsparkles@rainbow.io"
          method  = "EMAIL"
          contact = "78362637"
          }],
          [{
          type    = "Admin"
          addr    = "unicornsparkles@rainbow.io"
          method  = "EMAIL"
          contact = "78362637"
          }]
        ]
          type = "timebased"
         }
        ]
        description  = "LM Escalation Chain testing"
        cc_destinations = [
         {
         method = "EMAIL"
         contact = "string"
         type = "Admin"
         addr = "unicornsparkles@rainbow.io"
         }
        ]
        throttling_alerts = 40
        enable_throttling = true
        throttling_period = 30
}
```

### Sdt

```hcl
# create a new LogicMonitor sdt
resource "logicmonitor_sdt" "my_sdt" {
     type = "CollectorSDT"
     comment = "Scheduled maintenance window for system updates"
     sdt_type = "oneTime"
     start_date_time = 1767225600000  # November 30, 2025 00:00:00 UTC
     end_date_time = 1767232800000    # November 30, 2025 01:00:00 UTC  
     duration = 60
     timezone = "America/New_York"
     default_value = "2025-01-01T00:00:00Z"
     collector_id = "30329"
     #device_id = "3087872"
     #device_data_source_id = "25268320"
     #data_source_id = "552560846"
     #device_group_full_path = "DEVTS-21343 Dynamic Group"
     #device_group_id = "275527"
     #device_data_source_id = "24195849"
     #data_source_instance_id = "368309065"
     #device_data_source_instance_group_id = "27132393"
     #device_event_source_id = "2333036" 
     #device_batch_job_id = "64369"
}
```

### Website

```hcl
# create a new LogicMonitor website
resource "logicmonitor_website" "my_website" {
          type = "pingcheck"
          host = "google.com"
          overall_alert_level = "warn"
          schema = "http"
          polling_interval = 5
          description = "website test"
          disable_alerting = true
          stop_monitoring = true
          user_permission = "string"
          test_location = [
          {
          all = false
          collector_ids = [1, 2, 3]
          collectors = null
          smg_ids = [85, 90]
          }
          ]
          group_id = 8
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
          http_headers = "X-Version:3"
          follow_redirection = true
          post_data_edit_type = "raw"
          name = "string"
          req_type = "config"
          fullpage_load = false
          require_auth = var.require_auth
          auth = var.require_auth ? [{
            password  = "string"
            type      = "basic"
            domain    = "string"
            user_name = "string"
           }] : []
           path = "string"
           keyword = "DEVWRT-SANRT-JOB1-9127"
           http_body = "string"
           resp_script = "string"
           req_script = "string"
           label = "string"
           url = "/rest/version"
           type = "string"
           invert_match = false
           status_code = "200"
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
         trigger_s_s_l_expiration_alert = true
         trigger_s_s_l_status_alert = true
         ignore_s_s_l = false
         alert_expr = "< 100"
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
* `domain` - The company domain. Defaults to logicmonitor.com. For gov companies, specify the appropriate domain value (e.g., qa-lmgov.us).