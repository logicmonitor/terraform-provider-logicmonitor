package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func AlertWidgetDisplaySettingsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"columns": {
			Type: schema.TypeList, //GoType: []*AlertWidgetColumn  
			Elem: &schema.Resource{
				Schema: AlertWidgetColumnSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"fontsize": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"is_show_all": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"play_sound": {
			Type: schema.TypeList, //GoType: AlertWidgetPlaySoundSettings
			Elem: &schema.Resource{
				Schema: AlertWidgetPlaySoundSettingsSchema(),
			},
			Optional: true,
		},
		
		"show_filter": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"sort": {
			Type: schema.TypeString,
			Optional: true,
		},
		
	}
}

func SetAlertWidgetDisplaySettingsSubResourceData(m []*models.AlertWidgetDisplaySettings) (d []*map[string]interface{}) {
	for _, alertWidgetDisplaySettings := range m {
		if alertWidgetDisplaySettings != nil {
			properties := make(map[string]interface{})
			properties["columns"] = SetAlertWidgetColumnSubResourceData(alertWidgetDisplaySettings.Columns)
			properties["fontsize"] = alertWidgetDisplaySettings.Fontsize
			properties["is_show_all"] = alertWidgetDisplaySettings.IsShowAll
			properties["play_sound"] = SetAlertWidgetPlaySoundSettingsSubResourceData([]*models.AlertWidgetPlaySoundSettings{alertWidgetDisplaySettings.PlaySound})
			properties["show_filter"] = alertWidgetDisplaySettings.ShowFilter
			properties["sort"] = alertWidgetDisplaySettings.Sort
			d = append(d, &properties)
		}
	}
	return
}

func AlertWidgetDisplaySettingsModel(d map[string]interface{}) *models.AlertWidgetDisplaySettings {
	// assume that the incoming map only contains the relevant resource data
	columns := d["columns"].([]*models.AlertWidgetColumn)
	fontsize := d["fontsize"].(string)
	isShowAll := d["is_show_all"].(bool)
	var playSound *models.AlertWidgetPlaySoundSettings = nil
	playSoundList := d["play_sound"].([]interface{})
	if len(playSoundList) > 0 { // len(nil) = 0
		playSound = AlertWidgetPlaySoundSettingsModel(playSoundList[0].(map[string]interface{}))
	}
	showFilter := d["show_filter"].(bool)
	sort := d["sort"].(string)
	
	return &models.AlertWidgetDisplaySettings {
		Columns: columns,
		Fontsize: fontsize,
		IsShowAll: isShowAll,
		PlaySound: playSound,
		ShowFilter: showFilter,
		Sort: sort,
	}
}

func GetAlertWidgetDisplaySettingsPropertyFields() (t []string) {
	return []string{
		"columns",
		"fontsize",
		"is_show_all",
		"play_sound",
		"show_filter",
		"sort",
	}
}