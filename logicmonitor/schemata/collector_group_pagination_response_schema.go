package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func CollectorGroupPaginationResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"items": {
			Type: schema.TypeList, //GoType: []*CollectorGroup 
			Elem: &schema.Resource{
				Schema: CollectorGroupSchema(),
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

func SetCollectorGroupPaginationResponseSubResourceData(m []*models.CollectorGroupPaginationResponse) (d []*map[string]interface{}) {
	for _, collectorGroupPaginationResponse := range m {
		if collectorGroupPaginationResponse != nil {
			properties := make(map[string]interface{})
			properties["items"] = SetCollectorGroupSubResourceData(collectorGroupPaginationResponse.Items)
			properties["search_id"] = collectorGroupPaginationResponse.SearchID
			properties["total"] = collectorGroupPaginationResponse.Total
			d = append(d, &properties)
		}
	}
	return
}

func CollectorGroupPaginationResponseModel(d map[string]interface{}) *models.CollectorGroupPaginationResponse {
	// assume that the incoming map only contains the relevant resource data
	items := d["items"].([]*models.CollectorGroup)
	
	return &models.CollectorGroupPaginationResponse {
		Items: items,
	}
}

func GetCollectorGroupPaginationResponsePropertyFields() (t []string) {
	return []string{
		"items",
	}
}