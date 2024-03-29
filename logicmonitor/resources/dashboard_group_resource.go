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
	"terraform-provider-logicmonitor/client/dashboard_group"
	"terraform-provider-logicmonitor/logicmonitor/schemata"
	"terraform-provider-logicmonitor/logicmonitor/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
DashboardGroup dashboard group API
*/

func DashboardGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: addDashboardGroup,
		DeleteContext: deleteDashboardGroupById,
		ReadContext:   getDashboardGroupById,
		UpdateContext: updateDashboardGroupById,
		Importer: &schema.ResourceImporter{
			State: resourceDashboardGroupStateImporter,
		},
		Schema: schemata.DashboardGroupSchema(),
	}
}

func DataResourceDashboardGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext: getDashboardGroupList,
		Schema:      schemata.DataSourceDashboardGroupSchema(),
	}
}

func addDashboardGroup(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := schemata.DashboardGroupModel(d)
	params := dashboard_group.NewAddDashboardGroupParams()
	params.SetBody(model)

	client := m.(*client.LogicMonitorRESTAPI)

	resp, err := client.DashboardGroup.AddDashboardGroup(params)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return diags
	}

	respModel := resp.GetPayload()
	schemata.SetDashboardGroupResourceData(d, respModel)
	d.SetId(strconv.Itoa(int(resp.Payload.ID)))

	return diags
}

func deleteDashboardGroupById(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := dashboard_group.NewDeleteDashboardGroupByIDParams()

	allowNonEmptyGroupVal, allowNonEmptyGroupIsSet := d.GetOk("allow_non_empty_group")
	if allowNonEmptyGroupIsSet {
		params.AllowNonEmptyGroup = allowNonEmptyGroupVal.(*bool)
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

	resp, err := client.DashboardGroup.DeleteDashboardGroupByID(params)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return diags
	}

	d.SetId("")
	return diags
}

func getDashboardGroupById(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := dashboard_group.NewGetDashboardGroupByIDParams()

	fieldsVal, fieldsIsSet := d.GetOk("fields")
	if fieldsIsSet {
		params.Fields = fieldsVal.(*string)
	}

	formatVal, formatIsSet := d.GetOk("format")
	if formatIsSet {
		params.Format = formatVal.(*string)
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

	templateVal, templateIsSet := d.GetOk("template")
	if templateIsSet {
		params.Template = templateVal.(*bool)
	}

	client := m.(*client.LogicMonitorRESTAPI)

	resp, err := client.DashboardGroup.GetDashboardGroupByID(params)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return diags
	}

	respModel := resp.GetPayload()
	schemata.SetDashboardGroupResourceData(d, respModel)
	return diags
}

func getDashboardGroupList(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := dashboard_group.NewGetDashboardGroupListParams()

	filterVal, filterIsSet := d.GetOk("filter")
	if filterIsSet {
		stringVal := filterVal.(string)
		params.Filter = &stringVal
	}

	client := m.(*client.LogicMonitorRESTAPI)

	resp, err := client.DashboardGroup.GetDashboardGroupList(params)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return diags
	}

	respModel := resp.GetPayload()
	if len(respModel.Items) == 0 {
		diags = append(diags, diag.Errorf("no DashboardGroup found")...)
	} else {
		result := respModel.Items[0]
		d.SetId(strconv.Itoa(int(result.ID)))
		schemata.SetDashboardGroupResourceData(d, result)
	}
	return diags
}

func updateDashboardGroupById(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	d.Partial(true)

	model := schemata.DashboardGroupModel(d)
	params := dashboard_group.NewUpdateDashboardGroupByIDParams()

	params.SetBody(model)

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
	props := schemata.GetDashboardGroupPropertyFields()

	// loops through array of properties to see which one has changed, the ones that did not change are removed from the list
	for _, v := range props {
		if d.HasChange(v) {
		} else {
			props = utils.Remove(props, v)
		}
	}

	client := m.(*client.LogicMonitorRESTAPI)

	// makes a bulk update for all properties that were changed
	resp, err := client.DashboardGroup.UpdateDashboardGroupByID(params)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return diags
	}

	respModel := resp.GetPayload()
	schemata.SetDashboardGroupResourceData(d, respModel)
	d.Partial(false)

	return diags
}

func resourceDashboardGroupStateImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	client := m.(*client.LogicMonitorRESTAPI)

	// if user provides an ID, we will add the dashboard group directly
	if utils.IsID(d.Id()) {
		id, err := strconv.Atoi(d.Id())
		if err != nil {
			return nil, err
		}

		params := dashboard_group.NewGetDashboardGroupByIDParams()
		params.SetID(int32(id))

		resp, err := client.DashboardGroup.GetDashboardGroupByID(params)
		log.Printf("[TRACE] response: %v", resp)
		if err != nil {
			log.Printf("Failed to find dashboard group %q", err)
			return nil, err
		}

		respModel := resp.GetPayload()
		schemata.SetDashboardGroupResourceData(d, respModel)
	} else {
		// find dashboard_group by name
		params := dashboard_group.NewGetDashboardGroupListParams()
		filter := fmt.Sprintf("name:\"%s\"", d.Id())
		params.SetFilter(&filter)

		resp, err := client.DashboardGroup.GetDashboardGroupList(params)
		log.Printf("[TRACE] response: %v", resp)
		if err != nil {
			err := fmt.Errorf("unexpected: %s", err)
			return nil, err
		}

		if resp.Payload.Total > 1 {
			err := fmt.Errorf("found more than 1 dashboard group with filter %s, please make the filter more specific or import with ID", filter)
			return nil, err
		} else if resp.Payload.Total == 1 {
			schemata.SetDashboardGroupResourceData(d, resp.Payload.Items[0])
			d.SetId(strconv.Itoa(int(resp.Payload.Items[0].ID)))
		} else {
			err := fmt.Errorf("found no dashboard groups with filter '%s'", filter)
			return nil, err
		}
	}
	return []*schema.ResourceData{d}, nil
}
