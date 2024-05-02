package user

import (
	"blog_api/utils"
	"context"
	"time"

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
	return &Service{
		conf: conf,
		log:  log,
		Repo: Repo,
	}
}

// IsDBActive gets user data by her userID
func (s *Service) IsActive() (bool, error) {
	return s.Repo.IsActive()
}
func (s *Service) UpsertUserRegistration(dCtx context.Context, req *User) error {
	req.Password = utils.HashPassword(req.Password)
	req.IsActive = true
	req.CreatedAt = time.Now()
	req.UpdatedAt = time.Now()
	req.RegistrationDate = time.Now()
	return s.Repo.upsertUserRegistration(dCtx, req)
}

func (s *Service) FetchUserByMobileNumberOrEmail(dCtx context.Context, req User) (res *User, ok bool, err error) {
	res, err = s.Repo.fetchUserByMobileNumberOrEmail(dCtx, req)
	if err != nil {
		return nil, false, err
	}
	ok = utils.CheckPasswordHash(req.Password, res.Password)
	return res, ok, nil
}
