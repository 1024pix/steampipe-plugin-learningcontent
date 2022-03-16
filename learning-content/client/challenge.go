package learning_content_client

type Challenge struct {
	Id                     string   `json:"id"`
	Instruction            string   `json:"instruction"`
	Proposals              string   `json:"proposals"`
	Type                   string   `json:"type"`
	Solution               string   `json:"solution"`
	T1Status               bool     `json:"t1Status"`
	T2Status               bool     `json:"t2Status"`
	T3Status               bool     `json:"t3Status"`
	Status                 string   `json:"status"`
	SkillId                string   `json:"skillId"`
	CompetenceId           string   `json:"competenceId"`
	Format                 string   `json:"format"`
	AutoReply              bool     `json:"autoReply"`
	Locales                []string `json:"locales"`
	AlternativeInstruction string   `json:"alternativeInstruction"`
	Genealogy              string   `json:"genealogy"`
	Responsive             string   `json:"responsive"`
	EmbedUrl               string   `json:"embedUrl"`
	EmbedTitle             string   `json:"embedTitle"`
	EmbedHeight            int      `json:"embedHeight"`
	IllustrationUrl        string   `json:"illustrationUrl"`
	Attachments            []string `json:"attachments"`
	Focusable              bool     `json:"focusable"`
	SolutionToDisplay      string   `json:"solutionToDisplay"`
	IllustrationAlt        string   `json:"illustrationAlt"`
	Timer                  int      `json:"timer"`
	Delta                  float64  `json:"delta"`
	Alpha                  float64  `json:"alpha"`
}
