package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func CloudServicesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"api_g_a_t_e_w_a_y": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"api_m_a_n_a_g_e_m_e_n_t": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"a_p_p_l_i_c_a_t_i_o_n_e_l_b": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"a_p_p_l_i_c_a_t_i_o_n_g_a_t_e_w_a_y": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"a_p_p_l_i_c_a_t_i_o_n_i_n_s_i_g_h_t_s": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"a_p_p_s_e_r_v_i_c_e": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"a_p_p_s_e_r_v_i_c_e_p_l_a_n": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"a_p_p_s_t_r_e_a_m": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"a_t_h_e_n_a": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"a_u_t_o_m_a_t_i_o_n_a_c_c_o_u_n_t": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"a_u_t_o_s_c_a_l_i_n_g": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"b_a_c_k_u_p_p_r_o_t_e_c_t_e_d_i_t_e_m_s": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"b_l_o_b_s_t_o_r_a_g_e": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"c_l_o_u_d_f_r_o_n_t": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"c_l_o_u_d_s_e_a_r_c_h": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"c_o_d_e_b_ui_l_d": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"c_o_g_n_i_t_i_v_e_s_e_a_r_c_h": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"c_o_g_n_i_t_i_v_e_s_e_r_v_i_c_e_s": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"c_o_g_n_i_t_o": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"c_o_s_m_o_s_d_b": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"d_a_t_a_f_a_c_t_o_r_y": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"d_i_r_e_c_t_c_o_n_n_e_c_t": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"d_m_s_r_e_p_l_i_c_a_t_i_o_n": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"d_m_s_r_e_p_l_i_c_a_t_i_o_n_t_a_s_k_s": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"d_o_c_d_b": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"d_y_n_a_m_o_d_b": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"e_b_s": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"e_c2": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"e_c_s": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"e_f_s": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"e_l_a_s_t_i_c_a_c_h_e": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"e_l_a_s_t_i_c_b_e_a_n_s_t_a_l_k": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"e_l_a_s_t_i_c_s_e_a_r_c_h": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"e_l_a_s_t_i_c_t_r_a_n_s_c_o_d_e_r": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"e_l_b": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"e_m_r": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"e_v_e_n_t_b_r_id_g_e": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"e_v_e_n_t_h_u_b": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"e_x_p_r_e_s_s_r_o_u_t_e_c_i_r_c_ui_t": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"f_i_l_e_s_t_o_r_a_g_e": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"f_i_r_e_h_o_s_e": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"f_i_r_e_w_a_l_l": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"f_r_o_n_t_d_o_o_r_s": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"f_s_x": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"g_l_u_e": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"k_e_y_v_a_u_l_t": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"k_i_n_e_s_i_s": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"k_i_n_e_s_i_s_v_id_e_o": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"l_a_m_b_d_a": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"l_o_a_d_b_a_l_a_n_c_e_r_s": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"l_o_g_a_n_a_l_y_t_i_c_s_w_o_r_k_s_p_a_c_e_s": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"l_o_g_i_c_a_p_p_s": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"m_e_d_i_a_c_o_n_n_e_c_t": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"m_e_d_i_a_c_o_n_v_e_r_t": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"m_e_d_i_a_p_a_c_k_a_g_e_l_i_v_e": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"m_e_d_i_a_p_a_c_k_a_g_e_v_o_d": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"m_e_d_i_a_s_t_o_r_e": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"m_e_d_i_a_t_a_i_l_o_r": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"m_q": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"m_s_k_b_r_o_k_e_r": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"m_s_k_c_l_u_s_t_e_r": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"m_y_sql": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"n_a_t_g_a_t_e_w_a_y": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"n_e_t_w_o_r_k_e_l_b": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"n_e_t_w_o_r_k_i_n_t_e_r_f_a_c_e": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"o_p_s_w_o_r_k_s": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"p_o_s_t_g_r_e_sql": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"p_u_b_l_i_c_ip": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"q_u_e_u_e_s_t_o_r_a_g_e": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"r_d_s": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"r_e_c_o_v_e_r_y_p_r_o_t_e_c_t_e_d_i_t_e_m": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"r_e_c_o_v_e_r_y_s_e_r_v_i_c_e_s": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"r_e_d_i_s_c_a_c_h_e": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"r_e_d_s_h_i_f_t": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"r_o_u_t_e53": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"r_o_u_t_e53_r_e_s_o_l_v_e_r": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"s3": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"s_a_g_e_m_a_k_e_r": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"s_e_r_v_i_c_e_b_u_s": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"s_e_s": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"s_n_s": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"sql_d_a_t_a_b_a_s_e": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"sql_e_l_a_s_t_i_c_p_o_o_l": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"sql_m_a_n_a_g_e_d_i_n_s_t_a_n_c_e": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"s_q_s": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"s_t_e_p_f_u_n_c_t_i_o_n_s": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"s_t_o_r_a_g_e_a_c_c_o_u_n_t": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"s_w_f_a_c_t_i_v_i_t_y": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"s_w_f_w_o_r_k_f_l_o_w": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"s_y_n_a_p_s_e_w_o_r_k_s_p_a_c_e_s": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"t_a_b_l_e_s_t_o_r_a_g_e": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"t_r_a_f_f_i_c_m_a_n_a_g_e_r": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"t_r_a_n_s_i_t_g_a_t_e_w_a_y": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"v_i_r_t_u_a_l_d_e_s_k_t_o_p": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"v_i_r_t_u_a_l_m_a_c_h_i_n_e": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"v_i_r_t_u_a_l_m_a_c_h_i_n_e_s_c_a_l_e_s_e_t": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"v_i_r_t_u_a_l_m_a_c_h_i_n_e_s_c_a_l_e_s_e_t_vm": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"v_i_r_t_u_a_l_n_e_t_w_o_r_k_g_a_t_e_w_a_y": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"v_p_n": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"w_o_r_k_s_p_a_c_e": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
		"w_o_r_k_s_p_a_c_e_d_i_r_e_c_t_o_r_y": {
			Type: schema.TypeList, //GoType: CloudServiceSettings
			Elem: &schema.Resource{
				Schema: CloudServiceSettingsSchema(),
			},
			Optional: true,
		},
		
	}
}

func SetCloudServicesSubResourceData(m []*models.CloudServices) (d []*map[string]interface{}) {
	for _, cloudServices := range m {
		if cloudServices != nil {
			properties := make(map[string]interface{})
			properties["api_g_a_t_e_w_a_y"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.APIGATEWAY})
			properties["api_m_a_n_a_g_e_m_e_n_t"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.APIMANAGEMENT})
			properties["a_p_p_l_i_c_a_t_i_o_n_e_l_b"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.APPLICATIONELB})
			properties["a_p_p_l_i_c_a_t_i_o_n_g_a_t_e_w_a_y"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.APPLICATIONGATEWAY})
			properties["a_p_p_l_i_c_a_t_i_o_n_i_n_s_i_g_h_t_s"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.APPLICATIONINSIGHTS})
			properties["a_p_p_s_e_r_v_i_c_e"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.APPSERVICE})
			properties["a_p_p_s_e_r_v_i_c_e_p_l_a_n"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.APPSERVICEPLAN})
			properties["a_p_p_s_t_r_e_a_m"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.APPSTREAM})
			properties["a_t_h_e_n_a"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.ATHENA})
			properties["a_u_t_o_m_a_t_i_o_n_a_c_c_o_u_n_t"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.AUTOMATIONACCOUNT})
			properties["a_u_t_o_s_c_a_l_i_n_g"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.AUTOSCALING})
			properties["b_a_c_k_u_p_p_r_o_t_e_c_t_e_d_i_t_e_m_s"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.BACKUPPROTECTEDITEMS})
			properties["b_l_o_b_s_t_o_r_a_g_e"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.BLOBSTORAGE})
			properties["c_l_o_u_d_f_r_o_n_t"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.CLOUDFRONT})
			properties["c_l_o_u_d_s_e_a_r_c_h"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.CLOUDSEARCH})
			properties["c_o_d_e_b_ui_l_d"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.CODEBUILD})
			properties["c_o_g_n_i_t_i_v_e_s_e_a_r_c_h"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.COGNITIVESEARCH})
			properties["c_o_g_n_i_t_i_v_e_s_e_r_v_i_c_e_s"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.COGNITIVESERVICES})
			properties["c_o_g_n_i_t_o"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.COGNITO})
			properties["c_o_s_m_o_s_d_b"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.COSMOSDB})
			properties["d_a_t_a_f_a_c_t_o_r_y"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.DATAFACTORY})
			properties["d_i_r_e_c_t_c_o_n_n_e_c_t"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.DIRECTCONNECT})
			properties["d_m_s_r_e_p_l_i_c_a_t_i_o_n"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.DMSREPLICATION})
			properties["d_m_s_r_e_p_l_i_c_a_t_i_o_n_t_a_s_k_s"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.DMSREPLICATIONTASKS})
			properties["d_o_c_d_b"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.DOCDB})
			properties["d_y_n_a_m_o_d_b"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.DYNAMODB})
			properties["e_b_s"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.EBS})
			properties["e_c2"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.EC2})
			properties["e_c_s"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.ECS})
			properties["e_f_s"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.EFS})
			properties["e_l_a_s_t_i_c_a_c_h_e"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.ELASTICACHE})
			properties["e_l_a_s_t_i_c_b_e_a_n_s_t_a_l_k"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.ELASTICBEANSTALK})
			properties["e_l_a_s_t_i_c_s_e_a_r_c_h"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.ELASTICSEARCH})
			properties["e_l_a_s_t_i_c_t_r_a_n_s_c_o_d_e_r"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.ELASTICTRANSCODER})
			properties["e_l_b"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.ELB})
			properties["e_m_r"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.EMR})
			properties["e_v_e_n_t_b_r_id_g_e"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.EVENTBRIDGE})
			properties["e_v_e_n_t_h_u_b"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.EVENTHUB})
			properties["e_x_p_r_e_s_s_r_o_u_t_e_c_i_r_c_ui_t"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.EXPRESSROUTECIRCUIT})
			properties["f_i_l_e_s_t_o_r_a_g_e"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.FILESTORAGE})
			properties["f_i_r_e_h_o_s_e"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.FIREHOSE})
			properties["f_i_r_e_w_a_l_l"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.FIREWALL})
			properties["f_r_o_n_t_d_o_o_r_s"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.FRONTDOORS})
			properties["f_s_x"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.FSX})
			properties["g_l_u_e"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.GLUE})
			properties["k_e_y_v_a_u_l_t"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.KEYVAULT})
			properties["k_i_n_e_s_i_s"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.KINESIS})
			properties["k_i_n_e_s_i_s_v_id_e_o"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.KINESISVIDEO})
			properties["l_a_m_b_d_a"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.LAMBDA})
			properties["l_o_a_d_b_a_l_a_n_c_e_r_s"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.LOADBALANCERS})
			properties["l_o_g_a_n_a_l_y_t_i_c_s_w_o_r_k_s_p_a_c_e_s"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.LOGANALYTICSWORKSPACES})
			properties["l_o_g_i_c_a_p_p_s"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.LOGICAPPS})
			properties["m_e_d_i_a_c_o_n_n_e_c_t"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.MEDIACONNECT})
			properties["m_e_d_i_a_c_o_n_v_e_r_t"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.MEDIACONVERT})
			properties["m_e_d_i_a_p_a_c_k_a_g_e_l_i_v_e"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.MEDIAPACKAGELIVE})
			properties["m_e_d_i_a_p_a_c_k_a_g_e_v_o_d"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.MEDIAPACKAGEVOD})
			properties["m_e_d_i_a_s_t_o_r_e"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.MEDIASTORE})
			properties["m_e_d_i_a_t_a_i_l_o_r"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.MEDIATAILOR})
			properties["m_q"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.MQ})
			properties["m_s_k_b_r_o_k_e_r"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.MSKBROKER})
			properties["m_s_k_c_l_u_s_t_e_r"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.MSKCLUSTER})
			properties["m_y_sql"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.MYSQL})
			properties["n_a_t_g_a_t_e_w_a_y"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.NATGATEWAY})
			properties["n_e_t_w_o_r_k_e_l_b"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.NETWORKELB})
			properties["n_e_t_w_o_r_k_i_n_t_e_r_f_a_c_e"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.NETWORKINTERFACE})
			properties["o_p_s_w_o_r_k_s"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.OPSWORKS})
			properties["p_o_s_t_g_r_e_sql"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.POSTGRESQL})
			properties["p_u_b_l_i_c_ip"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.PUBLICIP})
			properties["q_u_e_u_e_s_t_o_r_a_g_e"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.QUEUESTORAGE})
			properties["r_d_s"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.RDS})
			properties["r_e_c_o_v_e_r_y_p_r_o_t_e_c_t_e_d_i_t_e_m"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.RECOVERYPROTECTEDITEM})
			properties["r_e_c_o_v_e_r_y_s_e_r_v_i_c_e_s"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.RECOVERYSERVICES})
			properties["r_e_d_i_s_c_a_c_h_e"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.REDISCACHE})
			properties["r_e_d_s_h_i_f_t"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.REDSHIFT})
			properties["r_o_u_t_e53"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.ROUTE53})
			properties["r_o_u_t_e53_r_e_s_o_l_v_e_r"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.ROUTE53RESOLVER})
			properties["s3"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.S3})
			properties["s_a_g_e_m_a_k_e_r"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.SAGEMAKER})
			properties["s_e_r_v_i_c_e_b_u_s"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.SERVICEBUS})
			properties["s_e_s"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.SES})
			properties["s_n_s"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.SNS})
			properties["sql_d_a_t_a_b_a_s_e"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.SQLDATABASE})
			properties["sql_e_l_a_s_t_i_c_p_o_o_l"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.SQLELASTICPOOL})
			properties["sql_m_a_n_a_g_e_d_i_n_s_t_a_n_c_e"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.SQLMANAGEDINSTANCE})
			properties["s_q_s"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.SQS})
			properties["s_t_e_p_f_u_n_c_t_i_o_n_s"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.STEPFUNCTIONS})
			properties["s_t_o_r_a_g_e_a_c_c_o_u_n_t"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.STORAGEACCOUNT})
			properties["s_w_f_a_c_t_i_v_i_t_y"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.SWFACTIVITY})
			properties["s_w_f_w_o_r_k_f_l_o_w"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.SWFWORKFLOW})
			properties["s_y_n_a_p_s_e_w_o_r_k_s_p_a_c_e_s"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.SYNAPSEWORKSPACES})
			properties["t_a_b_l_e_s_t_o_r_a_g_e"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.TABLESTORAGE})
			properties["t_r_a_f_f_i_c_m_a_n_a_g_e_r"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.TRAFFICMANAGER})
			properties["t_r_a_n_s_i_t_g_a_t_e_w_a_y"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.TRANSITGATEWAY})
			properties["v_i_r_t_u_a_l_d_e_s_k_t_o_p"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.VIRTUALDESKTOP})
			properties["v_i_r_t_u_a_l_m_a_c_h_i_n_e"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.VIRTUALMACHINE})
			properties["v_i_r_t_u_a_l_m_a_c_h_i_n_e_s_c_a_l_e_s_e_t"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.VIRTUALMACHINESCALESET})
			properties["v_i_r_t_u_a_l_m_a_c_h_i_n_e_s_c_a_l_e_s_e_t_vm"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.VIRTUALMACHINESCALESETVM})
			properties["v_i_r_t_u_a_l_n_e_t_w_o_r_k_g_a_t_e_w_a_y"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.VIRTUALNETWORKGATEWAY})
			properties["v_p_n"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.VPN})
			properties["w_o_r_k_s_p_a_c_e"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.WORKSPACE})
			properties["w_o_r_k_s_p_a_c_e_d_i_r_e_c_t_o_r_y"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.WORKSPACEDIRECTORY})
			d = append(d, &properties)
		}
	}
	return
}

func CloudServicesModel(d map[string]interface{}) *models.CloudServices {
	// assume that the incoming map only contains the relevant resource data
	var aPIGATEWAY *models.CloudServiceSettings = nil
	APIGATEWAYList := d["api_g_a_t_e_w_a_y"].([]interface{})
	if len(APIGATEWAYList) > 0 { // len(nil) = 0
		aPIGATEWAY = CloudServiceSettingsModel(APIGATEWAYList[0].(map[string]interface{}))
	}
	var aPIMANAGEMENT *models.CloudServiceSettings = nil
	APIMANAGEMENTList := d["api_m_a_n_a_g_e_m_e_n_t"].([]interface{})
	if len(APIMANAGEMENTList) > 0 { // len(nil) = 0
		aPIMANAGEMENT = CloudServiceSettingsModel(APIMANAGEMENTList[0].(map[string]interface{}))
	}
	var aPPLICATIONELB *models.CloudServiceSettings = nil
	APPLICATIONELBList := d["a_p_p_l_i_c_a_t_i_o_n_e_l_b"].([]interface{})
	if len(APPLICATIONELBList) > 0 { // len(nil) = 0
		aPPLICATIONELB = CloudServiceSettingsModel(APPLICATIONELBList[0].(map[string]interface{}))
	}
	var aPPLICATIONGATEWAY *models.CloudServiceSettings = nil
	APPLICATIONGATEWAYList := d["a_p_p_l_i_c_a_t_i_o_n_g_a_t_e_w_a_y"].([]interface{})
	if len(APPLICATIONGATEWAYList) > 0 { // len(nil) = 0
		aPPLICATIONGATEWAY = CloudServiceSettingsModel(APPLICATIONGATEWAYList[0].(map[string]interface{}))
	}
	var aPPLICATIONINSIGHTS *models.CloudServiceSettings = nil
	APPLICATIONINSIGHTSList := d["a_p_p_l_i_c_a_t_i_o_n_i_n_s_i_g_h_t_s"].([]interface{})
	if len(APPLICATIONINSIGHTSList) > 0 { // len(nil) = 0
		aPPLICATIONINSIGHTS = CloudServiceSettingsModel(APPLICATIONINSIGHTSList[0].(map[string]interface{}))
	}
	var aPPSERVICE *models.CloudServiceSettings = nil
	APPSERVICEList := d["a_p_p_s_e_r_v_i_c_e"].([]interface{})
	if len(APPSERVICEList) > 0 { // len(nil) = 0
		aPPSERVICE = CloudServiceSettingsModel(APPSERVICEList[0].(map[string]interface{}))
	}
	var aPPSERVICEPLAN *models.CloudServiceSettings = nil
	APPSERVICEPLANList := d["a_p_p_s_e_r_v_i_c_e_p_l_a_n"].([]interface{})
	if len(APPSERVICEPLANList) > 0 { // len(nil) = 0
		aPPSERVICEPLAN = CloudServiceSettingsModel(APPSERVICEPLANList[0].(map[string]interface{}))
	}
	var aPPSTREAM *models.CloudServiceSettings = nil
	APPSTREAMList := d["a_p_p_s_t_r_e_a_m"].([]interface{})
	if len(APPSTREAMList) > 0 { // len(nil) = 0
		aPPSTREAM = CloudServiceSettingsModel(APPSTREAMList[0].(map[string]interface{}))
	}
	var aTHENA *models.CloudServiceSettings = nil
	ATHENAList := d["a_t_h_e_n_a"].([]interface{})
	if len(ATHENAList) > 0 { // len(nil) = 0
		aTHENA = CloudServiceSettingsModel(ATHENAList[0].(map[string]interface{}))
	}
	var aUTOMATIONACCOUNT *models.CloudServiceSettings = nil
	AUTOMATIONACCOUNTList := d["a_u_t_o_m_a_t_i_o_n_a_c_c_o_u_n_t"].([]interface{})
	if len(AUTOMATIONACCOUNTList) > 0 { // len(nil) = 0
		aUTOMATIONACCOUNT = CloudServiceSettingsModel(AUTOMATIONACCOUNTList[0].(map[string]interface{}))
	}
	var aUTOSCALING *models.CloudServiceSettings = nil
	AUTOSCALINGList := d["a_u_t_o_s_c_a_l_i_n_g"].([]interface{})
	if len(AUTOSCALINGList) > 0 { // len(nil) = 0
		aUTOSCALING = CloudServiceSettingsModel(AUTOSCALINGList[0].(map[string]interface{}))
	}
	var bACKUPPROTECTEDITEMS *models.CloudServiceSettings = nil
	BACKUPPROTECTEDITEMSList := d["b_a_c_k_u_p_p_r_o_t_e_c_t_e_d_i_t_e_m_s"].([]interface{})
	if len(BACKUPPROTECTEDITEMSList) > 0 { // len(nil) = 0
		bACKUPPROTECTEDITEMS = CloudServiceSettingsModel(BACKUPPROTECTEDITEMSList[0].(map[string]interface{}))
	}
	var bLOBSTORAGE *models.CloudServiceSettings = nil
	BLOBSTORAGEList := d["b_l_o_b_s_t_o_r_a_g_e"].([]interface{})
	if len(BLOBSTORAGEList) > 0 { // len(nil) = 0
		bLOBSTORAGE = CloudServiceSettingsModel(BLOBSTORAGEList[0].(map[string]interface{}))
	}
	var cLOUDFRONT *models.CloudServiceSettings = nil
	CLOUDFRONTList := d["c_l_o_u_d_f_r_o_n_t"].([]interface{})
	if len(CLOUDFRONTList) > 0 { // len(nil) = 0
		cLOUDFRONT = CloudServiceSettingsModel(CLOUDFRONTList[0].(map[string]interface{}))
	}
	var cLOUDSEARCH *models.CloudServiceSettings = nil
	CLOUDSEARCHList := d["c_l_o_u_d_s_e_a_r_c_h"].([]interface{})
	if len(CLOUDSEARCHList) > 0 { // len(nil) = 0
		cLOUDSEARCH = CloudServiceSettingsModel(CLOUDSEARCHList[0].(map[string]interface{}))
	}
	var cODEBUILD *models.CloudServiceSettings = nil
	CODEBUILDList := d["c_o_d_e_b_ui_l_d"].([]interface{})
	if len(CODEBUILDList) > 0 { // len(nil) = 0
		cODEBUILD = CloudServiceSettingsModel(CODEBUILDList[0].(map[string]interface{}))
	}
	var cOGNITIVESEARCH *models.CloudServiceSettings = nil
	COGNITIVESEARCHList := d["c_o_g_n_i_t_i_v_e_s_e_a_r_c_h"].([]interface{})
	if len(COGNITIVESEARCHList) > 0 { // len(nil) = 0
		cOGNITIVESEARCH = CloudServiceSettingsModel(COGNITIVESEARCHList[0].(map[string]interface{}))
	}
	var cOGNITIVESERVICES *models.CloudServiceSettings = nil
	COGNITIVESERVICESList := d["c_o_g_n_i_t_i_v_e_s_e_r_v_i_c_e_s"].([]interface{})
	if len(COGNITIVESERVICESList) > 0 { // len(nil) = 0
		cOGNITIVESERVICES = CloudServiceSettingsModel(COGNITIVESERVICESList[0].(map[string]interface{}))
	}
	var cOGNITO *models.CloudServiceSettings = nil
	COGNITOList := d["c_o_g_n_i_t_o"].([]interface{})
	if len(COGNITOList) > 0 { // len(nil) = 0
		cOGNITO = CloudServiceSettingsModel(COGNITOList[0].(map[string]interface{}))
	}
	var cOSMOSDB *models.CloudServiceSettings = nil
	COSMOSDBList := d["c_o_s_m_o_s_d_b"].([]interface{})
	if len(COSMOSDBList) > 0 { // len(nil) = 0
		cOSMOSDB = CloudServiceSettingsModel(COSMOSDBList[0].(map[string]interface{}))
	}
	var dATAFACTORY *models.CloudServiceSettings = nil
	DATAFACTORYList := d["d_a_t_a_f_a_c_t_o_r_y"].([]interface{})
	if len(DATAFACTORYList) > 0 { // len(nil) = 0
		dATAFACTORY = CloudServiceSettingsModel(DATAFACTORYList[0].(map[string]interface{}))
	}
	var dIRECTCONNECT *models.CloudServiceSettings = nil
	DIRECTCONNECTList := d["d_i_r_e_c_t_c_o_n_n_e_c_t"].([]interface{})
	if len(DIRECTCONNECTList) > 0 { // len(nil) = 0
		dIRECTCONNECT = CloudServiceSettingsModel(DIRECTCONNECTList[0].(map[string]interface{}))
	}
	var dMSREPLICATION *models.CloudServiceSettings = nil
	DMSREPLICATIONList := d["d_m_s_r_e_p_l_i_c_a_t_i_o_n"].([]interface{})
	if len(DMSREPLICATIONList) > 0 { // len(nil) = 0
		dMSREPLICATION = CloudServiceSettingsModel(DMSREPLICATIONList[0].(map[string]interface{}))
	}
	var dMSREPLICATIONTASKS *models.CloudServiceSettings = nil
	DMSREPLICATIONTASKSList := d["d_m_s_r_e_p_l_i_c_a_t_i_o_n_t_a_s_k_s"].([]interface{})
	if len(DMSREPLICATIONTASKSList) > 0 { // len(nil) = 0
		dMSREPLICATIONTASKS = CloudServiceSettingsModel(DMSREPLICATIONTASKSList[0].(map[string]interface{}))
	}
	var dOCDB *models.CloudServiceSettings = nil
	DOCDBList := d["d_o_c_d_b"].([]interface{})
	if len(DOCDBList) > 0 { // len(nil) = 0
		dOCDB = CloudServiceSettingsModel(DOCDBList[0].(map[string]interface{}))
	}
	var dYNAMODB *models.CloudServiceSettings = nil
	DYNAMODBList := d["d_y_n_a_m_o_d_b"].([]interface{})
	if len(DYNAMODBList) > 0 { // len(nil) = 0
		dYNAMODB = CloudServiceSettingsModel(DYNAMODBList[0].(map[string]interface{}))
	}
	var eBS *models.CloudServiceSettings = nil
	EBSList := d["e_b_s"].([]interface{})
	if len(EBSList) > 0 { // len(nil) = 0
		eBS = CloudServiceSettingsModel(EBSList[0].(map[string]interface{}))
	}
	var eC2 *models.CloudServiceSettings = nil
	EC2List := d["e_c2"].([]interface{})
	if len(EC2List) > 0 { // len(nil) = 0
		eC2 = CloudServiceSettingsModel(EC2List[0].(map[string]interface{}))
	}
	var eCS *models.CloudServiceSettings = nil
	ECSList := d["e_c_s"].([]interface{})
	if len(ECSList) > 0 { // len(nil) = 0
		eCS = CloudServiceSettingsModel(ECSList[0].(map[string]interface{}))
	}
	var eFS *models.CloudServiceSettings = nil
	EFSList := d["e_f_s"].([]interface{})
	if len(EFSList) > 0 { // len(nil) = 0
		eFS = CloudServiceSettingsModel(EFSList[0].(map[string]interface{}))
	}
	var eLASTICACHE *models.CloudServiceSettings = nil
	ELASTICACHEList := d["e_l_a_s_t_i_c_a_c_h_e"].([]interface{})
	if len(ELASTICACHEList) > 0 { // len(nil) = 0
		eLASTICACHE = CloudServiceSettingsModel(ELASTICACHEList[0].(map[string]interface{}))
	}
	var eLASTICBEANSTALK *models.CloudServiceSettings = nil
	ELASTICBEANSTALKList := d["e_l_a_s_t_i_c_b_e_a_n_s_t_a_l_k"].([]interface{})
	if len(ELASTICBEANSTALKList) > 0 { // len(nil) = 0
		eLASTICBEANSTALK = CloudServiceSettingsModel(ELASTICBEANSTALKList[0].(map[string]interface{}))
	}
	var eLASTICSEARCH *models.CloudServiceSettings = nil
	ELASTICSEARCHList := d["e_l_a_s_t_i_c_s_e_a_r_c_h"].([]interface{})
	if len(ELASTICSEARCHList) > 0 { // len(nil) = 0
		eLASTICSEARCH = CloudServiceSettingsModel(ELASTICSEARCHList[0].(map[string]interface{}))
	}
	var eLASTICTRANSCODER *models.CloudServiceSettings = nil
	ELASTICTRANSCODERList := d["e_l_a_s_t_i_c_t_r_a_n_s_c_o_d_e_r"].([]interface{})
	if len(ELASTICTRANSCODERList) > 0 { // len(nil) = 0
		eLASTICTRANSCODER = CloudServiceSettingsModel(ELASTICTRANSCODERList[0].(map[string]interface{}))
	}
	var eLB *models.CloudServiceSettings = nil
	ELBList := d["e_l_b"].([]interface{})
	if len(ELBList) > 0 { // len(nil) = 0
		eLB = CloudServiceSettingsModel(ELBList[0].(map[string]interface{}))
	}
	var eMR *models.CloudServiceSettings = nil
	EMRList := d["e_m_r"].([]interface{})
	if len(EMRList) > 0 { // len(nil) = 0
		eMR = CloudServiceSettingsModel(EMRList[0].(map[string]interface{}))
	}
	var eVENTBRIDGE *models.CloudServiceSettings = nil
	EVENTBRIDGEList := d["e_v_e_n_t_b_r_id_g_e"].([]interface{})
	if len(EVENTBRIDGEList) > 0 { // len(nil) = 0
		eVENTBRIDGE = CloudServiceSettingsModel(EVENTBRIDGEList[0].(map[string]interface{}))
	}
	var eVENTHUB *models.CloudServiceSettings = nil
	EVENTHUBList := d["e_v_e_n_t_h_u_b"].([]interface{})
	if len(EVENTHUBList) > 0 { // len(nil) = 0
		eVENTHUB = CloudServiceSettingsModel(EVENTHUBList[0].(map[string]interface{}))
	}
	var eXPRESSROUTECIRCUIT *models.CloudServiceSettings = nil
	EXPRESSROUTECIRCUITList := d["e_x_p_r_e_s_s_r_o_u_t_e_c_i_r_c_ui_t"].([]interface{})
	if len(EXPRESSROUTECIRCUITList) > 0 { // len(nil) = 0
		eXPRESSROUTECIRCUIT = CloudServiceSettingsModel(EXPRESSROUTECIRCUITList[0].(map[string]interface{}))
	}
	var fILESTORAGE *models.CloudServiceSettings = nil
	FILESTORAGEList := d["f_i_l_e_s_t_o_r_a_g_e"].([]interface{})
	if len(FILESTORAGEList) > 0 { // len(nil) = 0
		fILESTORAGE = CloudServiceSettingsModel(FILESTORAGEList[0].(map[string]interface{}))
	}
	var fIREHOSE *models.CloudServiceSettings = nil
	FIREHOSEList := d["f_i_r_e_h_o_s_e"].([]interface{})
	if len(FIREHOSEList) > 0 { // len(nil) = 0
		fIREHOSE = CloudServiceSettingsModel(FIREHOSEList[0].(map[string]interface{}))
	}
	var fIREWALL *models.CloudServiceSettings = nil
	FIREWALLList := d["f_i_r_e_w_a_l_l"].([]interface{})
	if len(FIREWALLList) > 0 { // len(nil) = 0
		fIREWALL = CloudServiceSettingsModel(FIREWALLList[0].(map[string]interface{}))
	}
	var fRONTDOORS *models.CloudServiceSettings = nil
	FRONTDOORSList := d["f_r_o_n_t_d_o_o_r_s"].([]interface{})
	if len(FRONTDOORSList) > 0 { // len(nil) = 0
		fRONTDOORS = CloudServiceSettingsModel(FRONTDOORSList[0].(map[string]interface{}))
	}
	var fSX *models.CloudServiceSettings = nil
	FSXList := d["f_s_x"].([]interface{})
	if len(FSXList) > 0 { // len(nil) = 0
		fSX = CloudServiceSettingsModel(FSXList[0].(map[string]interface{}))
	}
	var gLUE *models.CloudServiceSettings = nil
	GLUEList := d["g_l_u_e"].([]interface{})
	if len(GLUEList) > 0 { // len(nil) = 0
		gLUE = CloudServiceSettingsModel(GLUEList[0].(map[string]interface{}))
	}
	var kEYVAULT *models.CloudServiceSettings = nil
	KEYVAULTList := d["k_e_y_v_a_u_l_t"].([]interface{})
	if len(KEYVAULTList) > 0 { // len(nil) = 0
		kEYVAULT = CloudServiceSettingsModel(KEYVAULTList[0].(map[string]interface{}))
	}
	var kINESIS *models.CloudServiceSettings = nil
	KINESISList := d["k_i_n_e_s_i_s"].([]interface{})
	if len(KINESISList) > 0 { // len(nil) = 0
		kINESIS = CloudServiceSettingsModel(KINESISList[0].(map[string]interface{}))
	}
	var kINESISVIDEO *models.CloudServiceSettings = nil
	KINESISVIDEOList := d["k_i_n_e_s_i_s_v_id_e_o"].([]interface{})
	if len(KINESISVIDEOList) > 0 { // len(nil) = 0
		kINESISVIDEO = CloudServiceSettingsModel(KINESISVIDEOList[0].(map[string]interface{}))
	}
	var lAMBDA *models.CloudServiceSettings = nil
	LAMBDAList := d["l_a_m_b_d_a"].([]interface{})
	if len(LAMBDAList) > 0 { // len(nil) = 0
		lAMBDA = CloudServiceSettingsModel(LAMBDAList[0].(map[string]interface{}))
	}
	var lOADBALANCERS *models.CloudServiceSettings = nil
	LOADBALANCERSList := d["l_o_a_d_b_a_l_a_n_c_e_r_s"].([]interface{})
	if len(LOADBALANCERSList) > 0 { // len(nil) = 0
		lOADBALANCERS = CloudServiceSettingsModel(LOADBALANCERSList[0].(map[string]interface{}))
	}
	var lOGANALYTICSWORKSPACES *models.CloudServiceSettings = nil
	LOGANALYTICSWORKSPACESList := d["l_o_g_a_n_a_l_y_t_i_c_s_w_o_r_k_s_p_a_c_e_s"].([]interface{})
	if len(LOGANALYTICSWORKSPACESList) > 0 { // len(nil) = 0
		lOGANALYTICSWORKSPACES = CloudServiceSettingsModel(LOGANALYTICSWORKSPACESList[0].(map[string]interface{}))
	}
	var lOGICAPPS *models.CloudServiceSettings = nil
	LOGICAPPSList := d["l_o_g_i_c_a_p_p_s"].([]interface{})
	if len(LOGICAPPSList) > 0 { // len(nil) = 0
		lOGICAPPS = CloudServiceSettingsModel(LOGICAPPSList[0].(map[string]interface{}))
	}
	var mEDIACONNECT *models.CloudServiceSettings = nil
	MEDIACONNECTList := d["m_e_d_i_a_c_o_n_n_e_c_t"].([]interface{})
	if len(MEDIACONNECTList) > 0 { // len(nil) = 0
		mEDIACONNECT = CloudServiceSettingsModel(MEDIACONNECTList[0].(map[string]interface{}))
	}
	var mEDIACONVERT *models.CloudServiceSettings = nil
	MEDIACONVERTList := d["m_e_d_i_a_c_o_n_v_e_r_t"].([]interface{})
	if len(MEDIACONVERTList) > 0 { // len(nil) = 0
		mEDIACONVERT = CloudServiceSettingsModel(MEDIACONVERTList[0].(map[string]interface{}))
	}
	var mEDIAPACKAGELIVE *models.CloudServiceSettings = nil
	MEDIAPACKAGELIVEList := d["m_e_d_i_a_p_a_c_k_a_g_e_l_i_v_e"].([]interface{})
	if len(MEDIAPACKAGELIVEList) > 0 { // len(nil) = 0
		mEDIAPACKAGELIVE = CloudServiceSettingsModel(MEDIAPACKAGELIVEList[0].(map[string]interface{}))
	}
	var mEDIAPACKAGEVOD *models.CloudServiceSettings = nil
	MEDIAPACKAGEVODList := d["m_e_d_i_a_p_a_c_k_a_g_e_v_o_d"].([]interface{})
	if len(MEDIAPACKAGEVODList) > 0 { // len(nil) = 0
		mEDIAPACKAGEVOD = CloudServiceSettingsModel(MEDIAPACKAGEVODList[0].(map[string]interface{}))
	}
	var mEDIASTORE *models.CloudServiceSettings = nil
	MEDIASTOREList := d["m_e_d_i_a_s_t_o_r_e"].([]interface{})
	if len(MEDIASTOREList) > 0 { // len(nil) = 0
		mEDIASTORE = CloudServiceSettingsModel(MEDIASTOREList[0].(map[string]interface{}))
	}
	var mEDIATAILOR *models.CloudServiceSettings = nil
	MEDIATAILORList := d["m_e_d_i_a_t_a_i_l_o_r"].([]interface{})
	if len(MEDIATAILORList) > 0 { // len(nil) = 0
		mEDIATAILOR = CloudServiceSettingsModel(MEDIATAILORList[0].(map[string]interface{}))
	}
	var mQ *models.CloudServiceSettings = nil
	MQList := d["m_q"].([]interface{})
	if len(MQList) > 0 { // len(nil) = 0
		mQ = CloudServiceSettingsModel(MQList[0].(map[string]interface{}))
	}
	var mSKBROKER *models.CloudServiceSettings = nil
	MSKBROKERList := d["m_s_k_b_r_o_k_e_r"].([]interface{})
	if len(MSKBROKERList) > 0 { // len(nil) = 0
		mSKBROKER = CloudServiceSettingsModel(MSKBROKERList[0].(map[string]interface{}))
	}
	var mSKCLUSTER *models.CloudServiceSettings = nil
	MSKCLUSTERList := d["m_s_k_c_l_u_s_t_e_r"].([]interface{})
	if len(MSKCLUSTERList) > 0 { // len(nil) = 0
		mSKCLUSTER = CloudServiceSettingsModel(MSKCLUSTERList[0].(map[string]interface{}))
	}
	var mYSQL *models.CloudServiceSettings = nil
	MYSQLList := d["m_y_sql"].([]interface{})
	if len(MYSQLList) > 0 { // len(nil) = 0
		mYSQL = CloudServiceSettingsModel(MYSQLList[0].(map[string]interface{}))
	}
	var nATGATEWAY *models.CloudServiceSettings = nil
	NATGATEWAYList := d["n_a_t_g_a_t_e_w_a_y"].([]interface{})
	if len(NATGATEWAYList) > 0 { // len(nil) = 0
		nATGATEWAY = CloudServiceSettingsModel(NATGATEWAYList[0].(map[string]interface{}))
	}
	var nETWORKELB *models.CloudServiceSettings = nil
	NETWORKELBList := d["n_e_t_w_o_r_k_e_l_b"].([]interface{})
	if len(NETWORKELBList) > 0 { // len(nil) = 0
		nETWORKELB = CloudServiceSettingsModel(NETWORKELBList[0].(map[string]interface{}))
	}
	var nETWORKINTERFACE *models.CloudServiceSettings = nil
	NETWORKINTERFACEList := d["n_e_t_w_o_r_k_i_n_t_e_r_f_a_c_e"].([]interface{})
	if len(NETWORKINTERFACEList) > 0 { // len(nil) = 0
		nETWORKINTERFACE = CloudServiceSettingsModel(NETWORKINTERFACEList[0].(map[string]interface{}))
	}
	var oPSWORKS *models.CloudServiceSettings = nil
	OPSWORKSList := d["o_p_s_w_o_r_k_s"].([]interface{})
	if len(OPSWORKSList) > 0 { // len(nil) = 0
		oPSWORKS = CloudServiceSettingsModel(OPSWORKSList[0].(map[string]interface{}))
	}
	var pOSTGRESQL *models.CloudServiceSettings = nil
	POSTGRESQLList := d["p_o_s_t_g_r_e_sql"].([]interface{})
	if len(POSTGRESQLList) > 0 { // len(nil) = 0
		pOSTGRESQL = CloudServiceSettingsModel(POSTGRESQLList[0].(map[string]interface{}))
	}
	var pUBLICIP *models.CloudServiceSettings = nil
	PUBLICIPList := d["p_u_b_l_i_c_ip"].([]interface{})
	if len(PUBLICIPList) > 0 { // len(nil) = 0
		pUBLICIP = CloudServiceSettingsModel(PUBLICIPList[0].(map[string]interface{}))
	}
	var qUEUESTORAGE *models.CloudServiceSettings = nil
	QUEUESTORAGEList := d["q_u_e_u_e_s_t_o_r_a_g_e"].([]interface{})
	if len(QUEUESTORAGEList) > 0 { // len(nil) = 0
		qUEUESTORAGE = CloudServiceSettingsModel(QUEUESTORAGEList[0].(map[string]interface{}))
	}
	var rDS *models.CloudServiceSettings = nil
	RDSList := d["r_d_s"].([]interface{})
	if len(RDSList) > 0 { // len(nil) = 0
		rDS = CloudServiceSettingsModel(RDSList[0].(map[string]interface{}))
	}
	var rECOVERYPROTECTEDITEM *models.CloudServiceSettings = nil
	RECOVERYPROTECTEDITEMList := d["r_e_c_o_v_e_r_y_p_r_o_t_e_c_t_e_d_i_t_e_m"].([]interface{})
	if len(RECOVERYPROTECTEDITEMList) > 0 { // len(nil) = 0
		rECOVERYPROTECTEDITEM = CloudServiceSettingsModel(RECOVERYPROTECTEDITEMList[0].(map[string]interface{}))
	}
	var rECOVERYSERVICES *models.CloudServiceSettings = nil
	RECOVERYSERVICESList := d["r_e_c_o_v_e_r_y_s_e_r_v_i_c_e_s"].([]interface{})
	if len(RECOVERYSERVICESList) > 0 { // len(nil) = 0
		rECOVERYSERVICES = CloudServiceSettingsModel(RECOVERYSERVICESList[0].(map[string]interface{}))
	}
	var rEDISCACHE *models.CloudServiceSettings = nil
	REDISCACHEList := d["r_e_d_i_s_c_a_c_h_e"].([]interface{})
	if len(REDISCACHEList) > 0 { // len(nil) = 0
		rEDISCACHE = CloudServiceSettingsModel(REDISCACHEList[0].(map[string]interface{}))
	}
	var rEDSHIFT *models.CloudServiceSettings = nil
	REDSHIFTList := d["r_e_d_s_h_i_f_t"].([]interface{})
	if len(REDSHIFTList) > 0 { // len(nil) = 0
		rEDSHIFT = CloudServiceSettingsModel(REDSHIFTList[0].(map[string]interface{}))
	}
	var rOUTE53 *models.CloudServiceSettings = nil
	ROUTE53List := d["r_o_u_t_e53"].([]interface{})
	if len(ROUTE53List) > 0 { // len(nil) = 0
		rOUTE53 = CloudServiceSettingsModel(ROUTE53List[0].(map[string]interface{}))
	}
	var rOUTE53RESOLVER *models.CloudServiceSettings = nil
	ROUTE53RESOLVERList := d["r_o_u_t_e53_r_e_s_o_l_v_e_r"].([]interface{})
	if len(ROUTE53RESOLVERList) > 0 { // len(nil) = 0
		rOUTE53RESOLVER = CloudServiceSettingsModel(ROUTE53RESOLVERList[0].(map[string]interface{}))
	}
	var s3 *models.CloudServiceSettings = nil
	S3List := d["s3"].([]interface{})
	if len(S3List) > 0 { // len(nil) = 0
		s3 = CloudServiceSettingsModel(S3List[0].(map[string]interface{}))
	}
	var sAGEMAKER *models.CloudServiceSettings = nil
	SAGEMAKERList := d["s_a_g_e_m_a_k_e_r"].([]interface{})
	if len(SAGEMAKERList) > 0 { // len(nil) = 0
		sAGEMAKER = CloudServiceSettingsModel(SAGEMAKERList[0].(map[string]interface{}))
	}
	var sERVICEBUS *models.CloudServiceSettings = nil
	SERVICEBUSList := d["s_e_r_v_i_c_e_b_u_s"].([]interface{})
	if len(SERVICEBUSList) > 0 { // len(nil) = 0
		sERVICEBUS = CloudServiceSettingsModel(SERVICEBUSList[0].(map[string]interface{}))
	}
	var sES *models.CloudServiceSettings = nil
	SESList := d["s_e_s"].([]interface{})
	if len(SESList) > 0 { // len(nil) = 0
		sES = CloudServiceSettingsModel(SESList[0].(map[string]interface{}))
	}
	var sNS *models.CloudServiceSettings = nil
	SNSList := d["s_n_s"].([]interface{})
	if len(SNSList) > 0 { // len(nil) = 0
		sNS = CloudServiceSettingsModel(SNSList[0].(map[string]interface{}))
	}
	var sQLDATABASE *models.CloudServiceSettings = nil
	SQLDATABASEList := d["sql_d_a_t_a_b_a_s_e"].([]interface{})
	if len(SQLDATABASEList) > 0 { // len(nil) = 0
		sQLDATABASE = CloudServiceSettingsModel(SQLDATABASEList[0].(map[string]interface{}))
	}
	var sQLELASTICPOOL *models.CloudServiceSettings = nil
	SQLELASTICPOOLList := d["sql_e_l_a_s_t_i_c_p_o_o_l"].([]interface{})
	if len(SQLELASTICPOOLList) > 0 { // len(nil) = 0
		sQLELASTICPOOL = CloudServiceSettingsModel(SQLELASTICPOOLList[0].(map[string]interface{}))
	}
	var sQLMANAGEDINSTANCE *models.CloudServiceSettings = nil
	SQLMANAGEDINSTANCEList := d["sql_m_a_n_a_g_e_d_i_n_s_t_a_n_c_e"].([]interface{})
	if len(SQLMANAGEDINSTANCEList) > 0 { // len(nil) = 0
		sQLMANAGEDINSTANCE = CloudServiceSettingsModel(SQLMANAGEDINSTANCEList[0].(map[string]interface{}))
	}
	var sQS *models.CloudServiceSettings = nil
	SQSList := d["s_q_s"].([]interface{})
	if len(SQSList) > 0 { // len(nil) = 0
		sQS = CloudServiceSettingsModel(SQSList[0].(map[string]interface{}))
	}
	var sTEPFUNCTIONS *models.CloudServiceSettings = nil
	STEPFUNCTIONSList := d["s_t_e_p_f_u_n_c_t_i_o_n_s"].([]interface{})
	if len(STEPFUNCTIONSList) > 0 { // len(nil) = 0
		sTEPFUNCTIONS = CloudServiceSettingsModel(STEPFUNCTIONSList[0].(map[string]interface{}))
	}
	var sTORAGEACCOUNT *models.CloudServiceSettings = nil
	STORAGEACCOUNTList := d["s_t_o_r_a_g_e_a_c_c_o_u_n_t"].([]interface{})
	if len(STORAGEACCOUNTList) > 0 { // len(nil) = 0
		sTORAGEACCOUNT = CloudServiceSettingsModel(STORAGEACCOUNTList[0].(map[string]interface{}))
	}
	var sWFACTIVITY *models.CloudServiceSettings = nil
	SWFACTIVITYList := d["s_w_f_a_c_t_i_v_i_t_y"].([]interface{})
	if len(SWFACTIVITYList) > 0 { // len(nil) = 0
		sWFACTIVITY = CloudServiceSettingsModel(SWFACTIVITYList[0].(map[string]interface{}))
	}
	var sWFWORKFLOW *models.CloudServiceSettings = nil
	SWFWORKFLOWList := d["s_w_f_w_o_r_k_f_l_o_w"].([]interface{})
	if len(SWFWORKFLOWList) > 0 { // len(nil) = 0
		sWFWORKFLOW = CloudServiceSettingsModel(SWFWORKFLOWList[0].(map[string]interface{}))
	}
	var sYNAPSEWORKSPACES *models.CloudServiceSettings = nil
	SYNAPSEWORKSPACESList := d["s_y_n_a_p_s_e_w_o_r_k_s_p_a_c_e_s"].([]interface{})
	if len(SYNAPSEWORKSPACESList) > 0 { // len(nil) = 0
		sYNAPSEWORKSPACES = CloudServiceSettingsModel(SYNAPSEWORKSPACESList[0].(map[string]interface{}))
	}
	var tABLESTORAGE *models.CloudServiceSettings = nil
	TABLESTORAGEList := d["t_a_b_l_e_s_t_o_r_a_g_e"].([]interface{})
	if len(TABLESTORAGEList) > 0 { // len(nil) = 0
		tABLESTORAGE = CloudServiceSettingsModel(TABLESTORAGEList[0].(map[string]interface{}))
	}
	var tRAFFICMANAGER *models.CloudServiceSettings = nil
	TRAFFICMANAGERList := d["t_r_a_f_f_i_c_m_a_n_a_g_e_r"].([]interface{})
	if len(TRAFFICMANAGERList) > 0 { // len(nil) = 0
		tRAFFICMANAGER = CloudServiceSettingsModel(TRAFFICMANAGERList[0].(map[string]interface{}))
	}
	var tRANSITGATEWAY *models.CloudServiceSettings = nil
	TRANSITGATEWAYList := d["t_r_a_n_s_i_t_g_a_t_e_w_a_y"].([]interface{})
	if len(TRANSITGATEWAYList) > 0 { // len(nil) = 0
		tRANSITGATEWAY = CloudServiceSettingsModel(TRANSITGATEWAYList[0].(map[string]interface{}))
	}
	var vIRTUALDESKTOP *models.CloudServiceSettings = nil
	VIRTUALDESKTOPList := d["v_i_r_t_u_a_l_d_e_s_k_t_o_p"].([]interface{})
	if len(VIRTUALDESKTOPList) > 0 { // len(nil) = 0
		vIRTUALDESKTOP = CloudServiceSettingsModel(VIRTUALDESKTOPList[0].(map[string]interface{}))
	}
	var vIRTUALMACHINE *models.CloudServiceSettings = nil
	VIRTUALMACHINEList := d["v_i_r_t_u_a_l_m_a_c_h_i_n_e"].([]interface{})
	if len(VIRTUALMACHINEList) > 0 { // len(nil) = 0
		vIRTUALMACHINE = CloudServiceSettingsModel(VIRTUALMACHINEList[0].(map[string]interface{}))
	}
	var vIRTUALMACHINESCALESET *models.CloudServiceSettings = nil
	VIRTUALMACHINESCALESETList := d["v_i_r_t_u_a_l_m_a_c_h_i_n_e_s_c_a_l_e_s_e_t"].([]interface{})
	if len(VIRTUALMACHINESCALESETList) > 0 { // len(nil) = 0
		vIRTUALMACHINESCALESET = CloudServiceSettingsModel(VIRTUALMACHINESCALESETList[0].(map[string]interface{}))
	}
	var vIRTUALMACHINESCALESETVM *models.CloudServiceSettings = nil
	VIRTUALMACHINESCALESETVMList := d["v_i_r_t_u_a_l_m_a_c_h_i_n_e_s_c_a_l_e_s_e_t_vm"].([]interface{})
	if len(VIRTUALMACHINESCALESETVMList) > 0 { // len(nil) = 0
		vIRTUALMACHINESCALESETVM = CloudServiceSettingsModel(VIRTUALMACHINESCALESETVMList[0].(map[string]interface{}))
	}
	var vIRTUALNETWORKGATEWAY *models.CloudServiceSettings = nil
	VIRTUALNETWORKGATEWAYList := d["v_i_r_t_u_a_l_n_e_t_w_o_r_k_g_a_t_e_w_a_y"].([]interface{})
	if len(VIRTUALNETWORKGATEWAYList) > 0 { // len(nil) = 0
		vIRTUALNETWORKGATEWAY = CloudServiceSettingsModel(VIRTUALNETWORKGATEWAYList[0].(map[string]interface{}))
	}
	var vPN *models.CloudServiceSettings = nil
	VPNList := d["v_p_n"].([]interface{})
	if len(VPNList) > 0 { // len(nil) = 0
		vPN = CloudServiceSettingsModel(VPNList[0].(map[string]interface{}))
	}
	var wORKSPACE *models.CloudServiceSettings = nil
	WORKSPACEList := d["w_o_r_k_s_p_a_c_e"].([]interface{})
	if len(WORKSPACEList) > 0 { // len(nil) = 0
		wORKSPACE = CloudServiceSettingsModel(WORKSPACEList[0].(map[string]interface{}))
	}
	var wORKSPACEDIRECTORY *models.CloudServiceSettings = nil
	WORKSPACEDIRECTORYList := d["w_o_r_k_s_p_a_c_e_d_i_r_e_c_t_o_r_y"].([]interface{})
	if len(WORKSPACEDIRECTORYList) > 0 { // len(nil) = 0
		wORKSPACEDIRECTORY = CloudServiceSettingsModel(WORKSPACEDIRECTORYList[0].(map[string]interface{}))
	}
	
	return &models.CloudServices {
		APIGATEWAY: aPIGATEWAY,
		APIMANAGEMENT: aPIMANAGEMENT,
		APPLICATIONELB: aPPLICATIONELB,
		APPLICATIONGATEWAY: aPPLICATIONGATEWAY,
		APPLICATIONINSIGHTS: aPPLICATIONINSIGHTS,
		APPSERVICE: aPPSERVICE,
		APPSERVICEPLAN: aPPSERVICEPLAN,
		APPSTREAM: aPPSTREAM,
		ATHENA: aTHENA,
		AUTOMATIONACCOUNT: aUTOMATIONACCOUNT,
		AUTOSCALING: aUTOSCALING,
		BACKUPPROTECTEDITEMS: bACKUPPROTECTEDITEMS,
		BLOBSTORAGE: bLOBSTORAGE,
		CLOUDFRONT: cLOUDFRONT,
		CLOUDSEARCH: cLOUDSEARCH,
		CODEBUILD: cODEBUILD,
		COGNITIVESEARCH: cOGNITIVESEARCH,
		COGNITIVESERVICES: cOGNITIVESERVICES,
		COGNITO: cOGNITO,
		COSMOSDB: cOSMOSDB,
		DATAFACTORY: dATAFACTORY,
		DIRECTCONNECT: dIRECTCONNECT,
		DMSREPLICATION: dMSREPLICATION,
		DMSREPLICATIONTASKS: dMSREPLICATIONTASKS,
		DOCDB: dOCDB,
		DYNAMODB: dYNAMODB,
		EBS: eBS,
		EC2: eC2,
		ECS: eCS,
		EFS: eFS,
		ELASTICACHE: eLASTICACHE,
		ELASTICBEANSTALK: eLASTICBEANSTALK,
		ELASTICSEARCH: eLASTICSEARCH,
		ELASTICTRANSCODER: eLASTICTRANSCODER,
		ELB: eLB,
		EMR: eMR,
		EVENTBRIDGE: eVENTBRIDGE,
		EVENTHUB: eVENTHUB,
		EXPRESSROUTECIRCUIT: eXPRESSROUTECIRCUIT,
		FILESTORAGE: fILESTORAGE,
		FIREHOSE: fIREHOSE,
		FIREWALL: fIREWALL,
		FRONTDOORS: fRONTDOORS,
		FSX: fSX,
		GLUE: gLUE,
		KEYVAULT: kEYVAULT,
		KINESIS: kINESIS,
		KINESISVIDEO: kINESISVIDEO,
		LAMBDA: lAMBDA,
		LOADBALANCERS: lOADBALANCERS,
		LOGANALYTICSWORKSPACES: lOGANALYTICSWORKSPACES,
		LOGICAPPS: lOGICAPPS,
		MEDIACONNECT: mEDIACONNECT,
		MEDIACONVERT: mEDIACONVERT,
		MEDIAPACKAGELIVE: mEDIAPACKAGELIVE,
		MEDIAPACKAGEVOD: mEDIAPACKAGEVOD,
		MEDIASTORE: mEDIASTORE,
		MEDIATAILOR: mEDIATAILOR,
		MQ: mQ,
		MSKBROKER: mSKBROKER,
		MSKCLUSTER: mSKCLUSTER,
		MYSQL: mYSQL,
		NATGATEWAY: nATGATEWAY,
		NETWORKELB: nETWORKELB,
		NETWORKINTERFACE: nETWORKINTERFACE,
		OPSWORKS: oPSWORKS,
		POSTGRESQL: pOSTGRESQL,
		PUBLICIP: pUBLICIP,
		QUEUESTORAGE: qUEUESTORAGE,
		RDS: rDS,
		RECOVERYPROTECTEDITEM: rECOVERYPROTECTEDITEM,
		RECOVERYSERVICES: rECOVERYSERVICES,
		REDISCACHE: rEDISCACHE,
		REDSHIFT: rEDSHIFT,
		ROUTE53: rOUTE53,
		ROUTE53RESOLVER: rOUTE53RESOLVER,
		S3: s3,
		SAGEMAKER: sAGEMAKER,
		SERVICEBUS: sERVICEBUS,
		SES: sES,
		SNS: sNS,
		SQLDATABASE: sQLDATABASE,
		SQLELASTICPOOL: sQLELASTICPOOL,
		SQLMANAGEDINSTANCE: sQLMANAGEDINSTANCE,
		SQS: sQS,
		STEPFUNCTIONS: sTEPFUNCTIONS,
		STORAGEACCOUNT: sTORAGEACCOUNT,
		SWFACTIVITY: sWFACTIVITY,
		SWFWORKFLOW: sWFWORKFLOW,
		SYNAPSEWORKSPACES: sYNAPSEWORKSPACES,
		TABLESTORAGE: tABLESTORAGE,
		TRAFFICMANAGER: tRAFFICMANAGER,
		TRANSITGATEWAY: tRANSITGATEWAY,
		VIRTUALDESKTOP: vIRTUALDESKTOP,
		VIRTUALMACHINE: vIRTUALMACHINE,
		VIRTUALMACHINESCALESET: vIRTUALMACHINESCALESET,
		VIRTUALMACHINESCALESETVM: vIRTUALMACHINESCALESETVM,
		VIRTUALNETWORKGATEWAY: vIRTUALNETWORKGATEWAY,
		VPN: vPN,
		WORKSPACE: wORKSPACE,
		WORKSPACEDIRECTORY: wORKSPACEDIRECTORY,
	}
}

func GetCloudServicesPropertyFields() (t []string) {
	return []string{
		"api_g_a_t_e_w_a_y",
		"api_m_a_n_a_g_e_m_e_n_t",
		"a_p_p_l_i_c_a_t_i_o_n_e_l_b",
		"a_p_p_l_i_c_a_t_i_o_n_g_a_t_e_w_a_y",
		"a_p_p_l_i_c_a_t_i_o_n_i_n_s_i_g_h_t_s",
		"a_p_p_s_e_r_v_i_c_e",
		"a_p_p_s_e_r_v_i_c_e_p_l_a_n",
		"a_p_p_s_t_r_e_a_m",
		"a_t_h_e_n_a",
		"a_u_t_o_m_a_t_i_o_n_a_c_c_o_u_n_t",
		"a_u_t_o_s_c_a_l_i_n_g",
		"b_a_c_k_u_p_p_r_o_t_e_c_t_e_d_i_t_e_m_s",
		"b_l_o_b_s_t_o_r_a_g_e",
		"c_l_o_u_d_f_r_o_n_t",
		"c_l_o_u_d_s_e_a_r_c_h",
		"c_o_d_e_b_ui_l_d",
		"c_o_g_n_i_t_i_v_e_s_e_a_r_c_h",
		"c_o_g_n_i_t_i_v_e_s_e_r_v_i_c_e_s",
		"c_o_g_n_i_t_o",
		"c_o_s_m_o_s_d_b",
		"d_a_t_a_f_a_c_t_o_r_y",
		"d_i_r_e_c_t_c_o_n_n_e_c_t",
		"d_m_s_r_e_p_l_i_c_a_t_i_o_n",
		"d_m_s_r_e_p_l_i_c_a_t_i_o_n_t_a_s_k_s",
		"d_o_c_d_b",
		"d_y_n_a_m_o_d_b",
		"e_b_s",
		"e_c2",
		"e_c_s",
		"e_f_s",
		"e_l_a_s_t_i_c_a_c_h_e",
		"e_l_a_s_t_i_c_b_e_a_n_s_t_a_l_k",
		"e_l_a_s_t_i_c_s_e_a_r_c_h",
		"e_l_a_s_t_i_c_t_r_a_n_s_c_o_d_e_r",
		"e_l_b",
		"e_m_r",
		"e_v_e_n_t_b_r_id_g_e",
		"e_v_e_n_t_h_u_b",
		"e_x_p_r_e_s_s_r_o_u_t_e_c_i_r_c_ui_t",
		"f_i_l_e_s_t_o_r_a_g_e",
		"f_i_r_e_h_o_s_e",
		"f_i_r_e_w_a_l_l",
		"f_r_o_n_t_d_o_o_r_s",
		"f_s_x",
		"g_l_u_e",
		"k_e_y_v_a_u_l_t",
		"k_i_n_e_s_i_s",
		"k_i_n_e_s_i_s_v_id_e_o",
		"l_a_m_b_d_a",
		"l_o_a_d_b_a_l_a_n_c_e_r_s",
		"l_o_g_a_n_a_l_y_t_i_c_s_w_o_r_k_s_p_a_c_e_s",
		"l_o_g_i_c_a_p_p_s",
		"m_e_d_i_a_c_o_n_n_e_c_t",
		"m_e_d_i_a_c_o_n_v_e_r_t",
		"m_e_d_i_a_p_a_c_k_a_g_e_l_i_v_e",
		"m_e_d_i_a_p_a_c_k_a_g_e_v_o_d",
		"m_e_d_i_a_s_t_o_r_e",
		"m_e_d_i_a_t_a_i_l_o_r",
		"m_q",
		"m_s_k_b_r_o_k_e_r",
		"m_s_k_c_l_u_s_t_e_r",
		"m_y_sql",
		"n_a_t_g_a_t_e_w_a_y",
		"n_e_t_w_o_r_k_e_l_b",
		"n_e_t_w_o_r_k_i_n_t_e_r_f_a_c_e",
		"o_p_s_w_o_r_k_s",
		"p_o_s_t_g_r_e_sql",
		"p_u_b_l_i_c_ip",
		"q_u_e_u_e_s_t_o_r_a_g_e",
		"r_d_s",
		"r_e_c_o_v_e_r_y_p_r_o_t_e_c_t_e_d_i_t_e_m",
		"r_e_c_o_v_e_r_y_s_e_r_v_i_c_e_s",
		"r_e_d_i_s_c_a_c_h_e",
		"r_e_d_s_h_i_f_t",
		"r_o_u_t_e53",
		"r_o_u_t_e53_r_e_s_o_l_v_e_r",
		"s3",
		"s_a_g_e_m_a_k_e_r",
		"s_e_r_v_i_c_e_b_u_s",
		"s_e_s",
		"s_n_s",
		"sql_d_a_t_a_b_a_s_e",
		"sql_e_l_a_s_t_i_c_p_o_o_l",
		"sql_m_a_n_a_g_e_d_i_n_s_t_a_n_c_e",
		"s_q_s",
		"s_t_e_p_f_u_n_c_t_i_o_n_s",
		"s_t_o_r_a_g_e_a_c_c_o_u_n_t",
		"s_w_f_a_c_t_i_v_i_t_y",
		"s_w_f_w_o_r_k_f_l_o_w",
		"s_y_n_a_p_s_e_w_o_r_k_s_p_a_c_e_s",
		"t_a_b_l_e_s_t_o_r_a_g_e",
		"t_r_a_f_f_i_c_m_a_n_a_g_e_r",
		"t_r_a_n_s_i_t_g_a_t_e_w_a_y",
		"v_i_r_t_u_a_l_d_e_s_k_t_o_p",
		"v_i_r_t_u_a_l_m_a_c_h_i_n_e",
		"v_i_r_t_u_a_l_m_a_c_h_i_n_e_s_c_a_l_e_s_e_t",
		"v_i_r_t_u_a_l_m_a_c_h_i_n_e_s_c_a_l_e_s_e_t_vm",
		"v_i_r_t_u_a_l_n_e_t_w_o_r_k_g_a_t_e_w_a_y",
		"v_p_n",
		"w_o_r_k_s_p_a_c_e",
		"w_o_r_k_s_p_a_c_e_d_i_r_e_c_t_o_r_y",
	}
}