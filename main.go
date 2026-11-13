package main

import (
	"github.com/turbot/steampipe-plugin-code/code"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: code.Plugin})
}
