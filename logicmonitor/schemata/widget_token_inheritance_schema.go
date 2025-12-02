package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-logicmonitor/models"
)

func WidgetTokenInheritanceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"fullpath": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"value": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func SetWidgetTokenInheritanceSubResourceData(m []*models.WidgetTokenInheritance) (d []*map[string]interface{}) {
	for _, widgetTokenInheritance := range m {
		if widgetTokenInheritance != nil {
			properties := make(map[string]interface{})
			properties["fullpath"] = widgetTokenInheritance.Fullpath
			properties["value"] = widgetTokenInheritance.Value
			d = append(d, &properties)
		}
	}
	return
}

func WidgetTokenInheritanceModel(d map[string]interface{}) *models.WidgetTokenInheritance {
	// assume that the incoming map only contains the relevant resource data

	return &models.WidgetTokenInheritance{}
}

func GetWidgetTokenInheritancePropertyFields() (t []string) {
	return []string{}
}
