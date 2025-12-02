package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-logicmonitor/models"
)

func SdtPaginationResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"items": {
			Type: schema.TypeList, //GoType: []*Sdt
			Elem: &schema.Resource{
				Schema: SdtSchema(),
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

func SetSdtPaginationResponseSubResourceData(m []*models.SdtPaginationResponse) (d []*map[string]interface{}) {
	for _, sdtPaginationResponse := range m {
		if sdtPaginationResponse != nil {
			properties := make(map[string]interface{})
			properties["items"] = SetSdtSubResourceData(sdtPaginationResponse.Items)
			properties["search_id"] = sdtPaginationResponse.SearchID
			properties["total"] = sdtPaginationResponse.Total
			d = append(d, &properties)
		}
	}
	return
}

func SdtPaginationResponseModel(d map[string]interface{}) *models.SdtPaginationResponse {
	// assume that the incoming map only contains the relevant resource data
	items := d["items"].([]*models.Sdt)

	return &models.SdtPaginationResponse{
		Items: items,
	}
}

func GetSdtPaginationResponsePropertyFields() (t []string) {
	return []string{
		"items",
	}
}
