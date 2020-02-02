package logicmonitor

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	lmclient "github.com/logicmonitor/lm-sdk-go/client"
	"github.com/logicmonitor/lm-sdk-go/client/lm"
)

func resourceDashboardGroup() *schema.Resource {
	return &schema.Resource{
		Create: addDashboardGroup,
		Read:   readDashboardGroup,
		Update: updateDashboardGroup,
		Delete: deleteDashboardGroup,
		Importer: &schema.ResourceImporter{
			State: resourceDashboardGroupStateImporter,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"parent_id": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"force_delete": {
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

// function to add dashboard groups
func addDashboardGroup(d *schema.ResourceData, m interface{}) error {
	log.Printf("Adding group %q", d.Get("name").(string))
	client := m.(*lmclient.LMSdkGo)
	params := lm.NewAddDashboardGroupParams()
	group := makeDashboardGroupObject(d)
	params.SetBody(&group)

	// calling API to add Device to portal
	restDashboardGroupResponse, err := client.LM.AddDashboardGroup(params)
	if err != nil {
		return err
	}

	// calling API to save groupId
	d.SetId(strconv.Itoa(int(restDashboardGroupResponse.Payload.ID)))
	return nil
}

// function to sync dashboard groups
func readDashboardGroup(d *schema.ResourceData, m interface{}) error {
	client := m.(*lmclient.LMSdkGo)
	id, err := strconv.Atoi(d.Id())
	t := true
	if err != nil {
		return err
	}
	params := lm.NewGetDashboardGroupByIDParams()
	params.SetID(int32(id))
	params.Template = &t
	restDashboardGroupResponse, err := client.LM.GetDashboardGroupByID(params)
	if err != nil {
		log.Printf("Failed to find device group %q", err)
		d.SetId("")
		return nil
	}
	log.Printf("JSON LOGS")
	log.Printf("%s", restDashboardGroupResponse.Payload.Template)
	d.Set("parent_id", restDashboardGroupResponse.Payload.ParentID)
	d.Set("name", restDashboardGroupResponse.Payload.Name)
	d.Set("description", restDashboardGroupResponse.Payload.Description)
	// until we fix template output
	//template := []byte(restDashboardGroupResponse.Payload.Template.(string))
	//var jsonOutput map[string]interface{}
	//json.Unmarshal(template, &jsonOutput)
	//d.Set("template", restDashboardGroupResponse.Payload.Template)
	properties := make(map[string]string)
	props := restDashboardGroupResponse.Payload.WidgetTokens
	for _, v := range props {
		properties[v.Name] = v.Value
	}

	d.Set("widget_tokens", properties)
	return nil
}

// function to update dashboard groups
func updateDashboardGroup(d *schema.ResourceData, m interface{}) error {
	deleteDashboardGroup(d, m)
	addDashboardGroup(d, m)
	/*
		d.Partial(true)
		client := m.(*lmclient.LMSdkGo)

		// get Id
		id, err := strconv.Atoi(d.Id())
		if err != nil {
			return err
		}

		device := makeDashboardGroupObject(d)
		params := lm.NewUpdateDashboardGroupByIDParams()
		params.SetID(int32(id))
		params.SetBody(&device)

		// list of available properties
		s := []string{"name", "parent_id", "widget_tokens", "description", "template"}

		// loops through array of properties to see which one has changed, the ones that did not change are removed from the list
		for _, v := range s {
			if d.HasChange(v) {
			} else {
				s = remove(s, v)
			}
		}

		// makes a bulk update for all properties that were changed
		nil, err := client.LM.UpdateDashboardGroupByID(params)
		if err != nil {
			return err
		}

		for _, v := range s {
			d.SetPartial(v)
		}

		d.Partial(false)
	*/
	return nil
}

// function to delete dashboard groups
func deleteDashboardGroup(d *schema.ResourceData, m interface{}) error {
	client := m.(*lmclient.LMSdkGo)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	forceDelete := d.Get("force_delete").(bool)
	params := lm.NewDeleteDashboardGroupByIDParams()
	params.SetID(int32(id))
	params.SetAllowNonEmptyGroup(&forceDelete)

	restDashboardGroupResponse, err := client.LM.DeleteDashboardGroupByID(params)
	if err != nil {
		log.Printf("Error deleting dashboard group %s", err)
		d.SetId("")
	}

	log.Printf("dashboard group payload response %v", restDashboardGroupResponse.Payload)
	d.SetId("")

	return nil
}

// function to import dashboard groups
func resourceDashboardGroupStateImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	client := m.(*lmclient.LMSdkGo)

	// if it is a groupId, we will add the group directly, currently there is no way to see this in the UI.
	if checkID(d.Id()) {
		id, err := strconv.Atoi(d.Id())
		if err != nil {
			return nil, err
		}
		params := lm.NewGetDashboardGroupByIDParams()
		restDashboardGroupResponse, err := client.LM.GetDashboardGroupByID(params)
		if err != nil {
			log.Printf("Failed to find dashboard group %q", err)
			return nil, err
		}
		d.Set("name", restDashboardGroupResponse.Payload.Name)
		d.Set("description", restDashboardGroupResponse.Payload.Description)
		d.Set("parent_id", restDashboardGroupResponse.Payload.ParentID)
		properties := make(map[string]string)
		props := restDashboardGroupResponse.Payload.WidgetTokens
		for _, v := range props {
			properties[v.Name] = v.Value
		}
		d.Set("widget_tokens", properties)
		params.SetID(int32(id))
	} else {

		// currently I just set it to add FullPath, I can give it the option later of specifying entire sets of groups if needed
		filter := fmt.Sprintf("fullPath:\"%s\"", d.Id())
		params := lm.NewGetDashboardGroupListParams()
		params.SetFilter(&filter)
		//looks for dashboard group
		restDashboardGroupPaginationResponse, err := client.LM.GetDashboardGroupList(params)
		if err != nil {
			return nil, err
		}

		if restDashboardGroupPaginationResponse.Payload.Total > 0 {
			d.Set("name", restDashboardGroupPaginationResponse.Payload.Items[0].Name)
			d.Set("description", restDashboardGroupPaginationResponse.Payload.Items[0].Description)
			d.Set("parent_id", restDashboardGroupPaginationResponse.Payload.Items[0].ParentID)
			d.SetId(strconv.Itoa(int(restDashboardGroupPaginationResponse.Payload.Items[0].ID)))
		} else {
			err := fmt.Errorf("Found no dashboard groups with filter %s", filter)
			return nil, err
		}
	}
	return []*schema.ResourceData{d}, nil
}
