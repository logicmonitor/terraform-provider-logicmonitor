---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_datasource"
sidebar_current: "docs-logicmonitor-resource-datasource"
description: |-
  Provides a LogicMonitor datasource resource. This can be used to create and manage LogicMonitor datasources.
---

# logicmonitor_datasource

Provides a LogicMonitor datasource resource. This can be used to create and manage LogicMonitor datasources.

## Example Usage
```hcl
# Create a LogicMonitor datasource
resource "logicmonitor_datasource" "my_datasource" {
  applies_to = "system.deviceId == \"22\""
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
  collect_interval = 100
  collect_method = "script"
  collector_attribute {
  name = "script"
  }
  data_points = [
   {
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

   }
  ]
  description = "string"
  display_name = "test"
  enable_auto_discovery = false
  enable_eri_discovery = false
  eri_discovery_config {
    name = "ad_script"
    win_script = "string"
    groovy_script = "string"
    type = "string"
    linux_cmdline = "string"
    linux_script = "string"
    win_cmdline = "string"
  }
  eri_discovery_interval = 10
  group = "string"
  has_multi_instances = 
  name = "datasource test"
  tags = "string"
  technology = "string"
}
```

## Argument Reference

The following arguments are **required**:
* `collect_interval` - The DataSource data collect interval
   (int32)
* `collect_method` - The  method to collect data. The values can be snmp|ping|exs|webpage|wmi|cim|datadump|dns|ipmi|jdbb|script|udp|tcp|xen
   (string)
* `collector_attribute` - Data collector's attributes to collector data. e.g. a ping data source has a ping collector attribute. 
 PingCollectorAttributeV1 has two fields. the ip to ping, the data size send to ping
   (CollectorAttribute)
  + `name` - The data collectors name
* `name` - The data source name
   (string)

The following arguments are **optional**:
* `applies_to` - The Applies To for the LMModule (string)
* `auto_discovery_config` -  (AutoDiscoveryConfiguration)
* `data_points` - The data point list ([]*DataPoint)
  + `alertForNoData` - The triggered alert level if we cannot collect data for this datapoint. The values can be 0-4 (0:unused alert, 1:alert ok, 2:warn alert, 2:error alert, 4:critical alert)
  + `postProcessorParam` - The post processor parameter, e.g. dataPoint1*2
  + `postProcessorMethod` - The post processor method for the data value. Currently support complex expression and groovy.
  + `maxDigits` - The max digits of the data value
  + `rawDataFieldName` - The name of the raw data field name used to fetch value, e.g. avgrtt, output
  + `description` - The datapoint description
  + `alertClearTransitionInterval` - The count that the alert must exist for this many poll cycles before the alert will be cleared
  + `type` - The data metric type. The values can be 0-7 (0:unknown, 1:counter, 2:gauge, 3:derive, 5:status, 6:compute, 7:counter32, 8:counter64)
  + `minValue` - The minimum value of the datapoint value range
  + `alertBody` - The customized alert message body define.  Empty string mean we will use the define in default template
  + `alertSubject` - The customized alert message subject define. Empty string mean we will use the define in default templateds
  + `alertTransitionInterval` - The count that the alert must exist for this many poll cycles before it will be triggered
  + `maxValue` - The max value of the datapoint value range
  + `dataType` - The data value type. The values can be 1-8 (1:boolean, 2:byte, 3:short, 4:int, 5:long, 6:float, 7:double, 8:ulong)
  + `alertExprNote` - alert expression note
  + `name` - The datapoint name
  + `alertExpr` - The alert threshold define for the datapoint. e.g.  60 80 90 mean it will: trigger warn alert if value  60 trigger error alert if value  80 trigger critical alert if value  90
* `description` - The description for the LMModule (string)
* `display_name` - The data source display name (string)
* `enable_auto_discovery` - Enable Auto Discovery or not when this data source has multi instance: false|true (bool)
* `enable_eri_discovery` - Enable ERI Discovery or not: false|true (bool)
* `eri_discovery_config` -  (ScriptERIDiscoveryAttributeV2)
* `eri_discovery_interval` - The DataSource data collect interval (int32)
* `group` - The group the LMModule is in (string)
* `has_multi_instances` - If the DataSource has multi instance: true|false (bool)
* `tags` - The Tags for the LMModule (string)
* `technology` - The Technical Notes for the LMModule (string)

## Import

datasources can be imported using their datasource ID or name
```
$ terraform import logicmonitor_datasource.my_datasource 66
$ terraform import logicmonitor_datasource.my_datasource datasource test
```