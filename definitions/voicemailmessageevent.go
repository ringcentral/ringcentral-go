package definitions

// VoicemailMessageEvent Voicemail Message Event
type VoicemailMessageEvent struct {
	Uuid string `json:"uuid"`
	Event string `json:"event"`
	Timestamp string `json:"timestamp"`
	SubscriptionId string `json:"subscriptionId"`
	OwnerId string `json:"ownerId"`
	Body VoicemailMessageEventBody `json:"body"`
}
