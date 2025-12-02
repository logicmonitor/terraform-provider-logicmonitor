package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-logicmonitor/models"
)

func NextUpgradeInfoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"major_version": {
			Type:     schema.TypeInt,
			Computed: true,
		},

		"mandatory": {
			Type:     schema.TypeBool,
			Computed: true,
		},

		"minor_version": {
			Type:     schema.TypeInt,
			Computed: true,
		},

		"stable": {
			Type:     schema.TypeBool,
			Computed: true,
		},

		"upgrade_time": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"upgrade_time_epoch": {
			Type:     schema.TypeInt,
			Computed: true,
		},
	}
}

func SetNextUpgradeInfoSubResourceData(m []*models.NextUpgradeInfo) (d []*map[string]interface{}) {
	for _, nextUpgradeInfo := range m {
		if nextUpgradeInfo != nil {
			properties := make(map[string]interface{})
			properties["major_version"] = nextUpgradeInfo.MajorVersion
			properties["mandatory"] = nextUpgradeInfo.Mandatory
			properties["minor_version"] = nextUpgradeInfo.MinorVersion
			properties["stable"] = nextUpgradeInfo.Stable
			properties["upgrade_time"] = nextUpgradeInfo.UpgradeTime
			properties["upgrade_time_epoch"] = nextUpgradeInfo.UpgradeTimeEpoch
			d = append(d, &properties)
		}
	}
	return
}

func NextUpgradeInfoModel(d map[string]interface{}) *models.NextUpgradeInfo {
	// assume that the incoming map only contains the relevant resource data

	return &models.NextUpgradeInfo{}
}

func GetNextUpgradeInfoPropertyFields() (t []string) {
	return []string{}
}
