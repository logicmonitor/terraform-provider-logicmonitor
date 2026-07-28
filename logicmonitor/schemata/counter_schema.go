package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func CounterSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"applies_to": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"name": {
			Type: schema.TypeString,
			Optional: true,
		},
		
	}
}

func SetCounterSubResourceData(m []*models.Counter) (d []*map[string]interface{}) {
	for _, counter := range m {
		if counter != nil {
			properties := make(map[string]interface{})
			properties["applies_to"] = counter.AppliesTo
			properties["name"] = counter.Name
			d = append(d, &properties)
		}
	}
	return
}

func CounterModel(d map[string]interface{}) *models.Counter {
	// assume that the incoming map only contains the relevant resource data
	appliesTo := d["applies_to"].(string)
	name := d["name"].(string)
	
	return &models.Counter {
		AppliesTo: appliesTo,
		Name: name,
	}
}

func GetCounterPropertyFields() (t []string) {
	return []string{
		"applies_to",
		"name",
	}
}