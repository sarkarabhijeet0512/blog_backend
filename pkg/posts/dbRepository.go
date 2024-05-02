package posts

import (
	model "blog_api/utils/models"
	"context"

	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type Repository interface {
	upsertPost(context.Context, *Post) error
	getPosts(context.Context, model.PostFilter) ([]Post, error)
	deletePosts(context.Context, int) error
}

// NewRepositoryIn is function param struct of func `NewRepository`
type NewRepositoryIn struct {
	fx.In

	Log *logrus.Logger
	DB  *gorm.DB `name:"pointoDB"`
}

// PGRepo is postgres implementation
type dbRepo struct {
	log *logrus.Logger
	db  *gorm.DB
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
	return nil
}

func (r *dbRepo) getPosts(ctx context.Context, filter model.PostFilter) ([]Post, error) {
	return nil, nil
}

func (r *dbRepo) deletePosts(ctx context.Context, ID int) error {
	return nil
}
