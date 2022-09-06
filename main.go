package main

import (
	"github.com/turbot/steampipe-plugin-code/code"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: code.Plugin})
}
