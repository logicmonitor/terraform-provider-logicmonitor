package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TableWidgetColumnSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"alternate_data_points": {
			Type: schema.TypeList, //GoType: []*TableWidgetDataPoint  
			Elem: &schema.Resource{
				Schema: TableWidgetDataPointSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"column_name": {
			Type: schema.TypeString,
			Required: true,
		},
		
		"data_point": {
			Type: schema.TypeList, //GoType: TableWidgetDataPoint
			Elem: &schema.Resource{
				Schema: TableWidgetDataPointSchema(),
			},
			Required: true,
		},
		
		"enable_forecast": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"id": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"index": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"position": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"properties_options": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"rounding_decimal": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"rpn": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"unit_label": {
			Type: schema.TypeString,
			Optional: true,
		},
		
	}
}

func SetTableWidgetColumnSubResourceData(m []*models.TableWidgetColumn) (d []*map[string]interface{}) {
	for _, tableWidgetColumn := range m {
		if tableWidgetColumn != nil {
			properties := make(map[string]interface{})
			properties["alternate_data_points"] = SetTableWidgetDataPointSubResourceData(tableWidgetColumn.AlternateDataPoints)
			properties["column_name"] = tableWidgetColumn.ColumnName
			properties["data_point"] = SetTableWidgetDataPointSubResourceData([]*models.TableWidgetDataPoint{tableWidgetColumn.DataPoint})
			properties["enable_forecast"] = tableWidgetColumn.EnableForecast
			properties["id"] = tableWidgetColumn.ID
			properties["index"] = tableWidgetColumn.Index
			properties["position"] = tableWidgetColumn.Position
			properties["properties_options"] = tableWidgetColumn.PropertiesOptions
			properties["rounding_decimal"] = tableWidgetColumn.RoundingDecimal
			properties["rpn"] = tableWidgetColumn.Rpn
			properties["unit_label"] = tableWidgetColumn.UnitLabel
			d = append(d, &properties)
		}
	}
	return
}

func TableWidgetColumnModel(d map[string]interface{}) *models.TableWidgetColumn {
	// assume that the incoming map only contains the relevant resource data
	alternateDataPoints := d["alternate_data_points"].([]*models.TableWidgetDataPoint)
	columnName := d["column_name"].(string)
	var dataPoint *models.TableWidgetDataPoint = nil
	dataPointList := d["data_point"].([]interface{})
	if len(dataPointList) > 0 { // len(nil) = 0
		dataPoint = TableWidgetDataPointModel(dataPointList[0].(map[string]interface{}))
	}
	enableForecast := d["enable_forecast"].(bool)
	id := d["id"].(string)
	index := int32(d["index"].(int))
	position := int32(d["position"].(int))
	propertiesOptions := d["properties_options"].(string)
	roundingDecimal := int32(d["rounding_decimal"].(int))
	rpn := d["rpn"].(string)
	unitLabel := d["unit_label"].(string)
	
	return &models.TableWidgetColumn {
		AlternateDataPoints: alternateDataPoints,
		ColumnName: &columnName,
		DataPoint: dataPoint,
		EnableForecast: enableForecast,
		ID: id,
		Index: index,
		Position: position,
		PropertiesOptions: propertiesOptions,
		RoundingDecimal: roundingDecimal,
		Rpn: rpn,
		UnitLabel: unitLabel,
	}
}

func GetTableWidgetColumnPropertyFields() (t []string) {
	return []string{
		"alternate_data_points",
		"column_name",
		"data_point",
		"enable_forecast",
		"id",
		"index",
		"position",
		"properties_options",
		"rounding_decimal",
		"rpn",
		"unit_label",
	}
}