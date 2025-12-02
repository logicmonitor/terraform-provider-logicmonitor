package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func PrivilegeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"object_id": {
			Type: schema.TypeString,
			Required: true,
		},
		
		"object_name": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"object_type": {
			Type: schema.TypeString,
			Required: true,
		},
		
		"operation": {
			Type: schema.TypeString,
			Required: true,
		},
		
		"sub_operation": {
			Type: schema.TypeString,
			Computed: true,
		},
		
	}
}

func SetPrivilegeSubResourceData(m []*models.Privilege) (d []*map[string]interface{}) {
	for _, privilege := range m {
		if privilege != nil {
			properties := make(map[string]interface{})
			properties["object_id"] = privilege.ObjectID
			properties["object_name"] = privilege.ObjectName
			properties["object_type"] = privilege.ObjectType
			properties["operation"] = privilege.Operation
			properties["sub_operation"] = privilege.SubOperation
			d = append(d, &properties)
		}
	}
	return
}

func PrivilegeModel(d map[string]interface{}) *models.Privilege {
	// assume that the incoming map only contains the relevant resource data
	objectID := d["object_id"].(string)
	objectType := d["object_type"].(string)
	operation := d["operation"].(string)
	
	return &models.Privilege {
		ObjectID: &objectID,
		ObjectType: &objectType,
		Operation: &operation,
	}
}

func GetPrivilegePropertyFields() (t []string) {
	return []string{
		"object_id",
		"object_type",
		"operation",
	}
}