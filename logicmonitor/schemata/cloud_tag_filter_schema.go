package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func CloudTagFilterSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type: schema.TypeString,
			Required: true,
		},
		
		"operation": {
			Type: schema.TypeString,
			Required: true,
		},
		
		"value": {
			Type: schema.TypeString,
			Required: true,
		},
		
	}
}

func SetCloudTagFilterSubResourceData(m []*models.CloudTagFilter) (d []*map[string]interface{}) {
	for _, cloudTagFilter := range m {
		if cloudTagFilter != nil {
			properties := make(map[string]interface{})
			properties["name"] = cloudTagFilter.Name
			properties["operation"] = cloudTagFilter.Operation
			properties["value"] = cloudTagFilter.Value
			d = append(d, &properties)
		}
	}
	return
}

func CloudTagFilterModel(d map[string]interface{}) *models.CloudTagFilter {
	// assume that the incoming map only contains the relevant resource data
	name := d["name"].(string)
	operation := d["operation"].(string)
	value := d["value"].(string)
	
	return &models.CloudTagFilter {
		Name: &name,
		Operation: &operation,
		Value: &value,
	}
}

func GetCloudTagFilterPropertyFields() (t []string) {
	return []string{
		"name",
		"operation",
		"value",
	}
}