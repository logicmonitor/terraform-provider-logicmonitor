package logicmonitor

import (
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	lmv1 "github.com/logicmonitor/lm-sdk-go"
)

func resourceCollector() *schema.Resource {
	return &schema.Resource{
		Create: addCollector,
		Read:   readCollector,
		Update: updateCollector,
		Delete: deleteCollector,

		Schema: map[string]*schema.Schema{
			"backup_collector_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"collector_group_id": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enable_failback": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enable_collector_device_failover": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"escalation_chain_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"resend_interval": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"suppress_alert_clear": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

//resource function to add collector
func addCollector(d *schema.ResourceData, m interface{}) error {
	client := m.(*lmv1.DefaultApi)
	collector := makeDeviceCollectorObject(d)

	// calling API to add collector to portal
	log.Printf("Adding collector")
	restCollectorResponse, apiResponse, e := client.AddCollector(collector)
	err := checkStatus(restCollectorResponse.Status, restCollectorResponse.Errmsg, apiResponse.StatusCode, apiResponse.Status, e)
	if err != nil {
		return err
	}

	// calling API to save deviceId
	d.SetId(strconv.Itoa(int(restCollectorResponse.Data.Id)))
	return nil
}

// function to sync collector data
func readCollector(d *schema.ResourceData, m interface{}) error {
	client := m.(*lmv1.DefaultApi)

	// get collector id
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}

	restCollectorResponse, apiResponse, e := client.GetCollectorById(int32(id), "")
	err = checkStatus(restCollectorResponse.Status, restCollectorResponse.Errmsg, apiResponse.StatusCode, apiResponse.Status, e)
	if err != nil {
		log.Printf("Failed to find collector %q", err)
		d.SetId("")
		return nil
	}

	d.Set("backup_collector_id", restCollectorResponse.Data.BackupAgentId)
	d.Set("collector_group_id", restCollectorResponse.Data.CollectorGroupId)
	d.Set("description", restCollectorResponse.Data.Description)
	d.Set("enable_failback", restCollectorResponse.Data.EnableFailBack)
	d.Set("enable_collector_device_failover", restCollectorResponse.Data.EnableFailOverOnCollectorDevice)
	d.Set("escalation_chain_id", restCollectorResponse.Data.EscalatingChainId)
	d.Set("resend_interval", restCollectorResponse.Data.ResendIval)
	d.Set("suppress_alert_clear", restCollectorResponse.Data.SuppressAlertClear)

	return nil
}

// function to update collector data
func updateCollector(d *schema.ResourceData, m interface{}) error {
	d.Partial(true)
	client := m.(*lmv1.DefaultApi)
	collector := makeDeviceCollectorObject(d)

	// get collector id
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}

	// list of available properties
	s := []string{
		"backup_collector_id",
		"collector_group_id",
		"description",
		"enable_failback",
		"enable_collector_device_failover",
		"escalation_chain_id",
		"resend_interval",
		"suppress_alert_clear",
	}

	// loops through array of properties to see which one has changed, the ones that did not change are removed from the list
	for _, v := range s {
		if d.HasChange(v) {
		} else {
			s = remove(s, v)
		}
	}

	// makes a bulk update for all properties that were changed
	restCollectorResponse, apiResponse, e := client.UpdateCollectorById(int32(id), collector)
	err = checkStatus(restCollectorResponse.Status, restCollectorResponse.Errmsg, apiResponse.StatusCode, apiResponse.Status, e)
	if err != nil {
		return err
	}

	for _, v := range s {
		d.SetPartial(v)
	}

	d.Partial(false)
	return nil
}

// function to delete collector
func deleteCollector(d *schema.ResourceData, m interface{}) error {
	client := m.(*lmv1.DefaultApi)

	// get collector id
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}

	restCollectorResponse, apiResponse, e := client.DeleteCollectorById(int32(id))
	err = checkStatus(restCollectorResponse.Status, restCollectorResponse.Errmsg, apiResponse.StatusCode, apiResponse.Status, e)
	if err != nil {
		log.Printf("Error deleting collector %s", err)
		d.SetId("")
	}

	d.SetId("")
	return nil
}
