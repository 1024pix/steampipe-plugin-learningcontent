package learning_content

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v2/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
)

func thematicsTable() *plugin.Table {
	return &plugin.Table{
		Name: "thematics",
		List: &plugin.ListConfig{
			Hydrate: hydrateThematicsList,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING},
			{Name: "name", Type: proto.ColumnType_STRING},
			{Name: "competenceId", Type: proto.ColumnType_STRING},
			{Name: "index", Type: proto.ColumnType_INT},
		},
	}
}

func hydrateThematicsList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	logger.Debug("hydrating thematics list")

	r, err := GetLatestRelease(d)
	if err != nil {
		return nil, err
	}

	for _, area := range r.Content.Thematics {
		d.StreamListItem(ctx, area)
	}

	return nil, nil
}
