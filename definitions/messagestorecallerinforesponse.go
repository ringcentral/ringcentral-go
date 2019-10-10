package definitions

// MessageStoreCallerInfoResponse Message Store Caller Info Response
type MessageStoreCallerInfoResponse struct {
	ExtensionNumber string `json:"extensionNumber"`
	Location string `json:"location"`
	MessageStatus string `json:"messageStatus"`
	FaxErrorCode string `json:"faxErrorCode"`
	Name string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
}
