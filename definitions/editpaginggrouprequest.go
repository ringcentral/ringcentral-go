package definitions

// EditPagingGroupRequest Edit Paging Group Request
type EditPagingGroupRequest struct {
	AddedUserIds []string `json:"addedUserIds"`
	RemovedUserIds []string `json:"removedUserIds"`
	AddedDeviceIds []string `json:"addedDeviceIds"`
	RemovedDeviceIds []string `json:"removedDeviceIds"`
}
