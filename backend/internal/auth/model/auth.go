package model

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"mado/internal"
	"net/http"
)

const (
	block   = "Блок данных на подпись"
	baseURL = "https://sigex.kz"
)

type AuthResponse struct {
	UserID                  string           `json:"userId"`
	BusinessID              string           `json:"businessId"`
	Email                   string           `json:"email"`
	Subject                 string           `json:"subject"`
	SubjectStructure        [][]Attribute    `json:"subjectStructure"`
	SubjectAltName          string           `json:"subjectAltName"`
	SubjectAltNameStructure []SubjectAltName `json:"subjectAltNameStructure"`
	SignAlgorithm           string           `json:"signAlgorithm"`
	KeyStorage              string           `json:"keyStorage"`
	PolicyIds               []string         `json:"policyIds"`
	ExtKeyUsages            []string         `json:"extKeyUsages"`
}

type Attribute struct {
	Oid        string `json:"oid"`
	Name       string `json:"name"`
	ValueInB64 bool   `json:"valueInB64"`
	Value      string `json:"value"`
}

type SubjectAltName struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type AuthRequest struct {
	Nonce     *string `json:"nonce"`
	Signature *string `json:"signature"`
	External  bool    `json:"external"`
}

func Authentification(request AuthRequest) (*AuthResponse, error) {
	// Convert the request to JSON
	requestData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	// Make a POST request to the endpoint
	response, err := http.Post(baseURL+"/api/auth", "application/json", bytes.NewReader(requestData))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Check the response status code
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Server returned status '%d: %s", response.StatusCode, response.Status)
	}

	// Parse the JSON response
	var responseJSON AuthResponse
	err = json.NewDecoder(response.Body).Decode(&responseJSON)
	if err != nil {
		return nil, err
	}

	return &responseJSON, nil
}

type LoginRequirements struct {
	Context  context.Context              `json:"context"`
	QrSigner *internal.QRSigningClientCMS `json:"qrsigner"`
	Nonce    *string                      `json:"nonce"`
}