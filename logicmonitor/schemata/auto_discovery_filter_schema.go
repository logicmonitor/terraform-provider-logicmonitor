package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-logicmonitor/models"
)

func AutoDiscoveryFilterSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"attribute": {
			Type:     schema.TypeString,
			Required: true,
		},

		"comment": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"operation": {
			Type:     schema.TypeString,
			Required: true,
		},

		"value": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func SetAutoDiscoveryFilterSubResourceData(m []*models.AutoDiscoveryFilter) (d []*map[string]interface{}) {
	for _, autoDiscoveryFilter := range m {
		if autoDiscoveryFilter != nil {
			properties := make(map[string]interface{})
			properties["attribute"] = autoDiscoveryFilter.Attribute
			properties["comment"] = autoDiscoveryFilter.Comment
			properties["operation"] = autoDiscoveryFilter.Operation
			properties["value"] = autoDiscoveryFilter.Value
			d = append(d, &properties)
		}
	}
	return
}

func AutoDiscoveryFilterModel(d map[string]interface{}) *models.AutoDiscoveryFilter {
	// assume that the incoming map only contains the relevant resource data
	attribute := d["attribute"].(string)
	comment := d["comment"].(string)
	operation := d["operation"].(string)
	value := d["value"].(string)

	return &models.AutoDiscoveryFilter{
		Attribute: &attribute,
		Comment:   comment,
		Operation: &operation,
		Value:     value,
	}
}

func GetAutoDiscoveryFilterPropertyFields() (t []string) {
	return []string{
		"attribute",
		"comment",
		"operation",
		"value",
	}
}
