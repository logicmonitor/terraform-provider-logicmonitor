package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func PieChartInfoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"counters": {
			Type: schema.TypeList, //GoType: []*Counter  
			Elem: &schema.Resource{
				Schema: CounterSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"data_points": {
			Type: schema.TypeList, //GoType: []*PieChartDataPoint  
			Elem: &schema.Resource{
				Schema: PieChartDataPointSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"group_remaining_as_others": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"hide_zero_percent_slices": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"max_slices_can_be_shown": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"pie_chart_items": {
			Type: schema.TypeList, //GoType: []*PieChartItem  
			Elem: &schema.Resource{
				Schema: PieChartItemSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Required: true,
		},
		
		"show_labels_and_lines_on_p_c": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"title": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"virtual_data_points": {
			Type: schema.TypeList, //GoType: []*VirtualDataPoint  
			Elem: &schema.Resource{
				Schema: VirtualDataPointSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
	}
}

func SetPieChartInfoSubResourceData(m []*models.PieChartInfo) (d []*map[string]interface{}) {
	for _, pieChartInfo := range m {
		if pieChartInfo != nil {
			properties := make(map[string]interface{})
			properties["counters"] = SetCounterSubResourceData(pieChartInfo.Counters)
			properties["data_points"] = SetPieChartDataPointSubResourceData(pieChartInfo.DataPoints)
			properties["group_remaining_as_others"] = pieChartInfo.GroupRemainingAsOthers
			properties["hide_zero_percent_slices"] = pieChartInfo.HideZeroPercentSlices
			properties["max_slices_can_be_shown"] = pieChartInfo.MaxSlicesCanBeShown
			properties["pie_chart_items"] = SetPieChartItemSubResourceData(pieChartInfo.PieChartItems)
			properties["show_labels_and_lines_on_p_c"] = pieChartInfo.ShowLabelsAndLinesOnPC
			properties["title"] = pieChartInfo.Title
			properties["virtual_data_points"] = SetVirtualDataPointSubResourceData(pieChartInfo.VirtualDataPoints)
			d = append(d, &properties)
		}
	}
	return
}

func PieChartInfoModel(d map[string]interface{}) *models.PieChartInfo {
	// assume that the incoming map only contains the relevant resource data
	counters := d["counters"].([]*models.Counter)
	dataPoints := d["data_points"].([]*models.PieChartDataPoint)
	groupRemainingAsOthers := d["group_remaining_as_others"].(bool)
	hideZeroPercentSlices := d["hide_zero_percent_slices"].(bool)
	maxSlicesCanBeShown := int32(d["max_slices_can_be_shown"].(int))
	pieChartItems := d["pie_chart_items"].([]*models.PieChartItem)
	showLabelsAndLinesOnPC := d["show_labels_and_lines_on_p_c"].(bool)
	title := d["title"].(string)
	virtualDataPoints := d["virtual_data_points"].([]*models.VirtualDataPoint)
	
	return &models.PieChartInfo {
		Counters: counters,
		DataPoints: dataPoints,
		GroupRemainingAsOthers: groupRemainingAsOthers,
		HideZeroPercentSlices: hideZeroPercentSlices,
		MaxSlicesCanBeShown: maxSlicesCanBeShown,
		PieChartItems: pieChartItems,
		ShowLabelsAndLinesOnPC: showLabelsAndLinesOnPC,
		Title: title,
		VirtualDataPoints: virtualDataPoints,
	}
}

func GetPieChartInfoPropertyFields() (t []string) {
	return []string{
		"counters",
		"data_points",
		"group_remaining_as_others",
		"hide_zero_percent_slices",
		"max_slices_can_be_shown",
		"pie_chart_items",
		"show_labels_and_lines_on_p_c",
		"title",
		"virtual_data_points",
	}
}