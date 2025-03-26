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
			Optional: true,
		},
		
		"billing_bucket_name": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"billing_path_prefix": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"client_id": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"collector_description": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"collector_id": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"country": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"external_id": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"schedule": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"secret_key": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"subscription_ids": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"tenant_id": {
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
			properties["client_id"] = cloudAccount.ClientID
			properties["collector_description"] = cloudAccount.CollectorDescription
			properties["collector_id"] = cloudAccount.CollectorID
			properties["country"] = cloudAccount.Country
			properties["external_id"] = cloudAccount.ExternalID
			properties["schedule"] = cloudAccount.Schedule
			properties["secret_key"] = cloudAccount.SecretKey
			properties["subscription_ids"] = cloudAccount.SubscriptionIds
			properties["tenant_id"] = cloudAccount.TenantID
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
	clientID := d["client_id"].(string)
	collectorDescription := d["collector_description"].(string)
	country := d["country"].(string)
	externalID := d["external_id"].(string)
	schedule := d["schedule"].(string)
	secretKey := d["secret_key"].(string)
	subscriptionIds := d["subscription_ids"].(string)
	tenantID := d["tenant_id"].(string)
	
	return &models.CloudAccount {
		AccountID: accountID,
		AssumedRoleArn: assumedRoleArn,
		ClientID: clientID,
		CollectorDescription: collectorDescription,
		Country: country,
		ExternalID: externalID,
		Schedule: schedule,
		SecretKey: secretKey,
		SubscriptionIds: subscriptionIds,
		TenantID: tenantID,
	}
}

func GetCloudAccountPropertyFields() (t []string) {
	return []string{
		"account_id",
		"assumed_role_arn",
		"client_id",
		"collector_description",
		"country",
		"external_id",
		"schedule",
		"secret_key",
		"subscription_ids",
		"tenant_id",
	}
}