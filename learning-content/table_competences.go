package learning_content

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func competencesTable() *plugin.Table {
	return &plugin.Table{
		Name: "competences",
		List: &plugin.ListConfig{
			Hydrate: hydrateCompetencesList,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING},
			{Name: "nameFrFr", Type: proto.ColumnType_STRING},
			{Name: "index", Type: proto.ColumnType_STRING},
			{Name: "areaId", Type: proto.ColumnType_STRING},
			{Name: "origin", Type: proto.ColumnType_STRING},
			{Name: "descriptionFrFr", Type: proto.ColumnType_STRING},
			{Name: "nameEnUs", Type: proto.ColumnType_STRING},
			{Name: "descriptionEnUs", Type: proto.ColumnType_STRING},
			{Name: "name", Type: proto.ColumnType_STRING},
			{Name: "description", Type: proto.ColumnType_STRING},
		},
	}
}

func hydrateCompetencesList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	logger.Debug("hydrating competences list")

	r, err := GetLatestRelease(d)
	if err != nil {
		return nil, err
	}

	for _, area := range r.Content.Competences {
		d.StreamListItem(ctx, area)
	}

	return nil, nil
}
