package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-logicmonitor/logicmonitor/utils"
	"terraform-provider-logicmonitor/models"
)

func WebsiteLocationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"all": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"collector_ids": {
			Type: schema.TypeList, //GoType: []int32
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"collectors": {
			Type: schema.TypeList, //GoType: []*WebsiteCollectorInfo
			Elem: &schema.Resource{
				Schema: WebsiteCollectorInfoSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"smg_ids": {
			Type: schema.TypeList, //GoType: []int32
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},
	}
}

func SetWebsiteLocationSubResourceData(m []*models.WebsiteLocation) (d []*map[string]interface{}) {
	for _, websiteLocation := range m {
		if websiteLocation != nil {
			properties := make(map[string]interface{})
			properties["all"] = websiteLocation.All
			properties["collector_ids"] = websiteLocation.CollectorIds
			properties["collectors"] = SetWebsiteCollectorInfoSubResourceData(websiteLocation.Collectors)
			properties["smg_ids"] = websiteLocation.SmgIds
			d = append(d, &properties)
		}
	}
	return
}

func WebsiteLocationModel(d map[string]interface{}) *models.WebsiteLocation {
	// assume that the incoming map only contains the relevant resource data
	all := d["all"].(bool)
	collectorIds := utils.ConvertSetToInt32Slice(d["collectorIds"].([]interface{}))
	collectors := d["collectors"].([]*models.WebsiteCollectorInfo)
	smgIds := utils.ConvertSetToInt32Slice(d["smgIds"].([]interface{}))

	return &models.WebsiteLocation{
		All:          all,
		CollectorIds: collectorIds,
		Collectors:   collectors,
		SmgIds:       smgIds,
	}
}

func GetWebsiteLocationPropertyFields() (t []string) {
	return []string{
		"all",
		"collector_ids",
		"collectors",
		"smg_ids",
	}
}
