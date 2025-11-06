---
layout: "logicmonitor"
page_title: "LogicMonitor: logicmonitor_sdt"
sidebar_current: "docs-logicmonitor-resource-sdt"
description: |-
  Provides a LogicMonitor sdt resource. This can be used to create and manage LogicMonitor sdts.
---

# logicmonitor_sdt

Provides a LogicMonitor sdt resource. This can be used to create and manage LogicMonitor sdts.

## Example Usage
```hcl
# Create a LogicMonitor sdt
resource "logicmonitor_sdt" "my_sdt" {
  batch_job_name = ""
  collector_id = 
  comment = "Emergency prod deployment"
  data_source_id = 
  data_source_instance_id = 
  data_source_instance_name = ""
  data_source_name = ""
  default_value = 
  device_batch_job_id = 
  device_data_source_id = 
  device_data_source_instance_group_id = 
  device_data_source_instance_group_name = ""
  device_display_name = ""
  device_event_source_id = 
  device_group_full_path = ""
  device_group_id = 
  device_id = 
  duration = 138
  end_date_time = 1534554000000
  end_hour = 5
  end_minute = 18
  event_source_name = ""
  hour = 3
  minute = 6
  month_day = 7
  sdt_type = "oneTime"
  start_date_time = 1534460400000
  timezone = "America/Los_Angeles"
  type = "ResourceGroupSDT"
  week_day = "Sunday"
  week_of_month = "1"
}
```

## Argument Reference

The following arguments are **required**:
* `type` - The type of resource that this SDT is for. The values can be CollectorSDT | DeviceDataSourceInstanceSDT | DeviceBatchJobSDT | DeviceClusterAlertDefSDT | DeviceDataSourceInstanceGroupSDT | DeviceDataSourceSDT | DeviceEventSourceSDT | ResourceGroupSDT | ResourceSDT | WebsiteCheckpointSDT | WebsiteGroupSDT | WebsiteSDT | DeviceLogPipeLineResourceSDT
   (string)

The following arguments are **optional**:
* `batch_job_name` - The name of the batchjob that the SDT will apply to (string)
* `collector_id` - The id of the collector that the SDT will be associated with (int32)
* `comment` - The notes associated with the SDT (string)
* `data_source_id` - The id of the datasource that this SDT will be associated with, for the specified group. dataSourceId 0 indicates all datasources (int32)
* `data_source_instance_id` - The id of the datasource instance that the SDT will be associated with (int32)
* `data_source_instance_name` - The name of the datasource instance that the SDT will be associated with (string)
* `data_source_name` - The name of the datasource that this SDT will be associated with, for the specified group. dataSourceName "All" indicates all datasources (string)
* `default_value` -  (strfmt.DateTime)
* `device_batch_job_id` - The id of the device batchjob that the SDT will be associated with (int32)
* `device_data_source_id` - The id of the device datasource that the SDT will be associated with (int32)
* `device_data_source_instance_group_id` - The name of the device datasource instance group that the SDT will be associated with (int32)
* `device_data_source_instance_group_name` - The name of the instance group (string)
* `device_display_name` - The name of the device that this SDT will be associated with (string)
* `device_event_source_id` - The id of the device eventsource that the SDT will be associated with (int32)
* `device_group_full_path` - The full path of the device group that this SDT will be associated with (string)
* `device_group_id` - The id of the device group that the SDT will be associated with (int32)
* `device_id` - The id of the device that the SDT will be associated with (int32)
* `duration` - The duration of the SDT in minutes (int32)
* `end_date_time` - The epoch time, in milliseconds, that the SDT will end (int64)
* `end_hour` - The values can be 1 | 2....| 24. Specifies the hour that the SDT ends for a repeating SDT (int32)
* `end_minute` - The values can be 1 | 2....| 60. Specifies the minute of the hour that the SDT ends for a repeating SDT (int32)
* `event_source_name` - The name of the eventsource that the SDT will apply to (string)
* `hour` - The values can be 1 | 2....| 24. Specifies the hour that the SDT will start for a repeating SDT (daily, weekly, or monthly) (int32)
* `minute` - The values can be 1 | 2....| 60. Specifies the minute of the hour that the SDT should begin for a repeating SDT (int32)
* `month_day` - The values can be 1 | 2....| 31. Specifies the day of the month that the SDT will be active for a monthly SDT (int32)
* `sdt_type` - The type of sdt. The values can be oneTime|weekly|monthly|daily|monthlyByWeek (string)
* `start_date_time` - The epoch time, in milliseconds, that the SDT will start (int64)
* `timezone` - The specific timezone for SDT (string)
* `week_day` - The week day of sdt. The values can be SUNDAY|MONDAY|TUESDAY|WEDNESDAY|THURSDAY|FRIDAY|SATURDAY (string)
* `week_of_month` - The week of the month that the SDT will be active for a monthly SDT (string)

## Import

sdts can be imported using their sdt ID or name
```
$ terraform import logicmonitor_sdt.my_sdt 66
$ terraform import logicmonitor_sdt.my_sdt
```