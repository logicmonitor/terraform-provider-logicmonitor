package logicmonitor

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	lmclient "github.com/logicmonitor/lm-sdk-go/client"
	"github.com/logicmonitor/lm-sdk-go/client/lm"
)

func resourceCollectorGroup() *schema.Resource {
	return &schema.Resource{
		Create: addCollectorGroup,
		Read:   readCollectorGroup,
		Update: updateCollectorGroup,
		Delete: deleteCollectorGroup,
		Importer: &schema.ResourceImporter{
			State: resourceCollectorGroupStateImporter,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
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

//resource function to add collector group
func addCollectorGroup(d *schema.ResourceData, m interface{}) error {
	client := m.(*lmclient.LMSdkGo)
	cgroup := makeDeviceCollectorGroupObject(d)
	params := lm.NewAddCollectorGroupParams()
	params.SetBody(&cgroup)

	// calling API to add Device to portal
	log.Printf("Adding collector group")
	restCollectorGroupResponse, err := client.LM.AddCollectorGroup(params)
	if err != nil {
		return err
	}

	// calling API to save deviceId
	d.SetId(strconv.Itoa(int(restCollectorGroupResponse.Payload.ID)))
	return nil
}

// function to sync collector group data
func readCollectorGroup(d *schema.ResourceData, m interface{}) error {
	client := m.(*lmclient.LMSdkGo)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	params := lm.NewGetCollectorGroupByIDParams()
	params.ID = (int32(id))

	restCollectorGroupResponse, err := client.LM.GetCollectorGroupByID(params)
	if err != nil {
		log.Printf("Failed to find collector group %q", err)
		d.SetId("")
		return nil
	}

	d.Set("name", restCollectorGroupResponse.Payload.Name)
	d.Set("description", restCollectorGroupResponse.Payload.Description)

	properties := make(map[*string]*string)
	props := restCollectorGroupResponse.Payload.CustomProperties
	for _, v := range props {
		properties[v.Name] = v.Value
	}
	d.Set("properties", properties)

	return nil
}

// function to update collector group data
func updateCollectorGroup(d *schema.ResourceData, m interface{}) error {
	d.Partial(true)
	client := m.(*lmclient.LMSdkGo)
	group := makeDeviceCollectorGroupObject(d)
	// get Id
	id, a := strconv.Atoi(d.Id())
	if a != nil {
		return a
	}

	// list of available properties
	s := []string{"name", "description", "properties"}

	// loops through array of properties to see which one has changed, the ones that did not change are removed from the list
	for _, v := range s {
		if d.HasChange(v) {
		} else {
			s = remove(s, v)
		}
	}

	params := lm.NewUpdateCollectorGroupByIDParams()
	params.ID = int32(id)
	params.SetBody(&group)

	// makes a bulk update for all properties that were changed
	restCollectorGroupResponse, err := client.LM.UpdateCollectorGroupByID(params)
	if err != nil {
		return err
	}

	for _, v := range s {
		d.SetPartial(v)
	}

	log.Printf("payload response %v", restCollectorGroupResponse.Payload)
	d.Partial(false)
	return nil
}

// function to delete collector groups
func deleteCollectorGroup(d *schema.ResourceData, m interface{}) error {
	client := m.(*lmclient.LMSdkGo)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}

	params := lm.NewDeleteCollectorGroupByIDParams()
	params.ID = int32(id)
	restCollectorGroupResponse, err := client.LM.DeleteCollectorGroupByID(params)

	if err != nil {
		log.Printf("Error deleting collector group %s", err)
		d.SetId("")
	}

	log.Printf("payload response %v", restCollectorGroupResponse.Payload)
	d.SetId("")
	return nil
}

func resourceCollectorGroupStateImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	client := m.(*lmclient.LMSdkGo)

	// if it is a groupId, we will add the group directly
	if checkID(d.Id()) {
		id, err := strconv.Atoi(d.Id())
		if err != nil {
			return nil, err
		}
		params := lm.NewGetCollectorGroupByIDParams()
		params.SetID(int32(id))
		restCollectorGroupResponse, err := client.LM.GetCollectorGroupByID(params)
		if err != nil {
			log.Printf("Failed to find collector group %q", err)
			return nil, err
		}
		d.Set("name", restCollectorGroupResponse.Payload.Name)
		d.Set("description", restCollectorGroupResponse.Payload.Description)

		properties := make(map[*string]*string)
		props := restCollectorGroupResponse.Payload.CustomProperties
		for _, v := range props {
			properties[v.Name] = v.Value
		}
		d.Set("properties", properties)
		d.SetId(d.Id())
	} else {
		// finding collector group by name
		filter := fmt.Sprintf("name:\"%s\"", d.Id())
		params := lm.NewGetCollectorGroupListParams()
		params.SetFilter(&filter)

		//looks for collector group
		restCollectorGroupPaginationResponse, err := client.LM.GetCollectorGroupList(params)
		if err != nil {
			return nil, err
		}

		if restCollectorGroupPaginationResponse.Payload.Total > 1 {
			err := fmt.Errorf("Found more than 1 group with filter %s, please make the filter more specific or import with Id", filter)
			return nil, err
		} else if restCollectorGroupPaginationResponse.Payload.Total == 1 {
			d.Set("name", restCollectorGroupPaginationResponse.Payload.Items[0].Name)
			d.Set("description", restCollectorGroupPaginationResponse.Payload.Items[0].Description)
			properties := make(map[*string]*string)
			props := restCollectorGroupPaginationResponse.Payload.Items[0].CustomProperties
			for _, v := range props {
				properties[v.Name] = v.Value
			}
			d.Set("properties", properties)
			d.SetId(strconv.Itoa(int(restCollectorGroupPaginationResponse.Payload.Items[0].ID)))
		} else {
			err := fmt.Errorf("Found no collector groups with filter %s", filter)
			return nil, err
		}
	}
	return []*schema.ResourceData{d}, nil
}
