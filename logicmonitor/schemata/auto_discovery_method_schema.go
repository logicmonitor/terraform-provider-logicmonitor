package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func AutoDiscoveryMethodSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type: schema.TypeString,
			Required: true,
		},
		
	}
}

func SetAutoDiscoveryMethodSubResourceData(m []*models.AutoDiscoveryMethod) (d []*map[string]interface{}) {
	for _, autoDiscoveryMethod := range m {
		if autoDiscoveryMethod != nil {
			properties := make(map[string]interface{})
			properties["name"] = autoDiscoveryMethod.Name
			d = append(d, &properties)
		}
	}
	return
}

func AutoDiscoveryMethodModel(d map[string]interface{}) *models.AutoDiscoveryMethod {
	// assume that the incoming map only contains the relevant resource data
	name := d["name"].(string)
	
	return &models.AutoDiscoveryMethod {
		Name: &name,
	}
}

func GetAutoDiscoveryMethodPropertyFields() (t []string) {
	return []string{
		"name",
	}
}