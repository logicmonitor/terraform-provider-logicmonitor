package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-logicmonitor/models"
)

func CollectorPaginationResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"items": {
			Type: schema.TypeList, //GoType: []*Collector
			Elem: &schema.Resource{
				Schema: CollectorSchema(),
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

func SetCollectorPaginationResponseSubResourceData(m []*models.CollectorPaginationResponse) (d []*map[string]interface{}) {
	for _, collectorPaginationResponse := range m {
		if collectorPaginationResponse != nil {
			properties := make(map[string]interface{})
			properties["items"] = SetCollectorSubResourceData(collectorPaginationResponse.Items)
			properties["search_id"] = collectorPaginationResponse.SearchID
			properties["total"] = collectorPaginationResponse.Total
			d = append(d, &properties)
		}
	}
	return
}

func CollectorPaginationResponseModel(d map[string]interface{}) *models.CollectorPaginationResponse {
	// assume that the incoming map only contains the relevant resource data
	items := d["items"].([]*models.Collector)

	return &models.CollectorPaginationResponse{
		Items: items,
	}
}

func GetCollectorPaginationResponsePropertyFields() (t []string) {
	return []string{
		"items",
	}
}
