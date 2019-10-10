package definitions

// MeetingsResource Meetings Resource
type MeetingsResource struct {
	Uri string `json:"uri"`
	Records []MeetingResponseResource `json:"records"`
	Paging MeetingsPagingInfo `json:"paging"`
	Navigation MeetingsNavigationInfo `json:"navigation"`
}
