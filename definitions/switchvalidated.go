package definitions

// SwitchValidated Switch Validated
type SwitchValidated struct {
	Id string `json:"id"`
	ChassisId string `json:"chassisId"`
	Status string `json:"status"`
	Errors []ValidationError `json:"errors"`
}
