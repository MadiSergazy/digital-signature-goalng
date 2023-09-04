package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"mado/internal"
)

const (
	block   = "Блок данных на подпись"
	baseUrl = "https://sigex.kz"
)

func PreparationStep() (nonce *string, signature *string) {

	body, _ := json.Marshal(map[string]interface{}{})

	response, err := http.Post(baseUrl+"/api/auth", "application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("Error getting аутентификация, подготовительный этап:", err)
		return nil, nil
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Printf("Server returned status '%d: %s\n'", response.StatusCode, response.Status)
		return nil, nil
	}

	var responseJSON map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&responseJSON)
	if err != nil {
		fmt.Println("decode err:", err)
		return nil, nil
	}
	fmt.Println(responseJSON)

	// var ErrorResp models.ErrorResponse
	// errMessage := ErrorResp.GetHumanReadableErrorMessageByResponse(responseJSON)
	// if errMessage != "" {
	// 	fmt.Println("err resp:", err)
	// 	return nil, nil
	// }

	if nonce, ok := responseJSON["nonce"].(string); ok {
		qrSigner := internal.NewQRSigningClientCMS("Тестовое подписание", false, baseUrl)
		err = qrSigner.AddDataToSign([]string{block, block, block}, nonce, nil, true)
		if err != nil {
			fmt.Println("Could not read file: ", err)
			return nil, nil
		}

		// Register QR signing
		dataURL, err := qrSigner.RegisterQRSinging()
		if err != nil {
			fmt.Println("RegisterQRSinging Error:", err)
			return nil, nil
		}
		fmt.Println("First man RegisterQRSinging dataURL: ", dataURL)

		eGovMobileLaunchLink := qrSigner.GetEGovMobileLaunchLink()
		fmt.Println("For log-in eGov Mobile Launch Link:", eGovMobileLaunchLink)
		signatures, err := qrSigner.GetSignatures(nil)
		if err != nil {
			fmt.Println("GetSignatures Error:", err)
			return nil, nil
		}

		return &nonce, &signatures[0]

	} else {
		fmt.Println("err did't have nonce in resp:", err)
		return nil, nil
	}

}
