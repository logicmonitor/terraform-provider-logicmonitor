package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func StageSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type: schema.TypeInt,
			Required: true,
		},
		
		"name": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"type": {
			Type: schema.TypeString,
			Required: true,
		},
		
	}
}

func SetStageSubResourceData(m []*models.Stage) (d []*map[string]interface{}) {
	for _, stage := range m {
		if stage != nil {
			properties := make(map[string]interface{})
			properties["id"] = stage.ID
			properties["name"] = stage.Name
			properties["type"] = stage.Type
			d = append(d, &properties)
		}
	}
	return
}

func StageModel(d map[string]interface{}) *models.Stage {
	// assume that the incoming map only contains the relevant resource data
	id := int32(d["id"].(int))
	typeVar := d["type"].(string)
	
	return &models.Stage {
		ID: &id,
		Type: &typeVar,
	}
}

func GetStagePropertyFields() (t []string) {
	return []string{
		"id",
		"type",
	}
}