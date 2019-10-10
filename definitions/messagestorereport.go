package definitions

// MessageStoreReport Message Store Report
type MessageStoreReport struct {
	Id string `json:"id"`
	Uri string `json:"uri"`
	Status string `json:"status"`
	AccountId string `json:"accountId"`
	ExtensionId string `json:"extensionId"`
	CreationTime string `json:"creationTime"`
	LastModifiedTime string `json:"lastModifiedTime"`
	DateTo string `json:"dateTo"`
	DateFrom string `json:"dateFrom"`
}
