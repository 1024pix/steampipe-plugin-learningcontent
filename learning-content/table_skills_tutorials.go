package learning_content

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func skillsTutorialsTable() *plugin.Table {
	return &plugin.Table{
		Name: "skills_tutorials",
		List: &plugin.ListConfig{
			Hydrate: hydrateSkillsTutorialsList,
		},
		Columns: []*plugin.Column{
			{Name: "skillId", Type: proto.ColumnType_STRING},
			{Name: "tutorialId", Type: proto.ColumnType_STRING},
		},
	}
}

func hydrateSkillsTutorialsList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	logger.Debug("hydrating skills_tutorials list")

	r, err := GetLatestRelease(d)
	if err != nil {
		return nil, err
	}

	for _, skill := range r.Content.Skills {
		for _, tutorialId := range skill.TutorialIds {
			d.StreamListItem(ctx, SkillTutorial{skill.Id, tutorialId})
		}
	}

	return nil, nil
}
