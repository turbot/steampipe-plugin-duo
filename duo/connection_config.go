package duo

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
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
	if connection == nil || connection.Config == nil {
		return duoConfig{}
	}
	config, _ := connection.Config.(duoConfig)
	return config
}
