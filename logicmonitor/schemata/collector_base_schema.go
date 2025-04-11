package schemata

import (
	"terraform-provider-logicmonitor/logicmonitor/utils"
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func CollectorBaseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"ack_comment": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"acked": {
			Type: schema.TypeBool,
			Computed: true,
		},
		
		"acked_by": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"acked_on": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"acked_on_local": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"arch": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"automatic_upgrade_info": {
			Type: schema.TypeList, //GoType: AutomaticUpgradeInfo
			Elem: &schema.Resource{
				Schema: AutomaticUpgradeInfoSchema(),
			},
			Optional: true,
		},
		
		"backup_agent_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"build": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"can_downgrade": {
			Type: schema.TypeBool,
			Computed: true,
		},
		
		"can_downgrade_reason": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"clear_sent": {
			Type: schema.TypeBool,
			Computed: true,
		},
		
		"collector_conf": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"collector_device_id": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"collector_group_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"collector_group_name": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"collector_size": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"conf_version": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"created_on": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"created_on_local": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"custom_properties": {
			Type: schema.TypeSet,
			Elem: &schema.Resource{
				Schema: NameAndValueSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"description": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"ea": {
			Type: schema.TypeBool,
			Computed: true,
		},
		
		"enable_fail_back": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"enable_fail_over_on_collector_device": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"escalating_chain_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"has_fail_over_device": {
			Type: schema.TypeBool,
			Computed: true,
		},
		
		"hostname": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"id": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"in_s_d_t": {
			Type: schema.TypeBool,
			Computed: true,
		},
		
		"is_down": {
			Type: schema.TypeBool,
			Computed: true,
		},
		
		"last_sent_notification_on": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"last_sent_notification_on_local": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"need_auto_create_collector_device": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"netscan_version": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"next_recipient": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"next_upgrade_info": {
			Type: schema.TypeList, //GoType: NextUpgradeInfo
			Elem: &schema.Resource{
				Schema: NextUpgradeInfoSchema(),
			},
			Computed: true,
		},
		
		"number_of_hosts": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"number_of_instances": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"onetime_downgrade_info": {
			Type: schema.TypeList, //GoType: OnetimeUpgradeInfo
			Elem: &schema.Resource{
				Schema: OnetimeUpgradeInfoSchema(),
			},
			Optional: true,
		},
		
		"onetime_upgrade_info": {
			Type: schema.TypeList, //GoType: OnetimeUpgradeInfo
			Elem: &schema.Resource{
				Schema: OnetimeUpgradeInfoSchema(),
			},
			Optional: true,
		},
		
		"platform": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"predefined_config": {
			Type: schema.TypeMap, //GoType: interface{}
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Computed: true,
		},
		
		"previous_version": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"resend_ival": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"sbproxy_conf": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"specified_collector_device_group_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"status": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"suppress_alert_clear": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"up_time": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"updated_on": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"updated_on_local": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"user_change_on": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"user_change_on_local": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"user_permission": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"user_visible_hosts_num": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"watchdog_conf": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"watchdog_updated_on": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"watchdog_updated_on_local": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"wrapper_conf": {
			Type: schema.TypeString,
			Computed: true,
		},
		
	}
}

func SetCollectorBaseSubResourceData(m []*models.CollectorBase) (d []*map[string]interface{}) {
	for _, collectorBase := range m {
		if collectorBase != nil {
			properties := make(map[string]interface{})
			properties["ack_comment"] = collectorBase.AckComment
			properties["acked"] = collectorBase.Acked
			properties["acked_by"] = collectorBase.AckedBy
			properties["acked_on"] = collectorBase.AckedOn
			properties["acked_on_local"] = collectorBase.AckedOnLocal
			properties["arch"] = collectorBase.Arch
			properties["automatic_upgrade_info"] = SetAutomaticUpgradeInfoSubResourceData([]*models.AutomaticUpgradeInfo{collectorBase.AutomaticUpgradeInfo})
			properties["backup_agent_id"] = collectorBase.BackupAgentID
			properties["build"] = collectorBase.Build
			properties["can_downgrade"] = collectorBase.CanDowngrade
			properties["can_downgrade_reason"] = collectorBase.CanDowngradeReason
			properties["clear_sent"] = collectorBase.ClearSent
			properties["collector_conf"] = collectorBase.CollectorConf
			properties["collector_device_id"] = collectorBase.CollectorDeviceID
			properties["collector_group_id"] = collectorBase.CollectorGroupID
			properties["collector_group_name"] = collectorBase.CollectorGroupName
			properties["collector_size"] = collectorBase.CollectorSize
			properties["conf_version"] = collectorBase.ConfVersion
			properties["created_on"] = collectorBase.CreatedOn
			properties["created_on_local"] = collectorBase.CreatedOnLocal
			properties["custom_properties"] = SetNameAndValueSubResourceData(collectorBase.CustomProperties)
			properties["description"] = collectorBase.Description
			properties["ea"] = collectorBase.Ea
			properties["enable_fail_back"] = collectorBase.EnableFailBack
			properties["enable_fail_over_on_collector_device"] = collectorBase.EnableFailOverOnCollectorDevice
			properties["escalating_chain_id"] = collectorBase.EscalatingChainID
			properties["has_fail_over_device"] = collectorBase.HasFailOverDevice
			properties["hostname"] = collectorBase.Hostname
			properties["id"] = collectorBase.ID
			properties["in_s_d_t"] = collectorBase.InSDT
			properties["is_down"] = collectorBase.IsDown
			properties["last_sent_notification_on"] = collectorBase.LastSentNotificationOn
			properties["last_sent_notification_on_local"] = collectorBase.LastSentNotificationOnLocal
			properties["need_auto_create_collector_device"] = collectorBase.NeedAutoCreateCollectorDevice
			properties["netscan_version"] = collectorBase.NetscanVersion
			properties["next_recipient"] = collectorBase.NextRecipient
			properties["next_upgrade_info"] = SetNextUpgradeInfoSubResourceData([]*models.NextUpgradeInfo{collectorBase.NextUpgradeInfo})
			properties["number_of_hosts"] = collectorBase.NumberOfHosts
			properties["number_of_instances"] = collectorBase.NumberOfInstances
			properties["onetime_downgrade_info"] = SetOnetimeUpgradeInfoSubResourceData([]*models.OnetimeUpgradeInfo{collectorBase.OnetimeDowngradeInfo})
			properties["onetime_upgrade_info"] = SetOnetimeUpgradeInfoSubResourceData([]*models.OnetimeUpgradeInfo{collectorBase.OnetimeUpgradeInfo})
			properties["platform"] = collectorBase.Platform
			properties["predefined_config"] = collectorBase.PredefinedConfig
			properties["previous_version"] = collectorBase.PreviousVersion
			properties["resend_ival"] = collectorBase.ResendIval
			properties["sbproxy_conf"] = collectorBase.SbproxyConf
			properties["specified_collector_device_group_id"] = collectorBase.SpecifiedCollectorDeviceGroupID
			properties["status"] = collectorBase.Status
			properties["suppress_alert_clear"] = collectorBase.SuppressAlertClear
			properties["up_time"] = collectorBase.UpTime
			properties["updated_on"] = collectorBase.UpdatedOn
			properties["updated_on_local"] = collectorBase.UpdatedOnLocal
			properties["user_change_on"] = collectorBase.UserChangeOn
			properties["user_change_on_local"] = collectorBase.UserChangeOnLocal
			properties["user_permission"] = collectorBase.UserPermission
			properties["user_visible_hosts_num"] = collectorBase.UserVisibleHostsNum
			properties["watchdog_conf"] = collectorBase.WatchdogConf
			properties["watchdog_updated_on"] = collectorBase.WatchdogUpdatedOn
			properties["watchdog_updated_on_local"] = collectorBase.WatchdogUpdatedOnLocal
			properties["wrapper_conf"] = collectorBase.WrapperConf
			d = append(d, &properties)
		}
	}
	return
}

func CollectorBaseModel(d map[string]interface{}) *models.CollectorBase {
	// assume that the incoming map only contains the relevant resource data
	var automaticUpgradeInfo *models.AutomaticUpgradeInfo = nil
	automaticUpgradeInfoList := d["automatic_upgrade_info"].([]interface{})
	if len(automaticUpgradeInfoList) > 0 { // len(nil) = 0
		automaticUpgradeInfo = AutomaticUpgradeInfoModel(automaticUpgradeInfoList[0].(map[string]interface{}))
	}
	backupAgentID := int32(d["backup_agent_id"].(int))
	collectorGroupID := int32(d["collector_group_id"].(int))
	customProperties := utils.GetPropertiesFromMap(d, "custom_properties")
	description := d["description"].(string)
	enableFailBack := d["enable_fail_back"].(bool)
	enableFailOverOnCollectorDevice := d["enable_fail_over_on_collector_device"].(bool)
	escalatingChainID := int32(d["escalating_chain_id"].(int))
	id := int32(d["id"].(int))
	needAutoCreateCollectorDevice := d["need_auto_create_collector_device"].(bool)
	numberOfInstances := int32(d["number_of_instances"].(int))
	var onetimeDowngradeInfo *models.OnetimeUpgradeInfo = nil
	onetimeDowngradeInfoList := d["onetime_downgrade_info"].([]interface{})
	if len(onetimeDowngradeInfoList) > 0 { // len(nil) = 0
		onetimeDowngradeInfo = OnetimeUpgradeInfoModel(onetimeDowngradeInfoList[0].(map[string]interface{}))
	}
	var onetimeUpgradeInfo *models.OnetimeUpgradeInfo = nil
	onetimeUpgradeInfoList := d["onetime_upgrade_info"].([]interface{})
	if len(onetimeUpgradeInfoList) > 0 { // len(nil) = 0
		onetimeUpgradeInfo = OnetimeUpgradeInfoModel(onetimeUpgradeInfoList[0].(map[string]interface{}))
	}
	resendIval := int32(d["resend_ival"].(int))
	specifiedCollectorDeviceGroupID := int32(d["specified_collector_device_group_id"].(int))
	suppressAlertClear := d["suppress_alert_clear"].(bool)
	
	return &models.CollectorBase {
		AutomaticUpgradeInfo: automaticUpgradeInfo,
		BackupAgentID: backupAgentID,
		CollectorGroupID: collectorGroupID,
		CustomProperties: customProperties,
		Description: description,
		EnableFailBack: enableFailBack,
		EnableFailOverOnCollectorDevice: enableFailOverOnCollectorDevice,
		EscalatingChainID: escalatingChainID,
		ID: id,
		NeedAutoCreateCollectorDevice: needAutoCreateCollectorDevice,
		NumberOfInstances: numberOfInstances,
		OnetimeDowngradeInfo: onetimeDowngradeInfo,
		OnetimeUpgradeInfo: onetimeUpgradeInfo,
		ResendIval: resendIval,
		SpecifiedCollectorDeviceGroupID: specifiedCollectorDeviceGroupID,
		SuppressAlertClear: suppressAlertClear,
	}
}

func GetCollectorBasePropertyFields() (t []string) {
	return []string{
		"automatic_upgrade_info",
		"backup_agent_id",
		"collector_group_id",
		"custom_properties",
		"description",
		"enable_fail_back",
		"enable_fail_over_on_collector_device",
		"escalating_chain_id",
		"id",
		"need_auto_create_collector_device",
		"number_of_instances",
		"onetime_downgrade_info",
		"onetime_upgrade_info",
		"resend_ival",
		"specified_collector_device_group_id",
		"suppress_alert_clear",
	}
}