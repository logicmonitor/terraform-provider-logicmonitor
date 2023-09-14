package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func WebsitePaginationResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"items": {
			Type: schema.TypeList, //GoType: []*Website  
			Elem: &schema.Resource{
				Schema: WebsiteSchema(),
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

func SetWebsitePaginationResponseSubResourceData(m []*models.WebsitePaginationResponse) (d []*map[string]interface{}) {
	for _, websitePaginationResponse := range m {
		if websitePaginationResponse != nil {
			properties := make(map[string]interface{})
			properties["items"] = SetWebsiteSubResourceData(websitePaginationResponse.Items)
			properties["search_id"] = websitePaginationResponse.SearchID
			properties["total"] = websitePaginationResponse.Total
			d = append(d, &properties)
		}
	}
	return
}

func WebsitePaginationResponseModel(d map[string]interface{}) *models.WebsitePaginationResponse {
	// assume that the incoming map only contains the relevant resource data
	items := d["items"].([]*models.Website)
	
	return &models.WebsitePaginationResponse {
		Items: items,
	}
}

func GetWebsitePaginationResponsePropertyFields() (t []string) {
	return []string{
		"items",
	}
}