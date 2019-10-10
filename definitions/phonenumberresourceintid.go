package definitions

// PhoneNumberResourceIntId Phone Number Resource Int Id
type PhoneNumberResourceIntId struct {
	Id int `json:"id"`
	Country CountryResource `json:"country"`
	Extension PhoneNumberExtensionInfo `json:"extension"`
	Label string `json:"label"`
	Location string `json:"location"`
	PaymentType string `json:"paymentType"`
	PhoneNumber string `json:"phoneNumber"`
	Status string `json:"status"`
	UsageType string `json:"usageType"`
	Type string `json:"type"`
	ReservationId string `json:"reservationId"`
}
