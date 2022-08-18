// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"terraform-provider-logicmonitor/client/collector"
	"terraform-provider-logicmonitor/client/collector_group"
	"terraform-provider-logicmonitor/client/dashboard"
	"terraform-provider-logicmonitor/client/dashboard_group"
	"terraform-provider-logicmonitor/client/data_resource_aws_external_id"
	"terraform-provider-logicmonitor/client/device"
	"terraform-provider-logicmonitor/client/device_group"
)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/santaba/rest"
)

// DefaultSchees are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// Config information for LogicMonitorRESTAPI client
type Config struct {
	AccessKey    *string
	AccessID     *string
	BulkResource *bool
	TransportCfg *TransportConfig
}

// NewConfig create a new empty client Config
func NewConfig() *Config {
	return &Config{
		TransportCfg: DefaultTransportConfig(),
	}
}

// SetAccessID for the client Config
func (c *Config) SetAccessID(accessID *string) {
	c.AccessID = accessID
}

// SetAccessKey for the client Config
func (c *Config) SetAccessKey(accessKey *string) {
	c.AccessKey = accessKey
}

// SetAccountDomain for the client Config
func (c *Config) SetAccountDomain(accountDomain *string) {
	if c.TransportCfg == nil {
		c.TransportCfg = DefaultTransportConfig()
	}
	domain := *accountDomain
	c.TransportCfg = c.TransportCfg.WithHost(domain)
}

// Set Bulk import feature
func (c *Config) SetBulkResource(bulkResource *bool) {
	c.BulkResource = bulkResource
}

// New creates a new logic monitor r e s t API client
func New(c *Config, httpClient *http.Client) *LogicMonitorRESTAPI {
	var transport *httptransport.Runtime
	if httpClient != nil {
		transport = httptransport.NewWithClient(c.TransportCfg.Host, c.TransportCfg.BasePath, c.TransportCfg.Schemes, httpClient)
	} else {
		transport = httptransport.New(c.TransportCfg.Host, c.TransportCfg.BasePath, c.TransportCfg.Schemes)
	}
	authInfo := LMv1Auth(*c.AccessID, *c.AccessKey)

	cli := new(LogicMonitorRESTAPI)
	cli.Transport = transport

	cli.Collector = collector.New(transport, strfmt.Default, authInfo)

	cli.CollectorGroup = collector_group.New(transport, strfmt.Default, authInfo)

	cli.Dashboard = dashboard.New(transport, strfmt.Default, authInfo)

	cli.DashboardGroup = dashboard_group.New(transport, strfmt.Default, authInfo)

	cli.DataResourceAwsExternalID = data_resource_aws_external_id.New(transport, strfmt.Default, authInfo)

	cli.Device = device.New(transport, strfmt.Default, authInfo)

	cli.DeviceGroup = device_group.New(transport, strfmt.Default, authInfo)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// LogicMonitorRESTAPI is a client for logic monitor r e s t API
type LogicMonitorRESTAPI struct {
	Collector *collector.Client

	CollectorGroup *collector_group.Client

	Dashboard *dashboard.Client

	DashboardGroup *dashboard_group.Client

	DataResourceAwsExternalID *data_resource_aws_external_id.Client

	Device *device.Client

	DeviceGroup *device_group.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *LogicMonitorRESTAPI) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Collector.SetTransport(transport)

	c.CollectorGroup.SetTransport(transport)

	c.Dashboard.SetTransport(transport)

	c.DashboardGroup.SetTransport(transport)

	c.DataResourceAwsExternalID.SetTransport(transport)

	c.Device.SetTransport(transport)

	c.DeviceGroup.SetTransport(transport)

}

//TODO: See if there is a way to move this out of Facade Template and into Main or Provider templates
func LMv1Auth(accessId, accessKey string) runtime.ClientAuthInfoWriter {
	return runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		// get epoch
		now := time.Now()
		nanos := now.UnixNano()
		epoch := strconv.FormatInt(nanos/1000000, 10)

		// build the signature
		h := hmac.New(sha256.New, []byte(accessKey))
		h.Write([]byte(r.GetMethod() + epoch))

		if r.GetBodyParam() != nil {
			buf := new(bytes.Buffer)
			enc := json.NewEncoder(buf)
			enc.SetEscapeHTML(false)
			_ = enc.Encode(r.GetBodyParam())
			h.Write(buf.Bytes())
		}

		if r.GetFileParam() != nil {
			for _, files := range r.GetFileParam() {
				for i, file := range files {
					buf := bytes.NewBuffer(nil)
					buf.ReadFrom(file)
					h.Write(buf.Bytes())
					file = runtime.NamedReader(file.Name(), bytes.NewReader(buf.Bytes()))
					files[i] = file
				}
			}
		}

		h.Write([]byte(r.GetPath()))
		hexDigest := hex.EncodeToString(h.Sum(nil))
		signature := base64.StdEncoding.EncodeToString([]byte(hexDigest))
		r.SetHeaderParam("Authorization", fmt.Sprintf("LMv1 %s:%s:%s", accessId, signature, epoch))
		//TODO Consider moving this up to terraform template level of config
		return r.SetHeaderParam("X-version", "3")
	})
}
