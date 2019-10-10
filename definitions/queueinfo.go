package definitions

// QueueInfo Queue Info
type QueueInfo struct {
	TransferMode string `json:"transferMode"`
	FixedOrderAgents []FixedOrderAgents `json:"fixedOrderAgents"`
	HoldAudioInterruptionMode string `json:"holdAudioInterruptionMode"`
	HoldAudioInterruptionPeriod int `json:"holdAudioInterruptionPeriod"`
	HoldTimeExpirationAction string `json:"holdTimeExpirationAction"`
	AgentTimeout int `json:"agentTimeout"`
	WrapUpTime int `json:"wrapUpTime"`
	HoldTime int `json:"holdTime"`
	MaxCallers int `json:"maxCallers"`
	MaxCallersAction string `json:"maxCallersAction"`
}
