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
	"terraform-provider-logicmonitor/client/datasource"
	"terraform-provider-logicmonitor/logicmonitor/schemata"
	"terraform-provider-logicmonitor/logicmonitor/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
Datasource datasource API
*/

func Datasource() *schema.Resource {
	return &schema.Resource{
		CreateContext: addDatasourceById,
		DeleteContext: deleteDatasourceById,
		ReadContext:   getDatasourceById,
		UpdateContext: updateDatasourceById,
		Importer: &schema.ResourceImporter{
			State: resourceDatasourceStateImporter,
		},
		Schema: schemata.DatasourceSchema(),
	}
}

func DataResourceDatasource() *schema.Resource {
	return &schema.Resource{
		ReadContext: getDatasourceList,
		Schema:      schemata.DataSourceDatasourceSchema(),
	}
}

func addDatasourceById(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := schemata.DatasourceModel(d)
	params := datasource.NewAddDatasourceByIDParams()

	params.SetBody(model)

	createGraphVal, createGraphIsSet := d.GetOk("create_graph")
	if createGraphIsSet {
		createGraph := createGraphVal.(bool)
		params.CreateGraph = &createGraph
	}

	client := m.(*client.LogicMonitorRESTAPI)

	resp, err := client.Datasource.AddDatasourceByID(params)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return diags
	}

	respModel := resp.GetPayload()
	schemata.SetDatasourceResourceData(d, respModel)
	d.SetId(strconv.Itoa(int(resp.Payload.ID)))

	return diags
}

func deleteDatasourceById(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := datasource.NewDeleteDatasourceByIDParams()

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

	resp, err := client.Datasource.DeleteDatasourceByID(params)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return diags
	}

	d.SetId("")
	return diags
}

func getDatasourceById(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := datasource.NewGetDatasourceByIDParams()

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

	client := m.(*client.LogicMonitorRESTAPI)

	resp, err := client.Datasource.GetDatasourceByID(params)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return diags
	}

	respModel := resp.GetPayload()
	schemata.SetDatasourceResourceData(d, respModel)
	return diags
}

func getDatasourceList(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := datasource.NewGetDatasourceListParams()

	filterVal, filterIsSet := d.GetOk("filter")
	if filterIsSet {
		stringVal := filterVal.(string)
		params.Filter = &stringVal
	}

	client := m.(*client.LogicMonitorRESTAPI)

	resp, err := client.Datasource.GetDatasourceList(params)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return diags
	}

	respModel := resp.GetPayload()
	if len(respModel.Items) == 0 {
		diags = append(diags, diag.Errorf("no Datasource found")...)
	} else {
		result := respModel.Items[0]
		d.SetId(strconv.Itoa(int(result.ID)))
		schemata.SetDatasourceResourceData(d, result)
	}
	return diags
}

func updateDatasourceById(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	d.Partial(true)

	model := schemata.DatasourceModel(d)
	params := datasource.NewUpdateDatasourceByIDParams()

	params.SetBody(model)

	forceRestrictedChangeKeyVal, forceRestrictedChangeKeyIsSet := d.GetOk("force_restricted_change_key")
	if forceRestrictedChangeKeyIsSet {
		params.ForceRestrictedChangeKey = forceRestrictedChangeKeyVal.(*string)
	}

	forceUniqueIdentifierVal, forceUniqueIdentifierIsSet := d.GetOk("force_unique_identifier")
	if forceUniqueIdentifierIsSet {
		forceUniqueIdentifier := forceUniqueIdentifierVal.(bool)
		params.ForceUniqueIdentifier = &forceUniqueIdentifier
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

	reasonVal, reasonIsSet := d.GetOk("reason")
	if reasonIsSet {
		params.Reason = reasonVal.(*string)
	}

	// list of available properties
	props := schemata.GetDatasourcePropertyFields()

	// loops through array of properties to see which one has changed, the ones that did not change are removed from the list
	for _, v := range props {
		if d.HasChange(v) {
		} else {
			props = utils.Remove(props, v)
		}
	}

	client := m.(*client.LogicMonitorRESTAPI)

	// makes a bulk update for all properties that were changed
	resp, err := client.Datasource.UpdateDatasourceByID(params)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return diags
	}

	respModel := resp.GetPayload()
	schemata.SetDatasourceResourceData(d, respModel)
	d.Partial(false)

	return diags
}

func resourceDatasourceStateImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	client := m.(*client.LogicMonitorRESTAPI)

	// if user provides an ID, we will add the datasource directly
	if utils.IsID(d.Id()) {
		id, err := strconv.Atoi(d.Id())
		if err != nil {
			return nil, err
		}

		params := datasource.NewGetDatasourceByIDParams()
		params.SetID(int32(id))

		resp, err := client.Datasource.GetDatasourceByID(params)
		log.Printf("[TRACE] response: %v", resp)
		if err != nil {
			log.Printf("Failed to find datasource %q", err)
			return nil, err
		}

		respModel := resp.GetPayload()
		schemata.SetDatasourceResourceData(d, respModel)
	} else {
		// find datasource by name
		params := datasource.NewGetDatasourceListParams()
		filter := fmt.Sprintf("name:\"%s\"", d.Id())
		params.SetFilter(&filter)

		resp, err := client.Datasource.GetDatasourceList(params)
		log.Printf("[TRACE] response: %v", resp)
		if err != nil {
			err := fmt.Errorf("unexpected: %s", err)
			return nil, err
		}

		if resp.Payload.Total > 1 {
			err := fmt.Errorf("found more than 1 datasource with filter %s, please make the filter more specific or import with ID", filter)
			return nil, err
		} else if resp.Payload.Total == 1 {
			schemata.SetDatasourceResourceData(d, resp.Payload.Items[0])
			d.SetId(strconv.Itoa(int(resp.Payload.Items[0].ID)))
		} else {
			err := fmt.Errorf("found no datasources with filter '%s'", filter)
			return nil, err
		}
	}
	return []*schema.ResourceData{d}, nil
}
