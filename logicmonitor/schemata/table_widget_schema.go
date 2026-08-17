package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TableWidgetSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
	}
}

func SetTableWidgetSubResourceData(m []*models.TableWidget) (d []*map[string]interface{}) {
	for _, tableWidget := range m {
		if tableWidget != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func TableWidgetModel(d map[string]interface{}) *models.TableWidget {
	// assume that the incoming map only contains the relevant resource data
	
	return &models.TableWidget {
	}
}

func GetTableWidgetPropertyFields() (t []string) {
	return []string{
	}
}