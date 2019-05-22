package Command

type CategoryCommand struct {
	Name string `form:"title" json:"name" binding:"required"`
}


