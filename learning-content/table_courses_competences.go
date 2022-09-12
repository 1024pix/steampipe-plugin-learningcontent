package learning_content

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func coursesCompetencesTable() *plugin.Table {
	return &plugin.Table{
		Name: "courses_competences",
		List: &plugin.ListConfig{
			Hydrate: hydrateCoursesCompetencesList,
		},
		Columns: []*plugin.Column{
			{Name: "courseId", Type: proto.ColumnType_STRING},
			{Name: "competenceId", Type: proto.ColumnType_STRING},
		},
	}
}

func hydrateCoursesCompetencesList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	logger.Debug("hydrating courses_competences list")

	r, err := GetLatestRelease(d)
	if err != nil {
		return nil, err
	}

	for _, course := range r.Content.Courses {
		for _, competenceId := range course.Competences {
			d.StreamListItem(ctx, CourseCompetence{course.Id, competenceId})
		}
	}

	return nil, nil
}
