package definitions

// GetCallRecordingResponse Get Call Recording Response
type GetCallRecordingResponse struct {
	Id string `json:"id"`
	ContentUri string `json:"contentUri"`
	ContentType string `json:"contentType"`
	Duration int `json:"duration"`
}
