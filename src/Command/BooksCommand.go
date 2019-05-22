package Command

type BookCommand struct {
	Name string `form:"name" json:"name" binding:"required"`
	Description string `form:"description" json:"description" binding:"required"`
	Author int `form:"author" json:"author" binding:"required"`
	Category int `form:"category" json:"category" binding:"required"`
}

