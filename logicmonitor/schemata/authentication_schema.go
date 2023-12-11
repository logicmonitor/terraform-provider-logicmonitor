package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func AuthenticationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"domain": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"password": {
			Type: schema.TypeString,
			Sensitive: true,
			Required: true,
		},
		
		"type": {
			Type: schema.TypeString,
			Required: true,
		},
		
		"user_name": {
			Type: schema.TypeString,
			Required: true,
		},
		
	}
}

func SetAuthenticationSubResourceData(m []*models.Authentication) (d []*map[string]interface{}) {
	for _, authentication := range m {
		if authentication != nil {
			properties := make(map[string]interface{})
			properties["domain"] = authentication.Domain
			properties["password"] = authentication.Password
			properties["type"] = authentication.Type
			properties["user_name"] = authentication.UserName
			d = append(d, &properties)
		}
	}
	return
}

func AuthenticationModel(d map[string]interface{}) *models.Authentication {
	// assume that the incoming map only contains the relevant resource data
	domain := d["domain"].(string)
	password := d["password"].(string)
	typeVar := d["type"].(string)
	userName := d["user_name"].(string)
	
	return &models.Authentication {
		Domain: domain,
		Password: &password,
		Type: &typeVar,
		UserName: &userName,
	}
}

func GetAuthenticationPropertyFields() (t []string) {
	return []string{
		"domain",
		"password",
		"type",
		"user_name",
	}
}