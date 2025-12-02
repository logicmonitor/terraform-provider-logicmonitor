package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-logicmonitor/models"
)

func AutomaticUpgradeInfoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"created_by": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"day_of_week": {
			Type:     schema.TypeString,
			Required: true,
		},

		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"hour": {
			Type:     schema.TypeInt,
			Required: true,
		},

		"level": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"minute": {
			Type:     schema.TypeInt,
			Required: true,
		},

		"occurrence": {
			Type:     schema.TypeString,
			Required: true,
		},

		"timezone": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"type": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"version": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

func SetAutomaticUpgradeInfoSubResourceData(m []*models.AutomaticUpgradeInfo) (d []*map[string]interface{}) {
	for _, automaticUpgradeInfo := range m {
		if automaticUpgradeInfo != nil {
			properties := make(map[string]interface{})
			properties["created_by"] = automaticUpgradeInfo.CreatedBy
			properties["day_of_week"] = automaticUpgradeInfo.DayOfWeek
			properties["description"] = automaticUpgradeInfo.Description
			properties["hour"] = automaticUpgradeInfo.Hour
			properties["level"] = automaticUpgradeInfo.Level
			properties["minute"] = automaticUpgradeInfo.Minute
			properties["occurrence"] = automaticUpgradeInfo.Occurrence
			properties["timezone"] = automaticUpgradeInfo.Timezone
			properties["type"] = automaticUpgradeInfo.Type
			properties["version"] = automaticUpgradeInfo.Version
			d = append(d, &properties)
		}
	}
	return
}

func AutomaticUpgradeInfoModel(d map[string]interface{}) *models.AutomaticUpgradeInfo {
	// assume that the incoming map only contains the relevant resource data
	dayOfWeek := d["day_of_week"].(string)
	description := d["description"].(string)
	hour := int32(d["hour"].(int))
	minute := int32(d["minute"].(int))
	occurrence := d["occurrence"].(string)
	timezone := d["timezone"].(string)
	version := d["version"].(string)

	return &models.AutomaticUpgradeInfo{
		DayOfWeek:   &dayOfWeek,
		Description: description,
		Hour:        &hour,
		Minute:      &minute,
		Occurrence:  &occurrence,
		Timezone:    timezone,
		Version:     &version,
	}
}

func GetAutomaticUpgradeInfoPropertyFields() (t []string) {
	return []string{
		"day_of_week",
		"description",
		"hour",
		"minute",
		"occurrence",
		"timezone",
		"version",
	}
}
