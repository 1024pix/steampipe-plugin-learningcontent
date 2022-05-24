package learning_content_client

type Thematic struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	CompetenceId string `json:"competenceId"`
	Index        int    `json:"index"`
	// TubeIds      []string `json:"tubeIds"`
}
