package logicmonitor

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccLogicMonitorDashboardGroup_import(t *testing.T) {
	resourceName := "logicmonitor_dashboard_group.Lakers"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
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
					"widget_tokens",
				},
			},
		},
	})
}

const testAccCheckLogicMonitorDashboardGroupImport = `
resource "logicmonitor_dashboard_group" "Lakers" {
  name = "Lakers"
	description = "testing dashboard group import"
}
`
