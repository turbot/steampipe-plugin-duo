package main

import (
	"github.com/turbot/steampipe-plugin-duo/duo"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: duo.Plugin})
}
