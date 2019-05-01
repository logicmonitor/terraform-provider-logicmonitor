package logicmonitor

import (
	lmclient "github.com/logicmonitor/lm-sdk-go/client"
)

//Config struct for setting up AccessKeys and Company name
type Config struct {
	AccessID  string
	AccessKey string
	Company   string
}

//NewLMClient initialization
func (c *Config) newLMClient() (*lmclient.LMSdkGo, error) {
	config := lmclient.NewConfig()

	config.SetAccessID(&c.AccessID)
	config.SetAccessKey(&c.AccessKey)
	config.SetAccountDomain(&c.Company)
	client := lmclient.New(config)
	//	config.SetUserAgent("LogicMonitor Terraform Provider")

	return client, nil
}
