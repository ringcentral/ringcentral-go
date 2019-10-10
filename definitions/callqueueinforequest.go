package definitions

// CallQueueInfoRequest Call Queue Info Request
type CallQueueInfoRequest struct {
	SlaGoal int `json:"slaGoal"`
	SlaThresholdSeconds int `json:"slaThresholdSeconds"`
	IncludeAbandonedCalls bool `json:"includeAbandonedCalls"`
	AbandonedThresholdSeconds int `json:"abandonedThresholdSeconds"`
}
