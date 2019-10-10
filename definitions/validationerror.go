package definitions

// ValidationError Validation Error
type ValidationError struct {
	ErrorCode string `json:"errorCode"`
	Message string `json:"message"`
	ParameterName string `json:"parameterName"`
}
