package logicmonitor

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	lmclient "github.com/logicmonitor/lm-sdk-go/client"
	"github.com/logicmonitor/lm-sdk-go/client/lm"
)

// terraform data source to look up collectors
func dataSourceFindCollectors() *schema.Resource {
	return &schema.Resource{
		Read: findDataSourceCollectors,

		Schema: map[string]*schema.Schema{
			"filters": {
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"property": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"operator": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"custom_property_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"custom_property_value": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"most_recent": {
				Type:     schema.TypeBool,
				Default:  false,
				Optional: true,
			},
			"size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  50,
			},
			"offset": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
		},
	}
}

// function to find collectors with certain filters
func findDataSourceCollectors(d *schema.ResourceData, m interface{}) error {
	client := m.(*lmclient.LMSdkGo)
	filter := getCollectorFilters(d)
	size := int32(d.Get("size").(int))
	offset := int32(d.Get("offset").(int))
	var collectorIds []string

	params := lm.NewGetCollectorListParams()
	params.SetFilter(&filter)
	params.SetSize(&size)
	params.SetOffset(&offset)

	//looks for collector list with selected filter
	// we sort by created date so its easier if we need to choose the latest created collector
	restCollectorPaginationResponse, err := client.LM.GetCollectorList(params)
	if err != nil {
		return err
	}

	for _, e := range restCollectorPaginationResponse.Payload.Items {
		log.Printf("Found collector with filter with status %v", *e.IsDown)
		if !*e.IsDown {
			collectorIds = append(collectorIds, strconv.Itoa(int(e.ID)))
		}
	}

	//comma separated string of device Ids
	var ids = strings.Join(collectorIds, ",")

	if len(collectorIds) == 1 {
		d.SetId(ids)
	} else if len(collectorIds) > 1 {
		if (d.Get("most_recent").(bool)) == true {
			d.SetId(collectorIds[len(collectorIds)-1])
			return nil
		}
		err = fmt.Errorf("Found more than 1 collector id matching this result, a device can only be matched to 1 collector, please change your device or choose a collector to add in tfvars.  %s", ids)
		return err
	} else {
		err = fmt.Errorf("Found no collectors with filter %s", filter)
		return err
	}
	return nil
}
