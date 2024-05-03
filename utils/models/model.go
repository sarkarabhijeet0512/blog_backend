package model

type (
	GenericRes struct {
		Success bool       `json:"success"`
		Message string     `json:"message"`
		Data    any        `json:"data,omitempty"`
		Meta    Pagination `json:"meta,omitempty"`
	}
	PostFilter struct {
		ID        int    `form:"id" json:"id" pg:"id"`
		Title     string `form:"title" json:"title" pg:"title"`
		Content   string `form:"content" json:"content" pg:"content"`
		Author    string `form:"author" json:"author" pg:"author"`
		Page      int    `form:"page,default=1"`
		Limit     int    `form:"limit,default=20"`
		IsActive  *bool  `form:"is_active" json:"is_active" pg:"is_active"`
		CreatedAt string `form:"created_at" json:"created_at" pg:"created_at"`
	}
	UserRoles struct {
		Resource []int `json:"resource_id" pg:"resource_id"`
	}
	Pagination struct {
		CurrentPage    int `json:"current_page,omitempty"`
		TotalPages     int `json:"total_pages,omitempty"`
		TotalDataCount int `json:"total_data_count,omitempty"`
	}
)
