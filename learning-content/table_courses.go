package learning_content

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func coursesTable() *plugin.Table {
	return &plugin.Table{
		Name: "courses",
		List: &plugin.ListConfig{
			Hydrate: hydrateCoursesList,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING},
			{Name: "name", Type: proto.ColumnType_STRING},
			{Name: "description", Type: proto.ColumnType_STRING},
			{Name: "imageUrl", Type: proto.ColumnType_STRING},
			// FIXME challenges
			// FIXME competences
		},
	}
}

func hydrateCoursesList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	logger.Debug("hydrating courses list")

	r, err := GetLatestRelease(d)
	if err != nil {
		return nil, err
	}

	for _, area := range r.Content.Courses {
		d.StreamListItem(ctx, area)
	}

	return nil, nil
}
