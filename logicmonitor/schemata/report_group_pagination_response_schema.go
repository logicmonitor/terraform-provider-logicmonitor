package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ReportGroupPaginationResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"items": {
			Type: schema.TypeList, //GoType: []*ReportGroup  
			Elem: &schema.Resource{
				Schema: ReportGroupSchema(),
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

func SetReportGroupPaginationResponseSubResourceData(m []*models.ReportGroupPaginationResponse) (d []*map[string]interface{}) {
	for _, reportGroupPaginationResponse := range m {
		if reportGroupPaginationResponse != nil {
			properties := make(map[string]interface{})
			properties["items"] = SetReportGroupSubResourceData(reportGroupPaginationResponse.Items)
			properties["search_id"] = reportGroupPaginationResponse.SearchID
			properties["total"] = reportGroupPaginationResponse.Total
			d = append(d, &properties)
		}
	}
	return
}

func ReportGroupPaginationResponseModel(d map[string]interface{}) *models.ReportGroupPaginationResponse {
	// assume that the incoming map only contains the relevant resource data
	items := d["items"].([]*models.ReportGroup)
	
	return &models.ReportGroupPaginationResponse {
		Items: items,
	}
}

func GetReportGroupPaginationResponsePropertyFields() (t []string) {
	return []string{
		"items",
	}
}