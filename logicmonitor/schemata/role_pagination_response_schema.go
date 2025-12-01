package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func RolePaginationResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"items": {
			Type: schema.TypeList, //GoType: []*Role  
			Elem: &schema.Resource{
				Schema: RoleSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"search_id": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"total": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
	}
}

func SetRolePaginationResponseSubResourceData(m []*models.RolePaginationResponse) (d []*map[string]interface{}) {
	for _, rolePaginationResponse := range m {
		if rolePaginationResponse != nil {
			properties := make(map[string]interface{})
			properties["items"] = SetRoleSubResourceData(rolePaginationResponse.Items)
			properties["search_id"] = rolePaginationResponse.SearchID
			properties["total"] = rolePaginationResponse.Total
			d = append(d, &properties)
		}
	}
	return
}

func RolePaginationResponseModel(d map[string]interface{}) *models.RolePaginationResponse {
	// assume that the incoming map only contains the relevant resource data
	items := d["items"].([]*models.Role)
	
	return &models.RolePaginationResponse {
		Items: items,
	}
}

func GetRolePaginationResponsePropertyFields() (t []string) {
	return []string{
		"items",
	}
}