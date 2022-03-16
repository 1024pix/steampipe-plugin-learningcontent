package learning_content_client

type Content struct {
	Areas       []*Area       `json:"areas"`
	Challenges  []*Challenge  `json:"challenges"`
	Competences []*Competence `json:"competences"`
	Courses     []*Course     `json:"courses"`
	Frameworks  []*Framework  `json:"frameworks"`
	Skills      []*Skill      `json:"skills"`
	Thematics   []*Thematic   `json:"thematics"`
	Tutorials   []*Tutorial   `json:"tutorials"`
	Tubes       []*Tube       `json:"tubes"`
}
