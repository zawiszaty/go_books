package Command

type AuthorCommand struct {
	Name string `form:"title" json:"name" binding:"required"`
}

