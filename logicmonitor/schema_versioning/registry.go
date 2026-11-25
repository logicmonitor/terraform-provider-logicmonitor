// Package schema_versioning holds hand-written Terraform state upgraders.
//
// Generated resource code calls Upgraders. The returned version is one
// higher than the last registered upgrader, or 0 when a resource has none.
// Released upgraders stay in this package so a state file from any earlier
// schema version can be upgraded in one step.
package schema_versioning

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-logicmonitor/logicmonitor/schema_versioning/upgraders"
)

// Upgraders returns the current schema version and the ordered state
// upgraders for a managed resource. Resources with no breaking schema change
// return version 0 and no upgraders, which leaves their state untouched.
func Upgraders(resourceType string) (int, []schema.StateUpgrader) {
	switch resourceType {
	case "logicmonitor_dashboard":
		dashboard := upgraders.DashboardUpgraders()
		return schemaVersion(dashboard), dashboard
	default:
		return 0, nil
	}
}

// schemaVersion is the last upgrader's version plus one. An empty list is version 0.
func schemaVersion(upgraders []schema.StateUpgrader) int {
	if len(upgraders) == 0 {
		return 0
	}
	return upgraders[len(upgraders)-1].Version + 1
}
