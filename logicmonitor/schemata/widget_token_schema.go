package schemata

import (
	"terraform-provider-logicmonitor/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func WidgetTokenSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"inherit_list": {
			Type: schema.TypeList, //GoType: []*WidgetTokenInheritance
			Elem: &schema.Resource{
				Schema: WidgetTokenInheritanceSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"type": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"value": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func SetWidgetTokenSubResourceData(m []*models.WidgetToken) (d []*map[string]interface{}) {
	for _, widgetToken := range m {
		if widgetToken != nil {
			properties := make(map[string]interface{})
			properties["inherit_list"] = SetWidgetTokenInheritanceSubResourceData(widgetToken.InheritList)
			properties["name"] = widgetToken.Name
			properties["type"] = widgetToken.Type
			properties["value"] = widgetToken.Value
			d = append(d, &properties)
		}
	}
	return
}

func WidgetTokenModel(d map[string]interface{}) *models.WidgetToken {
	// assume that the incoming map only contains the relevant resource data
	name := d["name"].(string)
	value := d["value"].(string)

	return &models.WidgetToken{
		Name:  name,
		Value: value,
	}
}

func GetWidgetTokenPropertyFields() (t []string) {
	return []string{
		"name",
		"value",
	}
}
