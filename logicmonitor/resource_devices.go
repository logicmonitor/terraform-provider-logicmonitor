package logicmonitor

import (
	"fmt"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/hashicorp/terraform/helper/schema"
	lmv1 "github.com/logicmonitor/lm-sdk-go"
)

// function to add a device
func addDevice(d *schema.ResourceData, m interface{}) error {
	client := m.(*lmv1.DefaultApi)
	device := makeDeviceObject(client, d)
	device.CustomProperties = getProperties(d)

	// calling API to add Device to portal
	log.Printf("Adding device %q", d.Get("ip_addr").(string))
	restResponse, apiResponse, e := client.AddDevice(device, false)
	err := checkStatus(restResponse.Status, restResponse.Errmsg, apiResponse.StatusCode, apiResponse.Status, e)
	if err != nil {
		return err
	}

	// calling API to save deviceId
	d.SetId(strconv.Itoa(int(restResponse.Data.Id)))
	return nil
}

// function to sync a device
func readDevice(d *schema.ResourceData, m interface{}) error {
	client := m.(*lmv1.DefaultApi)
	id, a := strconv.Atoi(d.Id())
	if a != nil {
		log.Error(a)
	}
	restResponse, apiResponse, e := client.GetDeviceById(int32(id), "")
	err := checkStatus(restResponse.Status, restResponse.Errmsg, apiResponse.StatusCode, apiResponse.Status, e)
	if err != nil {
		fmt.Printf("Failed to find collector group %q", err)
		d.SetId("")
		return nil
	}

	return nil
}

// function to update a device
func updateDevice(d *schema.ResourceData, m interface{}) error {
	d.Partial(true)
	client := m.(*lmv1.DefaultApi)
	device := makeDeviceObject(client, d)
	device.CustomProperties = getProperties(d)
	// get Id
	id, a := strconv.Atoi(d.Id())
	if a != nil {
		log.Error(a)
	}

	// list of available properties
	s := []string{"ip_addr", "display_name", "disable_alerting", "collector", "description", "hostgroup_id", "properties"}

	// loops through array of properties to see which one has changed, the ones that did not change are removed from the list
	for _, v := range s {
		if d.HasChange(v) {
		} else {
			s = remove(s, v)
		}
	}

	restResponse, apiResponse, e := client.UpdateDevice(device, int32(id), "refresh")
	err := checkStatus(restResponse.Status, restResponse.Errmsg, apiResponse.StatusCode, apiResponse.Status, e)
	if err != nil {
		return err
	}

	for _, v := range s {
		d.SetPartial(v)
	}

	d.Partial(false)
	return nil
}

// function to delete a device
func deleteDevice(d *schema.ResourceData, m interface{}) error {
	client := m.(*lmv1.DefaultApi)
	deviceID, a := strconv.Atoi(d.Id())
	if a != nil {
		log.Error(a)
	}
	restResponse, apiResponse, e := client.DeleteDevice(int32(deviceID))
	err := checkStatus(restResponse.Status, restResponse.Errmsg, apiResponse.StatusCode, apiResponse.Status, e)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func resourceDevices() *schema.Resource {
	return &schema.Resource{
		Create: addDevice,
		Read:   readDevice,
		Update: updateDevice,
		Delete: deleteDevice,

		Schema: map[string]*schema.Schema{
			"ip_addr": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"disable_alerting": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"collector": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"hostgroup_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
		},
	}
}
