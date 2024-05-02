package rbac

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Service struct {
	conf         *viper.Viper
	log          *logrus.Logger
	Repo         Repository
	ErrorChannel chan error
}

// NewService returns a user service object.
func NewService(conf *viper.Viper, log *logrus.Logger, Repo Repository, ErrorChannel chan error) *Service {
	return &Service{
		conf:         conf,
		log:          log,
		Repo:         Repo,
		ErrorChannel: ErrorChannel,
	}
}

// IsDBActive gets user data by her userID
func (s *Service) CreateUserRole(dCtx context.Context, r *Role) error {
	if err := s.Repo.upsertRole(dCtx, r); err != nil {
		return err
	}
	go func(rolePermissions []RolePermission, s *Service, dCtx context.Context) {
		for _, role := range r.RolePermission {
			if err := s.Repo.upsertPermission(dCtx, role.Permission); err != nil {
				s.ErrorChannel <- err // Send error to channel
				return
			}
			role.RoleID = r.ID
			role.PermissionID = role.Permission.ID
			if err := s.Repo.upsertRolePermission(dCtx, &role); err != nil {
				s.ErrorChannel <- err // Send error to channel
				return
			}
		}
	}(r.RolePermission, s, dCtx)
	return nil
}
func (s *Service) AssignRole(dCtx context.Context, r *UserRole) error {
	return s.Repo.upsertAssignRole(dCtx, r)
}
func (s *Service) FetchUserRole(dCtx context.Context, userID any) ([]UserRole, error) {
	return s.Repo.fetchUserRole(dCtx, userID)
}

func (s *Service) FetchRole(dCtx context.Context, roleID int) ([]Role, error) {
	return s.Repo.fetchRole(dCtx, roleID)
}
func (s *Service) FetchAllRoles(dCtx context.Context) ([]Role, error) {
	return s.Repo.fetchALLRoles(dCtx)
}
