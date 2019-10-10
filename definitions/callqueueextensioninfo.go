package definitions

// CallQueueExtensionInfo Call Queue Extension Info
type CallQueueExtensionInfo struct {
	SlaGoal int `json:"slaGoal"`
	SlaThresholdSeconds int `json:"slaThresholdSeconds"`
	IncludeAbandonedCalls bool `json:"includeAbandonedCalls"`
	AbandonedThresholdSeconds int `json:"abandonedThresholdSeconds"`
}
