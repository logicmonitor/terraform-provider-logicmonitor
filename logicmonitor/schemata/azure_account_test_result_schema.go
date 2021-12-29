package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func AzureAccountTestResultSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"detail_link": {
			Type: schema.TypeMap, //GoType: interface{}
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Computed: true,
		},
		
		"no_permission_services": {
			Type: schema.TypeMap, //GoType: interface{}
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Computed: true,
		},
		
	}
}

func SetAzureAccountTestResultSubResourceData(m []*models.AzureAccountTestResult) (d []*map[string]interface{}) {
	for _, azureAccountTestResult := range m {
		if azureAccountTestResult != nil {
			properties := make(map[string]interface{})
			properties["detail_link"] = azureAccountTestResult.DetailLink
			properties["no_permission_services"] = azureAccountTestResult.NoPermissionServices
			d = append(d, &properties)
		}
	}
	return
}

func AzureAccountTestResultModel(d map[string]interface{}) *models.AzureAccountTestResult {
	// assume that the incoming map only contains the relevant resource data
	
	return &models.AzureAccountTestResult {
	}
}

func GetAzureAccountTestResultPropertyFields() (t []string) {
	return []string{
	}
}