package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func WebsiteNOCItemSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
	}
}

func SetWebsiteNOCItemSubResourceData(m []*models.WebsiteNOCItem) (d []*map[string]interface{}) {
	for _, websiteNOCItem := range m {
		if websiteNOCItem != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func WebsiteNOCItemModel(d map[string]interface{}) *models.WebsiteNOCItem {
	// assume that the incoming map only contains the relevant resource data
	
	return &models.WebsiteNOCItem {
	}
}

func GetWebsiteNOCItemPropertyFields() (t []string) {
	return []string{
	}
}