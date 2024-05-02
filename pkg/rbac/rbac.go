package rbac

import (
	"time"

	"go.uber.org/fx"
)

// Module provides all constructor and invocation methods to facilitate credits module
var Module = fx.Options(
	fx.Provide(
		NewDBRepository,
		NewService,
		provideErrorChannel,
	),
)

type (
	// Role represents the role entity
	Role struct {
		tableName      struct{}         `pg:"roles,discard_unknown_columns"`
		ID             int              `json:"id" pg:"id"`
		RoleName       string           `json:"role_name" pg:"role_name,unique"`
		RolePermission []RolePermission `json:"role_permission" pg:"rel:has-many"`
		IsActive       bool             `json:"is_active" pg:"is_active"`
		CreatedAt      time.Time        `json:"created_at" pg:"created_at"`
		UpdatedAt      time.Time        `json:"updated_at" pg:"updated_at"`
	}

	// UserRole represents the many-to-many relationship between users and roles
	UserRole struct {
		tableName struct{}  `pg:"user_roles,discard_unknown_columns"`
		ID        int       `json:"id" pg:"id"`
		UserID    int       `json:"user_id" pg:"user_id,unique"`
		RoleID    []int     `json:"role_id" pg:"role_id"`
		Roles     []*Role   `json:"roles,omitempty" pg:"rel:has-many"`
		IsActive  bool      `json:"is_active" pg:"is_active"`
		CreatedAt time.Time `json:"created_at" pg:"created_at"`
		UpdatedAt time.Time `json:"updated_at" pg:"updated_at"`
	}

	// Resource represents the resource entity
	Resource struct {
		tableName      struct{}        `pg:"resources,discard_unknown_columns"`
		ID             int             `json:"id" pg:"id"`
		ResourceName   string          `json:"resource_name" pg:"resource_name"`
		Role           *Role           `json:"role,omitempty" pg:"rel:belongs-to"`
		IsActive       bool            `json:"is_active" pg:"is_active"`
		CreatedAt      time.Time       `json:"created_at" pg:"created_at"`
		UpdatedAt      time.Time       `json:"updated_at" pg:"updated_at"`
		RolePermission *RolePermission `json:"role_permission,omitempty" pg:"rel:belongs-to"`
	}

	// Permission represents the permission entity
	Permission struct {
		tableName      struct{}        `pg:"permissions,discard_unknown_columns"`
		ID             int             `json:"id" pg:"id"`
		Read           bool            `json:"read" pg:"read,default:false"`
		Write          bool            `json:"write" pg:"write,default:false"`
		Edit           bool            `json:"edit" pg:"edit,default:false"`
		Remove         bool            `json:"remove" pg:"remove,default:false"`
		RolePermission *RolePermission `json:"role_permission,omitempty" pg:"rel:belongs-to"`
		IsActive       bool            `json:"is_active" pg:"is_active,default:true"`
		CreatedAt      time.Time       `json:"created_at" pg:"created_at"`
		UpdatedAt      time.Time       `json:"updated_at" pg:"updated_at"`
	}

	// RolePermission represents the many-to-many relationship between roles and permissions
	RolePermission struct {
		tableName    struct{}    `pg:"role_permissions,discard_unknown_columns"`
		ID           int         `json:"id" pg:"id"`
		RoleID       int         `json:"role_id" pg:"role_id"`
		PermissionID int         `json:"permission_id" pg:"permission_id"`
		ResourceID   int         `json:"resourse_id" pg:"resourse_id"`
		Role         *Role       `json:"role,omitempty" pg:"rel:has-one"`       // Relationship with Role struct
		Permission   *Permission `json:"permission,omitempty" pg:"rel:has-one"` // Relationship with Permission struct
		Resource     []*Resource `json:"resourse" pg:"rel:has-many"`            // Relationship with Resource struct
		IsActive     bool        `json:"is_active" pg:"is_active"`
		CreatedAt    time.Time   `json:"created_at" pg:"created_at"`
		UpdatedAt    time.Time   `json:"updated_at" pg:"updated_at"`
	}
)

func provideErrorChannel() chan error {
	return make(chan error, 1)
}

func (p *Permission) SQL() string {
	return `
		CREATE UNIQUE INDEX idx_permissions_unique_combination
		ON permissions (read, write, edit, remove);
	`
}
func (p *RolePermission) SQL() string {
	return `
		CREATE UNIQUE INDEX idx_role_permissions_unique_combination
		ON role_permissions (role_id,permission_id,resourse_id);
	`
}
