package users

import (
	"foodcal/business/users"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Username  string
	Email     string
	Password  string
	Address   string
	Phone     string
	Dob       string
}

func (user User) ToDomain() users.Domain {
	return users.Domain{
		Id:        user.Id,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		Address:   user.Address,
		Phone:     user.Phone,
		Dob:       user.Dob,
	}
}

func FromDomain(domain users.Domain) User {
	return User{
		Id:        domain.Id,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		Username:  domain.Username,
		Email:     domain.Email,
		Password:  domain.Password,
		Address:   domain.Address,
		Phone:     domain.Phone,
		Dob:       domain.Dob,
	}
}
