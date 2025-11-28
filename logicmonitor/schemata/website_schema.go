package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"terraform-provider-logicmonitor/logicmonitor/utils"
	"terraform-provider-logicmonitor/models"
)

func WebsiteSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"alert_expr": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"disable_alerting": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"domain": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"global_sm_alert_cond": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"group_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"host": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"ignore_s_s_l": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"individual_alert_level": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"individual_sm_alert_enable": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"is_internal": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"last_updated": {
			Type:     schema.TypeInt,
			Computed: true,
		},

		"name": {
			Type:     schema.TypeString,
			Required: true,
		},

		"overall_alert_level": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"polling_interval": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"schema": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"status": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"steps": {
			Type: schema.TypeList, //GoType: []*WebCheckStep
			Elem: &schema.Resource{
				Schema: WebCheckStepSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"stop_monitoring": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"stop_monitoring_by_folder": {
			Type:     schema.TypeBool,
			Computed: true,
		},

		"template": {
			Type: schema.TypeMap, //GoType: interface{}
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"test_location": {
			Type: schema.TypeList, //GoType: WebsiteLocation
			Elem: &schema.Resource{
				Schema: WebsiteLocationSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"transition": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"trigger_s_s_l_expiration_alert": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"trigger_s_s_l_status_alert": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"type": {
			Type:     schema.TypeString,
			Required: true,
		},

		"use_default_alert_setting": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"use_default_location_setting": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"user_permission": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

// Schema mapping representing the resource's respective datasource object defined in Terraform configuration
// Only difference between this and WebsiteSchema() are the computabilty of the id field and the inclusion of a filter field for datasources
func DataSourceWebsiteSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"alert_expr": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"disable_alerting": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"domain": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"global_sm_alert_cond": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"group_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"host": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"id": {
			Type:     schema.TypeInt,
			Computed: true,
			Optional: true,
		},

		"ignore_s_s_l": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"individual_alert_level": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"individual_sm_alert_enable": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"is_internal": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"last_updated": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"overall_alert_level": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"polling_interval": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"schema": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"status": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"steps": {
			Type: schema.TypeList, //GoType: []*WebCheckStep
			Elem: &schema.Resource{
				Schema: WebCheckStepSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"stop_monitoring": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"stop_monitoring_by_folder": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"template": {
			Type: schema.TypeMap, //GoType: interface{}
			Elem: &schema.Schema{
				Type: schema.TypeString,
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

		"transition": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"trigger_s_s_l_expiration_alert": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"trigger_s_s_l_status_alert": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"type": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"use_default_alert_setting": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"use_default_location_setting": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"user_permission": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"filter": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func SetWebsiteResourceData(d *schema.ResourceData, m *models.Website) {
	d.Set("alert_expr", m.AlertExpr)
	d.Set("description", m.Description)
	d.Set("disable_alerting", m.DisableAlerting)
	d.Set("domain", m.Domain)
	d.Set("global_sm_alert_cond", m.GlobalSmAlertCond)
	d.Set("group_id", m.GroupID)
	d.Set("host", m.Host)
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("ignore_s_s_l", m.IgnoreSSL)
	d.Set("individual_alert_level", m.IndividualAlertLevel)
	d.Set("individual_sm_alert_enable", m.IndividualSmAlertEnable)
	d.Set("is_internal", m.IsInternal)
	d.Set("last_updated", m.LastUpdated)
	d.Set("name", m.Name)
	d.Set("overall_alert_level", m.OverallAlertLevel)
	d.Set("polling_interval", m.PollingInterval)
	d.Set("schema", m.Schema)
	d.Set("status", m.Status)
	d.Set("steps", SetWebCheckStepSubResourceData(m.Steps))
	d.Set("stop_monitoring", m.StopMonitoring)
	d.Set("stop_monitoring_by_folder", m.StopMonitoringByFolder)
	d.Set("template", m.Template)
	d.Set("test_location", SetWebsiteLocationSubResourceData([]*models.WebsiteLocation{m.TestLocation}))
	d.Set("transition", m.Transition)
	d.Set("trigger_s_s_l_expiration_alert", m.TriggerSSLExpirationAlert)
	d.Set("trigger_s_s_l_status_alert", m.TriggerSSLStatusAlert)
	d.Set("type", m.Type)
	d.Set("use_default_alert_setting", m.UseDefaultAlertSetting)
	d.Set("use_default_location_setting", m.UseDefaultLocationSetting)
	d.Set("user_permission", m.UserPermission)
}

func SetWebsiteSubResourceData(m []*models.Website) (d []*map[string]interface{}) {
	for _, website := range m {
		if website != nil {
			properties := make(map[string]interface{})
			properties["alert_expr"] = website.AlertExpr
			properties["description"] = website.Description
			properties["disable_alerting"] = website.DisableAlerting
			properties["domain"] = website.Domain
			properties["global_sm_alert_cond"] = website.GlobalSmAlertCond
			properties["group_id"] = website.GroupID
			properties["host"] = website.Host
			properties["id"] = website.ID
			properties["ignore_s_s_l"] = website.IgnoreSSL
			properties["individual_alert_level"] = website.IndividualAlertLevel
			properties["individual_sm_alert_enable"] = website.IndividualSmAlertEnable
			properties["is_internal"] = website.IsInternal
			properties["last_updated"] = website.LastUpdated
			properties["name"] = website.Name
			properties["overall_alert_level"] = website.OverallAlertLevel
			properties["polling_interval"] = website.PollingInterval
			properties["schema"] = website.Schema
			properties["status"] = website.Status
			properties["steps"] = SetWebCheckStepSubResourceData(website.Steps)
			properties["stop_monitoring"] = website.StopMonitoring
			properties["stop_monitoring_by_folder"] = website.StopMonitoringByFolder
			properties["template"] = website.Template
			properties["test_location"] = SetWebsiteLocationSubResourceData([]*models.WebsiteLocation{website.TestLocation})
			properties["transition"] = website.Transition
			properties["trigger_s_s_l_expiration_alert"] = website.TriggerSSLExpirationAlert
			properties["trigger_s_s_l_status_alert"] = website.TriggerSSLStatusAlert
			properties["type"] = website.Type
			properties["use_default_alert_setting"] = website.UseDefaultAlertSetting
			properties["use_default_location_setting"] = website.UseDefaultLocationSetting
			properties["user_permission"] = website.UserPermission
			d = append(d, &properties)
		}
	}
	return
}

func WebsiteModel(d *schema.ResourceData) *models.Website {
	alertExpr := d.Get("alert_expr").(string)
	description := d.Get("description").(string)
	disableAlerting := d.Get("disable_alerting").(bool)
	domain := d.Get("domain").(string)
	globalSmAlertCond := int32(d.Get("global_sm_alert_cond").(int))
	groupID := int32(d.Get("group_id").(int))
	host := d.Get("host").(string)
	id, _ := strconv.Atoi(d.Get("id").(string))
	ignoreSSL := d.Get("ignore_s_s_l").(bool)
	individualAlertLevel := d.Get("individual_alert_level").(string)
	individualSmAlertEnable := d.Get("individual_sm_alert_enable").(bool)
	isInternal := d.Get("is_internal").(bool)
	name := d.Get("name").(string)
	overallAlertLevel := d.Get("overall_alert_level").(string)
	pollingInterval := int32(d.Get("polling_interval").(int))
	schema := d.Get("schema").(string)
	steps := utils.GetPropFromSteps(d, "steps")
	stopMonitoring := d.Get("stop_monitoring").(bool)
	template := d.Get("template")
	testLocation := utils.GetPropFromLocationMap(d, "test_location")
	transition := int32(d.Get("transition").(int))
	triggerSSLExpirationAlert := d.Get("trigger_s_s_l_expiration_alert").(bool)
	triggerSSLStatusAlert := d.Get("trigger_s_s_l_status_alert").(bool)
	typeVar := d.Get("type").(string)
	useDefaultAlertSetting := d.Get("use_default_alert_setting").(bool)
	useDefaultLocationSetting := d.Get("use_default_location_setting").(bool)
	userPermission := d.Get("user_permission").(string)

	return &models.Website{
		AlertExpr:                 alertExpr,
		Description:               description,
		DisableAlerting:           disableAlerting,
		Domain:                    domain,
		GlobalSmAlertCond:         globalSmAlertCond,
		GroupID:                   groupID,
		Host:                      host,
		ID:                        int32(id),
		IgnoreSSL:                 ignoreSSL,
		IndividualAlertLevel:      individualAlertLevel,
		IndividualSmAlertEnable:   individualSmAlertEnable,
		IsInternal:                isInternal,
		Name:                      &name,
		OverallAlertLevel:         overallAlertLevel,
		PollingInterval:           pollingInterval,
		Schema:                    schema,
		Steps:                     steps,
		StopMonitoring:            stopMonitoring,
		Template:                  template,
		TestLocation:              testLocation,
		Transition:                transition,
		TriggerSSLExpirationAlert: triggerSSLExpirationAlert,
		TriggerSSLStatusAlert:     triggerSSLStatusAlert,
		Type:                      &typeVar,
		UseDefaultAlertSetting:    useDefaultAlertSetting,
		UseDefaultLocationSetting: useDefaultLocationSetting,
		UserPermission:            userPermission,
	}
}
func GetWebsitePropertyFields() (t []string) {
	return []string{
		"alert_expr",
		"description",
		"disable_alerting",
		"domain",
		"global_sm_alert_cond",
		"group_id",
		"host",
		"id",
		"ignore_s_s_l",
		"individual_alert_level",
		"individual_sm_alert_enable",
		"is_internal",
		"name",
		"overall_alert_level",
		"polling_interval",
		"schema",
		"steps",
		"stop_monitoring",
		"template",
		"test_location",
		"transition",
		"trigger_s_s_l_expiration_alert",
		"trigger_s_s_l_status_alert",
		"type",
		"use_default_alert_setting",
		"use_default_location_setting",
		"user_permission",
	}
}
