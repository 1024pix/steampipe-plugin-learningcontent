package learning_content_client

import (
	"time"
)

type Release struct {
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Content   *Content  `json:"content"`
}
