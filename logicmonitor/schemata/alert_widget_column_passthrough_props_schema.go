package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func AlertWidgetColumnPassthroughPropsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"has_display_name": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"has_icon": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"has_relative_time": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"type": {
			Type: schema.TypeString,
			Optional: true,
		},
		
	}
}

func SetAlertWidgetColumnPassthroughPropsSubResourceData(m []*models.AlertWidgetColumnPassthroughProps) (d []*map[string]interface{}) {
	for _, alertWidgetColumnPassthroughProps := range m {
		if alertWidgetColumnPassthroughProps != nil {
			properties := make(map[string]interface{})
			properties["has_display_name"] = alertWidgetColumnPassthroughProps.HasDisplayName
			properties["has_icon"] = alertWidgetColumnPassthroughProps.HasIcon
			properties["has_relative_time"] = alertWidgetColumnPassthroughProps.HasRelativeTime
			properties["type"] = alertWidgetColumnPassthroughProps.Type
			d = append(d, &properties)
		}
	}
	return
}

func AlertWidgetColumnPassthroughPropsModel(d map[string]interface{}) *models.AlertWidgetColumnPassthroughProps {
	// assume that the incoming map only contains the relevant resource data
	hasDisplayName := d["has_display_name"].(bool)
	hasIcon := d["has_icon"].(bool)
	hasRelativeTime := d["has_relative_time"].(bool)
	typeVar := d["type"].(string)
	
	return &models.AlertWidgetColumnPassthroughProps {
		HasDisplayName: hasDisplayName,
		HasIcon: hasIcon,
		HasRelativeTime: hasRelativeTime,
		Type: typeVar,
	}
}

func GetAlertWidgetColumnPassthroughPropsPropertyFields() (t []string) {
	return []string{
		"has_display_name",
		"has_icon",
		"has_relative_time",
		"type",
	}
}