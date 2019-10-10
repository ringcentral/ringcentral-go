package definitions

// ReplyParty Reply Party
type ReplyParty struct {
	Id string `json:"id"`
	Status CallStatusInfo `json:"status"`
	Muted bool `json:"muted"`
	StandAlone bool `json:"standAlone"`
	Park ParkInfo `json:"park"`
	From PartyInfo `json:"from"`
	To PartyInfo `json:"to"`
	Owner OwnerInfo `json:"owner"`
	Direction string `json:"direction"`
}
