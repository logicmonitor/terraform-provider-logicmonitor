package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func WebsiteDeviceCollectorInfoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"collector_group_id": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"collector_group_name": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"description": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"hostname": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"id": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"status": {
			Type: schema.TypeString,
			Computed: true,
		},
		
	}
}

func SetWebsiteDeviceCollectorInfoSubResourceData(m []*models.WebsiteDeviceCollectorInfo) (d []*map[string]interface{}) {
	for _, websiteDeviceCollectorInfo := range m {
		if websiteDeviceCollectorInfo != nil {
			properties := make(map[string]interface{})
			properties["collector_group_id"] = websiteDeviceCollectorInfo.CollectorGroupID
			properties["collector_group_name"] = websiteDeviceCollectorInfo.CollectorGroupName
			properties["description"] = websiteDeviceCollectorInfo.Description
			properties["hostname"] = websiteDeviceCollectorInfo.Hostname
			properties["id"] = websiteDeviceCollectorInfo.ID
			properties["status"] = websiteDeviceCollectorInfo.Status
			d = append(d, &properties)
		}
	}
	return
}

func WebsiteDeviceCollectorInfoModel(d map[string]interface{}) *models.WebsiteDeviceCollectorInfo {
	// assume that the incoming map only contains the relevant resource data
	
	return &models.WebsiteDeviceCollectorInfo {
	}
}

func GetWebsiteDeviceCollectorInfoPropertyFields() (t []string) {
	return []string{
		"id",
	}
}