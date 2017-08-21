package logicmonitor

import (
	lmv1 "github.com/logicmonitor/lm-sdk-go"
)

//Config struct for setting up AccessKeys and Company name
type Config struct {
	AccessID  string
	AccessKey string
	Company   string
}

//NewLMClient initialization
func (c *Config) NewLMClient() (*lmv1.DefaultApi, error) {
	config := lmv1.NewConfiguration()
	config.APIKey = map[string]map[string]string{
		"Authorization": {
			"AccessID":  c.AccessID,
			"AccessKey": c.AccessKey,
		},
	}
	config.BasePath = "https://" + c.Company + ".logicmonitor.com/santaba/rest"

	api := lmv1.NewDefaultApi()
	api.Configuration = config
	return api, nil
}
