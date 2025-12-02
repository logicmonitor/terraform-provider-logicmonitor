package schemata

import (
	"strconv"
	"terraform-provider-logicmonitor/logicmonitor/utils"
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func RoleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"acct_require_two_f_a": {
			Type: schema.TypeBool,
			Computed: true,
		},
		
		"associated_user_count": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"custom_help_label": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"custom_help_url": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"description": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"enable_remote_session_in_company_level": {
			Type: schema.TypeBool,
			Computed: true,
		},
		
		"id": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"name": {
			Type: schema.TypeString,
			Required: true,
		},
		
		"privileges": {
			Type: schema.TypeList, //GoType: []*Privilege  
			Elem: &schema.Resource{
				Schema: PrivilegeSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Required: true,
		},
		
		"require_e_u_l_a": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"role_group_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"two_f_a_required": {
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
// Only difference between this and RoleSchema() are the computabilty of the id field and the inclusion of a filter field for datasources
func DataSourceRoleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"acct_require_two_f_a": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"associated_user_count": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"custom_help_label": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"custom_help_url": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"description": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"enable_remote_session_in_company_level": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"id": {
			Type: schema.TypeInt,
			Computed: true,
			Optional: true,
		},
		
		"name": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"privileges": {
			Type: schema.TypeList, //GoType: []*Privilege 
			Elem: &schema.Resource{
				Schema: PrivilegeSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"require_e_u_l_a": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"role_group_id": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"two_f_a_required": {
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

func SetRoleResourceData(d *schema.ResourceData, m *models.Role) {
	d.Set("acct_require_two_f_a", m.AcctRequireTwoFA)
	d.Set("associated_user_count", m.AssociatedUserCount)
	d.Set("custom_help_label", m.CustomHelpLabel)
	d.Set("custom_help_url", m.CustomHelpURL)
	d.Set("description", m.Description)
	d.Set("enable_remote_session_in_company_level", m.EnableRemoteSessionInCompanyLevel)
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("name", m.Name)
	d.Set("privileges", SetPrivilegeSubResourceData(m.Privileges))
	d.Set("require_e_u_l_a", m.RequireEULA)
	d.Set("role_group_id", m.RoleGroupID)
	d.Set("two_f_a_required", m.TwoFARequired)
	d.Set("user_permission", m.UserPermission)
}

func SetRoleSubResourceData(m []*models.Role) (d []*map[string]interface{}) {
	for _, role := range m {
		if role != nil {
			properties := make(map[string]interface{})
			properties["acct_require_two_f_a"] = role.AcctRequireTwoFA
			properties["associated_user_count"] = role.AssociatedUserCount
			properties["custom_help_label"] = role.CustomHelpLabel
			properties["custom_help_url"] = role.CustomHelpURL
			properties["description"] = role.Description
			properties["enable_remote_session_in_company_level"] = role.EnableRemoteSessionInCompanyLevel
			properties["id"] = role.ID
			properties["name"] = role.Name
			properties["privileges"] = SetPrivilegeSubResourceData(role.Privileges)
			properties["require_e_u_l_a"] = role.RequireEULA
			properties["role_group_id"] = role.RoleGroupID
			properties["two_f_a_required"] = role.TwoFARequired
			properties["user_permission"] = role.UserPermission
			d = append(d, &properties)
		}
	}
	return
}

func RoleModel(d *schema.ResourceData) *models.Role {
	customHelpLabel := d.Get("custom_help_label").(string)
	customHelpURL := d.Get("custom_help_url").(string)
	description := d.Get("description").(string)
	id, _ := strconv.Atoi(d.Get("id").(string))
	name := d.Get("name").(string)
	privileges := utils.GetPropFromPriviligeMap(d, "privileges")
	requireEULA := d.Get("require_e_u_l_a").(bool)
	roleGroupID := int32(d.Get("role_group_id").(int))
	twoFARequired := d.Get("two_f_a_required").(bool)
	
	return &models.Role {
		CustomHelpLabel: customHelpLabel,
		CustomHelpURL: customHelpURL,
		Description: description,
		ID: int32(id),
		Name: &name,
		Privileges: privileges,
		RequireEULA: requireEULA,
		RoleGroupID: roleGroupID,
		TwoFARequired: twoFARequired,
	}
}
func GetRolePropertyFields() (t []string) {
	return []string{
		"custom_help_label",
		"custom_help_url",
		"description",
		"id",
		"name",
		"privileges",
		"require_e_u_l_a",
		"role_group_id",
		"two_f_a_required",
	}
}