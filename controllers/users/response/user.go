package response

import (
	"foodcal/business/users"
	"time"

	"gorm.io/gorm"
)

type UserResponseRegister struct {
	Id        uint           `json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
	Email     string         `json:"email"`
	Username  string         `json:"username"`
	Password  string         `json:"password"`
	Phone     string         `json:"phone"`
	Address   string         `json:"address"`
	Dob       string         `json:"dob"`
}

type UserResponLogin struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

func FromDomainRegister(domain users.Domain) UserResponseRegister {
	return UserResponseRegister{
		Id:        domain.Id,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		Username:  domain.Username,
		Address:   domain.Address,
		Phone:     domain.Phone,
		Email:     domain.Email,
		Password:  domain.Password,
		Dob:       domain.Dob,
	}
}

func FromDomainLogin(domain users.Domain) UserResponLogin {
	return UserResponLogin{
		Message: "Login success",
		Token:   domain.Token,
	}
}
