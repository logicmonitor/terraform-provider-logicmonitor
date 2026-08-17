package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func AlertWidgetPlaySoundSettingsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"critical_alert_audio_file_name": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"error_alert_audio_file_name": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"should_play": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"warning_alert_audio_file_name": {
			Type: schema.TypeString,
			Optional: true,
		},
		
	}
}

func SetAlertWidgetPlaySoundSettingsSubResourceData(m []*models.AlertWidgetPlaySoundSettings) (d []*map[string]interface{}) {
	for _, alertWidgetPlaySoundSettings := range m {
		if alertWidgetPlaySoundSettings != nil {
			properties := make(map[string]interface{})
			properties["critical_alert_audio_file_name"] = alertWidgetPlaySoundSettings.CriticalAlertAudioFileName
			properties["error_alert_audio_file_name"] = alertWidgetPlaySoundSettings.ErrorAlertAudioFileName
			properties["should_play"] = alertWidgetPlaySoundSettings.ShouldPlay
			properties["warning_alert_audio_file_name"] = alertWidgetPlaySoundSettings.WarningAlertAudioFileName
			d = append(d, &properties)
		}
	}
	return
}

func AlertWidgetPlaySoundSettingsModel(d map[string]interface{}) *models.AlertWidgetPlaySoundSettings {
	// assume that the incoming map only contains the relevant resource data
	criticalAlertAudioFileName := d["critical_alert_audio_file_name"].(string)
	errorAlertAudioFileName := d["error_alert_audio_file_name"].(string)
	shouldPlay := d["should_play"].(bool)
	warningAlertAudioFileName := d["warning_alert_audio_file_name"].(string)
	
	return &models.AlertWidgetPlaySoundSettings {
		CriticalAlertAudioFileName: criticalAlertAudioFileName,
		ErrorAlertAudioFileName: errorAlertAudioFileName,
		ShouldPlay: shouldPlay,
		WarningAlertAudioFileName: warningAlertAudioFileName,
	}
}

func GetAlertWidgetPlaySoundSettingsPropertyFields() (t []string) {
	return []string{
		"critical_alert_audio_file_name",
		"error_alert_audio_file_name",
		"should_play",
		"warning_alert_audio_file_name",
	}
}