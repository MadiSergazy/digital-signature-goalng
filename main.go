package main

import (
	"fmt"

	"mado/helpers"
	"mado/internal"
	"mado/models"
)

const baseURL = "https://sigex.kz"

// we did't need timeout because when 15min life time of QRCode is expired sigex will message to us
func main() {
	// Usage example:
	qrSigner := internal.NewQRSigningClientCMS("Тестовое подписание", false, baseURL)
	// Add data to sign (encoded in base64)
	// dataToSignBase64 := "MTEK"
	dataToSignBase64, dataBytes, err := helpers.ReadPdf("someFile.pdf")
	if err != nil {
		fmt.Println("Could not read file: ", err)
		return
	}

	err = qrSigner.AddDataToSign([]string{"Блок данных на подпись", "Блок данных на подпись", "Блок данных на подпись"}, dataToSignBase64, nil, true)
	if err != nil {
		fmt.Println("Could not read file: ", err)
		return
	}

	// Register QR signing
	dataURL, err := qrSigner.RegisterQRSinging()
	if err != nil {
		fmt.Println("RegisterQRSinging Error:", err)
		return
	}
	fmt.Println("dataURL: ", dataURL)

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
	fmt.Println("eGov Mobile Launch Link:", eGovMobileLaunchLink)
	// fmt.Println("eGov Business Launch Link:", eGovBusinessLaunchLink)

	/*
	   {"documentsToSign":[{"document":{"file":{"data":"MTEK","mime":"@file/pdf"}},"documentXml":"\u003cgroupId\u003e2\u003c/groupId\u003e","id":1,"meta":null,"nameEn":"Блок данных на подпись","nameKz":"Блок данных на подпись","nameRu":"Блок данных на подпись"}],"signMethod":"CMS_SIGN_ONLY"}

	*/

	// Get signatures
	signatures, err := qrSigner.GetSignatures(nil)
	if err != nil {
		fmt.Println("GetSignatures Error:", err)
		return
	}

	if len(signatures) > 0 {
		fmt.Println("Signature:", signatures[0])
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

		var docRes models.DocumentDataResponse
		docResponse, err := docRes.PostDocumentData(documentRegistrationResponse.DocumentID, dataBytes, baseURL) //the reason why we did't use []byte(documentRegistrationResponse) because it will return doc if it was senden with inside signature
		helpers.ErrorHandlingWithRerurn(err, "documentDataResponse Error: ")
		fmt.Println("documentDataResponse:", docResponse)

		///TODO automate this GETTING SECOND SIGNATURE
		qrSigner := internal.NewQRSigningClientCMS("Тестовое подписание", false, baseURL)

		err = qrSigner.AddDataToSign([]string{"Блок данных на подпись", "Блок данных на подпись", "Блок данных на подпись"}, dataToSignBase64, nil, true)
		if err != nil {
			fmt.Println("Could not read file: ", err)
			return
		}

		// Register QR signing
		_, err = qrSigner.RegisterQRSinging()
		if err != nil {
			fmt.Println("RegisterQRSinging Error:", err)
			return
		}

		eGovMobileLaunchLink := qrSigner.GetEGovMobileLaunchLink()
		// eGovBusinessLaunchLink := qrSigner.GetEGovBusinessLaunchLink()
		fmt.Println("eGov Mobile Launch Link2:", eGovMobileLaunchLink)

		newSignature, err := qrSigner.GetSignatures(nil)
		if err != nil {
			fmt.Println("GetSignatures Error:", err)
			return
		}

		fmt.Println("(len(newSignature: ", len(newSignature))
		// documentRegistrationResponse.DocumentID
		addSignatureResponse, err := internal.AddSignatureToDocument(docRes.DocumentID, newSignature[0], baseURL)
		helpers.ErrorHandlingWithRerurn(err, "addSignatureResponse Error: ")
		fmt.Println("addSignatureResponse:", addSignatureResponse)
	} else {
		fmt.Println("No signatures.")
	}
}
