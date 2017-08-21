package logicmonitor

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	lmv1 "github.com/logicmonitor/lm-sdk-go"
)

// this test assumes you have a collector installed already, currently this provider does not handle collector installation

func TestAccLogicMonitorCollectorDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccLogicMonitorCollectorDataSourceConfig,
				Check: resource.ComposeTestCheckFunc(
					testCollectorExists("data.logicmonitor_findcollectors.collectors"),
					resource.TestCheckResourceAttr(
						"data.logicmonitor_findcollectors.collectors", "most_recent", "true"),
				),
			},
		},
	})
}

func testCollectorExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find logicmonitor collector data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("logicmonitor collector data source ID not set")
		}

		client := testAccProvider.Meta().(*lmv1.DefaultApi)
		if err := testCollectorExistsHelper(s, client); err != nil {
			return err
		}

		return nil
	}
}

func testCollectorExistsHelper(s *terraform.State, client *lmv1.DefaultApi) error {
	for _, r := range s.RootModule().Resources {
		id, e := strconv.Atoi(r.Primary.ID)
		if e != nil {
			return e
		}

		restCollectorResponse, apiResponse, e := client.GetCollectorById(int32(id), "")
		err := checkStatus(restCollectorResponse.Status, restCollectorResponse.Errmsg, apiResponse.StatusCode, apiResponse.Status, e)
		if err != nil {
			return err
		}
	}
	return nil
}

// finds the most recent collector created that is online
const testAccLogicMonitorCollectorDataSourceConfig = `
data "logicmonitor_findcollectors" "collectors" {
  most_recent = true
}
`
