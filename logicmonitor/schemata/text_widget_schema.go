package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TextWidgetSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
	}
}

func SetTextWidgetSubResourceData(m []*models.TextWidget) (d []*map[string]interface{}) {
	for _, textWidget := range m {
		if textWidget != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func TextWidgetModel(d map[string]interface{}) *models.TextWidget {
	// assume that the incoming map only contains the relevant resource data
	
	return &models.TextWidget {
	}
}

func GetTextWidgetPropertyFields() (t []string) {
	return []string{
	}
}