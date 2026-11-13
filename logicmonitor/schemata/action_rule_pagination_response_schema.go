package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ActionRulePaginationResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"items": {
			Type: schema.TypeList, //GoType: []*ActionRule  
			Elem: &schema.Resource{
				Schema: ActionRuleSchema(),
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

func SetActionRulePaginationResponseSubResourceData(m []*models.ActionRulePaginationResponse) (d []*map[string]interface{}) {
	for _, actionRulePaginationResponse := range m {
		if actionRulePaginationResponse != nil {
			properties := make(map[string]interface{})
			properties["items"] = SetActionRuleSubResourceData(actionRulePaginationResponse.Items)
			properties["search_id"] = actionRulePaginationResponse.SearchID
			properties["total"] = actionRulePaginationResponse.Total
			d = append(d, &properties)
		}
	}
	return
}

func ActionRulePaginationResponseModel(d map[string]interface{}) *models.ActionRulePaginationResponse {
	// assume that the incoming map only contains the relevant resource data
	items := d["items"].([]*models.ActionRule)
	
	return &models.ActionRulePaginationResponse {
		Items: items,
	}
}

func GetActionRulePaginationResponsePropertyFields() (t []string) {
	return []string{
		"items",
	}
}