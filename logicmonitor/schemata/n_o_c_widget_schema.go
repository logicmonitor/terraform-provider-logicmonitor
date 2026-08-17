package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func NOCWidgetSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
	}
}

func SetNOCWidgetSubResourceData(m []*models.NOCWidget) (d []*map[string]interface{}) {
	for _, nOCWidget := range m {
		if nOCWidget != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func NOCWidgetModel(d map[string]interface{}) *models.NOCWidget {
	// assume that the incoming map only contains the relevant resource data
	
	return &models.NOCWidget {
	}
}

func GetNOCWidgetPropertyFields() (t []string) {
	return []string{
	}
}