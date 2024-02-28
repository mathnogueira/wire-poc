package paidrepository

import "github.com/mathnogueira/ioc/repository"

type userRepository struct{}

func NewUserRepository() (repository.UserRepository, error) {
	return &userRepository{}, nil
}

func (u *userRepository) Current() (repository.User, error) {
	return repository.User{Name: "Premium Jordan"}, nil
}
