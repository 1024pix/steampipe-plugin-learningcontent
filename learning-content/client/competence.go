package learning_content_client

type Competence struct {
	Id              string `json:"id"`
	NameFrFr        string `json:"nameFrFr"`
	Index           string `json:"index"`
	AreaId          string `json:"areaId"`
	Origin          string `json:"origin"`
	DescriptionFrFr string `json:"descriptionFrFr"`
	NameEnUs        string `json:"nameEnUs"`
	DescriptionEnUs string `json:"descriptionEnUs"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	// SkillIds        []string `json:"skillIds"`
	// ThematicIds     []string `json:"thematicIds"`
}
