package logicmonitor

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccLogicMonitorCollectorGroup_import(t *testing.T) {
	resourceName := "logicmonitor_collector_group.group1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCollectorGroupDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckLogicMonitorCollectorGroupImport,
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccCheckLogicMonitorCollectorGroupImport = `
resource "logicmonitor_collector_group" "group1" {
    name = "CollectorGroup"
		description = "testing collector group import"
}
`
