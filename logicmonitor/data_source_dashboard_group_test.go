package logicmonitor

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	lmclient "github.com/logicmonitor/lm-sdk-go/client"
)

func TestAccCheckLogicMonitorDashboardGroupDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TestAccCheckLogicMonitorDashboardGroupDataSourceConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDashboardGroupID("data.logicmonitor_dashboard_group.board1"),
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
resource "logicmonitor_dashboard_group" "group1" {
	name = "NewTestDashboardGroup"
	description = "this is a test for creating dashboard groups"
}

data "logicmonitor_dashboard_group" "board1" {
  filters {
    property = "name"
    operator = ":"
    value = "${logicmonitor_dashboard_group.group1.name}"
  }
}
`
