package logicmonitor

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	lmclient "github.com/logicmonitor/lm-sdk-go/client"
)

// this test assumes you have a collector installed already, currently this provider does not handle collector installation

func TestAccCheckLogicMonitorDeviceGroupDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testDeviceGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckLogicMonitorDeviceGroupDataSourceConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDeviceGroupID("data.logicmonitor_device_group.devicegroups"),
				),
			},
		},
	})
}

func testAccCheckDeviceGroupID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find logicmonitor device group data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("logicmonitor device group data source ID not set")
		}

		client := testAccProvider.Meta().(*lmclient.LMSdkGo)
		if err := testDeviceGroupExistsHelper(s, client); err != nil {
			return err
		}

		return nil
	}
}

const testAccCheckLogicMonitorDeviceGroupDataSourceConfig = `
resource "logicmonitor_device_group" "group1" {
    name = "KB824"
    disable_alerting = true
    description = "testing group"
    properties = {
     "group" = "test"
     "system.categories" = "a,b,c,d"
    }
}
data "logicmonitor_device_group" "devicegroups" {
  filters {
    property = "name"
    operator = ":"
    value = "${logicmonitor_device_group.group1.name}"
  }
}
`
