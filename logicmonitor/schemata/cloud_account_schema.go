package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func CloudAccountSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"assumed_role_arn": {
			Type: schema.TypeString,
			Required: true,
		},
		
		"billing_bucket_name": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"billing_path_prefix": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"collector_id": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"external_id": {
			Type: schema.TypeString,
			Required: true,
		},
		
		"schedule": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"type": {
			Type: schema.TypeString,
			Computed: true,
		},
		
	}
}

func SetCloudAccountSubResourceData(m []*models.CloudAccount) (d []*map[string]interface{}) {
	for _, cloudAccount := range m {
		if cloudAccount != nil {
			properties := make(map[string]interface{})
			properties["account_id"] = cloudAccount.AccountID
			properties["assumed_role_arn"] = cloudAccount.AssumedRoleArn
			properties["billing_bucket_name"] = cloudAccount.BillingBucketName
			properties["billing_path_prefix"] = cloudAccount.BillingPathPrefix
			properties["collector_id"] = cloudAccount.CollectorID
			properties["external_id"] = cloudAccount.ExternalID
			properties["schedule"] = cloudAccount.Schedule
			properties["type"] = cloudAccount.Type
			d = append(d, &properties)
		}
	}
	return
}

func CloudAccountModel(d map[string]interface{}) *models.CloudAccount {
	// assume that the incoming map only contains the relevant resource data
	accountID := d["account_id"].(string)
	assumedRoleArn := d["assumed_role_arn"].(string)
	externalID := d["external_id"].(string)
	schedule := d["schedule"].(string)
	
	return &models.CloudAccount {
		AccountID: accountID,
		AssumedRoleArn: &assumedRoleArn,
		ExternalID: &externalID,
		Schedule: schedule,
	}
}

func GetCloudAccountPropertyFields() (t []string) {
	return []string{
		"account_id",
		"assumed_role_arn",
		"external_id",
		"schedule",
	}
}