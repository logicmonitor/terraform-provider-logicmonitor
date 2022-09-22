package logicmonitor

import (
	"log"
	"net/http"
	"terraform-provider-logicmonitor/logicmonitor/internal/apimutex"
	"terraform-provider-logicmonitor/logicmonitor/internal/transport"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/logging"

	"github.com/hashicorp/go-cleanhttp"
)

func (adt *AddRoundTripTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return adt.T.RoundTrip(req)
}

type (
	AddRoundTripTransport struct {
		T http.RoundTripper
	}

	// ValidateClient contains our provider schema values and LogicMonitor clients
	ValidateClient struct {
		apiId      string
		apiKey     string
		company    string
		bulkImport int
		client     *http.Client
	}
)

func (c *ValidateClient) loadAndValidate(bulkResource bool) *http.Client {
	var httpClient *http.Client

	httpClient = cleanhttp.DefaultClient()
	httpClient.Transport = logging.NewTransport("LogicMonitor", httpClient.Transport)

	if bulkResource {
		log.Printf("Going to experimental mode for handling resources in bulk")
		apiMutex, err := apimutex.NewAPIMutex()
		if err != nil {
			log.Printf("[ERROR] Error occured")
		}
		httpClient.Transport = transport.NewGovernedTransport(httpClient.Transport, apiMutex)
		return httpClient
	}
	return nil
}
