package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-logicmonitor/models"
)

func WebsiteGroupPaginationResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"items": {
			Type: schema.TypeList, //GoType: []*WebsiteGroup
			Elem: &schema.Resource{
				Schema: WebsiteGroupSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"search_id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"total": {
			Type:     schema.TypeInt,
			Computed: true,
		},
	}
}

func SetWebsiteGroupPaginationResponseSubResourceData(m []*models.WebsiteGroupPaginationResponse) (d []*map[string]interface{}) {
	for _, websiteGroupPaginationResponse := range m {
		if websiteGroupPaginationResponse != nil {
			properties := make(map[string]interface{})
			properties["items"] = SetWebsiteGroupSubResourceData(websiteGroupPaginationResponse.Items)
			properties["search_id"] = websiteGroupPaginationResponse.SearchID
			properties["total"] = websiteGroupPaginationResponse.Total
			d = append(d, &properties)
		}
	}
	return
}

func WebsiteGroupPaginationResponseModel(d map[string]interface{}) *models.WebsiteGroupPaginationResponse {
	// assume that the incoming map only contains the relevant resource data
	items := d["items"].([]*models.WebsiteGroup)

	return &models.WebsiteGroupPaginationResponse{
		Items: items,
	}
}

func GetWebsiteGroupPaginationResponsePropertyFields() (t []string) {
	return []string{
		"items",
	}
}
