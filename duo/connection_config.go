package duo

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type duoConfig struct {
	APIHostname    *string `cty:"api_hostname"`
	IntegrationKey *string `cty:"integration_key"`
	SecretKey      *string `cty:"secret_key"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"api_hostname": {
		Type: schema.TypeString,
	},
	"integration_key": {
		Type: schema.TypeString,
	},
	"secret_key": {
		Type: schema.TypeString,
	},
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
