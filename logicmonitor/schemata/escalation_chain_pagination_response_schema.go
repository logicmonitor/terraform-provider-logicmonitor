package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-logicmonitor/models"
)

func EscalationChainPaginationResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"items": {
			Type: schema.TypeList, //GoType: []*EscalationChain
			Elem: &schema.Resource{
				Schema: EscalationChainSchema(),
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

func SetEscalationChainPaginationResponseSubResourceData(m []*models.EscalationChainPaginationResponse) (d []*map[string]interface{}) {
	for _, escalationChainPaginationResponse := range m {
		if escalationChainPaginationResponse != nil {
			properties := make(map[string]interface{})
			properties["items"] = SetEscalationChainSubResourceData(escalationChainPaginationResponse.Items)
			properties["search_id"] = escalationChainPaginationResponse.SearchID
			properties["total"] = escalationChainPaginationResponse.Total
			d = append(d, &properties)
		}
	}
	return
}

func EscalationChainPaginationResponseModel(d map[string]interface{}) *models.EscalationChainPaginationResponse {
	// assume that the incoming map only contains the relevant resource data
	items := d["items"].([]*models.EscalationChain)

	return &models.EscalationChainPaginationResponse{
		Items: items,
	}
}

func GetEscalationChainPaginationResponsePropertyFields() (t []string) {
	return []string{
		"items",
	}
}
