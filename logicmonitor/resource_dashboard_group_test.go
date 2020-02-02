package logicmonitor

import (
	"log"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	lmclient "github.com/logicmonitor/lm-sdk-go/client"
	"github.com/logicmonitor/lm-sdk-go/client/lm"
)

func TestAccLogicMonitorDashboardGroup(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckLogicMonitorDashboardGroup,
				Check: resource.ComposeTestCheckFunc(
					testDashboardGroupExists("logicmonitor_dashboard_group.lakersforever"),
					resource.TestCheckResourceAttr(
						"logicmonitor_dashboard_group.lakersforever", "name", "LoveLiveKobe"),
					resource.TestCheckResourceAttr(
						"logicmonitor_dashboard_group.lakersforever", "description", "long live the king"),
					resource.TestCheckResourceAttr(
						"logicmonitor_dashboard_group.lakersforever", "force_delete", "true"),
				),
			},
		},
	})
}

func testDashboardGroupExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*lmclient.LMSdkGo)
		if err := testDashboardGroupExistsHelper(s, client); err != nil {
			return err
		}
		return nil
	}
}

func testDashboardGroupExistsHelper(s *terraform.State, client *lmclient.LMSdkGo) error {
	for _, r := range s.RootModule().Resources {
		id, e := strconv.Atoi(r.Primary.ID)
		if e != nil {
			return e
		}
		params := lm.NewGetDashboardGroupByIDParams()
		params.SetID(int32(id))

		restDashboardGroupResponse, err := client.LM.GetDashboardGroupByID(params)
		if err != nil {
			return err
		}
		log.Printf("delete dashboardGroup response %v", restDashboardGroupResponse)
	}
	return nil
}

const testAccCheckLogicMonitorDashboardGroup = `
resource "logicmonitor_dashboard_group" "lakersforever" {
	name = "LoveLiveKobe"
	description = "long live the king"
}
`
