package comment

import "time"

type Comment struct {
	UUID          string
	Message       string
	UserUUID      string
	ParentComment *Comment
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
}
