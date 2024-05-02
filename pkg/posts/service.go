package posts

import (
	model "blog_api/utils/models"
	"context"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Service struct {
	conf *viper.Viper
	log  *logrus.Logger
	Repo Repository
}

// NewService returns a user service object.
func NewService(conf *viper.Viper, log *logrus.Logger, Repo Repository) *Service {
	return &Service{conf: conf, log: log, Repo: Repo}
}

func (s Service) UpsertPost(ctx context.Context, post *Post) error {
	return s.Repo.upsertPost(ctx, post)
}

func (s Service) GetPosts(ctx context.Context, filter model.PostFilter) ([]Post, error) {
	return s.Repo.getPosts(ctx, filter)
}
func (s Service) DeletePosts(ctx context.Context, ID int) error {
	return s.Repo.deletePosts(ctx, ID)
}
