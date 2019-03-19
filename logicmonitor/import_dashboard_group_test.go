package logicmonitor

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccLogicMonitorDashboardGroup_import(t *testing.T) {
	resourceName := "logicmonitor_dashboard_group.group1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testDashboardGroupDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckLogicMonitorDashboardGroupImport,
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"force_delete",
				},
			},
		},
	})
}

const testAccCheckLogicMonitorDashboardGroupImport = `
resource "logicmonitor_dashboard_group" "group1" {
    name = "DashboardGroup"
		description = "testing dashboard group import"
}
`
