package user

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
	// User represents the user entity
	User struct {
		tableName        struct{}  `pg:"users,discard_unknown_columns"`
		ID               int       `json:"id" pg:"id,pk"`
		Username         string    `json:"username" pg:"username"`
		Password         string    `json:"password" pg:"password"`
		Email            string    `json:"email" pg:"email,unique"`
		Mobile           string    `json:"mobile" pg:"mobile,unique"`
		FirstName        string    `json:"first_name" pg:"first_name"`
		LastName         string    `json:"last_name" pg:"last_name"`
		RegistrationDate time.Time `json:"registration_date" pg:"registration_date"`
		LastLoginDate    time.Time `json:"last_login_date" pg:"last_login_date"`
		IsActive         bool      `json:"is_active" pg:"is_active"`
		CreatedAt        time.Time `json:"created_at" pg:"created_at"`
		UpdatedAt        time.Time `json:"updated_at" pg:"updated_at"`
	}
	// Address represents the address entity
	UserAddress struct {
		tableName    struct{}  `pg:"user_address,discard_unknown_columns"`
		ID           int       `json:"id" pg:"id"`
		UserID       int       `json:"user_id" pg:"user_id"`
		User         *User     `pg:"fk:user_id"`
		AddressLine1 string    `json:"address_line1" pg:"address_line1"`
		AddressLine2 string    `json:"address_line2" pg:"address_line2"`
		City         string    `json:"city" pg:"city"`
		State        string    `json:"state" pg:"state"`
		Country      string    `json:"country" pg:"country"`
		PostalCode   string    `json:"postal_code" pg:"postal_code"`
		IsActive     bool      `json:"is_active" pg:"is_active"`
		CreatedAt    time.Time `json:"created_at" pg:"created_at"`
		UpdatedAt    time.Time `json:"updated_at" pg:"updated_at"`
	}
	// Payment represents the payment entity
	Payment struct {
		tableName       struct{}       `pg:"payment,discard_unknown_columns"`
		ID              int            `json:"id" pg:"id"`
		UserID          int            `json:"user_id" pg:"user_id"`
		User            *User          `pg:"fk:user_id"`
		PaymentMethodID int            `json:"payment_method_id" pg:"payment_method_id"`
		PaymentMethod   *PaymentMethod `pg:"fk:payment_method_id"`
		CardNumber      string         `json:"card_number" pg:"card_number"`
		ExpiryDate      string         `json:"expiry_date" pg:"expiry_date"`
		VPA             string         `json:"vpa" pg:"vpa"`
		IsActive        bool           `json:"is_active" pg:"is_active"`
		CreatedAt       time.Time      `json:"created_at" pg:"created_at"`
		UpdatedAt       time.Time      `json:"updated_at" pg:"updated_at"`
	}
	PaymentMethod struct {
		tableName struct{}  `pg:"payment_method,discard_unknown_columns"`
		ID        int       `json:"id" pg:"id"`
		Type      string    `json:"type" pg:"type"`
		IsActive  bool      `json:"is_active" pg:"is_active"`
		CreatedAt time.Time `json:"created_at" pg:"created_at"`
		UpdatedAt time.Time `json:"updated_at" pg:"updated_at"`
	}
)
