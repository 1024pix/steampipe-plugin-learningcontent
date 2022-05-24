package learning_content

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v2/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
)

func challengesTable() *plugin.Table {
	return &plugin.Table{
		Name: "challenges",
		List: &plugin.ListConfig{
			Hydrate: hydrateChallengesList,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING},
			{Name: "instruction", Type: proto.ColumnType_STRING},
			{Name: "proposals", Type: proto.ColumnType_STRING},
			{Name: "type", Type: proto.ColumnType_STRING},
			{Name: "solution", Type: proto.ColumnType_STRING},
			{Name: "t1Status", Type: proto.ColumnType_BOOL},
			{Name: "t2Status", Type: proto.ColumnType_BOOL},
			{Name: "t3Status", Type: proto.ColumnType_BOOL},
			{Name: "status", Type: proto.ColumnType_STRING},
			{Name: "skillId", Type: proto.ColumnType_STRING},
			{Name: "format", Type: proto.ColumnType_STRING},
			{Name: "autoReply", Type: proto.ColumnType_BOOL},
			{Name: "locales", Type: proto.ColumnType_STRING},
			{Name: "alternativeInstruction", Type: proto.ColumnType_STRING},
			{Name: "genealogy", Type: proto.ColumnType_STRING},
			{Name: "responsive", Type: proto.ColumnType_STRING},
			{Name: "embedUrl", Type: proto.ColumnType_STRING},
			{Name: "embedTitle", Type: proto.ColumnType_STRING},
			{Name: "embedHeight", Type: proto.ColumnType_INT},
			{Name: "illustrationUrl", Type: proto.ColumnType_STRING},
			{Name: "focusable", Type: proto.ColumnType_BOOL},
			{Name: "solutionToDisplay", Type: proto.ColumnType_STRING},
			{Name: "illustrationAlt", Type: proto.ColumnType_STRING},
			{Name: "timer", Type: proto.ColumnType_INT},
			{Name: "delta", Type: proto.ColumnType_DOUBLE},
			{Name: "alpha", Type: proto.ColumnType_DOUBLE},
		},
	}
}

func hydrateChallengesList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	logger.Debug("hydrating challenges list")

	r, err := GetLatestRelease(d)
	if err != nil {
		return nil, err
	}

	for _, challenge := range r.Content.Challenges {
		d.StreamListItem(ctx, challenge)
	}

	return nil, nil
}
