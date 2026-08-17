package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TableWidgetRowSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"device_display_name": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"device_id": {
			Type: schema.TypeInt,
			Required: true,
		},
		
		"group_full_path": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"group_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"id": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"instances": {
			Type: schema.TypeList, //GoType: []*TableWidgetInstanceCell  
			Elem: &schema.Resource{
				Schema: TableWidgetInstanceCellSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"label": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"position": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
	}
}

func SetTableWidgetRowSubResourceData(m []*models.TableWidgetRow) (d []*map[string]interface{}) {
	for _, tableWidgetRow := range m {
		if tableWidgetRow != nil {
			properties := make(map[string]interface{})
			properties["device_display_name"] = tableWidgetRow.DeviceDisplayName
			properties["device_id"] = tableWidgetRow.DeviceID
			properties["group_full_path"] = tableWidgetRow.GroupFullPath
			properties["group_id"] = tableWidgetRow.GroupID
			properties["id"] = tableWidgetRow.ID
			properties["instances"] = SetTableWidgetInstanceCellSubResourceData(tableWidgetRow.Instances)
			properties["label"] = tableWidgetRow.Label
			properties["position"] = tableWidgetRow.Position
			d = append(d, &properties)
		}
	}
	return
}

func TableWidgetRowModel(d map[string]interface{}) *models.TableWidgetRow {
	// assume that the incoming map only contains the relevant resource data
	deviceDisplayName := d["device_display_name"].(string)
	deviceID := int32(d["device_id"].(int))
	groupFullPath := d["group_full_path"].(string)
	groupID := int32(d["group_id"].(int))
	id := d["id"].(string)
	instances := d["instances"].([]*models.TableWidgetInstanceCell)
	label := d["label"].(string)
	position := int32(d["position"].(int))
	
	return &models.TableWidgetRow {
		DeviceDisplayName: deviceDisplayName,
		DeviceID: &deviceID,
		GroupFullPath: groupFullPath,
		GroupID: groupID,
		ID: id,
		Instances: instances,
		Label: label,
		Position: position,
	}
}

func GetTableWidgetRowPropertyFields() (t []string) {
	return []string{
		"device_display_name",
		"device_id",
		"group_full_path",
		"group_id",
		"id",
		"instances",
		"label",
		"position",
	}
}