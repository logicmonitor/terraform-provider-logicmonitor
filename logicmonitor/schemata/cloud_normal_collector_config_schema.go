package schemata

import (
	"terraform-provider-logicmonitor/logicmonitor/utils"
	"terraform-provider-logicmonitor/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func CloudNormalCollectorConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"collectors": {
			Type: schema.TypeList, //GoType: []*CloudCollectorConfig
			Elem: &schema.Resource{
				Schema: CloudCollectorConfigSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"enable": {
			Type:     schema.TypeBool,
			Required: true,
		},
	}
}

func SetCloudNormalCollectorConfigSubResourceData(m []*models.CloudNormalCollectorConfig) (d []*map[string]interface{}) {
	for _, cloudNormalCollectorConfig := range m {
		if cloudNormalCollectorConfig != nil {
			properties := make(map[string]interface{})
			properties["collectors"] = SetCloudCollectorConfigSubResourceData(cloudNormalCollectorConfig.Collectors)
			// Dereference the pointer - default to false if nil (should not happen with corrected JSON tag)
			enable := false
			if cloudNormalCollectorConfig.Enable != nil {
				enable = *cloudNormalCollectorConfig.Enable
			}
			properties["enable"] = enable
			d = append(d, &properties)
		}
	}
	return
}

func CloudNormalCollectorConfigModel(d map[string]interface{}) *models.CloudNormalCollectorConfig {
	// assume that the incoming map only contains the relevant resource data
	collectors := utils.GetCloudCollectorConfigs(d["collectors"].([]interface{}))
	enable := d["enable"].(bool)

	return &models.CloudNormalCollectorConfig{
		Collectors: collectors,
		Enable:     &enable,
	}
}

func GetCloudNormalCollectorConfigPropertyFields() (t []string) {
	return []string{
		"collectors",
		"enable",
	}
}
