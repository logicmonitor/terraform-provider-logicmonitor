package logicmonitor

import (
	"fmt"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/hashicorp/terraform/helper/schema"
	lmv1 "github.com/logicmonitor/lm-sdk-go"
)

//resource function to add collector group
func addCollectorGroup(d *schema.ResourceData, m interface{}) error {
	client := m.(*lmv1.DefaultApi)
	device := makeDeviceCollectorGroupObject(d)

	// calling API to add Device to portal
	log.Printf("Adding collector group")
	restCollectorGroupResponse, apiResponse, e := client.AddCollectorGroup(device)
	err := checkStatus(restCollectorGroupResponse.Status, restCollectorGroupResponse.Errmsg, apiResponse.StatusCode, apiResponse.Status, e)
	if err != nil {
		return err
	}

	// calling API to save deviceId
	d.SetId(strconv.Itoa(int(restCollectorGroupResponse.Data.Id)))
	return nil
}

// function to sync collector group data
func readCollectorGroup(d *schema.ResourceData, m interface{}) error {
	client := m.(*lmv1.DefaultApi)
	id, a := strconv.Atoi(d.Id())
	if a != nil {
		log.Error(a)
	}
	restCollectorGroupResponse, apiResponse, e := client.GetCollectorGroupById(int32(id), "")
	err := checkStatus(restCollectorGroupResponse.Status, restCollectorGroupResponse.Errmsg, apiResponse.StatusCode, apiResponse.Status, e)
	if err != nil {
		fmt.Printf("Failed to find collector group %q", err)
		d.SetId("")
		return nil
	}
	return nil
}

// function to update collector group data
func updateCollectorGroup(d *schema.ResourceData, m interface{}) error {
	d.Partial(true)
	client := m.(*lmv1.DefaultApi)
	group := makeDeviceCollectorGroupObject(d)
	// get Id
	id, a := strconv.Atoi(d.Id())
	if a != nil {
		log.Error(a)
	}

	// list of available properties
	s := []string{"name", "description"}

	// loops through array of properties to see which one has changed, the ones that did not change are removed from the list
	for _, v := range s {
		if d.HasChange(v) {
		} else {
			s = remove(s, v)
		}
	}

	// makes a bulk update for all properties that were changed
	restCollectorGroupResponse, apiResponse, e := client.UpdateCollectorGroupById(int32(id), group)
	err := checkStatus(restCollectorGroupResponse.Status, restCollectorGroupResponse.Errmsg, apiResponse.StatusCode, apiResponse.Status, e)
	if err != nil {
		return err
	}

	for _, v := range s {
		d.SetPartial(v)
	}

	d.Partial(false)
	return nil
}

// function to delete collector groups
func deleteCollectorGroup(d *schema.ResourceData, m interface{}) error {
	client := m.(*lmv1.DefaultApi)
	groupID, a := strconv.Atoi(d.Id())
	if a != nil {
		log.Error(a)
	}
	restCollectorGroupResponse, apiResponse, e := client.DeleteCollectorGroupById(int32(groupID))
	err := checkStatus(restCollectorGroupResponse.Status, restCollectorGroupResponse.Errmsg, apiResponse.StatusCode, apiResponse.Status, e)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func resourceCollectorGroup() *schema.Resource {
	return &schema.Resource{
		Create: addCollectorGroup,
		Read:   readCollectorGroup,
		Update: updateCollectorGroup,
		Delete: deleteCollectorGroup,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
