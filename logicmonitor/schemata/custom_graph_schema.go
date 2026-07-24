package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func CustomGraphSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"aggregate": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"data_points": {
			Type: schema.TypeList, //GoType: []*CustomFlexibleVirtualDataSourceEx  
			Elem: &schema.Resource{
				Schema: CustomFlexibleVirtualDataSourceExSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Required: true,
		},
		
		"desc": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"global_consolidate_function": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"max_value": {
			Type: schema.TypeMap, //GoType: interface{}
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},
		
		"min_value": {
			Type: schema.TypeMap, //GoType: interface{}
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},
		
		"scale_unit": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"top_x": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"vertical_label": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"virtual_data_points": {
			Type: schema.TypeList, //GoType: []*CustomVirtualDataPoint  
			Elem: &schema.Resource{
				Schema: CustomVirtualDataPointSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
	}
}

func SetCustomGraphSubResourceData(m []*models.CustomGraph) (d []*map[string]interface{}) {
	for _, customGraph := range m {
		if customGraph != nil {
			properties := make(map[string]interface{})
			properties["aggregate"] = customGraph.Aggregate
			properties["data_points"] = SetCustomFlexibleVirtualDataSourceExSubResourceData(customGraph.DataPoints)
			properties["desc"] = customGraph.Desc
			properties["global_consolidate_function"] = customGraph.GlobalConsolidateFunction
			properties["id"] = customGraph.ID
			properties["max_value"] = customGraph.MaxValue
			properties["min_value"] = customGraph.MinValue
			properties["scale_unit"] = customGraph.ScaleUnit
			properties["top_x"] = customGraph.TopX
			properties["vertical_label"] = customGraph.VerticalLabel
			properties["virtual_data_points"] = SetCustomVirtualDataPointSubResourceData(customGraph.VirtualDataPoints)
			d = append(d, &properties)
		}
	}
	return
}

func CustomGraphModel(d map[string]interface{}) *models.CustomGraph {
	// assume that the incoming map only contains the relevant resource data
	aggregate := d["aggregate"].(bool)
	dataPoints := d["data_points"].([]*models.CustomFlexibleVirtualDataSourceEx)
	desc := d["desc"].(bool)
	globalConsolidateFunction := d["global_consolidate_function"].(string)
	id := int32(d["id"].(int))
	maxValue := d["max_value"]
	minValue := d["min_value"]
	scaleUnit := int32(d["scale_unit"].(int))
	topX := int32(d["top_x"].(int))
	verticalLabel := d["vertical_label"].(string)
	virtualDataPoints := d["virtual_data_points"].([]*models.CustomVirtualDataPoint)
	
	return &models.CustomGraph {
		Aggregate: aggregate,
		DataPoints: dataPoints,
		Desc: desc,
		GlobalConsolidateFunction: globalConsolidateFunction,
		ID: id,
		MaxValue: maxValue,
		MinValue: minValue,
		ScaleUnit: scaleUnit,
		TopX: topX,
		VerticalLabel: verticalLabel,
		VirtualDataPoints: virtualDataPoints,
	}
}

func GetCustomGraphPropertyFields() (t []string) {
	return []string{
		"aggregate",
		"data_points",
		"desc",
		"global_consolidate_function",
		"id",
		"max_value",
		"min_value",
		"scale_unit",
		"top_x",
		"vertical_label",
		"virtual_data_points",
	}
}