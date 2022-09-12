package learning_content

import (
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"

	client "github.com/1024pix/steampipe-plugin-learning-content/learning-content/client"
)

func GetLatestRelease(d *plugin.QueryData) (*client.Release, error) {
	config := getConfig(d.Connection)

	c := client.New(
		client.WithApiURL(*config.ApiURL),
		client.WithApiKey(*config.ApiKey),
	)

	return c.GetLatestRelease()
}
