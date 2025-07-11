package schemata

import (
	"strconv"
	"terraform-provider-logicmonitor/logicmonitor/utils"
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func AlertRuleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"datapoint": {
			Type: schema.TypeString,
			Required: true,
		},
		
		"datasource": {
			Type: schema.TypeString,
			Required: true,
		},
		
		"device_groups": {
			Type: schema.TypeSet,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Set:      schema.HashString,
			Required: true,
		},
		
		"devices": {
			Type: schema.TypeSet,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Set:      schema.HashString,
			Required: true,
		},
		
		"escalating_chain": {
			Type: schema.TypeMap, //GoType: interface{}
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Computed: true,
		},
		
		"escalating_chain_id": {
			Type: schema.TypeInt,
			Required: true,
		},
		
		"escalation_interval": {
			Type: schema.TypeInt,
			Required: true,
		},
		
		"id": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"instance": {
			Type: schema.TypeString,
			Required: true,
		},
		
		"level_str": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"name": {
			Type: schema.TypeString,
			Required: true,
		},
		
		"priority": {
			Type: schema.TypeInt,
			Required: true,
		},
		
		"resource_properties": {
			Type: schema.TypeList, //GoType: []*DeviceProperty  
			Elem: &schema.Resource{
				Schema: DevicePropertySchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"send_anomaly_suppressed_alert": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"suppress_alert_ack_sdt": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"suppress_alert_clear": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
	}
}


// Schema mapping representing the resource's respective datasource object defined in Terraform configuration
// Only difference between this and AlertRuleSchema() are the computabilty of the id field and the inclusion of a filter field for datasources
func DataSourceAlertRuleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"datapoint": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"datasource": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"device_groups": {
			Type: schema.TypeSet,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Set:      schema.HashString,
			Optional: true,
		},
		
		"devices": {
			Type: schema.TypeSet,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Set:      schema.HashString,
			Optional: true,
		},
		
		"escalating_chain": {
			Type: schema.TypeMap, //GoType: interface{}
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},
		
		"escalating_chain_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"escalation_interval": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"id": {
			Type: schema.TypeInt,
			Computed: true,
			Optional: true,
		},
		
		"instance": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"level_str": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"name": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"priority": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"resource_properties": {
			Type: schema.TypeList, //GoType: []*DeviceProperty 
			Elem: &schema.Resource{
				Schema: DevicePropertySchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"send_anomaly_suppressed_alert": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"suppress_alert_ack_sdt": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"suppress_alert_clear": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"filter": {
			Type:     schema.TypeString,
            Optional: true,
		},
	}
}

func SetAlertRuleResourceData(d *schema.ResourceData, m *models.AlertRule) {
	d.Set("datapoint", m.Datapoint)
	d.Set("datasource", m.Datasource)
	d.Set("device_groups", m.DeviceGroups)
	d.Set("devices", m.Devices)
	d.Set("escalating_chain", m.EscalatingChain)
	d.Set("escalating_chain_id", m.EscalatingChainID)
	d.Set("escalation_interval", m.EscalationInterval)
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("instance", m.Instance)
	d.Set("level_str", m.LevelStr)
	d.Set("name", m.Name)
	d.Set("priority", m.Priority)
	d.Set("resource_properties", SetDevicePropertySubResourceData(m.ResourceProperties))
	d.Set("send_anomaly_suppressed_alert", m.SendAnomalySuppressedAlert)
	d.Set("suppress_alert_ack_sdt", m.SuppressAlertAckSdt)
	d.Set("suppress_alert_clear", m.SuppressAlertClear)
}

func SetAlertRuleSubResourceData(m []*models.AlertRule) (d []*map[string]interface{}) {
	for _, alertRule := range m {
		if alertRule != nil {
			properties := make(map[string]interface{})
			properties["datapoint"] = alertRule.Datapoint
			properties["datasource"] = alertRule.Datasource
			properties["device_groups"] = alertRule.DeviceGroups
			properties["devices"] = alertRule.Devices
			properties["escalating_chain"] = alertRule.EscalatingChain
			properties["escalating_chain_id"] = alertRule.EscalatingChainID
			properties["escalation_interval"] = alertRule.EscalationInterval
			properties["id"] = alertRule.ID
			properties["instance"] = alertRule.Instance
			properties["level_str"] = alertRule.LevelStr
			properties["name"] = alertRule.Name
			properties["priority"] = alertRule.Priority
			properties["resource_properties"] = SetDevicePropertySubResourceData(alertRule.ResourceProperties)
			properties["send_anomaly_suppressed_alert"] = alertRule.SendAnomalySuppressedAlert
			properties["suppress_alert_ack_sdt"] = alertRule.SuppressAlertAckSdt
			properties["suppress_alert_clear"] = alertRule.SuppressAlertClear
			d = append(d, &properties)
		}
	}
	return
}

func AlertRuleModel(d *schema.ResourceData) *models.AlertRule {
	datapoint := d.Get("datapoint").(string)
	datasource := d.Get("datasource").(string)
	deviceGroups := utils.ConvertSetToStringSlice(d.Get("device_groups").(*schema.Set))
	devices := utils.ConvertSetToStringSlice(d.Get("devices").(*schema.Set))
	escalatingChainID := int32(d.Get("escalating_chain_id").(int))
	escalationInterval := int32(d.Get("escalation_interval").(int))
	id, _ := strconv.Atoi(d.Get("id").(string))
	instance := d.Get("instance").(string)
	levelStr := d.Get("level_str").(string)
	name := d.Get("name").(string)
	priority := int32(d.Get("priority").(int))
	resourceProperties := utils.GetResourceProperties(d.Get("resource_properties").([]interface{}))
	sendAnomalySuppressedAlert := d.Get("send_anomaly_suppressed_alert").(bool)
	suppressAlertAckSdt := d.Get("suppress_alert_ack_sdt").(bool)
	suppressAlertClear := d.Get("suppress_alert_clear").(bool)
	
	return &models.AlertRule {
		Datapoint: &datapoint,
		Datasource: &datasource,
		DeviceGroups: deviceGroups,
		Devices: devices,
		EscalatingChainID: &escalatingChainID,
		EscalationInterval: &escalationInterval,
		ID: int32(id),
		Instance: &instance,
		LevelStr: levelStr,
		Name: &name,
		Priority: &priority,
		ResourceProperties: resourceProperties,
		SendAnomalySuppressedAlert: sendAnomalySuppressedAlert,
		SuppressAlertAckSdt: suppressAlertAckSdt,
		SuppressAlertClear: suppressAlertClear,
	}
}
func GetAlertRulePropertyFields() (t []string) {
	return []string{
		"datapoint",
		"datasource",
		"device_groups",
		"devices",
		"escalating_chain_id",
		"escalation_interval",
		"id",
		"instance",
		"level_str",
		"name",
		"priority",
		"resource_properties",
		"send_anomaly_suppressed_alert",
		"suppress_alert_ack_sdt",
		"suppress_alert_clear",
	}
}