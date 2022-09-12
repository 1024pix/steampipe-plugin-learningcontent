package learning_content

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func skillsLearningMoreTutorialsTable() *plugin.Table {
	return &plugin.Table{
		Name: "skills_learning_more_tutorials",
		List: &plugin.ListConfig{
			Hydrate: hydrateSkillsLearningMoreTutorialsList,
		},
		Columns: []*plugin.Column{
			{Name: "skillId", Type: proto.ColumnType_STRING},
			{Name: "tutorialId", Type: proto.ColumnType_STRING},
		},
	}
}

func hydrateSkillsLearningMoreTutorialsList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	logger.Debug("hydrating skills_learning_more_tutorials list")

	r, err := GetLatestRelease(d)
	if err != nil {
		return nil, err
	}

	for _, skill := range r.Content.Skills {
		for _, tutorialId := range skill.LearningMoreTutorialIds {
			d.StreamListItem(ctx, SkillTutorial{skill.Id, tutorialId})
		}
	}

	return nil, nil
}
