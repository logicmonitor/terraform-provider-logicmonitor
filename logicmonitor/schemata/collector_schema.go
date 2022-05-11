package schemata

import (
	"strconv"
	"terraform-provider-logicmonitor/logicmonitor/utils"
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func CollectorSchema() map[string]*schema.Schema {
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
			Default: "linux64",
			Optional: true,
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
			Optional: true,
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
			Default: "medium",
			Optional: true,
		},
		
		"company": {
			Type: schema.TypeString,
			Optional: true,
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
			Type: schema.TypeList, //GoType: []*NameAndValue 
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
			Optional: true,
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
			Type: schema.TypeString,
			Computed: true,
		},
		
		"in_s_d_t": {
			Type: schema.TypeBool,
			Computed: true,
		},
		
		"installer_url_cmds": {
			Type: schema.TypeMap, //GoType: interface{}
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
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
		
		"monitor_others": {
			Type: schema.TypeBool,
			Default: true,
			Optional: true,
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
		
		"number_of_websites": {
			Type: schema.TypeInt,
			Computed: true,
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
		
		"user_visible_websites_num": {
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
		
		"website_conf": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"wrapper_conf": {
			Type: schema.TypeString,
			Computed: true,
		},
		
	}
}


// Schema mapping representing the resource's respective datasource object defined in Terraform configuration
// Only difference between this and CollectorSchema() are the computabilty of the id field and the inclusion of a filter field for datasources
func DataSourceCollectorSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"ack_comment": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"acked": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"acked_by": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"acked_on": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"acked_on_local": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"arch": {
			Type: schema.TypeString,
			Default: "linux64",
			Optional: true,
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
			Optional: true,
		},
		
		"can_downgrade": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"can_downgrade_reason": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"clear_sent": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"collector_conf": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"collector_device_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"collector_group_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"collector_group_name": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"collector_size": {
			Type: schema.TypeString,
			Default: "medium",
			Optional: true,
		},
		
		"company": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"conf_version": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"created_on": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"created_on_local": {
			Type: schema.TypeString,
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
		
		"description": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"ea": {
			Type: schema.TypeBool,
			Optional: true,
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
			Optional: true,
		},
		
		"hostname": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"id": {
			Type: schema.TypeInt,
			Computed: true,
			Optional: true,
		},
		
		"in_s_d_t": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"installer_url_cmds": {
			Type: schema.TypeMap, //GoType: interface{}
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},
		
		"is_down": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"last_sent_notification_on": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"last_sent_notification_on_local": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"monitor_others": {
			Type: schema.TypeBool,
			Default: true,
			Optional: true,
		},
		
		"need_auto_create_collector_device": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"netscan_version": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"next_recipient": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"next_upgrade_info": {
			Type: schema.TypeList, //GoType: NextUpgradeInfo
			Elem: &schema.Resource{
				Schema: NextUpgradeInfoSchema(),
			},
			Optional: true,
		},
		
		"number_of_hosts": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"number_of_instances": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"number_of_websites": {
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
			Optional: true,
		},
		
		"predefined_config": {
			Type: schema.TypeMap, //GoType: interface{}
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},
		
		"previous_version": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"resend_ival": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"sbproxy_conf": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"specified_collector_device_group_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"status": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"suppress_alert_clear": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"up_time": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"updated_on": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"updated_on_local": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"user_change_on": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"user_change_on_local": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"user_permission": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"user_visible_hosts_num": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"user_visible_websites_num": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"watchdog_conf": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"watchdog_updated_on": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"watchdog_updated_on_local": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"website_conf": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"wrapper_conf": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"filter": {
			Type:     schema.TypeString,
            Optional: true,
		},
	}
}

func SetCollectorResourceData(d *schema.ResourceData, m *models.Collector) {
	d.Set("ack_comment", m.AckComment)
	d.Set("acked", m.Acked)
	d.Set("acked_by", m.AckedBy)
	d.Set("acked_on", m.AckedOn)
	d.Set("acked_on_local", m.AckedOnLocal)
	d.Set("arch", m.Arch)
	d.Set("automatic_upgrade_info", SetAutomaticUpgradeInfoSubResourceData([]*models.AutomaticUpgradeInfo{m.AutomaticUpgradeInfo}))
	d.Set("backup_agent_id", m.BackupAgentID)
	d.Set("build", m.Build)
	d.Set("can_downgrade", m.CanDowngrade)
	d.Set("can_downgrade_reason", m.CanDowngradeReason)
	d.Set("clear_sent", m.ClearSent)
	d.Set("collector_conf", m.CollectorConf)
	d.Set("collector_device_id", m.CollectorDeviceID)
	d.Set("collector_group_id", m.CollectorGroupID)
	d.Set("collector_group_name", m.CollectorGroupName)
	d.Set("collector_size", m.CollectorSize)
	d.Set("company", m.Company)
	d.Set("conf_version", m.ConfVersion)
	d.Set("created_on", m.CreatedOn)
	d.Set("created_on_local", m.CreatedOnLocal)
	d.Set("custom_properties", SetNameAndValueSubResourceData(m.CustomProperties))
	d.Set("description", m.Description)
	d.Set("ea", m.Ea)
	d.Set("enable_fail_back", m.EnableFailBack)
	d.Set("enable_fail_over_on_collector_device", m.EnableFailOverOnCollectorDevice)
	d.Set("escalating_chain_id", m.EscalatingChainID)
	d.Set("has_fail_over_device", m.HasFailOverDevice)
	d.Set("hostname", m.Hostname)
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("in_s_d_t", m.InSDT)
	d.Set("installer_url_cmds", m.InstallerURLCmds)
	d.Set("is_down", m.IsDown)
	d.Set("last_sent_notification_on", m.LastSentNotificationOn)
	d.Set("last_sent_notification_on_local", m.LastSentNotificationOnLocal)
	d.Set("monitor_others", m.MonitorOthers)
	d.Set("need_auto_create_collector_device", m.NeedAutoCreateCollectorDevice)
	d.Set("netscan_version", m.NetscanVersion)
	d.Set("next_recipient", m.NextRecipient)
	d.Set("next_upgrade_info", SetNextUpgradeInfoSubResourceData([]*models.NextUpgradeInfo{m.NextUpgradeInfo}))
	d.Set("number_of_hosts", m.NumberOfHosts)
	d.Set("number_of_instances", m.NumberOfInstances)
	d.Set("number_of_websites", m.NumberOfWebsites)
	d.Set("onetime_downgrade_info", SetOnetimeUpgradeInfoSubResourceData([]*models.OnetimeUpgradeInfo{m.OnetimeDowngradeInfo}))
	d.Set("onetime_upgrade_info", SetOnetimeUpgradeInfoSubResourceData([]*models.OnetimeUpgradeInfo{m.OnetimeUpgradeInfo}))
	d.Set("platform", m.Platform)
	d.Set("predefined_config", m.PredefinedConfig)
	d.Set("previous_version", m.PreviousVersion)
	d.Set("resend_ival", m.ResendIval)
	d.Set("sbproxy_conf", m.SbproxyConf)
	d.Set("specified_collector_device_group_id", m.SpecifiedCollectorDeviceGroupID)
	d.Set("status", m.Status)
	d.Set("suppress_alert_clear", m.SuppressAlertClear)
	d.Set("up_time", m.UpTime)
	d.Set("updated_on", m.UpdatedOn)
	d.Set("updated_on_local", m.UpdatedOnLocal)
	d.Set("user_change_on", m.UserChangeOn)
	d.Set("user_change_on_local", m.UserChangeOnLocal)
	d.Set("user_permission", m.UserPermission)
	d.Set("user_visible_hosts_num", m.UserVisibleHostsNum)
	d.Set("user_visible_websites_num", m.UserVisibleWebsitesNum)
	d.Set("watchdog_conf", m.WatchdogConf)
	d.Set("watchdog_updated_on", m.WatchdogUpdatedOn)
	d.Set("watchdog_updated_on_local", m.WatchdogUpdatedOnLocal)
	d.Set("website_conf", m.WebsiteConf)
	d.Set("wrapper_conf", m.WrapperConf)
}

func SetCollectorSubResourceData(m []*models.Collector) (d []*map[string]interface{}) {
	for _, collector := range m {
		if collector != nil {
			properties := make(map[string]interface{})
			properties["ack_comment"] = collector.AckComment
			properties["acked"] = collector.Acked
			properties["acked_by"] = collector.AckedBy
			properties["acked_on"] = collector.AckedOn
			properties["acked_on_local"] = collector.AckedOnLocal
			properties["arch"] = collector.Arch
			properties["automatic_upgrade_info"] = SetAutomaticUpgradeInfoSubResourceData([]*models.AutomaticUpgradeInfo{collector.AutomaticUpgradeInfo})
			properties["backup_agent_id"] = collector.BackupAgentID
			properties["build"] = collector.Build
			properties["can_downgrade"] = collector.CanDowngrade
			properties["can_downgrade_reason"] = collector.CanDowngradeReason
			properties["clear_sent"] = collector.ClearSent
			properties["collector_conf"] = collector.CollectorConf
			properties["collector_device_id"] = collector.CollectorDeviceID
			properties["collector_group_id"] = collector.CollectorGroupID
			properties["collector_group_name"] = collector.CollectorGroupName
			properties["collector_size"] = collector.CollectorSize
			properties["company"] = collector.Company
			properties["conf_version"] = collector.ConfVersion
			properties["created_on"] = collector.CreatedOn
			properties["created_on_local"] = collector.CreatedOnLocal
			properties["custom_properties"] = SetNameAndValueSubResourceData(collector.CustomProperties)
			properties["description"] = collector.Description
			properties["ea"] = collector.Ea
			properties["enable_fail_back"] = collector.EnableFailBack
			properties["enable_fail_over_on_collector_device"] = collector.EnableFailOverOnCollectorDevice
			properties["escalating_chain_id"] = collector.EscalatingChainID
			properties["has_fail_over_device"] = collector.HasFailOverDevice
			properties["hostname"] = collector.Hostname
			properties["id"] = collector.ID
			properties["in_s_d_t"] = collector.InSDT
			properties["installer_url_cmds"] = collector.InstallerURLCmds
			properties["is_down"] = collector.IsDown
			properties["last_sent_notification_on"] = collector.LastSentNotificationOn
			properties["last_sent_notification_on_local"] = collector.LastSentNotificationOnLocal
			properties["monitor_others"] = collector.MonitorOthers
			properties["need_auto_create_collector_device"] = collector.NeedAutoCreateCollectorDevice
			properties["netscan_version"] = collector.NetscanVersion
			properties["next_recipient"] = collector.NextRecipient
			properties["next_upgrade_info"] = SetNextUpgradeInfoSubResourceData([]*models.NextUpgradeInfo{collector.NextUpgradeInfo})
			properties["number_of_hosts"] = collector.NumberOfHosts
			properties["number_of_instances"] = collector.NumberOfInstances
			properties["number_of_websites"] = collector.NumberOfWebsites
			properties["onetime_downgrade_info"] = SetOnetimeUpgradeInfoSubResourceData([]*models.OnetimeUpgradeInfo{collector.OnetimeDowngradeInfo})
			properties["onetime_upgrade_info"] = SetOnetimeUpgradeInfoSubResourceData([]*models.OnetimeUpgradeInfo{collector.OnetimeUpgradeInfo})
			properties["platform"] = collector.Platform
			properties["predefined_config"] = collector.PredefinedConfig
			properties["previous_version"] = collector.PreviousVersion
			properties["resend_ival"] = collector.ResendIval
			properties["sbproxy_conf"] = collector.SbproxyConf
			properties["specified_collector_device_group_id"] = collector.SpecifiedCollectorDeviceGroupID
			properties["status"] = collector.Status
			properties["suppress_alert_clear"] = collector.SuppressAlertClear
			properties["up_time"] = collector.UpTime
			properties["updated_on"] = collector.UpdatedOn
			properties["updated_on_local"] = collector.UpdatedOnLocal
			properties["user_change_on"] = collector.UserChangeOn
			properties["user_change_on_local"] = collector.UserChangeOnLocal
			properties["user_permission"] = collector.UserPermission
			properties["user_visible_hosts_num"] = collector.UserVisibleHostsNum
			properties["user_visible_websites_num"] = collector.UserVisibleWebsitesNum
			properties["watchdog_conf"] = collector.WatchdogConf
			properties["watchdog_updated_on"] = collector.WatchdogUpdatedOn
			properties["watchdog_updated_on_local"] = collector.WatchdogUpdatedOnLocal
			properties["website_conf"] = collector.WebsiteConf
			properties["wrapper_conf"] = collector.WrapperConf
			d = append(d, &properties)
		}
	}
	return
}

func CollectorModel(d *schema.ResourceData) *models.Collector {
	arch := d.Get("arch").(string)
	var automaticUpgradeInfo *models.AutomaticUpgradeInfo = nil
	automaticUpgradeInfoInterface, automaticUpgradeInfoIsSet := d.GetOk("automatic_upgrade_info")
	if automaticUpgradeInfoIsSet {
		automaticUpgradeInfoMap := automaticUpgradeInfoInterface.([]interface{})[0].(map[string]interface{})
		automaticUpgradeInfo = AutomaticUpgradeInfoModel(automaticUpgradeInfoMap)
	}
	backupAgentID := int32(d.Get("backup_agent_id").(int))
	build := d.Get("build").(string)
	collectorGroupID := int32(d.Get("collector_group_id").(int))
	collectorSize := d.Get("collector_size").(string)
	company := d.Get("company").(string)
	customProperties := utils.GetPropertiesFromResource(d, "custom_properties")
	description := d.Get("description").(string)
	ea := d.Get("ea").(bool)
	enableFailBack := d.Get("enable_fail_back").(bool)
	enableFailOverOnCollectorDevice := d.Get("enable_fail_over_on_collector_device").(bool)
	escalatingChainID := int32(d.Get("escalating_chain_id").(int))
	id, _ := strconv.Atoi(d.Get("id").(string))
	monitorOthers := d.Get("monitor_others").(bool)
	needAutoCreateCollectorDevice := d.Get("need_auto_create_collector_device").(bool)
	numberOfInstances := int32(d.Get("number_of_instances").(int))
	var onetimeDowngradeInfo *models.OnetimeUpgradeInfo = nil
	onetimeDowngradeInfoInterface, onetimeDowngradeInfoIsSet := d.GetOk("onetime_downgrade_info")
	if onetimeDowngradeInfoIsSet {
		onetimeDowngradeInfoMap := onetimeDowngradeInfoInterface.([]interface{})[0].(map[string]interface{})
		onetimeDowngradeInfo = OnetimeUpgradeInfoModel(onetimeDowngradeInfoMap)
	}
	var onetimeUpgradeInfo *models.OnetimeUpgradeInfo = nil
	onetimeUpgradeInfoInterface, onetimeUpgradeInfoIsSet := d.GetOk("onetime_upgrade_info")
	if onetimeUpgradeInfoIsSet {
		onetimeUpgradeInfoMap := onetimeUpgradeInfoInterface.([]interface{})[0].(map[string]interface{})
		onetimeUpgradeInfo = OnetimeUpgradeInfoModel(onetimeUpgradeInfoMap)
	}
	resendIval := int32(d.Get("resend_ival").(int))
	specifiedCollectorDeviceGroupID := int32(d.Get("specified_collector_device_group_id").(int))
	suppressAlertClear := d.Get("suppress_alert_clear").(bool)
	
	return &models.Collector {
		Arch: &arch,
		AutomaticUpgradeInfo: automaticUpgradeInfo,
		BackupAgentID: backupAgentID,
		Build: build,
		CollectorGroupID: collectorGroupID,
		CollectorSize: &collectorSize,
		Company: company,
		CustomProperties: customProperties,
		Description: description,
		Ea: &ea,
		EnableFailBack: enableFailBack,
		EnableFailOverOnCollectorDevice: enableFailOverOnCollectorDevice,
		EscalatingChainID: escalatingChainID,
		ID: int32(id),
		MonitorOthers: &monitorOthers,
		NeedAutoCreateCollectorDevice: needAutoCreateCollectorDevice,
		NumberOfInstances: numberOfInstances,
		OnetimeDowngradeInfo: onetimeDowngradeInfo,
		OnetimeUpgradeInfo: onetimeUpgradeInfo,
		ResendIval: resendIval,
		SpecifiedCollectorDeviceGroupID: specifiedCollectorDeviceGroupID,
		SuppressAlertClear: suppressAlertClear,
	}
}
func GetCollectorPropertyFields() (t []string) {
	return []string{
		"arch",
		"automatic_upgrade_info",
		"backup_agent_id",
		"build",
		"collector_group_id",
		"collector_size",
		"company",
		"custom_properties",
		"description",
		"ea",
		"enable_fail_back",
		"enable_fail_over_on_collector_device",
		"escalating_chain_id",
		"id",
		"monitor_others",
		"need_auto_create_collector_device",
		"number_of_instances",
		"onetime_downgrade_info",
		"onetime_upgrade_info",
		"resend_ival",
		"specified_collector_device_group_id",
		"suppress_alert_clear",
	}
}