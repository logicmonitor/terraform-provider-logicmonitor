package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-logicmonitor/models"
)

func DeviceGroupDataSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"applies_to": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"aws_regions_info": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"azure_regions_info": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"description": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"full_path": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"gcp_regions_info": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"group_type": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"id": {
			Type:     schema.TypeInt,
			Computed: true,
		},

		"name": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"num_of_direct_devices": {
			Type:     schema.TypeInt,
			Computed: true,
		},

		"num_of_direct_sub_groups": {
			Type:     schema.TypeInt,
			Computed: true,
		},

		"num_of_hosts": {
			Type:     schema.TypeInt,
			Computed: true,
		},

		"user_permission": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func SetDeviceGroupDataSubResourceData(m []*models.DeviceGroupData) (d []*map[string]interface{}) {
	for _, deviceGroupData := range m {
		if deviceGroupData != nil {
			properties := make(map[string]interface{})
			properties["applies_to"] = deviceGroupData.AppliesTo
			properties["aws_regions_info"] = deviceGroupData.AwsRegionsInfo
			properties["azure_regions_info"] = deviceGroupData.AzureRegionsInfo
			properties["description"] = deviceGroupData.Description
			properties["full_path"] = deviceGroupData.FullPath
			properties["gcp_regions_info"] = deviceGroupData.GcpRegionsInfo
			properties["group_type"] = deviceGroupData.GroupType
			properties["id"] = deviceGroupData.ID
			properties["name"] = deviceGroupData.Name
			properties["num_of_direct_devices"] = deviceGroupData.NumOfDirectDevices
			properties["num_of_direct_sub_groups"] = deviceGroupData.NumOfDirectSubGroups
			properties["num_of_hosts"] = deviceGroupData.NumOfHosts
			properties["user_permission"] = deviceGroupData.UserPermission
			d = append(d, &properties)
		}
	}
	return
}

func DeviceGroupDataModel(d map[string]interface{}) *models.DeviceGroupData {
	// assume that the incoming map only contains the relevant resource data

	return &models.DeviceGroupData{}
}

func GetDeviceGroupDataPropertyFields() (t []string) {
	return []string{
		"id",
	}
}
