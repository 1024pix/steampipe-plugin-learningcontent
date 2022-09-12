package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"

	learning_content "github.com/1024pix/steampipe-plugin-learning-content/learning-content"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: learning_content.Plugin})
}
