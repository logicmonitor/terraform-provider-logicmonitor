package logicmonitor

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccLogicMonitorDevice_import(t *testing.T) {
	resourceName := "logicmonitor_device.device"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testDeviceGroupDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckLogicMonitorDeviceImport,
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"properties",
				},
			},
		},
	})
}

const testAccCheckLogicMonitorDeviceImport = `
	resource "logicmonitor_device" "device" {
		ip_addr = "test.logicmonitor.com"
		display_name = "testing"
		disable_alerting = true
		hostgroup_id = "1"
		collector = "${data.logicmonitor_collectors.collectors.id}"
		properties {
	 		"system.categories" = "snmp"
		}
	}
	data "logicmonitor_collectors" "collectors" {
	  most_recent = true
	}
	`
