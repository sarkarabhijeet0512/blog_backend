package posts

import (
	model "blog_api/utils/models"
	"context"
	"math"

	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

type Repository interface {
	upsertPost(context.Context, *Post) error
	updatePost(ctx context.Context, post *Post) error
	getPosts(context.Context, model.PostFilter) ([]Post, model.Pagination, error)
	deletePosts(context.Context, int) error
}

// NewRepositoryIn is function param struct of func `NewRepository`
type NewRepositoryIn struct {
	fx.In

	Log *logrus.Logger
	DB  *pg.DB `name:"blogdb"`
}

// PGRepo is postgres implementation
type dbRepo struct {
	log *logrus.Logger
	db  *pg.DB
}

// NewDBRepository returns a new persistence layer object which can be used for
// CRUD on db
func NewDBRepository(i NewRepositoryIn) (Repo Repository, err error) {
	Repo = &dbRepo{
		log: i.Log,
		db:  i.DB,
	}

	return
}

func (r *dbRepo) upsertPost(ctx context.Context, post *Post) error {
	_, err := r.db.ModelContext(ctx, post).OnConflict("(title,author) DO UPDATE").Insert()
	return err
}

func (r *dbRepo) updatePost(ctx context.Context, post *Post) error {
	query := r.db.ModelContext(ctx, post)
	if post.Title != "" {
		query.Set("title=?", post.Title)
	}
	if post.IsPublished != nil {
		query.Set("is_published=?", post.IsPublished)
	}
	if post.IsActive != nil {
		query.Set("is_active=?", post.IsActive)
	}
	if post.Likes != 0 {
		query.Set("likes=?", post.Likes)
	}
	if post.CommentsCount != 0 {
		query.Set("comment_count=?", post.CommentsCount)
	}
	if post.Author != "" {
		query.Set("author=?", post.Author)
	}
	if post.Category != "" {
		query.Set("category=?", post.Category)
	}
	if post.Content != "" {
		query.Set("content=?", post.Content)
	}
	if post.Tags != nil {
		query.Set("tags=?", post.Tags)
	}
	_, err := query.Where("id", post.ID).Update()
	return err
}

func (r *dbRepo) getPosts(ctx context.Context, filter model.PostFilter) ([]Post, model.Pagination, error) {
	var (
		posts []Post
		p     model.Pagination
		count int
		err   error
	)
	query := r.db.ModelContext(ctx, &posts)

	if filter.ID != 0 {
		query.Where("id=?", filter.ID)
	}
	if filter.Author != "" {
		query.Where("author=?", filter.Author)
	}
	if filter.Content != "" {
		query.Where(`content ILIKE '%` + filter.Content + `%'`)
	}
	if filter.CreatedAt != "" {
		query.Where("date(created_at)=?", filter.CreatedAt)
	}
	if filter.IsActive != nil {
		query.Where("is_active=?", filter.IsActive)
	}
	if filter.Title != "" {
		query.Where("title=?", filter.Title)
	}
	if filter.Limit != -1 {
		count, err = query.Limit(filter.Limit).Offset((filter.Page - 1) * filter.Limit).Order("rider.id desc").SelectAndCount(&posts)
	} else {
		err = query.Select(&posts)
	}
	if err != nil {
		if err == pg.ErrNoRows {
			r.log.WithContext(ctx).Info(ctx, err.Error())
			return posts, p, nil
		}
		return nil, p, err
	}
	p.TotalDataCount = count
	p.CurrentPage = filter.Page
	p.TotalPages = int(math.Ceil(float64(count) / float64(filter.Limit)))
	return posts, p, err
}

func (r *dbRepo) deletePosts(ctx context.Context, ID int) error {
	var post Post
	_, err := r.db.ModelContext(ctx, &post).Set("is_active=?", false).Where("id=?", ID).Update()
	return err
}
