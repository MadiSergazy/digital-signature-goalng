package user

import (
	"context"
	"fmt"

	// "mado/internal"
	"mado/internal/auth"
	"mado/internal/auth/model"
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
func (s Service) Login(requirements model.LoginRequirements) (*User, error) {
	// ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	// defer cancel()

	signature := auth.GetNonceSignature(requirements.QrSigner)
	// nonce

	req := model.AuthRequest{
		Nonce:     requirements.Nonce,
		Signature: signature,
		External:  true,
	}

	response, err := model.Authentification(req)
	if err != nil {
		fmt.Println("Authentication error:", err)
	}
	fmt.Println(response)

	return nil, nil
}

// todo do it properly
func (s Service) LogOut(ctx context.Context, user *User) error {
	return nil
}
