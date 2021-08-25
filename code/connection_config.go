package code

import (
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/schema"
)

type codeConfig struct {
	CustomPatterns []string `cty:"custom_patterns"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"custom_patterns": {
		Type: schema.TypeList,
		Elem: &schema.Attribute{Type: schema.TypeString},
	},
}

func ConfigInstance() interface{} {
	return &codeConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) codeConfig {
	if connection == nil || connection.Config == nil {
		return codeConfig{}
	}
	config, _ := connection.Config.(codeConfig)
	return config
}
