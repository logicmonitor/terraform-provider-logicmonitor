package schemata

import (
	"strconv"
	"time"
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func AwsExternalIDSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"created_at": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"external_id": {
			Type: schema.TypeString,
			Optional: true,
		},
		
	}
}

func SetAwsExternalIDResourceData(d *schema.ResourceData, m *models.AwsExternalID) {
	d.Set("created_at", m.CreatedAt)
	d.Set("external_id", m.ExternalID)
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
}

func SetAwsExternalIDSubResourceData(m []*models.AwsExternalID) (d []*map[string]interface{}) {
	for _, awsExternalId := range m {
		if awsExternalId != nil {
			properties := make(map[string]interface{})
			properties["created_at"] = awsExternalId.CreatedAt
			properties["external_id"] = awsExternalId.ExternalID
			d = append(d, &properties)
		}
	}
	return
}

func AwsExternalIDModel(d map[string]interface{}) *models.AwsExternalID {
	// assume that the incoming map only contains the relevant resource data
	createdAt := d["created_at"].(string)
	externalID := d["external_id"].(string)
	
	return &models.AwsExternalID {
		CreatedAt: createdAt,
		ExternalID: externalID,
	}
}

func GetAwsExternalIDPropertyFields() (t []string) {
	return []string{
		"created_at",
		"external_id",
	}
}