package schemata

import (
	"terraform-provider-logicmonitor/logicmonitor/utils"
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func AlertWidgetColumnSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"column_key": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"column_label": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"column_size": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"component_name": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"data_path": {
			Type: schema.TypeSet,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Set:      schema.HashString,
			Optional: true,
		},
		
		"is_sortable": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"min_size": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"passthrough_props": {
			Type: schema.TypeList, //GoType: AlertWidgetColumnPassthroughProps
			Elem: &schema.Resource{
				Schema: AlertWidgetColumnPassthroughPropsSchema(),
			},
			Optional: true,
		},
		
		"visible": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
	}
}

func SetAlertWidgetColumnSubResourceData(m []*models.AlertWidgetColumn) (d []*map[string]interface{}) {
	for _, alertWidgetColumn := range m {
		if alertWidgetColumn != nil {
			properties := make(map[string]interface{})
			properties["column_key"] = alertWidgetColumn.ColumnKey
			properties["column_label"] = alertWidgetColumn.ColumnLabel
			properties["column_size"] = alertWidgetColumn.ColumnSize
			properties["component_name"] = alertWidgetColumn.ComponentName
			properties["data_path"] = alertWidgetColumn.DataPath
			properties["is_sortable"] = alertWidgetColumn.IsSortable
			properties["min_size"] = alertWidgetColumn.MinSize
			properties["passthrough_props"] = SetAlertWidgetColumnPassthroughPropsSubResourceData([]*models.AlertWidgetColumnPassthroughProps{alertWidgetColumn.PassthroughProps})
			properties["visible"] = alertWidgetColumn.Visible
			d = append(d, &properties)
		}
	}
	return
}

func AlertWidgetColumnModel(d map[string]interface{}) *models.AlertWidgetColumn {
	// assume that the incoming map only contains the relevant resource data
	columnKey := d["column_key"].(string)
	columnLabel := d["column_label"].(string)
	columnSize := int32(d["column_size"].(int))
	componentName := d["component_name"].(string)
	dataPath := utils.ConvertSetToStringSlice(d["data_path"].(*schema.Set))
	isSortable := d["is_sortable"].(bool)
	minSize := int32(d["min_size"].(int))
	var passthroughProps *models.AlertWidgetColumnPassthroughProps = nil
	passthroughPropsList := d["passthrough_props"].([]interface{})
	if len(passthroughPropsList) > 0 { // len(nil) = 0
		passthroughProps = AlertWidgetColumnPassthroughPropsModel(passthroughPropsList[0].(map[string]interface{}))
	}
	visible := d["visible"].(bool)
	
	return &models.AlertWidgetColumn {
		ColumnKey: columnKey,
		ColumnLabel: columnLabel,
		ColumnSize: columnSize,
		ComponentName: componentName,
		DataPath: dataPath,
		IsSortable: isSortable,
		MinSize: minSize,
		PassthroughProps: passthroughProps,
		Visible: visible,
	}
}

func GetAlertWidgetColumnPropertyFields() (t []string) {
	return []string{
		"column_key",
		"column_label",
		"column_size",
		"component_name",
		"data_path",
		"is_sortable",
		"min_size",
		"passthrough_props",
		"visible",
	}
}