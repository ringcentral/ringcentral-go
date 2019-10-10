package definitions

// CallParty Call Party
type CallParty struct {
	Id string `json:"id"`
	Status CallStatusInfo `json:"status"`
	Muted bool `json:"muted"`
	StandAlone bool `json:"standAlone"`
	Park ParkInfo `json:"park"`
	From PartyInfo `json:"from"`
	To PartyInfo `json:"to"`
	Owner OwnerInfo `json:"owner"`
	Direction string `json:"direction"`
	ConferenceRole string `json:"conferenceRole"`
	RingOutRole string `json:"ringOutRole"`
	RingMeRole string `json:"ringMeRole"`
	Recordings []RecordingInfo `json:"recordings"`
}
