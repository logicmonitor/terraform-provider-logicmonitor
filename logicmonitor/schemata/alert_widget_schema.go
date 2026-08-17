package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func AlertWidgetSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
	}
}

func SetAlertWidgetSubResourceData(m []*models.AlertWidget) (d []*map[string]interface{}) {
	for _, alertWidget := range m {
		if alertWidget != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func AlertWidgetModel(d map[string]interface{}) *models.AlertWidget {
	// assume that the incoming map only contains the relevant resource data
	
	return &models.AlertWidget {
	}
}

func GetAlertWidgetPropertyFields() (t []string) {
	return []string{
	}
}