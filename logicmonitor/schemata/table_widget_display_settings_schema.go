package schemata

import (
	"terraform-provider-logicmonitor/logicmonitor/utils"
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TableWidgetDisplaySettingsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"columns_v4": {
			Type: schema.TypeList, //GoType: []*TableWidgetColumn  
			Elem: &schema.Resource{
				Schema: TableWidgetColumnSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"override_filter": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"page_size": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"property_columns": {
			Type: schema.TypeSet,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Set:      schema.HashString,
			Optional: true,
		},
		
		"show_filter": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
	}
}

func SetTableWidgetDisplaySettingsSubResourceData(m []*models.TableWidgetDisplaySettings) (d []*map[string]interface{}) {
	for _, tableWidgetDisplaySettings := range m {
		if tableWidgetDisplaySettings != nil {
			properties := make(map[string]interface{})
			properties["columns_v4"] = SetTableWidgetColumnSubResourceData(tableWidgetDisplaySettings.ColumnsV4)
			properties["override_filter"] = tableWidgetDisplaySettings.OverrideFilter
			properties["page_size"] = tableWidgetDisplaySettings.PageSize
			properties["property_columns"] = tableWidgetDisplaySettings.PropertyColumns
			properties["show_filter"] = tableWidgetDisplaySettings.ShowFilter
			d = append(d, &properties)
		}
	}
	return
}

func TableWidgetDisplaySettingsModel(d map[string]interface{}) *models.TableWidgetDisplaySettings {
	// assume that the incoming map only contains the relevant resource data
	columnsV4 := d["columns_v4"].([]*models.TableWidgetColumn)
	overrideFilter := d["override_filter"].(bool)
	pageSize := d["page_size"].(string)
	propertyColumns := utils.ConvertSetToStringSlice(d["property_columns"].(*schema.Set))
	showFilter := d["show_filter"].(bool)
	
	return &models.TableWidgetDisplaySettings {
		ColumnsV4: columnsV4,
		OverrideFilter: overrideFilter,
		PageSize: pageSize,
		PropertyColumns: propertyColumns,
		ShowFilter: showFilter,
	}
}

func GetTableWidgetDisplaySettingsPropertyFields() (t []string) {
	return []string{
		"columns_v4",
		"override_filter",
		"page_size",
		"property_columns",
		"show_filter",
	}
}