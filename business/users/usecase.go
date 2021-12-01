package users

import (
	"context"
	"errors"

	"foodcal/app/middlewares"
	"time"
)

type UserUseCases struct {
	repo    UserRepoInterface
	ctx     time.Duration
	jwtAuth *middlewares.ConfigJWT
}

func NewUseCase(userRepo UserRepoInterface, contextTimeout time.Duration, jwtauth *middlewares.ConfigJWT) UserUsecaseInterface {
	return &UserUseCases{
		repo:    userRepo,
		ctx:     contextTimeout,
		jwtAuth: jwtauth,
	}
}

func (usecase *UserUseCases) Register(domain *Domain) (Domain, error) {
	user, err := usecase.repo.Register(domain)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (usecase *UserUseCases) Login(domain Domain, ctx context.Context) (Domain, error) {
	if domain.Email == "" {
		return Domain{}, errors.New("Email empty")
	}
	if domain.Password == "" {
		return Domain{}, errors.New("Password empty")
	}
	user, err := usecase.repo.Login(domain, ctx)
	if err != nil {
		return Domain{}, err
	}
	user.Token = usecase.jwtAuth.GenererateToken(domain.Id)

	return user, nil
}
