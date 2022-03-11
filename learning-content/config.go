package learning_content

import (
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin/schema"
)

type config struct {
	ApiURL *string `cty:"apiURL"`
	ApiKey *string `cty:"apiKey"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"apiURL": {
		Type: schema.TypeString,
	},
	"apiKey": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return new(config)
}

func getConfig(connection *plugin.Connection) config {
	if connection == nil || connection.Config == nil {
		return config{}
	}
	return connection.Config.(config)
}
