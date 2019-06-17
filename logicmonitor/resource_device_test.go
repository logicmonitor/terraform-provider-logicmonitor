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

func TestAccLogicMonitorDevices(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckLogicMonitorDevice,
				Check: resource.ComposeTestCheckFunc(
					testDeviceExists("logicmonitor_device.device1.ip_addr"),
					resource.TestCheckResourceAttr(
						"logicmonitor_device.device1", "display_name", "testing"),
					resource.TestCheckResourceAttr(
						"logicmonitor_device.device1", "ip_addr", "10.32.12.34"),
					resource.TestCheckResourceAttr(
						"logicmonitor_device.device1", "disable_alerting", "true"),
					resource.TestCheckResourceAttr(
						"logicmonitor_device.device1", "hostgroup_id", "1"),
					resource.TestCheckResourceAttr(
						"logicmonitor_device.device1", "properties.app", "haproxy"),
					resource.TestCheckResourceAttr(
						"logicmonitor_device.device1", "properties.system.categories", "a,b,c,d"),
				)},
		},
	})
}

func testDeviceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*lmclient.LMSdkGo)
		if err := testDeviceExistsHelper(s, client); err != nil {
			return err
		}
		return nil
	}
}

func testDeviceExistsHelper(s *terraform.State, client *lmclient.LMSdkGo) error {
	deviceID := s.RootModule().Resources["logicmonitor_device.device1"]
	id, e := strconv.Atoi(deviceID.Primary.ID)
	if e != nil {
		return e
	}
	params := lm.NewGetDeviceByIDParams()
	params.SetID(int32(id))
	restResponse, err := client.LM.GetDeviceByID(params)
	if err != nil {
		return err
	}
	log.Printf("update device response %v", restResponse)
	return nil
}

const testAccCheckLogicMonitorDevice = `
	resource "logicmonitor_device" "device1" {
		ip_addr = "10.32.12.34"
		display_name = "testing"
		disable_alerting = true
		hostgroup_id = "1"
		collector = "${data.logicmonitor_collectors.collectors.id}"
		properties = {
	 		"app" = "haproxy"
	 		"system.categories" = "a,b,c,d"
		}
	}
	data "logicmonitor_collectors" "collectors" {
	  most_recent = true
	}
	`
