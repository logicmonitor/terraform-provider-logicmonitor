package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DeviceGroupClusterAlertConfPaginationResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"items": {
			Type: schema.TypeList, //GoType: []*DeviceGroupClusterAlertConf  
			Elem: &schema.Resource{
				Schema: DeviceGroupClusterAlertConfSchema(),
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

func SetDeviceGroupClusterAlertConfPaginationResponseSubResourceData(m []*models.DeviceGroupClusterAlertConfPaginationResponse) (d []*map[string]interface{}) {
	for _, deviceGroupClusterAlertConfPaginationResponse := range m {
		if deviceGroupClusterAlertConfPaginationResponse != nil {
			properties := make(map[string]interface{})
			properties["items"] = SetDeviceGroupClusterAlertConfSubResourceData(deviceGroupClusterAlertConfPaginationResponse.Items)
			properties["search_id"] = deviceGroupClusterAlertConfPaginationResponse.SearchID
			properties["total"] = deviceGroupClusterAlertConfPaginationResponse.Total
			d = append(d, &properties)
		}
	}
	return
}

func DeviceGroupClusterAlertConfPaginationResponseModel(d map[string]interface{}) *models.DeviceGroupClusterAlertConfPaginationResponse {
	// assume that the incoming map only contains the relevant resource data
	items := d["items"].([]*models.DeviceGroupClusterAlertConf)
	
	return &models.DeviceGroupClusterAlertConfPaginationResponse {
		Items: items,
	}
}

func GetDeviceGroupClusterAlertConfPaginationResponsePropertyFields() (t []string) {
	return []string{
		"items",
	}
}