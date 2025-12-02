package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-logicmonitor/models"
)

func DashboardDataSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeInt,
			Computed: true,
		},

		"name": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"sharable": {
			Type:     schema.TypeBool,
			Computed: true,
		},

		"user_permission": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func SetDashboardDataSubResourceData(m []*models.DashboardData) (d []*map[string]interface{}) {
	for _, dashboardData := range m {
		if dashboardData != nil {
			properties := make(map[string]interface{})
			properties["id"] = dashboardData.ID
			properties["name"] = dashboardData.Name
			properties["sharable"] = dashboardData.Sharable
			properties["user_permission"] = dashboardData.UserPermission
			d = append(d, &properties)
		}
	}
	return
}

func DashboardDataModel(d map[string]interface{}) *models.DashboardData {
	// assume that the incoming map only contains the relevant resource data

	return &models.DashboardData{}
}

func GetDashboardDataPropertyFields() (t []string) {
	return []string{
		"id",
	}
}
