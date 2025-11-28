package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-logicmonitor/models"
)

func DeviceGroupPaginationResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"items": {
			Type: schema.TypeList, //GoType: []*DeviceGroup
			Elem: &schema.Resource{
				Schema: DeviceGroupSchema(),
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

func SetDeviceGroupPaginationResponseSubResourceData(m []*models.DeviceGroupPaginationResponse) (d []*map[string]interface{}) {
	for _, deviceGroupPaginationResponse := range m {
		if deviceGroupPaginationResponse != nil {
			properties := make(map[string]interface{})
			properties["items"] = SetDeviceGroupSubResourceData(deviceGroupPaginationResponse.Items)
			properties["search_id"] = deviceGroupPaginationResponse.SearchID
			properties["total"] = deviceGroupPaginationResponse.Total
			d = append(d, &properties)
		}
	}
	return
}

func DeviceGroupPaginationResponseModel(d map[string]interface{}) *models.DeviceGroupPaginationResponse {
	// assume that the incoming map only contains the relevant resource data
	items := d["items"].([]*models.DeviceGroup)

	return &models.DeviceGroupPaginationResponse{
		Items: items,
	}
}

func GetDeviceGroupPaginationResponsePropertyFields() (t []string) {
	return []string{
		"items",
	}
}
