package models

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"mado/helpers"
)

type ErrorResponse struct {
	Message   string `json:"message,omitempty"`
	RequestID int64  `json:"requestID,omitempty"`
}

// Map of error messages to their human-readable versions
var errorMap = map[string]string{
	"Failed to build certificate chain":                      "Не удалось построить цепочку сертификатов.",
	"Failed to build digital document card":                  "Не удалось сформировать карточку электронного документа.",
	"Failed to parse JSON":                                   "Не удалось разобрать JSON запрос.",
	"Failed to parse digital document card":                  "Не удалось разобрать карточку электронного документа.",
	"Failed to parse signature":                              "Не удалось разобрать подпись.",
	"Failed to prepare archived data":                        "Ошибка подготовки архива.",
	"Feature is not available on current subscription level": "Функция недоступна на данном уровне подписки.",
	"HTTP Content-Length request header not set":             "Не указан HTTP заголовок Content-Length.",
	"Invalid API route":                                      "Некорректный маршрут.",
	"Invalid HTTP request headers":                           "Некорректный HTTP заголовок запроса.",
	"Invalid HTTP request method":                            "Некорректный HTTP метод.",
	"Invalid JSON request structure":                         "Некорректная структура JSON запроса.",
	"Invalid QR code logo format":                            "Неверный формат логотипа QR кода.",
	"Invalid QR code logo resolution":                        "Неподдерживаемое разрешение логотипа QR кода.",
	"Invalid QR code logo size":                              "Неверный размер логотипа QR кода.",
	"Invalid QR signing state":                               "Неверное состояние QR подписания.",
	"Invalid URL query parameter":                            "Некорректный URL параметр запроса.",
	"Invalid authentication nonce":                           "Неверный nonce аутентификации.",
	"Invalid certificate status":                             "Некорректный статус сертификата.",
	"Invalid document":                                       "Некорректный документ.",
	"Invalid document Title":                                 "Неверный заголовок документа.",
	"Invalid document identifier":                            "Некорректный идентификатор документа.",
	"Invalid file name":                                      "Некорректное имя файла.",
	"Invalid format of provided document settings":           "Некорректный формат объекта настроек документа.",
	"Invalid signature":                                      "Некорректная подпись.",
	"Invalid signature export format":                        "Некорректный формат экспорта.",
	"Invalid signature identifier":                           "Некорректный идентификатор подписи.",
	"Nonconforming order of elements in one of ASN.1 SET tags of the signature": "Неверный порядок элементов в одном из ASN.1 тегов SET подписи.",
	"Not enough TLS certificates":                   "Недостаточно TLS сертификатов.",
	"Not enough authorities":                        "Недостаточно полномочий.",
	"Not supported document type":                   "Неподдерживаемый тип документа.",
	"OCSP server problem":                           "Проблема с OCSP сервером.",
	"One of the provided IINs is in invalid format": "Некорректный формат одного из ИИН переданных в запросе.",
	"Access denied":                                 "Доступ запрещен.",
	"Already authenticated with TLS certificate":    "Аутентификация уже выполнена по TLS сертификату.",
	"Archive not found":                             "Архив не найден.",
	"Authentication required":                       "Требуется аутентификация.",
	"Bad signer certificate":                        "Плохой сертификат подписавшего.",
	"Certificate issuer is unknown":                 "Издатель сертификата неизвестен.",
	"Certificate signature algorithm not supported": "Алгоритм подписи сертификата не поддерживается.",
	"Digest algorithm not supported":                "Алгоритм хеширования не поддерживается.",
	"Digests values update required":                "Необходимо обновить хеши документа.",
	"Digital document card contains signatures that correspond to different registered documents": "В карточке электронного документа присутствуют подписи, зарегистрированные под разными электронными документами.",
	"Digital document card has signature in invalid format":                                       "В карточке электронного документа присутствует подпись не подходящего формата.",
	"Digital document card original document is not registered":                                   "Документ в карточке электронного документа не зарегистрирован на сервисе.",
	"Document data archival is started by another user":                                           "Архивирование данных выполняется другим пользователем.",
	"Document digests are already known":                                                          "Хеши документа уже известны.",
	"Document digests are not known":                                                              "Хеши документа не известны.",
	"Document does not have signatures in required format":                                        "Под документом не зарегистрировано ни одной подписи подходящего для данной операции формата.",
	"Document has more signatures than new limit":                                                 "Количество подписей под документом превышает устанавливаемое ограничение.",
	"Document has signatures in not supported by this operation format":                           "Под документом зарегистрированы подписи не подходящего для данной операции формата.",
	"Document not found":                "Документ не найден.",
	"Document signatures limit reached": "Превышено ограничение на количество подписей.",
	"Failed to archive document data":   "Ошибка архивирования документа.",
	"One of the provided TLS certificates can not be used for authentication":                     "Один из TLS сертификатов не может быть использован для аутентификации.",
	"One of the provided URL addresses is in invalid format":                                      "Некорректный формат одного из URL адресов.",
	"One of the provided authority OIDs is in invalid format":                                     "Некорректный формат одного из OID-ов.",
	"One of the provided certificate indices is invalid":                                          "Один из индексов сертификатов не верен.",
	"One of the provided notification email addresses is in invalid format":                       "Некорректный формат одного из адресов электронной почты для отправки уведомлений.",
	"Other endpoint connected":                                                                    "Подключена другая конечная точка.",
	"QR code versions lower than 11 are not allowed":                                              "Версии QR кодов ниже 11 не поддерживаются.",
	"QR signing operation is not active":                                                          "Операция QR подписания не активна.",
	"QR signing operation timeout":                                                                "Таймаут операции QR подписания.",
	"Request body is too large":                                                                   "Размер запроса слишком велик.",
	"Request field size is too large":                                                             "Размер поля слишком велик.",
	"Request rate limit reached":                                                                  "Достигнуто ограничение частоты запросов.",
	"Signature algorithm not supported":                                                           "Алгоритм подписи не поддерживается.",
	"Signature contains invalid OCSP data":                                                        "Подпись содержит поврежденные или некорректные данные OCSP.",
	"Signature contains invalid TSP time stamp":                                                   "Подпись содержит поврежденную или некорректную метку времени TSP.",
	"Signature does not conform to document settings requirements":                                "Подпись не удовлетворяет требованиям в настройках документа.",
	"Signature does not correspond to the document":                                               "Подпись не соответствует документу.",
	"Signature type is not supported":                                                             "Не поддерживаемый тип подписи.",
	"Signer certificate expired or not yet valid":                                                 "Срок действия сертификата уже истек, либо еще не наступил.",
	"Some of digital document card signatures are not registered":                                 "В карточке электронного документа присутствует не зарегистрированная на сервисе подпись.",
	"TLS certificate used for authentication was disabled":                                        "TLS сертификат, использованный для аутентификации, отключен.",
	"TSP server problem":                                                                          "Проблема с TSP сервером.",
	"The processed data size does not match the one passed in HTTP Content-Length request header": "Размер обработанных данных не равен размеру, переданному в HTTP заголовке Content-Length.",
	"This signature has already been submitted":                                                   "Данная подпись уже была зарегистрирована.",
	"Too many notification email addresses provided":                                              "Указано слишком много адресов электронной почты для рассылки уведомлений.",
	"Too much users use QR signing":                                                               "Слишком много пользователей выполняют QR подписание.",
	"Unexpected error":                                                                            "Непредвиденная ошибка.",
	"User does not represent an organization":                                                     "Пользователь не является представителем организации.",
	"eGov mobile data exchange failed":                                                            "Ошибка обмена данными с eGov mobile.",
}

func (errResp *ErrorResponse) ParseErrorResponse(jsonData []byte) (*ErrorResponse, error) {
	err := json.Unmarshal(jsonData, &errResp)
	if err != nil {
		return nil, err
	}

	return errResp, nil
}

// GetHumanReadableErrorMessage converts the given error message to its human-readable version
func (errResp *ErrorResponse) GetHumanReadableErrorMessageByResponse(response map[string]interface{}) string {
	if msg, ok := response["message"].(string); ok {
		if humanReadableMsg, exists := errorMap[msg]; exists {
			return humanReadableMsg
		}
		// If the error message is not found in the map, return the original error message
		return msg
	}
	// If the 'message' field does not exist in the response, return an empty string
	return ""
}

func (errResp *ErrorResponse) GetHumanReadableErrorMessage() string {
	if humanReadableMsg, exists := errorMap[errResp.Message]; exists {
		return humanReadableMsg
	}
	// If the error message is not found in the map, return the original error message
	return errResp.Message
}

// //////////////////////////////
type MetaData struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type DocumentFile struct {
	Mime string `json:"mime"`
	Data string `json:"data"`
}

type Document struct {
	File        DocumentFile `json:"file"`
	DocumentXml string       `json:"documentXml"`
}

type DocumentToSign struct {
	ID       int        `json:"id"`
	NameRu   string     `json:"nameRu"`
	NameKz   string     `json:"nameKz"`
	NameEn   string     `json:"nameEn"`
	Meta     []MetaData `json:"meta"`
	Document Document   `json:"document"`
}

type ResponseGettingSignatureData struct {
	Message         string           `json:"messgage,omitempty"`
	SignMethod      string           `json:"signMethod"`
	DocumentsToSign []DocumentToSign `json:"documentsToSign"`
}

func (responseData *ResponseGettingSignatureData) GetSignaturesFromResponse(response *http.Response) ([]string, error) {
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Server returned status '%d: %s'", response.StatusCode, response.Status)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return nil, err
	}

	if responseData.Message != "" {
		return nil, errors.Join(fmt.Errorf("Error in geting signature"), fmt.Errorf(responseData.Message))
	}

	if responseData.SignMethod != "CMS_SIGN_ONLY" {
		return nil, fmt.Errorf("Invalid sign method in the response")
	}

	// Extract the signatures from the responseData.DocumentsToSign
	signatures := make([]string, 0, len(responseData.DocumentsToSign))
	for _, doc := range responseData.DocumentsToSign {
		signatures = append(signatures, doc.Document.File.Data)
	}
	fmt.Printf("DocumentsToSign ID: %d\n", responseData.DocumentsToSign[0].ID)
	return signatures, nil
}

type QRResponse struct {
	ExpireAt               int64  `json:"expireAt"`
	DataURL                string `json:"dataURL"`
	SignURL                string `json:"signURL"`
	EGovMobileLaunchLink   string `json:"eGovMobileLaunchLink"`
	EGovBusinessLaunchLink string `json:"eGovBusinessLaunchLink"`
	QRCode                 string `json:"qrCode"`
}

func (qr *QRResponse) printQRResponse(response *http.Response) {
	// Print the status code and status text
	fmt.Printf("Status: %d %s\n", response.StatusCode, response.Status)

	// Print all the response headers
	for key, value := range response.Header {
		fmt.Printf("%s: %s\n", key, value)
	}

	// Read and print the response body
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	defer response.Body.Close()

	var qrResponse QRResponse
	err = json.Unmarshal(bodyBytes, &qrResponse)
	if err != nil {
		fmt.Println("Error unmarshaling response body:", err)
		return
	}
	fmt.Println("ExpireAt:", qrResponse.ExpireAt)
	helpers.PrintTime(qrResponse.ExpireAt)
	fmt.Println("DataURL:", qrResponse.DataURL)
	fmt.Println("SignURL:", qrResponse.SignURL)
	fmt.Println("EGovMobileLaunchLink:", qrResponse.EGovMobileLaunchLink)
	fmt.Println("EGovBusinessLaunchLink:", qrResponse.EGovBusinessLaunchLink)
	fmt.Println("QRCode:", qrResponse.QRCode)
}

/*

type SendDataToSignResponse struct {
	ExpireAt int64  `json:"expireAt"`
	SignURL  string `json:"signURL"`
}

func (qc *QRSigningClientCMS) SendDataToSign(qrId string) (string, error) {
	response, err := http.Post(qrId, "application/json", nil)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Server returned status '%d: %s'", response.StatusCode, response.Status)
	}

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "", err
	}
	defer response.Body.Close()

	var sendDataToSignResponse SendDataToSignResponse
	err = json.Unmarshal(bodyBytes, &sendDataToSignResponse)
	if err != nil {
		fmt.Println("Error unmarshaling response body:", err)
		return "", err
	}

	return sendDataToSignResponse.SignURL, nil

}

*/

//*!for  POST /api - регистрация нового документа в системе

type DocumentRegistrationRequest struct {
	Title              string                   `json:"title"`
	Description        string                   `json:"description"`
	SignType           string                   `json:"signType,omitempty"`
	Signature          string                   `json:"signature"`
	EmailNotifications EmailNotificationOptions `json:"emailNotifications,omitempty"` //todo maybe add email notifications
	Settings           DocumentSettings         `json:"settings,omitempty"`
}

type EmailNotificationOptions struct {
	To []string `json:"to"`
}

type DocumentSettings struct {
	Private                   bool     `json:"private,omitempty"`
	SignaturesLimit           int      `json:"signaturesLimit,omitempty"`
	SwitchToPrivateAfterLimit bool     `json:"switchToPrivateAfterLimitReached,omitempty"`
	Unique                    []string `json:"unique,omitempty"`
	StrictSignersRequirements bool     `json:"strictSignersRequirements,omitempty"`
	SignersRequirements       []struct {
		IIN string `json:"iin"`
	} `json:"signersRequirements,omitempty"`
}

type DocumentRegistrationResponse struct {
	DocumentID string `json:"documentId"`
	SignID     int    `json:"signId"`
	Data       string `json:"data,omitempty"`

	// ErrorResponse fields
	Message   string `json:"message,omitempty"`
	RequestID int64  `json:"requestID,omitempty"`
}

func NewDocumentRegistrationRequest(title, description, signType, signature string, emailNotifications []string, settings DocumentSettings) DocumentRegistrationRequest {
	return DocumentRegistrationRequest{
		Title:       title,       //* must have
		Description: description, //* must have
		SignType:    signType,    //* must have default cms
		Signature:   signature,   //* must have
		EmailNotifications: EmailNotificationOptions{
			To: emailNotifications,
		},
		Settings: settings,
	}
}

func (docRegReq *DocumentRegistrationRequest) RegisterDocument(baseURL string) (*DocumentRegistrationResponse, error) {
	// Create the request payload
	/*	requestBody := DocumentRegistrationRequest{
		Title:       "document title",
		Description: "document description",
		SignType:    "cms",
		Signature:   signature,
		//EmailNotifications: EmailNotificationOptions{
		//		To: []string{"user@example.com", "other@example.com"},
		//	},
		Settings: DocumentSettings{
			Private:                   false,
			SignaturesLimit:           0,
			SwitchToPrivateAfterLimit: false,
			Unique:                    []string{"iin"},
			StrictSignersRequirements: false,
			// SignersRequirements: []struct {
			// IIN string `json:"iin"`
			// }{
			// {IIN: "IIN112233445566"},
			// },
		},
	}*/

	// Marshal the request payload to JSON
	requestJSON, err := json.Marshal(docRegReq)
	if err != nil {
		return nil, err
	}

	// Create the HTTP request
	url := fmt.Sprintf("%s/api", baseURL) // Replace {id} with the appropriate ID in the actual URL
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestJSON))
	if err != nil {
		return nil, err
	}

	// Set the appropriate headers (Content-Type and Authorization if required)
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var errResp ErrorResponse
	// Check the response status code for errors
	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		errResp, err := errResp.ParseErrorResponse(bodyBytes)
		if err != nil {
			return nil, errors.Join(fmt.Errorf("Error parsing error response: "), err)
		}
		return nil, fmt.Errorf("API error: %s", errResp.GetHumanReadableErrorMessage())
	}

	// Parse the response
	var response DocumentRegistrationResponse
	// err = json.Unmarshal(bodyBytes, &response)
	// if err != nil {
	// 	return nil, err
	// }
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	} else if err == nil && response.Message != "" {
		// If there's an error message, return the ErrorResponse
		return nil, errors.Join(fmt.Errorf("RequestID: %s%s", strconv.FormatInt(response.RequestID, 10), " Error msg: "), fmt.Errorf(response.Message))
	}

	fmt.Println("Document ID:", response.DocumentID)
	fmt.Println("Sign ID:", response.SignID)
	fmt.Println("Data:", response.Data) ///todo ETO POLE EMPTY

	return &response, nil
}

// *Registration HESH
// Define the structure for the response
type DocumentDataResponse struct {
	DocumentID                       string                     `json:"documentId"`
	SignedDataSize                   int                        `json:"signedDataSize"`
	Digests                          map[string]string          `json:"digests"`
	EmailNotifications               *EmailNotificationResponse `json:"emailNotifications,omitempty"`
	AutomaticallyCreatedUserSettings *UserSettingsResponse      `json:"automaticallyCreatedUserSettings,omitempty"`
	DataArchived                     bool                       `json:"dataArchived"`

	// ErrorResponse fields
	Message   string `json:"message,omitempty"`
	RequestID int64  `json:"requestID,omitempty"`
}

type EmailNotificationResponse struct {
	Attached bool   `json:"attached"`
	Message  string `json:"message,omitempty"`
}

type UserSettingsResponse struct {
	UserID                    string `json:"userId"`
	EmailNotificationsEnabled bool   `json:"emailNotificationsEnabled"`
	Email                     string `json:"email"`
	ModifiedAt                int64  `json:"modifiedAt"`
}

// Function to perform the POST request and receive the response
func (docRes DocumentDataResponse) PostDocumentData(id string, document []byte, baseURL string) (*DocumentDataResponse, error) {
	// Construct the URL for the specific document ID
	url := fmt.Sprintf("%s/api/%s/data", baseURL, id)

	// Create a new request with the document data
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(document))
	if err != nil {
		return nil, err
	}

	// Set the Content-Type header to application/octet-stream
	req.Header.Set("Content-Type", "application/octet-stream")

	// Set the Content-Length header with the correct size of the document data
	req.Header.Set("Content-Length", fmt.Sprintf("%d", len(document)))

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check the response status code for errors
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-OK status code: %d", resp.StatusCode)
	}
	var response DocumentDataResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	} else if err == nil && response.Message != "" {
		// If there's an error message, return the ErrorResponse
		return nil, errors.Join(fmt.Errorf("RequestID: %s%s", strconv.FormatInt(response.RequestID, 10), " Error msg: "), fmt.Errorf(response.Message))

	}

	/*
		// Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		// Parse the response body into the DocumentDataResponse struct
		var dataResponse DocumentDataResponse
		err = json.Unmarshal(body, &dataResponse)
		if err != nil {
			return nil, err
		}*/

	return &response, nil
}
