package users

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	Id        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Username  string
	Address   string
	Phone     string
	Email     string
	Password  string
	Dob       string
	Token     string
}

type UserUsecaseInterface interface {
	Register(domain *Domain) (Domain, error)
	Login(domain Domain, ctx context.Context) (Domain, error)
	// GetAllUsers(ctx context.Context) ([]Domain, error)
}

type UserRepoInterface interface {
	Register(domain *Domain) (Domain, error)
	Login(domain Domain, ctx context.Context) (Domain, error)
	// GetAllUsers(ctx context.Context) ([]Domain, error)
}
