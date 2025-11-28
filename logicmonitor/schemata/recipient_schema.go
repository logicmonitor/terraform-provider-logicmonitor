package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-logicmonitor/models"
)

func RecipientSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"addr": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"contact": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"method": {
			Type:     schema.TypeString,
			Required: true,
		},

		"type": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

func SetRecipientSubResourceData(m []*models.Recipient) (d []*map[string]interface{}) {
	for _, recipient := range m {
		if recipient != nil {
			properties := make(map[string]interface{})
			properties["addr"] = recipient.Addr
			properties["contact"] = recipient.Contact
			properties["method"] = recipient.Method
			properties["type"] = recipient.Type
			d = append(d, &properties)
		}
	}
	return
}

func RecipientModel(d map[string]interface{}) *models.Recipient {
	// assume that the incoming map only contains the relevant resource data
	addr := d["addr"].(string)
	contact := d["contact"].(string)
	method := d["method"].(string)
	typeVar := d["type"].(string)

	return &models.Recipient{
		Addr:    addr,
		Contact: contact,
		Method:  &method,
		Type:    &typeVar,
	}
}

func GetRecipientPropertyFields() (t []string) {
	return []string{
		"addr",
		"contact",
		"method",
		"type",
	}
}
