package logicmonitor

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	lmclient "github.com/logicmonitor/lm-sdk-go/client"
)

// this test assumes you have a collector installed already, currently this provider does not handle collector installation

func TestAccCheckLogicMonitorDashboardGroupDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testDashboardGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: TestAccCheckLogicMonitorDashboardGroupDataSourceConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDashboardGroupID("data.logicmonitor_dashboard_group.DashboardGroups"),
				),
			},
		},
	})
}

func testAccCheckDashboardGroupID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find logicmonitor dashboard group data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("logicmonitor dashgroup group data source ID not set")
		}

		client := testAccProvider.Meta().(*lmclient.LMSdkGo)
		if err := testDashboardGroupExistsHelper(s, client); err != nil {
			return err
		}

		return nil
	}
}

const TestAccCheckLogicMonitorDashboardGroupDataSourceConfig = `
resource "logicmonitor_dashboard" "dash" {
    name = "KB824"
    disable_alerting = true
    description = "testing group"
    properties {
     "group" = "test"
     "system.categories" = "a,b,c,d"
    }
}

data "logicmonitor_device_group" "DashboardGroups" {
  filters {
    "property" = "name"
    "operator" = ":"
    "value" = "${logicmonitor_device_group.group1.name}"
  }
}
`
