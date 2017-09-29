package logicmonitor

import (
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	lmv1 "github.com/logicmonitor/lm-sdk-go"
)

func TestAccLogicMonitorDeviceGroup(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testDeviceGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckLogicMonitorDeviceGroup,
				Check: resource.ComposeTestCheckFunc(
					testDeviceGroupExists("logicmonitor_device_group.group1"),
					resource.TestCheckResourceAttr(
						"logicmonitor_device_group.group1", "name", "NewGroup"),
					resource.TestCheckResourceAttr(
						"logicmonitor_device_group.group1", "disable_alerting", "true"),
					resource.TestCheckResourceAttr(
						"logicmonitor_device_group.group1", "properties.group", "test"),
					resource.TestCheckResourceAttr(
						"logicmonitor_device_group.group1", "properties.system.categories", "a,b,c,d"),
				),
			},
		},
	})
}

func testDeviceGroupDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*lmv1.DefaultApi)
	if err := testCollectorGroupDestroyHelper(s, client); err != nil {
		return err
	}
	return nil
}

func testDeviceGroupDestroyHelper(s *terraform.State, client *lmv1.DefaultApi) error {
	for _, r := range s.RootModule().Resources {
		id, e := strconv.Atoi(r.Primary.ID)
		if e != nil {
			return e
		}

		restDeviceGroupResponse, apiResponse, e := client.DeleteDeviceGroupById(int32(id), false)
		err := checkStatus(restDeviceGroupResponse.Status, restDeviceGroupResponse.Errmsg, apiResponse.StatusCode, apiResponse.Status, e)
		if err != nil {
			return err
		}
	}
	return nil
}

func testDeviceGroupExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*lmv1.DefaultApi)
		if err := testDeviceGroupExistsHelper(s, client); err != nil {
			return err
		}
		return nil
	}
}

func testDeviceGroupExistsHelper(s *terraform.State, client *lmv1.DefaultApi) error {
	for _, r := range s.RootModule().Resources {
		id, e := strconv.Atoi(r.Primary.ID)
		if e != nil {
			return e
		}

		restDeviceGroupResponse, apiResponse, e := client.GetDeviceGroupById(int32(id), "")
		err := checkStatus(restDeviceGroupResponse.Status, restDeviceGroupResponse.Errmsg, apiResponse.StatusCode, apiResponse.Status, e)
		if err != nil {
			return err
		}
	}
	return nil
}

const testAccCheckLogicMonitorDeviceGroup = `
resource "logicmonitor_device_group" "group1" {
    name = "NewGroup"
    disable_alerting = true
    description = "testing group"
    applies_to = "system.displayname =~ \"test\""
    parent_id = 0
    properties {
     "group" = "test"
     "system.categories" = "a,b,c,d"
    }
}
`
