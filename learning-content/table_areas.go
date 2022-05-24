package learning_content

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v2/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
)

func areasTable() *plugin.Table {
	return &plugin.Table{
		Name: "areas",
		List: &plugin.ListConfig{
			Hydrate: hydrateAreasList,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING},
			{Name: "code", Type: proto.ColumnType_STRING},
			{Name: "titleFrFr", Type: proto.ColumnType_STRING},
			{Name: "name", Type: proto.ColumnType_STRING},
			{Name: "frameworkId", Type: proto.ColumnType_STRING},
			{Name: "titleEnUs", Type: proto.ColumnType_STRING},
			{Name: "color", Type: proto.ColumnType_STRING},
		},
	}
}

func hydrateAreasList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	logger.Debug("hydrating areas list")

	r, err := GetLatestRelease(d)
	if err != nil {
		return nil, err
	}

	for _, area := range r.Content.Areas {
		d.StreamListItem(ctx, area)
	}

	return nil, nil
}
