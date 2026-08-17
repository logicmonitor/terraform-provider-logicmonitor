package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TableWidgetInstanceCellSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"data_point_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"data_point_name": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"instance_id": {
			Type: schema.TypeInt,
			Required: true,
		},
		
		"instance_name": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"validation_status_code": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
	}
}

func SetTableWidgetInstanceCellSubResourceData(m []*models.TableWidgetInstanceCell) (d []*map[string]interface{}) {
	for _, tableWidgetInstanceCell := range m {
		if tableWidgetInstanceCell != nil {
			properties := make(map[string]interface{})
			properties["data_point_id"] = tableWidgetInstanceCell.DataPointID
			properties["data_point_name"] = tableWidgetInstanceCell.DataPointName
			properties["instance_id"] = tableWidgetInstanceCell.InstanceID
			properties["instance_name"] = tableWidgetInstanceCell.InstanceName
			properties["validation_status_code"] = tableWidgetInstanceCell.ValidationStatusCode
			d = append(d, &properties)
		}
	}
	return
}

func TableWidgetInstanceCellModel(d map[string]interface{}) *models.TableWidgetInstanceCell {
	// assume that the incoming map only contains the relevant resource data
	dataPointID := int32(d["data_point_id"].(int))
	dataPointName := d["data_point_name"].(string)
	instanceID := int32(d["instance_id"].(int))
	instanceName := d["instance_name"].(string)
	validationStatusCode := int32(d["validation_status_code"].(int))
	
	return &models.TableWidgetInstanceCell {
		DataPointID: dataPointID,
		DataPointName: dataPointName,
		InstanceID: &instanceID,
		InstanceName: instanceName,
		ValidationStatusCode: validationStatusCode,
	}
}

func GetTableWidgetInstanceCellPropertyFields() (t []string) {
	return []string{
		"data_point_id",
		"data_point_name",
		"instance_id",
		"instance_name",
		"validation_status_code",
	}
}