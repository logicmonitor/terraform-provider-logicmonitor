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
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testDashboardGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckLogicMonitorDashboardGroup,
				Check: resource.ComposeTestCheckFunc(
					testDashboardGroupExists("logicmonitor_dashboard_group.group1"),
					resource.TestCheckResourceAttr(
						"logicmonitor_dashboard_group.group1", "name", "TestDashboardGroup"),
					resource.TestCheckResourceAttr(
						"logicmonitor_dashboard_group.group1", "description", "this is a test for creating dashboard groups"),
				),
			},
		},
	})
}

func testDashboardGroupDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*lmclient.LMSdkGo)
	if err := testCollectorGroupDestroyHelper(s, client); err != nil {
		return err
	}
	return nil
}

func testDashboardGroupDestroyHelper(s *terraform.State, client *lmclient.LMSdkGo) error {
	for _, r := range s.RootModule().Resources {
		id, e := strconv.Atoi(r.Primary.ID)
		if e != nil {
			return e
		}
		forceDelete := true
		params := lm.NewDeleteDashboardGroupByIDParams()
		params.SetAllowNonEmptyGroup(&forceDelete)
		params.SetID(int32(id))

		nil, err := client.LM.DeleteDashboardGroupByID(params)
		if err != nil {
			return err
		}
	}
	return nil
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
resource "logicmonitor_dashboard_group" "group1" {
	name = "TestDashboardGroup"
	description = "this is a test for creating dashboard groups"
}
`
