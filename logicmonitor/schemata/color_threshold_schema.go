package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ColorThresholdSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"level": {
			Type: schema.TypeInt,
			Required: true,
		},
		
		"relation": {
			Type: schema.TypeString,
			Required: true,
		},
		
		"threshold": {
			Type: schema.TypeFloat,
			Required: true,
		},
		
	}
}

func SetColorThresholdSubResourceData(m []*models.ColorThreshold) (d []*map[string]interface{}) {
	for _, colorThreshold := range m {
		if colorThreshold != nil {
			properties := make(map[string]interface{})
			properties["level"] = colorThreshold.Level
			properties["relation"] = colorThreshold.Relation
			properties["threshold"] = colorThreshold.Threshold
			d = append(d, &properties)
		}
	}
	return
}

func ColorThresholdModel(d map[string]interface{}) *models.ColorThreshold {
	// assume that the incoming map only contains the relevant resource data
	level := int32(d["level"].(int))
	relation := d["relation"].(string)
	threshold := d["threshold"].(float64)
	
	return &models.ColorThreshold {
		Level: &level,
		Relation: &relation,
		Threshold: &threshold,
	}
}

func GetColorThresholdPropertyFields() (t []string) {
	return []string{
		"level",
		"relation",
		"threshold",
	}
}