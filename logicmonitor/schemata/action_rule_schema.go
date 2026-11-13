package schemata

import (
	"strconv"
	"terraform-provider-logicmonitor/logicmonitor/utils"
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ActionRuleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action_chain": {
			Type: schema.TypeList, //GoType: ActionChainSummary
			Elem: &schema.Resource{
				Schema: ActionChainSummarySchema(),
			},
			Computed: true,
		},
		
		"action_chain_id": {
			Type: schema.TypeInt,
			Required: true,
		},
		
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
			Required: true,
		},
		
		"devices": {
			Type: schema.TypeSet,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Set:      schema.HashString,
			Optional: true,
		},
		
		"enabled": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"id": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"instance": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"level_str": {
			Type: schema.TypeString,
			Required: true,
		},
		
		"name": {
			Type: schema.TypeString,
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
		
	}
}


// Schema mapping representing the resource's respective datasource object defined in Terraform configuration
// Only difference between this and ActionRuleSchema() are the computabilty of the id field and the inclusion of a filter field for datasources
func DataSourceActionRuleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action_chain": {
			Type: schema.TypeList, //GoType: ActionChainSummary
			Elem: &schema.Resource{
				Schema: ActionChainSummarySchema(),
			},
			Optional: true,
		},
		
		"action_chain_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
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
		
		"enabled": {
			Type: schema.TypeBool,
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
		
		"resource_properties": {
			Type: schema.TypeList, //GoType: []*DeviceProperty 
			Elem: &schema.Resource{
				Schema: DevicePropertySchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"filter": {
			Type:     schema.TypeString,
            Optional: true,
		},
	}
}

// DataSourceActionRuleListSchema is the schema of the plural "logicmonitor_action_rules" data source.
// "items" holds one read-only entry per ActionRule matching "filter", with the same attributes as DataSourceActionRuleSchema().
func DataSourceActionRuleListSchema() map[string]*schema.Schema {
	itemSchema := DataSourceActionRuleSchema()
	delete(itemSchema, "filter")

	return map[string]*schema.Schema{
		"filter": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"page_size": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  300,
		},
		"total": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"items": {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Resource{Schema: ComputedOnlySchema(itemSchema)},
		},
	}
}

func SetActionRuleResourceData(d *schema.ResourceData, m *models.ActionRule) {
	d.Set("action_chain", SetActionChainSummarySubResourceData([]*models.ActionChainSummary{m.ActionChain}))
	d.Set("action_chain_id", m.ActionChainID)
	d.Set("datapoint", m.Datapoint)
	d.Set("datasource", m.Datasource)
	d.Set("device_groups", m.DeviceGroups)
	d.Set("devices", m.Devices)
	d.Set("enabled", m.Enabled)
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("instance", m.Instance)
	d.Set("level_str", m.LevelStr)
	d.Set("name", m.Name)
	d.Set("resource_properties", SetDevicePropertySubResourceData(m.ResourceProperties))
}

func SetActionRuleSubResourceData(m []*models.ActionRule) (d []*map[string]interface{}) {
	for _, actionRule := range m {
		if actionRule != nil {
			properties := make(map[string]interface{})
			properties["action_chain"] = SetActionChainSummarySubResourceData([]*models.ActionChainSummary{actionRule.ActionChain})
			properties["action_chain_id"] = actionRule.ActionChainID
			properties["datapoint"] = actionRule.Datapoint
			properties["datasource"] = actionRule.Datasource
			properties["device_groups"] = actionRule.DeviceGroups
			properties["devices"] = actionRule.Devices
			properties["enabled"] = actionRule.Enabled
			properties["id"] = actionRule.ID
			properties["instance"] = actionRule.Instance
			properties["level_str"] = actionRule.LevelStr
			properties["name"] = actionRule.Name
			properties["resource_properties"] = SetDevicePropertySubResourceData(actionRule.ResourceProperties)
			d = append(d, &properties)
		}
	}
	return
}

func ActionRuleModel(d *schema.ResourceData) *models.ActionRule {
	actionChainID := int32(d.Get("action_chain_id").(int))
	datapoint := d.Get("datapoint").(string)
	datasource := d.Get("datasource").(string)
	deviceGroups := utils.ConvertSetToStringSlice(d.Get("device_groups").(*schema.Set))
	devices := utils.ConvertSetToStringSlice(d.Get("devices").(*schema.Set))
	enabled := d.Get("enabled").(bool)
	id, _ := strconv.Atoi(d.Get("id").(string))
	instance := d.Get("instance").(string)
	levelStr := d.Get("level_str").(string)
	name := d.Get("name").(string)
	resourceProperties := utils.GetResourceProperties(d.Get("resource_properties").([]interface{}))
	
	return &models.ActionRule {
		ActionChainID: &actionChainID,
		Datapoint: datapoint,
		Datasource: datasource,
		DeviceGroups: deviceGroups,
		Devices: devices,
		Enabled: enabled,
		ID: int32(id),
		Instance: instance,
		LevelStr: &levelStr,
		Name: &name,
		ResourceProperties: resourceProperties,
	}
}
func GetActionRulePropertyFields() (t []string) {
	return []string{
		"action_chain_id",
		"datapoint",
		"datasource",
		"device_groups",
		"devices",
		"enabled",
		"id",
		"instance",
		"level_str",
		"name",
		"resource_properties",
	}
}