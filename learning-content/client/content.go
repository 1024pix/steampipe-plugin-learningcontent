package learning_content_client

type Content struct {
	Areas      []*Area      `json:"areas"`
	Challenges []*Challenge `json:"challenges"`
	Skills     []*Skill     `json:"skills"`
	Tutorials  []*Tutorial  `json:"tutorials"`
}
