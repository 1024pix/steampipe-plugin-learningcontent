package learning_content

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	return &plugin.Plugin{
		Name: "steampipe-plugin-learning-content",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		DefaultTransform: transform.FromJSONTag(),
		TableMap: map[string]*plugin.Table{
			"areas":                          areasTable(),
			"challenges":                     challengesTable(),
			"competences":                    competencesTable(),
			"courses":                        coursesTable(),
			"courses_challenges":             coursesChallengesTable(),
			"courses_competences":            coursesCompetencesTable(),
			"frameworks":                     frameworksTable(),
			"skills":                         skillsTable(),
			"skills_tutorials":               skillsTutorialsTable(),
			"skills_learning_more_tutorials": skillsLearningMoreTutorialsTable(),
			"thematics":                      thematicsTable(),
			"tubes":                          tubesTable(),
			"tutorials":                      tutorialsTable(),
		},
	}
}
