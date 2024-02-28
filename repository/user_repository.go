package repository

type User struct {
	Name string
}

type UserRepository interface {
	Current() (User, error)
}

type userRepository struct{}

func NewUserRepository() (UserRepository, error) {
	return &userRepository{}, nil
}

func (u *userRepository) Current() (User, error) {
	return User{Name: "Jordan"}, nil
}
