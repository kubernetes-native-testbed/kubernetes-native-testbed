package rate

import "time"

type Rate struct {
	UUID        string
	Rating      int32
	UserUUID    string
	CommentUUID string
	ProductUUID string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
