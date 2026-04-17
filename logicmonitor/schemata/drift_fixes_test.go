// Description: Unit tests for configuration drift fixes across multiple schemas.
// Description: Validates DiffSuppressFunc, nil normalization, bounds checks, and Computed fields.
package schemata

import (
	"terraform-provider-logicmonitor/models"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// --- Bug 5: DiffSuppressFunc for masked sensitive values ---

func TestNameAndValueDiffSuppressMaskedValue(t *testing.T) {
	s := NameAndValueSchema()
	valueField := s["value"]

	if valueField.DiffSuppressFunc == nil {
		t.Fatal("value field must have a DiffSuppressFunc")
	}

	tests := []struct {
		name     string
		old      string
		new      string
		suppress bool
	}{
		{"masked old value suppresses diff", "**********", "realPassword123", true},
		{"masked old value suppresses any new value", "**********", "", true},
		{"non-masked values do not suppress", "oldValue", "newValue", false},
		{"identical values do not suppress", "same", "same", false},
		{"empty old does not suppress", "", "newValue", false},
		{"new value is mask does not suppress", "realPassword", "**********", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := valueField.DiffSuppressFunc("custom_properties.0.value", tt.old, tt.new, nil)
			if got != tt.suppress {
				t.Errorf("DiffSuppressFunc(%q, %q) = %v, want %v", tt.old, tt.new, got, tt.suppress)
			}
		})
	}
}

// --- Bug 4: Device group extra field Computed ---

func TestDeviceGroupExtraFieldIsComputed(t *testing.T) {
	s := DeviceGroupSchema()
	extra, ok := s["extra"]
	if !ok {
		t.Fatal("expected extra field in DeviceGroupSchema")
	}
	if !extra.Optional {
		t.Error("extra field should be Optional")
	}
	if !extra.Computed {
		t.Error("extra field should be Computed")
	}
}

// --- Bug 8: custom_properties Computed across resource schemas ---

func TestCustomPropertiesComputedInResourceSchemas(t *testing.T) {
	tests := []struct {
		name   string
		schema map[string]*schema.Schema
	}{
		{"DeviceGroupSchema", DeviceGroupSchema()},
		{"DeviceSchema", DeviceSchema()},
		{"CollectorSchema", CollectorSchema()},
		{"CollectorGroupSchema", CollectorGroupSchema()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cp, ok := tt.schema["custom_properties"]
			if !ok {
				t.Fatalf("expected custom_properties field in %s", tt.name)
			}
			if !cp.Optional {
				t.Errorf("custom_properties in %s should be Optional", tt.name)
			}
			if !cp.Computed {
				t.Errorf("custom_properties in %s should be Computed", tt.name)
			}
		})
	}
}

// --- Bug 6: escalating_chain nil normalization ---

func TestSetAlertRuleResourceDataNilEscalatingChain(t *testing.T) {
	s := AlertRuleSchema()
	d := schema.TestResourceDataRaw(t, s, map[string]interface{}{
		"datapoint":              "*",
		"datasource":            "*",
		"device_groups":         []interface{}{"*"},
		"devices":               []interface{}{"*"},
		"escalating_chain_id":   1,
		"escalation_interval":   15,
		"instance":              "*",
		"name":                  "test",
		"priority":              100,
	})

	model := &models.AlertRule{
		Datapoint:        strPtr("*"),
		Datasource:       strPtr("*"),
		EscalatingChain:  nil, // API returns null
		EscalatingChainID: int32Ptr(1),
		EscalationInterval: int32Ptr(15),
		Instance:         strPtr("*"),
		Name:             strPtr("test"),
		Priority:         int32Ptr(100),
	}

	SetAlertRuleResourceData(d, model)

	ec := d.Get("escalating_chain")
	if ec == nil {
		t.Error("escalating_chain should be empty map, not nil")
	}
	ecMap, ok := ec.(map[string]interface{})
	if !ok {
		t.Errorf("escalating_chain should be map[string]interface{}, got %T", ec)
	}
	if len(ecMap) != 0 {
		t.Errorf("escalating_chain should be empty map, got %v", ecMap)
	}
}

func TestSetAlertRuleResourceDataNonNilEscalatingChain(t *testing.T) {
	s := AlertRuleSchema()
	d := schema.TestResourceDataRaw(t, s, map[string]interface{}{
		"datapoint":              "*",
		"datasource":            "*",
		"device_groups":         []interface{}{"*"},
		"devices":               []interface{}{"*"},
		"escalating_chain_id":   1,
		"escalation_interval":   15,
		"instance":              "*",
		"name":                  "test",
		"priority":              100,
	})

	chainData := map[string]interface{}{"key": "value"}
	model := &models.AlertRule{
		Datapoint:          strPtr("*"),
		Datasource:         strPtr("*"),
		EscalatingChain:    chainData,
		EscalatingChainID:  int32Ptr(1),
		EscalationInterval: int32Ptr(15),
		Instance:           strPtr("*"),
		Name:               strPtr("test"),
		Priority:           int32Ptr(100),
	}

	SetAlertRuleResourceData(d, model)

	ec := d.Get("escalating_chain")
	if ec == nil {
		t.Fatal("escalating_chain should not be nil")
	}
}

// --- Bug 7: chain stages bounds check ---

func TestSetChainSubResourceDataEmptyStages(t *testing.T) {
	typeStr := "normal"
	chain := &models.Chain{
		Period: nil,
		Stages: [][]*models.Recipient{}, // empty stages
		Type:   &typeStr,
	}

	// Should not panic
	result := SetChainSubResourceData([]*models.Chain{chain})

	if len(result) != 1 {
		t.Fatalf("expected 1 result, got %d", len(result))
	}

	props := *result[0]
	stages, ok := props["stages"]
	if !ok {
		t.Fatal("expected stages key in result")
	}

	stageSlice, ok := stages.([]*map[string]interface{})
	if !ok {
		t.Fatalf("stages should be []*map[string]interface{}, got %T", stages)
	}
	if len(stageSlice) != 0 {
		t.Errorf("stages should be empty for chain with no stages, got %d", len(stageSlice))
	}
}

func TestSetChainSubResourceDataNilStages(t *testing.T) {
	typeStr := "normal"
	chain := &models.Chain{
		Period: nil,
		Stages: nil, // nil stages
		Type:   &typeStr,
	}

	// Should not panic
	result := SetChainSubResourceData([]*models.Chain{chain})

	if len(result) != 1 {
		t.Fatalf("expected 1 result, got %d", len(result))
	}
}

func TestSetChainSubResourceDataWithStages(t *testing.T) {
	typeStr := "normal"
	addr := "admin@example.com"
	method := "email"
	recipientType := "ADMIN"
	contact := "admin"

	chain := &models.Chain{
		Period: nil,
		Stages: [][]*models.Recipient{
			{
				{Addr: addr, Method: &method, Type: &recipientType, Contact: contact},
			},
		},
		Type: &typeStr,
	}

	result := SetChainSubResourceData([]*models.Chain{chain})

	if len(result) != 1 {
		t.Fatalf("expected 1 result, got %d", len(result))
	}

	props := *result[0]
	stages := props["stages"]
	stageSlice, ok := stages.([]*map[string]interface{})
	if !ok {
		t.Fatalf("stages should be []*map[string]interface{}, got %T", stages)
	}
	if len(stageSlice) != 1 {
		t.Errorf("expected 1 recipient in stages, got %d", len(stageSlice))
	}
}

func TestSetChainSubResourceDataNilChain(t *testing.T) {
	result := SetChainSubResourceData(nil)
	if result != nil {
		t.Errorf("expected nil for nil input, got %v", result)
	}

	result = SetChainSubResourceData([]*models.Chain{nil})
	if result != nil {
		t.Errorf("expected nil for slice with nil element, got %v", result)
	}
}

// --- Helpers ---

func strPtr(s string) *string    { return &s }
func int32Ptr(i int32) *int32    { return &i }
