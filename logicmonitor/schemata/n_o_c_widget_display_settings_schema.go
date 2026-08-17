package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func NOCWidgetDisplaySettingsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"display_as": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"override_filter": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"show_filter": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"show_type_icon": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
	}
}

func SetNOCWidgetDisplaySettingsSubResourceData(m []*models.NOCWidgetDisplaySettings) (d []*map[string]interface{}) {
	for _, nOCWidgetDisplaySettings := range m {
		if nOCWidgetDisplaySettings != nil {
			properties := make(map[string]interface{})
			properties["display_as"] = nOCWidgetDisplaySettings.DisplayAs
			properties["override_filter"] = nOCWidgetDisplaySettings.OverrideFilter
			properties["show_filter"] = nOCWidgetDisplaySettings.ShowFilter
			properties["show_type_icon"] = nOCWidgetDisplaySettings.ShowTypeIcon
			d = append(d, &properties)
		}
	}
	return
}

func NOCWidgetDisplaySettingsModel(d map[string]interface{}) *models.NOCWidgetDisplaySettings {
	// assume that the incoming map only contains the relevant resource data
	displayAs := d["display_as"].(string)
	overrideFilter := d["override_filter"].(bool)
	showFilter := d["show_filter"].(bool)
	showTypeIcon := d["show_type_icon"].(bool)
	
	return &models.NOCWidgetDisplaySettings {
		DisplayAs: displayAs,
		OverrideFilter: overrideFilter,
		ShowFilter: showFilter,
		ShowTypeIcon: showTypeIcon,
	}
}

func GetNOCWidgetDisplaySettingsPropertyFields() (t []string) {
	return []string{
		"display_as",
		"override_filter",
		"show_filter",
		"show_type_icon",
	}
}