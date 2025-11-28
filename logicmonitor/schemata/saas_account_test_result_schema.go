package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-logicmonitor/logicmonitor/utils"
	"terraform-provider-logicmonitor/models"
)

func SaasAccountTestResultSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"detail_link": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"invalid_status_urls": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"no_permission_apis": {
			Type:     schema.TypeSet,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Set:      schema.HashString,
			Optional: true,
		},

		"no_permission_service": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"non_permission_apis_errors": {
			Type:     schema.TypeSet,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Set:      schema.HashString,
			Optional: true,
		},

		"result_code": {
			Type:     schema.TypeInt,
			Optional: true,
		},
	}
}

func SetSaasAccountTestResultSubResourceData(m []*models.SaasAccountTestResult) (d []*map[string]interface{}) {
	for _, saasAccountTestResult := range m {
		if saasAccountTestResult != nil {
			properties := make(map[string]interface{})
			properties["detail_link"] = saasAccountTestResult.DetailLink
			properties["invalid_status_urls"] = saasAccountTestResult.InvalidStatusUrls
			properties["no_permission_apis"] = saasAccountTestResult.NoPermissionApis
			properties["no_permission_service"] = saasAccountTestResult.NoPermissionService
			properties["non_permission_apis_errors"] = saasAccountTestResult.NonPermissionApisErrors
			properties["result_code"] = saasAccountTestResult.ResultCode
			d = append(d, &properties)
		}
	}
	return
}

func SaasAccountTestResultModel(d map[string]interface{}) *models.SaasAccountTestResult {
	// assume that the incoming map only contains the relevant resource data
	detailLink := d["detail_link"].(string)
	invalidStatusUrls := d["invalid_status_urls"].(string)
	noPermissionApis := utils.ConvertSetToStringSlice(d["no_permission_apis"].(*schema.Set))
	noPermissionService := d["no_permission_service"].(string)
	nonPermissionApisErrors := utils.ConvertSetToStringSlice(d["non_permission_apis_errors"].(*schema.Set))
	resultCode := int32(d["result_code"].(int))

	return &models.SaasAccountTestResult{
		DetailLink:              detailLink,
		InvalidStatusUrls:       invalidStatusUrls,
		NoPermissionApis:        noPermissionApis,
		NoPermissionService:     noPermissionService,
		NonPermissionApisErrors: nonPermissionApisErrors,
		ResultCode:              resultCode,
	}
}

func GetSaasAccountTestResultPropertyFields() (t []string) {
	return []string{
		"detail_link",
		"invalid_status_urls",
		"no_permission_apis",
		"no_permission_service",
		"non_permission_apis_errors",
		"result_code",
	}
}
