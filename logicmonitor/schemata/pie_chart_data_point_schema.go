package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func PieChartDataPointSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"aggregate": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
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
		
		"glob_mode": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"instance_name": {
			Type: schema.TypeString,
			Required: true,
		},
		
		"name": {
			Type: schema.TypeString,
			Required: true,
		},
		
		"top10": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"type": {
			Type: schema.TypeString,
			Optional: true,
		},
		
	}
}

func SetPieChartDataPointSubResourceData(m []*models.PieChartDataPoint) (d []*map[string]interface{}) {
	for _, pieChartDataPoint := range m {
		if pieChartDataPoint != nil {
			properties := make(map[string]interface{})
			properties["aggregate"] = pieChartDataPoint.Aggregate
			properties["aggregate_function"] = pieChartDataPoint.AggregateFunction
			properties["data_point_id"] = pieChartDataPoint.DataPointID
			properties["data_point_name"] = pieChartDataPoint.DataPointName
			properties["data_source_full_name"] = pieChartDataPoint.DataSourceFullName
			properties["data_source_id"] = pieChartDataPoint.DataSourceID
			properties["device_display_name"] = pieChartDataPoint.DeviceDisplayName
			properties["device_group_full_path"] = pieChartDataPoint.DeviceGroupFullPath
			properties["glob_mode"] = pieChartDataPoint.GlobMode
			properties["instance_name"] = pieChartDataPoint.InstanceName
			properties["name"] = pieChartDataPoint.Name
			properties["top10"] = pieChartDataPoint.Top10
			properties["type"] = pieChartDataPoint.Type
			d = append(d, &properties)
		}
	}
	return
}

func PieChartDataPointModel(d map[string]interface{}) *models.PieChartDataPoint {
	// assume that the incoming map only contains the relevant resource data
	aggregate := d["aggregate"].(bool)
	aggregateFunction := d["aggregate_function"].(string)
	dataPointID := int32(d["data_point_id"].(int))
	dataPointName := d["data_point_name"].(string)
	dataSourceFullName := d["data_source_full_name"].(string)
	dataSourceID := int32(d["data_source_id"].(int))
	deviceDisplayName := d["device_display_name"].(string)
	deviceGroupFullPath := d["device_group_full_path"].(string)
	globMode := d["glob_mode"].(bool)
	instanceName := d["instance_name"].(string)
	name := d["name"].(string)
	top10 := d["top10"].(bool)
	typeVar := d["type"].(string)
	
	return &models.PieChartDataPoint {
		Aggregate: aggregate,
		AggregateFunction: aggregateFunction,
		DataPointID: dataPointID,
		DataPointName: dataPointName,
		DataSourceFullName: dataSourceFullName,
		DataSourceID: dataSourceID,
		DeviceDisplayName: &deviceDisplayName,
		DeviceGroupFullPath: &deviceGroupFullPath,
		GlobMode: globMode,
		InstanceName: &instanceName,
		Name: &name,
		Top10: top10,
		Type: typeVar,
	}
}

func GetPieChartDataPointPropertyFields() (t []string) {
	return []string{
		"aggregate",
		"aggregate_function",
		"data_point_id",
		"data_point_name",
		"data_source_full_name",
		"data_source_id",
		"device_display_name",
		"device_group_full_path",
		"glob_mode",
		"instance_name",
		"name",
		"top10",
		"type",
	}
}