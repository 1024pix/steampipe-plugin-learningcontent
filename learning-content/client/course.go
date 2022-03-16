package learning_content_client

type Course struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Challenges  []string `json:"challenges"`
	Description string   `json:"description"`
	ImageUrl    string   `json:"imageUrl"`
	Competences []string `json:"competences"`
}
