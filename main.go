package main

import (
	"github.com/turbot/steampipe-plugin-code/code"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: code.Plugin})
}
