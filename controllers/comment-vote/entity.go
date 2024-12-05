package commentvote

type InputCommentVote struct {
	CommentID uint `json:"comment_id" binding:"required"`
	Value     int `json:"value" binding:"required"`
}
