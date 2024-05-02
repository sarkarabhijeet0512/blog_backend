package posts

import (
	"time"

	"go.uber.org/fx"
)

// Module provides all constructor and invocation methods to facilitate credits module
var Module = fx.Options(
	fx.Provide(
		NewDBRepository,
		NewService,
	),
)

type (
	Post struct {
		ID            int       `json:"id" pg:"id"`
		Title         string    `json:"title" pg:"title"`
		Content       string    `json:"content" pg:"content"`
		Author        string    `json:"author" pg:"author"`
		IsActive      bool      `json:"is_active" pg:"is_active"`
		Tags          []string  `json:"tags" pg:"tags"`
		Category      string    `json:"category" pg:"category"`
		Views         int       `json:"views" pg:"views"`
		Likes         int       `json:"likes" pg:"likes"`
		CommentsCount int       `json:"comments_count" pg:"comments"`
		IsPublished   bool      `json:"is_published" pg:"is_published"`
		CreatedAt     time.Time `json:"created_at" pg:"created_at"`
		UpdatedAt     time.Time `json:"updated_at" pg:"updated_at"`
	}
)
