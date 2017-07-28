package logicmonitor

import (
	"fmt"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/hashicorp/terraform/helper/schema"
	lmv1 "github.com/logicmonitor/lm-sdk-go"
)

// function to add device groups
func addDeviceGroup(d *schema.ResourceData, m interface{}) error {
	log.Printf("Adding group %q", d.Get("name").(string))
	client := m.(*lmv1.DefaultApi)
	group := makeDeviceGroupObject(d)
	group.CustomProperties = getGroupProperties(d)

	// calling API to add Device to portal
	restDeviceGroupResponse, apiResponse, e := client.AddDeviceGroup(group)
	err := checkStatus(restDeviceGroupResponse.Status, restDeviceGroupResponse.Errmsg, apiResponse.StatusCode, apiResponse.Status, e)
	if err != nil {
		return err
	}

	// calling API to save groupId
	d.SetId(strconv.Itoa(int(restDeviceGroupResponse.Data.Id)))
	return nil
}

// function to sync device groups
func readDeviceGroup(d *schema.ResourceData, m interface{}) error {
	client := m.(*lmv1.DefaultApi)
	id, a := strconv.Atoi(d.Id())
	if a != nil {
		log.Error(a)
	}
	restDeviceGroupResponse, apiResponse, e := client.GetDeviceGroupById(int32(id), "")
	err := checkStatus(restDeviceGroupResponse.Status, restDeviceGroupResponse.Errmsg, apiResponse.StatusCode, apiResponse.Status, e)
	if err != nil {
		fmt.Printf("Failed to find device group %q", err)
		d.SetId("")
		return nil
	}

	return nil
}

// function to update device groups
func updateDeviceGroup(d *schema.ResourceData, m interface{}) error {
	d.Partial(true)
	client := m.(*lmv1.DefaultApi)
	device := makeDeviceGroupObject(d)
	device.CustomProperties = getProperties(d)

	// get Id
	id, _ := strconv.Atoi(d.Id())

	// list of available properties
	s := []string{"name", "parent_id", "applies_to", "disable_alerting", "description", "properties", "full_path"}

	// loops through array of properties to see which one has changed, the ones that did not change are removed from the list
	for _, v := range s {
		if d.HasChange(v) {
		} else {
			s = remove(s, v)
		}
	}

	// makes a bulk update for all properties that were changed
	restDeviceGroupResponse, apiResponse, e := client.UpdateDeviceGroupById(int32(id), device)
	err := checkStatus(restDeviceGroupResponse.Status, restDeviceGroupResponse.Errmsg, apiResponse.StatusCode, apiResponse.Status, e)
	if err != nil {
		return err
	}

	for _, v := range s {
		d.SetPartial(v)
	}

	d.Partial(false)
	return nil
}

// function to delete device groups
func deleteDeviceGroup(d *schema.ResourceData, m interface{}) error {
	client := m.(*lmv1.DefaultApi)
	groupID, a := strconv.Atoi(d.Id())
	if a != nil {
		log.Error(a)
	}
	restDeviceGroupResponse, apiResponse, e := client.DeleteDeviceGroupById(int32(groupID), false)
	err := checkStatus(restDeviceGroupResponse.Status, restDeviceGroupResponse.Errmsg, apiResponse.StatusCode, apiResponse.Status, e)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceDeviceGroup() *schema.Resource {
	return &schema.Resource{
		Create: addDeviceGroup,
		Read:   readDeviceGroup,
		Update: updateDeviceGroup,
		Delete: deleteDeviceGroup,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"parent_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"applies_to": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"disable_alerting": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"properties": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"full_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
