package learning_content_client

type Skill struct {
	Id                      string   `json:"id"`
	Name                    string   `json:"name"`
	HintFrFr                string   `json:"hintFrFr"`
	HintEnUs                string   `json:"hintEnUs"`
	HintStatus              string   `json:"hintStatus"`
	TutorialIds             []string `json:"tutorialIds"`
	LearningMoreTutorialIds []string `json:"learningMoreTutorialIds"`
	PixValue                int      `json:"pixValue"`
	CompetenceId            string   `json:"competenceId"`
	Status                  string   `json:"status"`
	TubeId                  string   `json:"tubeId"`
	Version                 int      `json:"version"`
}
