package learning_content

import (
	"context"

	client "github.com/1024pix/steampipe-plugin-learning-content/learning-content/client"
	"github.com/turbot/steampipe-plugin-sdk/v2/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
)

func tutorialsTable() *plugin.Table {
	return &plugin.Table{
		Name: "tutorials",
		List: &plugin.ListConfig{
			Hydrate: hydrateTutorialsList,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING},
			{Name: "duration", Type: proto.ColumnType_STRING},
			{Name: "format", Type: proto.ColumnType_STRING},
			{Name: "link", Type: proto.ColumnType_STRING},
			{Name: "source", Type: proto.ColumnType_STRING},
			{Name: "title", Type: proto.ColumnType_STRING},
			{Name: "locale", Type: proto.ColumnType_STRING},
		},
	}
}

func hydrateTutorialsList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	logger.Debug("hydrating tutorials list")

	config := getConfig(d.Connection)

	c := client.New(
		client.WithApiURL(*config.ApiURL),
		client.WithApiKey(*config.ApiKey),
	)

	r, err := c.GetLatestRelease()
	if err != nil {
		return nil, err
	}

	for _, tutorial := range r.Content.Tutorials {
		d.StreamListItem(ctx, tutorial)
	}

	return nil, nil
}
