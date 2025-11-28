package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"terraform-provider-logicmonitor/logicmonitor/utils"
	"terraform-provider-logicmonitor/models"
)

func EscalationChainSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cc_destinations": {
			Type: schema.TypeList, //GoType: []*Recipient
			Elem: &schema.Resource{
				Schema: RecipientSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"destinations": {
			Type: schema.TypeList, //GoType: []*Chain
			Elem: &schema.Resource{
				Schema: ChainSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Required:   true,
		},

		"enable_throttling": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"in_alerting": {
			Type:     schema.TypeBool,
			Computed: true,
		},

		"name": {
			Type:     schema.TypeString,
			Required: true,
		},

		"throttling_alerts": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"throttling_period": {
			Type:     schema.TypeInt,
			Optional: true,
		},
	}
}

// Schema mapping representing the resource's respective datasource object defined in Terraform configuration
// Only difference between this and EscalationChainSchema() are the computabilty of the id field and the inclusion of a filter field for datasources
func DataSourceEscalationChainSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cc_destinations": {
			Type: schema.TypeList, //GoType: []*Recipient
			Elem: &schema.Resource{
				Schema: RecipientSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"destinations": {
			Type: schema.TypeList, //GoType: []*Chain
			Elem: &schema.Resource{
				Schema: ChainSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"enable_throttling": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"id": {
			Type:     schema.TypeInt,
			Computed: true,
			Optional: true,
		},

		"in_alerting": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"throttling_alerts": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"throttling_period": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"filter": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func SetEscalationChainResourceData(d *schema.ResourceData, m *models.EscalationChain) {
	d.Set("cc_destinations", SetRecipientSubResourceData(m.CcDestinations))
	d.Set("description", m.Description)
	d.Set("destinations", SetChainSubResourceData(m.Destinations))
	d.Set("enable_throttling", m.EnableThrottling)
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("in_alerting", m.InAlerting)
	d.Set("name", m.Name)
	d.Set("throttling_alerts", m.ThrottlingAlerts)
	d.Set("throttling_period", m.ThrottlingPeriod)
}

func SetEscalationChainSubResourceData(m []*models.EscalationChain) (d []*map[string]interface{}) {
	for _, escalationChain := range m {
		if escalationChain != nil {
			properties := make(map[string]interface{})
			properties["cc_destinations"] = SetRecipientSubResourceData(escalationChain.CcDestinations)
			properties["description"] = escalationChain.Description
			properties["destinations"] = SetChainSubResourceData(escalationChain.Destinations)
			properties["enable_throttling"] = escalationChain.EnableThrottling
			properties["id"] = escalationChain.ID
			properties["in_alerting"] = escalationChain.InAlerting
			properties["name"] = escalationChain.Name
			properties["throttling_alerts"] = escalationChain.ThrottlingAlerts
			properties["throttling_period"] = escalationChain.ThrottlingPeriod
			d = append(d, &properties)
		}
	}
	return
}

func EscalationChainModel(d *schema.ResourceData) *models.EscalationChain {
	ccDestinations := utils.GetPropFromCCMap(d, "cc_destinations")
	description := d.Get("description").(string)
	destinations := utils.GetPropFromDesMap(d, "destinations")
	enableThrottling := d.Get("enable_throttling").(bool)
	id, _ := strconv.Atoi(d.Get("id").(string))
	name := d.Get("name").(string)
	throttlingAlerts := int32(d.Get("throttling_alerts").(int))
	throttlingPeriod := int32(d.Get("throttling_period").(int))

	return &models.EscalationChain{
		CcDestinations:   ccDestinations,
		Description:      description,
		Destinations:     destinations,
		EnableThrottling: enableThrottling,
		ID:               int32(id),
		Name:             &name,
		ThrottlingAlerts: throttlingAlerts,
		ThrottlingPeriod: throttlingPeriod,
	}
}
func GetEscalationChainPropertyFields() (t []string) {
	return []string{
		"cc_destinations",
		"description",
		"destinations",
		"enable_throttling",
		"id",
		"name",
		"throttling_alerts",
		"throttling_period",
	}
}
