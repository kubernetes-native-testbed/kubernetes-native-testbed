package rate

import "encoding/json"

type Rate struct {
	UUID        string `json:"uuid"`
	Rating      int32  `json:"rating"`
	UserUUID    string `json:"user_uuid"`
	CommentUUID string `json:"comment_uuid"`
	ProductUUID string `json:"product_uuid"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
	DeletedAt   int64  `json:"deleted_at"`
}

func (r *Rate) String() string {
	b, _ := json.Marshal(r)
	return string(b)
}
