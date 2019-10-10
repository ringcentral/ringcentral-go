package definitions

// IncomingCallEvent Incoming Call Event
type IncomingCallEvent struct {
	Aps APSInfo `json:"aps"`
	Event string `json:"event"`
	Uuid string `json:"uuid"`
	SubscriptionId string `json:"subscriptionId"`
	Timestamp string `json:"timestamp"`
	ExtensionId string `json:"extensionId"`
	Action string `json:"action"`
	SessionId string `json:"sessionId"`
	ServerId string `json:"serverId"`
	From string `json:"from"`
	FromName string `json:"fromName"`
	To string `json:"to"`
	ToName string `json:"toName"`
	Sid string `json:"sid"`
	ToUrl string `json:"toUrl"`
	SrvLvl string `json:"srvLvl"`
	SrvLvlExt string `json:"srvLvlExt"`
	RecUrl string `json:"recUrl"`
	PnTtl int `json:"pn_ttl"`
	OwnerId string `json:"ownerId"`
}
