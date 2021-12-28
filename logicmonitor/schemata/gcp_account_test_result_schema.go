package schemata

import (
	"terraform-provider-logicmonitor/logicmonitor/utils"
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func GcpAccountTestResultSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"detail_link": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"no_permission_services": {
			Type: schema.TypeSet,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Set:      schema.HashString,
			Optional: true,
		},
		
		"non_permission_errors": {
			Type: schema.TypeSet,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Set:      schema.HashString,
			Optional: true,
		},
		
	}
}

func SetGcpAccountTestResultSubResourceData(m []*models.GcpAccountTestResult) (d []*map[string]interface{}) {
	for _, gcpAccountTestResult := range m {
		if gcpAccountTestResult != nil {
			properties := make(map[string]interface{})
			properties["detail_link"] = gcpAccountTestResult.DetailLink
			properties["no_permission_services"] = gcpAccountTestResult.NoPermissionServices
			properties["non_permission_errors"] = gcpAccountTestResult.NonPermissionErrors
			d = append(d, &properties)
		}
	}
	return
}

func GcpAccountTestResultModel(d map[string]interface{}) *models.GcpAccountTestResult {
	// assume that the incoming map only contains the relevant resource data
	detailLink := d["detail_link"].(string)
	noPermissionServices := utils.ConvertSetToStringSlice(d["no_permission_services"].(*schema.Set))
	nonPermissionErrors := utils.ConvertSetToStringSlice(d["non_permission_errors"].(*schema.Set))
	
	return &models.GcpAccountTestResult {
		DetailLink: detailLink,
		NoPermissionServices: noPermissionServices,
		NonPermissionErrors: nonPermissionErrors,
	}
}

func GetGcpAccountTestResultPropertyFields() (t []string) {
	return []string{
		"detail_link",
		"no_permission_services",
		"non_permission_errors",
	}
}