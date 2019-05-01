package logicmonitor

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	lmclient "github.com/logicmonitor/lm-sdk-go/client"
	"github.com/logicmonitor/lm-sdk-go/client/lm"
)

func resourceDeviceGroup() *schema.Resource {
	return &schema.Resource{
		Create: addDeviceGroup,
		Read:   readDeviceGroup,
		Update: updateDeviceGroup,
		Delete: deleteDeviceGroup,
		Importer: &schema.ResourceImporter{
			State: resourceDeviceGroupStateImporter,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"parent_id": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
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
	client := m.(*lmclient.LMSdkGo)
	params := lm.NewAddDeviceGroupParams()
	group := makeDeviceGroupObject(d)
	params.SetBody(&group)

	// calling API to add Device to portal
	restDeviceGroupResponse, err := client.LM.AddDeviceGroup(params)
	if err != nil {
		return err
	}

	// calling API to save groupId
	d.SetId(strconv.Itoa(int(restDeviceGroupResponse.Payload.ID)))
	return nil
}

// function to sync device groups
func readDeviceGroup(d *schema.ResourceData, m interface{}) error {
	client := m.(*lmclient.LMSdkGo)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	params := lm.NewGetDeviceGroupByIDParams()
	params.SetID(int32(id))

	restDeviceGroupResponse, err := client.LM.GetDeviceGroupByID(params)
	if err != nil {
		log.Printf("Failed to find device group %q", err)
		d.SetId("")
		return nil
	}

	d.Set("parent_id", restDeviceGroupResponse.Payload.ParentID)
	d.Set("name", restDeviceGroupResponse.Payload.Name)
	d.Set("description", restDeviceGroupResponse.Payload.Description)
	d.Set("disable_alerting", restDeviceGroupResponse.Payload.DisableAlerting)
	d.Set("applies_to", restDeviceGroupResponse.Payload.AppliesTo)

	properties := make(map[*string]*string)
	props := restDeviceGroupResponse.Payload.CustomProperties
	for _, v := range props {
		properties[v.Name] = v.Value
	}

	d.Set("properties", properties)
	return nil
}

// function to update device groups
func updateDeviceGroup(d *schema.ResourceData, m interface{}) error {
	d.Partial(true)
	client := m.(*lmclient.LMSdkGo)

	// get Id
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}

	device := makeDeviceGroupObject(d)
	params := lm.NewUpdateDeviceGroupByIDParams()
	params.SetID(int32(id))
	params.SetBody(&device)

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
	nil, err := client.LM.UpdateDeviceGroupByID(params)
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
	client := m.(*lmclient.LMSdkGo)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}

	params := lm.NewDeleteDeviceGroupByIDParams()
	params.SetID(int32(id))

	restDeviceGroupResponse, err := client.LM.DeleteDeviceGroupByID(params)
	if err != nil {
		log.Printf("Error deleting device group %s", err)
		d.SetId("")
	}
	log.Printf("payload response %q", restDeviceGroupResponse.Payload)
	d.SetId("")
	return nil
}

// function to import device groups
func resourceDeviceGroupStateImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	client := m.(*lmclient.LMSdkGo)

	// if it is a groupId, we will add the group directly
	if checkID(d.Id()) {
		id, err := strconv.Atoi(d.Id())
		if err != nil {
			return nil, err
		}
		params := lm.NewGetDeviceGroupByIDParams()
		params.SetID(int32(id))
		restDeviceGroupResponse, err := client.LM.GetDeviceGroupByID(params)
		if err != nil {
			log.Printf("Failed to find device group %q", err)
			return nil, err
		}
		d.Set("disable_alerting", restDeviceGroupResponse.Payload.DisableAlerting)
		d.Set("name", restDeviceGroupResponse.Payload.Name)
		d.SetId(d.Id())
	} else {

		// currently I just set it to add FullPath, I can give it the option later of specifying entire sets of groups if needed
		filter := fmt.Sprintf("fullPath:\"%s\"", d.Id())
		params := lm.NewGetDeviceGroupListParams()
		params.SetFilter(&filter)
		//looks for device Group
		restDeviceGroupPaginationResponse, err := client.LM.GetDeviceGroupList(params)
		if err != nil {
			return nil, err
		}

		if restDeviceGroupPaginationResponse.Payload.Total > 0 {
			d.Set("disable_alerting", restDeviceGroupPaginationResponse.Payload.Items[0].DisableAlerting)
			d.Set("name", restDeviceGroupPaginationResponse.Payload.Items[0].Name)
			d.SetId(strconv.Itoa(int(restDeviceGroupPaginationResponse.Payload.Items[0].ID)))
		} else {
			err := fmt.Errorf("Found no device groups with filter %s", filter)
			return nil, err
		}
	}
	return []*schema.ResourceData{d}, nil
}
