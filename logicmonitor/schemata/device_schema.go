package schemata

import (
	"strconv"
	"terraform-provider-logicmonitor/logicmonitor/utils"
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DeviceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"alert_expr": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"auto_balanced_collector_group_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"auto_properties": {
			Type: schema.TypeSet,
			Elem: &schema.Resource{
				Schema: NameAndValueSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Computed: true,
		},
		
		"auto_props_assigned_on": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"auto_props_updated_on": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"aws_state": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"azure_state": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"checkpoints": {
			Type: schema.TypeList, //GoType: []*WebsiteCheckPoint  
			Elem: &schema.Resource{
				Schema: WebsiteCheckPointSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"cloned_from_host_id": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"collector_description": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"collectors": {
			Type: schema.TypeList, //GoType: []*WebsiteDeviceCollectorInfo  
			Elem: &schema.Resource{
				Schema: WebsiteDeviceCollectorInfoSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Computed: true,
		},
		
		"contains_multi_value": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"created_on": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"current_collector_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"current_log_collector_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"custom_properties": {
			Type: schema.TypeSet,
			Elem: &schema.Resource{
				Schema: NameAndValueSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"deleted_time_in_ms": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"description": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"device_type": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"disable_alerting": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"display_name": {
			Type: schema.TypeString,
			Required: true,
		},
		
		"domain": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"enable_netflow": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"gcp_state": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"global_sm_alert_cond": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"group_ids": {
			Type: schema.TypeList, //GoType: []int32
			Elem: &schema.Schema{
                 Type: schema.TypeInt,
            },
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"host": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"host_group_ids": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"host_status": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"id": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"ignore_s_s_l": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"individual_alert_level": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"individual_sm_alert_enable": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"inherited_properties": {
			Type: schema.TypeSet,
			Elem: &schema.Resource{
				Schema: NameAndValueSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Computed: true,
		},
		
		"is_internal": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"is_preferred_log_collector_configured": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"last_data_time": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"last_rawdata_time": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"last_updated": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"link": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"log_collector_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"model": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"name": {
			Type: schema.TypeString,
			Required: true,
		},
		
		"need_stc_grp_and_sorted_c_p": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"netflow_collector_description": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"netflow_collector_group_id": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"netflow_collector_group_name": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"netflow_collector_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"op": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"overall_alert_level": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"page_load_alert_time_in_m_s": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"percent_pkts_not_receive_in_time": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"polling_interval": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"preferred_collector_group_id": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"preferred_collector_group_name": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"preferred_collector_id": {
			Type: schema.TypeInt,
			Required: true,
		},
		
		"properties": {
			Type: schema.TypeSet,
			Elem: &schema.Resource{
				Schema: NameAndValueSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"related_device_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"resource_ids": {
			Type: schema.TypeSet,
			Elem: &schema.Resource{
				Schema: NameAndValueSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"scan_config_id": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"schema": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"status": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"steps": {
			Type: schema.TypeList, //GoType: []*UptimeWebCheckStep  
			Elem: &schema.Resource{
				Schema: UptimeWebCheckStepSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"stop_monitoring": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"stop_monitoring_by_folder": {
			Type: schema.TypeBool,
			Computed: true,
		},
		
		"synthetics_collector_ids": {
			Type: schema.TypeList, //GoType: []int32
			Elem: &schema.Schema{
                 Type: schema.TypeInt,
            },
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"system_properties": {
			Type: schema.TypeSet,
			Elem: &schema.Resource{
				Schema: NameAndValueSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Computed: true,
		},
		
		"template": {
			Type: schema.TypeList, //GoType: JSONObject
			Elem: &schema.Resource{
				Schema: JSONObjectSchema(),
			},
			Optional: true,
		},
		
		"test_location": {
			Type: schema.TypeList, //GoType: WebsiteLocation 
            Elem: &schema.Resource{
				Schema: WebsiteLocationSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"timeout_in_m_s_pkts_not_receive": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"to_delete_time_in_ms": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"transition": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"trigger_s_s_l_expiration_alert": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"trigger_s_s_l_status_alert": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"type": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"up_time_in_seconds": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"updated_on": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"use_default_alert_setting": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"use_default_location_setting": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"user_permission": {
			Type: schema.TypeString,
			Computed: true,
		},
		
	}
}


// Schema mapping representing the resource's respective datasource object defined in Terraform configuration
// Only difference between this and DeviceSchema() are the computabilty of the id field and the inclusion of a filter field for datasources
func DataSourceDeviceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"alert_expr": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"auto_balanced_collector_group_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"auto_properties": {
			Type: schema.TypeList, //GoType: []*NameAndValue 
			Elem: &schema.Resource{
				Schema: NameAndValueSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"auto_props_assigned_on": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"auto_props_updated_on": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"aws_state": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"azure_state": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"checkpoints": {
			Type: schema.TypeList, //GoType: []*WebsiteCheckPoint 
			Elem: &schema.Resource{
				Schema: WebsiteCheckPointSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"cloned_from_host_id": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"collector_description": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"collectors": {
			Type: schema.TypeList, //GoType: []*WebsiteDeviceCollectorInfo 
			Elem: &schema.Resource{
				Schema: WebsiteDeviceCollectorInfoSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"contains_multi_value": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"created_on": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"current_collector_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"current_log_collector_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"custom_properties": {
			Type: schema.TypeList, //GoType: []*NameAndValue 
			Elem: &schema.Resource{
				Schema: NameAndValueSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"deleted_time_in_ms": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"description": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"device_type": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"disable_alerting": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"display_name": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"domain": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"enable_netflow": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"gcp_state": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"global_sm_alert_cond": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"group_ids": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"host": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"host_group_ids": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"host_status": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"id": {
			Type: schema.TypeInt,
			Computed: true,
			Optional: true,
		},
		
		"ignore_s_s_l": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"individual_alert_level": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"individual_sm_alert_enable": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"inherited_properties": {
			Type: schema.TypeList, //GoType: []*NameAndValue 
			Elem: &schema.Resource{
				Schema: NameAndValueSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"is_internal": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"is_preferred_log_collector_configured": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"last_data_time": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"last_rawdata_time": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"last_updated": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"link": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"log_collector_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"model": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"name": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"need_stc_grp_and_sorted_c_p": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"netflow_collector_description": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"netflow_collector_group_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"netflow_collector_group_name": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"netflow_collector_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"op": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"overall_alert_level": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"page_load_alert_time_in_m_s": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"percent_pkts_not_receive_in_time": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"polling_interval": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"preferred_collector_group_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"preferred_collector_group_name": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"preferred_collector_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"properties": {
			Type: schema.TypeList, //GoType: []*NameAndValue 
			Elem: &schema.Resource{
				Schema: NameAndValueSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"related_device_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"resource_ids": {
			Type: schema.TypeList, //GoType: []*NameAndValue 
			Elem: &schema.Resource{
				Schema: NameAndValueSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"scan_config_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"schema": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"status": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"steps": {
			Type: schema.TypeList, //GoType: []*UptimeWebCheckStep 
			Elem: &schema.Resource{
				Schema: UptimeWebCheckStepSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"stop_monitoring": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"stop_monitoring_by_folder": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"synthetics_collector_ids": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"system_properties": {
			Type: schema.TypeList, //GoType: []*NameAndValue 
			Elem: &schema.Resource{
				Schema: NameAndValueSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"template": {
			Type: schema.TypeList, //GoType: JSONObject
			Elem: &schema.Resource{
				Schema: JSONObjectSchema(),
			},
			Optional: true,
		},
		
		"test_location": {
			Type: schema.TypeList, //GoType: WebsiteLocation
			Elem: &schema.Resource{
				Schema: WebsiteLocationSchema(),
			},
			Optional: true,
		},
		
		"timeout_in_m_s_pkts_not_receive": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"to_delete_time_in_ms": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"transition": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"trigger_s_s_l_expiration_alert": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"trigger_s_s_l_status_alert": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"type": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"up_time_in_seconds": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"updated_on": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"use_default_alert_setting": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"use_default_location_setting": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"user_permission": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"filter": {
			Type:     schema.TypeString,
            Optional: true,
		},
	}
}

func SetDeviceResourceData(d *schema.ResourceData, m *models.Device) {
	d.Set("alert_expr", m.AlertExpr)
	d.Set("auto_balanced_collector_group_id", m.AutoBalancedCollectorGroupID)
	d.Set("auto_properties", SetNameAndValueSubResourceData(m.AutoProperties))
	d.Set("auto_props_assigned_on", m.AutoPropsAssignedOn)
	d.Set("auto_props_updated_on", m.AutoPropsUpdatedOn)
	d.Set("aws_state", m.AwsState)
	d.Set("azure_state", m.AzureState)
	d.Set("checkpoints", SetWebsiteCheckPointSubResourceData(m.Checkpoints))
	d.Set("cloned_from_host_id", m.ClonedFromHostID)
	d.Set("collector_description", m.CollectorDescription)
	d.Set("collectors", SetWebsiteDeviceCollectorInfoSubResourceData(m.Collectors))
	d.Set("contains_multi_value", m.ContainsMultiValue)
	d.Set("created_on", m.CreatedOn)
	d.Set("current_collector_id", m.CurrentCollectorID)
	d.Set("current_log_collector_id", m.CurrentLogCollectorID)
	d.Set("custom_properties", SetNameAndValueSubResourceData(m.CustomProperties))
	d.Set("deleted_time_in_ms", m.DeletedTimeInMs)
	d.Set("description", m.Description)
	d.Set("device_type", m.DeviceType)
	d.Set("disable_alerting", m.DisableAlerting)
	d.Set("display_name", m.DisplayName)
	d.Set("domain", m.Domain)
	d.Set("enable_netflow", m.EnableNetflow)
	d.Set("gcp_state", m.GcpState)
	d.Set("global_sm_alert_cond", m.GlobalSmAlertCond)
	d.Set("group_ids", m.GroupIds)
	d.Set("host", m.Host)
	d.Set("host_group_ids", m.HostGroupIds)
	d.Set("host_status", m.HostStatus)
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("ignore_s_s_l", m.IgnoreSSL)
	d.Set("individual_alert_level", m.IndividualAlertLevel)
	d.Set("individual_sm_alert_enable", m.IndividualSmAlertEnable)
	d.Set("inherited_properties", SetNameAndValueSubResourceData(m.InheritedProperties))
	d.Set("is_internal", m.IsInternal)
	d.Set("is_preferred_log_collector_configured", m.IsPreferredLogCollectorConfigured)
	d.Set("last_data_time", m.LastDataTime)
	d.Set("last_rawdata_time", m.LastRawdataTime)
	d.Set("last_updated", m.LastUpdated)
	d.Set("link", m.Link)
	d.Set("log_collector_id", m.LogCollectorID)
	d.Set("model", m.Model)
	d.Set("name", m.Name)
	// Special handling for 'need_stc_grp_and_sorted_c_p' as it is a query param and not returned by API
	if val, ok := d.GetOk("need_stc_grp_and_sorted_c_p"); ok {
		d.Set("need_stc_grp_and_sorted_c_p", val)
	} else {
		d.Set("need_stc_grp_and_sorted_c_p", m.NeedStcGrpAndSortedCP)
	}
	d.Set("netflow_collector_description", m.NetflowCollectorDescription)
	d.Set("netflow_collector_group_id", m.NetflowCollectorGroupID)
	d.Set("netflow_collector_group_name", m.NetflowCollectorGroupName)
	d.Set("netflow_collector_id", m.NetflowCollectorID)
	d.Set("op", m.Op)
	d.Set("overall_alert_level", m.OverallAlertLevel)
	d.Set("page_load_alert_time_in_m_s", m.PageLoadAlertTimeInMS)
	d.Set("percent_pkts_not_receive_in_time", m.PercentPktsNotReceiveInTime)
	d.Set("polling_interval", m.PollingInterval)
	d.Set("preferred_collector_group_id", m.PreferredCollectorGroupID)
	d.Set("preferred_collector_group_name", m.PreferredCollectorGroupName)
	d.Set("preferred_collector_id", m.PreferredCollectorID)
	d.Set("properties", SetNameAndValueSubResourceData(m.Properties))
	d.Set("related_device_id", m.RelatedDeviceID)
	d.Set("resource_ids", SetNameAndValueSubResourceData(m.ResourceIds))
	d.Set("scan_config_id", m.ScanConfigID)
	d.Set("schema", m.Schema)
	d.Set("status", m.Status)
	d.Set("steps", SetUptimeWebCheckStepSubResourceData(m.Steps))
	d.Set("stop_monitoring", m.StopMonitoring)
	d.Set("stop_monitoring_by_folder", m.StopMonitoringByFolder)
	d.Set("synthetics_collector_ids", m.SyntheticsCollectorIds)
	d.Set("system_properties", SetNameAndValueSubResourceData(m.SystemProperties))
	d.Set("template", SetJSONObjectSubResourceData([]*models.JSONObject{m.Template}))
	d.Set("test_location", SetWebsiteLocationSubResourceData([]*models.WebsiteLocation{m.TestLocation}))
	d.Set("timeout_in_m_s_pkts_not_receive", m.TimeoutInMSPktsNotReceive)
	d.Set("to_delete_time_in_ms", m.ToDeleteTimeInMs)
	d.Set("transition", m.Transition)
	d.Set("trigger_s_s_l_expiration_alert", m.TriggerSSLExpirationAlert)
	d.Set("trigger_s_s_l_status_alert", m.TriggerSSLStatusAlert)
	d.Set("type", m.Type)
	d.Set("up_time_in_seconds", m.UpTimeInSeconds)
	d.Set("updated_on", m.UpdatedOn)
	d.Set("use_default_alert_setting", m.UseDefaultAlertSetting)
	d.Set("use_default_location_setting", m.UseDefaultLocationSetting)
	d.Set("user_permission", m.UserPermission)
}

func SetDeviceSubResourceData(m []*models.Device) (d []*map[string]interface{}) {
	for _, device := range m {
		if device != nil {
			properties := make(map[string]interface{})
			properties["alert_expr"] = device.AlertExpr
			properties["auto_balanced_collector_group_id"] = device.AutoBalancedCollectorGroupID
			properties["auto_properties"] = SetNameAndValueSubResourceData(device.AutoProperties)
			properties["auto_props_assigned_on"] = device.AutoPropsAssignedOn
			properties["auto_props_updated_on"] = device.AutoPropsUpdatedOn
			properties["aws_state"] = device.AwsState
			properties["azure_state"] = device.AzureState
			properties["checkpoints"] = SetWebsiteCheckPointSubResourceData(device.Checkpoints)
			properties["cloned_from_host_id"] = device.ClonedFromHostID
			properties["collector_description"] = device.CollectorDescription
			properties["collectors"] = SetWebsiteDeviceCollectorInfoSubResourceData(device.Collectors)
			properties["contains_multi_value"] = device.ContainsMultiValue
			properties["created_on"] = device.CreatedOn
			properties["current_collector_id"] = device.CurrentCollectorID
			properties["current_log_collector_id"] = device.CurrentLogCollectorID
			properties["custom_properties"] = SetNameAndValueSubResourceData(device.CustomProperties)
			properties["deleted_time_in_ms"] = device.DeletedTimeInMs
			properties["description"] = device.Description
			properties["device_type"] = device.DeviceType
			properties["disable_alerting"] = device.DisableAlerting
			properties["display_name"] = device.DisplayName
			properties["domain"] = device.Domain
			properties["enable_netflow"] = device.EnableNetflow
			properties["gcp_state"] = device.GcpState
			properties["global_sm_alert_cond"] = device.GlobalSmAlertCond
			properties["group_ids"] = device.GroupIds
			properties["host"] = device.Host
			properties["host_group_ids"] = device.HostGroupIds
			properties["host_status"] = device.HostStatus
			properties["id"] = device.ID
			properties["ignore_s_s_l"] = device.IgnoreSSL
			properties["individual_alert_level"] = device.IndividualAlertLevel
			properties["individual_sm_alert_enable"] = device.IndividualSmAlertEnable
			properties["inherited_properties"] = SetNameAndValueSubResourceData(device.InheritedProperties)
			properties["is_internal"] = device.IsInternal
			properties["is_preferred_log_collector_configured"] = device.IsPreferredLogCollectorConfigured
			properties["last_data_time"] = device.LastDataTime
			properties["last_rawdata_time"] = device.LastRawdataTime
			properties["last_updated"] = device.LastUpdated
			properties["link"] = device.Link
			properties["log_collector_id"] = device.LogCollectorID
			properties["model"] = device.Model
			properties["name"] = device.Name
			properties["need_stc_grp_and_sorted_c_p"] = device.NeedStcGrpAndSortedCP
			properties["netflow_collector_description"] = device.NetflowCollectorDescription
			properties["netflow_collector_group_id"] = device.NetflowCollectorGroupID
			properties["netflow_collector_group_name"] = device.NetflowCollectorGroupName
			properties["netflow_collector_id"] = device.NetflowCollectorID
			properties["op"] = device.Op
			properties["overall_alert_level"] = device.OverallAlertLevel
			properties["page_load_alert_time_in_m_s"] = device.PageLoadAlertTimeInMS
			properties["percent_pkts_not_receive_in_time"] = device.PercentPktsNotReceiveInTime
			properties["polling_interval"] = device.PollingInterval
			properties["preferred_collector_group_id"] = device.PreferredCollectorGroupID
			properties["preferred_collector_group_name"] = device.PreferredCollectorGroupName
			properties["preferred_collector_id"] = device.PreferredCollectorID
			properties["properties"] = SetNameAndValueSubResourceData(device.Properties)
			properties["related_device_id"] = device.RelatedDeviceID
			properties["resource_ids"] = SetNameAndValueSubResourceData(device.ResourceIds)
			properties["scan_config_id"] = device.ScanConfigID
			properties["schema"] = device.Schema
			properties["status"] = device.Status
			properties["steps"] = SetUptimeWebCheckStepSubResourceData(device.Steps)
			properties["stop_monitoring"] = device.StopMonitoring
			properties["stop_monitoring_by_folder"] = device.StopMonitoringByFolder
			properties["synthetics_collector_ids"] = device.SyntheticsCollectorIds
			properties["system_properties"] = SetNameAndValueSubResourceData(device.SystemProperties)
			properties["template"] = SetJSONObjectSubResourceData([]*models.JSONObject{device.Template})
			properties["test_location"] = SetWebsiteLocationSubResourceData([]*models.WebsiteLocation{device.TestLocation})
			properties["timeout_in_m_s_pkts_not_receive"] = device.TimeoutInMSPktsNotReceive
			properties["to_delete_time_in_ms"] = device.ToDeleteTimeInMs
			properties["transition"] = device.Transition
			properties["trigger_s_s_l_expiration_alert"] = device.TriggerSSLExpirationAlert
			properties["trigger_s_s_l_status_alert"] = device.TriggerSSLStatusAlert
			properties["type"] = device.Type
			properties["up_time_in_seconds"] = device.UpTimeInSeconds
			properties["updated_on"] = device.UpdatedOn
			properties["use_default_alert_setting"] = device.UseDefaultAlertSetting
			properties["use_default_location_setting"] = device.UseDefaultLocationSetting
			properties["user_permission"] = device.UserPermission
			d = append(d, &properties)
		}
	}
	return
}

func DeviceModel(d *schema.ResourceData) *models.Device {
	alertExpr := d.Get("alert_expr").(string)
	autoBalancedCollectorGroupID := int32(d.Get("auto_balanced_collector_group_id").(int))
	checkpointsTempSlice := d.Get("checkpoints").([]interface{})
	checkpoints := []*models.WebsiteCheckPoint{}
	for _, val := range checkpointsTempSlice {
		checkpoints = append(checkpoints, val.(*models.WebsiteCheckPoint))
	}
	containsMultiValue := d.Get("contains_multi_value").(bool)
	currentCollectorID := int32(d.Get("current_collector_id").(int))
	currentLogCollectorID := int32(d.Get("current_log_collector_id").(int))
	customProperties := utils.GetPropertiesFromResource(d, "custom_properties")
	description := d.Get("description").(string)
	deviceType := int32(d.Get("device_type").(int))
	disableAlerting := d.Get("disable_alerting").(bool)
	displayName := d.Get("display_name").(string)
	domain := d.Get("domain").(string)
	enableNetflow := d.Get("enable_netflow").(bool)
	globalSmAlertCond := int32(d.Get("global_sm_alert_cond").(int))
	groupIds := utils.ConvertSetToInt32Slice(d.Get("group_ids").([]interface{}))
	host := d.Get("host").(string)
	hostGroupIds := d.Get("host_group_ids").(string)
	id, _ := strconv.Atoi(d.Get("id").(string))
	ignoreSSL := d.Get("ignore_s_s_l").(bool)
	individualAlertLevel := d.Get("individual_alert_level").(string)
	individualSmAlertEnable := d.Get("individual_sm_alert_enable").(bool)
	isInternal := d.Get("is_internal").(bool)
	isPreferredLogCollectorConfigured := d.Get("is_preferred_log_collector_configured").(bool)
	link := d.Get("link").(string)
	logCollectorID := int32(d.Get("log_collector_id").(int))
	model := d.Get("model").(string)
	name := d.Get("name").(string)
	needStcGrpAndSortedCP := d.Get("need_stc_grp_and_sorted_c_p").(bool)
	netflowCollectorID := int32(d.Get("netflow_collector_id").(int))
	op := d.Get("op").(string)
	overallAlertLevel := d.Get("overall_alert_level").(string)
	pageLoadAlertTimeInMS := int32(d.Get("page_load_alert_time_in_m_s").(int))
	percentPktsNotReceiveInTime := int32(d.Get("percent_pkts_not_receive_in_time").(int))
	pollingInterval := int32(d.Get("polling_interval").(int))
	preferredCollectorID := int32(d.Get("preferred_collector_id").(int))
	properties := utils.GetPropertiesFromResource(d, "properties")
	relatedDeviceID := int32(d.Get("related_device_id").(int))
	resourceIds := utils.GetPropertiesFromResource(d, "resource_ids")
	schema := d.Get("schema").(string)
	steps := utils.GetPropFromUptimeSteps(d, "steps")
	stopMonitoring := d.Get("stop_monitoring").(bool)
	syntheticsCollectorIds := utils.ConvertSetToInt32Slice(d.Get("synthetics_collector_ids").([]interface{}))
	var template *models.JSONObject = nil
	templateInterface, templateIsSet := d.GetOk("template")
	if templateIsSet {
		templateMap := templateInterface.([]interface{})[0].(map[string]interface{})
		template = JSONObjectModel(templateMap)
	}
    testLocation := utils.GetPropFromLocationMap(d, "test_location")

	          
	timeoutInMSPktsNotReceive := int32(d.Get("timeout_in_m_s_pkts_not_receive").(int))
	transition := int32(d.Get("transition").(int))
	triggerSSLExpirationAlert := d.Get("trigger_s_s_l_expiration_alert").(bool)
	triggerSSLStatusAlert := d.Get("trigger_s_s_l_status_alert").(bool)
	typeVar := d.Get("type").(string)
	useDefaultAlertSetting := d.Get("use_default_alert_setting").(bool)
	useDefaultLocationSetting := d.Get("use_default_location_setting").(bool)
	
	return &models.Device {
		AlertExpr: alertExpr,
		AutoBalancedCollectorGroupID: autoBalancedCollectorGroupID,
		Checkpoints: checkpoints,
		ContainsMultiValue: containsMultiValue,
		CurrentCollectorID: currentCollectorID,
		CurrentLogCollectorID: currentLogCollectorID,
		CustomProperties: customProperties,
		Description: description,
		DeviceType: deviceType,
		DisableAlerting: disableAlerting,
		DisplayName: &displayName,
		Domain: domain,
		EnableNetflow: enableNetflow,
		GlobalSmAlertCond: globalSmAlertCond,
		GroupIds: groupIds,
		Host: host,
		HostGroupIds: hostGroupIds,
		ID: int32(id),
		IgnoreSSL: ignoreSSL,
		IndividualAlertLevel: individualAlertLevel,
		IndividualSmAlertEnable: individualSmAlertEnable,
		IsInternal: isInternal,
		IsPreferredLogCollectorConfigured: isPreferredLogCollectorConfigured,
		Link: link,
		LogCollectorID: logCollectorID,
		Model: model,
		Name: &name,
		NeedStcGrpAndSortedCP: &needStcGrpAndSortedCP,
		NetflowCollectorID: netflowCollectorID,
		Op: op,
		OverallAlertLevel: overallAlertLevel,
		PageLoadAlertTimeInMS: pageLoadAlertTimeInMS,
		PercentPktsNotReceiveInTime: &percentPktsNotReceiveInTime,
		PollingInterval: pollingInterval,
		PreferredCollectorID: &preferredCollectorID,
		Properties: properties,
		RelatedDeviceID: relatedDeviceID,
		ResourceIds: resourceIds,
		Schema: schema,
		Steps: steps,
		StopMonitoring: stopMonitoring,
		SyntheticsCollectorIds: syntheticsCollectorIds,
		Template: template,
		TestLocation: testLocation,
		TimeoutInMSPktsNotReceive: &timeoutInMSPktsNotReceive,
		Transition: transition,
		TriggerSSLExpirationAlert: triggerSSLExpirationAlert,
		TriggerSSLStatusAlert: triggerSSLStatusAlert,
		Type: typeVar,
		UseDefaultAlertSetting: useDefaultAlertSetting,
		UseDefaultLocationSetting: useDefaultLocationSetting,
	}
}
func GetDevicePropertyFields() (t []string) {
	return []string{
		"alert_expr",
		"auto_balanced_collector_group_id",
		"checkpoints",
		"contains_multi_value",
		"current_collector_id",
		"current_log_collector_id",
		"custom_properties",
		"description",
		"device_type",
		"disable_alerting",
		"display_name",
		"domain",
		"enable_netflow",
		"global_sm_alert_cond",
		"group_ids",
		"host",
		"host_group_ids",
		"id",
		"ignore_s_s_l",
		"individual_alert_level",
		"individual_sm_alert_enable",
		"is_internal",
		"is_preferred_log_collector_configured",
		"link",
		"log_collector_id",
		"model",
		"name",
		"need_stc_grp_and_sorted_c_p",
		"netflow_collector_id",
		"op",
		"overall_alert_level",
		"page_load_alert_time_in_m_s",
		"percent_pkts_not_receive_in_time",
		"polling_interval",
		"preferred_collector_id",
		"properties",
		"related_device_id",
		"resource_ids",
		"schema",
		"steps",
		"stop_monitoring",
		"synthetics_collector_ids",
		"template",
		"test_location",
		"timeout_in_m_s_pkts_not_receive",
		"transition",
		"trigger_s_s_l_expiration_alert",
		"trigger_s_s_l_status_alert",
		"type",
		"use_default_alert_setting",
		"use_default_location_setting",
	}
}