package learning_content_client

type Area struct {
	Id          string `json:"id"`
	Code        string `json:"code"`
	TitleFrFr   string `json:"titleFrFr"`
	Name        string `json:"name"`
	FrameworkId string `json:"frameworkId"`
	TitleEnUs   string `json:"titleEnUs"`
	Color       string `json:"color"`
	// CompetenceIds         []string `json:"competenceIds"`
	// CompetenceAirtableIds []string `json:"competenceAirtableIds"`
}
