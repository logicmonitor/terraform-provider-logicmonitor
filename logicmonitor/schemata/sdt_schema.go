package schemata

import (
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-logicmonitor/models"
)

func SdtSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"admin": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"batch_job_name": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"collector_description": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"collector_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"comment": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"data_source_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"data_source_instance_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"data_source_instance_name": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"data_source_name": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"default_value": {
			Type:     schema.TypeString, //GoType: strfmt.DateTime
			Optional: true,
		},

		"device_batch_job_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"device_data_source_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"device_data_source_instance_group_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"device_data_source_instance_group_name": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"device_display_name": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"device_event_source_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"device_group_full_path": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"device_group_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"device_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"duration": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"end_date_time": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"end_date_time_on_local": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"end_hour": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"end_minute": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"event_source_name": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"hour": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"is_effective": {
			Type:     schema.TypeBool,
			Computed: true,
		},

		"minute": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"month_day": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"sdt_type": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"start_date_time": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"start_date_time_on_local": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"timezone": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"type": {
			Type:     schema.TypeString,
			Required: true,
		},

		"week_day": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"week_of_month": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

// Schema mapping representing the resource's respective datasource object defined in Terraform configuration
// Only difference between this and SdtSchema() are the computabilty of the id field and the inclusion of a filter field for datasources
func DataSourceSdtSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"admin": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"batch_job_name": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"collector_description": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"collector_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"comment": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"data_source_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"data_source_instance_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"data_source_instance_name": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"data_source_name": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"default_value": {
			Type:     schema.TypeString, //GoType: strfmt.DateTime
			Optional: true,
		},

		"device_batch_job_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"device_data_source_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"device_data_source_instance_group_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"device_data_source_instance_group_name": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"device_display_name": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"device_event_source_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"device_group_full_path": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"device_group_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"device_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"duration": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"end_date_time": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"end_date_time_on_local": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"end_hour": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"end_minute": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"event_source_name": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"hour": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"id": {
			Type:     schema.TypeInt,
			Computed: true,
			Optional: true,
		},

		"is_effective": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"minute": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"month_day": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"sdt_type": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"start_date_time": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"start_date_time_on_local": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"timezone": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"type": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"week_day": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"week_of_month": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"filter": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func SetSdtResourceData(d *schema.ResourceData, m *models.Sdt) {
	d.Set("admin", m.Admin)
	d.Set("batch_job_name", m.BatchJobName)
	d.Set("collector_description", m.CollectorDescription)
	d.Set("collector_id", m.CollectorID)
	d.Set("comment", m.Comment)
	d.Set("data_source_id", m.DataSourceID)
	d.Set("data_source_instance_id", m.DataSourceInstanceID)
	d.Set("data_source_instance_name", m.DataSourceInstanceName)
	d.Set("data_source_name", m.DataSourceName)
	d.Set("default_value", m.DefaultValue.String())
	d.Set("device_batch_job_id", m.DeviceBatchJobID)
	d.Set("device_data_source_id", m.DeviceDataSourceID)
	d.Set("device_data_source_instance_group_id", m.DeviceDataSourceInstanceGroupID)
	d.Set("device_data_source_instance_group_name", m.DeviceDataSourceInstanceGroupName)
	d.Set("device_display_name", m.DeviceDisplayName)
	d.Set("device_event_source_id", m.DeviceEventSourceID)
	d.Set("device_group_full_path", m.DeviceGroupFullPath)
	d.Set("device_group_id", m.DeviceGroupID)
	d.Set("device_id", m.DeviceID)
	d.Set("duration", m.Duration)
	d.Set("end_date_time", m.EndDateTime)
	d.Set("end_date_time_on_local", m.EndDateTimeOnLocal)
	d.Set("end_hour", m.EndHour)
	d.Set("end_minute", m.EndMinute)
	d.Set("event_source_name", m.EventSourceName)
	d.Set("hour", m.Hour)
	d.Set("id", m.ID)
	d.Set("is_effective", m.IsEffective)
	d.Set("minute", m.Minute)
	d.Set("month_day", m.MonthDay)
	d.Set("sdt_type", m.SdtType)
	d.Set("start_date_time", m.StartDateTime)
	d.Set("start_date_time_on_local", m.StartDateTimeOnLocal)
	d.Set("timezone", m.Timezone)
	d.Set("type", m.Type)
	d.Set("week_day", m.WeekDay)
	d.Set("week_of_month", m.WeekOfMonth)
}

func SetSdtSubResourceData(m []*models.Sdt) (d []*map[string]interface{}) {
	for _, sdt := range m {
		if sdt != nil {
			properties := make(map[string]interface{})
			properties["admin"] = sdt.Admin
			properties["batch_job_name"] = sdt.BatchJobName
			properties["collector_description"] = sdt.CollectorDescription
			properties["collector_id"] = sdt.CollectorID
			properties["comment"] = sdt.Comment
			properties["data_source_id"] = sdt.DataSourceID
			properties["data_source_instance_id"] = sdt.DataSourceInstanceID
			properties["data_source_instance_name"] = sdt.DataSourceInstanceName
			properties["data_source_name"] = sdt.DataSourceName
			properties["default_value"] = sdt.DefaultValue.String()
			properties["device_batch_job_id"] = sdt.DeviceBatchJobID
			properties["device_data_source_id"] = sdt.DeviceDataSourceID
			properties["device_data_source_instance_group_id"] = sdt.DeviceDataSourceInstanceGroupID
			properties["device_data_source_instance_group_name"] = sdt.DeviceDataSourceInstanceGroupName
			properties["device_display_name"] = sdt.DeviceDisplayName
			properties["device_event_source_id"] = sdt.DeviceEventSourceID
			properties["device_group_full_path"] = sdt.DeviceGroupFullPath
			properties["device_group_id"] = sdt.DeviceGroupID
			properties["device_id"] = sdt.DeviceID
			properties["duration"] = sdt.Duration
			properties["end_date_time"] = sdt.EndDateTime
			properties["end_date_time_on_local"] = sdt.EndDateTimeOnLocal
			properties["end_hour"] = sdt.EndHour
			properties["end_minute"] = sdt.EndMinute
			properties["event_source_name"] = sdt.EventSourceName
			properties["hour"] = sdt.Hour
			properties["id"] = sdt.ID
			properties["is_effective"] = sdt.IsEffective
			properties["minute"] = sdt.Minute
			properties["month_day"] = sdt.MonthDay
			properties["sdt_type"] = sdt.SdtType
			properties["start_date_time"] = sdt.StartDateTime
			properties["start_date_time_on_local"] = sdt.StartDateTimeOnLocal
			properties["timezone"] = sdt.Timezone
			properties["type"] = sdt.Type
			properties["week_day"] = sdt.WeekDay
			properties["week_of_month"] = sdt.WeekOfMonth
			d = append(d, &properties)
		}
	}
	return
}

func SdtModel(d *schema.ResourceData) *models.Sdt {
	batchJobName := d.Get("batch_job_name").(string)
	collectorID := int32(d.Get("collector_id").(int))
	comment := d.Get("comment").(string)
	dataSourceID := int32(d.Get("data_source_id").(int))
	dataSourceInstanceID := int32(d.Get("data_source_instance_id").(int))
	dataSourceInstanceName := d.Get("data_source_instance_name").(string)
	dataSourceName := d.Get("data_source_name").(string)
	defaultValueStr := d.Get("default_value").(string)
	defaultValue, _ := strfmt.ParseDateTime(defaultValueStr)
	deviceBatchJobID := int32(d.Get("device_batch_job_id").(int))
	deviceDataSourceID := int32(d.Get("device_data_source_id").(int))
	deviceDataSourceInstanceGroupID := int32(d.Get("device_data_source_instance_group_id").(int))
	deviceDataSourceInstanceGroupName := d.Get("device_data_source_instance_group_name").(string)
	deviceDisplayName := d.Get("device_display_name").(string)
	deviceEventSourceID := int32(d.Get("device_event_source_id").(int))
	deviceGroupFullPath := d.Get("device_group_full_path").(string)
	deviceGroupID := int32(d.Get("device_group_id").(int))
	deviceID := int32(d.Get("device_id").(int))
	duration := int32(d.Get("duration").(int))
	endDateTime := int64(d.Get("end_date_time").(int))
	endHour := int32(d.Get("end_hour").(int))
	endMinute := int32(d.Get("end_minute").(int))
	eventSourceName := d.Get("event_source_name").(string)
	hour := int32(d.Get("hour").(int))
	id := d.Get("id").(string)
	minute := int32(d.Get("minute").(int))
	monthDay := int32(d.Get("month_day").(int))
	sdtType := d.Get("sdt_type").(string)
	startDateTime := int64(d.Get("start_date_time").(int))
	timezone := d.Get("timezone").(string)
	typeVar := d.Get("type").(string)
	weekDay := d.Get("week_day").(string)
	weekOfMonth := d.Get("week_of_month").(string)

	return &models.Sdt{
		BatchJobName:                      batchJobName,
		CollectorID:                       collectorID,
		Comment:                           comment,
		DataSourceID:                      dataSourceID,
		DataSourceInstanceID:              dataSourceInstanceID,
		DataSourceInstanceName:            dataSourceInstanceName,
		DataSourceName:                    dataSourceName,
		DefaultValue:                      defaultValue,
		DeviceBatchJobID:                  deviceBatchJobID,
		DeviceDataSourceID:                deviceDataSourceID,
		DeviceDataSourceInstanceGroupID:   deviceDataSourceInstanceGroupID,
		DeviceDataSourceInstanceGroupName: deviceDataSourceInstanceGroupName,
		DeviceDisplayName:                 deviceDisplayName,
		DeviceEventSourceID:               deviceEventSourceID,
		DeviceGroupFullPath:               deviceGroupFullPath,
		DeviceGroupID:                     deviceGroupID,
		DeviceID:                          deviceID,
		Duration:                          duration,
		EndDateTime:                       endDateTime,
		EndHour:                           endHour,
		EndMinute:                         endMinute,
		EventSourceName:                   eventSourceName,
		Hour:                              hour,
		ID:                                id,
		Minute:                            minute,
		MonthDay:                          monthDay,
		SdtType:                           sdtType,
		StartDateTime:                     startDateTime,
		Timezone:                          timezone,
		Type:                              &typeVar,
		WeekDay:                           weekDay,
		WeekOfMonth:                       weekOfMonth,
	}
}
func GetSdtPropertyFields() (t []string) {
	return []string{
		"batch_job_name",
		"collector_id",
		"comment",
		"data_source_id",
		"data_source_instance_id",
		"data_source_instance_name",
		"data_source_name",
		"default_value",
		"device_batch_job_id",
		"device_data_source_id",
		"device_data_source_instance_group_id",
		"device_data_source_instance_group_name",
		"device_display_name",
		"device_event_source_id",
		"device_group_full_path",
		"device_group_id",
		"device_id",
		"duration",
		"end_date_time",
		"end_hour",
		"end_minute",
		"event_source_name",
		"hour",
		"id",
		"minute",
		"month_day",
		"sdt_type",
		"start_date_time",
		"timezone",
		"type",
		"week_day",
		"week_of_month",
	}
}
