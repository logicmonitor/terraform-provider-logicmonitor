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
		
		"a_p_p_l_i_c_a_t_i_o_n_e_l_b": {
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
		
		"a_u_t_o_s_c_a_l_i_n_g": {
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
		
		"c_o_g_n_i_t_o": {
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
		
		"f_i_r_e_h_o_s_e": {
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
		
		"o_p_s_w_o_r_k_s": {
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
		
		"t_r_a_n_s_i_t_g_a_t_e_w_a_y": {
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
			properties["a_p_p_l_i_c_a_t_i_o_n_e_l_b"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.APPLICATIONELB})
			properties["a_p_p_s_t_r_e_a_m"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.APPSTREAM})
			properties["a_t_h_e_n_a"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.ATHENA})
			properties["a_u_t_o_s_c_a_l_i_n_g"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.AUTOSCALING})
			properties["c_l_o_u_d_f_r_o_n_t"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.CLOUDFRONT})
			properties["c_l_o_u_d_s_e_a_r_c_h"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.CLOUDSEARCH})
			properties["c_o_d_e_b_ui_l_d"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.CODEBUILD})
			properties["c_o_g_n_i_t_o"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.COGNITO})
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
			properties["f_i_r_e_h_o_s_e"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.FIREHOSE})
			properties["f_s_x"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.FSX})
			properties["g_l_u_e"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.GLUE})
			properties["k_i_n_e_s_i_s"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.KINESIS})
			properties["k_i_n_e_s_i_s_v_id_e_o"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.KINESISVIDEO})
			properties["l_a_m_b_d_a"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.LAMBDA})
			properties["m_e_d_i_a_c_o_n_n_e_c_t"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.MEDIACONNECT})
			properties["m_e_d_i_a_c_o_n_v_e_r_t"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.MEDIACONVERT})
			properties["m_e_d_i_a_p_a_c_k_a_g_e_l_i_v_e"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.MEDIAPACKAGELIVE})
			properties["m_e_d_i_a_p_a_c_k_a_g_e_v_o_d"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.MEDIAPACKAGEVOD})
			properties["m_e_d_i_a_s_t_o_r_e"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.MEDIASTORE})
			properties["m_e_d_i_a_t_a_i_l_o_r"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.MEDIATAILOR})
			properties["m_q"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.MQ})
			properties["m_s_k_b_r_o_k_e_r"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.MSKBROKER})
			properties["m_s_k_c_l_u_s_t_e_r"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.MSKCLUSTER})
			properties["n_a_t_g_a_t_e_w_a_y"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.NATGATEWAY})
			properties["n_e_t_w_o_r_k_e_l_b"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.NETWORKELB})
			properties["o_p_s_w_o_r_k_s"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.OPSWORKS})
			properties["r_d_s"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.RDS})
			properties["r_e_d_s_h_i_f_t"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.REDSHIFT})
			properties["r_o_u_t_e53"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.ROUTE53})
			properties["r_o_u_t_e53_r_e_s_o_l_v_e_r"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.ROUTE53RESOLVER})
			properties["s3"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.S3})
			properties["s_a_g_e_m_a_k_e_r"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.SAGEMAKER})
			properties["s_e_s"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.SES})
			properties["s_n_s"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.SNS})
			properties["s_q_s"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.SQS})
			properties["s_t_e_p_f_u_n_c_t_i_o_n_s"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.STEPFUNCTIONS})
			properties["s_w_f_a_c_t_i_v_i_t_y"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.SWFACTIVITY})
			properties["s_w_f_w_o_r_k_f_l_o_w"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.SWFWORKFLOW})
			properties["t_r_a_n_s_i_t_g_a_t_e_w_a_y"] = SetCloudServiceSettingsSubResourceData([]*models.CloudServiceSettings{cloudServices.TRANSITGATEWAY})
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
	var aPPLICATIONELB *models.CloudServiceSettings = nil
	APPLICATIONELBList := d["a_p_p_l_i_c_a_t_i_o_n_e_l_b"].([]interface{})
	if len(APPLICATIONELBList) > 0 { // len(nil) = 0
		aPPLICATIONELB = CloudServiceSettingsModel(APPLICATIONELBList[0].(map[string]interface{}))
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
	var aUTOSCALING *models.CloudServiceSettings = nil
	AUTOSCALINGList := d["a_u_t_o_s_c_a_l_i_n_g"].([]interface{})
	if len(AUTOSCALINGList) > 0 { // len(nil) = 0
		aUTOSCALING = CloudServiceSettingsModel(AUTOSCALINGList[0].(map[string]interface{}))
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
	var cOGNITO *models.CloudServiceSettings = nil
	COGNITOList := d["c_o_g_n_i_t_o"].([]interface{})
	if len(COGNITOList) > 0 { // len(nil) = 0
		cOGNITO = CloudServiceSettingsModel(COGNITOList[0].(map[string]interface{}))
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
	var fIREHOSE *models.CloudServiceSettings = nil
	FIREHOSEList := d["f_i_r_e_h_o_s_e"].([]interface{})
	if len(FIREHOSEList) > 0 { // len(nil) = 0
		fIREHOSE = CloudServiceSettingsModel(FIREHOSEList[0].(map[string]interface{}))
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
	var oPSWORKS *models.CloudServiceSettings = nil
	OPSWORKSList := d["o_p_s_w_o_r_k_s"].([]interface{})
	if len(OPSWORKSList) > 0 { // len(nil) = 0
		oPSWORKS = CloudServiceSettingsModel(OPSWORKSList[0].(map[string]interface{}))
	}
	var rDS *models.CloudServiceSettings = nil
	RDSList := d["r_d_s"].([]interface{})
	if len(RDSList) > 0 { // len(nil) = 0
		rDS = CloudServiceSettingsModel(RDSList[0].(map[string]interface{}))
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
	var tRANSITGATEWAY *models.CloudServiceSettings = nil
	TRANSITGATEWAYList := d["t_r_a_n_s_i_t_g_a_t_e_w_a_y"].([]interface{})
	if len(TRANSITGATEWAYList) > 0 { // len(nil) = 0
		tRANSITGATEWAY = CloudServiceSettingsModel(TRANSITGATEWAYList[0].(map[string]interface{}))
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
		APPLICATIONELB: aPPLICATIONELB,
		APPSTREAM: aPPSTREAM,
		ATHENA: aTHENA,
		AUTOSCALING: aUTOSCALING,
		CLOUDFRONT: cLOUDFRONT,
		CLOUDSEARCH: cLOUDSEARCH,
		CODEBUILD: cODEBUILD,
		COGNITO: cOGNITO,
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
		FIREHOSE: fIREHOSE,
		FSX: fSX,
		GLUE: gLUE,
		KINESIS: kINESIS,
		KINESISVIDEO: kINESISVIDEO,
		LAMBDA: lAMBDA,
		MEDIACONNECT: mEDIACONNECT,
		MEDIACONVERT: mEDIACONVERT,
		MEDIAPACKAGELIVE: mEDIAPACKAGELIVE,
		MEDIAPACKAGEVOD: mEDIAPACKAGEVOD,
		MEDIASTORE: mEDIASTORE,
		MEDIATAILOR: mEDIATAILOR,
		MQ: mQ,
		MSKBROKER: mSKBROKER,
		MSKCLUSTER: mSKCLUSTER,
		NATGATEWAY: nATGATEWAY,
		NETWORKELB: nETWORKELB,
		OPSWORKS: oPSWORKS,
		RDS: rDS,
		REDSHIFT: rEDSHIFT,
		ROUTE53: rOUTE53,
		ROUTE53RESOLVER: rOUTE53RESOLVER,
		S3: s3,
		SAGEMAKER: sAGEMAKER,
		SES: sES,
		SNS: sNS,
		SQS: sQS,
		STEPFUNCTIONS: sTEPFUNCTIONS,
		SWFACTIVITY: sWFACTIVITY,
		SWFWORKFLOW: sWFWORKFLOW,
		TRANSITGATEWAY: tRANSITGATEWAY,
		VPN: vPN,
		WORKSPACE: wORKSPACE,
		WORKSPACEDIRECTORY: wORKSPACEDIRECTORY,
	}
}

func GetCloudServicesPropertyFields() (t []string) {
	return []string{
		"api_g_a_t_e_w_a_y",
		"a_p_p_l_i_c_a_t_i_o_n_e_l_b",
		"a_p_p_s_t_r_e_a_m",
		"a_t_h_e_n_a",
		"a_u_t_o_s_c_a_l_i_n_g",
		"c_l_o_u_d_f_r_o_n_t",
		"c_l_o_u_d_s_e_a_r_c_h",
		"c_o_d_e_b_ui_l_d",
		"c_o_g_n_i_t_o",
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
		"f_i_r_e_h_o_s_e",
		"f_s_x",
		"g_l_u_e",
		"k_i_n_e_s_i_s",
		"k_i_n_e_s_i_s_v_id_e_o",
		"l_a_m_b_d_a",
		"m_e_d_i_a_c_o_n_n_e_c_t",
		"m_e_d_i_a_c_o_n_v_e_r_t",
		"m_e_d_i_a_p_a_c_k_a_g_e_l_i_v_e",
		"m_e_d_i_a_p_a_c_k_a_g_e_v_o_d",
		"m_e_d_i_a_s_t_o_r_e",
		"m_e_d_i_a_t_a_i_l_o_r",
		"m_q",
		"m_s_k_b_r_o_k_e_r",
		"m_s_k_c_l_u_s_t_e_r",
		"n_a_t_g_a_t_e_w_a_y",
		"n_e_t_w_o_r_k_e_l_b",
		"o_p_s_w_o_r_k_s",
		"r_d_s",
		"r_e_d_s_h_i_f_t",
		"r_o_u_t_e53",
		"r_o_u_t_e53_r_e_s_o_l_v_e_r",
		"s3",
		"s_a_g_e_m_a_k_e_r",
		"s_e_s",
		"s_n_s",
		"s_q_s",
		"s_t_e_p_f_u_n_c_t_i_o_n_s",
		"s_w_f_a_c_t_i_v_i_t_y",
		"s_w_f_w_o_r_k_f_l_o_w",
		"t_r_a_n_s_i_t_g_a_t_e_w_a_y",
		"v_p_n",
		"w_o_r_k_s_p_a_c_e",
		"w_o_r_k_s_p_a_c_e_d_i_r_e_c_t_o_r_y",
	}
}