package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func GraphDisplaySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"color": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"legend": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"option": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"type": {
			Type: schema.TypeString,
			Optional: true,
		},
		
	}
}

func SetGraphDisplaySubResourceData(m []*models.GraphDisplay) (d []*map[string]interface{}) {
	for _, graphDisplay := range m {
		if graphDisplay != nil {
			properties := make(map[string]interface{})
			properties["color"] = graphDisplay.Color
			properties["legend"] = graphDisplay.Legend
			properties["option"] = graphDisplay.Option
			properties["type"] = graphDisplay.Type
			d = append(d, &properties)
		}
	}
	return
}

func GraphDisplayModel(d map[string]interface{}) *models.GraphDisplay {
	// assume that the incoming map only contains the relevant resource data
	color := d["color"].(string)
	legend := d["legend"].(string)
	option := d["option"].(string)
	typeVar := d["type"].(string)
	
	return &models.GraphDisplay {
		Color: color,
		Legend: legend,
		Option: option,
		Type: typeVar,
	}
}

func GetGraphDisplayPropertyFields() (t []string) {
	return []string{
		"color",
		"legend",
		"option",
		"type",
	}
}