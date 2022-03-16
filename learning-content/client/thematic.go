package learning_content_client

type Thematic struct {
	Id           string   `json:"id"`
	Name         string   `json:"name"`
	CompetenceId string   `json:"competenceId"`
	TubeIds      []string `json:"tubeIds"`
	Index        int      `json:"index"`
}
