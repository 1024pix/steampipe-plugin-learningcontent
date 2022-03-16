package learning_content_client

type Tube struct {
	Id                       string `json:"id"`
	Name                     string `json:"name"`
	Title                    string `json:"title"`
	Description              string `json:"description"`
	CompetenceId             string `json:"competenceId"`
	PracticalTitleFrFr       string `json:"practicalTitleFrFr"`
	PracticalDescriptionFrFr string `json:"practicalDescriptionFrFr"`
	PracticalTitleEnUs       string `json:"practicalTitleEnUs"`
	PracticalDescriptionEnUs string `json:"practicalDescriptionEnUs"`
}
