package post

type InputPost struct {
	Title string   `json:"title" binding:"required"` // required field
	Body  string   `json:"body" binding:"required"`  // required field
	Tags  []string `json:"tags" binding:"required"`  // required field
}

