package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"terraform-provider-logicmonitor/logicmonitor/utils"
	"terraform-provider-logicmonitor/models"
)

func DashboardSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"full_name": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"group_full_path": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"group_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"group_name": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"name": {
			Type:     schema.TypeString,
			Required: true,
		},

		"owner": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"sharable": {
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

		"user_permission": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"widget_tokens": {
			Type: schema.TypeList, //GoType: []*WidgetToken
			Elem: &schema.Resource{
				Schema: WidgetTokenSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"widgets_config": {
			Type: schema.TypeMap, //GoType: interface{}
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},
	}
}

// Schema mapping representing the resource's respective datasource object defined in Terraform configuration
// Only difference between this and DashboardSchema() are the computabilty of the id field and the inclusion of a filter field for datasources
func DataSourceDashboardSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"full_name": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"group_full_path": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"group_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"group_name": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"id": {
			Type:     schema.TypeInt,
			Computed: true,
			Optional: true,
		},

		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"owner": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"sharable": {
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

		"user_permission": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"widget_tokens": {
			Type: schema.TypeList, //GoType: []*WidgetToken
			Elem: &schema.Resource{
				Schema: WidgetTokenSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"widgets_config": {
			Type: schema.TypeMap, //GoType: interface{}
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"filter": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func SetDashboardResourceData(d *schema.ResourceData, m *models.Dashboard) {
	d.Set("description", m.Description)
	d.Set("full_name", m.FullName)
	d.Set("group_full_path", m.GroupFullPath)
	d.Set("group_id", m.GroupID)
	d.Set("group_name", m.GroupName)
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("name", m.Name)
	d.Set("owner", m.Owner)
	d.Set("sharable", m.Sharable)
	d.Set("template", m.Template)
	d.Set("user_permission", m.UserPermission)
	d.Set("widget_tokens", SetWidgetTokenSubResourceData(m.WidgetTokens))
	d.Set("widgets_config", m.WidgetsConfig)
}

func SetDashboardSubResourceData(m []*models.Dashboard) (d []*map[string]interface{}) {
	for _, dashboard := range m {
		if dashboard != nil {
			properties := make(map[string]interface{})
			properties["description"] = dashboard.Description
			properties["full_name"] = dashboard.FullName
			properties["group_full_path"] = dashboard.GroupFullPath
			properties["group_id"] = dashboard.GroupID
			properties["group_name"] = dashboard.GroupName
			properties["id"] = dashboard.ID
			properties["name"] = dashboard.Name
			properties["owner"] = dashboard.Owner
			properties["sharable"] = dashboard.Sharable
			properties["template"] = dashboard.Template
			properties["user_permission"] = dashboard.UserPermission
			properties["widget_tokens"] = SetWidgetTokenSubResourceData(dashboard.WidgetTokens)
			properties["widgets_config"] = dashboard.WidgetsConfig
			d = append(d, &properties)
		}
	}
	return
}

func DashboardModel(d *schema.ResourceData) *models.Dashboard {
	description := d.Get("description").(string)
	groupID := int32(d.Get("group_id").(int))
	groupName := d.Get("group_name").(string)
	id, _ := strconv.Atoi(d.Get("id").(string))
	name := d.Get("name").(string)
	owner := d.Get("owner").(string)
	sharable := d.Get("sharable").(bool)
	template := d.Get("template")
	widgetTokens := utils.GetPropFromWTMap(d, "widget_tokens")
	widgetsConfig := d.Get("widgets_config")

	return &models.Dashboard{
		Description:   description,
		GroupID:       groupID,
		GroupName:     groupName,
		ID:            int32(id),
		Name:          &name,
		Owner:         owner,
		Sharable:      sharable,
		Template:      template,
		WidgetTokens:  widgetTokens,
		WidgetsConfig: widgetsConfig,
	}
}
func GetDashboardPropertyFields() (t []string) {
	return []string{
		"description",
		"group_id",
		"group_name",
		"id",
		"name",
		"owner",
		"sharable",
		"template",
		"widget_tokens",
		"widgets_config",
	}
}
