package comment

import (
	"encoding/json"
	"time"
)

type Comment struct {
	UUID              string     `json:"uuid"`
	Message           string     `json:"message"`
	UserUUID          string     `json:"user_uuid"`
	ParentCommentUUID string     `json:"parent_comment_uuid"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	DeletedAt         *time.Time `json:"deleted_at"`
}

func (c *Comment) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}
