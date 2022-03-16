package learning_content_client

type Tutorial struct {
	Id       string `json:"id"`
	Duration string `json:"duration"`
	Format   string `json:"format"`
	Link     string `json:"link"`
	Source   string `json:"source"`
	Title    string `json:"title"`
	Locale   string `json:"locale"`
}
