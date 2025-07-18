package logicmonitor

import (
	"context"
	"fmt"
	"net/http"
	"terraform-provider-logicmonitor/client"
	"terraform-provider-logicmonitor/logicmonitor/resources"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var ProviderVersion string

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_id": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("LM_API_ID", nil),
			},
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("LM_API_KEY", nil),
			},
			"domain": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "logicmonitor.com",
			},
			"company": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("LM_COMPANY", nil),
			},
			"bulk_resource": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "true if going for bulk resource, default is false",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"logicmonitor_alert_rule":       resources.AlertRule(),
			"logicmonitor_collector":        resources.Collector(),
			"logicmonitor_collector_group":  resources.CollectorGroup(),
			"logicmonitor_dashboard":        resources.Dashboard(),
			"logicmonitor_dashboard_group":  resources.DashboardGroup(),
			"logicmonitor_datasource":       resources.Datasource(),
			"logicmonitor_device":           resources.Device(),
			"logicmonitor_device_group":     resources.DeviceGroup(),
			"logicmonitor_escalation_chain": resources.EscalationChain(),
			"logicmonitor_website":          resources.Website(),
			"logicmonitor_website_group":    resources.WebsiteGroup(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"logicmonitor_alert_rule":                    resources.DataResourceAlertRule(),
			"logicmonitor_collector":                     resources.DataResourceCollector(),
			"logicmonitor_collector_group":               resources.DataResourceCollectorGroup(),
			"logicmonitor_dashboard":                     resources.DataResourceDashboard(),
			"logicmonitor_dashboard_group":               resources.DataResourceDashboardGroup(),
			"logicmonitor_data_resource_aws_external_id": resources.DataResourceAwsExternalID(),

			"logicmonitor_datasource":       resources.DataResourceDatasource(),
			"logicmonitor_device":           resources.DataResourceDevice(),
			"logicmonitor_device_group":     resources.DataResourceDeviceGroup(),
			"logicmonitor_escalation_chain": resources.DataResourceEscalationChain(),
			"logicmonitor_website":          resources.DataResourceWebsite(),
			"logicmonitor_website_group":    resources.DataResourceWebsiteGroup(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	id := d.Get("api_id").(string)
	key := d.Get("api_key").(string)
	domain := d.Get("domain").(string)
	company := d.Get("company").(string) + "." + domain
	bulkResource := d.Get("bulk_resource").(bool)

	config := client.NewConfig()
	config.SetAccessKey(&key)
	config.SetAccessID(&id)
	config.SetAccountDomain(&company)
	config.SetBulkResource(&bulkResource)

	// Create the HTTP client with a custom User-Agent
	httpClient := &http.Client{
		Transport: &userAgentTransport{
			underlyingTransport: http.DefaultTransport,
			userAgent:           fmt.Sprintf("logicmonitor-terraform-provider/v%s", ProviderVersion),
		},
	}
	c := ValidateClient{}
	httpClient = c.loadAndValidate(httpClient, bulkResource)

	//TODO: Find out what errors this can throw and capture them.
	client := client.New(config, httpClient)

	return client, diags
}

// userAgentTransport is a custom HTTP transport to add the User-Agent header
type userAgentTransport struct {
	underlyingTransport http.RoundTripper
	userAgent           string
}

func (t *userAgentTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", t.userAgent)
	return t.underlyingTransport.RoundTrip(req)
}
