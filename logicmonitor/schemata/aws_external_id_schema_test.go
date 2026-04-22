// Description: Unit tests for AWS External ID schema and resource data functions.
// Description: Validates Computed fields, resource ID assignment, and nil handling.
package schemata

import (
	"terraform-provider-logicmonitor/models"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestAwsExternalIDSchemaFieldsAreComputed(t *testing.T) {
	s := AwsExternalIDSchema()

	tests := []struct {
		field string
	}{
		{"created_at"},
		{"external_id"},
	}

	for _, tt := range tests {
		f, ok := s[tt.field]
		if !ok {
			t.Errorf("expected field %q in schema", tt.field)
			continue
		}
		if !f.Computed {
			t.Errorf("field %q should be Computed", tt.field)
		}
		if f.Optional {
			t.Errorf("field %q should not be Optional", tt.field)
		}
	}
}

func TestSetAwsExternalIDResourceDataSetsID(t *testing.T) {
	resourceSchema := map[string]*schema.Schema{
		"created_at":  {Type: schema.TypeString, Computed: true},
		"external_id": {Type: schema.TypeString, Computed: true},
	}
	d := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})

	model := &models.AwsExternalID{
		CreatedAt:  "2026-04-02T12:07:00Z",
		ExternalID: "0dfc2264-482c-4f74-b8b9-ceb7604c7492",
	}

	SetAwsExternalIDResourceData(d, model)

	if got := d.Id(); got != "0dfc2264-482c-4f74-b8b9-ceb7604c7492" {
		t.Errorf("expected ID to be the ExternalID UUID, got %q", got)
	}
	if got := d.Get("created_at").(string); got != "2026-04-02T12:07:00Z" {
		t.Errorf("expected created_at %q, got %q", "2026-04-02T12:07:00Z", got)
	}
	if got := d.Get("external_id").(string); got != "0dfc2264-482c-4f74-b8b9-ceb7604c7492" {
		t.Errorf("expected external_id %q, got %q", "0dfc2264-482c-4f74-b8b9-ceb7604c7492", got)
	}
}

func TestSetAwsExternalIDResourceDataIDIsStable(t *testing.T) {
	resourceSchema := map[string]*schema.Schema{
		"created_at":  {Type: schema.TypeString, Computed: true},
		"external_id": {Type: schema.TypeString, Computed: true},
	}

	model := &models.AwsExternalID{
		CreatedAt:  "2026-04-02T12:07:00Z",
		ExternalID: "abc-123-def-456",
	}

	// Call twice and verify the ID is the same (not timestamp-based)
	d1 := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
	SetAwsExternalIDResourceData(d1, model)

	d2 := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
	SetAwsExternalIDResourceData(d2, model)

	if d1.Id() != d2.Id() {
		t.Errorf("resource ID should be stable across calls, got %q and %q", d1.Id(), d2.Id())
	}
}

func TestSetAwsExternalIDSubResourceDataHandlesNil(t *testing.T) {
	result := SetAwsExternalIDSubResourceData(nil)
	if result != nil {
		t.Errorf("expected nil for nil input, got %v", result)
	}

	result = SetAwsExternalIDSubResourceData([]*models.AwsExternalID{nil})
	if result != nil {
		t.Errorf("expected nil for slice with nil element, got %v", result)
	}
}

func TestSetAwsExternalIDSubResourceDataPopulatesFields(t *testing.T) {
	input := []*models.AwsExternalID{
		{
			CreatedAt:  "2026-04-02T12:07:00Z",
			ExternalID: "test-uuid-123",
		},
	}

	result := SetAwsExternalIDSubResourceData(input)
	if len(result) != 1 {
		t.Fatalf("expected 1 result, got %d", len(result))
	}

	props := *result[0]
	if props["created_at"] != "2026-04-02T12:07:00Z" {
		t.Errorf("expected created_at %q, got %q", "2026-04-02T12:07:00Z", props["created_at"])
	}
	if props["external_id"] != "test-uuid-123" {
		t.Errorf("expected external_id %q, got %q", "test-uuid-123", props["external_id"])
	}
}
