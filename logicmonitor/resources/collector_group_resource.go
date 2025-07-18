//// Code generated by go-swagger; DO NOT EDIT.

package resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"terraform-provider-logicmonitor/client"
	"terraform-provider-logicmonitor/client/collector_group"
	"terraform-provider-logicmonitor/logicmonitor/schemata"
	"terraform-provider-logicmonitor/logicmonitor/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
CollectorGroup collector group API
*/

func CollectorGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: addCollectorGroup,
		DeleteContext: deleteCollectorGroupById,
		ReadContext:   getCollectorGroupById,
		UpdateContext: updateCollectorGroupById,
		Importer: &schema.ResourceImporter{
			State: resourceCollectorGroupStateImporter,
		},
		Schema: schemata.CollectorGroupSchema(),
	}
}

func DataResourceCollectorGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext: getCollectorGroupList,
		Schema:      schemata.DataSourceCollectorGroupSchema(),
	}
}

func addCollectorGroup(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := schemata.CollectorGroupModel(d)
	params := collector_group.NewAddCollectorGroupParams()

	params.SetBody(model)

	client := m.(*client.LogicMonitorRESTAPI)

	resp, err := client.CollectorGroup.AddCollectorGroup(params)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return diags
	}

	respModel := resp.GetPayload()
	schemata.SetCollectorGroupResourceData(d, respModel)
	d.SetId(strconv.Itoa(int(resp.Payload.ID)))

	return diags
}

func deleteCollectorGroupById(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := collector_group.NewDeleteCollectorGroupByIDParams()

	idVal, idIsSet := d.GetOk("id")
	if idIsSet {
		id, _ := strconv.Atoi(idVal.(string))
		params.ID = int32(id)
	} else {
		diags = append(diags, diag.Errorf("unexpected: Missing parameter - id")...)
		diags = append(diags, diag.Errorf("ending operation")...)
		return diags
	}

	client := m.(*client.LogicMonitorRESTAPI)

	resp, err := client.CollectorGroup.DeleteCollectorGroupByID(params)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return diags
	}

	d.SetId("")
	return diags
}

func getCollectorGroupById(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := collector_group.NewGetCollectorGroupByIDParams()

	fieldsVal, fieldsIsSet := d.GetOk("fields")
	if fieldsIsSet {
		params.Fields = fieldsVal.(*string)
	}

	idVal, idIsSet := d.GetOk("id")
	if idIsSet {
		id, _ := strconv.Atoi(idVal.(string))
		params.ID = int32(id)
	} else {
		diags = append(diags, diag.Errorf("unexpected: Missing parameter - id")...)
		diags = append(diags, diag.Errorf("ending operation")...)
		return diags
	}

	client := m.(*client.LogicMonitorRESTAPI)

	resp, err := client.CollectorGroup.GetCollectorGroupByID(params)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return diags
	}

	respModel := resp.GetPayload()
	schemata.SetCollectorGroupResourceData(d, respModel)
	return diags
}

func getCollectorGroupList(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := collector_group.NewGetCollectorGroupListParams()

	filterVal, filterIsSet := d.GetOk("filter")
	if filterIsSet {
		stringVal := filterVal.(string)
		params.Filter = &stringVal
	}

	client := m.(*client.LogicMonitorRESTAPI)

	resp, err := client.CollectorGroup.GetCollectorGroupList(params)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return diags
	}

	respModel := resp.GetPayload()
	if len(respModel.Items) == 0 {
		diags = append(diags, diag.Errorf("no CollectorGroup found")...)
	} else {
		result := respModel.Items[0]
		d.SetId(strconv.Itoa(int(result.ID)))
		schemata.SetCollectorGroupResourceData(d, result)
	}
	return diags
}

func updateCollectorGroupById(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	d.Partial(true)

	model := schemata.CollectorGroupModel(d)
	params := collector_group.NewUpdateCollectorGroupByIDParams()

	params.SetBody(model)

	forceUpdateFailedOverDevicesVal, forceUpdateFailedOverDevicesIsSet := d.GetOk("force_update_failed_over_devices")
	if forceUpdateFailedOverDevicesIsSet {
		forceUpdateFailedOverDevices := forceUpdateFailedOverDevicesVal.(bool)
		params.ForceUpdateFailedOverDevices = &forceUpdateFailedOverDevices
	}

	idVal, idIsSet := d.GetOk("id")
	if idIsSet {
		id, _ := strconv.Atoi(idVal.(string))
		params.ID = int32(id)
	} else {
		diags = append(diags, diag.Errorf("unexpected: Missing parameter - id")...)
		diags = append(diags, diag.Errorf("ending operation")...)
		return diags
	}

	// list of available properties
	props := schemata.GetCollectorGroupPropertyFields()

	// loops through array of properties to see which one has changed, the ones that did not change are removed from the list
	for _, v := range props {
		if d.HasChange(v) {
		} else {
			props = utils.Remove(props, v)
		}
	}

	client := m.(*client.LogicMonitorRESTAPI)

	// makes a bulk update for all properties that were changed
	resp, err := client.CollectorGroup.UpdateCollectorGroupByID(params)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return diags
	}

	respModel := resp.GetPayload()
	schemata.SetCollectorGroupResourceData(d, respModel)
	d.Partial(false)

	return diags
}

func resourceCollectorGroupStateImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	client := m.(*client.LogicMonitorRESTAPI)

	// if user provides an ID, we will add the collector group directly
	if utils.IsID(d.Id()) {
		id, err := strconv.Atoi(d.Id())
		if err != nil {
			return nil, err
		}

		params := collector_group.NewGetCollectorGroupByIDParams()
		params.SetID(int32(id))

		resp, err := client.CollectorGroup.GetCollectorGroupByID(params)
		log.Printf("[TRACE] response: %v", resp)
		if err != nil {
			log.Printf("Failed to find collector group %q", err)
			return nil, err
		}

		respModel := resp.GetPayload()
		schemata.SetCollectorGroupResourceData(d, respModel)
	} else {
		// find collector_group by name
		params := collector_group.NewGetCollectorGroupListParams()
		filter := fmt.Sprintf("name:\"%s\"", d.Id())
		params.SetFilter(&filter)

		resp, err := client.CollectorGroup.GetCollectorGroupList(params)
		log.Printf("[TRACE] response: %v", resp)
		if err != nil {
			err := fmt.Errorf("unexpected: %s", err)
			return nil, err
		}

		if resp.Payload.Total > 1 {
			err := fmt.Errorf("found more than 1 collector group with filter %s, please make the filter more specific or import with ID", filter)
			return nil, err
		} else if resp.Payload.Total == 1 {
			schemata.SetCollectorGroupResourceData(d, resp.Payload.Items[0])
			d.SetId(strconv.Itoa(int(resp.Payload.Items[0].ID)))
		} else {
			err := fmt.Errorf("found no collector groups with filter '%s'", filter)
			return nil, err
		}
	}
	return []*schema.ResourceData{d}, nil
}
