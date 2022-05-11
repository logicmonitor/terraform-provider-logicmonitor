package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DashboardGroupPaginationResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"items": {
			Type: schema.TypeList, //GoType: []*DashboardGroup 
			Elem: &schema.Resource{
				Schema: DashboardGroupSchema(),
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

func SetDashboardGroupPaginationResponseSubResourceData(m []*models.DashboardGroupPaginationResponse) (d []*map[string]interface{}) {
	for _, dashboardGroupPaginationResponse := range m {
		if dashboardGroupPaginationResponse != nil {
			properties := make(map[string]interface{})
			properties["items"] = SetDashboardGroupSubResourceData(dashboardGroupPaginationResponse.Items)
			properties["search_id"] = dashboardGroupPaginationResponse.SearchID
			properties["total"] = dashboardGroupPaginationResponse.Total
			d = append(d, &properties)
		}
	}
	return
}

func DashboardGroupPaginationResponseModel(d map[string]interface{}) *models.DashboardGroupPaginationResponse {
	// assume that the incoming map only contains the relevant resource data
	items := d["items"].([]*models.DashboardGroup)
	
	return &models.DashboardGroupPaginationResponse {
		Items: items,
	}
}

func GetDashboardGroupPaginationResponsePropertyFields() (t []string) {
	return []string{
		"items",
	}
}