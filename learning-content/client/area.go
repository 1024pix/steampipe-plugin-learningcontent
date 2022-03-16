package learning_content_client

type Area struct {
	Id                    string   `json:"id"`
	Code                  string   `json:"code"`
	TitleFrFr             string   `json:"titleFrFr"`
	CompetenceIds         []string `json:"competenceIds"`
	Name                  string   `json:"name"`
	CompetenceAirtableIds []string `json:"competenceAirtableIds"`
	FrameworkId           string   `json:"frameworkId"`
	TitleEnUs             string   `json:"titleEnUs"`
	Color                 string   `json:"color"`
}
