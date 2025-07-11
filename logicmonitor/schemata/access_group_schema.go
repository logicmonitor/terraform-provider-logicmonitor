package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func AccessGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"created_by": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"created_on": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"description": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"id": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"name": {
			Type: schema.TypeString,
			Required: true,
		},
		
		"tenant_id": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"updated_on": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
	}
}

func SetAccessGroupSubResourceData(m []*models.AccessGroup) (d []*map[string]interface{}) {
	for _, accessGroup := range m {
		if accessGroup != nil {
			properties := make(map[string]interface{})
			properties["created_by"] = accessGroup.CreatedBy
			properties["created_on"] = accessGroup.CreatedOn
			properties["description"] = accessGroup.Description
			properties["id"] = accessGroup.ID
			properties["name"] = accessGroup.Name
			properties["tenant_id"] = accessGroup.TenantID
			properties["updated_on"] = accessGroup.UpdatedOn
			d = append(d, &properties)
		}
	}
	return
}

func AccessGroupModel(d map[string]interface{}) *models.AccessGroup {
	// assume that the incoming map only contains the relevant resource data
	description := d["description"].(string)
	id := int32(d["id"].(int))
	name := d["name"].(string)
	tenantID := d["tenant_id"].(string)
	
	return &models.AccessGroup {
		Description: description,
		ID: id,
		Name: &name,
		TenantID: tenantID,
	}
}

func GetAccessGroupPropertyFields() (t []string) {
	return []string{
		"description",
		"id",
		"name",
		"tenant_id",
	}
}