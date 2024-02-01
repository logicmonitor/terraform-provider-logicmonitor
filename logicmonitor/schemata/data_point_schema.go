package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataPointSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"alert_body": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"alert_clear_transition_interval": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"alert_expr": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"alert_expr_note": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"alert_for_no_data": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"alert_subject": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"alert_transition_interval": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"data_type": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"description": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"max_digits": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"max_value": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"min_value": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"name": {
			Type: schema.TypeString,
			Required: true,
		},
		
		"post_processor_method": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"post_processor_param": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"raw_data_field_name": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"type": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
	}
}

func SetDataPointSubResourceData(m []*models.DataPoint) (d []*map[string]interface{}) {
	for _, dataPoint := range m {
		if dataPoint != nil {
			properties := make(map[string]interface{})
			properties["alert_body"] = dataPoint.AlertBody
			properties["alert_clear_transition_interval"] = dataPoint.AlertClearTransitionInterval
			properties["alert_expr"] = dataPoint.AlertExpr
			properties["alert_expr_note"] = dataPoint.AlertExprNote
			properties["alert_for_no_data"] = dataPoint.AlertForNoData
			properties["alert_subject"] = dataPoint.AlertSubject
			properties["alert_transition_interval"] = dataPoint.AlertTransitionInterval
			properties["data_type"] = dataPoint.DataType
			properties["description"] = dataPoint.Description
			properties["max_digits"] = dataPoint.MaxDigits
			properties["max_value"] = dataPoint.MaxValue
			properties["min_value"] = dataPoint.MinValue
			properties["name"] = dataPoint.Name
			properties["post_processor_method"] = dataPoint.PostProcessorMethod
			properties["post_processor_param"] = dataPoint.PostProcessorParam
			properties["raw_data_field_name"] = dataPoint.RawDataFieldName
			properties["type"] = dataPoint.Type
			d = append(d, &properties)
		}
	}
	return
}

func DataPointModel(d map[string]interface{}) *models.DataPoint {
	// assume that the incoming map only contains the relevant resource data
	alertBody := d["alert_body"].(string)
	alertClearTransitionInterval := int32(d["alert_clear_transition_interval"].(int))
	alertExpr := d["alert_expr"].(string)
	alertExprNote := d["alert_expr_note"].(string)
	alertForNoData := int32(d["alert_for_no_data"].(int))
	alertSubject := d["alert_subject"].(string)
	alertTransitionInterval := int32(d["alert_transition_interval"].(int))
	dataType := int32(d["data_type"].(int))
	description := d["description"].(string)
	maxDigits := int32(d["max_digits"].(int))
	maxValue := d["max_value"].(string)
	minValue := d["min_value"].(string)
	name := d["name"].(string)
	postProcessorMethod := d["post_processor_method"].(string)
	postProcessorParam := d["post_processor_param"].(string)
	rawDataFieldName := d["raw_data_field_name"].(string)
	typeVar := int32(d["type"].(int))
	
	return &models.DataPoint {
		AlertBody: alertBody,
		AlertClearTransitionInterval: alertClearTransitionInterval,
		AlertExpr: alertExpr,
		AlertExprNote: alertExprNote,
		AlertForNoData: alertForNoData,
		AlertSubject: alertSubject,
		AlertTransitionInterval: alertTransitionInterval,
		DataType: dataType,
		Description: description,
		MaxDigits: maxDigits,
		MaxValue: maxValue,
		MinValue: minValue,
		Name: &name,
		PostProcessorMethod: postProcessorMethod,
		PostProcessorParam: postProcessorParam,
		RawDataFieldName: rawDataFieldName,
		Type: typeVar,
	}
}

func GetDataPointPropertyFields() (t []string) {
	return []string{
		"alert_body",
		"alert_clear_transition_interval",
		"alert_expr",
		"alert_expr_note",
		"alert_for_no_data",
		"alert_subject",
		"alert_transition_interval",
		"data_type",
		"description",
		"max_digits",
		"max_value",
		"min_value",
		"name",
		"post_processor_method",
		"post_processor_param",
		"raw_data_field_name",
		"type",
	}
}