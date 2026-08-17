package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func WidgetRefreshFrequencySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"offset": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"units": {
			Type: schema.TypeString,
			Optional: true,
		},
		
	}
}

func SetWidgetRefreshFrequencySubResourceData(m []*models.WidgetRefreshFrequency) (d []*map[string]interface{}) {
	for _, widgetRefreshFrequency := range m {
		if widgetRefreshFrequency != nil {
			properties := make(map[string]interface{})
			properties["offset"] = widgetRefreshFrequency.Offset
			properties["units"] = widgetRefreshFrequency.Units
			d = append(d, &properties)
		}
	}
	return
}

func WidgetRefreshFrequencyModel(d map[string]interface{}) *models.WidgetRefreshFrequency {
	// assume that the incoming map only contains the relevant resource data
	offset := int32(d["offset"].(int))
	units := d["units"].(string)
	
	return &models.WidgetRefreshFrequency {
		Offset: offset,
		Units: units,
	}
}

func GetWidgetRefreshFrequencyPropertyFields() (t []string) {
	return []string{
		"offset",
		"units",
	}
}