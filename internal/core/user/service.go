package user

import (
	"context"
)

// Repository is a user repository.
type Repository interface {
	Create(ctx context.Context, dto *User) (*User, error)
}

// Service is a user service interface.
type Service struct {
	userRepository Repository
}

// NewService creates a new user service.
func NewService(userRepository Repository) Service {
	return Service{
		userRepository: userRepository,
	}
}

// todo do it properly
// Create creates a new user.
func (s Service) Create(ctx context.Context, dto *User) (*User, error) {

	// user := User{
	// 	Email:    dto.Email,
	// 	Username: dto.Username,
	// 	IIN:      dto.IIN,
	// 	BIN:      dto.BIN,
	// }

	// user, err := s.userRepository.Create(ctx, user)
	// if err != nil {
	// 	return &User{}, fmt.Errorf("can not create user: %w", err)
	// }

	// return &user, nil
	return nil, nil
}

// todo do it properly
// Login provides user login.
func (s Service) Login(context.Context, *User) (*User, error) {

	return nil, nil
}

// todo do it properly
func (s Service) LogOut(ctx context.Context, user *User) error {
	return nil
}
