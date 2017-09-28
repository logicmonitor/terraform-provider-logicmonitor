package logicmonitor

import (
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	lmv1 "github.com/logicmonitor/lm-sdk-go"
)

func resourceDeviceGroup() *schema.Resource {
	return &schema.Resource{
		Create: addDeviceGroup,
		Read:   readDeviceGroup,
		Update: updateDeviceGroup,
		Delete: deleteDeviceGroup,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"parent_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"applies_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"disable_alerting": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"properties": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

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
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	restDeviceGroupResponse, apiResponse, e := client.GetDeviceGroupById(int32(id), "")
	err = checkStatus(restDeviceGroupResponse.Status, restDeviceGroupResponse.Errmsg, apiResponse.StatusCode, apiResponse.Status, e)
	if err != nil {
		log.Printf("Failed to find device group %q", err)
		d.SetId("")
		return nil
	}

	// known issue with v1 Go SDK, fix in v2
	if restDeviceGroupResponse.Data.ParentId == 1 {
		log.Printf("Host Group already at root level")
	} else {
		d.Set("parent_id", restDeviceGroupResponse.Data.ParentId)
	}

	d.Set("name", restDeviceGroupResponse.Data.Name)
	d.Set("description", restDeviceGroupResponse.Data.Description)
	d.Set("disable_alerting", restDeviceGroupResponse.Data.DisableAlerting)
	d.Set("applies_to", restDeviceGroupResponse.Data.AppliesTo)

	properties := make(map[string]string)
	props := restDeviceGroupResponse.Data.CustomProperties
	for _, v := range props {
		properties[v.Name] = v.Value
	}

	d.Set("properties", properties)
	return nil
}

// function to update device groups
func updateDeviceGroup(d *schema.ResourceData, m interface{}) error {
	d.Partial(true)
	client := m.(*lmv1.DefaultApi)
	device := makeDeviceGroupObject(d)
	device.CustomProperties = getProperties(d)

	// get Id
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	// list of available properties
	s := []string{"name", "parent_id", "applies_to", "disable_alerting", "description", "properties"}

	// loops through array of properties to see which one has changed, the ones that did not change are removed from the list
	for _, v := range s {
		if d.HasChange(v) {
		} else {
			s = remove(s, v)
		}
	}

	// makes a bulk update for all properties that were changed
	restDeviceGroupResponse, apiResponse, e := client.UpdateDeviceGroupById(int32(id), device)
	err = checkStatus(restDeviceGroupResponse.Status, restDeviceGroupResponse.Errmsg, apiResponse.StatusCode, apiResponse.Status, e)
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
	groupID, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	restDeviceGroupResponse, apiResponse, e := client.DeleteDeviceGroupById(int32(groupID), false)
	err = checkStatus(restDeviceGroupResponse.Status, restDeviceGroupResponse.Errmsg, apiResponse.StatusCode, apiResponse.Status, e)
	if err != nil {
		log.Printf("Error deleting device group %s", err)
		d.SetId("")
	}
	d.SetId("")
	return nil
}
