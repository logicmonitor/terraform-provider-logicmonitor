package logicmonitor

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	lmclient "github.com/logicmonitor/lm-sdk-go/client"
)

func TestAccCheckLogicMonitorDashboardDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TestAccCheckLogicMonitorDashboardDataSourceConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDashboardID("data.logicmonitor_dashboard.board1"),
				),
			},
		},
	})
}

func testAccCheckDashboardID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find logicmonitor dashboard  data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("logicmonitor dash  data source ID not set")
		}

		client := testAccProvider.Meta().(*lmclient.LMSdkGo)
		if err := testDashboardExistsHelper(s, client); err != nil {
			return err
		}

		return nil
	}
}

const TestAccCheckLogicMonitorDashboardDataSourceConfig = `
resource "logicmonitor_dashboard" "brd" {
	name = "Kobe Doin Work"
	description = "Another Kobe Tribute Dashboard"
}

data "logicmonitor_dashboard" "board1" {
  filters {
    property = "name"
    operator = ":"
    value = "${logicmonitor_dashboard.brd.name}"
  }
}
`
