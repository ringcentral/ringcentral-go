package definitions

// SuperviseCallSession Supervise Call Session
type SuperviseCallSession struct {
	From PartyInfo `json:"from"`
	To PartyInfo `json:"to"`
	Direction string `json:"direction"`
	Id string `json:"id"`
	Muted bool `json:"muted"`
	Owner OwnerInfo `json:"owner"`
	StandAlone bool `json:"standAlone"`
	Status CallStatusInfo `json:"status"`
}
