package request

import "foodcal/business/users"

type Users struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Dob      string `json:"dob"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user UserLogin) ToDomainLogin() *users.Domain {
	return &users.Domain{
		Email:    user.Email,
		Password: user.Password,
	}
}

func (user Users) ToDomainRegister() *users.Domain {
	return &users.Domain{
		Username: user.Username,
		Address:  user.Address,
		Phone:    user.Phone,
		Password: user.Password,
		Email:    user.Email,
		Dob:      user.Dob,
	}
}
