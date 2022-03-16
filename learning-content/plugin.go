package learning_content

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin/transform"
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
			"skills":                         skillsTable(),
			"skills_tutorials":               skillsTutorialsTable(),
			"skills_learning_more_tutorials": skillsLearningMoreTutorialsTable(),
			"tutorials":                      tutorialsTable(),
		},
	}
}
