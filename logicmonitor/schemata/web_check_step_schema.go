package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-logicmonitor/models"
)

func WebCheckStepSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"http_body": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"http_headers": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"http_method": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"http_version": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"auth": {
			Type: schema.TypeList, //GoType: Authentication
			Elem: &schema.Resource{
				Schema: AuthenticationSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"enable": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"follow_redirection": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"fullpage_load": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"invert_match": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"keyword": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"label": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"match_type": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"path": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"post_data_edit_type": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"req_script": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"req_type": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"require_auth": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"resp_script": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"resp_type": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"schema": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"status_code": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"timeout": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"type": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"url": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"use_default_root": {
			Type:     schema.TypeBool,
			Required: true,
		},
	}
}

func SetWebCheckStepSubResourceData(m []*models.WebCheckStep) (d []*map[string]interface{}) {
	for _, webCheckStep := range m {
		if webCheckStep != nil {
			properties := make(map[string]interface{})
			properties["http_body"] = webCheckStep.HTTPBody
			properties["http_headers"] = webCheckStep.HTTPHeaders
			properties["http_method"] = webCheckStep.HTTPMethod
			properties["http_version"] = webCheckStep.HTTPVersion
			properties["auth"] = SetAuthenticationSubResourceData([]*models.Authentication{webCheckStep.Auth})
			properties["description"] = webCheckStep.Description
			properties["enable"] = webCheckStep.Enable
			properties["follow_redirection"] = webCheckStep.FollowRedirection
			properties["fullpage_load"] = webCheckStep.FullpageLoad
			properties["invert_match"] = webCheckStep.InvertMatch
			properties["keyword"] = webCheckStep.Keyword
			properties["label"] = webCheckStep.Label
			properties["match_type"] = webCheckStep.MatchType
			properties["name"] = webCheckStep.Name
			properties["path"] = webCheckStep.Path
			properties["post_data_edit_type"] = webCheckStep.PostDataEditType
			properties["req_script"] = webCheckStep.ReqScript
			properties["req_type"] = webCheckStep.ReqType
			properties["require_auth"] = webCheckStep.RequireAuth
			properties["resp_script"] = webCheckStep.RespScript
			properties["resp_type"] = webCheckStep.RespType
			properties["schema"] = webCheckStep.Schema
			properties["status_code"] = webCheckStep.StatusCode
			properties["timeout"] = webCheckStep.Timeout
			properties["type"] = webCheckStep.Type
			properties["url"] = webCheckStep.URL
			properties["use_default_root"] = webCheckStep.UseDefaultRoot
			d = append(d, &properties)
		}
	}
	return
}

func WebCheckStepModel(d map[string]interface{}) *models.WebCheckStep {
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

	return &models.WebCheckStep{
		HTTPBody:          hTTPBody,
		HTTPHeaders:       hTTPHeaders,
		HTTPMethod:        hTTPMethod,
		HTTPVersion:       hTTPVersion,
		Auth:              auth,
		Description:       description,
		Enable:            enable,
		FollowRedirection: followRedirection,
		FullpageLoad:      fullpageLoad,
		InvertMatch:       invertMatch,
		Keyword:           keyword,
		Label:             label,
		MatchType:         matchType,
		Name:              name,
		Path:              path,
		PostDataEditType:  postDataEditType,
		ReqScript:         reqScript,
		ReqType:           reqType,
		RequireAuth:       requireAuth,
		RespScript:        respScript,
		RespType:          respType,
		Schema:            schema,
		StatusCode:        statusCode,
		Timeout:           timeout,
		Type:              typeVar,
		URL:               url,
		UseDefaultRoot:    &useDefaultRoot,
	}
}

func GetWebCheckStepPropertyFields() (t []string) {
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
