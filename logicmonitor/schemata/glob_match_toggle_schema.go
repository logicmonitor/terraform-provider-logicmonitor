package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func GlobMatchToggleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"is_glob": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"value": {
			Type: schema.TypeString,
			Optional: true,
		},
		
	}
}

func SetGlobMatchToggleSubResourceData(m []*models.GlobMatchToggle) (d []*map[string]interface{}) {
	for _, globMatchToggle := range m {
		if globMatchToggle != nil {
			properties := make(map[string]interface{})
			properties["is_glob"] = globMatchToggle.IsGlob
			properties["value"] = globMatchToggle.Value
			d = append(d, &properties)
		}
	}
	return
}

func GlobMatchToggleModel(d map[string]interface{}) *models.GlobMatchToggle {
	// assume that the incoming map only contains the relevant resource data
	isGlob := d["is_glob"].(bool)
	value := d["value"].(string)
	
	return &models.GlobMatchToggle {
		IsGlob: isGlob,
		Value: value,
	}
}

func GetGlobMatchTogglePropertyFields() (t []string) {
	return []string{
		"is_glob",
		"value",
	}
}