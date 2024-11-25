package comment

type InputComment struct {
	PostID uint   `json:"post_id" binding:"required"`
	Body   string `json:"body" 	binding:"required"`
}
