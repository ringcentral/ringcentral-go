package definitions

// GetAccountInfoResponse Get Account Info Response
type GetAccountInfoResponse struct {
	Id string `json:"id"`
	Uri string `json:"uri"`
	MainNumber string `json:"mainNumber"`
	Operator GetExtensionInfoResponse `json:"operator"`
	PartnerId string `json:"partnerId"`
	ServiceInfo ServiceInfo `json:"serviceInfo"`
	SetupWizardState string `json:"setupWizardState"`
	Status string `json:"status"`
	StatusInfo AccountStatusInfo `json:"statusInfo"`
	RegionalSettings RegionalSettings `json:"regionalSettings"`
	Federated bool `json:"federated"`
	OutboundCallPrefix int `json:"outboundCallPrefix"`
	Cfid string `json:"cfid"`
	Limits AccountLimits `json:"limits"`
}
