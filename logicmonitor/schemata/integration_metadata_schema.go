package schemata

import (
	"terraform-provider-logicmonitor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func IntegrationMetadataSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"audited_registry_id": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"audited_version": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"is_changed_from_origin": {
			Type: schema.TypeBool,
			Computed: true,
		},
		
		"is_changed_from_target_last_published": {
			Type: schema.TypeBool,
			Computed: true,
		},
		
		"local_module_id": {
			Type: schema.TypeInt,
			Computed: true,
		},
		
		"logic_module_type": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"origin_author_company_uuid": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"origin_author_namespace": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"origin_checksum": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"origin_lineage_id": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"origin_locator": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"origin_name": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"origin_registry_id": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"origin_version": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"target_last_published_checksum": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"target_last_published_id": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"target_last_published_version": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"target_lineage_id": {
			Type: schema.TypeString,
			Computed: true,
		},
		
	}
}

func SetIntegrationMetadataSubResourceData(m []*models.IntegrationMetadata) (d []*map[string]interface{}) {
	for _, integrationMetadata := range m {
		if integrationMetadata != nil {
			properties := make(map[string]interface{})
			properties["audited_registry_id"] = integrationMetadata.AuditedRegistryID
			properties["audited_version"] = integrationMetadata.AuditedVersion
			properties["is_changed_from_origin"] = integrationMetadata.IsChangedFromOrigin
			properties["is_changed_from_target_last_published"] = integrationMetadata.IsChangedFromTargetLastPublished
			properties["local_module_id"] = integrationMetadata.LocalModuleID
			properties["logic_module_type"] = integrationMetadata.LogicModuleType
			properties["origin_author_company_uuid"] = integrationMetadata.OriginAuthorCompanyUUID
			properties["origin_author_namespace"] = integrationMetadata.OriginAuthorNamespace
			properties["origin_checksum"] = integrationMetadata.OriginChecksum
			properties["origin_lineage_id"] = integrationMetadata.OriginLineageID
			properties["origin_locator"] = integrationMetadata.OriginLocator
			properties["origin_name"] = integrationMetadata.OriginName
			properties["origin_registry_id"] = integrationMetadata.OriginRegistryID
			properties["origin_version"] = integrationMetadata.OriginVersion
			properties["target_last_published_checksum"] = integrationMetadata.TargetLastPublishedChecksum
			properties["target_last_published_id"] = integrationMetadata.TargetLastPublishedID
			properties["target_last_published_version"] = integrationMetadata.TargetLastPublishedVersion
			properties["target_lineage_id"] = integrationMetadata.TargetLineageID
			d = append(d, &properties)
		}
	}
	return
}

func IntegrationMetadataModel(d map[string]interface{}) *models.IntegrationMetadata {
	// assume that the incoming map only contains the relevant resource data
	
	return &models.IntegrationMetadata {
	}
}

func GetIntegrationMetadataPropertyFields() (t []string) {
	return []string{
	}
}