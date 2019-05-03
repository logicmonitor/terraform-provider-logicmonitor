package logicmonitor

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	lmclient "github.com/logicmonitor/lm-sdk-go/client"
	"github.com/logicmonitor/lm-sdk-go/client/lm"
)

func resourceDevices() *schema.Resource {
	return &schema.Resource{
		Create: addDevice,
		Read:   readDevice,
		Update: updateDevice,
		Delete: deleteDevice,
		Importer: &schema.ResourceImporter{
			State: resourceDeviceStateImporter,
		},

		Schema: map[string]*schema.Schema{
			"ip_addr": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"disable_alerting": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"collector": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"hostgroup_id": {
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

// function to add a device
func addDevice(d *schema.ResourceData, m interface{}) error {
	client := m.(*lmclient.LMSdkGo)
	params := lm.NewAddDeviceParams()
	device := makeDeviceObject(d)
	params.SetBody(&device)
	// calling API to add Device to portal
	log.Printf("Adding device %q", d.Get("ip_addr").(string))
	restResponse, err := client.LM.AddDevice(params)
	if err != nil {
		return err
	}
	// calling API to save deviceId
	d.SetId(strconv.Itoa(int(restResponse.Payload.ID)))
	return nil
}

// function to sync a device
func readDevice(d *schema.ResourceData, m interface{}) error {
	client := m.(*lmclient.LMSdkGo)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	params := lm.NewGetDeviceByIDParams()
	params.SetID(int32(id))

	restResponse, err := client.LM.GetDeviceByID(params)
	if err != nil {
		log.Printf("Failed to find device %q", err)
		d.SetId("")
		return nil
	}

	if *restResponse.Payload.DisplayName == *restResponse.Payload.Name || restResponse.Payload.DisplayName == nil {
		log.Printf("Don't need to set displayname")
		//d.Set("display_name", restResponse.Payload.Name)
	} else {
		d.Set("display_name", restResponse.Payload.DisplayName)
	}

	d.Set("ip_addr", restResponse.Payload.Name)
	d.Set("collector", restResponse.Payload.PreferredCollectorID)
	d.Set("disable_alerting", restResponse.Payload.DisableAlerting)
	d.Set("description", restResponse.Payload.Description)
	d.Set("hostgroup_id", restResponse.Payload.HostGroupIds)

	properties := make(map[*string]*string)
	props := restResponse.Payload.CustomProperties
	for _, v := range props {
		properties[v.Name] = v.Value
	}
	d.Set("properties", properties)

	return nil
}

// function to update a device
func updateDevice(d *schema.ResourceData, m interface{}) error {
	d.Partial(true)
	client := m.(*lmclient.LMSdkGo)

	// get Id
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return errors.New(err.Error())
	}

	params := lm.NewUpdateDeviceParams()
	device := makeDeviceObject(d)
	device.CustomProperties = getProperties(d)
	params.SetBody(&device)
	params.SetID(int32(id))

	// list of available properties
	s := []string{"ip_addr", "display_name", "disable_alerting", "collector", "description", "hostgroup_id", "properties"}

	// loops through array of properties to see which one has changed, the ones that did not change are removed from the list
	for _, v := range s {
		if d.HasChange(v) {
		} else {
			s = remove(s, v)
		}
	}

	restResponse, err := client.LM.UpdateDevice(params)
	if err != nil {
		return err
	}
	log.Printf("update device response %v", restResponse.Payload)

	for _, v := range s {
		d.SetPartial(v)
	}

	d.Partial(false)
	return nil
}

// function to delete a device
func deleteDevice(d *schema.ResourceData, m interface{}) error {
	client := m.(*lmclient.LMSdkGo)
	deviceID, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	params := lm.NewDeleteDeviceByIDParams()
	params.SetID(int32(deviceID))
	restResponse, err := client.LM.DeleteDeviceByID(params)
	if err != nil {
		log.Printf("Error deleting device %s", err)
		d.SetId("")
	}

	log.Printf("delete device response %v", restResponse.Payload)

	d.SetId("")
	return nil
}

// function to import devices
func resourceDeviceStateImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	client := m.(*lmclient.LMSdkGo)

	// if it is a deviceId, we will add the group directly
	if checkID(d.Id()) {
		id, err := strconv.Atoi(d.Id())
		if err != nil {
			return nil, err
		}
		params := lm.NewGetDeviceByIDParams()
		params.SetID(int32(id))

		restResponse, err := client.LM.GetDeviceByID(params)
		if err != nil {
			log.Printf("Failed to find device %q", err)
			return nil, err
		}
		d.Set("ip_addr", restResponse.Payload.Name)
		d.Set("collector", restResponse.Payload.PreferredCollectorID)
		d.Set("disable_alerting", restResponse.Payload.DisableAlerting)
		d.Set("description", restResponse.Payload.Description)
		d.Set("hostgroup_id", restResponse.Payload.HostGroupIds)
		properties := make(map[*string]*string)
		props := restResponse.Payload.CustomProperties
		for _, v := range props {
			properties[v.Name] = v.Value
		}
		d.Set("properties", properties)
	} else {

		// currently set to add displayname
		filter := fmt.Sprintf("displayName:\"%s\"", d.Id())
		//looks for device
		params := lm.NewGetDeviceListParams()
		params.SetFilter(&filter)
		restResponse, err := client.LM.GetDeviceList(params)
		if err != nil {
			return nil, err
		}
		if restResponse.Payload.Total > 1 {
			err := fmt.Errorf("Found more than 1 device with filter %s, please make the filter more specific or import with deviceId", filter)
			return nil, err
		} else if restResponse.Payload.Total > 0 {
			d.Set("ip_addr", restResponse.Payload.Items[0].Name)
			d.Set("display_name", restResponse.Payload.Items[0].DisplayName)
			d.Set("collector", restResponse.Payload.Items[0].PreferredCollectorID)
			d.Set("disable_alerting", restResponse.Payload.Items[0].DisableAlerting)
			d.SetId(strconv.Itoa(int(restResponse.Payload.Items[0].ID)))
			properties := make(map[*string]*string)
			props := restResponse.Payload.Items[0].CustomProperties
			for _, v := range props {
				properties[v.Name] = v.Value
			}
			d.Set("properties", properties)
		} else {
			err := fmt.Errorf("Found no device with filter %s", filter)
			return nil, err
		}
	}
	return []*schema.ResourceData{d}, nil
}
