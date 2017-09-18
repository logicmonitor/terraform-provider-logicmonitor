package logicmonitor

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	lmv1 "github.com/logicmonitor/lm-sdk-go"
)

// terraform data source to look up device groups
func dataSourceFindDeviceGroups() *schema.Resource {
	return &schema.Resource{
		Read: findDeviceGroup,

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

// function to find device groups with certain filters
func findDeviceGroup(d *schema.ResourceData, m interface{}) error {
	client := m.(*lmv1.DefaultApi)
	filter := getFilters(d)
	size := int32(d.Get("size").(int))
	offset := int32(d.Get("offset").(int))

	//looks for device Group
	restDeviceGroupPaginationResponse, apiResponse, e := client.GetDeviceGroupList("name,id,customProperties", size, offset, filter)
	err := checkStatus(restDeviceGroupPaginationResponse.Status, restDeviceGroupPaginationResponse.Errmsg, apiResponse.StatusCode, apiResponse.Status, e)
	if err != nil {
		return err
	}

	var deviceIds []string
	for _, d := range restDeviceGroupPaginationResponse.Data.Items {
		log.Printf("Found device group with filter %q", filter)
		deviceIds = append(deviceIds, strconv.Itoa(int(d.Id)))
	}

	if len(deviceIds) > 0 {
		//comma separated string of device Ids
		var ids = strings.Join(deviceIds, ",")
		d.SetId(ids)
	} else {
		err := fmt.Errorf("Found no device groups with filter %s", filter)
		return err
	}
	return nil
}
