// Package upgraders holds the state transformations for each schema version.
package upgraders

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// jsonEncodedTemplateKeys are the template map entries that provider builds
// before commit 246bead8 stored as JSON text. NormalizeJSONMap parsed them
// before the API call, so the v0 to v1 upgrader parses them the same way
// before the map is stored as one JSON string.
var jsonEncodedTemplateKeys = map[string]bool{
	"widgets":       true,
	"widgetTokens":  true,
	"widgetsConfig": true,
}

// DashboardUpgraders migrates logicmonitor_dashboard from schema version 0
// (template stored as a map of strings) to version 1 (template stored as a
// JSON string).
func DashboardUpgraders() []schema.StateUpgrader {
	return []schema.StateUpgrader{
		{
			Version: 0,
			Type:    dashboardV0Type(),
			Upgrade: DashboardV0ToV1,
		},
	}
}

// dashboardV0Type is the cty type of the dashboard schema before template
// changed from a map of strings to a string. It is a copy, not a call into
// the generated schemata package, so regeneration cannot change it.
func dashboardV0Type() cty.Type {
	widgetTokenInheritance := cty.Object(map[string]cty.Type{
		"fullpath": cty.String,
		"value":    cty.String,
	})
	widgetToken := cty.Object(map[string]cty.Type{
		"inherit_list": cty.List(widgetTokenInheritance),
		"name":         cty.String,
		"type":         cty.String,
		"value":        cty.String,
	})
	return cty.Object(map[string]cty.Type{
		"description":     cty.String,
		"full_name":       cty.String,
		"group_full_path": cty.String,
		"group_id":        cty.Number,
		"group_name":      cty.String,
		"id":              cty.String,
		"name":            cty.String,
		"owner":           cty.String,
		"sharable":        cty.Bool,
		"template":        cty.Map(cty.String),
		"user_permission": cty.String,
		"widget_tokens":   cty.List(widgetToken),
		"widgets_config":  cty.Map(cty.String),
	})
}

// DashboardV0ToV1 converts a dashboard state written at schema version 0 into
// schema version 1. It does not call the LogicMonitor API.
//
// template is rewritten as follows:
//   - missing, null, or an empty map becomes null
//   - a string is left unchanged, because builds after the type change still
//     wrote version 0 state
//   - a map is serialized to a JSON string, after JSON-encoded widgets,
//     widgetTokens, and widgetsConfig values are parsed
//   - any other type returns an error
func DashboardV0ToV1(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	raw, ok := rawState["template"]
	if !ok || raw == nil {
		rawState["template"] = nil
		return rawState, nil
	}

	switch value := raw.(type) {
	case string:
		return rawState, nil
	case map[string]interface{}:
		if len(value) == 0 {
			rawState["template"] = nil
			return rawState, nil
		}
		encoded, err := json.Marshal(normalizeTemplateMap(value))
		if err != nil {
			return nil, fmt.Errorf("logicmonitor_dashboard: attribute \"template\": %w", err)
		}
		rawState["template"] = string(encoded)
		return rawState, nil
	default:
		return nil, fmt.Errorf("logicmonitor_dashboard: attribute \"template\" has type %T, want a string or a map", raw)
	}
}

// normalizeTemplateMap returns a copy of a v0 template map with the
// historically JSON-encoded keys parsed into structured values.
func normalizeTemplateMap(template map[string]interface{}) map[string]interface{} {
	normalized := make(map[string]interface{}, len(template))
	for key, value := range template {
		text, isString := value.(string)
		if isString && jsonEncodedTemplateKeys[key] {
			var parsed interface{}
			if err := json.Unmarshal([]byte(text), &parsed); err == nil {
				normalized[key] = parsed
				continue
			}
		}
		normalized[key] = value
	}
	return normalized
}
