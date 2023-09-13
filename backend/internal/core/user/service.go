package user

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"time"

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

const (
	block   = "Блок данных на подпись"
	baseURL = "https://sigex.kz"
)

// NewService creates a new user service.
func NewService(userRepository Repository) Service {
	return Service{
		userRepository: userRepository,
	}
}

func (s Service) Login(requirements model.LoginRequirements) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	signature := auth.GetNonceSignature(requirements.QrSigner)

	req := model.AuthRequest{
		Nonce:     requirements.Nonce,
		Signature: signature,
		External:  true,
	}

	response, err := authentification(req)
	if err != nil {
		fmt.Println("Authentication error:", err)
	}
	fmt.Println(response)
	fmt.Println("email:", response.Email)
	fmt.Println("IIN:", response.UserID)
	fmt.Println("BIN:", response.BusinessID)
	fmt.Println("Name:", getName(response.Subject))
	user := &User{Username: getName(response.Subject), IIN: &response.UserID, Email: &response.Email, BIN: &response.BusinessID}
	user, err = s.userRepository.Create(ctx, user)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return user, nil
}

func authentification(request model.AuthRequest) (*model.AuthResponse, error) {

	requestData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	response, err := http.Post(baseURL+"/api/auth", "application/json", bytes.NewReader(requestData))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Server returned status '%d: %s", response.StatusCode, response.Status)
	}

	var responseJSON model.AuthResponse
	err = json.NewDecoder(response.Body).Decode(&responseJSON)
	if err != nil {
		return nil, err
	}

	return &responseJSON, nil
}

func getName(input string) *string {

	// Define regular expressions to match CN and GIVENNAME values
	cnRegex := regexp.MustCompile(`CN=([^,]+)`)
	givenNameRegex := regexp.MustCompile(`GIVENNAME=([^,]+)`)

	// Find CN and GIVENNAME values using regular expressions
	cnMatch := cnRegex.FindStringSubmatch(input)
	givenNameMatch := givenNameRegex.FindStringSubmatch(input)

	// Check if both CN and GIVENNAME values were found
	if len(cnMatch) > 1 && len(givenNameMatch) > 1 {
		cnValue := cnMatch[1]
		givenNameValue := givenNameMatch[1]

		// Print the extracted values
		result := fmt.Sprintf("%s %s", cnValue, givenNameValue)
		return &result
	} else {
		fmt.Println("CN and/or GIVENNAME not found in the input string.")
	}
	return nil
}
