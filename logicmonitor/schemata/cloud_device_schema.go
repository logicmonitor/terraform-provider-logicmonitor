package schemata

import (
	"terraform-provider-logicmonitor/logicmonitor/utils"
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func CloudDeviceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"device_type": {
			Type: schema.TypeInt,
			Required: true,
		},
		
		"required_props": {
			Type: schema.TypeSet,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Set:      schema.HashString,
			Required: true,
		},
		
	}
}

func SetCloudDeviceSubResourceData(m []*models.CloudDevice) (d []*map[string]interface{}) {
	for _, cloudDevice := range m {
		if cloudDevice != nil {
			properties := make(map[string]interface{})
			properties["device_type"] = cloudDevice.DeviceType
			properties["required_props"] = cloudDevice.RequiredProps
			d = append(d, &properties)
		}
	}
	return
}

func CloudDeviceModel(d map[string]interface{}) *models.CloudDevice {
	// assume that the incoming map only contains the relevant resource data
	deviceType := int32(d["device_type"].(int))
	requiredProps := utils.ConvertSetToStringSlice(d["required_props"].(*schema.Set))
	
	return &models.CloudDevice {
		DeviceType: &deviceType,
		RequiredProps: requiredProps,
	}
}

func GetCloudDevicePropertyFields() (t []string) {
	return []string{
		"device_type",
		"required_props",
	}
}