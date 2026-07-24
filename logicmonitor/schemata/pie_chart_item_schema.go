package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func PieChartItemSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"color": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"data_point_name": {
			Type: schema.TypeString,
			Required: true,
		},
		
		"legend": {
			Type: schema.TypeString,
			Required: true,
		},
		
	}
}

func SetPieChartItemSubResourceData(m []*models.PieChartItem) (d []*map[string]interface{}) {
	for _, pieChartItem := range m {
		if pieChartItem != nil {
			properties := make(map[string]interface{})
			properties["color"] = pieChartItem.Color
			properties["data_point_name"] = pieChartItem.DataPointName
			properties["legend"] = pieChartItem.Legend
			d = append(d, &properties)
		}
	}
	return
}

func PieChartItemModel(d map[string]interface{}) *models.PieChartItem {
	// assume that the incoming map only contains the relevant resource data
	color := d["color"].(string)
	dataPointName := d["data_point_name"].(string)
	legend := d["legend"].(string)
	
	return &models.PieChartItem {
		Color: color,
		DataPointName: &dataPointName,
		Legend: &legend,
	}
}

func GetPieChartItemPropertyFields() (t []string) {
	return []string{
		"color",
		"data_point_name",
		"legend",
	}
}