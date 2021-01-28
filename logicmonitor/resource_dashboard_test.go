package logicmonitor

import (
	"log"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	lmclient "github.com/logicmonitor/lm-sdk-go/client"
	"github.com/logicmonitor/lm-sdk-go/client/lm"
)

func TestAccLogicMonitorDashboard(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckLogicMonitorDashboard,
				Check: resource.ComposeTestCheckFunc(
					testDashboardExists("logicmonitor_dashboard.kb24"),
					resource.TestCheckResourceAttr(
						"logicmonitor_dashboard.kb24", "name", "81Points"),
					resource.TestCheckResourceAttr(
						"logicmonitor_dashboard.kb24", "description", "January 22, 2006"),
				),
			},
		},
	})
}

func testDashboardExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*lmclient.LMSdkGo)
		if err := testDashboardExistsHelper(s, client); err != nil {
			return err
		}
		return nil
	}
}

func testDashboardExistsHelper(s *terraform.State, client *lmclient.LMSdkGo) error {
	for _, r := range s.RootModule().Resources {
		id, e := strconv.Atoi(r.Primary.ID)
		if e != nil {
			return e
		}
		params := lm.NewGetDashboardByIDParams()
		params.SetID(int32(id))

		restDashboardResponse, err := client.LM.GetDashboardByID(params)
		if err != nil {
			return err
		}
		log.Printf("delete Dashboard response %v", restDashboardResponse)
	}
	return nil
}

const testAccCheckLogicMonitorDashboard = `
resource "logicmonitor_dashboard" "kb24" {
	name = "81Points"
	description = "January 22, 2006"
}
`
