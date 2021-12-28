package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func OnetimeUpgradeInfoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"created_by": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"description": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"end_epoch": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"level": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"major_version": {
			Type: schema.TypeInt,
			Required: true,
		},
		
		"minor_version": {
			Type: schema.TypeInt,
			Required: true,
		},
		
		"start_epoch": {
			Type: schema.TypeInt,
			Required: true,
		},
		
		"timezone": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"type": {
			Type: schema.TypeString,
			Computed: true,
		},
		
	}
}

func SetOnetimeUpgradeInfoSubResourceData(m []*models.OnetimeUpgradeInfo) (d []*map[string]interface{}) {
	for _, onetimeUpgradeInfo := range m {
		if onetimeUpgradeInfo != nil {
			properties := make(map[string]interface{})
			properties["created_by"] = onetimeUpgradeInfo.CreatedBy
			properties["description"] = onetimeUpgradeInfo.Description
			properties["end_epoch"] = onetimeUpgradeInfo.EndEpoch
			properties["level"] = onetimeUpgradeInfo.Level
			properties["major_version"] = onetimeUpgradeInfo.MajorVersion
			properties["minor_version"] = onetimeUpgradeInfo.MinorVersion
			properties["start_epoch"] = onetimeUpgradeInfo.StartEpoch
			properties["timezone"] = onetimeUpgradeInfo.Timezone
			properties["type"] = onetimeUpgradeInfo.Type
			d = append(d, &properties)
		}
	}
	return
}

func OnetimeUpgradeInfoModel(d map[string]interface{}) *models.OnetimeUpgradeInfo {
	// assume that the incoming map only contains the relevant resource data
	description := d["description"].(string)
	majorVersion := int32(d["major_version"].(int))
	minorVersion := int32(d["minor_version"].(int))
	startEpoch := int64(d["start_epoch"].(int))
	timezone := d["timezone"].(string)
	
	return &models.OnetimeUpgradeInfo {
		Description: description,
		MajorVersion: &majorVersion,
		MinorVersion: &minorVersion,
		StartEpoch: &startEpoch,
		Timezone: timezone,
	}
}

func GetOnetimeUpgradeInfoPropertyFields() (t []string) {
	return []string{
		"description",
		"major_version",
		"minor_version",
		"start_epoch",
		"timezone",
	}
}