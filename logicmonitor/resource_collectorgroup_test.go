package logicmonitor

import (
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	lmv1 "github.com/logicmonitor/lm-sdk-go"
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
					testCollectorGroupExists("logicmonitor_collectorgroup.id"),
					resource.TestCheckResourceAttr(
						"logicmonitor_collectorgroup.group1", "name", "collectorGroupTest"),
					resource.TestCheckResourceAttr(
						"logicmonitor_collectorgroup.group1", "description", "testing group"),
				),
			},
		},
	})
}

func testCollectorGroupDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*lmv1.DefaultApi)
	if err := testCollectorGroupDestroyHelper(s, client); err != nil {
		return err
	}
	return nil
}

func testCollectorGroupDestroyHelper(s *terraform.State, client *lmv1.DefaultApi) error {
	for _, r := range s.RootModule().Resources {
		id, e := strconv.Atoi(r.Primary.ID)
		if e != nil {
			return e
		}

		restCollectorGroupResponse, apiResponse, e := client.DeleteCollectorGroupById(int32(id))
		err := checkStatus(restCollectorGroupResponse.Status, restCollectorGroupResponse.Errmsg, apiResponse.StatusCode, apiResponse.Status, e)
		if err != nil {
			return err
		}
	}
	return nil
}

func testCollectorGroupExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*lmv1.DefaultApi)
		if err := testCollectorGroupExistsHelper(s, client); err != nil {
			return err
		}
		return nil
	}
}

func testCollectorGroupExistsHelper(s *terraform.State, client *lmv1.DefaultApi) error {
	for _, r := range s.RootModule().Resources {
		id, e := strconv.Atoi(r.Primary.ID)
		if e != nil {
			return e
		}

		restCollectorGroupResponse, apiResponse, e := client.GetCollectorGroupById(int32(id), "")
		err := checkStatus(restCollectorGroupResponse.Status, restCollectorGroupResponse.Errmsg, apiResponse.StatusCode, apiResponse.Status, e)
		if err != nil {
			return err
		}
	}
	return nil
}

const testAccCheckLogicMonitorConfigCollectorGroup = `
resource "logicmonitor_collectorgroup" "group1" {
    name = "collectorGroupTest"
    description = "testing group"
}
`
