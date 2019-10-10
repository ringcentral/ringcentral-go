package definitions

// SuperviseCallSessionRequest Supervise Call Session Request
type SuperviseCallSessionRequest struct {
	Mode string `json:"mode"`
	SupervisorDeviceId string `json:"supervisorDeviceId"`
	AgentExtensionNumber string `json:"agentExtensionNumber"`
	AgentExtensionId string `json:"agentExtensionId"`
}
