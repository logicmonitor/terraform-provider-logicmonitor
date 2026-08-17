package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TableWidgetDataPointSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"data_point_id": {
			Type: schema.TypeInt,
			Required: true,
		},
		
		"data_point_name": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"data_source_full_name": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"data_source_id": {
			Type: schema.TypeString,
			Required: true,
		},
		
		"is_multiple": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
	}
}

func SetTableWidgetDataPointSubResourceData(m []*models.TableWidgetDataPoint) (d []*map[string]interface{}) {
	for _, tableWidgetDataPoint := range m {
		if tableWidgetDataPoint != nil {
			properties := make(map[string]interface{})
			properties["data_point_id"] = tableWidgetDataPoint.DataPointID
			properties["data_point_name"] = tableWidgetDataPoint.DataPointName
			properties["data_source_full_name"] = tableWidgetDataPoint.DataSourceFullName
			properties["data_source_id"] = tableWidgetDataPoint.DataSourceID
			properties["is_multiple"] = tableWidgetDataPoint.IsMultiple
			d = append(d, &properties)
		}
	}
	return
}

func TableWidgetDataPointModel(d map[string]interface{}) *models.TableWidgetDataPoint {
	// assume that the incoming map only contains the relevant resource data
	dataPointID := int32(d["data_point_id"].(int))
	dataPointName := d["data_point_name"].(string)
	dataSourceFullName := d["data_source_full_name"].(string)
	dataSourceID := d["data_source_id"].(string)
	isMultiple := d["is_multiple"].(bool)
	
	return &models.TableWidgetDataPoint {
		DataPointID: &dataPointID,
		DataPointName: dataPointName,
		DataSourceFullName: dataSourceFullName,
		DataSourceID: &dataSourceID,
		IsMultiple: isMultiple,
	}
}

func GetTableWidgetDataPointPropertyFields() (t []string) {
	return []string{
		"data_point_id",
		"data_point_name",
		"data_source_full_name",
		"data_source_id",
		"is_multiple",
	}
}