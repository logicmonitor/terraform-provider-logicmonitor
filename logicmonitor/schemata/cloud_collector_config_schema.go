package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-logicmonitor/models"
)

func CloudCollectorConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"applies_to": {
			Type:     schema.TypeString,
			Required: true,
		},

		"auto_balanced_collector_group_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"collector_id": {
			Type:     schema.TypeInt,
			Required: true,
		},

		"priority": {
			Type:     schema.TypeInt,
			Required: true,
		},

		"use_public_ip": {
			Type:     schema.TypeBool,
			Optional: true,
		},
	}
}

func SetCloudCollectorConfigSubResourceData(m []*models.CloudCollectorConfig) (d []*map[string]interface{}) {
	for _, cloudCollectorConfig := range m {
		if cloudCollectorConfig != nil {
			properties := make(map[string]interface{})
			properties["applies_to"] = cloudCollectorConfig.AppliesTo
			properties["auto_balanced_collector_group_id"] = cloudCollectorConfig.AutoBalancedCollectorGroupID
			properties["collector_id"] = cloudCollectorConfig.CollectorID
			properties["priority"] = cloudCollectorConfig.Priority
			properties["use_public_ip"] = cloudCollectorConfig.UsePublicIP
			d = append(d, &properties)
		}
	}
	return
}

func CloudCollectorConfigModel(d map[string]interface{}) *models.CloudCollectorConfig {
	// assume that the incoming map only contains the relevant resource data
	appliesTo := d["applies_to"].(string)
	autoBalancedCollectorGroupID := int32(d["auto_balanced_collector_group_id"].(int))
	collectorID := int32(d["collector_id"].(int))
	priority := int32(d["priority"].(int))
	usePublicIP := d["use_public_ip"].(bool)

	return &models.CloudCollectorConfig{
		AppliesTo:                    &appliesTo,
		AutoBalancedCollectorGroupID: autoBalancedCollectorGroupID,
		CollectorID:                  &collectorID,
		Priority:                     &priority,
		UsePublicIP:                  usePublicIP,
	}
}

func GetCloudCollectorConfigPropertyFields() (t []string) {
	return []string{
		"applies_to",
		"auto_balanced_collector_group_id",
		"collector_id",
		"priority",
		"use_public_ip",
	}
}
