package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func VirtualDataPointSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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

func SetVirtualDataPointSubResourceData(m []*models.VirtualDataPoint) (d []*map[string]interface{}) {
	for _, virtualDataPoint := range m {
		if virtualDataPoint != nil {
			properties := make(map[string]interface{})
			properties["name"] = virtualDataPoint.Name
			properties["rpn"] = virtualDataPoint.Rpn
			d = append(d, &properties)
		}
	}
	return
}

func VirtualDataPointModel(d map[string]interface{}) *models.VirtualDataPoint {
	// assume that the incoming map only contains the relevant resource data
	name := d["name"].(string)
	rpn := d["rpn"].(string)
	
	return &models.VirtualDataPoint {
		Name: name,
		Rpn: rpn,
	}
}

func GetVirtualDataPointPropertyFields() (t []string) {
	return []string{
		"name",
		"rpn",
	}
}