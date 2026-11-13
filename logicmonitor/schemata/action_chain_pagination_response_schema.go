package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ActionChainPaginationResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"items": {
			Type: schema.TypeList, //GoType: []*ActionChain  
			Elem: &schema.Resource{
				Schema: ActionChainSchema(),
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

func SetActionChainPaginationResponseSubResourceData(m []*models.ActionChainPaginationResponse) (d []*map[string]interface{}) {
	for _, actionChainPaginationResponse := range m {
		if actionChainPaginationResponse != nil {
			properties := make(map[string]interface{})
			properties["items"] = SetActionChainSubResourceData(actionChainPaginationResponse.Items)
			properties["search_id"] = actionChainPaginationResponse.SearchID
			properties["total"] = actionChainPaginationResponse.Total
			d = append(d, &properties)
		}
	}
	return
}

func ActionChainPaginationResponseModel(d map[string]interface{}) *models.ActionChainPaginationResponse {
	// assume that the incoming map only contains the relevant resource data
	items := d["items"].([]*models.ActionChain)
	
	return &models.ActionChainPaginationResponse {
		Items: items,
	}
}

func GetActionChainPaginationResponsePropertyFields() (t []string) {
	return []string{
		"items",
	}
}