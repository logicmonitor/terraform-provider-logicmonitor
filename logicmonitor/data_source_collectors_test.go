package logicmonitor

import (
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	lmclient "github.com/logicmonitor/lm-sdk-go/client"
	"github.com/logicmonitor/lm-sdk-go/client/lm"
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
					testDataSourceCollectorExists("data.logicmonitor_collectors.collectors"),
					resource.TestCheckResourceAttr(
						"data.logicmonitor_collectors.collectors", "most_recent", "true"),
				),
			},
		},
	})
}

func testDataSourceCollectorExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find logicmonitor collector data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("logicmonitor collector data source ID not set")
		}

		client := testAccProvider.Meta().(*lmclient.LMSdkGo)
		if err := testDataSourceCollectorExistsHelper(s, client); err != nil {
			return err
		}

		return nil
	}
}

func testDataSourceCollectorExistsHelper(s *terraform.State, client *lmclient.LMSdkGo) error {
	for _, r := range s.RootModule().Resources {
		id, e := strconv.Atoi(r.Primary.ID)
		if e != nil {
			return e
		}

		params := lm.NewGetCollectorByIDParams()
		params.SetID(int32(id))

		restCollectorResponse, err := client.LM.GetCollectorByID(params)
		if err != nil {
			return err
		}
		log.Printf("get collector id response %v", restCollectorResponse.Payload)
	}
	return nil
}

// finds the most recent collector created that is online
const testAccLogicMonitorCollectorDataSourceConfig = `
data "logicmonitor_collectors" "collectors" {
  most_recent = true
}
`
