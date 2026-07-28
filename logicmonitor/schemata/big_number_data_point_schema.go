package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func BigNumberDataPointSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"aggregate_function": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"data_point_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"data_point_name": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"data_source_full_name": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"data_source_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"device_display_name": {
			Type: schema.TypeString,
			Required: true,
		},
		
		"device_group_full_path": {
			Type: schema.TypeString,
			Required: true,
		},
		
		"instance_name": {
			Type: schema.TypeString,
			Required: true,
		},
		
		"name": {
			Type: schema.TypeString,
			Required: true,
		},
		
		"type": {
			Type: schema.TypeString,
			Optional: true,
		},
		
	}
}

func SetBigNumberDataPointSubResourceData(m []*models.BigNumberDataPoint) (d []*map[string]interface{}) {
	for _, bigNumberDataPoint := range m {
		if bigNumberDataPoint != nil {
			properties := make(map[string]interface{})
			properties["aggregate_function"] = bigNumberDataPoint.AggregateFunction
			properties["data_point_id"] = bigNumberDataPoint.DataPointID
			properties["data_point_name"] = bigNumberDataPoint.DataPointName
			properties["data_source_full_name"] = bigNumberDataPoint.DataSourceFullName
			properties["data_source_id"] = bigNumberDataPoint.DataSourceID
			properties["device_display_name"] = bigNumberDataPoint.DeviceDisplayName
			properties["device_group_full_path"] = bigNumberDataPoint.DeviceGroupFullPath
			properties["instance_name"] = bigNumberDataPoint.InstanceName
			properties["name"] = bigNumberDataPoint.Name
			properties["type"] = bigNumberDataPoint.Type
			d = append(d, &properties)
		}
	}
	return
}

func BigNumberDataPointModel(d map[string]interface{}) *models.BigNumberDataPoint {
	// assume that the incoming map only contains the relevant resource data
	aggregateFunction := d["aggregate_function"].(string)
	dataPointID := int32(d["data_point_id"].(int))
	dataPointName := d["data_point_name"].(string)
	dataSourceFullName := d["data_source_full_name"].(string)
	dataSourceID := int32(d["data_source_id"].(int))
	deviceDisplayName := d["device_display_name"].(string)
	deviceGroupFullPath := d["device_group_full_path"].(string)
	instanceName := d["instance_name"].(string)
	name := d["name"].(string)
	typeVar := d["type"].(string)
	
	return &models.BigNumberDataPoint {
		AggregateFunction: aggregateFunction,
		DataPointID: dataPointID,
		DataPointName: dataPointName,
		DataSourceFullName: dataSourceFullName,
		DataSourceID: dataSourceID,
		DeviceDisplayName: &deviceDisplayName,
		DeviceGroupFullPath: &deviceGroupFullPath,
		InstanceName: &instanceName,
		Name: &name,
		Type: typeVar,
	}
}

func GetBigNumberDataPointPropertyFields() (t []string) {
	return []string{
		"aggregate_function",
		"data_point_id",
		"data_point_name",
		"data_source_full_name",
		"data_source_id",
		"device_display_name",
		"device_group_full_path",
		"instance_name",
		"name",
		"type",
	}
}