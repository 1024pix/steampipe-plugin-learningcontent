package learning_content

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v2/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
)

func frameworksTable() *plugin.Table {
	return &plugin.Table{
		Name: "frameworks",
		List: &plugin.ListConfig{
			Hydrate: hydrateFrameworksList,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING},
			{Name: "name", Type: proto.ColumnType_STRING},
		},
	}
}

func hydrateFrameworksList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	logger.Debug("hydrating frameworks list")

	r, err := GetLatestRelease(d)
	if err != nil {
		return nil, err
	}

	for _, framework := range r.Content.Frameworks {
		d.StreamListItem(ctx, framework)
	}

	return nil, nil
}
