package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-logicmonitor/models"
)

func DevicePaginationResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"items": {
			Type: schema.TypeList, //GoType: []*Device
			Elem: &schema.Resource{
				Schema: DeviceSchema(),
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

func SetDevicePaginationResponseSubResourceData(m []*models.DevicePaginationResponse) (d []*map[string]interface{}) {
	for _, devicePaginationResponse := range m {
		if devicePaginationResponse != nil {
			properties := make(map[string]interface{})
			properties["items"] = SetDeviceSubResourceData(devicePaginationResponse.Items)
			properties["search_id"] = devicePaginationResponse.SearchID
			properties["total"] = devicePaginationResponse.Total
			d = append(d, &properties)
		}
	}
	return
}

func DevicePaginationResponseModel(d map[string]interface{}) *models.DevicePaginationResponse {
	// assume that the incoming map only contains the relevant resource data
	items := d["items"].([]*models.Device)

	return &models.DevicePaginationResponse{
		Items: items,
	}
}

func GetDevicePaginationResponsePropertyFields() (t []string) {
	return []string{
		"items",
	}
}
