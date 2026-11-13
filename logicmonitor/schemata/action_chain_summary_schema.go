package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ActionChainSummarySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"name": {
			Type: schema.TypeString,
			Computed: true,
		},
		
	}
}

func SetActionChainSummarySubResourceData(m []*models.ActionChainSummary) (d []*map[string]interface{}) {
	for _, actionChainSummary := range m {
		if actionChainSummary != nil {
			properties := make(map[string]interface{})
			properties["id"] = actionChainSummary.ID
			properties["name"] = actionChainSummary.Name
			d = append(d, &properties)
		}
	}
	return
}

func ActionChainSummaryModel(d map[string]interface{}) *models.ActionChainSummary {
	// assume that the incoming map only contains the relevant resource data
	
	return &models.ActionChainSummary {
	}
}

func GetActionChainSummaryPropertyFields() (t []string) {
	return []string{
		"id",
	}
}