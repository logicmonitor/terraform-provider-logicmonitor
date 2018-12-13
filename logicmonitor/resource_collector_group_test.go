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

func TestAccLogicMonitorCollectorGroup(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCollectorGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckLogicMonitorConfigCollectorGroup,
				Check: resource.ComposeTestCheckFunc(
					testCollectorGroupExists("logicmonitor_collector_group.id"),
					resource.TestCheckResourceAttr(
						"logicmonitor_collector_group.group1", "name", "collectorGroupTest"),
					resource.TestCheckResourceAttr(
						"logicmonitor_collector_group.group1", "description", "testing group"),
				),
			},
		},
	})
}

func testCollectorGroupDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*lmclient.LMSdkGo)
	if err := testCollectorGroupDestroyHelper(s, client); err != nil {
		return err
	}
	return nil
}

func testCollectorGroupDestroyHelper(s *terraform.State, client *lmclient.LMSdkGo) error {
	for _, r := range s.RootModule().Resources {
		id, e := strconv.Atoi(r.Primary.ID)
		if e != nil {
			return e
		}

		params := lm.NewDeleteCollectorGroupByIDParams()
		params.ID = int32(id)
		restCollectorGroupResponse, err := client.LM.DeleteCollectorGroupByID(params)
		if err != nil {
			return err
		}
		log.Printf("delete collector group payload response %v", restCollectorGroupResponse.Payload)
	}
	return nil
}

func testCollectorGroupExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*lmclient.LMSdkGo)
		if err := testCollectorGroupExistsHelper(s, client); err != nil {
			return err
		}
		return nil
	}
}

func testCollectorGroupExistsHelper(s *terraform.State, client *lmclient.LMSdkGo) error {
	for _, r := range s.RootModule().Resources {
		id, e := strconv.Atoi(r.Primary.ID)
		if e != nil {
			return e
		}

		params := lm.NewGetCollectorGroupByIDParams()
		params.ID = (int32(id))

		restCollectorGroupResponse, err := client.LM.GetCollectorGroupByID(params)
		if err != nil {
			return err
		}
		log.Printf("payload response %v", restCollectorGroupResponse.Payload)
	}
	return nil
}

const testAccCheckLogicMonitorConfigCollectorGroup = `
resource "logicmonitor_collector_group" "group1" {
    name = "collectorGroupTest"
    description = "testing group"
}
`
