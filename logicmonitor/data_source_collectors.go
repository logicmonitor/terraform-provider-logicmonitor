package logicmonitor

import (
	"fmt"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/hashicorp/terraform/helper/schema"
	lmv1 "github.com/logicmonitor/lm-sdk-go"
)

// terraform data source to look up collectors
func dataSourceFindCollectors() *schema.Resource {
	return &schema.Resource{
		Read: findCollectors,

		Schema: map[string]*schema.Schema{
			"filters": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"property": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"operator": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"most_recent": &schema.Schema{
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
func findCollectors(d *schema.ResourceData, m interface{}) error {
	client := m.(*lmv1.DefaultApi)
	filter := getCollectorFilters(d)
	size := int32(d.Get("size").(int))
	offset := int32(d.Get("offset").(int))
	var collectorIds []string
	//looks for collector list with selected filter
	// we sort by created date so its easier if we need to choose the latest created collector
	restCollectorPaginationResponse, apiResponse, e := client.GetCollectorList("sort=-createdOn&hostname,id,createdOn,isDown", size, offset, filter)
	err := checkStatus(restCollectorPaginationResponse.Status, restCollectorPaginationResponse.Errmsg, apiResponse.StatusCode, apiResponse.Status, e)
	if err != nil {
		return err
	}

	if (d.Get("most_recent").(bool)) == true {
		for _, e := range restCollectorPaginationResponse.Data.Items {
			log.Infof("Found collector with filter %s", filter)
			if !e.IsDown {
				collectorIds = append(collectorIds, strconv.Itoa(int(e.Id)))
				break
			}
		}
	} else {
		for _, e := range restCollectorPaginationResponse.Data.Items {
			log.Infof("Found collector with filter with status %s", e.IsDown)
			if !e.IsDown {
				collectorIds = append(collectorIds, strconv.Itoa(int(e.Id)))
			}
		}

	}

	//comma separated string of device Ids
	var ids = strings.Join(collectorIds, ",")

	if len(collectorIds) == 1 {
		d.SetId(ids)
	} else if len(collectorIds) > 1 {
		err = fmt.Errorf("Found more than 1 collector id matching this result, a device can only be matched to 1 collector, please change your device or choose a collector to add in tfvars.  %s", ids)
		return err
	} else {
		err = fmt.Errorf("Found no collectors with filter %s", filter)
		return err
	}
	return nil
}
