package resources

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"terraform-provider-logicmonitor/client"
	"terraform-provider-logicmonitor/logicmonitor/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
DataSource graphs and overview graphs. In the REST API they are sub-resources of a datasource:
/setting/datasources/{id}/graphs and /setting/datasources/{id}/ographs. A graph datapoint points at a datasource
datapoint by id; here it is written by name and the id is looked up, so a graph follows its datapoints.
Requests go through PolymorphicREST (raw JSON), like the widget resource; there are no swagger models for graphs.
*/

// DatasourceGraph is logicmonitor_datasource_graph (instance graphs).
func DatasourceGraph() *schema.Resource { return datasourceGraphResource("graphs") }

// DatasourceOverviewGraph is logicmonitor_datasource_overview_graph (overview graphs).
func DatasourceOverviewGraph() *schema.Resource { return datasourceGraphResource("ographs") }

func datasourceGraphResource(kind string) *schema.Resource {
	return &schema.Resource{
		CreateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			return writeDatasourceGraph(ctx, d, m, kind, http.MethodPost)
		},
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			return readDatasourceGraph(ctx, d, m, kind)
		},
		UpdateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			return writeDatasourceGraph(ctx, d, m, kind, http.MethodPut)
		},
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			_, err := m.(*client.LogicMonitorRESTAPI).PolymorphicREST(ctx, http.MethodDelete, graphPath(d, kind), nil)
			if err != nil && !utils.IsNotFoundError(err) {
				return diag.Errorf("unexpected: %s", err)
			}
			d.SetId("")
			return nil
		},
		Importer: &schema.ResourceImporter{
			// import id: <datasource id>/<graph id>
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				parts := strings.SplitN(d.Id(), "/", 2)
				if len(parts) != 2 || !utils.IsID(parts[0]) || !utils.IsID(parts[1]) {
					return nil, fmt.Errorf("import id must be <datasource id>/<graph id>, got %q", d.Id())
				}
				d.Set("datasource_id", parts[0])
				d.SetId(parts[1])
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: datasourceGraphSchema(kind == "ographs"),
	}
}

func datasourceGraphSchema(overview bool) map[string]*schema.Schema {
	optComputed := func(t schema.ValueType) *schema.Schema {
		return &schema.Schema{Type: t, Optional: true, Computed: true}
	}
	dataPoint := map[string]*schema.Schema{
		"name": {Type: schema.TypeString, Required: true},
		// the datasource datapoint this graph datapoint shows, when it differs from name
		"data_point_name":  {Type: schema.TypeString, Optional: true},
		"consolidate_func": {Type: schema.TypeInt, Optional: true, Default: 1},
	}
	s := map[string]*schema.Schema{
		"datasource_id":  {Type: schema.TypeString, Required: true, ForceNew: true},
		"name":           {Type: schema.TypeString, Required: true},
		"title":          optComputed(schema.TypeString),
		"vertical_label": optComputed(schema.TypeString),
		"time_scale":     optComputed(schema.TypeString),
		"width":          optComputed(schema.TypeInt),
		"height":         optComputed(schema.TypeInt),
		"display_prio":   optComputed(schema.TypeInt),
		"rigid":          optComputed(schema.TypeBool),
		"base1024":       optComputed(schema.TypeBool),
		// numbers, or "NaN" for automatic
		"min_value":   optComputed(schema.TypeString),
		"max_value":   optComputed(schema.TypeString),
		"data_points": {Type: schema.TypeSet, Optional: true, Elem: &schema.Resource{Schema: dataPoint}},
		"virtual_data_points": {Type: schema.TypeSet, Optional: true, Elem: &schema.Resource{Schema: map[string]*schema.Schema{
			"name": {Type: schema.TypeString, Required: true},
			"rpn":  {Type: schema.TypeString, Required: true},
		}}},
		// LogicMonitor does not keep the line order, hence a set
		"lines": {Type: schema.TypeSet, Optional: true, Elem: &schema.Resource{Schema: map[string]*schema.Schema{
			"data_point_name":       {Type: schema.TypeString, Required: true},
			"legend":                {Type: schema.TypeString, Optional: true},
			"type":                  {Type: schema.TypeInt, Optional: true, Default: 1},
			"color_name":            {Type: schema.TypeString, Optional: true, Default: "Auto"},
			"is_virtual_data_point": {Type: schema.TypeBool, Optional: true},
		}}},
	}
	if overview {
		s["aggregated"] = optComputed(schema.TypeBool)
		dataPoint["aggregate_method"] = &schema.Schema{Type: schema.TypeString, Optional: true}
	}
	return s
}

func graphPath(d *schema.ResourceData, kind string) string {
	p := fmt.Sprintf("/setting/datasources/%s/%s", d.Get("datasource_id").(string), kind)
	if d.Id() != "" {
		p += "/" + d.Id()
	}
	return p
}

// datasourceDataPointIDs maps the datasource's datapoint names to ids.
func datasourceDataPointIDs(ctx context.Context, c *client.LogicMonitorRESTAPI, dsID string) (map[string]float64, error) {
	ds, err := c.PolymorphicREST(ctx, http.MethodGet, "/setting/datasources/"+dsID+"?fields=dataPoints", nil)
	if err != nil {
		return nil, err
	}
	ids := map[string]float64{}
	dps, _ := ds["dataPoints"].([]interface{})
	for _, v := range dps {
		if dp, ok := v.(map[string]interface{}); ok {
			if name, ok := dp["name"].(string); ok {
				ids[name], _ = dp["id"].(float64)
			}
		}
	}
	return ids, nil
}

func writeDatasourceGraph(ctx context.Context, d *schema.ResourceData, m interface{}, kind, method string) diag.Diagnostics {
	c := m.(*client.LogicMonitorRESTAPI)
	overview := kind == "ographs"
	ids, err := datasourceDataPointIDs(ctx, c, d.Get("datasource_id").(string))
	if err != nil {
		return diag.Errorf("unexpected: %s", err)
	}
	body := map[string]interface{}{"name": d.Get("name").(string)}
	for attr, key := range map[string]string{"title": "title", "vertical_label": "verticalLabel", "time_scale": "timeScale"} {
		if v := d.Get(attr).(string); v != "" {
			body[key] = v
		}
	}
	for attr, key := range map[string]string{"width": "width", "height": "height", "display_prio": "displayPrio"} {
		if v := d.Get(attr).(int); v != 0 {
			body[key] = v
		}
	}
	body["rigid"] = d.Get("rigid").(bool)
	body["base1024"] = d.Get("base1024").(bool)
	if overview {
		body["aggregated"] = d.Get("aggregated").(bool)
	}
	for attr, key := range map[string]string{"min_value": "minValue", "max_value": "maxValue"} {
		if v := d.Get(attr).(string); v != "" {
			if f, err := strconv.ParseFloat(v, 64); err == nil && !math.IsNaN(f) {
				body[key] = f
			} else {
				body[key] = v
			}
		}
	}
	dataPoints := []interface{}{}
	for _, v := range d.Get("data_points").(*schema.Set).List() {
		dp := v.(map[string]interface{})
		source := dp["data_point_name"].(string)
		if source == "" {
			source = dp["name"].(string)
		}
		id, ok := ids[source]
		if !ok {
			return diag.Errorf("graph %q: datasource has no datapoint %q", d.Get("name"), source)
		}
		out := map[string]interface{}{"dataPointId": id, "dataSourceDataPointId": id, "consolidateFunc": dp["consolidate_func"]}
		if overview {
			out["dataPointName"] = dp["name"]
			out["aggregateMethod"] = dp["aggregate_method"]
		} else {
			out["name"] = dp["name"]
		}
		dataPoints = append(dataPoints, out)
	}
	body["dataPoints"] = dataPoints
	vdps := []interface{}{}
	for _, v := range d.Get("virtual_data_points").(*schema.Set).List() {
		vdp := v.(map[string]interface{})
		vdps = append(vdps, map[string]interface{}{"name": vdp["name"], "rpn": vdp["rpn"]})
	}
	body["virtualDataPoints"] = vdps
	lines := []interface{}{}
	for _, v := range d.Get("lines").(*schema.Set).List() {
		l := v.(map[string]interface{})
		lines = append(lines, map[string]interface{}{"dataPointName": l["data_point_name"], "legend": l["legend"], "type": l["type"],
			"colorName": l["color_name"], "isVirtualDataPoint": l["is_virtual_data_point"]})
	}
	body["lines"] = lines

	if method == http.MethodPost {
		d.SetId("")
	}
	resp, err := c.PolymorphicREST(ctx, method, graphPath(d, kind), body)
	if err != nil {
		return diag.Errorf("unexpected: %s", err)
	}
	if method == http.MethodPost {
		id, _ := resp["id"].(float64)
		d.SetId(strconv.Itoa(int(id)))
	}
	return readDatasourceGraph(ctx, d, m, kind)
}

func readDatasourceGraph(ctx context.Context, d *schema.ResourceData, m interface{}, kind string) diag.Diagnostics {
	c := m.(*client.LogicMonitorRESTAPI)
	overview := kind == "ographs"
	g, err := c.PolymorphicREST(ctx, http.MethodGet, graphPath(d, kind), nil)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("unexpected: %s", err)
	}
	ids, err := datasourceDataPointIDs(ctx, c, d.Get("datasource_id").(string))
	if err != nil {
		return diag.Errorf("unexpected: %s", err)
	}
	names := map[float64]string{}
	for n, id := range ids {
		names[id] = n
	}
	str := func(k string) string { s, _ := g[k].(string); return s }
	num := func(k string) int { f, _ := g[k].(float64); return int(f) }
	boolean := func(k string) bool { b, _ := g[k].(bool); return b }
	value := func(k string) string {
		if f, ok := g[k].(float64); ok {
			return strconv.FormatFloat(f, 'f', -1, 64)
		}
		return str(k)
	}
	d.Set("name", str("name"))
	d.Set("title", str("title"))
	d.Set("vertical_label", str("verticalLabel"))
	d.Set("time_scale", str("timeScale"))
	d.Set("width", num("width"))
	d.Set("height", num("height"))
	d.Set("display_prio", num("displayPrio"))
	d.Set("rigid", boolean("rigid"))
	d.Set("base1024", boolean("base1024"))
	d.Set("min_value", value("minValue"))
	d.Set("max_value", value("maxValue"))
	if overview {
		d.Set("aggregated", boolean("aggregated"))
	}
	var dataPoints, vdps, lines []interface{}
	items := func(k string) []interface{} { l, _ := g[k].([]interface{}); return l }
	toInt := func(v interface{}) int { f, _ := v.(float64); return int(f) }
	for _, v := range items("dataPoints") {
		dp, _ := v.(map[string]interface{})
		id, _ := dp["dataPointId"].(float64)
		out := map[string]interface{}{"consolidate_func": toInt(dp["consolidateFunc"])}
		if overview {
			out["name"] = dp["dataPointName"]
			out["aggregate_method"] = dp["aggregateMethod"]
		} else {
			out["name"] = dp["name"]
			if names[id] != dp["name"] {
				out["data_point_name"] = names[id]
			}
		}
		dataPoints = append(dataPoints, out)
	}
	for _, v := range items("virtualDataPoints") {
		vdp, _ := v.(map[string]interface{})
		vdps = append(vdps, map[string]interface{}{"name": vdp["name"], "rpn": vdp["rpn"]})
	}
	for _, v := range items("lines") {
		l, _ := v.(map[string]interface{})
		lines = append(lines, map[string]interface{}{"data_point_name": l["dataPointName"], "legend": l["legend"],
			"type": toInt(l["type"]), "color_name": l["colorName"], "is_virtual_data_point": l["isVirtualDataPoint"]})
	}
	d.Set("data_points", dataPoints)
	d.Set("virtual_data_points", vdps)
	d.Set("lines", lines)
	return nil
}
