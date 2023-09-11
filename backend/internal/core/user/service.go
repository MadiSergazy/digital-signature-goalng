package user

import (
	"context"
	"errors"
	"fmt"

	"mado/internal"
	"mado/internal/auth"
	"mado/internal/auth/model"
)

var parsingError = errors.New("Error parsing JSON: ")

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
func (s Service) Login(ctx context.Context, qrSigner *internal.QRSigningClientCMS, nonce *string) (*User, error) {
	// ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	// defer cancel()

	signature := auth.GetNonceSignature(qrSigner)
	// nonce

	req := model.AuthRequest{
		Nonce:     nonce,
		Signature: signature,
		External:  true,
	}

	response, err := model.Authentification(req)
	if err != nil {
		fmt.Println("Authentication error:", err)
	}
	fmt.Println(response)

	return s.userRepository.Create(ctx, &User{ // temporary it will return nil
		Username: &response.Subject,
		Email:    &response.Email,
		IIN:      &response.UserID,
		BIN:      &response.BusinessID,
	})

}

// todo do it properly
func (s Service) LogOut(ctx context.Context, user *User) error {
	return nil
}
