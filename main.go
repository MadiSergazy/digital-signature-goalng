package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"mado/helpers"
	"mado/internal"
	"mado/models"
)

const baseURL = "https://sigex.kz"

type ResponseStruct struct {
	Nonce string `json:"nonce"`
}

// we did't need timeout because when 15min life time of QRCode is expired sigex will message to us
func main() {
	nonce, err := MakeRequest()
	if err != nil {
		fmt.Println("ERROR OCCURED")
	}
	fmt.Println("nonce:", nonce)
}

func MakeRequest() (string, error) {
	// Define the URL you want to send the POST request to
	url := baseURL + "/api/auth"

	// Create a map or struct containing the data you want to send as JSON
	data := map[string]interface{}{}

	// Convert the data to JSON format
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return "", err
	}

	// Create a new HTTP request with the POST method, URL, and request body
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "", err
	}

	// Set the request headers if needed (e.g., content-type)
	req.Header.Set("Content-Type", "application/json")

	// Create an HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return "", err
	}
	defer resp.Body.Close()

	// Read and process the response
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return "", err
	}

	// Print the response status code and body
	fmt.Println("Status Code:", resp.Status)
	fmt.Println("Response Body:", string(responseBody))
	var responseObj ResponseStruct
	if err := json.Unmarshal(responseBody, &responseObj); err != nil {
		fmt.Println("Error unmarshaling response:", err)
		return "", err
	}
	nonceValue := responseObj.Nonce
	return nonceValue, nil
}

func dinit() {
	signatures, dataBytes, dataToSignBase64 := firstThreStep() //regist send data and get signatures

	if len(signatures) > 0 {
		// fmt.Println("Signature:", signatures[0])
		signature := signatures[0]
		documentRegistrationRequest := models.NewDocumentRegistrationRequest(
			"document title",
			"document description",
			"cms",
			signature,
			[]string{"saitamenter@gmail.com"}, //nil,
			models.DocumentSettings{
				Private:                   false,
				SignaturesLimit:           0,
				SwitchToPrivateAfterLimit: false,
				Unique:                    []string{"iin"},
				StrictSignersRequirements: false,
				// SignersRequirements:  ,

			},
		)

		documentRegistrationResponse, err := documentRegistrationRequest.RegisterDocument(baseURL)
		if err != nil {
			fmt.Println("documentRegistrationRequest Error:", err)
			return
		}
		fmt.Println("Registered DocumentID: ", documentRegistrationResponse.DocumentID)

		var docRes models.DocumentHashesResponse
		docResponse, err := docRes.FixingDocumentHashes(documentRegistrationResponse.DocumentID, dataBytes, baseURL) //the reason why we did't use []byte(documentRegistrationResponse) because it will return doc if it was senden with inside signature
		helpers.ErrorHandlingWithRerurn(err, "FixingDocumentHashes Error: ")
		fmt.Println("FixingDocumentHashes DocumentID: ", docResponse.DocumentID)
		fmt.Println("FixingDocumentHashes Digests: ", docResponse.Digests)

		///TODO automate this GETTING SECOND SIGNATURE
		qrSigner := internal.NewQRSigningClientCMS("Тестовое подписание", false, baseURL)

		err = qrSigner.AddDataToSign([]string{"Блок данных на подпись", "Блок данных на подпись", "Блок данных на подпись"}, dataToSignBase64, nil, true)
		if err != nil {
			fmt.Println("Could not read file: ", err)
			return
		}

		// Register QR signing
		dateUrl, err := qrSigner.RegisterQRSinging()
		if err != nil {
			fmt.Println("RegisterQRSinging Error:", err)
			return
		}
		fmt.Println("Second man DateURL: ", dateUrl)
		eGovMobileLaunchLink := qrSigner.GetEGovMobileLaunchLink()
		// eGovBusinessLaunchLink := qrSigner.GetEGovBusinessLaunchLink()
		fmt.Println("Second maneGov Mobile Launch Link2:", eGovMobileLaunchLink)

		newSignature, err := qrSigner.GetSignatures(nil)
		if err != nil {
			fmt.Println("GetSignatures Error:", err)
			return
		}

		fmt.Println("(len(newSignature: ", len(newSignature))
		// documentRegistrationResponse.DocumentID

		fmt.Println("SECOND MAN docResponse.DocumentID: ", docResponse.DocumentID)
		// fmt.Println("SECOND MAN docRes.DocumentID: ", docRes.DocumentID) // docRes had nothing

		addSignatureResponse, err := internal.AddSignatureToDocument(docResponse.DocumentID, newSignature[0], baseURL) //todo  docRes.DocumentID before was it
		helpers.ErrorHandlingWithRerurn(err, "addSignatureResponse Error: ")
		fmt.Println("Second man addSignatureResponse DocumentID:", addSignatureResponse.DocumentID)

		// docID: BqHfcYdPvidBIvl8 //addSignatureResponse.DocumentID, //TODO: CHECK
		documentCard := internal.NewGenerateElectronicDocumentCardRequest(
			addSignatureResponse.DocumentID,
			"Petition.pdf",
			false,
			false,
			false,
			"kk/ru",
			dataBytes,
		)

		//Document ID: PLViD43c6HgkbC1x
		documentCardResponse, err := documentCard.GenerateElectronicDocumentCard(baseURL)
		helpers.ErrorHandlingWithRerurn(err, "documentCardResponse Error: ")
		fmt.Println("documentCardResponse:", documentCardResponse)

		helpers.DecodeBase64ToPDF(documentCardResponse.DDC, "output.pdf")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("PDF file has been saved !")

	} else {
		fmt.Println("No signatures.")
	}
}

func firstThreStep() ([]string, []byte, string) {
	// Usage example:
	qrSigner := internal.NewQRSigningClientCMS("Тестовое подписание", false, baseURL)
	// Add data to sign (encoded in base64)
	// dataToSignBase64 := "MTEK"
	dataToSignBase64, dataBytes, err := helpers.ReadPdf("someFile.pdf")
	if err != nil {
		fmt.Println("Could not read file: ", err)
		return nil, nil, ""
	}

	err = qrSigner.AddDataToSign([]string{"Блок данных на подпись", "Блок данных на подпись", "Блок данных на подпись"}, dataToSignBase64, nil, true)
	if err != nil {
		fmt.Println("Could not read file: ", err)
		return nil, nil, ""
	}

	// Register QR signing
	dataURL, err := qrSigner.RegisterQRSinging()
	if err != nil {
		fmt.Println("RegisterQRSinging Error:", err)
		return nil, nil, ""
	}
	fmt.Println("First man RegisterQRSinging dataURL: ", dataURL)

	/*//todo maybe split GetSignatures
	signURL, err := qrSigner.SendDataToSign(dataURL)
	if err != nil {
		fmt.Println("SendDataToSign Error:", err)
		return
	}
	fmt.Println("signURL: ", signURL)
	*/

	// qrCodeDataString := "data:image/gif;base64," + qrCode
	// fmt.Println("QR Code Image Data URL:", qrCodeDataString)

	// Get launch links for eGov mobile and eGov Business

	eGovMobileLaunchLink := qrSigner.GetEGovMobileLaunchLink()
	// eGovBusinessLaunchLink := qrSigner.GetEGovBusinessLaunchLink()
	fmt.Println("FIRST eGov Mobile Launch Link:", eGovMobileLaunchLink)
	// fmt.Println("eGov Business Launch Link:", eGovBusinessLaunchLink)

	/*
	   {"documentsToSign":[{"document":{"file":{"data":"MTEK","mime":"@file/pdf"}},"documentXml":"\u003cgroupId\u003e2\u003c/groupId\u003e","id":1,"meta":null,"nameEn":"Блок данных на подпись","nameKz":"Блок данных на подпись","nameRu":"Блок данных на подпись"}],"signMethod":"CMS_SIGN_ONLY"}

	*/

	// Get signatures
	signatures, err := qrSigner.GetSignatures(nil)
	if err != nil {
		fmt.Println("GetSignatures Error:", err)
		return nil, nil, ""
	}
	return signatures, dataBytes, dataToSignBase64
}
