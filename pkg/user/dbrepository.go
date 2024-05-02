package user

import (
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

type Repository interface {
	IsActive() (ok bool, err error)
	GetDBConnection(dCtx context.Context) *pg.DB
	upsertUserRegistration(dCtx context.Context, req *User) error
	fetchUserByMobileNumberOrEmail(dCtx context.Context, req User) (res *User, err error)
}

// NewRepositoryIn is function param struct of func `NewRepository`
type NewRepositoryIn struct {
	fx.In

	Log *logrus.Logger
	DB  *pg.DB `name:"userdb"`
}

// PGRepo is postgres implementation
type PGRepo struct {
	log *logrus.Logger
	db  *pg.DB
}

// NewDBRepository returns a new persistence layer object which can be used for
// CRUD on db
func NewDBRepository(i NewRepositoryIn) (Repo Repository, err error) {

	Repo = &PGRepo{
		log: i.Log,
		db:  i.DB,
	}

	return
}

// IsActive checks if DB is connected
func (r *PGRepo) IsActive() (ok bool, err error) {

	ctx := context.Background()
	err = r.db.Ping(ctx)
	if err == nil {
		ok = true
	}
	return
}
func (r *PGRepo) GetDBConnection(dCtx context.Context) *pg.DB {
	return r.db
}

func (r *PGRepo) upsertUserRegistration(dCtx context.Context, req *User) error {
	_, err := r.db.ModelContext(dCtx, req).OnConflict("(mobile,email) DO UPDATE").Insert()
	return err
}

func (r *PGRepo) fetchUserByMobileNumberOrEmail(dCtx context.Context, req User) (res *User, err error) {
	res = &User{
		Email: req.Email,
	}
	query := r.db.ModelContext(dCtx, res)
	switch {
	case req.Mobile != "":
		query.Where("mobile=?", req.Mobile).Select()
	case req.Email != "":
		query.Where("email=?", req.Email).Select()
	}
	return res, nil
}
