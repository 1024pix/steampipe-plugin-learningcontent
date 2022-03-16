package learning_content_client

type Content struct {
	Skills    []*Skill    `json:"skills"`
	Tutorials []*Tutorial `json:"tutorials"`
}
