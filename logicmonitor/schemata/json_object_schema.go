package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func JSONObjectSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"empty": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
	}
}

func SetJSONObjectSubResourceData(m []*models.JSONObject) (d []*map[string]interface{}) {
	for _, jsonObject := range m {
		if jsonObject != nil {
			properties := make(map[string]interface{})
			properties["empty"] = jsonObject.Empty
			d = append(d, &properties)
		}
	}
	return
}

func JSONObjectModel(d map[string]interface{}) *models.JSONObject {
	// assume that the incoming map only contains the relevant resource data
	empty := d["empty"].(bool)
	
	return &models.JSONObject {
		Empty: empty,
	}
}

func GetJSONObjectPropertyFields() (t []string) {
	return []string{
		"empty",
	}
}