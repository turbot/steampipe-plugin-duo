package duo

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

type duoConfig struct {
	APIHostname    *string `hcl:"api_hostname"`
	IntegrationKey *string `hcl:"integration_key"`
	SecretKey      *string `hcl:"secret_key"`
}

func ConfigInstance() interface{} {
	return &duoConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) duoConfig {
	if connection == nil || connection.GetConfig() == nil {
		return duoConfig{}
	}
	config, _ := connection.GetConfig().(duoConfig)
	return config
}
