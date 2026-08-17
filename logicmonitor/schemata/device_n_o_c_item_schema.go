package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DeviceNOCItemSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
	}
}

func SetDeviceNOCItemSubResourceData(m []*models.DeviceNOCItem) (d []*map[string]interface{}) {
	for _, deviceNOCItem := range m {
		if deviceNOCItem != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func DeviceNOCItemModel(d map[string]interface{}) *models.DeviceNOCItem {
	// assume that the incoming map only contains the relevant resource data
	
	return &models.DeviceNOCItem {
	}
}

func GetDeviceNOCItemPropertyFields() (t []string) {
	return []string{
	}
}