package logicmonitor 

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"terraform-provider-logicmonitor/client"
	"terraform-provider-logicmonitor/logicmonitor/resources"
	"time"

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
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("LM_API_KEY", nil),
			},
			"domain": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "logicmonitor.com",
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
			"logicmonitor_alert_rule": resources.AlertRule(),
			"logicmonitor_collector": resources.Collector(),
			"logicmonitor_collector_group": resources.CollectorGroup(),
			"logicmonitor_dashboard": resources.Dashboard(),
			"logicmonitor_dashboard_group": resources.DashboardGroup(),
			"logicmonitor_datasource": resources.Datasource(),
			"logicmonitor_device": resources.Device(),
			"logicmonitor_device_group": resources.DeviceGroup(),
			"logicmonitor_device_group_cluster_alert_conf": resources.DeviceGroupClusterAlertConf(),
			"logicmonitor_escalation_chain": resources.EscalationChain(),
			"logicmonitor_report_group": resources.ReportGroup(),
			"logicmonitor_role": resources.Role(),
			"logicmonitor_sdt": resources.Sdt(),
			"logicmonitor_website": resources.Website(),
			"logicmonitor_website_group": resources.WebsiteGroup(),
			"logicmonitor_widget": resources.Widget(),
		},
		DataSourcesMap: map[string]*schema.Resource{
						"logicmonitor_alert_rule": resources.DataResourceAlertRule(),
						"logicmonitor_collector": resources.DataResourceCollector(),
						"logicmonitor_collector_group": resources.DataResourceCollectorGroup(),
						"logicmonitor_dashboard": resources.DataResourceDashboard(),
						"logicmonitor_dashboard_group": resources.DataResourceDashboardGroup(),
					"logicmonitor_data_resource_aws_external_id": resources.DataResourceAwsExternalID(),
				
						"logicmonitor_datasource": resources.DataResourceDatasource(),
						"logicmonitor_device": resources.DataResourceDevice(),
						"logicmonitor_device_group": resources.DataResourceDeviceGroup(),
						"logicmonitor_device_group_cluster_alert_conf": resources.DataResourceDeviceGroupClusterAlertConf(),
						"logicmonitor_escalation_chain": resources.DataResourceEscalationChain(),
						"logicmonitor_report_group": resources.DataResourceReportGroup(),
						"logicmonitor_role": resources.DataResourceRole(),
						"logicmonitor_sdt": resources.DataResourceSdt(),
						"logicmonitor_website": resources.DataResourceWebsite(),
						"logicmonitor_website_group": resources.DataResourceWebsiteGroup(),
						"logicmonitor_widget": resources.DataResourceWidget(),
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

	// Create the HTTP client with a custom User-Agent and 429 retry handling.
    httpClient := &http.Client{
        Transport: &retryTransport{
            underlyingTransport: &userAgentTransport{
                underlyingTransport: http.DefaultTransport,
                userAgent:           fmt.Sprintf("logicmonitor-terraform-provider/v%s", ProviderVersion),
            },
            maxRetries: 6,
            baseDelay:  time.Second,
            maxDelay:   60 * time.Second,
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

// retryTransport retries requests that fail with HTTP 429 (Too Many Requests).
// It honors the server's Retry-After header when present and otherwise falls
// back to exponential backoff with jitter. This allows Terraform workflows that
// issue many GET requests (e.g. per-widget dashboard reads) to complete despite
// hitting the API rate limit instead of failing outright.
type retryTransport struct {
	underlyingTransport http.RoundTripper
	maxRetries          int
	baseDelay           time.Duration
	maxDelay            time.Duration
}

func (t *retryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Buffer the body so the request can be replayed on retry when the caller
	// did not provide a rewindable GetBody (GET requests have no body).
	if req.Body != nil && req.GetBody == nil {
		bodyBytes, err := io.ReadAll(req.Body)
		req.Body.Close()
		if err != nil {
			return nil, err
		}
		req.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		req.GetBody = func() (io.ReadCloser, error) {
			return io.NopCloser(bytes.NewReader(bodyBytes)), nil
		}
	}

	for attempt := 0; ; attempt++ {
		if attempt > 0 && req.GetBody != nil {
			body, err := req.GetBody()
			if err != nil {
				return nil, err
			}
			req.Body = body
		}

		resp, err := t.underlyingTransport.RoundTrip(req)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode != http.StatusTooManyRequests || attempt >= t.maxRetries {
			return resp, nil
		}

		delay := t.backoff(attempt, resp.Header.Get("Retry-After"))
		log.Printf("[WARN] %s %s returned HTTP 429; retrying in %s (attempt %d/%d)",
			req.Method, req.URL.Path, delay, attempt+1, t.maxRetries)

		// Drain and close the body so the underlying connection can be reused.
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()

		select {
		case <-time.After(delay):
		case <-req.Context().Done():
			return nil, req.Context().Err()
		}
	}
}

// backoff computes the wait before the next retry, preferring the Retry-After
// header (delta seconds or HTTP date) and falling back to capped exponential
// backoff with jitter.
func (t *retryTransport) backoff(attempt int, retryAfter string) time.Duration {
	if retryAfter = strings.TrimSpace(retryAfter); retryAfter != "" {
		if secs, err := strconv.Atoi(retryAfter); err == nil && secs >= 0 {
			return t.capDelay(time.Duration(secs) * time.Second)
		}
		if when, err := http.ParseTime(retryAfter); err == nil {
			if d := time.Until(when); d > 0 {
				return t.capDelay(d)
			}
		}
	}

	d := t.baseDelay * time.Duration(1<<uint(attempt))
	d = t.capDelay(d)
	if d > 0 {
		d += time.Duration(rand.Int63n(int64(d)/2 + 1))
	}
	return t.capDelay(d)
}

func (t *retryTransport) capDelay(d time.Duration) time.Duration {
	if t.maxDelay > 0 && d > t.maxDelay {
		return t.maxDelay
	}
	return d
}