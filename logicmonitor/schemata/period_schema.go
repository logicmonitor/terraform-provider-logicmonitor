package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-logicmonitor/logicmonitor/utils"
	"terraform-provider-logicmonitor/models"
)

func PeriodSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"end_minutes": {
			Type:     schema.TypeInt,
			Required: true,
		},

		"start_minutes": {
			Type:     schema.TypeInt,
			Required: true,
		},

		"timezone": {
			Type:     schema.TypeString,
			Required: true,
		},

		"week_days": {
			Type: schema.TypeList, //GoType: []int32
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Required:   true,
		},
	}
}

func SetPeriodSubResourceData(m []*models.Period) (d []*map[string]interface{}) {
	for _, period := range m {
		if period != nil {
			properties := make(map[string]interface{})
			properties["end_minutes"] = period.EndMinutes
			properties["start_minutes"] = period.StartMinutes
			properties["timezone"] = period.Timezone
			properties["week_days"] = period.WeekDays
			d = append(d, &properties)
		}
	}
	return
}

func PeriodModel(d map[string]interface{}) *models.Period {
	// assume that the incoming map only contains the relevant resource data
	endMinutes := int32(d["end_minutes"].(int))
	startMinutes := int32(d["start_minutes"].(int))
	timezone := d["timezone"].(string)
	weekDays := utils.ConvertSetToInt32Slice(d["weekDays"].([]interface{}))

	return &models.Period{
		EndMinutes:   &endMinutes,
		StartMinutes: &startMinutes,
		Timezone:     &timezone,
		WeekDays:     weekDays,
	}
}

func GetPeriodPropertyFields() (t []string) {
	return []string{
		"end_minutes",
		"start_minutes",
		"timezone",
		"week_days",
	}
}
