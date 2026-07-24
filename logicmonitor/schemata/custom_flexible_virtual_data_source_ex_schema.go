package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func CustomFlexibleVirtualDataSourceExSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"aggregate_function": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"consolidate_function": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"custom_graph_id": {
			Type: schema.TypeInt,
			Computed: true,
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
			Type: schema.TypeList, //GoType: GlobMatchToggle
			Elem: &schema.Resource{
				Schema: GlobMatchToggleSchema(),
			},
			Required: true,
		},
		
		"device_group_full_path": {
			Type: schema.TypeList, //GoType: GlobMatchToggle
			Elem: &schema.Resource{
				Schema: GlobMatchToggleSchema(),
			},
			Required: true,
		},
		
		"display": {
			Type: schema.TypeList, //GoType: GraphDisplay
			Elem: &schema.Resource{
				Schema: GraphDisplaySchema(),
			},
			Required: true,
		},
		
		"id": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"instance_name": {
			Type: schema.TypeList, //GoType: GlobMatchToggle
			Elem: &schema.Resource{
				Schema: GlobMatchToggleSchema(),
			},
			Required: true,
		},
		
		"is_valid_exp": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"is_virtual_datapoint": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"name": {
			Type: schema.TypeString,
			Required: true,
		},
		
	}
}

func SetCustomFlexibleVirtualDataSourceExSubResourceData(m []*models.CustomFlexibleVirtualDataSourceEx) (d []*map[string]interface{}) {
	for _, customFlexibleVirtualDataSourceEx := range m {
		if customFlexibleVirtualDataSourceEx != nil {
			properties := make(map[string]interface{})
			properties["aggregate_function"] = customFlexibleVirtualDataSourceEx.AggregateFunction
			properties["consolidate_function"] = customFlexibleVirtualDataSourceEx.ConsolidateFunction
			properties["custom_graph_id"] = customFlexibleVirtualDataSourceEx.CustomGraphID
			properties["data_point_id"] = customFlexibleVirtualDataSourceEx.DataPointID
			properties["data_point_name"] = customFlexibleVirtualDataSourceEx.DataPointName
			properties["data_source_full_name"] = customFlexibleVirtualDataSourceEx.DataSourceFullName
			properties["data_source_id"] = customFlexibleVirtualDataSourceEx.DataSourceID
			properties["device_display_name"] = SetGlobMatchToggleSubResourceData([]*models.GlobMatchToggle{customFlexibleVirtualDataSourceEx.DeviceDisplayName})
			properties["device_group_full_path"] = SetGlobMatchToggleSubResourceData([]*models.GlobMatchToggle{customFlexibleVirtualDataSourceEx.DeviceGroupFullPath})
			properties["display"] = SetGraphDisplaySubResourceData([]*models.GraphDisplay{customFlexibleVirtualDataSourceEx.Display})
			properties["id"] = customFlexibleVirtualDataSourceEx.ID
			properties["instance_name"] = SetGlobMatchToggleSubResourceData([]*models.GlobMatchToggle{customFlexibleVirtualDataSourceEx.InstanceName})
			properties["is_valid_exp"] = customFlexibleVirtualDataSourceEx.IsValidExp
			properties["is_virtual_datapoint"] = customFlexibleVirtualDataSourceEx.IsVirtualDatapoint
			properties["name"] = customFlexibleVirtualDataSourceEx.Name
			d = append(d, &properties)
		}
	}
	return
}

func CustomFlexibleVirtualDataSourceExModel(d map[string]interface{}) *models.CustomFlexibleVirtualDataSourceEx {
	// assume that the incoming map only contains the relevant resource data
	aggregateFunction := d["aggregate_function"].(string)
	consolidateFunction := d["consolidate_function"].(string)
	dataPointID := int32(d["data_point_id"].(int))
	dataPointName := d["data_point_name"].(string)
	dataSourceFullName := d["data_source_full_name"].(string)
	dataSourceID := int32(d["data_source_id"].(int))
	var deviceDisplayName *models.GlobMatchToggle = nil
	deviceDisplayNameList := d["device_display_name"].([]interface{})
	if len(deviceDisplayNameList) > 0 { // len(nil) = 0
		deviceDisplayName = GlobMatchToggleModel(deviceDisplayNameList[0].(map[string]interface{}))
	}
	var deviceGroupFullPath *models.GlobMatchToggle = nil
	deviceGroupFullPathList := d["device_group_full_path"].([]interface{})
	if len(deviceGroupFullPathList) > 0 { // len(nil) = 0
		deviceGroupFullPath = GlobMatchToggleModel(deviceGroupFullPathList[0].(map[string]interface{}))
	}
	var display *models.GraphDisplay = nil
	displayList := d["display"].([]interface{})
	if len(displayList) > 0 { // len(nil) = 0
		display = GraphDisplayModel(displayList[0].(map[string]interface{}))
	}
	id := int32(d["id"].(int))
	var instanceName *models.GlobMatchToggle = nil
	instanceNameList := d["instance_name"].([]interface{})
	if len(instanceNameList) > 0 { // len(nil) = 0
		instanceName = GlobMatchToggleModel(instanceNameList[0].(map[string]interface{}))
	}
	isValidExp := d["is_valid_exp"].(bool)
	isVirtualDatapoint := d["is_virtual_datapoint"].(bool)
	name := d["name"].(string)
	
	return &models.CustomFlexibleVirtualDataSourceEx {
		AggregateFunction: aggregateFunction,
		ConsolidateFunction: consolidateFunction,
		DataPointID: dataPointID,
		DataPointName: dataPointName,
		DataSourceFullName: dataSourceFullName,
		DataSourceID: dataSourceID,
		DeviceDisplayName: deviceDisplayName,
		DeviceGroupFullPath: deviceGroupFullPath,
		Display: display,
		ID: id,
		InstanceName: instanceName,
		IsValidExp: isValidExp,
		IsVirtualDatapoint: isVirtualDatapoint,
		Name: &name,
	}
}

func GetCustomFlexibleVirtualDataSourceExPropertyFields() (t []string) {
	return []string{
		"aggregate_function",
		"consolidate_function",
		"data_point_id",
		"data_point_name",
		"data_source_full_name",
		"data_source_id",
		"device_display_name",
		"device_group_full_path",
		"display",
		"id",
		"instance_name",
		"is_valid_exp",
		"is_virtual_datapoint",
		"name",
	}
}