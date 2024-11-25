package vote

type InputVote struct {
	PostID uint `json:"post_id" binding:"required"`
	Value  int  `json:"value"`
}
