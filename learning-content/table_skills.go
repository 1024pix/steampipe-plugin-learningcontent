package learning_content

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func skillsTable() *plugin.Table {
	return &plugin.Table{
		Name: "skills",
		List: &plugin.ListConfig{
			Hydrate: hydrateSkillsList,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING},
			{Name: "name", Type: proto.ColumnType_STRING},
			{Name: "hintFrFr", Type: proto.ColumnType_STRING},
			{Name: "hintEnUs", Type: proto.ColumnType_STRING},
			{Name: "hintStatus", Type: proto.ColumnType_STRING},
			{Name: "pixValue", Type: proto.ColumnType_INT},
			{Name: "competenceId", Type: proto.ColumnType_STRING},
			{Name: "status", Type: proto.ColumnType_STRING},
			{Name: "tubeId", Type: proto.ColumnType_STRING},
			{Name: "version", Type: proto.ColumnType_STRING},
		},
	}
}

func hydrateSkillsList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	logger.Debug("hydrating skills list")

	r, err := GetLatestRelease(d)
	if err != nil {
		return nil, err
	}

	for _, skill := range r.Content.Skills {
		d.StreamListItem(ctx, skill)
	}

	return nil, nil
}
