package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-logicmonitor/models"
)

func DevicePropertySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},

		"value": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

func SetDevicePropertySubResourceData(m []*models.DeviceProperty) (d []*map[string]interface{}) {
	for _, deviceProperty := range m {
		if deviceProperty != nil {
			properties := make(map[string]interface{})
			properties["name"] = deviceProperty.Name
			properties["value"] = deviceProperty.Value
			d = append(d, &properties)
		}
	}
	return
}

func DevicePropertyModel(d map[string]interface{}) *models.DeviceProperty {
	// assume that the incoming map only contains the relevant resource data
	name := d["name"].(string)
	value := d["value"].(string)

	return &models.DeviceProperty{
		Name:  &name,
		Value: &value,
	}
}

func GetDevicePropertyPropertyFields() (t []string) {
	return []string{
		"name",
		"value",
	}
}
