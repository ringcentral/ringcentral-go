package definitions

// GlipEventsInfo Glip Events Info
type GlipEventsInfo struct {
	Records []GlipEventInfo `json:"records"`
	Navigation GlipNavigationInfo `json:"navigation"`
}
