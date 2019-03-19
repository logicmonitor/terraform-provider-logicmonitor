package logicmonitor

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	lmclient "github.com/logicmonitor/lm-sdk-go/client"
	"github.com/logicmonitor/lm-sdk-go/client/lm"
)

func resourceDashboard() *schema.Resource {
	return &schema.Resource{
		Create: addDashboard,
		Read:   readDashboard,
		Update: updateDashboard,
		Delete: deleteDashboard,
		Importer: &schema.ResourceImporter{
			State: resourceDashboardStateImporter,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"group_id": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"public": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"template": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"widget_tokens": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

// function to add dashboards
func addDashboard(d *schema.ResourceData, m interface{}) error {
	log.Printf("Adding group %q", d.Get("name").(string))
	client := m.(*lmclient.LMSdkGo)
	params := lm.NewAddDashboardParams()
	group := makeDashboardObject(d)
	params.SetBody(&group)

	// calling API to add Device to portal
	restDashboardResponse, err := client.LM.AddDashboard(params)
	if err != nil {
		return err
	}

	// calling API to save groupId
	d.SetId(strconv.Itoa(int(restDashboardResponse.Payload.ID)))
	return nil
}

// function to sync dashboards
func readDashboard(d *schema.ResourceData, m interface{}) error {
	client := m.(*lmclient.LMSdkGo)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	params := lm.NewGetDashboardByIDParams()
	params.SetID(int32(id))
	restDashboardResponse, err := client.LM.GetDashboardByID(params)
	if err != nil {
		log.Printf("Failed to find device group %q", err)
		d.SetId("")
		return nil
	}

	d.Set("group_id", restDashboardResponse.Payload.GroupID)
	d.Set("name", restDashboardResponse.Payload.Name)
	d.Set("description", restDashboardResponse.Payload.Description)
	// until we fix template output
	//	d.Set("template", restDashboardResponse.Payload.Template)
	properties := make(map[string]string)
	props := restDashboardResponse.Payload.WidgetTokens
	for _, v := range props {
		properties[v.Name] = v.Value
	}

	d.Set("widget_tokens", properties)
	return nil
}

// function to update dashboards
func updateDashboard(d *schema.ResourceData, m interface{}) error {
	d.Partial(true)
	client := m.(*lmclient.LMSdkGo)

	// get Id
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}

	device := makeDashboardObject(d)
	params := lm.NewUpdateDashboardByIDParams()
	params.SetID(int32(id))
	params.SetBody(&device)

	// list of available properties
	s := []string{"name", "parent_id", "widget_tokens", "description"}

	// loops through array of properties to see which one has changed, the ones that did not change are removed from the list
	for _, v := range s {
		if d.HasChange(v) {
		} else {
			s = remove(s, v)
		}
	}

	// makes a bulk update for all properties that were changed
	nil, err := client.LM.UpdateDashboardByID(params)
	if err != nil {
		return err
	}

	for _, v := range s {
		d.SetPartial(v)
	}

	d.Partial(false)
	return nil
}

// function to delete dashboards
func deleteDashboard(d *schema.ResourceData, m interface{}) error {
	client := m.(*lmclient.LMSdkGo)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	params := lm.NewDeleteDashboardByIDParams()
	params.SetID(int32(id))

	restDashboardResponse, err := client.LM.DeleteDashboardByID(params)
	if err != nil {
		log.Printf("Error deleting dashboard group %s", err)
		d.SetId("")
	}

	log.Printf("dashboard group payload response %v", restDashboardResponse.Payload)
	d.SetId("")

	return nil
}

// function to import dashboard groups
func resourceDashboardStateImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	client := m.(*lmclient.LMSdkGo)

	// if it is a groupId, we will add the group directly, currently there is no way to see this in the UI.
	if checkID(d.Id()) {
		id, err := strconv.Atoi(d.Id())
		if err != nil {
			return nil, err
		}
		params := lm.NewGetDashboardByIDParams()
		restDashboardResponse, err := client.LM.GetDashboardByID(params)
		if err != nil {
			log.Printf("Failed to find dashboard group %q", err)
			return nil, err
		}
		d.Set("name", restDashboardResponse.Payload.Name)
		d.Set("description", restDashboardResponse.Payload.Description)
		d.Set("group_id", restDashboardResponse.Payload.GroupID)
		properties := make(map[string]string)
		props := restDashboardResponse.Payload.WidgetTokens
		for _, v := range props {
			properties[v.Name] = v.Value
		}
		d.Set("widget_tokens", properties)
		params.SetID(int32(id))
	} else {

		// currently I just set it to add FullPath, I can give it the option later of specifying entire sets of groups if needed
		filter := fmt.Sprintf("fullPath:\"%s\"", d.Id())
		params := lm.NewGetDashboardListParams()
		params.SetFilter(&filter)
		//looks for device Group
		restDashboardPaginationResponse, err := client.LM.GetDashboardList(params)
		if err != nil {
			return nil, err
		}

		if restDashboardPaginationResponse.Payload.Total > 0 {
			d.Set("name", restDashboardPaginationResponse.Payload.Items[0].Name)
			d.Set("description", restDashboardPaginationResponse.Payload.Items[0].Description)
			d.Set("group_id", restDashboardPaginationResponse.Payload.Items[0].GroupID)
			d.SetId(strconv.Itoa(int(restDashboardPaginationResponse.Payload.Items[0].ID)))
		} else {
			err := fmt.Errorf("Found no dashboards with filter %s", filter)
			return nil, err
		}
	}
	return []*schema.ResourceData{d}, nil
}
