package model

type (
	GenericRes struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Data    interface{} `json:"data,omitempty"`
		Meta    interface{} `json:"meta,omitempty"`
	}
	PostFilter struct {
		ID        int    `json:"id" pg:"id"`
		Title     string `json:"title" pg:"title"`
		Content   string `json:"content" pg:"content"`
		Author    string `json:"author" pg:"author"`
		IsActive  *bool  `json:"is_active" pg:"is_active"`
		CreatedAt string `json:"created_at" pg:"created_at"`
	}
	UserRoles struct {
		Resource []int `json:"resource_id" pg:"resource_id"`
	}
)
