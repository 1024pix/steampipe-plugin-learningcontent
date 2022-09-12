package learning_content

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func coursesChallengesTable() *plugin.Table {
	return &plugin.Table{
		Name: "courses_challenges",
		List: &plugin.ListConfig{
			Hydrate: hydrateCoursesChallengesList,
		},
		Columns: []*plugin.Column{
			{Name: "courseId", Type: proto.ColumnType_STRING},
			{Name: "challengeId", Type: proto.ColumnType_STRING},
		},
	}
}

func hydrateCoursesChallengesList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	logger.Debug("hydrating courses_challenges list")

	r, err := GetLatestRelease(d)
	if err != nil {
		return nil, err
	}

	for _, course := range r.Content.Courses {
		for _, challengeId := range course.Challenges {
			d.StreamListItem(ctx, CourseChallenge{course.Id, challengeId})
		}
	}

	return nil, nil
}
