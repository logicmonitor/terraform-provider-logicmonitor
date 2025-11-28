package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-logicmonitor/logicmonitor/utils"
	"terraform-provider-logicmonitor/models"
)

func AwsAccountTestResultSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"detail_link": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"no_permission_services": {
			Type:     schema.TypeSet,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Set:      schema.HashString,
			Optional: true,
		},

		"non_permission_errors": {
			Type:     schema.TypeSet,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Set:      schema.HashString,
			Optional: true,
		},
	}
}

func SetAwsAccountTestResultSubResourceData(m []*models.AwsAccountTestResult) (d []*map[string]interface{}) {
	for _, awsAccountTestResult := range m {
		if awsAccountTestResult != nil {
			properties := make(map[string]interface{})
			properties["detail_link"] = awsAccountTestResult.DetailLink
			properties["no_permission_services"] = awsAccountTestResult.NoPermissionServices
			properties["non_permission_errors"] = awsAccountTestResult.NonPermissionErrors
			d = append(d, &properties)
		}
	}
	return
}

func AwsAccountTestResultModel(d map[string]interface{}) *models.AwsAccountTestResult {
	// assume that the incoming map only contains the relevant resource data
	detailLink := d["detail_link"].(string)
	noPermissionServices := utils.ConvertSetToStringSlice(d["no_permission_services"].(*schema.Set))
	nonPermissionErrors := utils.ConvertSetToStringSlice(d["non_permission_errors"].(*schema.Set))

	return &models.AwsAccountTestResult{
		DetailLink:           detailLink,
		NoPermissionServices: noPermissionServices,
		NonPermissionErrors:  nonPermissionErrors,
	}
}

func GetAwsAccountTestResultPropertyFields() (t []string) {
	return []string{
		"detail_link",
		"no_permission_services",
		"non_permission_errors",
	}
}
