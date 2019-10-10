package definitions

// BulkAssignAutomaticaLocationUpdatesUsers Bulk Assign Automatica Location Updates Users
type BulkAssignAutomaticaLocationUpdatesUsers struct {
	EnabledUserIds []string `json:"enabledUserIds"`
	DisabledUserIds []string `json:"disabledUserIds"`
}
