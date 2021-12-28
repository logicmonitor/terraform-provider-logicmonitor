package schemata

import (
	"terraform-provider-logicmonitor/logicmonitor/utils"
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func CloudServiceSettingsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"custom_n_s_p_schedule": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"dead_operation": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"device_display_name_template": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"disable_stop_terminate_host_monitor": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"disable_terminated_host_alerting": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"monitoring_region_infos": {
			Type: schema.TypeSet,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Set:      schema.HashString,
			Optional: true,
		},
		
		"monitoring_regions": {
			Type: schema.TypeSet,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Set:      schema.HashString,
			Computed: true,
		},
		
		"name_filter": {
			Type: schema.TypeSet,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Set:      schema.HashString,
			Optional: true,
		},
		
		"normal_collector_config": {
			Type: schema.TypeList, //GoType: CloudNormalCollectorConfig
			Elem: &schema.Resource{
				Schema: CloudNormalCollectorConfigSchema(),
			},
			Optional: true,
		},
		
		"select_all": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"tags": {
			Type: schema.TypeList, //GoType: []*CloudTagFilter 
			Elem: &schema.Resource{
				Schema: CloudTagFilterSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"use_default": {
			Type: schema.TypeBool,
			Required: true,
		},
		
	}
}

func SetCloudServiceSettingsSubResourceData(m []*models.CloudServiceSettings) (d []*map[string]interface{}) {
	for _, cloudServiceSettings := range m {
		if cloudServiceSettings != nil {
			properties := make(map[string]interface{})
			properties["custom_n_s_p_schedule"] = cloudServiceSettings.CustomNSPSchedule
			properties["dead_operation"] = cloudServiceSettings.DeadOperation
			properties["device_display_name_template"] = cloudServiceSettings.DeviceDisplayNameTemplate
			properties["disable_stop_terminate_host_monitor"] = cloudServiceSettings.DisableStopTerminateHostMonitor
			properties["disable_terminated_host_alerting"] = cloudServiceSettings.DisableTerminatedHostAlerting
			properties["monitoring_region_infos"] = cloudServiceSettings.MonitoringRegionInfos
			properties["monitoring_regions"] = cloudServiceSettings.MonitoringRegions
			properties["name_filter"] = cloudServiceSettings.NameFilter
			properties["normal_collector_config"] = SetCloudNormalCollectorConfigSubResourceData([]*models.CloudNormalCollectorConfig{cloudServiceSettings.NormalCollectorConfig})
			properties["select_all"] = cloudServiceSettings.SelectAll
			properties["tags"] = SetCloudTagFilterSubResourceData(cloudServiceSettings.Tags)
			properties["use_default"] = cloudServiceSettings.UseDefault
			d = append(d, &properties)
		}
	}
	return
}

func CloudServiceSettingsModel(d map[string]interface{}) *models.CloudServiceSettings {
	// assume that the incoming map only contains the relevant resource data
	customNSPSchedule := d["custom_n_s_p_schedule"].(string)
	deadOperation := d["dead_operation"].(string)
	deviceDisplayNameTemplate := d["device_display_name_template"].(string)
	disableStopTerminateHostMonitor := d["disable_stop_terminate_host_monitor"].(bool)
	disableTerminatedHostAlerting := d["disable_terminated_host_alerting"].(bool)
	monitoringRegionInfos := utils.ConvertSetToStringSlice(d["monitoring_region_infos"].(*schema.Set))
	nameFilter := utils.ConvertSetToStringSlice(d["name_filter"].(*schema.Set))
	var normalCollectorConfig *models.CloudNormalCollectorConfig = nil
	normalCollectorConfigList := d["normal_collector_config"].([]interface{})
	if len(normalCollectorConfigList) > 0 { // len(nil) = 0
		normalCollectorConfig = CloudNormalCollectorConfigModel(normalCollectorConfigList[0].(map[string]interface{}))
	}
	selectAll := d["select_all"].(bool)
	tags := utils.GetCloudTagFilters(d["tags"].([]interface{}))
	useDefault := d["use_default"].(bool)
	
	return &models.CloudServiceSettings {
		CustomNSPSchedule: customNSPSchedule,
		DeadOperation: deadOperation,
		DeviceDisplayNameTemplate: deviceDisplayNameTemplate,
		DisableStopTerminateHostMonitor: disableStopTerminateHostMonitor,
		DisableTerminatedHostAlerting: disableTerminatedHostAlerting,
		MonitoringRegionInfos: monitoringRegionInfos,
		NameFilter: nameFilter,
		NormalCollectorConfig: normalCollectorConfig,
		SelectAll: selectAll,
		Tags: tags,
		UseDefault: &useDefault,
	}
}

func GetCloudServiceSettingsPropertyFields() (t []string) {
	return []string{
		"custom_n_s_p_schedule",
		"dead_operation",
		"device_display_name_template",
		"disable_stop_terminate_host_monitor",
		"disable_terminated_host_alerting",
		"monitoring_region_infos",
		"name_filter",
		"normal_collector_config",
		"select_all",
		"tags",
		"use_default",
	}
}