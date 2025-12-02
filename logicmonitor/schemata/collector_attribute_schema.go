package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-logicmonitor/models"
)

func CollectorAttributeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

func SetCollectorAttributeSubResourceData(m []*models.CollectorAttribute) (d []*map[string]interface{}) {
	for _, collectorAttribute := range m {
		if collectorAttribute != nil {
			properties := make(map[string]interface{})
			properties["name"] = collectorAttribute.Name
			d = append(d, &properties)
		}
	}
	return
}

func CollectorAttributeModel(d map[string]interface{}) *models.CollectorAttribute {
	// assume that the incoming map only contains the relevant resource data
	name := d["name"].(string)

	return &models.CollectorAttribute{
		Name: &name,
	}
}

func GetCollectorAttributePropertyFields() (t []string) {
	return []string{
		"name",
	}
}
