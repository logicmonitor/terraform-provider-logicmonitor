package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-logicmonitor/models"
)

func AlertRulePaginationResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"items": {
			Type: schema.TypeList, //GoType: []*AlertRule
			Elem: &schema.Resource{
				Schema: AlertRuleSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"search_id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"total": {
			Type:     schema.TypeInt,
			Computed: true,
		},
	}
}

func SetAlertRulePaginationResponseSubResourceData(m []*models.AlertRulePaginationResponse) (d []*map[string]interface{}) {
	for _, alertRulePaginationResponse := range m {
		if alertRulePaginationResponse != nil {
			properties := make(map[string]interface{})
			properties["items"] = SetAlertRuleSubResourceData(alertRulePaginationResponse.Items)
			properties["search_id"] = alertRulePaginationResponse.SearchID
			properties["total"] = alertRulePaginationResponse.Total
			d = append(d, &properties)
		}
	}
	return
}

func AlertRulePaginationResponseModel(d map[string]interface{}) *models.AlertRulePaginationResponse {
	// assume that the incoming map only contains the relevant resource data
	items := d["items"].([]*models.AlertRule)

	return &models.AlertRulePaginationResponse{
		Items: items,
	}
}

func GetAlertRulePaginationResponsePropertyFields() (t []string) {
	return []string{
		"items",
	}
}
