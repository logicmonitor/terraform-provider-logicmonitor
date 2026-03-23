package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func UptimeWebCheckStepSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"http_body": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"http_headers": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"http_method": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"http_version": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"auth": {
			Type: schema.TypeList, //GoType: Authentication 
            Elem: &schema.Resource{
				Schema: AuthenticationSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
		
		"description": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"enable": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"follow_redirection": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"fullpage_load": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"invert_match": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"keyword": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"label": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"match_type": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"name": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"path": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"post_data_edit_type": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"req_script": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"req_type": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"require_auth": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"resp_script": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"resp_type": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"schema": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"status_code": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"timeout": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"type": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"url": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"use_default_root": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
	}
}

func SetUptimeWebCheckStepSubResourceData(m []*models.UptimeWebCheckStep) (d []*map[string]interface{}) {
	for _, uptimeWebCheckStep := range m {
		if uptimeWebCheckStep != nil {
			properties := make(map[string]interface{})
			properties["http_body"] = uptimeWebCheckStep.HTTPBody
			properties["http_headers"] = uptimeWebCheckStep.HTTPHeaders
			properties["http_method"] = uptimeWebCheckStep.HTTPMethod
			properties["http_version"] = uptimeWebCheckStep.HTTPVersion
			properties["auth"] = SetAuthenticationSubResourceData([]*models.Authentication{uptimeWebCheckStep.Auth})
			properties["description"] = uptimeWebCheckStep.Description
			properties["enable"] = uptimeWebCheckStep.Enable
			properties["follow_redirection"] = uptimeWebCheckStep.FollowRedirection
			properties["fullpage_load"] = uptimeWebCheckStep.FullpageLoad
			properties["invert_match"] = uptimeWebCheckStep.InvertMatch
			properties["keyword"] = uptimeWebCheckStep.Keyword
			properties["label"] = uptimeWebCheckStep.Label
			properties["match_type"] = uptimeWebCheckStep.MatchType
			properties["name"] = uptimeWebCheckStep.Name
			properties["path"] = uptimeWebCheckStep.Path
			properties["post_data_edit_type"] = uptimeWebCheckStep.PostDataEditType
			properties["req_script"] = uptimeWebCheckStep.ReqScript
			properties["req_type"] = uptimeWebCheckStep.ReqType
			properties["require_auth"] = uptimeWebCheckStep.RequireAuth
			properties["resp_script"] = uptimeWebCheckStep.RespScript
			properties["resp_type"] = uptimeWebCheckStep.RespType
			properties["schema"] = uptimeWebCheckStep.Schema
			properties["status_code"] = uptimeWebCheckStep.StatusCode
			properties["timeout"] = uptimeWebCheckStep.Timeout
			properties["type"] = uptimeWebCheckStep.Type
			properties["url"] = uptimeWebCheckStep.URL
			properties["use_default_root"] = uptimeWebCheckStep.UseDefaultRoot
			d = append(d, &properties)
		}
	}
	return
}

func UptimeWebCheckStepModel(d map[string]interface{}) *models.UptimeWebCheckStep {
	// assume that the incoming map only contains the relevant resource data
	hTTPBody := d["http_body"].(string)
	hTTPHeaders := d["http_headers"].(string)
	hTTPMethod := d["http_method"].(string)
	hTTPVersion := d["http_version"].(string)
	var auth *models.Authentication = nil
	authList := d["auth"].([]interface{})
	if len(authList) > 0 { // len(nil) = 0
		auth = AuthenticationModel(authList[0].(map[string]interface{}))
	}
	description := d["description"].(string)
	enable := d["enable"].(bool)
	followRedirection := d["follow_redirection"].(bool)
	fullpageLoad := d["fullpage_load"].(bool)
	invertMatch := d["invert_match"].(bool)
	keyword := d["keyword"].(string)
	label := d["label"].(string)
	matchType := d["match_type"].(string)
	name := d["name"].(string)
	path := d["path"].(string)
	postDataEditType := d["post_data_edit_type"].(string)
	reqScript := d["req_script"].(string)
	reqType := d["req_type"].(string)
	requireAuth := d["require_auth"].(bool)
	respScript := d["resp_script"].(string)
	respType := d["resp_type"].(string)
	schema := d["schema"].(string)
	statusCode := d["status_code"].(string)
	timeout := int32(d["timeout"].(int))
	typeVar := d["type"].(string)
	url := d["url"].(string)
	useDefaultRoot := d["use_default_root"].(bool)
	
	return &models.UptimeWebCheckStep {
		HTTPBody: hTTPBody,
		HTTPHeaders: hTTPHeaders,
		HTTPMethod: hTTPMethod,
		HTTPVersion: hTTPVersion,
		Auth: auth,
		Description: description,
		Enable: enable,
		FollowRedirection: followRedirection,
		FullpageLoad: fullpageLoad,
		InvertMatch: invertMatch,
		Keyword: keyword,
		Label: label,
		MatchType: matchType,
		Name: name,
		Path: path,
		PostDataEditType: postDataEditType,
		ReqScript: reqScript,
		ReqType: reqType,
		RequireAuth: requireAuth,
		RespScript: respScript,
		RespType: respType,
		Schema: schema,
		StatusCode: statusCode,
		Timeout: timeout,
		Type: typeVar,
		URL: url,
		UseDefaultRoot: useDefaultRoot,
	}
}

func GetUptimeWebCheckStepPropertyFields() (t []string) {
	return []string{
		"http_body",
		"http_headers",
		"http_method",
		"http_version",
		"auth",
		"description",
		"enable",
		"follow_redirection",
		"fullpage_load",
		"invert_match",
		"keyword",
		"label",
		"match_type",
		"name",
		"path",
		"post_data_edit_type",
		"req_script",
		"req_type",
		"require_auth",
		"resp_script",
		"resp_type",
		"schema",
		"status_code",
		"timeout",
		"type",
		"url",
		"use_default_root",
	}
}