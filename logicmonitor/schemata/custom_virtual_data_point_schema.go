package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func CustomVirtualDataPointSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"custom_graph_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"display": {
			Type: schema.TypeList, //GoType: GraphDisplay
			Elem: &schema.Resource{
				Schema: GraphDisplaySchema(),
			},
			Optional: true,
		},
		
		"id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"name": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"rpn": {
			Type: schema.TypeString,
			Optional: true,
		},
		
	}
}

func SetCustomVirtualDataPointSubResourceData(m []*models.CustomVirtualDataPoint) (d []*map[string]interface{}) {
	for _, customVirtualDataPoint := range m {
		if customVirtualDataPoint != nil {
			properties := make(map[string]interface{})
			properties["custom_graph_id"] = customVirtualDataPoint.CustomGraphID
			properties["display"] = SetGraphDisplaySubResourceData([]*models.GraphDisplay{customVirtualDataPoint.Display})
			properties["id"] = customVirtualDataPoint.ID
			properties["name"] = customVirtualDataPoint.Name
			properties["rpn"] = customVirtualDataPoint.Rpn
			d = append(d, &properties)
		}
	}
	return
}

func CustomVirtualDataPointModel(d map[string]interface{}) *models.CustomVirtualDataPoint {
	// assume that the incoming map only contains the relevant resource data
	customGraphID := int32(d["custom_graph_id"].(int))
	var display *models.GraphDisplay = nil
	displayList := d["display"].([]interface{})
	if len(displayList) > 0 { // len(nil) = 0
		display = GraphDisplayModel(displayList[0].(map[string]interface{}))
	}
	id := int32(d["id"].(int))
	name := d["name"].(string)
	rpn := d["rpn"].(string)
	
	return &models.CustomVirtualDataPoint {
		CustomGraphID: customGraphID,
		Display: display,
		ID: id,
		Name: name,
		Rpn: rpn,
	}
}

func GetCustomVirtualDataPointPropertyFields() (t []string) {
	return []string{
		"custom_graph_id",
		"display",
		"id",
		"name",
		"rpn",
	}
}