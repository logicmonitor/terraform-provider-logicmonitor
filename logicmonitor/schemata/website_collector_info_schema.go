package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func WebsiteCollectorInfoSchema() map[string]*schema.Schema {
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

func SetWebsiteCollectorInfoSubResourceData(m []*models.WebsiteCollectorInfo) (d []*map[string]interface{}) {
	for _, websiteCollectorInfo := range m {
		if websiteCollectorInfo != nil {
			properties := make(map[string]interface{})
			properties["collector_group_id"] = websiteCollectorInfo.CollectorGroupID
			properties["collector_group_name"] = websiteCollectorInfo.CollectorGroupName
			properties["description"] = websiteCollectorInfo.Description
			properties["hostname"] = websiteCollectorInfo.Hostname
			properties["id"] = websiteCollectorInfo.ID
			properties["status"] = websiteCollectorInfo.Status
			d = append(d, &properties)
		}
	}
	return
}

func WebsiteCollectorInfoModel(d map[string]interface{}) *models.WebsiteCollectorInfo {
	// assume that the incoming map only contains the relevant resource data
	
	return &models.WebsiteCollectorInfo {
	}
}

func GetWebsiteCollectorInfoPropertyFields() (t []string) {
	return []string{
		"id",
	}
}