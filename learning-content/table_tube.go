package learning_content

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v2/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
)

func tubeTable() *plugin.Table {
	return &plugin.Table{
		Name: "tube",
		List: &plugin.ListConfig{
			Hydrate: hydrateTubeList,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING},
			{Name: "name", Type: proto.ColumnType_STRING},
			{Name: "title", Type: proto.ColumnType_STRING},
			{Name: "description", Type: proto.ColumnType_STRING},
			{Name: "competenceId", Type: proto.ColumnType_STRING},
			{Name: "practicalTitleFrFr", Type: proto.ColumnType_STRING},
			{Name: "practicalDescriptionFrFr", Type: proto.ColumnType_STRING},
			{Name: "practicalTitleEnUs", Type: proto.ColumnType_STRING},
			{Name: "practicalDescriptionEnUs", Type: proto.ColumnType_STRING},
		},
	}
}

func hydrateTubeList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	logger.Debug("hydrating tube list")

	r, err := GetLatestRelease(d)
	if err != nil {
		return nil, err
	}

	for _, tube := range r.Content.Tubes {
		d.StreamListItem(ctx, tube)
	}

	return nil, nil
}
