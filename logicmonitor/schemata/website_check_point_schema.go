package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func WebsiteCheckPointSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"geo_info": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"id": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"smg_id": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
	}
}

func SetWebsiteCheckPointSubResourceData(m []*models.WebsiteCheckPoint) (d []*map[string]interface{}) {
	for _, websiteCheckPoint := range m {
		if websiteCheckPoint != nil {
			properties := make(map[string]interface{})
			properties["geo_info"] = websiteCheckPoint.GeoInfo
			properties["id"] = websiteCheckPoint.ID
			properties["smg_id"] = websiteCheckPoint.SmgID
			d = append(d, &properties)
		}
	}
	return
}

func WebsiteCheckPointModel(d map[string]interface{}) *models.WebsiteCheckPoint {
	// assume that the incoming map only contains the relevant resource data
	
	return &models.WebsiteCheckPoint {
	}
}

func GetWebsiteCheckPointPropertyFields() (t []string) {
	return []string{
		"id",
	}
}