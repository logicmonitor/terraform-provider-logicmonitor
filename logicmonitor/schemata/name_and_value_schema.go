package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func NameAndValueSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type: schema.TypeString,
			Required: true,
		},
		
		"value": {
			Type: schema.TypeString,
			Required: true,
		},
		
	}
}

func SetNameAndValueSubResourceData(m []*models.NameAndValue) (d []*map[string]interface{}) {
	for _, nameAndValue := range m {
		if nameAndValue != nil {
			properties := make(map[string]interface{})
			properties["name"] = nameAndValue.Name
			properties["value"] = nameAndValue.Value
			d = append(d, &properties)
		}
	}
	return
}

func NameAndValueModel(d map[string]interface{}) *models.NameAndValue {
	// assume that the incoming map only contains the relevant resource data
	name := d["name"].(string)
	value := d["value"].(string)
	
	return &models.NameAndValue {
		Name: &name,
		Value: &value,
	}
}

func GetNameAndValuePropertyFields() (t []string) {
	return []string{
		"name",
		"value",
	}
}