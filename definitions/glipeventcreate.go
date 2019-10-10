package definitions

// GlipEventCreate Glip Event Create
type GlipEventCreate struct {
	Id string `json:"id"`
	CreatorId string `json:"creatorId"`
	Title string `json:"title"`
	StartTime string `json:"startTime"`
	EndTime string `json:"endTime"`
	AllDay bool `json:"allDay"`
	Recurrence string `json:"recurrence"`
	EndingCondition string `json:"endingCondition"`
	EndingAfter int `json:"endingAfter"`
	EndingOn string `json:"endingOn"`
	Color string `json:"color"`
	Location string `json:"location"`
	Description string `json:"description"`
}
