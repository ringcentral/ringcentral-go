package definitions

// AutomaticLocationUpdatesTaskInfo Automatic Location Updates Task Info
type AutomaticLocationUpdatesTaskInfo struct {
	Id string `json:"id"`
	Status string `json:"status"`
	CreationTime string `json:"creationTime"`
	LastModifiedTime string `json:"lastModifiedTime"`
	Type string `json:"type"`
	Result TaskResultInfo `json:"result"`
}
