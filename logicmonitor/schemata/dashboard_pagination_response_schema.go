package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DashboardPaginationResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"items": {
			Type: schema.TypeList, //GoType: []*Dashboard 
			Elem: &schema.Resource{
				Schema: DashboardSchema(),
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

func SetDashboardPaginationResponseSubResourceData(m []*models.DashboardPaginationResponse) (d []*map[string]interface{}) {
	for _, dashboardPaginationResponse := range m {
		if dashboardPaginationResponse != nil {
			properties := make(map[string]interface{})
			properties["items"] = SetDashboardSubResourceData(dashboardPaginationResponse.Items)
			properties["search_id"] = dashboardPaginationResponse.SearchID
			properties["total"] = dashboardPaginationResponse.Total
			d = append(d, &properties)
		}
	}
	return
}

func DashboardPaginationResponseModel(d map[string]interface{}) *models.DashboardPaginationResponse {
	// assume that the incoming map only contains the relevant resource data
	items := d["items"].([]*models.Dashboard)
	
	return &models.DashboardPaginationResponse {
		Items: items,
	}
}

func GetDashboardPaginationResponsePropertyFields() (t []string) {
	return []string{
		"items",
	}
}