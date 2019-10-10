package definitions

// EditGroupRequest Edit Group Request
type EditGroupRequest struct {
	AddedPersonIds []string `json:"addedPersonIds"`
	AddedPersonEmails []string `json:"addedPersonEmails"`
	RemovedPersonIds []string `json:"removedPersonIds"`
}
