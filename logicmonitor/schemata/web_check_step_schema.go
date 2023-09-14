package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func WebCheckStepSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"http_method": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"http_version": {
			Type: schema.TypeString,
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
		
		"match_type": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"name": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"post_data_edit_type": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"req_type": {
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
		
		"timeout": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"use_default_root": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
	}
}

func SetWebCheckStepSubResourceData(m []*models.WebCheckStep) (d []*map[string]interface{}) {
	for _, webCheckStep := range m {
		if webCheckStep != nil {
			properties := make(map[string]interface{})
			properties["http_method"] = webCheckStep.HTTPMethod
			properties["http_version"] = webCheckStep.HTTPVersion
			properties["description"] = webCheckStep.Description
			properties["enable"] = webCheckStep.Enable
			properties["follow_redirection"] = webCheckStep.FollowRedirection
			properties["fullpage_load"] = webCheckStep.FullpageLoad
			properties["match_type"] = webCheckStep.MatchType
			properties["name"] = webCheckStep.Name
			properties["post_data_edit_type"] = webCheckStep.PostDataEditType
			properties["req_type"] = webCheckStep.ReqType
			properties["resp_type"] = webCheckStep.RespType
			properties["schema"] = webCheckStep.Schema
			properties["timeout"] = webCheckStep.Timeout
			properties["use_default_root"] = webCheckStep.UseDefaultRoot
			d = append(d, &properties)
		}
	}
	return
}

func WebCheckStepModel(d map[string]interface{}) *models.WebCheckStep {
	// assume that the incoming map only contains the relevant resource data
	hTTPMethod := d["http_method"].(string)
	hTTPVersion := d["http_version"].(string)
	description := d["description"].(string)
	enable := d["enable"].(bool)
	followRedirection := d["follow_redirection"].(bool)
	fullpageLoad := d["fullpage_load"].(bool)
	matchType := d["match_type"].(string)
	name := d["name"].(string)
	postDataEditType := d["post_data_edit_type"].(string)
	reqType := d["req_type"].(string)
	respType := d["resp_type"].(string)
	schema := d["schema"].(string)
	timeout := int32(d["timeout"].(int))
	useDefaultRoot := d["use_default_root"].(bool)
	
	return &models.WebCheckStep {
		HTTPMethod: hTTPMethod,
		HTTPVersion: hTTPVersion,
		Description: description,
		Enable: enable,
		FollowRedirection: followRedirection,
		FullpageLoad: fullpageLoad,
		MatchType: matchType,
		Name: name,
		PostDataEditType: postDataEditType,
		ReqType: reqType,
		RespType: respType,
		Schema: schema,
		Timeout: timeout,
		UseDefaultRoot: useDefaultRoot,
	}
}

func GetWebCheckStepPropertyFields() (t []string) {
	return []string{
		"http_method",
		"http_version",
		"description",
		"enable",
		"follow_redirection",
		"fullpage_load",
		"match_type",
		"name",
		"post_data_edit_type",
		"req_type",
		"resp_type",
		"schema",
		"timeout",
		"use_default_root",
	}
}