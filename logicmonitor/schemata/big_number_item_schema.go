package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func BigNumberItemSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"bottom_label": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"change_threshold_color_toggle": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"color_thresholds": {
			Type: schema.TypeList, //GoType: []*ColorThreshold  
			Elem: &schema.Resource{
				Schema: ColorThresholdSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"data_point_name": {
			Type: schema.TypeString,
			Required: true,
		},
		
		"position": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"right_label": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"rounding": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"use_comma_separators": {
			Type: schema.TypeBool,
			Required: true,
		},
		
	}
}

func SetBigNumberItemSubResourceData(m []*models.BigNumberItem) (d []*map[string]interface{}) {
	for _, bigNumberItem := range m {
		if bigNumberItem != nil {
			properties := make(map[string]interface{})
			properties["bottom_label"] = bigNumberItem.BottomLabel
			properties["change_threshold_color_toggle"] = bigNumberItem.ChangeThresholdColorToggle
			properties["color_thresholds"] = SetColorThresholdSubResourceData(bigNumberItem.ColorThresholds)
			properties["data_point_name"] = bigNumberItem.DataPointName
			properties["position"] = bigNumberItem.Position
			properties["right_label"] = bigNumberItem.RightLabel
			properties["rounding"] = bigNumberItem.Rounding
			properties["use_comma_separators"] = bigNumberItem.UseCommaSeparators
			d = append(d, &properties)
		}
	}
	return
}

func BigNumberItemModel(d map[string]interface{}) *models.BigNumberItem {
	// assume that the incoming map only contains the relevant resource data
	bottomLabel := d["bottom_label"].(string)
	changeThresholdColorToggle := d["change_threshold_color_toggle"].(bool)
	colorThresholds := d["color_thresholds"].([]*models.ColorThreshold)
	dataPointName := d["data_point_name"].(string)
	position := int32(d["position"].(int))
	rightLabel := d["right_label"].(string)
	rounding := int32(d["rounding"].(int))
	useCommaSeparators := d["use_comma_separators"].(bool)
	
	return &models.BigNumberItem {
		BottomLabel: bottomLabel,
		ChangeThresholdColorToggle: changeThresholdColorToggle,
		ColorThresholds: colorThresholds,
		DataPointName: &dataPointName,
		Position: position,
		RightLabel: rightLabel,
		Rounding: rounding,
		UseCommaSeparators: &useCommaSeparators,
	}
}

func GetBigNumberItemPropertyFields() (t []string) {
	return []string{
		"bottom_label",
		"change_threshold_color_toggle",
		"color_thresholds",
		"data_point_name",
		"position",
		"right_label",
		"rounding",
		"use_comma_separators",
	}
}