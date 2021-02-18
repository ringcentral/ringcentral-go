package ringcentral

// AccountPresenceInfo Account Presence Info
type AccountPresenceInfo struct {
	Uri        string                 `json:"uri"`
	Records    []GetPresenceInfo      `json:"records"`
	Navigation PresenceNavigationInfo `json:"navigation"`
	Paging     PresencePagingInfo     `json:"paging"`
}

// ActiveCallInfo Active Call Info
type ActiveCallInfo struct {
	Id                 string           `json:"id"`
	Direction          string           `json:"direction"`
	QueueCall          bool             `json:"queueCall"`
	From               string           `json:"from"`
	FromName           string           `json:"fromName"`
	To                 string           `json:"to"`
	ToName             string           `json:"toName"`
	StartTime          string           `json:"startTime"`
	TelephonyStatus    string           `json:"telephonyStatus"`
	SipData            DetailedCallInfo `json:"sipData"`
	SessionId          string           `json:"sessionId"`
	TelephonySessionId string           `json:"telephonySessionId"`
	OnBehalfOf         string           `json:"onBehalfOf"`
	PartyId            string           `json:"partyId"`
	TerminationType    string           `json:"terminationType"`
	CallInfo           CallInfoCQ       `json:"callInfo"`
}

// CallInfoCQ Call Info CQ
type CallInfoCQ struct {
	Primary    PrimaryCQInfo    `json:"primary"`
	Additional AdditionalCQInfo `json:"additional"`
}

// PrimaryCQInfo Primary CQInfo
type PrimaryCQInfo struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// AdditionalCQInfo Additional CQInfo
type AdditionalCQInfo struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// DetailedCallInfo Detailed Call Info
type DetailedCallInfo struct {
	CallId      string `json:"callId"`
	ToTag       string `json:"toTag"`
	FromTag     string `json:"fromTag"`
	RemoteUri   string `json:"remoteUri"`
	LocalUri    string `json:"localUri"`
	RcSessionId string `json:"rcSessionId"`
}

// GetPresenceExtensionInfo Get Presence Extension Info
type GetPresenceExtensionInfo struct {
	Id              int    `json:"id"`
	Uri             string `json:"uri"`
	ExtensionNumber string `json:"extensionNumber"`
}

// GetPresenceInfo Get Presence Info
type GetPresenceInfo struct {
	Uri                 string                   `json:"uri"`
	AllowSeeMyPresence  bool                     `json:"allowSeeMyPresence"`
	DndStatus           string                   `json:"dndStatus"`
	Extension           GetPresenceExtensionInfo `json:"extension"`
	Message             string                   `json:"message"`
	PickUpCallsOnHold   bool                     `json:"pickUpCallsOnHold"`
	PresenceStatus      string                   `json:"presenceStatus"`
	RingOnMonitoredCall bool                     `json:"ringOnMonitoredCall"`
	TelephonyStatus     string                   `json:"telephonyStatus"`
	UserStatus          string                   `json:"userStatus"`
	MeetingStatus       string                   `json:"meetingStatus"`
	ActiveCalls         []ActiveCallInfo         `json:"activeCalls"`
}

// PresenceNavigationInfo Presence Navigation Info
type PresenceNavigationInfo struct {
	FirstPage    PresenceNavigationInfoURI `json:"firstPage"`
	NextPage     PresenceNavigationInfoURI `json:"nextPage"`
	PreviousPage PresenceNavigationInfoURI `json:"previousPage"`
	LastPage     PresenceNavigationInfoURI `json:"lastPage"`
}

// PresenceNavigationInfoURI Presence Navigation Info URI
type PresenceNavigationInfoURI struct {
	Uri string `json:"uri"`
}

// PresencePagingInfo Presence Paging Info
type PresencePagingInfo struct {
	Page          int `json:"page"`
	PerPage       int `json:"perPage"`
	PageStart     int `json:"pageStart"`
	PageEnd       int `json:"pageEnd"`
	TotalPages    int `json:"totalPages"`
	TotalElements int `json:"totalElements"`
}

// PresenceInfoResource Presence Info Resource
type PresenceInfoResource struct {
	UserStatus          string           `json:"userStatus"`
	DndStatus           string           `json:"dndStatus"`
	Message             string           `json:"message"`
	AllowSeeMyPresence  bool             `json:"allowSeeMyPresence"`
	RingOnMonitoredCall bool             `json:"ringOnMonitoredCall"`
	PickUpCallsOnHold   bool             `json:"pickUpCallsOnHold"`
	ActiveCalls         []ActiveCallInfo `json:"activeCalls"`
}

// PresenceInfoResponse Presence Info Response
type PresenceInfoResponse struct {
	Uri                 string                   `json:"uri"`
	UserStatus          string                   `json:"userStatus"`
	DndStatus           string                   `json:"dndStatus"`
	Message             string                   `json:"message"`
	AllowSeeMyPresence  bool                     `json:"allowSeeMyPresence"`
	RingOnMonitoredCall bool                     `json:"ringOnMonitoredCall"`
	PickUpCallsOnHold   bool                     `json:"pickUpCallsOnHold"`
	ActiveCalls         []ActiveCallInfo         `json:"activeCalls"`
	Extension           GetPresenceExtensionInfo `json:"extension"`
	MeetingStatus       string                   `json:"meetingStatus"`
	TelephonyStatus     string                   `json:"telephonyStatus"`
	PresenceStatus      string                   `json:"presenceStatus"`
}

// AccountPresenceEvent Account Presence Event
type AccountPresenceEvent struct {
	Uuid           string                   `json:"uuid"`
	Event          string                   `json:"@event"`
	Timestamp      string                   `json:"timestamp"`
	SubscriptionId string                   `json:"subscriptionId"`
	Body           AccountPresenceEventBody `json:"body"`
}

// AccountPresenceEventBody Account Presence Event Body
type AccountPresenceEventBody struct {
	ExtensionId         string `json:"extensionId"`
	TelephonyStatus     string `json:"telephonyStatus"`
	Sequence            int    `json:"sequence"`
	PresenceStatus      string `json:"presenceStatus"`
	UserStatus          string `json:"userStatus"`
	DndStatus           string `json:"dndStatus"`
	AllowSeeMyPresence  bool   `json:"allowSeeMyPresence"`
	RingOnMonitoredCall bool   `json:"ringOnMonitoredCall"`
	PickUpCallsOnHold   bool   `json:"pickUpCallsOnHold"`
	OwnerId             string `json:"ownerId"`
}

// ExtensionPresenceEvent Extension Presence Event
type ExtensionPresenceEvent struct {
	Uuid           string                     `json:"uuid"`
	Event          string                     `json:"@event"`
	Timestamp      string                     `json:"timestamp"`
	SubscriptionId string                     `json:"subscriptionId"`
	Body           ExtensionPresenceEventBody `json:"body"`
}

// ExtensionPresenceEventBody Extension Presence Event Body
type ExtensionPresenceEventBody struct {
	ExtensionId         string `json:"extensionId"`
	TelephonyStatus     string `json:"telephonyStatus"`
	Sequence            int    `json:"sequence"`
	PresenceStatus      string `json:"presenceStatus"`
	UserStatus          string `json:"userStatus"`
	DndStatus           string `json:"dndStatus"`
	AllowSeeMyPresence  bool   `json:"allowSeeMyPresence"`
	RingOnMonitoredCall bool   `json:"ringOnMonitoredCall"`
	PickUpCallsOnHold   bool   `json:"pickUpCallsOnHold"`
	OwnerId             string `json:"ownerId"`
}

// DetailedExtensionPresenceEvent Detailed Extension Presence Event
type DetailedExtensionPresenceEvent struct {
	Uuid           string                             `json:"uuid"`
	Event          string                             `json:"@event"`
	Timestamp      string                             `json:"timestamp"`
	SubscriptionId string                             `json:"subscriptionId"`
	Body           DetailedExtensionPresenceEventBody `json:"body"`
}

// DetailedExtensionPresenceEventBody Detailed Extension Presence Event Body
type DetailedExtensionPresenceEventBody struct {
	ExtensionId         string                     `json:"extensionId"`
	TelephonyStatus     string                     `json:"telephonyStatus"`
	ActiveCalls         []ActiveCallInfoWithoutSIP `json:"activeCalls"`
	Sequence            int                        `json:"sequence"`
	PresenceStatus      string                     `json:"presenceStatus"`
	UserStatus          string                     `json:"userStatus"`
	MeetingStatus       string                     `json:"meetingStatus"`
	DndStatus           string                     `json:"dndStatus"`
	AllowSeeMyPresence  bool                       `json:"allowSeeMyPresence"`
	RingOnMonitoredCall bool                       `json:"ringOnMonitoredCall"`
	PickUpCallsOnHold   bool                       `json:"pickUpCallsOnHold"`
	TotalActiveCalls    int                        `json:"totalActiveCalls"`
	OwnerId             string                     `json:"ownerId"`
}

// DetailedExtensionPresenceWithSIPEvent Detailed Extension Presence With SIPEvent
type DetailedExtensionPresenceWithSIPEvent struct {
	Uuid           string                                    `json:"uuid"`
	Event          string                                    `json:"@event"`
	Timestamp      string                                    `json:"timestamp"`
	SubscriptionId string                                    `json:"subscriptionId"`
	Body           DetailedExtensionPresenceWithSIPEventBody `json:"body"`
}

// DetailedExtensionPresenceWithSIPEventBody Detailed Extension Presence With SIPEvent Body
type DetailedExtensionPresenceWithSIPEventBody struct {
	ExtensionId         string           `json:"extensionId"`
	TelephonyStatus     string           `json:"telephonyStatus"`
	ActiveCalls         []ActiveCallInfo `json:"activeCalls"`
	Sequence            int              `json:"sequence"`
	PresenceStatus      string           `json:"presenceStatus"`
	UserStatus          string           `json:"userStatus"`
	MeetingStatus       string           `json:"meetingStatus"`
	DndStatus           string           `json:"dndStatus"`
	AllowSeeMyPresence  bool             `json:"allowSeeMyPresence"`
	RingOnMonitoredCall bool             `json:"ringOnMonitoredCall"`
	PickUpCallsOnHold   bool             `json:"pickUpCallsOnHold"`
	TotalActiveCalls    int              `json:"totalActiveCalls"`
	OwnerId             string           `json:"ownerId"`
}

// ActiveCallInfoWithoutSIP Active Call Info Without SIP
type ActiveCallInfoWithoutSIP struct {
	Id                 string     `json:"id"`
	Direction          string     `json:"direction"`
	QueueCall          bool       `json:"queueCall"`
	From               string     `json:"from"`
	FromName           string     `json:"fromName"`
	To                 string     `json:"to"`
	ToName             string     `json:"toName"`
	PartyId            string     `json:"partyId"`
	StartTime          string     `json:"startTime"`
	SessionId          string     `json:"sessionId"`
	TelephonyStatus    string     `json:"telephonyStatus"`
	TelephonySessionId string     `json:"telephonySessionId"`
	TerminationType    string     `json:"terminationType"`
	CallInfo           CallInfoCQ `json:"callInfo"`
}

// IncomingCallEvent Incoming Call Event
type IncomingCallEvent struct {
	Aps            APSInfo `json:"aps"`
	Event          string  `json:"@event"`
	Uuid           string  `json:"uuid"`
	SubscriptionId string  `json:"subscriptionId"`
	Timestamp      string  `json:"timestamp"`
	ExtensionId    string  `json:"extensionId"`
	Action         string  `json:"action"`
	SessionId      string  `json:"sessionId"`
	ServerId       string  `json:"serverId"`
	From           string  `json:"from"`
	FromName       string  `json:"fromName"`
	To             string  `json:"to"`
	ToName         string  `json:"toName"`
	Sid            string  `json:"sid"`
	ToUrl          string  `json:"toUrl"`
	SrvLvl         string  `json:"srvLvl"`
	SrvLvlExt      string  `json:"srvLvlExt"`
	RecUrl         string  `json:"recUrl"`
	PnTtl          int     `json:"pn_ttl"`
	OwnerId        string  `json:"ownerId"`
}

// MissedCallEvent Missed Call Event
type MissedCallEvent struct {
	Uuid           string   `json:"uuid"`
	PnApns         APNSInfo `json:"pn_apns"`
	Event          string   `json:"@event"`
	SubscriptionId string   `json:"subscriptionId"`
	Timestamp      string   `json:"timestamp"`
	ExtensionId    string   `json:"extensionId"`
	Action         string   `json:"action"`
	SessionId      string   `json:"sessionId"`
	ServerId       string   `json:"serverId"`
	From           string   `json:"from"`
	FromName       string   `json:"fromName"`
	To             string   `json:"to"`
	ToName         string   `json:"toName"`
	Sid            string   `json:"sid"`
	ToUrl          string   `json:"toUrl"`
	SrvLvl         string   `json:"srvLvl"`
	SrvLvlExt      string   `json:"srvLvlExt"`
	RecUrl         string   `json:"recUrl"`
	PnTtl          int      `json:"pn_ttl"`
	OwnerId        string   `json:"ownerId"`
}

// APSInfo APSInfo
type APSInfo struct {
	ContentAvailable int `json:"content-available"`
}

// APNSInfo APNSInfo
type APNSInfo struct {
	Aps APSInfo `json:"aps"`
}

// CallQueuePresence Call Queue Presence
type CallQueuePresence struct {
	Records []CallQueueMemberPresence `json:"records"`
}

// CallQueueUpdatePresence Call Queue Update Presence
type CallQueueUpdatePresence struct {
	Records []CallQueueUpdateMemberPresence `json:"records"`
}

// ExtensionCallQueuePresenceList Extension Call Queue Presence List
type ExtensionCallQueuePresenceList struct {
	Records []ExtensionCallQueuePresence `json:"records"`
}

// ExtensionCallQueueUpdatePresenceList Extension Call Queue Update Presence List
type ExtensionCallQueueUpdatePresenceList struct {
	Records []ExtensionCallQueueUpdatePresence `json:"records"`
}

// CallQueueMemberPresence Call Queue Member Presence
type CallQueueMemberPresence struct {
	Member                  CallQueueMember `json:"member"`
	AcceptQueueCalls        bool            `json:"acceptQueueCalls"`
	AcceptCurrentQueueCalls bool            `json:"acceptCurrentQueueCalls"`
}

// CallQueueUpdateMemberPresence Call Queue Update Member Presence
type CallQueueUpdateMemberPresence struct {
	Member                  CallQueueMemberId `json:"member"`
	AcceptCurrentQueueCalls bool              `json:"acceptCurrentQueueCalls"`
}

// ExtensionCallQueuePresence Extension Call Queue Presence
type ExtensionCallQueuePresence struct {
	CallQueue   CallQueueInfo `json:"callQueue"`
	AcceptCalls bool          `json:"acceptCalls"`
}

// ExtensionCallQueueUpdatePresence Extension Call Queue Update Presence
type ExtensionCallQueueUpdatePresence struct {
	CallQueue   CallQueueId `json:"callQueue"`
	AcceptCalls bool        `json:"acceptCalls"`
}

// CallQueueMemberId Call Queue Member Id
type CallQueueMemberId struct {
	Id string `json:"id"`
}

// CallQueueId Call Queue Id
type CallQueueId struct {
	Id string `json:"id"`
}

// CallQueueMember Call Queue Member
type CallQueueMember struct {
	Id              string       `json:"id"`
	Name            string       `json:"name"`
	ExtensionNumber string       `json:"extensionNumber"`
	Site            SiteResource `json:"site"`
}

// CallQueueInfo Call Queue Info
type CallQueueInfo struct {
	Uri             string `json:"uri"`
	Id              string `json:"id"`
	ExtensionNumber string `json:"extensionNumber"`
	Name            string `json:"name"`
}

// SiteResource Site Resource
type SiteResource struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// CallQueuePresenceEvent Call Queue Presence Event
type CallQueuePresenceEvent struct {
	Uuid           string                     `json:"uuid"`
	Event          string                     `json:"@event"`
	Timestamp      string                     `json:"timestamp"`
	SubscriptionId string                     `json:"subscriptionId"`
	Body           CallQueuePresenceEventBody `json:"body"`
}

// CallQueuePresenceEventBody Call Queue Presence Event Body
type CallQueuePresenceEventBody struct {
	ExtensionId string `json:"extensionId"`
	CallQueueId string `json:"callQueueId"`
	AcceptCalls bool   `json:"acceptCalls"`
}

// RecordsCollectionResourceSubscriptionResponse Records Collection Resource Subscription Response
type RecordsCollectionResourceSubscriptionResponse struct {
	Uri     string             `json:"uri"`
	Records []SubscriptionInfo `json:"records"`
}

// CreateSubscriptionRequest Create Subscription Request
type CreateSubscriptionRequest struct {
	EventFilters []string                        `json:"eventFilters"`
	DeliveryMode NotificationDeliveryModeRequest `json:"deliveryMode"`
	ExpiresIn    int                             `json:"expiresIn"`
}

// NotificationDeliveryModeRequest Notification Delivery Mode Request
type NotificationDeliveryModeRequest struct {
	TransportType     string `json:"transportType"`
	Address           string `json:"address"`
	Encryption        bool   `json:"encryption"`
	CertificateName   string `json:"certificateName"`
	RegistrationId    string `json:"registrationId"`
	VerificationToken string `json:"verificationToken"`
}

// SubscriptionInfo Subscription Info
type SubscriptionInfo struct {
	Id              string                      `json:"id"`
	Uri             string                      `json:"uri"`
	EventFilters    []string                    `json:"eventFilters"`
	DisabledFilters []DisabledFilterInfo        `json:"disabledFilters"`
	ExpirationTime  string                      `json:"expirationTime"`
	ExpiresIn       int                         `json:"expiresIn"`
	Status          string                      `json:"status"`
	CreationTime    string                      `json:"creationTime"`
	DeliveryMode    NotificationDeliveryMode    `json:"deliveryMode"`
	BlacklistedData NotificationBlacklistedData `json:"blacklistedData"`
	TransportType   string                      `json:"transportType"`
	CertificateName string                      `json:"certificateName"`
	RegistrationId  string                      `json:"registrationId"`
}

// DisabledFilterInfo Disabled Filter Info
type DisabledFilterInfo struct {
	Filter  string `json:"filter"`
	Reason  string `json:"reason"`
	Message string `json:"message"`
}

// NotificationDeliveryMode Notification Delivery Mode
type NotificationDeliveryMode struct {
	Encryption          bool   `json:"encryption"`
	Address             string `json:"address"`
	SubscriberKey       string `json:"subscriberKey"`
	SecretKey           string `json:"secretKey"`
	EncryptionAlgorithm string `json:"encryptionAlgorithm"`
	EncryptionKey       string `json:"encryptionKey"`
	TransportType       string `json:"transportType"`
	CertificateName     string `json:"certificateName"`
	RegistrationId      string `json:"registrationId"`
}

// NotificationBlacklistedData Notification Blacklisted Data
type NotificationBlacklistedData struct {
	BlacklistedAt string `json:"blacklistedAt"`
	Reason        string `json:"reason"`
}

// ModifySubscriptionRequest Modify Subscription Request
type ModifySubscriptionRequest struct {
	EventFilters []string                        `json:"eventFilters"`
	DeliveryMode NotificationDeliveryModeRequest `json:"deliveryMode"`
	ExpiresIn    int                             `json:"expiresIn"`
}

// CreateDataExportTaskRequest Create Data Export Task Request
type CreateDataExportTaskRequest struct {
	TimeFrom string                      `json:"timeFrom"`
	TimeTo   string                      `json:"timeTo"`
	Contacts []DataExportTaskContactInfo `json:"contacts"`
	ChatIds  []string                    `json:"chatIds"`
}

// DataExportTaskContactInfo Data Export Task Contact Info
type DataExportTaskContactInfo struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

// DataExportTaskList Data Export Task List
type DataExportTaskList struct {
	Tasks      []DataExportTask             `json:"tasks"`
	Navigation GlipDataExportNavigationInfo `json:"navigation"`
	Paging     GlipDataExportPagingInfo     `json:"paging"`
}

// DataExportTask Data Export Task
type DataExportTask struct {
	Uri              string                 `json:"uri"`
	Id               string                 `json:"id"`
	CreationTime     string                 `json:"creationTime"`
	LastModifiedTime string                 `json:"lastModifiedTime"`
	Status           string                 `json:"status"`
	Creator          CreatorInfo            `json:"creator"`
	Specific         SpecificInfo           `json:"specific"`
	Datasets         []ExportTaskResultInfo `json:"datasets"`
}

// CreatorInfo Creator Info
type CreatorInfo struct {
	Id string `json:"id"`
}

// SpecificInfo Specific Info
type SpecificInfo struct {
	TimeFrom string                      `json:"timeFrom"`
	TimeTo   string                      `json:"timeTo"`
	Contacts []DataExportTaskContactInfo `json:"contacts"`
	ChatIds  []string                    `json:"chatIds"`
}

// ExportTaskResultInfo Export Task Result Info
type ExportTaskResultInfo struct {
	Id   string `json:"id"`
	Uri  string `json:"uri"`
	Size int    `json:"size"`
}

// GlipDataExportNavigationInfo Glip Data Export Navigation Info
type GlipDataExportNavigationInfo struct {
	FirstPage    GlipDataExportNavigationInfoUri `json:"firstPage"`
	NextPage     GlipDataExportNavigationInfoUri `json:"nextPage"`
	PreviousPage GlipDataExportNavigationInfoUri `json:"previousPage"`
	LastPage     GlipDataExportNavigationInfoUri `json:"lastPage"`
}

// GlipDataExportPagingInfo Glip Data Export Paging Info
type GlipDataExportPagingInfo struct {
	Page          int `json:"page"`
	PerPage       int `json:"perPage"`
	PageStart     int `json:"pageStart"`
	PageEnd       int `json:"pageEnd"`
	TotalPages    int `json:"totalPages"`
	TotalElements int `json:"totalElements"`
}

// GlipDataExportNavigationInfoUri Glip Data Export Navigation Info Uri
type GlipDataExportNavigationInfoUri struct {
	Uri string `json:"uri"`
}

// GlipNotesInfo Glip Notes Info
type GlipNotesInfo struct {
	Records    []GlipNoteInfo     `json:"records"`
	Navigation GlipNavigationInfo `json:"navigation"`
}

// GlipNoteInfo Glip Note Info
type GlipNoteInfo struct {
	Id               string             `json:"id"`
	Title            string             `json:"title"`
	ChatIds          []string           `json:"chatIds"`
	Preview          string             `json:"preview"`
	Creator          CreatorInfo        `json:"creator"`
	LastModifiedBy   LastModifiedByInfo `json:"lastModifiedBy"`
	LockedBy         LockedByInfo       `json:"lockedBy"`
	Status           string             `json:"status"`
	CreationTime     string             `json:"creationTime"`
	LastModifiedTime string             `json:"lastModifiedTime"`
	Type             string             `json:"type"`
}

// GetGlipNoteInfo Get Glip Note Info
type GetGlipNoteInfo struct {
	Id               string             `json:"id"`
	Title            string             `json:"title"`
	ChatIds          []string           `json:"chatIds"`
	Preview          string             `json:"preview"`
	Body             string             `json:"body"`
	Creator          CreatorInfo        `json:"creator"`
	LastModifiedBy   LastModifiedByInfo `json:"lastModifiedBy"`
	LockedBy         LockedByInfo       `json:"lockedBy"`
	Status           string             `json:"status"`
	CreationTime     string             `json:"creationTime"`
	LastModifiedTime string             `json:"lastModifiedTime"`
	Type             string             `json:"type"`
}

// LockedByInfo Locked By Info
type LockedByInfo struct {
	Id string `json:"id"`
}

// LastModifiedByInfo Last Modified By Info
type LastModifiedByInfo struct {
	Id string `json:"id"`
}

// GlipNoteCreate Glip Note Create
type GlipNoteCreate struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// GlipEventsInfo Glip Events Info
type GlipEventsInfo struct {
	Records    []GlipEventInfo    `json:"records"`
	Navigation GlipNavigationInfo `json:"navigation"`
}

// GlipEventInfo Glip Event Info
type GlipEventInfo struct {
	Id              string `json:"id"`
	CreatorId       string `json:"creatorId"`
	Title           string `json:"title"`
	StartTime       string `json:"startTime"`
	EndTime         string `json:"endTime"`
	AllDay          bool   `json:"allDay"`
	Recurrence      string `json:"recurrence"`
	EndingCondition string `json:"endingCondition"`
	EndingAfter     int    `json:"endingAfter"`
	EndingOn        string `json:"endingOn"`
	Color           string `json:"color"`
	Location        string `json:"location"`
	Description     string `json:"description"`
}

// GlipEventCreate Glip Event Create
type GlipEventCreate struct {
	Id              string `json:"id"`
	CreatorId       string `json:"creatorId"`
	Title           string `json:"title"`
	StartTime       string `json:"startTime"`
	EndTime         string `json:"endTime"`
	AllDay          bool   `json:"allDay"`
	Recurrence      string `json:"recurrence"`
	EndingCondition string `json:"endingCondition"`
	EndingAfter     int    `json:"endingAfter"`
	EndingOn        string `json:"endingOn"`
	Color           string `json:"color"`
	Location        string `json:"location"`
	Description     string `json:"description"`
}

// GlipPersonInfo Glip Person Info
type GlipPersonInfo struct {
	Id               string `json:"id"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	Email            string `json:"email"`
	Avatar           string `json:"avatar"`
	CompanyId        string `json:"companyId"`
	CreationTime     string `json:"creationTime"`
	LastModifiedTime string `json:"lastModifiedTime"`
}

// GlipCompany Glip Company
type GlipCompany struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	Domain           string `json:"domain"`
	CreationTime     string `json:"creationTime"`
	LastModifiedTime string `json:"lastModifiedTime"`
}

// GlipMessageAttachmentInfo Glip Message Attachment Info
type GlipMessageAttachmentInfo struct {
	Id              string                            `json:"id"`
	Type            string                            `json:"type"`
	Fallback        string                            `json:"fallback"`
	Intro           string                            `json:"intro"`
	Author          GlipMessageAttachmentAuthorInfo   `json:"author"`
	Title           string                            `json:"title"`
	Text            string                            `json:"text"`
	ImageUri        string                            `json:"imageUri"`
	ThumbnailUri    string                            `json:"thumbnailUri"`
	Fields          []GlipMessageAttachmentFieldsInfo `json:"fields"`
	Footnote        GlipMessageAttachmentFootnoteInfo `json:"footnote"`
	CreatorId       string                            `json:"creatorId"`
	StartTime       string                            `json:"startTime"`
	EndTime         string                            `json:"endTime"`
	AllDay          bool                              `json:"allDay"`
	Recurrence      string                            `json:"recurrence"`
	EndingCondition string                            `json:"endingCondition"`
	EndingAfter     int                               `json:"endingAfter"`
	EndingOn        string                            `json:"endingOn"`
	Color           string                            `json:"color"`
	Location        string                            `json:"location"`
	Description     string                            `json:"description"`
}

// GlipMessageAttachmentAuthorInfo Glip Message Attachment Author Info
type GlipMessageAttachmentAuthorInfo struct {
	Name    string `json:"name"`
	Uri     string `json:"uri"`
	IconUri string `json:"iconUri"`
}

// GlipMessageAttachmentFootnoteInfo Glip Message Attachment Footnote Info
type GlipMessageAttachmentFootnoteInfo struct {
	Text    string `json:"text"`
	IconUri string `json:"iconUri"`
	Time    string `json:"time"`
}

// GlipMessageAttachmentFieldsInfo Glip Message Attachment Fields Info
type GlipMessageAttachmentFieldsInfo struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Style string `json:"style"`
}

// GlipAttachmentInfoRequest Glip Attachment Info Request
type GlipAttachmentInfoRequest struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

// GlipPostInfo Glip Post Info
type GlipPostInfo struct {
	Id               string                      `json:"id"`
	GroupId          string                      `json:"groupId"`
	Type             string                      `json:"type"`
	Text             string                      `json:"text"`
	CreatorId        string                      `json:"creatorId"`
	AddedPersonIds   []string                    `json:"addedPersonIds"`
	CreationTime     string                      `json:"creationTime"`
	LastModifiedTime string                      `json:"lastModifiedTime"`
	Attachments      []GlipMessageAttachmentInfo `json:"attachments"`
	Mentions         []GlipMentionsInfo          `json:"mentions"`
	Activity         string                      `json:"activity"`
	Title            string                      `json:"title"`
	IconUri          string                      `json:"iconUri"`
	IconEmoji        string                      `json:"iconEmoji"`
}

// GlipNavigationInfo Glip Navigation Info
type GlipNavigationInfo struct {
	PrevPageToken string `json:"prevPageToken"`
	NextPageToken string `json:"nextPageToken"`
}

// GlipMentionsInfo Glip Mentions Info
type GlipMentionsInfo struct {
	Id   string `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}

// GlipWebhookList Glip Webhook List
type GlipWebhookList struct {
	Records []GlipWebhookInfo `json:"records"`
}

// GlipWebhookInfo Glip Webhook Info
type GlipWebhookInfo struct {
	Id               string   `json:"id"`
	CreatorId        string   `json:"creatorId"`
	GroupIds         []string `json:"groupIds"`
	CreationTime     string   `json:"creationTime"`
	LastModifiedTime string   `json:"lastModifiedTime"`
	Uri              string   `json:"uri"`
	Status           string   `json:"status"`
}

// GlipGroupsEvent Glip Groups Event
type GlipGroupsEvent struct {
	Id               string           `json:"id"`
	Type             string           `json:"type"`
	IsPublic         bool             `json:"isPublic"`
	Name             string           `json:"name"`
	Description      string           `json:"description"`
	Members          []GlipMemberInfo `json:"members"`
	CreationTime     string           `json:"creationTime"`
	LastModifiedTime string           `json:"lastModifiedTime"`
	EventType        string           `json:"eventType"`
}

// GlipPostEvent Glip Post Event
type GlipPostEvent struct {
	Id               string             `json:"id"`
	EventType        string             `json:"eventType"`
	GroupId          string             `json:"groupId"`
	Type             string             `json:"type"`
	Text             string             `json:"text"`
	CreatorId        string             `json:"creatorId"`
	AddedPersonIds   []string           `json:"addedPersonIds"`
	RemovedPersonIds []string           `json:"removedPersonIds"`
	Mentions         []GlipMentionsInfo `json:"mentions"`
	CreationTime     string             `json:"creationTime"`
	LastModifiedTime string             `json:"lastModifiedTime"`
}

// GlipUnreadMessageCountEvent Glip Unread Message Count Event
type GlipUnreadMessageCountEvent struct {
	Uuid    string       `json:"uuid"`
	PnApns  GlipAPNSInfo `json:"pn_apns"`
	PnGcm   GCMInfo      `json:"pn_gcm"`
	OwnerId string       `json:"ownerId"`
}

// GlipAPNSInfo Glip APNSInfo
type GlipAPNSInfo struct {
	Aps            APSInfo                    `json:"aps"`
	Timestamp      string                     `json:"timestamp"`
	Uuid           string                     `json:"uuid"`
	Event          string                     `json:"@event"`
	SubscriptionId string                     `json:"subscriptionId"`
	Body           GlipUnreadMessageCountInfo `json:"body"`
}

// GlipUnreadMessageCountInfo Glip Unread Message Count Info
type GlipUnreadMessageCountInfo struct {
	Unread int `json:"unread"`
}

// GCMInfo GCMInfo
type GCMInfo struct {
	Priority   string  `json:"priority"`
	TimeToLive int     `json:"time_to_live"`
	Data       GCMData `json:"data"`
}

// GCMData GCMData
type GCMData struct {
	Event          string `json:"@event"`
	SubscriptionId string `json:"subscriptionId"`
	Timestamp      string `json:"timestamp"`
	ExtensionId    string `json:"extensionId"`
	Action         string `json:"action"`
	SessionId      string `json:"sessionId"`
	ServerId       string `json:"serverId"`
	From           string `json:"_from"`
	FromName       string `json:"fromName"`
	To             string `json:"to"`
	ToName         string `json:"toName"`
	Sid            string `json:"sid"`
	ToUrl          string `json:"toUrl"`
	SrvLvl         string `json:"srvLvl"`
	SrvLvlExt      string `json:"srvLvlExt"`
}

// GlipChatsListWithoutNavigation Glip Chats List Without Navigation
type GlipChatsListWithoutNavigation struct {
	Records []GlipChatInfo `json:"records"`
}

// GlipChatsList Glip Chats List
type GlipChatsList struct {
	Records    []GlipChatInfo     `json:"records"`
	Navigation GlipNavigationInfo `json:"navigation"`
}

// GlipTeamsList Glip Teams List
type GlipTeamsList struct {
	Records    []GlipTeamInfo     `json:"records"`
	Navigation GlipNavigationInfo `json:"navigation"`
}

// GlipConversationsList Glip Conversations List
type GlipConversationsList struct {
	Records    []GlipConversationInfo `json:"records"`
	Navigation GlipNavigationInfo     `json:"navigation"`
}

// GlipChatInfo Glip Chat Info
type GlipChatInfo struct {
	Id               string           `json:"id"`
	Type             string           `json:"type"`
	Public           bool             `json:"@public"`
	Name             string           `json:"name"`
	Description      string           `json:"description"`
	Status           string           `json:"status"`
	CreationTime     string           `json:"creationTime"`
	LastModifiedTime string           `json:"lastModifiedTime"`
	Members          []GlipMemberInfo `json:"members"`
}

// GlipMemberInfo Glip Member Info
type GlipMemberInfo struct {
	Id string `json:"id"`
}

// GlipTeamInfo Glip Team Info
type GlipTeamInfo struct {
	Id               string `json:"id"`
	Type             string `json:"type"`
	Public           bool   `json:"@public"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	Status           string `json:"status"`
	CreationTime     string `json:"creationTime"`
	LastModifiedTime string `json:"lastModifiedTime"`
}

// GlipConversationInfo Glip Conversation Info
type GlipConversationInfo struct {
	Id               string             `json:"id"`
	Type             string             `json:"type"`
	CreationTime     string             `json:"creationTime"`
	LastModifiedTime string             `json:"lastModifiedTime"`
	Members          []CreateGlipMember `json:"members"`
}

// GlipEveryoneInfo Glip Everyone Info
type GlipEveryoneInfo struct {
	Id               string `json:"id"`
	Type             string `json:"type"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	CreationTime     string `json:"creationTime"`
	LastModifiedTime string `json:"lastModifiedTime"`
}

// GlipPostMembersListBody Glip Post Members List Body
type GlipPostMembersListBody struct {
	Members []CreateGlipMember `json:"members"`
}

// CreateGlipMember Create Glip Member
type CreateGlipMember struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

// CreateGlipConversationRequest Create Glip Conversation Request
type CreateGlipConversationRequest struct {
	Members []CreateGlipMember `json:"members"`
}

// GlipPostMembersIdsListBody Glip Post Members Ids List Body
type GlipPostMembersIdsListBody struct {
	Members []GlipMemberInfo `json:"members"`
}

// GlipPostTeamBody Glip Post Team Body
type GlipPostTeamBody struct {
	Public      bool               `json:"@public"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Members     []CreateGlipMember `json:"members"`
}

// GlipPatchTeamBody Glip Patch Team Body
type GlipPatchTeamBody struct {
	Public      bool   `json:"@public"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// GlipPatchPostBody Glip Patch Post Body
type GlipPatchPostBody struct {
	Text string `json:"text"`
}

// GlipPostPostBody Glip Post Post Body
type GlipPostPostBody struct {
	Text        string                      `json:"text"`
	Attachments []GlipAttachmentInfoRequest `json:"attachments"`
}

// GlipPostsList Glip Posts List
type GlipPostsList struct {
	Records    []GlipPostInfo     `json:"records"`
	Navigation GlipNavigationInfo `json:"navigation"`
}

// GlipPreferencesInfo Glip Preferences Info
type GlipPreferencesInfo struct {
	Chats GlipPreferencesChats `json:"chats"`
}

// GlipPreferencesChats Glip Preferences Chats
type GlipPreferencesChats struct {
	MaxCount     int    `json:"maxCount"`
	LeftRailMode string `json:"leftRailMode"`
}

// UpdateGlipEveryoneRequest Update Glip Everyone Request
type UpdateGlipEveryoneRequest struct {
	Name        int    `json:"name"`
	Description string `json:"description"`
}

// GlipTaskList Glip Task List
type GlipTaskList struct {
	Records []GlipTaskInfo `json:"records"`
}

// GlipTaskInfo Glip Task Info
type GlipTaskInfo struct {
	Id                     string                 `json:"id"`
	CreationTime           string                 `json:"creationTime"`
	LastModifiedTime       string                 `json:"lastModifiedTime"`
	Type                   string                 `json:"type"`
	Creator                CreatorInfo            `json:"creator"`
	ChatIds                []string               `json:"chatIds"`
	Status                 string                 `json:"status"`
	Subject                string                 `json:"subject"`
	Assignees              []TaskAssigneeInfo     `json:"assignees"`
	CompletenessCondition  string                 `json:"completenessCondition"`
	CompletenessPercentage int                    `json:"completenessPercentage"`
	StartDate              string                 `json:"startDate"`
	DueDate                string                 `json:"dueDate"`
	Color                  string                 `json:"color"`
	Section                string                 `json:"section"`
	Description            string                 `json:"description"`
	Recurrence             GlipTaskRecurrenceInfo `json:"recurrence"`
	Attachments            []TaskAttachment       `json:"attachments"`
}

// TaskAttachment Task Attachment
type TaskAttachment struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Name       string `json:"name"`
	ContentUri string `json:"contentUri"`
}

// GlipCreateTask Glip Create Task
type GlipCreateTask struct {
	Subject               string                      `json:"subject"`
	Assignees             []AssigneeInfo              `json:"assignees"`
	CompletenessCondition string                      `json:"completenessCondition"`
	StartDate             string                      `json:"startDate"`
	DueDate               string                      `json:"dueDate"`
	Color                 string                      `json:"color"`
	Section               string                      `json:"section"`
	Description           string                      `json:"description"`
	Recurrence            GlipTaskRecurrenceInfo      `json:"recurrence"`
	Attachments           []GlipAttachmentInfoRequest `json:"attachments"`
}

// GlipTaskRecurrenceInfo Glip Task Recurrence Info
type GlipTaskRecurrenceInfo struct {
	Schedule        string `json:"schedule"`
	EndingCondition string `json:"endingCondition"`
	EndingAfter     int    `json:"endingAfter"`
	EndingOn        string `json:"endingOn"`
}

// GlipUpdateTask Glip Update Task
type GlipUpdateTask struct {
	Subject               string                      `json:"subject"`
	Assignees             []AssigneeInfo              `json:"assignees"`
	CompletenessCondition string                      `json:"completenessCondition"`
	StartDate             string                      `json:"startDate"`
	DueDate               string                      `json:"dueDate"`
	Color                 string                      `json:"color"`
	Section               string                      `json:"section"`
	Description           string                      `json:"description"`
	Recurrence            GlipTaskRecurrenceInfo      `json:"recurrence"`
	Attachments           []GlipAttachmentInfoRequest `json:"attachments"`
}

// AssigneeInfo Assignee Info
type AssigneeInfo struct {
	Id string `json:"id"`
}

// GlipCompleteTask Glip Complete Task
type GlipCompleteTask struct {
	Status                 string         `json:"status"`
	Assignees              []AssigneeInfo `json:"assignees"`
	CompletenessPercentage int            `json:"completenessPercentage"`
}

// UnifiedPresenceListEntry Unified Presence List Entry
type UnifiedPresenceListEntry struct {
	ResourceId string          `json:"resourceId"`
	Status     int             `json:"status"`
	Body       UnifiedPresence `json:"body"`
}

// UpdateUnifiedPresence Update Unified Presence
type UpdateUnifiedPresence struct {
	Glip      UpdateUnifiedPresenceGlip      `json:"glip"`
	Telephony UpdateUnifiedPresenceTelephony `json:"telephony"`
}

// UnifiedPresence Unified Presence
type UnifiedPresence struct {
	Status    string                   `json:"status"`
	Glip      UnifiedPresenceGlip      `json:"glip"`
	Telephony UnifiedPresenceTelephony `json:"telephony"`
	Meeting   UnifiedPresenceMeeting   `json:"meeting"`
}

// UpdateUnifiedPresenceGlip Update Unified Presence Glip
type UpdateUnifiedPresenceGlip struct {
	Visibility   string `json:"visibility"`
	Availability string `json:"availability"`
}

// UnifiedPresenceGlip Unified Presence Glip
type UnifiedPresenceGlip struct {
	Status       string `json:"status"`
	Visibility   string `json:"visibility"`
	Availability string `json:"availability"`
}

// UnifiedPresenceTelephony Unified Presence Telephony
type UnifiedPresenceTelephony struct {
	Status       string `json:"status"`
	Visibility   string `json:"visibility"`
	Availability string `json:"availability"`
}

// UpdateUnifiedPresenceTelephony Update Unified Presence Telephony
type UpdateUnifiedPresenceTelephony struct {
	Availability string `json:"availability"`
}

// UnifiedPresenceMeeting Unified Presence Meeting
type UnifiedPresenceMeeting struct {
	Status string `json:"status"`
}

// TaskAssigneeInfo Task Assignee Info
type TaskAssigneeInfo struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

// TokenInfo Token Info
type TokenInfo struct {
	AccessToken           string `json:"access_token"`
	ExpiresIn             int    `json:"expires_in"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiresIn int    `json:"refresh_token_expires_in"`
	Scope                 string `json:"scope"`
	TokenType             string `json:"token_type"`
	OwnerId               string `json:"owner_id"`
	EndpointId            string `json:"endpoint_id"`
	IdToken               string `json:"id_token"`
}

// ContactList Contact List
type ContactList struct {
	Uri        string                     `json:"uri"`
	Records    []PersonalContactResource  `json:"records"`
	Navigation UserContactsNavigationInfo `json:"navigation"`
	Paging     UserContactsPagingInfo     `json:"paging"`
	Groups     UserContactsGroupsInfo     `json:"groups"`
}

// UserContactsGroupsInfo User Contacts Groups Info
type UserContactsGroupsInfo struct {
	Uri string `json:"uri"`
}

// PersonalContactRequest Personal Contact Request
type PersonalContactRequest struct {
	FirstName       string             `json:"firstName"`
	LastName        string             `json:"lastName"`
	MiddleName      string             `json:"middleName"`
	NickName        string             `json:"nickName"`
	Company         string             `json:"company"`
	JobTitle        string             `json:"jobTitle"`
	Email           string             `json:"email"`
	Email2          string             `json:"email2"`
	Email3          string             `json:"email3"`
	Birthday        string             `json:"birthday"`
	WebPage         string             `json:"webPage"`
	Notes           string             `json:"notes"`
	HomePhone       string             `json:"homePhone"`
	HomePhone2      string             `json:"homePhone2"`
	BusinessPhone   string             `json:"businessPhone"`
	BusinessPhone2  string             `json:"businessPhone2"`
	MobilePhone     string             `json:"mobilePhone"`
	BusinessFax     string             `json:"businessFax"`
	CompanyPhone    string             `json:"companyPhone"`
	AssistantPhone  string             `json:"assistantPhone"`
	CarPhone        string             `json:"carPhone"`
	OtherPhone      string             `json:"otherPhone"`
	OtherFax        string             `json:"otherFax"`
	CallbackPhone   string             `json:"callbackPhone"`
	HomeAddress     ContactAddressInfo `json:"homeAddress"`
	BusinessAddress ContactAddressInfo `json:"businessAddress"`
	OtherAddress    ContactAddressInfo `json:"otherAddress"`
	RingtoneIndex   string             `json:"ringtoneIndex"`
}

// PersonalContactResource Personal Contact Resource
type PersonalContactResource struct {
	Uri             string             `json:"uri"`
	Availability    string             `json:"availability"`
	Email           string             `json:"email"`
	Id              int                `json:"id"`
	Notes           string             `json:"notes"`
	Company         string             `json:"company"`
	FirstName       string             `json:"firstName"`
	LastName        string             `json:"lastName"`
	JobTitle        string             `json:"jobTitle"`
	Birthday        string             `json:"birthday"`
	WebPage         string             `json:"webPage"`
	MiddleName      string             `json:"middleName"`
	NickName        string             `json:"nickName"`
	Email2          string             `json:"email2"`
	Email3          string             `json:"email3"`
	HomePhone       string             `json:"homePhone"`
	HomePhone2      string             `json:"homePhone2"`
	BusinessPhone   string             `json:"businessPhone"`
	BusinessPhone2  string             `json:"businessPhone2"`
	MobilePhone     string             `json:"mobilePhone"`
	BusinessFax     string             `json:"businessFax"`
	CompanyPhone    string             `json:"companyPhone"`
	AssistantPhone  string             `json:"assistantPhone"`
	CarPhone        string             `json:"carPhone"`
	OtherPhone      string             `json:"otherPhone"`
	OtherFax        string             `json:"otherFax"`
	CallbackPhone   string             `json:"callbackPhone"`
	BusinessAddress ContactAddressInfo `json:"businessAddress"`
	HomeAddress     ContactAddressInfo `json:"homeAddress"`
	OtherAddress    ContactAddressInfo `json:"otherAddress"`
	RingtoneIndex   string             `json:"ringtoneIndex"`
}

// AddressBookSync Address Book Sync
type AddressBookSync struct {
	Uri         string                    `json:"uri"`
	Records     []PersonalContactResource `json:"records"`
	SyncInfo    SyncInfo                  `json:"syncInfo"`
	NextPageId  int                       `json:"nextPageId"`
	NextPageUri string                    `json:"nextPageUri"`
}

// SyncInfo Sync Info
type SyncInfo struct {
	SyncType          string `json:"syncType"`
	SyncToken         string `json:"syncToken"`
	SyncTime          string `json:"syncTime"`
	OlderRecordsExist bool   `json:"olderRecordsExist"`
}

// UserContactsNavigationInfo User Contacts Navigation Info
type UserContactsNavigationInfo struct {
	FirstPage    UserContactsNavigationInfoUri `json:"firstPage"`
	NextPage     UserContactsNavigationInfoUri `json:"nextPage"`
	PreviousPage UserContactsNavigationInfoUri `json:"previousPage"`
	LastPage     UserContactsNavigationInfoUri `json:"lastPage"`
}

// UserContactsPagingInfo User Contacts Paging Info
type UserContactsPagingInfo struct {
	Page          int `json:"page"`
	PerPage       int `json:"perPage"`
	PageStart     int `json:"pageStart"`
	PageEnd       int `json:"pageEnd"`
	TotalPages    int `json:"totalPages"`
	TotalElements int `json:"totalElements"`
}

// UserContactsNavigationInfoUri User Contacts Navigation Info Uri
type UserContactsNavigationInfoUri struct {
	Uri string `json:"uri"`
}

// ContactAddressInfo Contact Address Info
type ContactAddressInfo struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country"`
	State   string `json:"state"`
	Zip     string `json:"zip"`
}

// FavoriteCollection Favorite Collection
type FavoriteCollection struct {
	Records []FavoriteContactResource `json:"records"`
}

// FavoriteContactList Favorite Contact List
type FavoriteContactList struct {
	Uri     string                    `json:"uri"`
	Records []FavoriteContactResource `json:"records"`
}

// FavoriteContactResource Favorite Contact Resource
type FavoriteContactResource struct {
	Id          int    `json:"id"`
	ExtensionId string `json:"extensionId"`
	AccountId   string `json:"accountId"`
	ContactId   string `json:"contactId"`
}

// ExtensionFavoritesEvent Extension Favorites Event
type ExtensionFavoritesEvent struct {
	Uuid           string                      `json:"uuid"`
	Event          string                      `json:"@event"`
	Timestamp      string                      `json:"timestamp"`
	SubscriptionId string                      `json:"subscriptionId"`
	Body           ExtensionFavoritesEventBody `json:"body"`
}

// ExtensionFavoritesEventBody Extension Favorites Event Body
type ExtensionFavoritesEventBody struct {
	ExtensionId string `json:"extensionId"`
	OwnerId     string `json:"ownerId"`
}

// AccountResource Account Resource
type AccountResource struct {
	CompanyName   string              `json:"companyName"`
	FederatedName string              `json:"federatedName"`
	Id            string              `json:"id"`
	MainNumber    PhoneNumberResource `json:"mainNumber"`
}

// BusinessSiteResource Business Site Resource
type BusinessSiteResource struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

// ContactResource Contact Resource
type ContactResource struct {
	Account         AccountResource                      `json:"account"`
	Department      string                               `json:"department"`
	Email           string                               `json:"email"`
	ExtensionNumber string                               `json:"extensionNumber"`
	FirstName       string                               `json:"firstName"`
	LastName        string                               `json:"lastName"`
	Name            string                               `json:"name"`
	Id              string                               `json:"id"`
	JobTitle        string                               `json:"jobTitle"`
	PhoneNumbers    []PhoneNumberResource                `json:"phoneNumbers"`
	ProfileImage    AccountDirectoryProfileImageResource `json:"profileImage"`
	Site            BusinessSiteResource                 `json:"site"`
	Status          string                               `json:"status"`
	Type            string                               `json:"type"`
}

// DirectoryResource Directory Resource
type DirectoryResource struct {
	Paging  CompanyContactsPagingInfo `json:"paging"`
	Records []ContactResource         `json:"records"`
}

// ErrorResponse Error Response
type ErrorResponse struct {
	ErrorCode   string `json:"errorCode"`
	Description string `json:"description"`
}

// FederatedAccountResource Federated Account Resource
type FederatedAccountResource struct {
	CompanyName      string              `json:"companyName"`
	ConflictCount    int                 `json:"conflictCount"`
	FederatedName    string              `json:"federatedName"`
	Id               string              `json:"id"`
	LinkCreationTime string              `json:"linkCreationTime"`
	MainNumber       PhoneNumberResource `json:"mainNumber"`
}

// FederationResource Federation Resource
type FederationResource struct {
	Accounts         []FederatedAccountResource `json:"accounts"`
	CreationTime     string                     `json:"creationTime"`
	DisplayName      string                     `json:"displayName"`
	Id               string                     `json:"id"`
	LastModifiedTime string                     `json:"lastModifiedTime"`
}

// CompanyContactsPagingInfo Company Contacts Paging Info
type CompanyContactsPagingInfo struct {
	Page          int `json:"page"`
	TotalPages    int `json:"totalPages"`
	PerPage       int `json:"perPage"`
	TotalElements int `json:"totalElements"`
	PageStart     int `json:"pageStart"`
	PageEnd       int `json:"pageEnd"`
}

// PhoneNumberResource Phone Number Resource
type PhoneNumberResource struct {
	Id          string                   `json:"id"`
	Country     CountryResource          `json:"country"`
	Extension   PhoneNumberExtensionInfo `json:"extension"`
	Label       string                   `json:"label"`
	Location    string                   `json:"location"`
	PaymentType string                   `json:"paymentType"`
	PhoneNumber string                   `json:"phoneNumber"`
	Status      string                   `json:"status"`
	UsageType   string                   `json:"usageType"`
	Type        string                   `json:"type"`
}

// AccountDirectoryProfileImageResource Account Directory Profile Image Resource
type AccountDirectoryProfileImageResource struct {
	Etag string `json:"etag"`
	Uri  string `json:"uri"`
}

// ContactDirectoryEvent Contact Directory Event
type ContactDirectoryEvent struct {
	Id              string                           `json:"id"`
	EventType       string                           `json:"eventType"`
	Type            string                           `json:"type"`
	Status          string                           `json:"status"`
	FirstName       string                           `json:"firstName"`
	LastName        string                           `json:"lastName"`
	Department      string                           `json:"department"`
	Email           string                           `json:"email"`
	ExtensionNumber string                           `json:"extensionNumber"`
	Account         CompanyDirectoryAccountInfo      `json:"account"`
	PhoneNumbers    CompanyDirectoryPhoneNumberInfo  `json:"phoneNumbers"`
	Site            ContactDirectorySiteInfo         `json:"site"`
	ProfileImage    CompanyDirectoryProfileImageInfo `json:"profileImage"`
	OwnerId         string                           `json:"ownerId"`
}

// CompanyDirectoryAccountInfo Company Directory Account Info
type CompanyDirectoryAccountInfo struct {
	Id string `json:"id"`
}

// CompanyDirectoryPhoneNumberInfo Company Directory Phone Number Info
type CompanyDirectoryPhoneNumberInfo struct {
	PhoneNumber string `json:"phoneNumber"`
	Type        string `json:"type"`
	Hidden      bool   `json:"hidden"`
	UsageType   string `json:"usageType"`
}

// CompanyDirectoryProfileImageInfo Company Directory Profile Image Info
type CompanyDirectoryProfileImageInfo struct {
	Uri  string `json:"uri"`
	Etag string `json:"etag"`
}

// ContactDirectorySiteInfo Contact Directory Site Info
type ContactDirectorySiteInfo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// SearchDirectoryEntriesRequest Search Directory Entries Request
type SearchDirectoryEntriesRequest struct {
	SearchString  string    `json:"searchString"`
	SearchFields  []string  `json:"searchFields"`
	ShowFederated bool      `json:"showFederated"`
	ExtensionType string    `json:"extensionType"`
	OrderBy       []OrderBy `json:"orderBy"`
	Page          int       `json:"page"`
	PerPage       int       `json:"perPage"`
}

// OrderBy Order By
type OrderBy struct {
	Index     int    `json:"index"`
	FieldName string `json:"fieldName"`
	Direction string `json:"direction"`
}

// CompanyDirectoryEvent Company Directory Event
type CompanyDirectoryEvent struct {
	Uuid           string                    `json:"uuid"`
	Event          string                    `json:"@event"`
	Timestamp      string                    `json:"timestamp"`
	SubscriptionId string                    `json:"subscriptionId"`
	OwnerId        string                    `json:"ownerId"`
	Body           CompanyDirectoryEventBody `json:"body"`
}

// CompanyDirectoryEventBody Company Directory Event Body
type CompanyDirectoryEventBody struct {
	Id              string                            `json:"id"`
	EventType       string                            `json:"eventType"`
	Type            string                            `json:"type"`
	Status          string                            `json:"status"`
	FirstName       string                            `json:"firstName"`
	LastName        string                            `json:"lastName"`
	Name            string                            `json:"name"`
	Department      string                            `json:"department"`
	Email           string                            `json:"email"`
	ExtensionNumber string                            `json:"extensionNumber"`
	Account         CompanyDirectoryAccountInfo       `json:"account"`
	PhoneNumbers    []CompanyDirectoryPhoneNumberInfo `json:"phoneNumbers"`
	Site            ContactDirectorySiteInfo          `json:"site"`
	ProfileImage    CompanyDirectoryProfileImageInfo  `json:"profileImage"`
}

// AnsweringRuleInfo Answering Rule Info
type AnsweringRuleInfo struct {
	Uri                     string                      `json:"uri"`
	Id                      string                      `json:"id"`
	Type                    string                      `json:"type"`
	Name                    string                      `json:"name"`
	Enabled                 bool                        `json:"enabled"`
	Schedule                ScheduleInfo                `json:"schedule"`
	CalledNumbers           []CalledNumberInfo          `json:"calledNumbers"`
	Callers                 []CallersInfo               `json:"callers"`
	CallHandlingAction      string                      `json:"callHandlingAction"`
	Forwarding              ForwardingInfo              `json:"forwarding"`
	UnconditionalForwarding UnconditionalForwardingInfo `json:"unconditionalForwarding"`
	Queue                   QueueInfo                   `json:"queue"`
	Transfer                TransferredExtensionInfo    `json:"transfer"`
	Voicemail               VoicemailInfo               `json:"voicemail"`
	Greetings               []GreetingInfo              `json:"greetings"`
	Screening               string                      `json:"screening"`
	SharedLines             SharedLinesInfo             `json:"sharedLines"`
}

// CustomAnsweringRuleInfo Custom Answering Rule Info
type CustomAnsweringRuleInfo struct {
	Uri                     string                      `json:"uri"`
	Id                      string                      `json:"id"`
	Type                    string                      `json:"type"`
	Name                    string                      `json:"name"`
	Enabled                 bool                        `json:"enabled"`
	Schedule                ScheduleInfo                `json:"schedule"`
	CalledNumbers           []CalledNumberInfo          `json:"calledNumbers"`
	Callers                 []CallersInfo               `json:"callers"`
	CallHandlingAction      string                      `json:"callHandlingAction"`
	Forwarding              ForwardingInfo              `json:"forwarding"`
	UnconditionalForwarding UnconditionalForwardingInfo `json:"unconditionalForwarding"`
	Queue                   QueueInfo                   `json:"queue"`
	Transfer                TransferredExtensionInfo    `json:"transfer"`
	Voicemail               VoicemailInfo               `json:"voicemail"`
	Greetings               []GreetingInfo              `json:"greetings"`
	Screening               string                      `json:"screening"`
	SharedLines             SharedLinesInfo             `json:"sharedLines"`
}

// SharedLinesInfo Shared Lines Info
type SharedLinesInfo struct {
	Timeout int `json:"timeout"`
}

// AudioPromptInfo Audio Prompt Info
type AudioPromptInfo struct {
	Uri string `json:"uri"`
	Id  string `json:"id"`
}

// CalledNumberInfo Called Number Info
type CalledNumberInfo struct {
	PhoneNumber string `json:"phoneNumber"`
}

// CallersInfo Callers Info
type CallersInfo struct {
	CallerId string `json:"callerId"`
	Name     string `json:"name"`
}

// CallersInfoRequest Callers Info Request
type CallersInfoRequest struct {
	CallerId string `json:"callerId"`
	Name     string `json:"name"`
}

// CompanyAnsweringRuleCalledNumberInfo Company Answering Rule Called Number Info
type CompanyAnsweringRuleCalledNumberInfo struct {
	Id          string `json:"id"`
	PhoneNumber string `json:"phoneNumber"`
}

// CompanyAnsweringRuleCalledNumberInfoRequest Company Answering Rule Called Number Info Request
type CompanyAnsweringRuleCalledNumberInfoRequest struct {
	Id          string `json:"id"`
	PhoneNumber string `json:"phoneNumber"`
}

// CompanyAnsweringRuleCallersInfoRequest Company Answering Rule Callers Info Request
type CompanyAnsweringRuleCallersInfoRequest struct {
	CallerId string `json:"callerId"`
	Name     string `json:"name"`
}

// CompanyAnsweringRuleExtensionInfoRequest Company Answering Rule Extension Info Request
type CompanyAnsweringRuleExtensionInfoRequest struct {
	Id string `json:"id"`
}

// CompanyAnsweringRuleInfo Company Answering Rule Info
type CompanyAnsweringRuleInfo struct {
	Id                 string                                        `json:"id"`
	Uri                string                                        `json:"uri"`
	Enabled            bool                                          `json:"enabled"`
	Type               string                                        `json:"type"`
	Name               string                                        `json:"name"`
	Callers            []CompanyAnsweringRuleCallersInfoRequest      `json:"callers"`
	CalledNumbers      []CompanyAnsweringRuleCalledNumberInfoRequest `json:"calledNumbers"`
	Schedule           CompanyAnsweringRuleScheduleInfo              `json:"schedule"`
	CallHandlingAction string                                        `json:"callHandlingAction"`
	Extension          CompanyAnsweringRuleExtensionInfoRequest      `json:"extension"`
	Greetings          []GreetingInfo                                `json:"greetings"`
}

// CompanyAnsweringRuleList Company Answering Rule List
type CompanyAnsweringRuleList struct {
	Uri        string                         `json:"uri"`
	Records    []ListCompanyAnsweringRuleInfo `json:"records"`
	Paging     CallHandlingPagingInfo         `json:"paging"`
	Navigation CallHandlingNavigationInfo     `json:"navigation"`
}

// CompanyAnsweringRuleRequest Company Answering Rule Request
type CompanyAnsweringRuleRequest struct {
	Name               string                                   `json:"name"`
	Enabled            bool                                     `json:"enabled"`
	Type               string                                   `json:"type"`
	Callers            []CompanyAnsweringRuleCallersInfoRequest `json:"callers"`
	CalledNumbers      []CompanyAnsweringRuleCalledNumberInfo   `json:"calledNumbers"`
	Schedule           CompanyAnsweringRuleScheduleInfoRequest  `json:"schedule"`
	CallHandlingAction string                                   `json:"callHandlingAction"`
	Extension          string                                   `json:"extension"`
	Greetings          []GreetingInfo                           `json:"greetings"`
}

// CompanyAnsweringRuleScheduleInfo Company Answering Rule Schedule Info
type CompanyAnsweringRuleScheduleInfo struct {
	WeeklyRanges CompanyAnsweringRuleWeeklyScheduleInfoRequest `json:"weeklyRanges"`
	Ranges       []RangesInfo                                  `json:"ranges"`
	Ref          string                                        `json:"@ref"`
}

// CompanyAnsweringRuleScheduleInfoRequest Company Answering Rule Schedule Info Request
type CompanyAnsweringRuleScheduleInfoRequest struct {
	WeeklyRanges CompanyAnsweringRuleWeeklyScheduleInfoRequest `json:"weeklyRanges"`
	Ranges       []RangesInfo                                  `json:"ranges"`
	Ref          string                                        `json:"@ref"`
}

// CompanyAnsweringRuleTimeIntervalRequest Company Answering Rule Time Interval Request
type CompanyAnsweringRuleTimeIntervalRequest struct {
	From string `json:"from"`
	To   string `json:"to"`
}

// CompanyAnsweringRuleWeeklyScheduleInfoRequest Company Answering Rule Weekly Schedule Info Request
type CompanyAnsweringRuleWeeklyScheduleInfoRequest struct {
	Monday    []CompanyAnsweringRuleTimeIntervalRequest `json:"monday"`
	Tuesday   []CompanyAnsweringRuleTimeIntervalRequest `json:"tuesday"`
	Wednesday []CompanyAnsweringRuleTimeIntervalRequest `json:"wednesday"`
	Thursday  []CompanyAnsweringRuleTimeIntervalRequest `json:"thursday"`
	Friday    []CompanyAnsweringRuleTimeIntervalRequest `json:"friday"`
	Saturday  []CompanyAnsweringRuleTimeIntervalRequest `json:"saturday"`
	Sunday    []CompanyAnsweringRuleTimeIntervalRequest `json:"sunday"`
}

// CompanyAnsweringRuleUpdate Company Answering Rule Update
type CompanyAnsweringRuleUpdate struct {
	Enabled            bool                                     `json:"enabled"`
	Name               string                                   `json:"name"`
	Callers            []CompanyAnsweringRuleCallersInfoRequest `json:"callers"`
	CalledNumbers      []CompanyAnsweringRuleCalledNumberInfo   `json:"calledNumbers"`
	Schedule           CompanyAnsweringRuleScheduleInfoRequest  `json:"schedule"`
	CallHandlingAction string                                   `json:"callHandlingAction"`
	Type               string                                   `json:"type"`
	Extension          string                                   `json:"extension"`
	Greetings          []GreetingInfo                           `json:"greetings"`
}

// CompanyBusinessHours Company Business Hours
type CompanyBusinessHours struct {
	Uri      string                           `json:"uri"`
	Schedule CompanyBusinessHoursScheduleInfo `json:"schedule"`
}

// CompanyBusinessHoursScheduleInfo Company Business Hours Schedule Info
type CompanyBusinessHoursScheduleInfo struct {
	WeeklyRanges WeeklyScheduleInfo `json:"weeklyRanges"`
}

// CompanyBusinessHoursUpdateRequest Company Business Hours Update Request
type CompanyBusinessHoursUpdateRequest struct {
	Schedule CompanyBusinessHoursScheduleInfo `json:"schedule"`
}

// CreateAnsweringRuleRequest Create Answering Rule Request
type CreateAnsweringRuleRequest struct {
	Enabled                 bool                        `json:"enabled"`
	Type                    string                      `json:"type"`
	Name                    string                      `json:"name"`
	Callers                 []CallersInfoRequest        `json:"callers"`
	CalledNumbers           []CalledNumberInfo          `json:"calledNumbers"`
	Schedule                ScheduleInfo                `json:"schedule"`
	CallHandlingAction      string                      `json:"callHandlingAction"`
	Forwarding              ForwardingInfo              `json:"forwarding"`
	UnconditionalForwarding UnconditionalForwardingInfo `json:"unconditionalForwarding"`
	Queue                   QueueInfo                   `json:"queue"`
	Transfer                TransferredExtensionInfo    `json:"transfer"`
	Voicemail               VoicemailInfo               `json:"voicemail"`
	Greetings               []GreetingInfo              `json:"greetings"`
	Screening               string                      `json:"screening"`
}

// CreateForwardingNumberRequest Create Forwarding Number Request
type CreateForwardingNumberRequest struct {
	FlipNumber  int                              `json:"flipNumber"`
	PhoneNumber string                           `json:"phoneNumber"`
	Label       string                           `json:"label"`
	Type        string                           `json:"type"`
	Device      CreateForwardingNumberDeviceInfo `json:"device"`
}

// CreateForwardingNumberDeviceInfo Create Forwarding Number Device Info
type CreateForwardingNumberDeviceInfo struct {
	Id string `json:"id"`
}

// CustomCompanyGreetingInfo Custom Company Greeting Info
type CustomCompanyGreetingInfo struct {
	Uri           string                            `json:"uri"`
	Id            string                            `json:"id"`
	Type          string                            `json:"type"`
	ContentType   string                            `json:"contentType"`
	ContentUri    string                            `json:"contentUri"`
	AnsweringRule CustomGreetingAnsweringRuleInfo   `json:"answeringRule"`
	Language      CustomCompanyGreetingLanguageInfo `json:"language"`
}

// CustomGreetingAnsweringRuleInfo Custom Greeting Answering Rule Info
type CustomGreetingAnsweringRuleInfo struct {
	Uri string `json:"uri"`
	Id  string `json:"id"`
}

// DictionaryGreetingInfo Dictionary Greeting Info
type DictionaryGreetingInfo struct {
	Id         string                     `json:"id"`
	Uri        string                     `json:"uri"`
	Name       string                     `json:"name"`
	UsageType  string                     `json:"usageType"`
	Text       string                     `json:"text"`
	ContentUri string                     `json:"contentUri"`
	Type       string                     `json:"type"`
	Category   string                     `json:"category"`
	Navigation CallHandlingNavigationInfo `json:"navigation"`
	Paging     CallHandlingPagingInfo     `json:"paging"`
}

// DictionaryGreetingList Dictionary Greeting List
type DictionaryGreetingList struct {
	Uri        string                     `json:"uri"`
	Records    []DictionaryGreetingInfo   `json:"records"`
	Navigation CallHandlingNavigationInfo `json:"navigation"`
	Paging     CallHandlingPagingInfo     `json:"paging"`
}

// ExtensionInfo Extension Info
type ExtensionInfo struct {
	Id              string `json:"id"`
	Uri             string `json:"uri"`
	Name            string `json:"name"`
	ExtensionNumber string `json:"extensionNumber"`
	PartnerId       string `json:"partnerId"`
}

// FixedOrderAgents Fixed Order Agents
type FixedOrderAgents struct {
	Extension FixedOrderAgentsExtensionInfo `json:"extension"`
	Index     int                           `json:"index"`
}

// FixedOrderAgentsExtensionInfo Fixed Order Agents Extension Info
type FixedOrderAgentsExtensionInfo struct {
	Id              string `json:"id"`
	Uri             string `json:"uri"`
	ExtensionNumber string `json:"extensionNumber"`
	Name            string `json:"name"`
}

// ForwardingInfo Forwarding Info
type ForwardingInfo struct {
	NotifyMySoftPhones    bool       `json:"notifyMySoftPhones"`
	NotifyAdminSoftPhones bool       `json:"notifyAdminSoftPhones"`
	SoftPhonesRingCount   int        `json:"softPhonesRingCount"`
	RingingMode           string     `json:"ringingMode"`
	Rules                 []RuleInfo `json:"rules"`
	MobileTimeout         bool       `json:"mobileTimeout"`
}

// ForwardingInfoCreateRuleRequest Forwarding Info Create Rule Request
type ForwardingInfoCreateRuleRequest struct {
	NotifyMySoftPhones    bool                        `json:"notifyMySoftPhones"`
	NotifyAdminSoftPhones bool                        `json:"notifyAdminSoftPhones"`
	SoftPhonesRingCount   int                         `json:"softPhonesRingCount"`
	RingingMode           string                      `json:"ringingMode"`
	Rules                 []RuleInfoCreateRuleRequest `json:"rules"`
	MobileTimeout         bool                        `json:"mobileTimeout"`
}

// ForwardingNumberInfo Forwarding Number Info
type ForwardingNumberInfo struct {
	Id          string                           `json:"id"`
	Uri         string                           `json:"uri"`
	PhoneNumber string                           `json:"phoneNumber"`
	Label       string                           `json:"label"`
	Features    []string                         `json:"features"`
	FlipNumber  string                           `json:"flipNumber"`
	Device      CreateForwardingNumberDeviceInfo `json:"device"`
	Type        string                           `json:"type"`
}

// CreateAnsweringRuleForwardingNumberInfo Create Answering Rule Forwarding Number Info
type CreateAnsweringRuleForwardingNumberInfo struct {
	Id          string `json:"id"`
	Uri         string `json:"uri"`
	PhoneNumber string `json:"phoneNumber"`
	Label       string `json:"label"`
	Type        string `json:"type"`
}

// ForwardingNumberInfoRulesCreateRuleRequest Forwarding Number Info Rules Create Rule Request
type ForwardingNumberInfoRulesCreateRuleRequest struct {
	Id          string `json:"id"`
	Type        string `json:"type"`
	Uri         string `json:"uri"`
	PhoneNumber string `json:"phoneNumber"`
	Label       string `json:"label"`
}

// GetExtensionForwardingNumberListResponse Get Extension Forwarding Number List Response
type GetExtensionForwardingNumberListResponse struct {
	Uri        string                     `json:"uri"`
	Records    []ForwardingNumberInfo     `json:"records"`
	Navigation CallHandlingNavigationInfo `json:"navigation"`
	Paging     CallHandlingPagingInfo     `json:"paging"`
}

// GetUserBusinessHoursResponse Get User Business Hours Response
type GetUserBusinessHoursResponse struct {
	Uri      string                        `json:"uri"`
	Schedule ScheduleInfoUserBusinessHours `json:"schedule"`
}

// GreetingInfo Greeting Info
type GreetingInfo struct {
	Type   string                    `json:"type"`
	Preset PresetInfo                `json:"preset"`
	Custom CustomGreetingInfoRequest `json:"custom"`
}

// IVRMenuActionsInfo IVRMenu Actions Info
type IVRMenuActionsInfo struct {
	Input       string               `json:"input"`
	Action      string               `json:"action"`
	Extension   IVRMenuExtensionInfo `json:"extension"`
	PhoneNumber string               `json:"phoneNumber"`
}

// IVRMenuExtensionInfo IVRMenu Extension Info
type IVRMenuExtensionInfo struct {
	Uri  string `json:"uri"`
	Id   string `json:"id"`
	Name string `json:"name"`
}

// IVRMenuInfo IVRMenu Info
type IVRMenuInfo struct {
	Id              string               `json:"id"`
	Uri             string               `json:"uri"`
	Name            string               `json:"name"`
	ExtensionNumber string               `json:"extensionNumber"`
	Site            IVRMenuSiteInfo      `json:"site"`
	Prompt          IVRMenuPromptInfo    `json:"prompt"`
	Actions         []IVRMenuActionsInfo `json:"actions"`
}

// IVRMenuSiteInfo IVRMenu Site Info
type IVRMenuSiteInfo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// IVRMenuPromptInfo IVRMenu Prompt Info
type IVRMenuPromptInfo struct {
	Mode     string             `json:"mode"`
	Audio    AudioPromptInfo    `json:"audio"`
	Text     string             `json:"text"`
	Language PromptLanguageInfo `json:"language"`
}

// IVRPrompts IVRPrompts
type IVRPrompts struct {
	Uri        string                     `json:"uri"`
	Records    []PromptInfo               `json:"records"`
	Navigation CallHandlingNavigationInfo `json:"navigation"`
	Paging     CallHandlingPagingInfo     `json:"paging"`
}

// ListCompanyAnsweringRuleInfo List Company Answering Rule Info
type ListCompanyAnsweringRuleInfo struct {
	Id            string                            `json:"id"`
	Uri           string                            `json:"uri"`
	Enabled       bool                              `json:"enabled"`
	Type          string                            `json:"type"`
	Name          string                            `json:"name"`
	CalledNumbers []CalledNumberInfo                `json:"calledNumbers"`
	Extension     CompanyAnsweringRuleExtensionInfo `json:"extension"`
}

// CallHandlingNavigationInfo Call Handling Navigation Info
type CallHandlingNavigationInfo struct {
	FirstPage    CallHandlingNavigationInfoUri `json:"firstPage"`
	NextPage     CallHandlingNavigationInfoUri `json:"nextPage"`
	PreviousPage CallHandlingNavigationInfoUri `json:"previousPage"`
	LastPage     CallHandlingNavigationInfoUri `json:"lastPage"`
}

// CallHandlingNavigationInfoUri Call Handling Navigation Info Uri
type CallHandlingNavigationInfoUri struct {
	Uri string `json:"uri"`
}

// CallHandlingPagingInfo Call Handling Paging Info
type CallHandlingPagingInfo struct {
	Page          int `json:"page"`
	PerPage       int `json:"perPage"`
	PageStart     int `json:"pageStart"`
	PageEnd       int `json:"pageEnd"`
	TotalPages    int `json:"totalPages"`
	TotalElements int `json:"totalElements"`
}

// PresetInfo Preset Info
type PresetInfo struct {
	Uri  string `json:"uri"`
	Id   string `json:"id"`
	Name string `json:"name"`
}

// PromptInfo Prompt Info
type PromptInfo struct {
	Uri         string `json:"uri"`
	Id          string `json:"id"`
	ContentType string `json:"contentType"`
	ContentUri  string `json:"contentUri"`
	Filename    string `json:"filename"`
}

// PromptLanguageInfo Prompt Language Info
type PromptLanguageInfo struct {
	Uri        string `json:"uri"`
	Id         string `json:"id"`
	Name       string `json:"name"`
	LocaleCode string `json:"localeCode"`
}

// QueueInfo Queue Info
type QueueInfo struct {
	TransferMode                string                        `json:"transferMode"`
	Transfer                    []TransferInfo                `json:"transfer"`
	NoAnswerAction              string                        `json:"noAnswerAction"`
	FixedOrderAgents            []FixedOrderAgents            `json:"fixedOrderAgents"`
	HoldAudioInterruptionMode   string                        `json:"holdAudioInterruptionMode"`
	HoldAudioInterruptionPeriod int                           `json:"holdAudioInterruptionPeriod"`
	HoldTimeExpirationAction    string                        `json:"holdTimeExpirationAction"`
	AgentTimeout                int                           `json:"agentTimeout"`
	WrapUpTime                  int                           `json:"wrapUpTime"`
	HoldTime                    int                           `json:"holdTime"`
	MaxCallers                  int                           `json:"maxCallers"`
	MaxCallersAction            string                        `json:"maxCallersAction"`
	UnconditionalForwarding     []UnconditionalForwardingInfo `json:"unconditionalForwarding"`
}

// TransferInfo Transfer Info
type TransferInfo struct {
	Extension TransferExtensionInfo `json:"extension"`
	Action    string                `json:"action"`
}

// TransferExtensionInfo Transfer Extension Info
type TransferExtensionInfo struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	ExtensionNumber string `json:"extensionNumber"`
}

// RangesInfo Ranges Info
type RangesInfo struct {
	From string `json:"from"`
	To   string `json:"to"`
}

// RecipientInfo Recipient Info
type RecipientInfo struct {
	Uri string `json:"uri"`
	Id  int    `json:"id"`
}

// RuleInfo Rule Info
type RuleInfo struct {
	Index             int                                       `json:"index"`
	RingCount         int                                       `json:"ringCount"`
	Enabled           bool                                      `json:"enabled"`
	ForwardingNumbers []CreateAnsweringRuleForwardingNumberInfo `json:"forwardingNumbers"`
}

// RuleInfoCreateRuleRequest Rule Info Create Rule Request
type RuleInfoCreateRuleRequest struct {
	Index             int                                          `json:"index"`
	RingCount         int                                          `json:"ringCount"`
	Enabled           bool                                         `json:"enabled"`
	ForwardingNumbers []ForwardingNumberInfoRulesCreateRuleRequest `json:"forwardingNumbers"`
}

// ScheduleInfo Schedule Info
type ScheduleInfo struct {
	WeeklyRanges WeeklyScheduleInfo `json:"weeklyRanges"`
	Ranges       []RangesInfo       `json:"ranges"`
	Ref          string             `json:"@ref"`
}

// ScheduleInfoUserBusinessHours Schedule Info User Business Hours
type ScheduleInfoUserBusinessHours struct {
	WeeklyRanges WeeklyScheduleInfo `json:"weeklyRanges"`
}

// TimeInterval Time Interval
type TimeInterval struct {
	From string `json:"from"`
	To   string `json:"to"`
}

// TransferredExtensionInfo Transferred Extension Info
type TransferredExtensionInfo struct {
	Extension ExtensionInfo `json:"extension"`
}

// UnconditionalForwardingInfo Unconditional Forwarding Info
type UnconditionalForwardingInfo struct {
	PhoneNumber string `json:"phoneNumber"`
	Action      string `json:"action"`
}

// UpdateAnsweringRuleRequest Update Answering Rule Request
type UpdateAnsweringRuleRequest struct {
	Id                      string                          `json:"id"`
	Forwarding              ForwardingInfoCreateRuleRequest `json:"forwarding"`
	Enabled                 bool                            `json:"enabled"`
	Name                    string                          `json:"name"`
	Callers                 []CallersInfoRequest            `json:"callers"`
	CalledNumbers           []CalledNumberInfo              `json:"calledNumbers"`
	Schedule                ScheduleInfo                    `json:"schedule"`
	CallHandlingAction      string                          `json:"callHandlingAction"`
	Type                    string                          `json:"type"`
	UnconditionalForwarding UnconditionalForwardingInfo     `json:"unconditionalForwarding"`
	Queue                   QueueInfo                       `json:"queue"`
	Voicemail               VoicemailInfo                   `json:"voicemail"`
	Greetings               []GreetingInfo                  `json:"greetings"`
	Screening               string                          `json:"screening"`
	ShowInactiveNumbers     bool                            `json:"showInactiveNumbers"`
	Transfer                TransferredExtensionInfo        `json:"transfer"`
}

// UpdateForwardingNumberRequest Update Forwarding Number Request
type UpdateForwardingNumberRequest struct {
	PhoneNumber string `json:"phoneNumber"`
	Label       string `json:"label"`
	FlipNumber  string `json:"flipNumber"`
	Type        string `json:"type"`
}

// UserBusinessHoursScheduleInfo User Business Hours Schedule Info
type UserBusinessHoursScheduleInfo struct {
	WeeklyRanges WeeklyScheduleInfo `json:"weeklyRanges"`
}

// UserBusinessHoursUpdateResponse User Business Hours Update Response
type UserBusinessHoursUpdateResponse struct {
	Uri      string                        `json:"uri"`
	Schedule UserBusinessHoursScheduleInfo `json:"schedule"`
}

// UserBusinessHoursUpdateRequest User Business Hours Update Request
type UserBusinessHoursUpdateRequest struct {
	Schedule UserBusinessHoursScheduleInfo `json:"schedule"`
}

// VoicemailInfo Voicemail Info
type VoicemailInfo struct {
	Enabled   bool          `json:"enabled"`
	Recipient RecipientInfo `json:"recipient"`
}

// WeeklyScheduleInfo Weekly Schedule Info
type WeeklyScheduleInfo struct {
	Monday    []TimeInterval `json:"monday"`
	Tuesday   []TimeInterval `json:"tuesday"`
	Wednesday []TimeInterval `json:"wednesday"`
	Thursday  []TimeInterval `json:"thursday"`
	Friday    []TimeInterval `json:"friday"`
	Saturday  []TimeInterval `json:"saturday"`
	Sunday    []TimeInterval `json:"sunday"`
}

// BulkAccountCallRecordingsResource Bulk Account Call Recordings Resource
type BulkAccountCallRecordingsResource struct {
	AddedExtensions   []CallRecordingExtensionResource `json:"addedExtensions"`
	UpdatedExtensions []CallRecordingExtensionResource `json:"updatedExtensions"`
	RemovedExtensions []CallRecordingExtensionResource `json:"removedExtensions"`
}

// CallRecordingExtensionResource Call Recording Extension Resource
type CallRecordingExtensionResource struct {
	Id              string `json:"id"`
	Uri             string `json:"uri"`
	ExtensionNumber string `json:"extensionNumber"`
	Type            string `json:"type"`
	CallDirection   string `json:"callDirection"`
}

// CallRecordingSettingsResource Call Recording Settings Resource
type CallRecordingSettingsResource struct {
	OnDemand  OnDemandResource           `json:"onDemand"`
	Automatic AutomaticRecordingResource `json:"automatic"`
	Greetings []GreetingResource         `json:"greetings"`
}

// GreetingResource Greeting Resource
type GreetingResource struct {
	Type string `json:"type"`
	Mode string `json:"mode"`
}

// AutomaticRecordingResource Automatic Recording Resource
type AutomaticRecordingResource struct {
	Enabled                  bool `json:"enabled"`
	OutboundCallTones        bool `json:"outboundCallTones"`
	OutboundCallAnnouncement bool `json:"outboundCallAnnouncement"`
	AllowMute                bool `json:"allowMute"`
	ExtensionCount           int  `json:"extensionCount"`
}

// OnDemandResource On Demand Resource
type OnDemandResource struct {
	Enabled bool `json:"enabled"`
}

// CustomUserGreetingInfo Custom User Greeting Info
type CustomUserGreetingInfo struct {
	Uri           string                          `json:"uri"`
	Id            string                          `json:"id"`
	Type          string                          `json:"type"`
	ContentType   string                          `json:"contentType"`
	ContentUri    string                          `json:"contentUri"`
	AnsweringRule CustomGreetingAnsweringRuleInfo `json:"answeringRule"`
}

// CustomCompanyGreetingLanguageInfo Custom Company Greeting Language Info
type CustomCompanyGreetingLanguageInfo struct {
	Id         string `json:"id"`
	Uri        string `json:"uri"`
	Name       string `json:"name"`
	LocaleCode string `json:"localeCode"`
}

// CallerBlockingSettings Caller Blocking Settings
type CallerBlockingSettings struct {
	Mode       string                      `json:"mode"`
	NoCallerId string                      `json:"noCallerId"`
	PayPhones  string                      `json:"payPhones"`
	Greetings  []BlockedCallerGreetingInfo `json:"greetings"`
}

// BlockedCallerGreetingInfo Blocked Caller Greeting Info
type BlockedCallerGreetingInfo struct {
	Type   string     `json:"type"`
	Preset PresetInfo `json:"preset"`
}

// CallerBlockingSettingsUpdate Caller Blocking Settings Update
type CallerBlockingSettingsUpdate struct {
	Mode       string                      `json:"mode"`
	NoCallerId string                      `json:"noCallerId"`
	PayPhones  string                      `json:"payPhones"`
	Greetings  []BlockedCallerGreetingInfo `json:"greetings"`
}

// BlockedAllowedPhoneNumbersList Blocked Allowed Phone Numbers List
type BlockedAllowedPhoneNumbersList struct {
	Uri        string                          `json:"uri"`
	Records    []BlockedAllowedPhoneNumberInfo `json:"records"`
	Navigation CallHandlingNavigationInfo      `json:"navigation"`
	Paging     CallHandlingPagingInfo          `json:"paging"`
}

// BlockedAllowedPhoneNumberInfo Blocked Allowed Phone Number Info
type BlockedAllowedPhoneNumberInfo struct {
	Uri         string `json:"uri"`
	Id          string `json:"id"`
	PhoneNumber string `json:"phoneNumber"`
	Label       string `json:"label"`
	Status      string `json:"status"`
}

// AddBlockedAllowedPhoneNumber Add Blocked Allowed Phone Number
type AddBlockedAllowedPhoneNumber struct {
	PhoneNumber string `json:"phoneNumber"`
	Label       string `json:"label"`
	Status      string `json:"status"`
}

// CallRecordingCustomGreetings Call Recording Custom Greetings
type CallRecordingCustomGreetings struct {
	Records []CallRecordingCustomGreeting `json:"records"`
}

// CallRecordingCustomGreeting Call Recording Custom Greeting
type CallRecordingCustomGreeting struct {
	Type     string                              `json:"type"`
	Custom   CallRecordingCustomGreetingData     `json:"custom"`
	Language CallRecordingCustomGreetingLanguage `json:"language"`
}

// CallRecordingCustomGreetingData Call Recording Custom Greeting Data
type CallRecordingCustomGreetingData struct {
	Uri string `json:"uri"`
	Id  string `json:"id"`
}

// CustomGreetingInfoRequest Custom Greeting Info Request
type CustomGreetingInfoRequest struct {
	Id string `json:"id"`
}

// CallRecordingCustomGreetingLanguage Call Recording Custom Greeting Language
type CallRecordingCustomGreetingLanguage struct {
	Uri        string `json:"uri"`
	Id         string `json:"id"`
	Name       string `json:"name"`
	LocaleCode string `json:"localeCode"`
}

// UserAnsweringRuleList User Answering Rule List
type UserAnsweringRuleList struct {
	Uri        string                          `json:"uri"`
	Records    []UserAnsweringRuleListRecord   `json:"records"`
	Paging     UserAnsweringRuleListPaging     `json:"paging"`
	Navigation UserAnsweringRuleListNavigation `json:"navigation"`
}

// UserAnsweringRuleListRecord User Answering Rule List Record
type UserAnsweringRuleListRecord struct {
	Uri         string          `json:"uri"`
	Id          string          `json:"id"`
	Type        string          `json:"type"`
	Name        string          `json:"name"`
	Enabled     bool            `json:"enabled"`
	SharedLines SharedLinesInfo `json:"sharedLines"`
}

// UserAnsweringRuleListPaging User Answering Rule List Paging
type UserAnsweringRuleListPaging struct {
	Page          int `json:"page"`
	TotalPages    int `json:"totalPages"`
	PerPage       int `json:"perPage"`
	TotalElements int `json:"totalElements"`
	PageStart     int `json:"pageStart"`
	PageEnd       int `json:"pageEnd"`
}

// UserAnsweringRuleListNavigation User Answering Rule List Navigation
type UserAnsweringRuleListNavigation struct {
	FirstPage    UserAnsweringRuleListNavigationPage `json:"firstPage"`
	NextPage     UserAnsweringRuleListNavigationPage `json:"nextPage"`
	PreviousPage UserAnsweringRuleListNavigationPage `json:"previousPage"`
	LastPage     UserAnsweringRuleListNavigationPage `json:"lastPage"`
}

// UserAnsweringRuleListNavigationPage User Answering Rule List Navigation Page
type UserAnsweringRuleListNavigationPage struct {
	Uri string `json:"uri"`
}

// CallRecordingExtensions Call Recording Extensions
type CallRecordingExtensions struct {
	Uri        string                       `json:"uri"`
	Records    []CallRecordingExtensionInfo `json:"records"`
	Navigation CallHandlingNavigationInfo   `json:"navigation"`
	Paging     CallHandlingPagingInfo       `json:"paging"`
}

// CallRecordingExtensionInfo Call Recording Extension Info
type CallRecordingExtensionInfo struct {
	Id              string `json:"id"`
	Uri             string `json:"uri"`
	ExtensionNumber string `json:"extensionNumber"`
	Name            string `json:"name"`
}

// UpdateIVRPromptRequest Update IVRPrompt Request
type UpdateIVRPromptRequest struct {
	Filename string `json:"filename"`
}

// CompanyAnsweringRuleExtensionInfo Company Answering Rule Extension Info
type CompanyAnsweringRuleExtensionInfo struct {
	Id string `json:"id"`
}

// GetVersionsResponse Get Versions Response
type GetVersionsResponse struct {
	Uri            string        `json:"uri"`
	ApiVersions    []VersionInfo `json:"apiVersions"`
	ServerVersion  string        `json:"serverVersion"`
	ServerRevision string        `json:"serverRevision"`
}

// VersionInfo Version Info
type VersionInfo struct {
	Uri           string `json:"uri"`
	VersionString string `json:"versionString"`
	ReleaseDate   string `json:"releaseDate"`
	UriString     string `json:"uriString"`
}

// GetVersionResponse Get Version Response
type GetVersionResponse struct {
	Uri           string `json:"uri"`
	VersionString string `json:"versionString"`
	ReleaseDate   string `json:"releaseDate"`
	UriString     string `json:"uriString"`
}

// CreateSMSMessageBatchRequest Create SMSMessage Batch Request
type CreateSMSMessageBatchRequest struct {
	From     string                 `json:"from"`
	Text     string                 `json:"text"`
	Messages []MessageCreateRequest `json:"messages"`
}

// MessageBatchResponse Message Batch Response
type MessageBatchResponse struct {
	Id               string `json:"id"`
	From             string `json:"from"`
	BatchSize        int    `json:"batchSize"`
	ProcessedCount   int    `json:"processedCount"`
	LastModifiedTime string `json:"lastModifiedTime"`
	Status           string `json:"status"`
	CreationTime     string `json:"creationTime"`
}

// MessageCreateRequest Message Create Request
type MessageCreateRequest struct {
	To   []string `json:"to"`
	Text string   `json:"text"`
}

// MessageDetailsResponse Message Details Response
type MessageDetailsResponse struct {
	Id               string   `json:"id"`
	From             string   `json:"from"`
	To               []string `json:"to"`
	Text             string   `json:"text"`
	CreationTime     string   `json:"creationTime"`
	LastModifiedTime string   `json:"lastModifiedTime"`
	MessageStatus    string   `json:"messageStatus"`
	SegmentCount     int      `json:"segmentCount"`
	Cost             float64  `json:"cost"`
	BatchId          string   `json:"batchId"`
	Direction        string   `json:"direction"`
	ErrorCode        string   `json:"errorCode"`
}

// MessageListResponse Message List Response
type MessageListResponse struct {
	Records []MessageListMessageResponse `json:"records"`
	Paging  PagingResource               `json:"paging"`
}

// MessageListMessageResponse Message List Message Response
type MessageListMessageResponse struct {
	Id               int      `json:"id"`
	BatchId          string   `json:"batchId"`
	From             string   `json:"from"`
	To               []string `json:"to"`
	CreationTime     string   `json:"creationTime"`
	LastModifiedTime string   `json:"lastModifiedTime"`
	MessageStatus    string   `json:"messageStatus"`
	SegmentCount     int      `json:"segmentCount"`
	Text             string   `json:"text"`
	Cost             float64  `json:"cost"`
	Direction        string   `json:"direction"`
	ErrorCode        string   `json:"errorCode"`
}

// OptOutListResponse Opt Out List Response
type OptOutListResponse struct {
	Records []OptOutResponse `json:"records"`
	Paging  PagingResource   `json:"paging"`
}

// OptOutResponse Opt Out Response
type OptOutResponse struct {
	From string `json:"from"`
	To   string `json:"to"`
}

// PagingResource Paging Resource
type PagingResource struct {
	PageToken         string `json:"pageToken"`
	PerPage           int    `json:"perPage"`
	FirstPageToken    string `json:"firstPageToken"`
	PreviousPageToken string `json:"previousPageToken"`
	NextPageToken     string `json:"nextPageToken"`
}

// UserCallLogResponse User Call Log Response
type UserCallLogResponse struct {
	Records    []UserCallLogRecord   `json:"records"`
	Navigation CallLogNavigationInfo `json:"navigation"`
	Paging     CallLogPagingInfo     `json:"paging"`
}

// CallLogSync Call Log Sync
type CallLogSync struct {
	Uri      string              `json:"uri"`
	Records  []UserCallLogRecord `json:"records"`
	SyncInfo SyncInfoCallLog     `json:"syncInfo"`
}

// UserCallLogRecord User Call Log Record
type UserCallLogRecord struct {
	Id                 string                 `json:"id"`
	Uri                string                 `json:"uri"`
	SessionId          string                 `json:"sessionId"`
	TelephonySessionId string                 `json:"telephonySessionId"`
	From               CallLogCallerInfo      `json:"from"`
	To                 CallLogCallerInfo      `json:"to"`
	Extension          ExtensionInfoCallLog   `json:"extension"`
	Type               string                 `json:"type"`
	Transport          string                 `json:"transport"`
	Legs               []CallLogRecordLegInfo `json:"legs"`
	Billing            BillingInfo            `json:"billing"`
	Direction          string                 `json:"direction"`
	Message            CallLogRecordMessage   `json:"message"`
	StartTime          string                 `json:"startTime"`
	Delegate           DelegateInfo           `json:"@delegate"`
	Deleted            bool                   `json:"deleted"`
	Duration           int                    `json:"duration"`
	LastModifiedTime   string                 `json:"lastModifiedTime"`
	Recording          CallLogRecordingInfo   `json:"recording"`
	ShortRecording     bool                   `json:"shortRecording"`
	Action             string                 `json:"action"`
	Result             string                 `json:"result"`
	Reason             string                 `json:"reason"`
	ReasonDescription  string                 `json:"reasonDescription"`
}

// CompanyActiveCallsResponse Company Active Calls Response
type CompanyActiveCallsResponse struct {
	Uri        string                 `json:"uri"`
	Records    []CompanyCallLogRecord `json:"records"`
	Navigation CallLogNavigationInfo  `json:"navigation"`
	Paging     CallLogPagingInfo      `json:"paging"`
}

// UserActiveCallsResponse User Active Calls Response
type UserActiveCallsResponse struct {
	Uri        string                `json:"uri"`
	Records    []UserCallLogRecord   `json:"records"`
	Navigation CallLogNavigationInfo `json:"navigation"`
	Paging     CallLogPagingInfo     `json:"paging"`
}

// BillingInfo Billing Info
type BillingInfo struct {
	CostIncluded  float64 `json:"costIncluded"`
	CostPurchased float64 `json:"costPurchased"`
}

// AccountCallLogResponse Account Call Log Response
type AccountCallLogResponse struct {
	Uri        string                 `json:"uri"`
	Records    []CompanyCallLogRecord `json:"records"`
	Navigation CallLogNavigationInfo  `json:"navigation"`
	Paging     CallLogPagingInfo      `json:"paging"`
}

// GetCallRecordingResponse Get Call Recording Response
type GetCallRecordingResponse struct {
	Id          string `json:"id"`
	ContentUri  string `json:"contentUri"`
	ContentType string `json:"contentType"`
	Duration    int    `json:"duration"`
}

// DelegateInfo Delegate Info
type DelegateInfo struct {
	Extension    DelegateExtensionInfo `json:"extension"`
	CallPickup   bool                  `json:"callPickup"`
	Conferencing bool                  `json:"conferencing"`
}

// CallLogRecordMessage Call Log Record Message
type CallLogRecordMessage struct {
	Id   string `json:"id"`
	Type string `json:"type"`
	Uri  string `json:"uri"`
}

// SyncInfoCallLog Sync Info Call Log
type SyncInfoCallLog struct {
	SyncType  string `json:"syncType"`
	SyncToken string `json:"syncToken"`
	SyncTime  string `json:"syncTime"`
}

// CallLogCallerInfo Call Log Caller Info
type CallLogCallerInfo struct {
	PhoneNumber     string                  `json:"phoneNumber"`
	ExtensionNumber string                  `json:"extensionNumber"`
	ExtensionId     string                  `json:"extensionId"`
	Location        string                  `json:"location"`
	Name            string                  `json:"name"`
	Device          CallLogRecordDeviceInfo `json:"device"`
}

// CallLogRecordingInfo Call Log Recording Info
type CallLogRecordingInfo struct {
	Id         string `json:"id"`
	Uri        string `json:"uri"`
	Type       string `json:"type"`
	ContentUri string `json:"contentUri"`
}

// CallLogRecordDeviceInfo Call Log Record Device Info
type CallLogRecordDeviceInfo struct {
	Id  string `json:"id"`
	Uri string `json:"uri"`
}

// CallLogNavigationInfo Call Log Navigation Info
type CallLogNavigationInfo struct {
	FirstPage    CallLogNavigationInfoURI `json:"firstPage"`
	NextPage     CallLogNavigationInfoURI `json:"nextPage"`
	PreviousPage CallLogNavigationInfoURI `json:"previousPage"`
	LastPage     CallLogNavigationInfoURI `json:"lastPage"`
}

// CallLogPagingInfo Call Log Paging Info
type CallLogPagingInfo struct {
	Page          int `json:"page"`
	PerPage       int `json:"perPage"`
	PageStart     int `json:"pageStart"`
	PageEnd       int `json:"pageEnd"`
	TotalPages    int `json:"totalPages"`
	TotalElements int `json:"totalElements"`
}

// CallLogNavigationInfoURI Call Log Navigation Info URI
type CallLogNavigationInfoURI struct {
	Uri string `json:"uri"`
}

// CallLogRecordLegInfo Call Log Record Leg Info
type CallLogRecordLegInfo struct {
	Action             string               `json:"action"`
	Direction          string               `json:"direction"`
	Billing            BillingInfo          `json:"billing"`
	Delegate           DelegateInfo         `json:"@delegate"`
	ExtensionId        string               `json:"extensionId"`
	Duration           int                  `json:"duration"`
	Extension          ExtensionInfoCallLog `json:"extension"`
	LegType            string               `json:"legType"`
	StartTime          string               `json:"startTime"`
	Type               string               `json:"type"`
	Result             string               `json:"result"`
	Reason             string               `json:"reason"`
	ReasonDescription  string               `json:"reasonDescription"`
	From               CallLogCallerInfo    `json:"from"`
	To                 CallLogCallerInfo    `json:"to"`
	Transport          string               `json:"transport"`
	Recording          CallLogRecordingInfo `json:"recording"`
	ShortRecording     bool                 `json:"shortRecording"`
	Master             bool                 `json:"master"`
	Message            CallLogRecordMessage `json:"message"`
	TelephonySessionId string               `json:"telephonySessionId"`
}

// ExtensionInfoCallLog Extension Info Call Log
type ExtensionInfoCallLog struct {
	Id  int    `json:"id"`
	Uri string `json:"uri"`
}

// AccountCallLogSyncResponse Account Call Log Sync Response
type AccountCallLogSyncResponse struct {
	Uri      string                 `json:"uri"`
	Records  []CompanyCallLogRecord `json:"records"`
	SyncInfo CompanyCallLogSyncInfo `json:"syncInfo"`
}

// CompanyCallLogRecord Company Call Log Record
type CompanyCallLogRecord struct {
	Id                 string                 `json:"id"`
	Uri                string                 `json:"uri"`
	SessionId          string                 `json:"sessionId"`
	Extension          ExtensionInfoCallLog   `json:"extension"`
	TelephonySessionId string                 `json:"telephonySessionId"`
	Transport          string                 `json:"transport"`
	From               CallLogCallerInfo      `json:"from"`
	To                 CallLogCallerInfo      `json:"to"`
	Type               string                 `json:"type"`
	Direction          string                 `json:"direction"`
	Message            CallLogRecordMessage   `json:"message"`
	Delegate           DelegateInfo           `json:"@delegate"`
	Deleted            bool                   `json:"deleted"`
	Action             string                 `json:"action"`
	Result             string                 `json:"result"`
	Reason             string                 `json:"reason"`
	ReasonDescription  string                 `json:"reasonDescription"`
	StartTime          string                 `json:"startTime"`
	Duration           int                    `json:"duration"`
	Recording          CallLogRecordingInfo   `json:"recording"`
	ShortRecording     bool                   `json:"shortRecording"`
	Legs               []CallLogRecordLegInfo `json:"legs"`
	Billing            BillingInfo            `json:"billing"`
	LastModifiedTime   string                 `json:"lastModifiedTime"`
}

// CompanyCallLogSyncInfo Company Call Log Sync Info
type CompanyCallLogSyncInfo struct {
	SyncType  string `json:"syncType"`
	SyncToken string `json:"syncToken"`
	SyncTime  string `json:"syncTime"`
}

// CreateSipRegistrationResponse Create Sip Registration Response
type CreateSipRegistrationResponse struct {
	Device        SipRegistrationDeviceInfo `json:"device"`
	SipInfo       []SIPInfoResponse         `json:"sipInfo"`
	SipInfoPstn   []SIPInfoResponse         `json:"sipInfoPstn"`
	SipFlags      SIPFlagsResponse          `json:"sipFlags"`
	SipErrorCodes []string                  `json:"sipErrorCodes"`
}

// SipRegistrationDeviceInfo Sip Registration Device Info
type SipRegistrationDeviceInfo struct {
	Uri                     string                                `json:"uri"`
	Id                      string                                `json:"id"`
	Type                    string                                `json:"type"`
	Sku                     string                                `json:"sku"`
	Status                  string                                `json:"status"`
	Name                    string                                `json:"name"`
	Serial                  string                                `json:"serial"`
	ComputerName            string                                `json:"computerName"`
	Model                   DeviceModelInfo                       `json:"model"`
	Extension               DeviceExtensionInfo                   `json:"extension"`
	EmergencyServiceAddress DeviceEmergencyServiceAddressResource `json:"emergencyServiceAddress"`
	Emergency               SipRegistrationDeviceEmergencyInfo    `json:"emergency"`
	Shipping                Shipping                              `json:"shipping"`
	PhoneLines              []DevicePhoneLinesInfo                `json:"phoneLines"`
	BoxBillingId            int                                   `json:"boxBillingId"`
	UseAsCommonPhone        bool                                  `json:"useAsCommonPhone"`
	LinePooling             string                                `json:"linePooling"`
	InCompanyNet            bool                                  `json:"inCompanyNet"`
	Site                    DeviceSiteInfo                        `json:"site"`
	LastLocationReportTime  string                                `json:"lastLocationReportTime"`
}

// SipRegistrationDeviceEmergencyInfo Sip Registration Device Emergency Info
type SipRegistrationDeviceEmergencyInfo struct {
	Address               DeviceEmergencyServiceAddressResource `json:"address"`
	Location              SipRegistrationDeviceLocationInfo     `json:"location"`
	OutOfCountry          bool                                  `json:"outOfCountry"`
	AddressStatus         string                                `json:"addressStatus"`
	SyncStatus            string                                `json:"syncStatus"`
	AddressEditableStatus string                                `json:"addressEditableStatus"`
	AddressRequired       bool                                  `json:"addressRequired"`
	AddressLocationOnly   bool                                  `json:"addressLocationOnly"`
}

// SipRegistrationDeviceLocationInfo Sip Registration Device Location Info
type SipRegistrationDeviceLocationInfo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// DeviceSiteInfo Device Site Info
type DeviceSiteInfo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// DeviceModelInfo Device Model Info
type DeviceModelInfo struct {
	Id       string            `json:"id"`
	Name     string            `json:"name"`
	Addons   []DeviceAddonInfo `json:"addons"`
	Features []string          `json:"features"`
}

// DeviceAddonInfo Device Addon Info
type DeviceAddonInfo struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Count string `json:"count"`
}

// DeviceExtensionInfo Device Extension Info
type DeviceExtensionInfo struct {
	Id              int    `json:"id"`
	Uri             string `json:"uri"`
	ExtensionNumber string `json:"extensionNumber"`
}

// DeviceEmergencyServiceAddressResource Device Emergency Service Address Resource
type DeviceEmergencyServiceAddressResource struct {
	Street         string `json:"street"`
	Street2        string `json:"street2"`
	City           string `json:"city"`
	Zip            string `json:"zip"`
	CustomerName   string `json:"customerName"`
	State          string `json:"state"`
	StateId        string `json:"stateId"`
	StateIsoCode   string `json:"stateIsoCode"`
	StateName      string `json:"stateName"`
	CountryId      string `json:"countryId"`
	CountryIsoCode string `json:"countryIsoCode"`
	Country        string `json:"country"`
	CountryName    string `json:"countryName"`
	OutOfCountry   bool   `json:"outOfCountry"`
}

// SIPFlagsResponse SIPFlags Response
type SIPFlagsResponse struct {
	VoipFeatureEnabled   string `json:"voipFeatureEnabled"`
	VoipCountryBlocked   string `json:"voipCountryBlocked"`
	OutboundCallsEnabled string `json:"outboundCallsEnabled"`
	DscpEnabled          bool   `json:"dscpEnabled"`
	DscpSignaling        int    `json:"dscpSignaling"`
	DscpVoice            int    `json:"dscpVoice"`
	DscpVideo            int    `json:"dscpVideo"`
}

// SIPInfoResponse SIPInfo Response
type SIPInfoResponse struct {
	Username                string `json:"username"`
	Password                string `json:"password"`
	AuthorizationId         string `json:"authorizationId"`
	Domain                  string `json:"domain"`
	OutboundProxy           string `json:"outboundProxy"`
	OutboundProxyIPv6       string `json:"outboundProxyIPv6"`
	OutboundProxyBackup     string `json:"outboundProxyBackup"`
	OutboundProxyIPv6Backup string `json:"outboundProxyIPv6Backup"`
	Transport               string `json:"transport"`
	Certificate             string `json:"certificate"`
	SwitchBackInterval      int    `json:"switchBackInterval"`
}

// CreateSipRegistrationRequest Create Sip Registration Request
type CreateSipRegistrationRequest struct {
	Device  DeviceInfoRequest `json:"device"`
	SipInfo []SIPInfoRequest  `json:"sipInfo"`
}

// DeviceInfoRequest Device Info Request
type DeviceInfoRequest struct {
	Id            string `json:"id"`
	AppExternalId string `json:"appExternalId"`
	ComputerName  string `json:"computerName"`
	Serial        string `json:"serial"`
}

// SIPInfoRequest SIPInfo Request
type SIPInfoRequest struct {
	Transport string `json:"transport"`
}

// DevicePhoneLinesInfo Device Phone Lines Info
type DevicePhoneLinesInfo struct {
	Id               string                               `json:"id"`
	LineType         string                               `json:"lineType"`
	EmergencyAddress DevicePhoneLinesEmergencyAddressInfo `json:"emergencyAddress"`
	PhoneInfo        DevicePhoneNumberInfo                `json:"phoneInfo"`
}

// DevicePhoneLinesEmergencyAddressInfo Device Phone Lines Emergency Address Info
type DevicePhoneLinesEmergencyAddressInfo struct {
	Required  bool `json:"required"`
	LocalOnly bool `json:"localOnly"`
}

// DevicePhoneNumberInfo Device Phone Number Info
type DevicePhoneNumberInfo struct {
	Id          int                          `json:"id"`
	Country     DevicePhoneNumberCountryInfo `json:"country"`
	PaymentType string                       `json:"paymentType"`
	PhoneNumber string                       `json:"phoneNumber"`
	UsageType   string                       `json:"usageType"`
	Type        string                       `json:"type"`
}

// DevicePhoneNumberCountryInfo Device Phone Number Country Info
type DevicePhoneNumberCountryInfo struct {
	Id   string `json:"id"`
	Uri  string `json:"uri"`
	Name string `json:"name"`
}

// Shipping Shipping
type Shipping struct {
	Address        DeviceEmergencyServiceAddressResource `json:"address"`
	Method         MethodResource                        `json:"method"`
	Status         string                                `json:"status"`
	Carrier        string                                `json:"carrier"`
	TrackingNumber string                                `json:"trackingNumber"`
}

// MethodResource Method Resource
type MethodResource struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// CustomFieldsResource Custom Fields Resource
type CustomFieldsResource struct {
	Records []CustomFieldResource `json:"records"`
}

// CustomFieldResource Custom Field Resource
type CustomFieldResource struct {
	Id          string `json:"id"`
	Category    string `json:"category"`
	DisplayName string `json:"displayName"`
}

// CustomFieldCreateRequest Custom Field Create Request
type CustomFieldCreateRequest struct {
	Category    string `json:"category"`
	DisplayName string `json:"displayName"`
}

// CustomFieldUpdateRequest Custom Field Update Request
type CustomFieldUpdateRequest struct {
	DisplayName string `json:"displayName"`
}

// GetDeviceInfoResponse Get Device Info Response
type GetDeviceInfoResponse struct {
	Id                      string                          `json:"id"`
	Uri                     string                          `json:"uri"`
	Sku                     string                          `json:"sku"`
	Type                    string                          `json:"type"`
	Name                    string                          `json:"name"`
	Serial                  string                          `json:"serial"`
	Status                  string                          `json:"status"`
	ComputerName            string                          `json:"computerName"`
	Model                   ModelInfo                       `json:"model"`
	Extension               ExtensionInfoIntId              `json:"extension"`
	Emergency               DeviceEmergencyInfo             `json:"emergency"`
	EmergencyServiceAddress EmergencyServiceAddressResource `json:"emergencyServiceAddress"`
	PhoneLines              []PhoneLinesInfo                `json:"phoneLines"`
	Shipping                ShippingInfo                    `json:"shipping"`
	BoxBillingId            int                             `json:"boxBillingId"`
	UseAsCommonPhone        bool                            `json:"useAsCommonPhone"`
	InCompanyNet            bool                            `json:"inCompanyNet"`
	Site                    DeviceSiteInfo                  `json:"site"`
	LastLocationReportTime  string                          `json:"lastLocationReportTime"`
	LinePooling             string                          `json:"linePooling"`
	BillingStatement        BillingStatementInfo            `json:"billingStatement"`
}

// BillingStatementInfo Billing Statement Info
type BillingStatementInfo struct {
	Currency               string                    `json:"currency"`
	Charges                []BillingStatementCharges `json:"charges"`
	Fees                   []BillingStatementFees    `json:"fees"`
	TotalCharged           float64                   `json:"totalCharged"`
	TotalCharges           float64                   `json:"totalCharges"`
	TotalFees              float64                   `json:"totalFees"`
	Subtotal               float64                   `json:"subtotal"`
	TotalFreeServiceCredit float64                   `json:"totalFreeServiceCredit"`
}

// BillingStatementCharges Billing Statement Charges
type BillingStatementCharges struct {
	Description       string  `json:"description"`
	Amount            float64 `json:"amount"`
	Feature           string  `json:"feature"`
	FreeServiceCredit float64 `json:"freeServiceCredit"`
}

// BillingStatementFees Billing Statement Fees
type BillingStatementFees struct {
	Description       string  `json:"description"`
	Amount            float64 `json:"amount"`
	FreeServiceCredit float64 `json:"freeServiceCredit"`
}

// ModelInfo Model Info
type ModelInfo struct {
	Id        string      `json:"id"`
	Name      string      `json:"name"`
	Addons    []AddonInfo `json:"addons"`
	Features  []string    `json:"features"`
	LineCount int         `json:"lineCount"`
}

// AddonInfo Addon Info
type AddonInfo struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

// EmergencyAddress Emergency Address
type EmergencyAddress struct {
	Required  bool `json:"required"`
	LocalOnly bool `json:"localOnly"`
}

// DeviceEmergencyAddress Device Emergency Address
type DeviceEmergencyAddress struct {
	CustomerName   string `json:"customerName"`
	Street         string `json:"street"`
	Street2        string `json:"street2"`
	City           string `json:"city"`
	Zip            string `json:"zip"`
	State          string `json:"state"`
	StateId        string `json:"stateId"`
	StateIsoCode   string `json:"stateIsoCode"`
	StateName      string `json:"stateName"`
	CountryId      string `json:"countryId"`
	CountryIsoCode string `json:"countryIsoCode"`
	Country        string `json:"country"`
	CountryName    string `json:"countryName"`
}

// PhoneLinesInfo Phone Lines Info
type PhoneLinesInfo struct {
	Id               string               `json:"id"`
	LineType         string               `json:"lineType"`
	PhoneInfo        PhoneNumberInfoIntId `json:"phoneInfo"`
	EmergencyAddress EmergencyAddress     `json:"emergencyAddress"`
}

// ShippingInfo Shipping Info
type ShippingInfo struct {
	Status         string              `json:"status"`
	Carrier        string              `json:"carrier"`
	TrackingNumber string              `json:"trackingNumber"`
	Method         MethodInfo          `json:"method"`
	Address        ShippingAddressInfo `json:"address"`
}

// ShippingAddressInfo Shipping Address Info
type ShippingAddressInfo struct {
	CustomerName            string `json:"customerName"`
	AdditionalCustomerName  string `json:"additionalCustomerName"`
	CustomerEmail           string `json:"customerEmail"`
	AdditionalCustomerEmail string `json:"additionalCustomerEmail"`
	CustomerPhone           string `json:"customerPhone"`
	AdditionalCustomerPhone string `json:"additionalCustomerPhone"`
	Street                  string `json:"street"`
	Street2                 string `json:"street2"`
	City                    string `json:"city"`
	State                   string `json:"state"`
	StateId                 string `json:"stateId"`
	StateIsoCode            string `json:"stateIsoCode"`
	StateName               string `json:"stateName"`
	CountryId               string `json:"countryId"`
	CountryIsoCode          string `json:"countryIsoCode"`
	Country                 string `json:"country"`
	CountryName             string `json:"countryName"`
	Zip                     string `json:"zip"`
	TaxId                   string `json:"taxId"`
}

// MethodInfo Method Info
type MethodInfo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// GetExtensionDevicesResponse Get Extension Devices Response
type GetExtensionDevicesResponse struct {
	Uri        string                           `json:"uri"`
	Records    []ExtensionDeviceResponse        `json:"records"`
	Navigation DeviceProvisioningNavigationInfo `json:"navigation"`
	Paging     DeviceProvisioningPagingInfo     `json:"paging"`
}

// ExtensionDeviceResponse Extension Device Response
type ExtensionDeviceResponse struct {
	Id                      string                          `json:"id"`
	Uri                     string                          `json:"uri"`
	Sku                     string                          `json:"sku"`
	Type                    string                          `json:"type"`
	Name                    string                          `json:"name"`
	Status                  string                          `json:"status"`
	Serial                  string                          `json:"serial"`
	ComputerName            string                          `json:"computerName"`
	Model                   ModelInfo                       `json:"model"`
	Extension               ExtensionInfoIntId              `json:"extension"`
	EmergencyServiceAddress EmergencyServiceAddressResource `json:"emergencyServiceAddress"`
	Emergency               DeviceEmergencyInfo             `json:"emergency"`
	PhoneLines              []PhoneLinesInfo                `json:"phoneLines"`
	Shipping                ShippingInfo                    `json:"shipping"`
	BoxBillingId            int                             `json:"boxBillingId"`
	UseAsCommonPhone        bool                            `json:"useAsCommonPhone"`
	LinePooling             string                          `json:"linePooling"`
	InCompanyNet            bool                            `json:"inCompanyNet"`
	Site                    DeviceSiteInfo                  `json:"site"`
	LastLocationReportTime  string                          `json:"lastLocationReportTime"`
}

// DeviceProvisioningNavigationInfo Device Provisioning Navigation Info
type DeviceProvisioningNavigationInfo struct {
	FirstPage    DeviceProvisioningNavigationInfoUri `json:"firstPage"`
	NextPage     DeviceProvisioningNavigationInfoUri `json:"nextPage"`
	PreviousPage DeviceProvisioningNavigationInfoUri `json:"previousPage"`
	LastPage     DeviceProvisioningNavigationInfoUri `json:"lastPage"`
}

// DeviceProvisioningPagingInfo Device Provisioning Paging Info
type DeviceProvisioningPagingInfo struct {
	Page          int `json:"page"`
	PerPage       int `json:"perPage"`
	PageStart     int `json:"pageStart"`
	PageEnd       int `json:"pageEnd"`
	TotalPages    int `json:"totalPages"`
	TotalElements int `json:"totalElements"`
}

// DeviceProvisioningNavigationInfoUri Device Provisioning Navigation Info Uri
type DeviceProvisioningNavigationInfoUri struct {
	Uri string `json:"uri"`
}

// PhoneNumberInfoIntId Phone Number Info Int Id
type PhoneNumberInfoIntId struct {
	Id          int                             `json:"id"`
	Country     PhoneNumberCountryInfo          `json:"country"`
	Extension   DeviceProvisioningExtensionInfo `json:"extension"`
	Label       string                          `json:"label"`
	Location    string                          `json:"location"`
	PaymentType string                          `json:"paymentType"`
	PhoneNumber string                          `json:"phoneNumber"`
	Status      string                          `json:"status"`
	Type        string                          `json:"type"`
	UsageType   string                          `json:"usageType"`
}

// DeviceProvisioningExtensionInfo Device Provisioning Extension Info
type DeviceProvisioningExtensionInfo struct {
	Id              string `json:"id"`
	Uri             string `json:"uri"`
	ExtensionNumber string `json:"extensionNumber"`
	PartnerId       string `json:"partnerId"`
}

// ExtensionInfoIntId Extension Info Int Id
type ExtensionInfoIntId struct {
	Id              int    `json:"id"`
	Uri             string `json:"uri"`
	ExtensionNumber string `json:"extensionNumber"`
	PartnerId       string `json:"partnerId"`
}

// PhoneNumberCountryInfo Phone Number Country Info
type PhoneNumberCountryInfo struct {
	Id   string `json:"id"`
	Uri  string `json:"uri"`
	Name string `json:"name"`
}

// EmergencyServiceAddressResourceRequest Emergency Service Address Resource Request
type EmergencyServiceAddressResourceRequest struct {
	Street       string `json:"street"`
	Street2      string `json:"street2"`
	City         string `json:"city"`
	Zip          string `json:"zip"`
	CustomerName string `json:"customerName"`
	State        string `json:"state"`
	StateId      string `json:"stateId"`
	Country      string `json:"country"`
	CountryId    string `json:"countryId"`
}

// EmergencyServiceAddressResource Emergency Service Address Resource
type EmergencyServiceAddressResource struct {
	Street                  string `json:"street"`
	Street2                 string `json:"street2"`
	City                    string `json:"city"`
	Zip                     string `json:"zip"`
	CustomerName            string `json:"customerName"`
	State                   string `json:"state"`
	StateId                 string `json:"stateId"`
	StateIsoCode            string `json:"stateIsoCode"`
	StateName               string `json:"stateName"`
	CountryId               string `json:"countryId"`
	CountryIsoCode          string `json:"countryIsoCode"`
	Country                 string `json:"country"`
	CountryName             string `json:"countryName"`
	OutOfCountry            bool   `json:"outOfCountry"`
	SyncStatus              string `json:"syncStatus"`
	AdditionalCustomerName  string `json:"additionalCustomerName"`
	CustomerEmail           string `json:"customerEmail"`
	AdditionalCustomerEmail string `json:"additionalCustomerEmail"`
	CustomerPhone           string `json:"customerPhone"`
	AdditionalCustomerPhone string `json:"additionalCustomerPhone"`
	TaxId                   string `json:"taxId"`
}

// PhoneNumberExtensionInfo Phone Number Extension Info
type PhoneNumberExtensionInfo struct {
	Uri              string                       `json:"uri"`
	Id               string                       `json:"id"`
	PartnerId        string                       `json:"partnerId"`
	ExtensionNumber  string                       `json:"extensionNumber"`
	LoginName        string                       `json:"loginName"`
	Contact          ExtensionContactInfo         `json:"contact"`
	References       []Reference                  `json:"references"`
	Name             string                       `json:"name"`
	Type             string                       `json:"type"`
	Status           string                       `json:"status"`
	StatusInfo       StatusInfo                   `json:"statusInfo"`
	Departments      []DepartmentResource         `json:"departments"`
	ServiceFeatures  []ServiceFeatureValue        `json:"serviceFeatures"`
	RegionalSettings RegionalSettingsInfo         `json:"regionalSettings"`
	SetupWizardState string                       `json:"setupWizardState"`
	Permissions      ExtensionPermissionsResource `json:"permissions"`
	Password         string                       `json:"password"`
	IvrPin           string                       `json:"ivrPin"`
	ProfileImage     ProfileImageResource         `json:"profileImage"`
}

// ProfileImageResource Profile Image Resource
type ProfileImageResource struct {
	Uri          string                       `json:"uri"`
	Etag         string                       `json:"etag"`
	ContentType  string                       `json:"contentType"`
	LastModified string                       `json:"lastModified"`
	Scales       []ScaledProfileImageResource `json:"scales"`
}

// ExtensionPermissionsResource Extension Permissions Resource
type ExtensionPermissionsResource struct {
	Uri                      string     `json:"uri"`
	Admin                    Permission `json:"admin"`
	InternationalCalling     Permission `json:"internationalCalling"`
	FreeSoftPhoneDigitalLine Permission `json:"freeSoftPhoneDigitalLine"`
}

// RegionalSettingsInfo Regional Settings Info
type RegionalSettingsInfo struct {
	Timezone         TimezoneResource `json:"timezone"`
	HomeCountry      CountryResource  `json:"homeCountry"`
	Language         LanguageResource `json:"language"`
	GreetingLanguage LanguageResource `json:"greetingLanguage"`
	FormattingLocale LanguageResource `json:"formattingLocale"`
}

// StatusInfo Status Info
type StatusInfo struct {
	Reason  string `json:"reason"`
	Till    string `json:"till"`
	Comment string `json:"comment"`
}

// ExtensionContactInfo Extension Contact Info
type ExtensionContactInfo struct {
	FirstName       string                    `json:"firstName"`
	LastName        string                    `json:"lastName"`
	Company         string                    `json:"company"`
	Email           string                    `json:"email"`
	BusinessPhone   string                    `json:"businessPhone"`
	BusinessAddress ContactAddressInfoDevices `json:"businessAddress"`
}

// ContactAddressInfoDevices Contact Address Info Devices
type ContactAddressInfoDevices struct {
	Country string `json:"country"`
	State   string `json:"state"`
	City    string `json:"city"`
	Street  string `json:"street"`
	Zip     string `json:"zip"`
}

// LanguageResource Language Resource
type LanguageResource struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	LocaleCode string `json:"localeCode"`
}

// CountryResource Country Resource
type CountryResource struct {
	Uri              string `json:"uri"`
	Id               string `json:"id"`
	Name             string `json:"name"`
	IsoCode          string `json:"isoCode"`
	CallingCode      string `json:"callingCode"`
	EmergencyCalling bool   `json:"emergencyCalling"`
	NumberSelling    bool   `json:"numberSelling"`
	LoginAllowed     bool   `json:"loginAllowed"`
}

// TimezoneResource Timezone Resource
type TimezoneResource struct {
	Uri         string `json:"uri"`
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Permission Permission
type Permission struct {
	Enabled bool `json:"enabled"`
}

// Reference Reference
type Reference struct {
	Type string `json:"type"`
	Ref  string `json:"@ref"`
}

// ScaledProfileImageResource Scaled Profile Image Resource
type ScaledProfileImageResource struct {
	Uri string `json:"uri"`
}

// DepartmentResource Department Resource
type DepartmentResource struct {
	Uri             string `json:"uri"`
	Id              string `json:"id"`
	ExtensionNumber string `json:"extensionNumber"`
}

// ServiceFeatureValue Service Feature Value
type ServiceFeatureValue struct {
	FeatureName string `json:"featureName"`
	Enabled     bool   `json:"enabled"`
	Reason      string `json:"reason"`
}

// AccountDeviceUpdate Account Device Update
type AccountDeviceUpdate struct {
	EmergencyServiceAddress EmergencyServiceAddressResourceRequest `json:"emergencyServiceAddress"`
	Emergency               DeviceEmergencyInfo                    `json:"emergency"`
	Extension               DeviceUpdateExtensionInfo              `json:"extension"`
	PhoneLines              DeviceUpdatePhoneLinesInfo             `json:"phoneLines"`
	UseAsCommonPhone        bool                                   `json:"useAsCommonPhone"`
}

// DeviceEmergencyInfo Device Emergency Info
type DeviceEmergencyInfo struct {
	Address               DeviceEmergencyAddress      `json:"address"`
	Location              DeviceEmergencyLocationInfo `json:"location"`
	OutOfCountry          bool                        `json:"outOfCountry"`
	AddressStatus         string                      `json:"addressStatus"`
	SyncStatus            string                      `json:"syncStatus"`
	AddressEditableStatus string                      `json:"addressEditableStatus"`
}

// DeviceEmergencyLocationInfo Device Emergency Location Info
type DeviceEmergencyLocationInfo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// DeviceUpdateExtensionInfo Device Update Extension Info
type DeviceUpdateExtensionInfo struct {
	Id string `json:"id"`
}

// DeviceUpdatePhoneLinesInfo Device Update Phone Lines Info
type DeviceUpdatePhoneLinesInfo struct {
	PhoneLines []UpdateDevicePhoneInfo `json:"phoneLines"`
}

// UpdateDevicePhoneInfo Update Device Phone Info
type UpdateDevicePhoneInfo struct {
	Id string `json:"id"`
}

// ListMeetingRecordingsResponse List Meeting Recordings Response
type ListMeetingRecordingsResponse struct {
	Records    []MeetingRecording              `json:"records"`
	Paging     MeetingRecordingsPagingInfo     `json:"paging"`
	Navigation MeetingRecordingsNavigationInfo `json:"navigation"`
}

// MeetingRecording Meeting Recording
type MeetingRecording struct {
	Meeting   MeetingInfo            `json:"meeting"`
	Recording []MeetingRecordingInfo `json:"recording"`
}

// MeetingInfo Meeting Info
type MeetingInfo struct {
	Id        string `json:"id"`
	Topic     string `json:"topic"`
	StartTime string `json:"startTime"`
}

// MeetingRecordingInfo Meeting Recording Info
type MeetingRecordingInfo struct {
	Id                 string `json:"id"`
	ContentDownloadUri string `json:"contentDownloadUri"`
	ContentType        string `json:"contentType"`
	Size               int    `json:"size"`
	StartTime          string `json:"startTime"`
	EndTime            string `json:"endTime"`
	Status             string `json:"status"`
}

// MeetingRecordingsNavigationInfo Meeting Recordings Navigation Info
type MeetingRecordingsNavigationInfo struct {
	FirstPage    MeetingRecordingsNavigationInfoUri `json:"firstPage"`
	NextPage     MeetingRecordingsNavigationInfoUri `json:"nextPage"`
	PreviousPage MeetingRecordingsNavigationInfoUri `json:"previousPage"`
	LastPage     MeetingRecordingsNavigationInfoUri `json:"lastPage"`
}

// MeetingRecordingsNavigationInfoUri Meeting Recordings Navigation Info Uri
type MeetingRecordingsNavigationInfoUri struct {
	Uri string `json:"uri"`
}

// MeetingRecordingsPagingInfo Meeting Recordings Paging Info
type MeetingRecordingsPagingInfo struct {
	Page          int `json:"page"`
	PerPage       int `json:"perPage"`
	PageStart     int `json:"pageStart"`
	PageEnd       int `json:"pageEnd"`
	TotalPages    int `json:"totalPages"`
	TotalElements int `json:"totalElements"`
}

// AccountLockedSettingResponse Account Locked Setting Response
type AccountLockedSettingResponse struct {
	ScheduleMeeting ScheduleUserMeetingInfo     `json:"scheduleMeeting"`
	Recording       UserMeetingRecordingSetting `json:"recording"`
}

// PublicMeetingInvitationResponse Public Meeting Invitation Response
type PublicMeetingInvitationResponse struct {
	Invitation string `json:"invitation"`
}

// MeetingsResource Meetings Resource
type MeetingsResource struct {
	Uri        string                    `json:"uri"`
	Records    []MeetingResponseResource `json:"records"`
	Paging     MeetingsPagingInfo        `json:"paging"`
	Navigation MeetingsNavigationInfo    `json:"navigation"`
}

// MeetingResponseResource Meeting Response Resource
type MeetingResponseResource struct {
	Uri                     string                  `json:"uri"`
	Uuid                    string                  `json:"uuid"`
	Id                      string                  `json:"id"`
	Topic                   string                  `json:"topic"`
	MeetingType             string                  `json:"meetingType"`
	Password                string                  `json:"password"`
	H323Password            string                  `json:"h323Password"`
	Status                  string                  `json:"status"`
	Links                   MeetingLinks            `json:"links"`
	Schedule                MeetingScheduleResource `json:"schedule"`
	Host                    HostInfoRequest         `json:"host"`
	AllowJoinBeforeHost     bool                    `json:"allowJoinBeforeHost"`
	StartHostVideo          bool                    `json:"startHostVideo"`
	StartParticipantsVideo  bool                    `json:"startParticipantsVideo"`
	AudioOptions            []string                `json:"audioOptions"`
	Reccurence              RecurrenceInfo          `json:"reccurence"`
	AutoRecordType          string                  `json:"autoRecordType"`
	EnforceLogin            bool                    `json:"enforceLogin"`
	MuteParticipantsOnEntry bool                    `json:"muteParticipantsOnEntry"`
	EnableWaitingRoom       bool                    `json:"enableWaitingRoom"`
	GlobalDialInCountries   []string                `json:"globalDialInCountries"`
}

// MeetingLinks Meeting Links
type MeetingLinks struct {
	StartUri string `json:"startUri"`
	JoinUri  string `json:"joinUri"`
}

// MeetingScheduleResource Meeting Schedule Resource
type MeetingScheduleResource struct {
	StartTime         string           `json:"startTime"`
	DurationInMinutes int              `json:"durationInMinutes"`
	TimeZone          TimezoneResource `json:"timeZone"`
}

// MeetingsNavigationInfo Meetings Navigation Info
type MeetingsNavigationInfo struct {
	NextPage     MeetingsNavigationInfoUri `json:"nextPage"`
	PreviousPage MeetingsNavigationInfoUri `json:"previousPage"`
	FirstPage    MeetingsNavigationInfoUri `json:"firstPage"`
	LastPage     MeetingsNavigationInfoUri `json:"lastPage"`
}

// MeetingsNavigationInfoUri Meetings Navigation Info Uri
type MeetingsNavigationInfoUri struct {
	Uri string `json:"uri"`
}

// MeetingRequestResource Meeting Request Resource
type MeetingRequestResource struct {
	Topic                   string                  `json:"topic"`
	MeetingType             string                  `json:"meetingType"`
	Schedule                MeetingScheduleResource `json:"schedule"`
	Password                string                  `json:"password"`
	Host                    HostInfoRequest         `json:"host"`
	AllowJoinBeforeHost     bool                    `json:"allowJoinBeforeHost"`
	StartHostVideo          bool                    `json:"startHostVideo"`
	StartParticipantsVideo  bool                    `json:"startParticipantsVideo"`
	UsePersonalMeetingId    bool                    `json:"usePersonalMeetingId"`
	AudioOptions            []string                `json:"audioOptions"`
	Recurrence              RecurrenceInfo          `json:"recurrence"`
	AutoRecordType          string                  `json:"autoRecordType"`
	EnforceLogin            bool                    `json:"enforceLogin"`
	MuteParticipantsOnEntry bool                    `json:"muteParticipantsOnEntry"`
	EnableWaitingRoom       bool                    `json:"enableWaitingRoom"`
	GlobalDialInCountries   []string                `json:"globalDialInCountries"`
}

// RecurrenceInfo Recurrence Info
type RecurrenceInfo struct {
	Frequency        string `json:"frequency"`
	Interval         int    `json:"interval"`
	MonthlyByWeek    string `json:"monthlyByWeek"`
	WeeklyByDay      string `json:"weeklyByDay"`
	WeeklyByDays     string `json:"weeklyByDays"`
	MonthlyByDay     int    `json:"monthlyByDay"`
	MonthlyByWeekDay int    `json:"monthlyByWeekDay"`
	Count            int    `json:"count"`
	Until            string `json:"until"`
}

// HostInfoRequest Host Info Request
type HostInfoRequest struct {
	Uri string `json:"uri"`
	Id  string `json:"id"`
}

// MeetingServiceInfoResource Meeting Service Info Resource
type MeetingServiceInfoResource struct {
	Uri                  string                          `json:"uri"`
	SupportUri           string                          `json:"supportUri"`
	IntlDialInNumbersUri string                          `json:"intlDialInNumbersUri"`
	ExternalUserInfo     MeetingExternalUserInfoResource `json:"externalUserInfo"`
	DialInNumbers        []DialInNumberResource          `json:"dialInNumbers"`
}

// MeetingExternalUserInfoResource Meeting External User Info Resource
type MeetingExternalUserInfoResource struct {
	Uri                      string `json:"uri"`
	UserId                   string `json:"userId"`
	AccountId                string `json:"accountId"`
	UserType                 int    `json:"userType"`
	UserToken                string `json:"userToken"`
	HostKey                  string `json:"hostKey"`
	PersonalMeetingId        string `json:"personalMeetingId"`
	PersonalLink             string `json:"personalLink"`
	UsePmiForInstantMeetings bool   `json:"usePmiForInstantMeetings"`
}

// DialInNumberResource Dial In Number Resource
type DialInNumberResource struct {
	PhoneNumber     string          `json:"phoneNumber"`
	FormattedNumber string          `json:"formattedNumber"`
	Location        string          `json:"location"`
	Country         CountryResource `json:"country"`
}

// MeetingsPagingInfo Meetings Paging Info
type MeetingsPagingInfo struct {
	Page          int `json:"page"`
	TotalPages    int `json:"totalPages"`
	PerPage       int `json:"perPage"`
	TotalElements int `json:"totalElements"`
	PageStart     int `json:"pageStart"`
	PageEnd       int `json:"pageEnd"`
}

// AssistantsResource Assistants Resource
type AssistantsResource struct {
	Records []AssistantResource `json:"records"`
}

// AssistantResource Assistant Resource
type AssistantResource struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// AssistedUsersResource Assisted Users Resource
type AssistedUsersResource struct {
	Records []AssistedUserResource `json:"records"`
}

// AssistedUserResource Assisted User Resource
type AssistedUserResource struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// MeetingServiceInfoRequest Meeting Service Info Request
type MeetingServiceInfoRequest struct {
	ExternalUserInfo MeetingExternalUserInfoResource `json:"externalUserInfo"`
}

// MeetingUserSettingsResponse Meeting User Settings Response
type MeetingUserSettingsResponse struct {
	Recording       UserMeetingRecordingSetting `json:"recording"`
	ScheduleMeeting ScheduleUserMeetingInfo     `json:"scheduleMeeting"`
}

// ScheduleUserMeetingInfo Schedule User Meeting Info
type ScheduleUserMeetingInfo struct {
	StartHostVideo                          bool     `json:"startHostVideo"`
	StartParticipantsVideo                  bool     `json:"startParticipantsVideo"`
	AudioOptions                            []string `json:"audioOptions"`
	AllowJoinBeforeHost                     bool     `json:"allowJoinBeforeHost"`
	UsePmiForScheduledMeetings              bool     `json:"usePmiForScheduledMeetings"`
	UsePmiForInstantMeetings                bool     `json:"usePmiForInstantMeetings"`
	RequirePasswordForSchedulingNewMeetings bool     `json:"requirePasswordForSchedulingNewMeetings"`
	RequirePasswordForScheduledMeetings     bool     `json:"requirePasswordForScheduledMeetings"`
	DefaultPasswordForScheduledMeetings     string   `json:"defaultPasswordForScheduledMeetings"`
	RequirePasswordForInstantMeetings       bool     `json:"requirePasswordForInstantMeetings"`
	RequirePasswordForPmiMeetings           string   `json:"requirePasswordForPmiMeetings"`
	PmiPassword                             string   `json:"pmiPassword"`
	PstnPasswordProtected                   bool     `json:"pstnPasswordProtected"`
}

// UserMeetingRecordingSetting User Meeting Recording Setting
type UserMeetingRecordingSetting struct {
	LocalRecording    bool   `json:"localRecording"`
	CloudRecording    bool   `json:"cloudRecording"`
	RecordSpeakerView bool   `json:"recordSpeakerView"`
	RecordGalleryView bool   `json:"recordGalleryView"`
	RecordAudioFile   bool   `json:"recordAudioFile"`
	SaveChatText      bool   `json:"saveChatText"`
	ShowTimestamp     bool   `json:"showTimestamp"`
	AutoRecording     string `json:"autoRecording"`
	AutoDeleteCmr     string `json:"autoDeleteCmr"`
	AutoDeleteCmrDays int    `json:"autoDeleteCmrDays"`
}

// GetMessageInfoResponse Get Message Info Response
type GetMessageInfoResponse struct {
	Id                      int                                `json:"id"`
	Uri                     string                             `json:"uri"`
	Attachments             []MessageAttachmentInfo            `json:"attachments"`
	Availability            string                             `json:"availability"`
	ConversationId          int                                `json:"conversationId"`
	Conversation            ConversationInfo                   `json:"conversation"`
	CreationTime            string                             `json:"creationTime"`
	DeliveryErrorCode       string                             `json:"deliveryErrorCode"`
	Direction               string                             `json:"direction"`
	FaxPageCount            int                                `json:"faxPageCount"`
	FaxResolution           string                             `json:"faxResolution"`
	From                    MessageStoreCallerInfoResponseFrom `json:"from"`
	LastModifiedTime        string                             `json:"lastModifiedTime"`
	MessageStatus           string                             `json:"messageStatus"`
	PgToDepartment          bool                               `json:"pgToDepartment"`
	Priority                string                             `json:"priority"`
	ReadStatus              string                             `json:"readStatus"`
	SmsDeliveryTime         string                             `json:"smsDeliveryTime"`
	SmsSendingAttemptsCount int                                `json:"smsSendingAttemptsCount"`
	Subject                 string                             `json:"subject"`
	To                      []MessageStoreCallerInfoResponseTo `json:"to"`
	Type                    string                             `json:"type"`
	VmTranscriptionStatus   string                             `json:"vmTranscriptionStatus"`
	CoverIndex              int                                `json:"coverIndex"`
	CoverPageText           string                             `json:"coverPageText"`
}

// GetInternalTextMessageInfoResponse Get Internal Text Message Info Response
type GetInternalTextMessageInfoResponse struct {
	Id               int                                `json:"id"`
	Uri              string                             `json:"uri"`
	Attachments      []MessageAttachmentInfo            `json:"attachments"`
	Availability     string                             `json:"availability"`
	ConversationId   int                                `json:"conversationId"`
	Conversation     ConversationInfo                   `json:"conversation"`
	CreationTime     string                             `json:"creationTime"`
	Direction        string                             `json:"direction"`
	From             MessageStoreCallerInfoResponseFrom `json:"from"`
	LastModifiedTime string                             `json:"lastModifiedTime"`
	MessageStatus    string                             `json:"messageStatus"`
	PgToDepartment   bool                               `json:"pgToDepartment"`
	Priority         string                             `json:"priority"`
	ReadStatus       string                             `json:"readStatus"`
	Subject          string                             `json:"subject"`
	To               []MessageStoreCallerInfoResponseTo `json:"to"`
	Type             string                             `json:"type"`
}

// GetSMSMessageInfoResponse Get SMSMessage Info Response
type GetSMSMessageInfoResponse struct {
	Id                      int                                `json:"id"`
	Uri                     string                             `json:"uri"`
	Attachments             []MessageAttachmentInfo            `json:"attachments"`
	Availability            string                             `json:"availability"`
	ConversationId          int                                `json:"conversationId"`
	Conversation            ConversationInfo                   `json:"conversation"`
	CreationTime            string                             `json:"creationTime"`
	DeliveryErrorCode       string                             `json:"deliveryErrorCode"`
	Direction               string                             `json:"direction"`
	From                    MessageStoreCallerInfoResponseFrom `json:"from"`
	LastModifiedTime        string                             `json:"lastModifiedTime"`
	MessageStatus           string                             `json:"messageStatus"`
	Priority                string                             `json:"priority"`
	ReadStatus              string                             `json:"readStatus"`
	SmsDeliveryTime         string                             `json:"smsDeliveryTime"`
	SmsSendingAttemptsCount int                                `json:"smsSendingAttemptsCount"`
	Subject                 string                             `json:"subject"`
	To                      []MessageStoreCallerInfoResponseTo `json:"to"`
	Type                    string                             `json:"type"`
}

// GetMessageInfoMultiResponse Get Message Info Multi Response
type GetMessageInfoMultiResponse struct {
	ResourceId string      `json:"resourceId"`
	Status     int         `json:"status"`
	Body       MessageBody `json:"body"`
}

// MessageBody Message Body
type MessageBody struct {
	Uri                     string                  `json:"uri"`
	Id                      string                  `json:"id"`
	Attachments             []MessageAttachmentInfo `json:"attachments"`
	Availability            string                  `json:"availability"`
	ConversationId          int                     `json:"conversationId"`
	Conversation            ConversationInfo        `json:"conversation"`
	CreationTime            string                  `json:"creationTime"`
	DeliveryErrorCode       string                  `json:"deliveryErrorCode"`
	Direction               string                  `json:"direction"`
	FaxPageCount            int                     `json:"faxPageCount"`
	FaxResolution           string                  `json:"faxResolution"`
	From                    MessageSenderInfo       `json:"from"`
	LastModifiedTime        string                  `json:"lastModifiedTime"`
	MessageStatus           string                  `json:"messageStatus"`
	PgToDepartment          bool                    `json:"pgToDepartment"`
	Priority                string                  `json:"priority"`
	ReadStatus              string                  `json:"readStatus"`
	SmsDeliveryTime         string                  `json:"smsDeliveryTime"`
	SmsSendingAttemptsCount int                     `json:"smsSendingAttemptsCount"`
	Subject                 string                  `json:"subject"`
	To                      []MessageRecipientInfo  `json:"to"`
	Type                    string                  `json:"type"`
	VmTranscriptionStatus   string                  `json:"vmTranscriptionStatus"`
}

// MessageRecipientInfo Message Recipient Info
type MessageRecipientInfo struct {
	ExtensionNumber string `json:"extensionNumber"`
	ExtensionId     string `json:"extensionId"`
	Name            string `json:"name"`
}

// MessageSenderInfo Message Sender Info
type MessageSenderInfo struct {
	ExtensionNumber string `json:"extensionNumber"`
	ExtensionId     string `json:"extensionId"`
	Name            string `json:"name"`
}

// MessageStoreCallerInfoResponseTo Message Store Caller Info Response To
type MessageStoreCallerInfoResponseTo struct {
	ExtensionNumber string `json:"extensionNumber"`
	ExtensionId     string `json:"extensionId"`
	Location        string `json:"location"`
	Target          bool   `json:"target"`
	MessageStatus   string `json:"messageStatus"`
	FaxErrorCode    string `json:"faxErrorCode"`
	Name            string `json:"name"`
	PhoneNumber     string `json:"phoneNumber"`
}

// MessageStoreCallerInfoResponseFrom Message Store Caller Info Response From
type MessageStoreCallerInfoResponseFrom struct {
	ExtensionNumber string `json:"extensionNumber"`
	ExtensionId     string `json:"extensionId"`
	Location        string `json:"location"`
	Name            string `json:"name"`
	PhoneNumber     string `json:"phoneNumber"`
}

// MessageAttachmentInfo Message Attachment Info
type MessageAttachmentInfo struct {
	Id          int    `json:"id"`
	Uri         string `json:"uri"`
	Type        string `json:"type"`
	ContentType string `json:"contentType"`
	VmDuration  int    `json:"vmDuration"`
	FileName    string `json:"fileName"`
	Size        int    `json:"size"`
	Height      int    `json:"height"`
	Width       int    `json:"width"`
}

// MessageAttachmentInfoIntId Message Attachment Info Int Id
type MessageAttachmentInfoIntId struct {
	Id          int    `json:"id"`
	Uri         string `json:"uri"`
	Type        string `json:"type"`
	ContentType string `json:"contentType"`
	VmDuration  int    `json:"vmDuration"`
	Filename    string `json:"filename"`
	Size        int    `json:"size"`
}

// CreateInternalTextMessageRequest Create Internal Text Message Request
type CreateInternalTextMessageRequest struct {
	From    PagerCallerInfoRequest   `json:"from"`
	ReplyOn int                      `json:"replyOn"`
	Text    string                   `json:"text"`
	To      []PagerCallerInfoRequest `json:"to"`
}

// PagerCallerInfoRequest Pager Caller Info Request
type PagerCallerInfoRequest struct {
	ExtensionId     string `json:"extensionId"`
	ExtensionNumber string `json:"extensionNumber"`
}

// MessageStoreCallerInfoRequest Message Store Caller Info Request
type MessageStoreCallerInfoRequest struct {
	PhoneNumber string `json:"phoneNumber"`
}

// UpdateMessageRequest Update Message Request
type UpdateMessageRequest struct {
	ReadStatus string `json:"readStatus"`
}

// FaxResponse Fax Response
type FaxResponse struct {
	Id               int                          `json:"id"`
	Uri              string                       `json:"uri"`
	Type             string                       `json:"type"`
	From             CallerInfoFrom               `json:"from"`
	To               []CallerInfoTo               `json:"to"`
	CreationTime     string                       `json:"creationTime"`
	ReadStatus       string                       `json:"readStatus"`
	Priority         string                       `json:"priority"`
	Attachments      []MessageAttachmentInfoIntId `json:"attachments"`
	Direction        string                       `json:"direction"`
	Availability     string                       `json:"availability"`
	MessageStatus    string                       `json:"messageStatus"`
	FaxResolution    string                       `json:"faxResolution"`
	FaxPageCount     int                          `json:"faxPageCount"`
	LastModifiedTime string                       `json:"lastModifiedTime"`
	CoverIndex       int                          `json:"coverIndex"`
	CoverPageText    string                       `json:"coverPageText"`
}

// GetMessageSyncResponse Get Message Sync Response
type GetMessageSyncResponse struct {
	Uri      string                   `json:"uri"`
	Records  []GetMessageInfoResponse `json:"records"`
	SyncInfo SyncInfoMessages         `json:"syncInfo"`
}

// SyncInfoMessages Sync Info Messages
type SyncInfoMessages struct {
	SyncType          string `json:"syncType"`
	SyncToken         string `json:"syncToken"`
	SyncTime          string `json:"syncTime"`
	OlderRecordsExist bool   `json:"olderRecordsExist"`
}

// CallerInfoFrom Caller Info From
type CallerInfoFrom struct {
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
	Location    string `json:"location"`
}

// CallerInfoTo Caller Info To
type CallerInfoTo struct {
	PhoneNumber   string `json:"phoneNumber"`
	Name          string `json:"name"`
	Location      string `json:"location"`
	MessageStatus string `json:"messageStatus"`
	FaxErrorCode  string `json:"faxErrorCode"`
}

// GetMessageList Get Message List
type GetMessageList struct {
	Uri        string                   `json:"uri"`
	Records    []GetMessageInfoResponse `json:"records"`
	Navigation MessagingNavigationInfo  `json:"navigation"`
	Paging     MessagingPagingInfo      `json:"paging"`
}

// MessagingNavigationInfo Messaging Navigation Info
type MessagingNavigationInfo struct {
	FirstPage    MessagingNavigationInfoURI `json:"firstPage"`
	NextPage     MessagingNavigationInfoURI `json:"nextPage"`
	PreviousPage MessagingNavigationInfoURI `json:"previousPage"`
	LastPage     MessagingNavigationInfoURI `json:"lastPage"`
}

// MessagingPagingInfo Messaging Paging Info
type MessagingPagingInfo struct {
	Page          int `json:"page"`
	PerPage       int `json:"perPage"`
	PageStart     int `json:"pageStart"`
	PageEnd       int `json:"pageEnd"`
	TotalPages    int `json:"totalPages"`
	TotalElements int `json:"totalElements"`
}

// MessagingNavigationInfoURI Messaging Navigation Info URI
type MessagingNavigationInfoURI struct {
	Uri string `json:"uri"`
}

// CreateMMSMessage Create MMSMessage
type CreateMMSMessage struct {
	From        MessageStoreCallerInfoRequest   `json:"from"`
	To          []MessageStoreCallerInfoRequest `json:"to"`
	Text        string                          `json:"text"`
	Country     MessageCountryInfo              `json:"country"`
	Attachments []Attachment                    `json:"attachments"`
}

// CreateSMSMessage Create SMSMessage
type CreateSMSMessage struct {
	From    MessageStoreCallerInfoRequest   `json:"from"`
	To      []MessageStoreCallerInfoRequest `json:"to"`
	Text    string                          `json:"text"`
	Country MessageCountryInfo              `json:"country"`
}

// InboundMessageEvent Inbound Message Event
type InboundMessageEvent struct {
	Aps            NotificationInfo `json:"aps"`
	MessageId      string           `json:"messageId"`
	ConversationId string           `json:"conversationId"`
	From           string           `json:"from"`
	To             string           `json:"to"`
	OwnerId        string           `json:"ownerId"`
}

// NotificationInfo Notification Info
type NotificationInfo struct {
	Alert            AlertInfo `json:"alert"`
	Badge            string    `json:"badge"`
	Sound            string    `json:"sound"`
	ContentAvailable string    `json:"content-available"`
	Category         string    `json:"category"`
}

// AlertInfo Alert Info
type AlertInfo struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// InstantMessageEvent Instant Message Event
type InstantMessageEvent struct {
	Uuid           string                  `json:"uuid"`
	Event          string                  `json:"@event"`
	Timestamp      string                  `json:"timestamp"`
	SubscriptionId string                  `json:"subscriptionId"`
	Body           InstantMessageEventBody `json:"body"`
}

// InstantMessageEventBody Instant Message Event Body
type InstantMessageEventBody struct {
	Id               string                      `json:"id"`
	To               []NotificationRecipientInfo `json:"to"`
	From             SenderInfo                  `json:"from"`
	Type             string                      `json:"type"`
	CreationTime     string                      `json:"creationTime"`
	LastModifiedTime string                      `json:"lastModifiedTime"`
	ReadStatus       string                      `json:"readStatus"`
	Priority         string                      `json:"priority"`
	Attachments      []MessageAttachmentInfo     `json:"attachments"`
	Direction        string                      `json:"direction"`
	Availability     string                      `json:"availability"`
	Subject          string                      `json:"subject"`
	MessageStatus    string                      `json:"messageStatus"`
	ConversationId   string                      `json:"conversationId"`
	Conversation     ConversationInfo            `json:"conversation"`
	OwnerId          string                      `json:"ownerId"`
}

// NotificationRecipientInfo Notification Recipient Info
type NotificationRecipientInfo struct {
	PhoneNumber     string `json:"phoneNumber"`
	ExtensionNumber string `json:"extensionNumber"`
	Target          bool   `json:"target"`
	Location        string `json:"location"`
	Name            string `json:"name"`
}

// SenderInfo Sender Info
type SenderInfo struct {
	PhoneNumber     string `json:"phoneNumber"`
	ExtensionNumber string `json:"extensionNumber"`
	Location        string `json:"location"`
	Name            string `json:"name"`
}

// MessageEvent Message Event
type MessageEvent struct {
	Uuid           string           `json:"uuid"`
	Event          string           `json:"@event"`
	Timestamp      string           `json:"timestamp"`
	SubscriptionId string           `json:"subscriptionId"`
	Body           MessageEventBody `json:"body"`
}

// MessageEventBody Message Event Body
type MessageEventBody struct {
	ExtensionId string           `json:"extensionId"`
	LastUpdated string           `json:"lastUpdated"`
	Changes     []MessageChanges `json:"changes"`
	OwnerId     string           `json:"ownerId"`
}

// MessageChanges Message Changes
type MessageChanges struct {
	Type         string `json:"type"`
	NewCount     int    `json:"newCount"`
	UpdatedCount int    `json:"updatedCount"`
}

// ConversationInfo Conversation Info
type ConversationInfo struct {
	Id  string `json:"id"`
	Uri string `json:"uri"`
}

// MessageStoreConfiguration Message Store Configuration
type MessageStoreConfiguration struct {
	RetentionPeriod int `json:"retentionPeriod"`
}

// ListFaxCoverPagesResponse List Fax Cover Pages Response
type ListFaxCoverPagesResponse struct {
	Uri        string                  `json:"uri"`
	Records    []FaxCoverPageInfo      `json:"records"`
	Navigation MessagingNavigationInfo `json:"navigation"`
	Paging     MessagingPagingInfo     `json:"paging"`
}

// FaxCoverPageInfo Fax Cover Page Info
type FaxCoverPageInfo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// VoicemailMessageEvent Voicemail Message Event
type VoicemailMessageEvent struct {
	Uuid           string                    `json:"uuid"`
	Event          string                    `json:"@event"`
	Timestamp      string                    `json:"timestamp"`
	SubscriptionId string                    `json:"subscriptionId"`
	OwnerId        string                    `json:"ownerId"`
	Body           VoicemailMessageEventBody `json:"body"`
}

// VoicemailMessageEventBody Voicemail Message Event Body
type VoicemailMessageEventBody struct {
	Id                    string                      `json:"id"`
	To                    []NotificationRecipientInfo `json:"to"`
	From                  SenderInfo                  `json:"from"`
	Type                  string                      `json:"type"`
	CreationTime          string                      `json:"creationTime"`
	LastModifiedTime      string                      `json:"lastModifiedTime"`
	ReadStatus            string                      `json:"readStatus"`
	Priority              string                      `json:"priority"`
	Attachments           []MessageAttachmentInfo     `json:"attachments"`
	Direction             string                      `json:"direction"`
	Availability          string                      `json:"availability"`
	Subject               string                      `json:"subject"`
	MessageStatus         string                      `json:"messageStatus"`
	ConversationId        string                      `json:"conversationId"`
	VmTranscriptionStatus string                      `json:"vmTranscriptionStatus"`
}

// FaxMessageEvent Fax Message Event
type FaxMessageEvent struct {
	Uuid           string              `json:"uuid"`
	Event          string              `json:"@event"`
	Timestamp      string              `json:"timestamp"`
	SubscriptionId string              `json:"subscriptionId"`
	OwnerId        string              `json:"ownerId"`
	Body           FaxMessageEventBody `json:"body"`
}

// FaxMessageEventBody Fax Message Event Body
type FaxMessageEventBody struct {
	Id               string                      `json:"id"`
	To               []NotificationRecipientInfo `json:"to"`
	From             SenderInfo                  `json:"from"`
	Type             string                      `json:"type"`
	CreationTime     string                      `json:"creationTime"`
	LastModifiedTime string                      `json:"lastModifiedTime"`
	ReadStatus       string                      `json:"readStatus"`
	Priority         string                      `json:"priority"`
	Attachments      []FaxMessageAttachment      `json:"attachments"`
	Direction        string                      `json:"direction"`
	Availability     string                      `json:"availability"`
	Subject          string                      `json:"subject"`
	MessageStatus    string                      `json:"messageStatus"`
	ConversationId   string                      `json:"conversationId"`
	FaxResolution    string                      `json:"faxResolution"`
	FaxPageCount     int                         `json:"faxPageCount"`
}

// FaxMessageAttachment Fax Message Attachment
type FaxMessageAttachment struct {
	Id          string `json:"id"`
	Uri         string `json:"uri"`
	Type        string `json:"type"`
	ContentType string `json:"contentType"`
	Filename    string `json:"filename"`
	Size        int    `json:"size"`
}

// MessageCountryInfo Message Country Info
type MessageCountryInfo struct {
	Id          string `json:"id"`
	Uri         string `json:"uri"`
	Name        string `json:"name"`
	IsoCode     string `json:"isoCode"`
	CallingCode string `json:"callingCode"`
}

// ParsePhoneNumberRequest Parse Phone Number Request
type ParsePhoneNumberRequest struct {
	OriginalStrings []string `json:"originalStrings"`
}

// ParsePhoneNumberResponse Parse Phone Number Response
type ParsePhoneNumberResponse struct {
	Uri          string                        `json:"uri"`
	HomeCountry  GetCountryInfoNumberParser    `json:"homeCountry"`
	PhoneNumbers []PhoneNumberInfoNumberParser `json:"phoneNumbers"`
}

// GetCountryInfoNumberParser Get Country Info Number Parser
type GetCountryInfoNumberParser struct {
	Id               string `json:"id"`
	Uri              string `json:"uri"`
	CallingCode      string `json:"callingCode"`
	EmergencyCalling bool   `json:"emergencyCalling"`
	IsoCode          string `json:"isoCode"`
	Name             string `json:"name"`
}

// PhoneNumberInfoNumberParser Phone Number Info Number Parser
type PhoneNumberInfoNumberParser struct {
	AreaCode               string                     `json:"areaCode"`
	Country                GetCountryInfoNumberParser `json:"country"`
	Dialable               string                     `json:"dialable"`
	E164                   string                     `json:"e164"`
	FormattedInternational string                     `json:"formattedInternational"`
	FormattedNational      string                     `json:"formattedNational"`
	OriginalString         string                     `json:"originalString"`
	Special                bool                       `json:"special"`
	Normalized             string                     `json:"normalized"`
	TollFree               bool                       `json:"tollFree"`
	SubAddress             string                     `json:"subAddress"`
	SubscriberNumber       string                     `json:"subscriberNumber"`
	DtmfPostfix            string                     `json:"dtmfPostfix"`
}

// CreateMessageStoreReportRequest Create Message Store Report Request
type CreateMessageStoreReportRequest struct {
	DateFrom string `json:"dateFrom"`
	DateTo   string `json:"dateTo"`
}

// MessageStoreReport Message Store Report
type MessageStoreReport struct {
	Id               string `json:"id"`
	Uri              string `json:"uri"`
	Status           string `json:"status"`
	AccountId        string `json:"accountId"`
	ExtensionId      string `json:"extensionId"`
	CreationTime     string `json:"creationTime"`
	LastModifiedTime string `json:"lastModifiedTime"`
	DateTo           string `json:"dateTo"`
	DateFrom         string `json:"dateFrom"`
}

// MessageStoreReportArchive Message Store Report Archive
type MessageStoreReportArchive struct {
	Records []ArchiveInfo `json:"records"`
}

// ArchiveInfo Archive Info
type ArchiveInfo struct {
	Size int    `json:"size"`
	Uri  string `json:"uri"`
}

// SwitchesList Switches List
type SwitchesList struct {
	Uri        string                     `json:"uri"`
	Records    []SwitchInfo               `json:"records"`
	Navigation ProvisioningNavigationInfo `json:"navigation"`
	Paging     ProvisioningPagingInfo     `json:"paging"`
}

// SwitchInfo Switch Info
type SwitchInfo struct {
	Uri                 string                              `json:"uri"`
	Id                  string                              `json:"id"`
	ChassisId           string                              `json:"chassisId"`
	Name                string                              `json:"name"`
	Site                SwitchSiteInfo                      `json:"site"`
	EmergencyAddress    LocationUpdatesEmergencyAddressInfo `json:"emergencyAddress"`
	EmergencyLocationId string                              `json:"emergencyLocationId"`
	EmergencyLocation   ERLLocationInfo                     `json:"emergencyLocation"`
}

// CreateSwitchInfo Create Switch Info
type CreateSwitchInfo struct {
	ChassisId           string                                     `json:"chassisId"`
	Name                string                                     `json:"name"`
	Site                SwitchSiteInfo                             `json:"site"`
	EmergencyAddress    LocationUpdatesEmergencyAddressInfoRequest `json:"emergencyAddress"`
	EmergencyLocationId string                                     `json:"emergencyLocationId"`
	EmergencyLocation   ERLLocationInfo                            `json:"emergencyLocation"`
}

// UpdateSwitchInfo Update Switch Info
type UpdateSwitchInfo struct {
	Id                  string                                     `json:"id"`
	ChassisId           string                                     `json:"chassisId"`
	Name                string                                     `json:"name"`
	Site                SwitchSiteInfo                             `json:"site"`
	EmergencyAddress    LocationUpdatesEmergencyAddressInfoRequest `json:"emergencyAddress"`
	EmergencyLocationId string                                     `json:"emergencyLocationId"`
	EmergencyLocation   ERLLocationInfo                            `json:"emergencyLocation"`
}

// LocationUpdatesEmergencyAddressInfoRequest Location Updates Emergency Address Info Request
type LocationUpdatesEmergencyAddressInfoRequest struct {
	Country        string `json:"country"`
	CountryId      string `json:"countryId"`
	CountryIsoCode string `json:"countryIsoCode"`
	CountryName    string `json:"countryName"`
	CustomerName   string `json:"customerName"`
	State          string `json:"state"`
	StateId        string `json:"stateId"`
	StateIsoCode   string `json:"stateIsoCode"`
	StateName      string `json:"stateName"`
	City           string `json:"city"`
	Street         string `json:"street"`
	Street2        string `json:"street2"`
	Zip            string `json:"zip"`
}

// LocationUpdatesEmergencyAddressInfo Location Updates Emergency Address Info
type LocationUpdatesEmergencyAddressInfo struct {
	Country        string `json:"country"`
	CountryId      string `json:"countryId"`
	CountryIsoCode string `json:"countryIsoCode"`
	CountryName    string `json:"countryName"`
	CustomerName   string `json:"customerName"`
	State          string `json:"state"`
	StateId        string `json:"stateId"`
	StateIsoCode   string `json:"stateIsoCode"`
	StateName      string `json:"stateName"`
	City           string `json:"city"`
	Street         string `json:"street"`
	Street2        string `json:"street2"`
	Zip            string `json:"zip"`
}

// SwitchSiteInfo Switch Site Info
type SwitchSiteInfo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// GetExtensionPhoneNumbersResponse Get Extension Phone Numbers Response
type GetExtensionPhoneNumbersResponse struct {
	Uri        string                     `json:"uri"`
	Records    []UserPhoneNumberInfo      `json:"records"`
	Navigation ProvisioningNavigationInfo `json:"navigation"`
	Paging     ProvisioningPagingInfo     `json:"paging"`
}

// UserPhoneNumberInfo User Phone Number Info
type UserPhoneNumberInfo struct {
	Uri                   string                       `json:"uri"`
	Id                    int                          `json:"id"`
	Country               CountryInfo                  `json:"country"`
	ContactCenterProvider ContactCenterProvider        `json:"contactCenterProvider"`
	Extension             UserPhoneNumberExtensionInfo `json:"extension"`
	Label                 string                       `json:"label"`
	Location              string                       `json:"location"`
	PaymentType           string                       `json:"paymentType"`
	PhoneNumber           string                       `json:"phoneNumber"`
	Status                string                       `json:"status"`
	Type                  string                       `json:"type"`
	UsageType             string                       `json:"usageType"`
	Features              []string                     `json:"features"`
}

// CompanyPhoneNumberInfo Company Phone Number Info
type CompanyPhoneNumberInfo struct {
	Uri                   string                `json:"uri"`
	Id                    int                   `json:"id"`
	Country               CountryInfo           `json:"country"`
	Extension             ExtensionInfo         `json:"extension"`
	Label                 string                `json:"label"`
	Location              string                `json:"location"`
	PaymentType           string                `json:"paymentType"`
	PhoneNumber           string                `json:"phoneNumber"`
	Status                string                `json:"status"`
	Type                  string                `json:"type"`
	UsageType             string                `json:"usageType"`
	TemporaryNumber       TemporaryNumberInfo   `json:"temporaryNumber"`
	ContactCenterProvider ContactCenterProvider `json:"contactCenterProvider"`
	VanityPattern         string                `json:"vanityPattern"`
}

// ContactCenterProvider Contact Center Provider
type ContactCenterProvider struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// TemporaryNumberInfo Temporary Number Info
type TemporaryNumberInfo struct {
	Id          string `json:"id"`
	PhoneNumber string `json:"phoneNumber"`
}

// CountryInfo Country Info
type CountryInfo struct {
	Id          string `json:"id"`
	Uri         string `json:"uri"`
	Name        string `json:"name"`
	IsoCode     string `json:"isoCode"`
	CallingCode string `json:"callingCode"`
}

// UserPhoneNumberExtensionInfo User Phone Number Extension Info
type UserPhoneNumberExtensionInfo struct {
	Id                    int                   `json:"id"`
	Uri                   string                `json:"uri"`
	ExtensionNumber       string                `json:"extensionNumber"`
	PartnerId             string                `json:"partnerId"`
	Type                  string                `json:"type"`
	ContactCenterProvider ContactCenterProvider `json:"contactCenterProvider"`
	Name                  string                `json:"name"`
}

// ExtensionCreationResponse Extension Creation Response
type ExtensionCreationResponse struct {
	Id               int                              `json:"id"`
	Uri              string                           `json:"uri"`
	Contact          ContactInfo                      `json:"contact"`
	CustomFields     []CustomFieldInfo                `json:"customFields"`
	ExtensionNumber  string                           `json:"extensionNumber"`
	Name             string                           `json:"name"`
	PartnerId        string                           `json:"partnerId"`
	Permissions      ExtensionPermissions             `json:"permissions"`
	ProfileImage     ProfileImageInfo                 `json:"profileImage"`
	References       []ReferenceInfo                  `json:"references"`
	RegionalSettings RegionalSettings                 `json:"regionalSettings"`
	ServiceFeatures  []ExtensionServiceFeatureInfo    `json:"serviceFeatures"`
	SetupWizardState string                           `json:"setupWizardState"`
	Site             AutomaticLocationUpdatesSiteInfo `json:"site"`
	Status           string                           `json:"status"`
	StatusInfo       ExtensionStatusInfo              `json:"statusInfo"`
	Type             string                           `json:"type"`
	Hidden           bool                             `json:"hidden"`
}

// ProvisioningNavigationInfo Provisioning Navigation Info
type ProvisioningNavigationInfo struct {
	FirstPage    ProvisioningNavigationInfoUri `json:"firstPage"`
	NextPage     ProvisioningNavigationInfoUri `json:"nextPage"`
	PreviousPage ProvisioningNavigationInfoUri `json:"previousPage"`
	LastPage     ProvisioningNavigationInfoUri `json:"lastPage"`
}

// ProvisioningNavigationInfoUri Provisioning Navigation Info Uri
type ProvisioningNavigationInfoUri struct {
	Uri string `json:"uri"`
}

// ProvisioningPagingInfo Provisioning Paging Info
type ProvisioningPagingInfo struct {
	Page          int `json:"page"`
	PerPage       int `json:"perPage"`
	PageStart     int `json:"pageStart"`
	PageEnd       int `json:"pageEnd"`
	TotalPages    int `json:"totalPages"`
	TotalElements int `json:"totalElements"`
}

// GetExtensionListResponse Get Extension List Response
type GetExtensionListResponse struct {
	Uri        string                     `json:"uri"`
	Records    []GetExtensionInfoResponse `json:"records"`
	Navigation ProvisioningNavigationInfo `json:"navigation"`
	Paging     ProvisioningPagingInfo     `json:"paging"`
}

// GetExtensionInfoResponse Get Extension Info Response
type GetExtensionInfoResponse struct {
	Id               int                              `json:"id"`
	Uri              string                           `json:"uri"`
	Account          GetExtensionAccountInfo          `json:"account"`
	Contact          ContactInfo                      `json:"contact"`
	CustomFields     []CustomFieldInfo                `json:"customFields"`
	Departments      []DepartmentInfo                 `json:"departments"`
	ExtensionNumber  string                           `json:"extensionNumber"`
	ExtensionNumbers []string                         `json:"extensionNumbers"`
	Name             string                           `json:"name"`
	PartnerId        string                           `json:"partnerId"`
	Permissions      ExtensionPermissions             `json:"permissions"`
	ProfileImage     ProfileImageInfo                 `json:"profileImage"`
	References       []ReferenceInfo                  `json:"references"`
	Roles            []Roles                          `json:"roles"`
	RegionalSettings RegionalSettings                 `json:"regionalSettings"`
	ServiceFeatures  []ExtensionServiceFeatureInfo    `json:"serviceFeatures"`
	SetupWizardState string                           `json:"setupWizardState"`
	Status           string                           `json:"status"`
	StatusInfo       ExtensionStatusInfo              `json:"statusInfo"`
	Type             string                           `json:"type"`
	CallQueueInfo    CallQueueExtensionInfo           `json:"callQueueInfo"`
	Hidden           bool                             `json:"hidden"`
	Site             AutomaticLocationUpdatesSiteInfo `json:"site"`
}

// GetExtensionAccountInfo Get Extension Account Info
type GetExtensionAccountInfo struct {
	Id  string `json:"id"`
	Uri string `json:"uri"`
}

// ContactInfo Contact Info
type ContactInfo struct {
	FirstName        string                     `json:"firstName"`
	LastName         string                     `json:"lastName"`
	Company          string                     `json:"company"`
	JobTitle         string                     `json:"jobTitle"`
	Email            string                     `json:"email"`
	BusinessPhone    string                     `json:"businessPhone"`
	MobilePhone      string                     `json:"mobilePhone"`
	BusinessAddress  ContactBusinessAddressInfo `json:"businessAddress"`
	EmailAsLoginName bool                       `json:"emailAsLoginName"`
	PronouncedName   PronouncedNameInfo         `json:"pronouncedName"`
	Department       string                     `json:"department"`
}

// ContactBusinessAddressInfo Contact Business Address Info
type ContactBusinessAddressInfo struct {
	Country string `json:"country"`
	State   string `json:"state"`
	City    string `json:"city"`
	Street  string `json:"street"`
	Zip     string `json:"zip"`
}

// DepartmentInfo Department Info
type DepartmentInfo struct {
	Id              string `json:"id"`
	Uri             string `json:"uri"`
	ExtensionNumber string `json:"extensionNumber"`
}

// ExtensionPermissions Extension Permissions
type ExtensionPermissions struct {
	Admin                PermissionInfoAdmin `json:"admin"`
	InternationalCalling PermissionInfoInt   `json:"internationalCalling"`
}

// PermissionInfoAdmin Permission Info Admin
type PermissionInfoAdmin struct {
	Enabled bool `json:"enabled"`
}

// PermissionInfoInt Permission Info Int
type PermissionInfoInt struct {
	Enabled bool `json:"enabled"`
}

// ProfileImageInfo Profile Image Info
type ProfileImageInfo struct {
	Uri          string                `json:"uri"`
	Etag         string                `json:"etag"`
	LastModified string                `json:"lastModified"`
	ContentType  string                `json:"contentType"`
	Scales       []ProfileImageInfoURI `json:"scales"`
}

// ProfileImageInfoURI Profile Image Info URI
type ProfileImageInfoURI struct {
	Uri string `json:"uri"`
}

// ReferenceInfo Reference Info
type ReferenceInfo struct {
	Ref  string `json:"@ref"`
	Type string `json:"type"`
}

// RegionalSettings Regional Settings
type RegionalSettings struct {
	HomeCountry      CountryInfo          `json:"homeCountry"`
	Timezone         TimezoneInfo         `json:"timezone"`
	Language         LanguageInfo         `json:"language"`
	GreetingLanguage GreetingLanguageInfo `json:"greetingLanguage"`
	FormattingLocale FormattingLocaleInfo `json:"formattingLocale"`
	TimeFormat       string               `json:"timeFormat"`
}

// AccountRegionalSettings Account Regional Settings
type AccountRegionalSettings struct {
	HomeCountry      CountryInfo          `json:"homeCountry"`
	Timezone         TimezoneInfo         `json:"timezone"`
	Language         LanguageInfo         `json:"language"`
	GreetingLanguage GreetingLanguageInfo `json:"greetingLanguage"`
	FormattingLocale FormattingLocaleInfo `json:"formattingLocale"`
	TimeFormat       string               `json:"timeFormat"`
	Currency         CurrencyInfo         `json:"currency"`
}

// TimezoneInfo Timezone Info
type TimezoneInfo struct {
	Id          string `json:"id"`
	Uri         string `json:"uri"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Bias        string `json:"bias"`
}

// LanguageInfo Language Info
type LanguageInfo struct {
	Id               string `json:"id"`
	Uri              string `json:"uri"`
	Greeting         bool   `json:"greeting"`
	FormattingLocale bool   `json:"formattingLocale"`
	LocaleCode       string `json:"localeCode"`
	IsoCode          string `json:"isoCode"`
	Name             string `json:"name"`
	Ui               bool   `json:"ui"`
}

// GreetingLanguageInfo Greeting Language Info
type GreetingLanguageInfo struct {
	Id         string `json:"id"`
	LocaleCode string `json:"localeCode"`
	Name       string `json:"name"`
}

// FormattingLocaleInfo Formatting Locale Info
type FormattingLocaleInfo struct {
	Id         string `json:"id"`
	LocaleCode string `json:"localeCode"`
	Name       string `json:"name"`
}

// ExtensionServiceFeatureInfo Extension Service Feature Info
type ExtensionServiceFeatureInfo struct {
	Enabled     bool   `json:"enabled"`
	FeatureName string `json:"featureName"`
	Reason      string `json:"reason"`
}

// ExtensionStatusInfo Extension Status Info
type ExtensionStatusInfo struct {
	Comment string `json:"comment"`
	Reason  string `json:"reason"`
}

// CallQueueExtensionInfo Call Queue Extension Info
type CallQueueExtensionInfo struct {
	SlaGoal                   int  `json:"slaGoal"`
	SlaThresholdSeconds       int  `json:"slaThresholdSeconds"`
	IncludeAbandonedCalls     bool `json:"includeAbandonedCalls"`
	AbandonedThresholdSeconds int  `json:"abandonedThresholdSeconds"`
}

// ExtensionCreationRequest Extension Creation Request
type ExtensionCreationRequest struct {
	Contact          ContactInfoCreationRequest `json:"contact"`
	ExtensionNumber  string                     `json:"extensionNumber"`
	CustomFields     []CustomFieldInfo          `json:"customFields"`
	Password         string                     `json:"password"`
	References       []ReferenceInfo            `json:"references"`
	RegionalSettings RegionalSettings           `json:"regionalSettings"`
	PartnerId        string                     `json:"partnerId"`
	IvrPin           string                     `json:"ivrPin"`
	SetupWizardState string                     `json:"setupWizardState"`
	Site             SiteInfo                   `json:"site"`
	Status           string                     `json:"status"`
	StatusInfo       ExtensionStatusInfo        `json:"statusInfo"`
	Type             string                     `json:"type"`
	Hidden           bool                       `json:"hidden"`
}

// ContactInfoCreationRequest Contact Info Creation Request
type ContactInfoCreationRequest struct {
	FirstName        string                     `json:"firstName"`
	LastName         string                     `json:"lastName"`
	Company          string                     `json:"company"`
	JobTitle         string                     `json:"jobTitle"`
	Email            string                     `json:"email"`
	BusinessPhone    string                     `json:"businessPhone"`
	MobilePhone      string                     `json:"mobilePhone"`
	BusinessAddress  ContactBusinessAddressInfo `json:"businessAddress"`
	EmailAsLoginName bool                       `json:"emailAsLoginName"`
	PronouncedName   PronouncedNameInfo         `json:"pronouncedName"`
	Department       string                     `json:"department"`
}

// PronouncedNameInfo Pronounced Name Info
type PronouncedNameInfo struct {
	Type   string                   `json:"type"`
	Text   string                   `json:"text"`
	Prompt PronouncedNamePromptInfo `json:"prompt"`
}

// PronouncedNamePromptInfo Pronounced Name Prompt Info
type PronouncedNamePromptInfo struct {
	Id          string `json:"id"`
	ContentUri  string `json:"contentUri"`
	ContentType string `json:"contentType"`
}

// ExtensionUpdateRequest Extension Update Request
type ExtensionUpdateRequest struct {
	Status           string                           `json:"status"`
	StatusInfo       ExtensionStatusInfo              `json:"statusInfo"`
	Reason           string                           `json:"reason"`
	Comment          string                           `json:"comment"`
	ExtensionNumber  string                           `json:"extensionNumber"`
	Contact          ContactInfoUpdateRequest         `json:"contact"`
	RegionalSettings ExtensionRegionalSettingRequest  `json:"regionalSettings"`
	SetupWizardState string                           `json:"setupWizardState"`
	PartnerId        string                           `json:"partnerId"`
	IvrPin           string                           `json:"ivrPin"`
	Password         string                           `json:"password"`
	CallQueueInfo    CallQueueInfoRequest             `json:"callQueueInfo"`
	Transition       []UserTransitionInfo             `json:"transition"`
	CustomFields     []CustomFieldInfo                `json:"customFields"`
	Hidden           bool                             `json:"hidden"`
	Site             AutomaticLocationUpdatesSiteInfo `json:"site"`
	Type             string                           `json:"type"`
	References       []ReferenceInfo                  `json:"references"`
}

// UserTransitionInfo User Transition Info
type UserTransitionInfo struct {
	SendWelcomeEmailsToUsers bool `json:"sendWelcomeEmailsToUsers"`
	SendWelcomeEmail         bool `json:"sendWelcomeEmail"`
}

// CustomFieldInfo Custom Field Info
type CustomFieldInfo struct {
	Id          string `json:"id"`
	Value       string `json:"value"`
	DisplayName string `json:"displayName"`
}

// ContactInfoUpdateRequest Contact Info Update Request
type ContactInfoUpdateRequest struct {
	FirstName        string                     `json:"firstName"`
	LastName         string                     `json:"lastName"`
	Company          string                     `json:"company"`
	JobTitle         string                     `json:"jobTitle"`
	Email            string                     `json:"email"`
	BusinessPhone    string                     `json:"businessPhone"`
	MobilePhone      string                     `json:"mobilePhone"`
	BusinessAddress  ContactBusinessAddressInfo `json:"businessAddress"`
	EmailAsLoginName bool                       `json:"emailAsLoginName"`
	PronouncedName   PronouncedNameInfo         `json:"pronouncedName"`
	Department       string                     `json:"department"`
}

// CallQueueInfoRequest Call Queue Info Request
type CallQueueInfoRequest struct {
	SlaGoal                   int  `json:"slaGoal"`
	SlaThresholdSeconds       int  `json:"slaThresholdSeconds"`
	IncludeAbandonedCalls     bool `json:"includeAbandonedCalls"`
	AbandonedThresholdSeconds int  `json:"abandonedThresholdSeconds"`
}

// ExtensionRegionalSettingRequest Extension Regional Setting Request
type ExtensionRegionalSettingRequest struct {
	HomeCountry      ExtensionCountryInfoRequest          `json:"homeCountry"`
	Timezone         ExtensionTimezoneInfoRequest         `json:"timezone"`
	Language         ExtensionLanguageInfoRequest         `json:"language"`
	GreetingLanguage ExtensionGreetingLanguageInfoRequest `json:"greetingLanguage"`
	FormattingLocale ExtensionFormattingLocaleInfoRequest `json:"formattingLocale"`
	TimeFormat       string                               `json:"timeFormat"`
}

// ExtensionCountryInfoRequest Extension Country Info Request
type ExtensionCountryInfoRequest struct {
	Id string `json:"id"`
}

// ExtensionTimezoneInfoRequest Extension Timezone Info Request
type ExtensionTimezoneInfoRequest struct {
	Id string `json:"id"`
}

// ExtensionLanguageInfoRequest Extension Language Info Request
type ExtensionLanguageInfoRequest struct {
	Id string `json:"id"`
}

// ExtensionGreetingLanguageInfoRequest Extension Greeting Language Info Request
type ExtensionGreetingLanguageInfoRequest struct {
	Id string `json:"id"`
}

// ExtensionFormattingLocaleInfoRequest Extension Formatting Locale Info Request
type ExtensionFormattingLocaleInfoRequest struct {
	Id string `json:"id"`
}

// ExtensionCallerIdInfo Extension Caller Id Info
type ExtensionCallerIdInfo struct {
	Uri                             string              `json:"uri"`
	ByDevice                        []CallerIdByDevice  `json:"byDevice"`
	ByFeature                       []CallerIdByFeature `json:"byFeature"`
	ExtensionNameForOutboundCalls   bool                `json:"extensionNameForOutboundCalls"`
	ExtensionNumberForInternalCalls bool                `json:"extensionNumberForInternalCalls"`
}

// CallerIdByDevice Caller Id By Device
type CallerIdByDevice struct {
	Device   CallerIdDeviceInfo   `json:"device"`
	CallerId CallerIdByDeviceInfo `json:"callerId"`
}

// CallerIdByFeature Caller Id By Feature
type CallerIdByFeature struct {
	Feature  string                `json:"feature"`
	CallerId CallerIdByFeatureInfo `json:"callerId"`
}

// CallerIdDeviceInfo Caller Id Device Info
type CallerIdDeviceInfo struct {
	Id   string `json:"id"`
	Uri  string `json:"uri"`
	Name string `json:"name"`
}

// CallerIdByDeviceInfo Caller Id By Device Info
type CallerIdByDeviceInfo struct {
	Type      string            `json:"type"`
	PhoneInfo CallerIdPhoneInfo `json:"phoneInfo"`
}

// CallerIdByFeatureInfo Caller Id By Feature Info
type CallerIdByFeatureInfo struct {
	Type      string            `json:"type"`
	PhoneInfo CallerIdPhoneInfo `json:"phoneInfo"`
}

// CallerIdPhoneInfo Caller Id Phone Info
type CallerIdPhoneInfo struct {
	Id          string `json:"id"`
	Uri         string `json:"uri"`
	PhoneNumber string `json:"phoneNumber"`
}

// GetConferencingInfoResponse Get Conferencing Info Response
type GetConferencingInfoResponse struct {
	Uri                 string                        `json:"uri"`
	AllowJoinBeforeHost bool                          `json:"allowJoinBeforeHost"`
	HostCode            string                        `json:"hostCode"`
	Mode                string                        `json:"mode"`
	ParticipantCode     string                        `json:"participantCode"`
	PhoneNumber         string                        `json:"phoneNumber"`
	TapToJoinUri        string                        `json:"tapToJoinUri"`
	PhoneNumbers        []PhoneNumberInfoConferencing `json:"phoneNumbers"`
}

// PhoneNumberInfoConferencing Phone Number Info Conferencing
type PhoneNumberInfoConferencing struct {
	Country     GetCountryInfoConferencing `json:"country"`
	Default     bool                       `json:"@default"`
	HasGreeting bool                       `json:"hasGreeting"`
	Location    string                     `json:"location"`
	PhoneNumber string                     `json:"phoneNumber"`
	Premium     bool                       `json:"premium"`
}

// GetCountryInfoConferencing Get Country Info Conferencing
type GetCountryInfoConferencing struct {
	Id               string `json:"id"`
	Uri              string `json:"uri"`
	CallingCode      string `json:"callingCode"`
	EmergencyCalling bool   `json:"emergencyCalling"`
	IsoCode          string `json:"isoCode"`
	Name             string `json:"name"`
}

// GetExtensionGrantListResponse Get Extension Grant List Response
type GetExtensionGrantListResponse struct {
	Uri        string                     `json:"uri"`
	Records    []GrantInfo                `json:"records"`
	Navigation ProvisioningNavigationInfo `json:"navigation"`
	Paging     ProvisioningPagingInfo     `json:"paging"`
}

// GrantInfo Grant Info
type GrantInfo struct {
	Uri                   string              `json:"uri"`
	Extension             ExtensionInfoGrants `json:"extension"`
	CallPickup            bool                `json:"callPickup"`
	CallMonitoring        bool                `json:"callMonitoring"`
	CallOnBehalfOf        bool                `json:"callOnBehalfOf"`
	CallDelegation        bool                `json:"callDelegation"`
	GroupPaging           bool                `json:"groupPaging"`
	CallQueueSetup        bool                `json:"callQueueSetup"`
	CallQueueMembersSetup bool                `json:"callQueueMembersSetup"`
}

// ExtensionInfoGrants Extension Info Grants
type ExtensionInfoGrants struct {
	Id              string `json:"id"`
	Uri             string `json:"uri"`
	ExtensionNumber string `json:"extensionNumber"`
	Name            string `json:"name"`
	Type            string `json:"type"`
}

// UpdateConferencingInfoRequest Update Conferencing Info Request
type UpdateConferencingInfoRequest struct {
	PhoneNumbers        []ConferencePhoneNumberInfo `json:"phoneNumbers"`
	AllowJoinBeforeHost bool                        `json:"allowJoinBeforeHost"`
}

// ConferencePhoneNumberInfo Conference Phone Number Info
type ConferencePhoneNumberInfo struct {
	PhoneNumber string `json:"phoneNumber"`
	Default     bool   `json:"@default"`
}

// NotificationSettings Notification Settings
type NotificationSettings struct {
	Uri               string               `json:"uri"`
	EmailRecipients   []EmailRecipientInfo `json:"emailRecipients"`
	EmailAddresses    []string             `json:"emailAddresses"`
	SmsEmailAddresses []string             `json:"smsEmailAddresses"`
	AdvancedMode      bool                 `json:"advancedMode"`
	Voicemails        VoicemailsInfo       `json:"voicemails"`
	InboundFaxes      InboundFaxesInfo     `json:"inboundFaxes"`
	OutboundFaxes     OutboundFaxesInfo    `json:"outboundFaxes"`
	InboundTexts      InboundTextsInfo     `json:"inboundTexts"`
	MissedCalls       MissedCallsInfo      `json:"missedCalls"`
}

// EmailRecipientInfo Email Recipient Info
type EmailRecipientInfo struct {
	ExtensionId     string   `json:"extensionId"`
	FullName        string   `json:"fullName"`
	ExtensionNumber string   `json:"extensionNumber"`
	Status          string   `json:"status"`
	EmailAddresses  []string `json:"emailAddresses"`
	Permission      string   `json:"permission"`
}

// NotificationSettingsUpdateRequest Notification Settings Update Request
type NotificationSettingsUpdateRequest struct {
	EmailAddresses    []string          `json:"emailAddresses"`
	SmsEmailAddresses []string          `json:"smsEmailAddresses"`
	AdvancedMode      bool              `json:"advancedMode"`
	Voicemails        VoicemailsInfo    `json:"voicemails"`
	InboundFaxes      InboundFaxesInfo  `json:"inboundFaxes"`
	OutboundFaxes     OutboundFaxesInfo `json:"outboundFaxes"`
	InboundTexts      InboundTextsInfo  `json:"inboundTexts"`
	MissedCalls       MissedCallsInfo   `json:"missedCalls"`
}

// VoicemailsInfo Voicemails Info
type VoicemailsInfo struct {
	NotifyByEmail             bool     `json:"notifyByEmail"`
	NotifyBySms               bool     `json:"notifyBySms"`
	AdvancedEmailAddresses    []string `json:"advancedEmailAddresses"`
	AdvancedSmsEmailAddresses []string `json:"advancedSmsEmailAddresses"`
	IncludeAttachment         bool     `json:"includeAttachment"`
	IncludeTranscription      bool     `json:"includeTranscription"`
	MarkAsRead                bool     `json:"markAsRead"`
}

// InboundFaxesInfo Inbound Faxes Info
type InboundFaxesInfo struct {
	NotifyByEmail             bool     `json:"notifyByEmail"`
	NotifyBySms               bool     `json:"notifyBySms"`
	AdvancedEmailAddresses    []string `json:"advancedEmailAddresses"`
	AdvancedSmsEmailAddresses []string `json:"advancedSmsEmailAddresses"`
	IncludeAttachment         bool     `json:"includeAttachment"`
	MarkAsRead                bool     `json:"markAsRead"`
}

// OutboundFaxesInfo Outbound Faxes Info
type OutboundFaxesInfo struct {
	NotifyByEmail             bool     `json:"notifyByEmail"`
	NotifyBySms               bool     `json:"notifyBySms"`
	AdvancedEmailAddresses    []string `json:"advancedEmailAddresses"`
	AdvancedSmsEmailAddresses []string `json:"advancedSmsEmailAddresses"`
}

// InboundTextsInfo Inbound Texts Info
type InboundTextsInfo struct {
	NotifyByEmail             bool     `json:"notifyByEmail"`
	NotifyBySms               bool     `json:"notifyBySms"`
	AdvancedEmailAddresses    []string `json:"advancedEmailAddresses"`
	AdvancedSmsEmailAddresses []string `json:"advancedSmsEmailAddresses"`
}

// MissedCallsInfo Missed Calls Info
type MissedCallsInfo struct {
	NotifyByEmail             bool     `json:"notifyByEmail"`
	NotifyBySms               bool     `json:"notifyBySms"`
	AdvancedEmailAddresses    []string `json:"advancedEmailAddresses"`
	AdvancedSmsEmailAddresses []string `json:"advancedSmsEmailAddresses"`
}

// GetAccountInfoResponse Get Account Info Response
type GetAccountInfoResponse struct {
	Id                 int                     `json:"id"`
	Uri                string                  `json:"uri"`
	Bsid               string                  `json:"bsid"`
	MainNumber         string                  `json:"mainNumber"`
	Operator           AccountOperatorInfo     `json:"@operator"`
	PartnerId          string                  `json:"partnerId"`
	ServiceInfo        ServiceInfo             `json:"serviceInfo"`
	SetupWizardState   string                  `json:"setupWizardState"`
	SignupInfo         SignupInfoResource      `json:"signupInfo"`
	Status             string                  `json:"status"`
	StatusInfo         AccountStatusInfo       `json:"statusInfo"`
	RegionalSettings   AccountRegionalSettings `json:"regionalSettings"`
	Federated          bool                    `json:"federated"`
	OutboundCallPrefix int                     `json:"outboundCallPrefix"`
	Cfid               string                  `json:"cfid"`
	Limits             AccountLimits           `json:"limits"`
}

// ServiceInfo Service Info
type ServiceInfo struct {
	Uri               string                `json:"uri"`
	BillingPlan       BillingPlanInfo       `json:"billingPlan"`
	Brand             BrandInfo             `json:"brand"`
	ServicePlan       ServicePlanInfo       `json:"servicePlan"`
	TargetServicePlan TargetServicePlanInfo `json:"targetServicePlan"`
	ContractedCountry ContractedCountryInfo `json:"contractedCountry"`
}

// ContractedCountryInfo Contracted Country Info
type ContractedCountryInfo struct {
	Id  string `json:"id"`
	Uri string `json:"uri"`
}

// BillingPlanInfo Billing Plan Info
type BillingPlanInfo struct {
	Id                 string `json:"id"`
	Name               string `json:"name"`
	DurationUnit       string `json:"durationUnit"`
	Duration           int    `json:"duration"`
	Type               string `json:"type"`
	IncludedPhoneLines int    `json:"includedPhoneLines"`
}

// BrandInfo Brand Info
type BrandInfo struct {
	Id          string      `json:"id"`
	Name        string      `json:"name"`
	HomeCountry CountryInfo `json:"homeCountry"`
}

// ServicePlanInfo Service Plan Info
type ServicePlanInfo struct {
	Id                  string `json:"id"`
	Name                string `json:"name"`
	Edition             string `json:"edition"`
	FreemiumProductType string `json:"freemiumProductType"`
}

// TargetServicePlanInfo Target Service Plan Info
type TargetServicePlanInfo struct {
	Id                  string `json:"id"`
	Name                string `json:"name"`
	Edition             string `json:"edition"`
	FreemiumProductType string `json:"freemiumProductType"`
}

// AccountStatusInfo Account Status Info
type AccountStatusInfo struct {
	Comment string `json:"comment"`
	Reason  string `json:"reason"`
	Till    string `json:"till"`
}

// GetServiceInfoResponse Get Service Info Response
type GetServiceInfoResponse struct {
	Uri               string                `json:"uri"`
	ServicePlanName   string                `json:"servicePlanName"`
	Brand             BrandInfo             `json:"brand"`
	ContractedCountry ContractedCountryInfo `json:"contractedCountry"`
	ServicePlan       ServicePlanInfo       `json:"servicePlan"`
	TargetServicePlan TargetServicePlanInfo `json:"targetServicePlan"`
	BillingPlan       BillingPlanInfo       `json:"billingPlan"`
	ServiceFeatures   []ServiceFeatureInfo  `json:"serviceFeatures"`
	Limits            AccountLimits         `json:"limits"`
	Package           PackageInfo           `json:"package"`
}

// PackageInfo Package Info
type PackageInfo struct {
	Version string `json:"version"`
	Id      string `json:"id"`
}

// AccountLimits Account Limits
type AccountLimits struct {
	FreeSoftPhoneLinesPerExtension int `json:"freeSoftPhoneLinesPerExtension"`
	MeetingSize                    int `json:"meetingSize"`
	CloudRecordingStorage          int `json:"cloudRecordingStorage"`
	MaxMonitoredExtensionsPerUser  int `json:"maxMonitoredExtensionsPerUser"`
	MaxExtensionNumberLength       int `json:"maxExtensionNumberLength"`
	SiteCodeLength                 int `json:"siteCodeLength"`
	ShortExtensionNumberLength     int `json:"shortExtensionNumberLength"`
}

// ServiceFeatureInfo Service Feature Info
type ServiceFeatureInfo struct {
	FeatureName string `json:"featureName"`
	Enabled     bool   `json:"enabled"`
}

// AccountBusinessAddressResource Account Business Address Resource
type AccountBusinessAddressResource struct {
	Uri             string                     `json:"uri"`
	BusinessAddress ContactBusinessAddressInfo `json:"businessAddress"`
	Company         string                     `json:"company"`
	Email           string                     `json:"email"`
	MainSiteName    string                     `json:"mainSiteName"`
}

// ModifyAccountBusinessAddressRequest Modify Account Business Address Request
type ModifyAccountBusinessAddressRequest struct {
	Company         string              `json:"company"`
	Email           string              `json:"email"`
	BusinessAddress BusinessAddressInfo `json:"businessAddress"`
	MainSiteName    string              `json:"mainSiteName"`
}

// BusinessAddressInfo Business Address Info
type BusinessAddressInfo struct {
	Country string `json:"country"`
	State   string `json:"state"`
	City    string `json:"city"`
	Street  string `json:"street"`
	Zip     string `json:"zip"`
}

// SignupInfoResource Signup Info Resource
type SignupInfoResource struct {
	TosAccepted        bool     `json:"tosAccepted"`
	SignupState        []string `json:"signupState"`
	VerificationReason string   `json:"verificationReason"`
	MarketingAccepted  bool     `json:"marketingAccepted"`
}

// LanguageList Language List
type LanguageList struct {
	Uri        string                     `json:"uri"`
	Records    []LanguageInfo             `json:"records"`
	Navigation ProvisioningNavigationInfo `json:"navigation"`
	Paging     ProvisioningPagingInfo     `json:"paging"`
}

// GetCountryListResponse Get Country List Response
type GetCountryListResponse struct {
	Uri        string                             `json:"uri"`
	Records    []GetCountryInfoDictionaryResponse `json:"records"`
	Navigation ProvisioningNavigationInfo         `json:"navigation"`
	Paging     ProvisioningPagingInfo             `json:"paging"`
}

// GetCountryInfoDictionaryResponse Get Country Info Dictionary Response
type GetCountryInfoDictionaryResponse struct {
	Id                string `json:"id"`
	Uri               string `json:"uri"`
	CallingCode       string `json:"callingCode"`
	EmergencyCalling  bool   `json:"emergencyCalling"`
	IsoCode           string `json:"isoCode"`
	Name              string `json:"name"`
	NumberSelling     bool   `json:"numberSelling"`
	LoginAllowed      bool   `json:"loginAllowed"`
	SignupAllowed     bool   `json:"signupAllowed"`
	FreeSoftphoneLine bool   `json:"freeSoftphoneLine"`
}

// GetLocationListResponse Get Location List Response
type GetLocationListResponse struct {
	Uri        string                     `json:"uri"`
	Records    []LocationInfo             `json:"records"`
	Navigation ProvisioningNavigationInfo `json:"navigation"`
	Paging     ProvisioningPagingInfo     `json:"paging"`
}

// LocationInfo Location Info
type LocationInfo struct {
	Uri      string            `json:"uri"`
	AreaCode string            `json:"areaCode"`
	City     string            `json:"city"`
	Npa      string            `json:"npa"`
	Nxx      string            `json:"nxx"`
	State    LocationStateInfo `json:"state"`
}

// LocationStateInfo Location State Info
type LocationStateInfo struct {
	Id  string `json:"id"`
	Uri string `json:"uri"`
}

// GetStateListResponse Get State List Response
type GetStateListResponse struct {
	Uri        string                     `json:"uri"`
	Records    []GetStateInfoResponse     `json:"records"`
	Navigation ProvisioningNavigationInfo `json:"navigation"`
	Paging     ProvisioningPagingInfo     `json:"paging"`
}

// GetStateInfoResponse Get State Info Response
type GetStateInfoResponse struct {
	Id      string              `json:"id"`
	Uri     string              `json:"uri"`
	Country GetCountryInfoState `json:"country"`
	IsoCode string              `json:"isoCode"`
	Name    string              `json:"name"`
}

// GetCountryInfoState Get Country Info State
type GetCountryInfoState struct {
	Id  string `json:"id"`
	Uri string `json:"uri"`
}

// GetTimezoneListResponse Get Timezone List Response
type GetTimezoneListResponse struct {
	Uri        string                     `json:"uri"`
	Records    []GetTimezoneInfoResponse  `json:"records"`
	Navigation ProvisioningNavigationInfo `json:"navigation"`
	Paging     ProvisioningPagingInfo     `json:"paging"`
}

// GetTimezoneInfoResponse Get Timezone Info Response
type GetTimezoneInfoResponse struct {
	Id          string `json:"id"`
	Uri         string `json:"uri"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Bias        string `json:"bias"`
}

// CurrencyInfo Currency Info
type CurrencyInfo struct {
	Id          string `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	MinorSymbol string `json:"minorSymbol"`
}

// AccountOperatorInfo Account Operator Info
type AccountOperatorInfo struct {
	Uri             string `json:"uri"`
	Id              int    `json:"id"`
	ExtensionNumber string `json:"extensionNumber"`
}

// AccountPhoneNumbers Account Phone Numbers
type AccountPhoneNumbers struct {
	Uri        string                     `json:"uri"`
	Records    []CompanyPhoneNumberInfo   `json:"records"`
	Navigation ProvisioningNavigationInfo `json:"navigation"`
	Paging     ProvisioningPagingInfo     `json:"paging"`
}

// Roles Roles
type Roles struct {
	Uri string `json:"uri"`
	Id  string `json:"id"`
}

// PagingOnlyGroupUsers Paging Only Group Users
type PagingOnlyGroupUsers struct {
	Uri        string                     `json:"uri"`
	Records    []PagingGroupExtensionInfo `json:"records"`
	Navigation ProvisioningNavigationInfo `json:"navigation"`
	Paging     ProvisioningPagingInfo     `json:"paging"`
}

// PagingGroupExtensionInfo Paging Group Extension Info
type PagingGroupExtensionInfo struct {
	Id              string `json:"id"`
	Uri             string `json:"uri"`
	ExtensionNumber string `json:"extensionNumber"`
	Name            string `json:"name"`
}

// PagingOnlyGroupDevices Paging Only Group Devices
type PagingOnlyGroupDevices struct {
	Uri        string                     `json:"uri"`
	Records    []PagingDeviceInfo         `json:"records"`
	Navigation ProvisioningNavigationInfo `json:"navigation"`
	Paging     ProvisioningPagingInfo     `json:"paging"`
}

// PagingDeviceInfo Paging Device Info
type PagingDeviceInfo struct {
	Id   string `json:"id"`
	Uri  string `json:"uri"`
	Name string `json:"name"`
}

// EditPagingGroupRequest Edit Paging Group Request
type EditPagingGroupRequest struct {
	AddedUserIds     []string `json:"addedUserIds"`
	RemovedUserIds   []string `json:"removedUserIds"`
	AddedDeviceIds   []string `json:"addedDeviceIds"`
	RemovedDeviceIds []string `json:"removedDeviceIds"`
}

// ErrorEntity Error Entity
type ErrorEntity struct {
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
}

// EmergencyLocationList Emergency Location List
type EmergencyLocationList struct {
	Records    []EmergencyLocationInfo    `json:"records"`
	Navigation ProvisioningNavigationInfo `json:"navigation"`
	Paging     ProvisioningPagingInfo     `json:"paging"`
}

// EmergencyLocationInfoRequest Emergency Location Info Request
type EmergencyLocationInfoRequest struct {
	Id            string                       `json:"id"`
	Address       EmergencyLocationAddressInfo `json:"address"`
	Name          string                       `json:"name"`
	Site          ShortSiteInfo                `json:"site"`
	AddressStatus string                       `json:"addressStatus"`
	UsageStatus   string                       `json:"usageStatus"`
	Visibility    string                       `json:"visibility"`
	Owners        []LocationOwnerInfo          `json:"owners"`
}

// EmergencyLocationInfo Emergency Location Info
type EmergencyLocationInfo struct {
	Id            string                       `json:"id"`
	Address       EmergencyLocationAddressInfo `json:"address"`
	Name          string                       `json:"name"`
	Site          ShortSiteInfo                `json:"site"`
	AddressStatus string                       `json:"addressStatus"`
	UsageStatus   string                       `json:"usageStatus"`
	SyncStatus    string                       `json:"syncStatus"`
	Visibility    string                       `json:"visibility"`
	Owners        []LocationOwnerInfo          `json:"owners"`
}

// LocationOwnerInfo Location Owner Info
type LocationOwnerInfo struct {
	Id string `json:"id"`
}

// EmergencyLocationAddressInfo Emergency Location Address Info
type EmergencyLocationAddressInfo struct {
	Country        string `json:"country"`
	CountryId      string `json:"countryId"`
	CountryIsoCode string `json:"countryIsoCode"`
	CountryName    string `json:"countryName"`
	State          string `json:"state"`
	StateId        string `json:"stateId"`
	StateIsoCode   string `json:"stateIsoCode"`
	StateName      string `json:"stateName"`
	City           string `json:"city"`
	Street         string `json:"street"`
	Street2        string `json:"street2"`
	Zip            string `json:"zip"`
	CustomerName   string `json:"customerName"`
}

// CallQueueBulkAssignResource Call Queue Bulk Assign Resource
type CallQueueBulkAssignResource struct {
	AddedExtensionIds   []string `json:"addedExtensionIds"`
	RemovedExtensionIds []string `json:"removedExtensionIds"`
}

// BulkAssignAutomaticLocationUpdatesUsers Bulk Assign Automatic Location Updates Users
type BulkAssignAutomaticLocationUpdatesUsers struct {
	EnabledUserIds  []string `json:"enabledUserIds"`
	DisabledUserIds []string `json:"disabledUserIds"`
}

// CreateWirelessPoint Create Wireless Point
type CreateWirelessPoint struct {
	Bssid               string                                     `json:"bssid"`
	Name                string                                     `json:"name"`
	Site                AutomaticLocationUpdatesSiteInfo           `json:"site"`
	EmergencyAddress    LocationUpdatesEmergencyAddressInfoRequest `json:"emergencyAddress"`
	EmergencyLocationId string                                     `json:"emergencyLocationId"`
	EmergencyLocation   ERLLocationInfo                            `json:"emergencyLocation"`
}

// UpdateWirelessPoint Update Wireless Point
type UpdateWirelessPoint struct {
	Id                  string                                     `json:"id"`
	Bssid               string                                     `json:"bssid"`
	Name                string                                     `json:"name"`
	Site                AutomaticLocationUpdatesSiteInfo           `json:"site"`
	EmergencyAddress    LocationUpdatesEmergencyAddressInfoRequest `json:"emergencyAddress"`
	EmergencyLocationId string                                     `json:"emergencyLocationId"`
	EmergencyLocation   ERLLocationInfo                            `json:"emergencyLocation"`
}

// WirelessPointsList Wireless Points List
type WirelessPointsList struct {
	Uri        string                     `json:"uri"`
	Records    []WirelessPointInfo        `json:"records"`
	Navigation ProvisioningNavigationInfo `json:"navigation"`
	Paging     ProvisioningPagingInfo     `json:"paging"`
}

// WirelessPointInfo Wireless Point Info
type WirelessPointInfo struct {
	Uri                 string                              `json:"uri"`
	Id                  string                              `json:"id"`
	Bssid               string                              `json:"bssid"`
	Name                string                              `json:"name"`
	Site                AutomaticLocationUpdatesSiteInfo    `json:"site"`
	EmergencyAddress    LocationUpdatesEmergencyAddressInfo `json:"emergencyAddress"`
	EmergencyLocation   ERLLocationInfo                     `json:"emergencyLocation"`
	EmergencyLocationId string                              `json:"emergencyLocationId"`
}

// ERLLocationInfo ERLLocation Info
type ERLLocationInfo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// WirelessPointInfoRequest Wireless Point Info Request
type WirelessPointInfoRequest struct {
	Id                  string                                     `json:"id"`
	Bssid               string                                     `json:"bssid"`
	Name                string                                     `json:"name"`
	Site                AutomaticLocationUpdatesSiteInfo           `json:"site"`
	EmergencyAddress    LocationUpdatesEmergencyAddressInfoRequest `json:"emergencyAddress"`
	EmergencyLocationId string                                     `json:"emergencyLocationId"`
}

// UpdateNetworkRequest Update Network Request
type UpdateNetworkRequest struct {
	Name              string                           `json:"name"`
	Site              AutomaticLocationUpdatesSiteInfo `json:"site"`
	PublicIpRanges    []PublicIpRangeInfo              `json:"publicIpRanges"`
	PrivateIpRanges   []PrivateIpRangeInfoRequest      `json:"privateIpRanges"`
	EmergencyLocation ERLLocationInfo                  `json:"emergencyLocation"`
}

// CreateNetworkRequest Create Network Request
type CreateNetworkRequest struct {
	Name              string                           `json:"name"`
	Site              AutomaticLocationUpdatesSiteInfo `json:"site"`
	PublicIpRanges    []PublicIpRangeInfo              `json:"publicIpRanges"`
	PrivateIpRanges   []PrivateIpRangeInfoRequest      `json:"privateIpRanges"`
	EmergencyLocation ERLLocationInfo                  `json:"emergencyLocation"`
}

// PublicIpRangeInfo Public Ip Range Info
type PublicIpRangeInfo struct {
	Id      string `json:"id"`
	StartIp string `json:"startIp"`
	EndIp   string `json:"endIp"`
	Matched bool   `json:"matched"`
}

// PrivateIpRangeInfo Private Ip Range Info
type PrivateIpRangeInfo struct {
	Id                  string                              `json:"id"`
	StartIp             string                              `json:"startIp"`
	EndIp               string                              `json:"endIp"`
	Name                string                              `json:"name"`
	EmergencyAddress    LocationUpdatesEmergencyAddressInfo `json:"emergencyAddress"`
	EmergencyLocationId string                              `json:"emergencyLocationId"`
	Matched             bool                                `json:"matched"`
}

// PrivateIpRangeInfoRequest Private Ip Range Info Request
type PrivateIpRangeInfoRequest struct {
	Id                  string                                     `json:"id"`
	StartIp             string                                     `json:"startIp"`
	EndIp               string                                     `json:"endIp"`
	Name                string                                     `json:"name"`
	EmergencyAddress    LocationUpdatesEmergencyAddressInfoRequest `json:"emergencyAddress"`
	EmergencyLocationId string                                     `json:"emergencyLocationId"`
}

// CallMonitoringBulkAssign Call Monitoring Bulk Assign
type CallMonitoringBulkAssign struct {
	AddedExtensions   []CallMonitoringExtensionInfo `json:"addedExtensions"`
	UpdatedExtensions []CallMonitoringExtensionInfo `json:"updatedExtensions"`
	RemovedExtensions []CallMonitoringExtensionInfo `json:"removedExtensions"`
}

// CallMonitoringExtensionInfo Call Monitoring Extension Info
type CallMonitoringExtensionInfo struct {
	Id          string   `json:"id"`
	Permissions []string `json:"permissions"`
}

// CallMonitoringGroups Call Monitoring Groups
type CallMonitoringGroups struct {
	Uri        string                     `json:"uri"`
	Records    []CallMonitoringGroup      `json:"records"`
	Navigation ProvisioningNavigationInfo `json:"navigation"`
	Paging     ProvisioningPagingInfo     `json:"paging"`
}

// CreateCallMonitoringGroupRequest Create Call Monitoring Group Request
type CreateCallMonitoringGroupRequest struct {
	Name string `json:"name"`
}

// CallMonitoringGroup Call Monitoring Group
type CallMonitoringGroup struct {
	Uri  string `json:"uri"`
	Id   string `json:"id"`
	Name string `json:"name"`
}

// CallMonitoringGroupMemberList Call Monitoring Group Member List
type CallMonitoringGroupMemberList struct {
	Uri        string                          `json:"uri"`
	Records    []CallMonitoringGroupMemberInfo `json:"records"`
	Navigation ProvisioningNavigationInfo      `json:"navigation"`
	Paging     ProvisioningPagingInfo          `json:"paging"`
}

// CallMonitoringGroupMemberInfo Call Monitoring Group Member Info
type CallMonitoringGroupMemberInfo struct {
	Uri             string   `json:"uri"`
	Id              string   `json:"id"`
	ExtensionNumber string   `json:"extensionNumber"`
	Permissions     []string `json:"permissions"`
}

// CallQueueMembers Call Queue Members
type CallQueueMembers struct {
	Uri        string                     `json:"uri"`
	Records    []CallQueueMemberInfo      `json:"records"`
	Navigation ProvisioningNavigationInfo `json:"navigation"`
	Paging     ProvisioningPagingInfo     `json:"paging"`
}

// CallQueueMemberInfo Call Queue Member Info
type CallQueueMemberInfo struct {
	Uri             string `json:"uri"`
	Id              int    `json:"id"`
	ExtensionNumber string `json:"extensionNumber"`
}

// CallQueues Call Queues
type CallQueues struct {
	Uri        string                     `json:"uri"`
	Records    []CallQueueInfo            `json:"records"`
	Navigation ProvisioningNavigationInfo `json:"navigation"`
	Paging     ProvisioningPagingInfo     `json:"paging"`
}

// UserTemplates User Templates
type UserTemplates struct {
	Uri        string                     `json:"uri"`
	Records    []TemplateInfo             `json:"records"`
	Navigation ProvisioningNavigationInfo `json:"navigation"`
	Paging     ProvisioningPagingInfo     `json:"paging"`
}

// TemplateInfo Template Info
type TemplateInfo struct {
	Uri              string `json:"uri"`
	Id               string `json:"id"`
	Type             string `json:"type"`
	Name             string `json:"name"`
	CreationTime     string `json:"creationTime"`
	LastModifiedTime string `json:"lastModifiedTime"`
}

// ExtensionGrantListEvent Extension Grant List Event
type ExtensionGrantListEvent struct {
	Uuid           string                      `json:"uuid"`
	Event          string                      `json:"@event"`
	Timestamp      string                      `json:"timestamp"`
	SubscriptionId string                      `json:"subscriptionId"`
	Body           ExtensionGrantListEventBody `json:"body"`
}

// ExtensionGrantListEventBody Extension Grant List Event Body
type ExtensionGrantListEventBody struct {
	ExtensionId string `json:"extensionId"`
	OwnerId     string `json:"ownerId"`
}

// ExtensionInfoEvent Extension Info Event
type ExtensionInfoEvent struct {
	Uuid           string                 `json:"uuid"`
	Event          string                 `json:"@event"`
	Timestamp      string                 `json:"timestamp"`
	SubscriptionId string                 `json:"subscriptionId"`
	Body           ExtensionInfoEventBody `json:"body"`
}

// ExtensionInfoEventBody Extension Info Event Body
type ExtensionInfoEventBody struct {
	ExtensionId string   `json:"extensionId"`
	EventType   string   `json:"eventType"`
	Hints       []string `json:"hints"`
	OwnerId     string   `json:"ownerId"`
}

// ExtensionListEvent Extension List Event
type ExtensionListEvent struct {
	Uuid           string                 `json:"uuid"`
	Event          string                 `json:"@event"`
	Timestamp      string                 `json:"timestamp"`
	SubscriptionId string                 `json:"subscriptionId"`
	Body           ExtensionListEventBody `json:"body"`
}

// ExtensionListEventBody Extension List Event Body
type ExtensionListEventBody struct {
	ExtensionId string `json:"extensionId"`
	EventType   string `json:"eventType"`
	OwnerId     string `json:"ownerId"`
}

// OperatorInfo Operator Info
type OperatorInfo struct {
	Id              string `json:"id"`
	Uri             string `json:"uri"`
	ExtensionNumber string `json:"extensionNumber"`
	Name            string `json:"name"`
}

// ShortSiteInfo Short Site Info
type ShortSiteInfo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// SiteInfo Site Info
type SiteInfo struct {
	Id               string                     `json:"id"`
	Uri              string                     `json:"uri"`
	Name             string                     `json:"name"`
	ExtensionNumber  string                     `json:"extensionNumber"`
	CallerIdName     string                     `json:"callerIdName"`
	Email            string                     `json:"email"`
	BusinessAddress  ContactBusinessAddressInfo `json:"businessAddress"`
	RegionalSettings RegionalSettings           `json:"regionalSettings"`
	Operator         OperatorInfo               `json:"@operator"`
	Code             string                     `json:"code"`
}

// UserCallQueues User Call Queues
type UserCallQueues struct {
	Records []QueueShortInfoResource `json:"records"`
}

// QueueShortInfoResource Queue Short Info Resource
type QueueShortInfoResource struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// CreateMultipleSwitchesRequest Create Multiple Switches Request
type CreateMultipleSwitchesRequest struct {
	Records []CreateSwitchInfo `json:"records"`
}

// CreateMultipleSwitchesResponse Create Multiple Switches Response
type CreateMultipleSwitchesResponse struct {
	Task BulkTaskInfoSwUpdate `json:"task"`
}

// BulkTaskInfoSwUpdate Bulk Task Info Sw Update
type BulkTaskInfoSwUpdate struct {
	Id               string `json:"id"`
	Status           string `json:"status"`
	CreationTime     string `json:"creationTime"`
	LastModifiedTime string `json:"lastModifiedTime"`
}

// BulkTaskInfoWiUpdate Bulk Task Info Wi Update
type BulkTaskInfoWiUpdate struct {
	Id               string `json:"id"`
	Status           string `json:"status"`
	CreationTime     string `json:"creationTime"`
	LastModifiedTime string `json:"lastModifiedTime"`
}

// BulkTaskInfoWiCreate Bulk Task Info Wi Create
type BulkTaskInfoWiCreate struct {
	Id               string `json:"id"`
	Status           string `json:"status"`
	CreationTime     string `json:"creationTime"`
	LastModifiedTime string `json:"lastModifiedTime"`
}

// UpdateMultipleSwitchesRequest Update Multiple Switches Request
type UpdateMultipleSwitchesRequest struct {
	Records []UpdateSwitchInfo `json:"records"`
}

// UpdateMultipleSwitchesResponse Update Multiple Switches Response
type UpdateMultipleSwitchesResponse struct {
	Task BulkTaskInfoSwUpdate `json:"task"`
}

// CreateMultipleWirelessPointsRequest Create Multiple Wireless Points Request
type CreateMultipleWirelessPointsRequest struct {
	Records []CreateWirelessPoint `json:"records"`
}

// CreateMultipleWirelessPointsResponse Create Multiple Wireless Points Response
type CreateMultipleWirelessPointsResponse struct {
	Task BulkTaskInfoWiCreate `json:"task"`
}

// UpdateMultipleWirelessPointsRequest Update Multiple Wireless Points Request
type UpdateMultipleWirelessPointsRequest struct {
	Records []UpdateWirelessPoint `json:"records"`
}

// UpdateMultipleWirelessPointsResponse Update Multiple Wireless Points Response
type UpdateMultipleWirelessPointsResponse struct {
	Task BulkTaskInfoWiUpdate `json:"task"`
}

// ValidateMultipleWirelessPointsRequest Validate Multiple Wireless Points Request
type ValidateMultipleWirelessPointsRequest struct {
	Records []WirelessPointInfoRequest `json:"records"`
}

// ValidateMultipleWirelessPointsResponse Validate Multiple Wireless Points Response
type ValidateMultipleWirelessPointsResponse struct {
	Records []WirelessPointValidated `json:"records"`
}

// WirelessPointValidated Wireless Point Validated
type WirelessPointValidated struct {
	Id     string            `json:"id"`
	Bssid  string            `json:"bssid"`
	Status string            `json:"status"`
	Errors []ValidationError `json:"errors"`
}

// ValidationError Validation Error
type ValidationError struct {
	ErrorCode      string `json:"errorCode"`
	Message        string `json:"message"`
	ParameterName  string `json:"parameterName"`
	FeatureName    string `json:"featureName"`
	ParameterValue string `json:"parameterValue"`
}

// ValidateMultipleSwitchesRequest Validate Multiple Switches Request
type ValidateMultipleSwitchesRequest struct {
	Records []SwitchInfo `json:"records"`
}

// ValidateMultipleSwitchesResponse Validate Multiple Switches Response
type ValidateMultipleSwitchesResponse struct {
	Records []SwitchValidated `json:"records"`
}

// SwitchValidated Switch Validated
type SwitchValidated struct {
	Id        string            `json:"id"`
	ChassisId string            `json:"chassisId"`
	Status    string            `json:"status"`
	Errors    []ValidationError `json:"errors"`
}

// UserVideoConfiguration User Video Configuration
type UserVideoConfiguration struct {
	Provider        string `json:"provider"`
	UserLicenseType string `json:"userLicenseType"`
}

// ListDevicesAutomaticLocationUpdates List Devices Automatic Location Updates
type ListDevicesAutomaticLocationUpdates struct {
	Uri        string                               `json:"uri"`
	Records    []AutomaticLocationUpdatesDeviceInfo `json:"records"`
	Navigation ProvisioningNavigationInfo           `json:"navigation"`
	Paging     ProvisioningPagingInfo               `json:"paging"`
}

// AutomaticLocationUpdatesDeviceInfo Automatic Location Updates Device Info
type AutomaticLocationUpdatesDeviceInfo struct {
	Id             string                              `json:"id"`
	Type           string                              `json:"type"`
	Serial         string                              `json:"serial"`
	FeatureEnabled bool                                `json:"featureEnabled"`
	Name           string                              `json:"name"`
	Model          AutomaticLocationUpdatesModelInfo   `json:"model"`
	Site           AutomaticLocationUpdatesSiteInfo    `json:"site"`
	PhoneLines     []AutomaticLocationUpdatesPhoneLine `json:"phoneLines"`
}

// AutomaticLocationUpdatesPhoneLine Automatic Location Updates Phone Line
type AutomaticLocationUpdatesPhoneLine struct {
	LineType  string                                  `json:"lineType"`
	PhoneInfo AutomaticLocationUpdatesPhoneNumberInfo `json:"phoneInfo"`
}

// AutomaticLocationUpdatesPhoneNumberInfo Automatic Location Updates Phone Number Info
type AutomaticLocationUpdatesPhoneNumberInfo struct {
	Id          int    `json:"id"`
	PhoneNumber string `json:"phoneNumber"`
}

// AutomaticLocationUpdatesModelInfo Automatic Location Updates Model Info
type AutomaticLocationUpdatesModelInfo struct {
	Id       string   `json:"id"`
	Name     string   `json:"name"`
	Features []string `json:"features"`
}

// AutomaticLocationUpdatesSiteInfo Automatic Location Updates Site Info
type AutomaticLocationUpdatesSiteInfo struct {
	Id   string `json:"id"`
	Uri  string `json:"uri"`
	Name string `json:"name"`
	Code string `json:"code"`
}

// AssignMultipleDevicesAutomaticLocationUpdates Assign Multiple Devices Automatic Location Updates
type AssignMultipleDevicesAutomaticLocationUpdates struct {
	EnabledDeviceIds  []string `json:"enabledDeviceIds"`
	DisabledDeviceIds []string `json:"disabledDeviceIds"`
}

// NetworksList Networks List
type NetworksList struct {
	Uri        string                     `json:"uri"`
	Records    []NetworkInfo              `json:"records"`
	Navigation ProvisioningNavigationInfo `json:"navigation"`
	Paging     ProvisioningPagingInfo     `json:"paging"`
}

// NetworkInfo Network Info
type NetworkInfo struct {
	Id                string                           `json:"id"`
	Uri               string                           `json:"uri"`
	Name              string                           `json:"name"`
	Site              AutomaticLocationUpdatesSiteInfo `json:"site"`
	PublicIpRanges    []PublicIpRangeInfo              `json:"publicIpRanges"`
	PrivateIpRanges   []PrivateIpRangeInfo             `json:"privateIpRanges"`
	EmergencyLocation ERLLocationInfo                  `json:"emergencyLocation"`
}

// AutomaticLocationUpdatesUserList Automatic Location Updates User List
type AutomaticLocationUpdatesUserList struct {
	Uri        string                             `json:"uri"`
	Records    []AutomaticLocationUpdatesUserInfo `json:"records"`
	Navigation ProvisioningNavigationInfo         `json:"navigation"`
	Paging     ProvisioningPagingInfo             `json:"paging"`
}

// AutomaticLocationUpdatesUserInfo Automatic Location Updates User Info
type AutomaticLocationUpdatesUserInfo struct {
	Id              string                           `json:"id"`
	FullName        string                           `json:"fullName"`
	ExtensionNumber string                           `json:"extensionNumber"`
	FeatureEnabled  bool                             `json:"featureEnabled"`
	Type            string                           `json:"type"`
	Site            AutomaticLocationUpdatesSiteInfo `json:"site"`
	Department      string                           `json:"department"`
}

// AutomaticLocationUpdatesTaskInfo Automatic Location Updates Task Info
type AutomaticLocationUpdatesTaskInfo struct {
	Id               string         `json:"id"`
	Status           string         `json:"status"`
	CreationTime     string         `json:"creationTime"`
	LastModifiedTime string         `json:"lastModifiedTime"`
	Type             string         `json:"type"`
	Result           TaskResultInfo `json:"result"`
}

// TaskResultInfo Task Result Info
type TaskResultInfo struct {
	Records []TaskResultRecord `json:"records"`
}

// TaskResultRecord Task Result Record
type TaskResultRecord struct {
	Id        string                       `json:"id"`
	Bssid     string                       `json:"bssid"`
	ChassisId string                       `json:"chassisId"`
	Status    string                       `json:"status"`
	Errors    []TaskResultRecordErrorsInfo `json:"errors"`
}

// TaskResultRecordErrorsInfo Task Result Record Errors Info
type TaskResultRecordErrorsInfo struct {
	ErrorCode     string `json:"errorCode"`
	Message       string `json:"message"`
	ParameterName string `json:"parameterName"`
	Description   string `json:"description"`
}

// FeatureList Feature List
type FeatureList struct {
	Records []FeatureInfo `json:"records"`
}

// FeatureInfo Feature Info
type FeatureInfo struct {
	Id        string       `json:"id"`
	Available bool         `json:"available"`
	Params    []ParamsInfo `json:"@params"`
	Reason    ReasonInfo   `json:"reason"`
}

// ParamsInfo Params Info
type ParamsInfo struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// ReasonInfo Reason Info
type ReasonInfo struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	Permission string `json:"permission"`
}

// CallQueueDetails Call Queue Details
type CallQueueDetails struct {
	Id                   string                        `json:"id"`
	Name                 string                        `json:"name"`
	ExtensionNumber      string                        `json:"extensionNumber"`
	Status               string                        `json:"status"`
	ServiceLevelSettings CallQueueServiceLevelSettings `json:"serviceLevelSettings"`
	EditableMemberStatus bool                          `json:"editableMemberStatus"`
}

// CallQueueServiceLevelSettings Call Queue Service Level Settings
type CallQueueServiceLevelSettings struct {
	SlaGoal                   int  `json:"slaGoal"`
	SlaThresholdSeconds       int  `json:"slaThresholdSeconds"`
	IncludeAbandonedCalls     bool `json:"includeAbandonedCalls"`
	AbandonedThresholdSeconds int  `json:"abandonedThresholdSeconds"`
}

// CallQueueUpdateDetails Call Queue Update Details
type CallQueueUpdateDetails struct {
	ServiceLevelSettings CallQueueServiceLevelSettings `json:"serviceLevelSettings"`
	EditableMemberStatus bool                          `json:"editableMemberStatus"`
}

// DelegateExtensionInfo Delegate Extension Info
type DelegateExtensionInfo struct {
	Uri             string `json:"uri"`
	Name            string `json:"name"`
	ExtensionNumber string `json:"extensionNumber"`
}

// GetLocationDeletionMultiResponse Get Location Deletion Multi Response
type GetLocationDeletionMultiResponse struct {
	Deletion           string                 `json:"deletion"`
	Reassignment       string                 `json:"reassignment"`
	EmergencyLocations []LocationDeletionInfo `json:"emergencyLocations"`
}

// LocationDeletionInfo Location Deletion Info
type LocationDeletionInfo struct {
	Id       string                    `json:"id"`
	Name     string                    `json:"name"`
	Deletion string                    `json:"deletion"`
	Errors   LocationDeletionErrorInfo `json:"errors"`
}

// LocationDeletionErrorInfo Location Deletion Error Info
type LocationDeletionErrorInfo struct {
	ErrorCode      string `json:"errorCode"`
	Message        string `json:"message"`
	AdditionalInfo string `json:"additionalInfo"`
}

// UserEmergencyLocationList User Emergency Location List
type UserEmergencyLocationList struct {
	Id string `json:"id"`
}

// ExtensionBulkUpdateRequest Extension Bulk Update Request
type ExtensionBulkUpdateRequest struct {
	Records []ExtensionBulkUpdateInfo `json:"records"`
}

// ExtensionBulkUpdateInfo Extension Bulk Update Info
type ExtensionBulkUpdateInfo struct {
	Id               string                           `json:"id"`
	Status           string                           `json:"status"`
	StatusInfo       ExtensionStatusInfo              `json:"statusInfo"`
	Reason           string                           `json:"reason"`
	Comment          string                           `json:"comment"`
	ExtensionNumber  string                           `json:"extensionNumber"`
	Contact          ContactInfoUpdateRequest         `json:"contact"`
	RegionalSettings ExtensionRegionalSettingRequest  `json:"regionalSettings"`
	SetupWizardState string                           `json:"setupWizardState"`
	PartnerId        string                           `json:"partnerId"`
	IvrPin           string                           `json:"ivrPin"`
	Password         string                           `json:"password"`
	CallQueueInfo    CallQueueInfoRequest             `json:"callQueueInfo"`
	Transition       []UserTransitionInfo             `json:"transition"`
	CustomFields     []CustomFieldInfo                `json:"customFields"`
	Hidden           bool                             `json:"hidden"`
	Site             AutomaticLocationUpdatesSiteInfo `json:"site"`
	Type             string                           `json:"type"`
	References       []ReferenceInfo                  `json:"references"`
}

// ExtensionBulkUpdateTaskResource Extension Bulk Update Task Resource
type ExtensionBulkUpdateTaskResource struct {
	Uri              string                        `json:"uri"`
	Id               string                        `json:"id"`
	Status           string                        `json:"status"`
	CreationTime     string                        `json:"creationTime"`
	LastModifiedTime string                        `json:"lastModifiedTime"`
	Result           ExtensionBulkUpdateTaskResult `json:"result"`
}

// ExtensionBulkUpdateTaskResult Extension Bulk Update Task Result
type ExtensionBulkUpdateTaskResult struct {
	AffectedItems []ExtensionUpdateShortResult `json:"affectedItems"`
	Errors        []ErrorEntity                `json:"errors"`
}

// ExtensionUpdateShortResult Extension Update Short Result
type ExtensionUpdateShortResult struct {
	ExtensionId string        `json:"extensionId"`
	Status      string        `json:"status"`
	Errors      []ErrorEntity `json:"errors"`
}

// AddressBookBulkUploadRequest Address Book Bulk Upload Request
type AddressBookBulkUploadRequest struct {
	Records []AddressBookBulkUploadResource `json:"records"`
}

// AddressBookBulkUploadResource Address Book Bulk Upload Resource
type AddressBookBulkUploadResource struct {
	ExtensionId string                           `json:"extensionId"`
	Contacts    []AddressBookBulkContactResource `json:"contacts"`
}

// AddressBookBulkContactResource Address Book Bulk Contact Resource
type AddressBookBulkContactResource struct {
	Email           string                            `json:"email"`
	Notes           string                            `json:"notes"`
	Company         string                            `json:"company"`
	FirstName       string                            `json:"firstName"`
	LastName        string                            `json:"lastName"`
	JobTitle        string                            `json:"jobTitle"`
	Birthday        string                            `json:"birthday"`
	WebPage         string                            `json:"webPage"`
	MiddleName      string                            `json:"middleName"`
	NickName        string                            `json:"nickName"`
	Email2          string                            `json:"email2"`
	Email3          string                            `json:"email3"`
	HomePhone       string                            `json:"homePhone"`
	HomePhone2      string                            `json:"homePhone2"`
	BusinessPhone   string                            `json:"businessPhone"`
	BusinessPhone2  string                            `json:"businessPhone2"`
	MobilePhone     string                            `json:"mobilePhone"`
	BusinessFax     string                            `json:"businessFax"`
	CompanyPhone    string                            `json:"companyPhone"`
	AssistantPhone  string                            `json:"assistantPhone"`
	CarPhone        string                            `json:"carPhone"`
	OtherPhone      string                            `json:"otherPhone"`
	OtherFax        string                            `json:"otherFax"`
	CallbackPhone   string                            `json:"callbackPhone"`
	BusinessAddress AddressBookBulkContactAddressInfo `json:"businessAddress"`
	HomeAddress     AddressBookBulkContactAddressInfo `json:"homeAddress"`
	OtherAddress    AddressBookBulkContactAddressInfo `json:"otherAddress"`
}

// AddressBookBulkContactAddressInfo Address Book Bulk Contact Address Info
type AddressBookBulkContactAddressInfo struct {
	Country string `json:"country"`
	State   string `json:"state"`
	City    string `json:"city"`
	Street  string `json:"street"`
	Zip     string `json:"zip"`
}

// AddressBookBulkUploadResponse Address Book Bulk Upload Response
type AddressBookBulkUploadResponse struct {
	Id               string                          `json:"id"`
	Uri              string                          `json:"uri"`
	Status           string                          `json:"status"`
	CreationTime     string                          `json:"creationTime"`
	LastModifiedTime string                          `json:"lastModifiedTime"`
	Results          AddressBookBulkUploadTaskResult `json:"results"`
}

// AddressBookBulkUploadTaskResult Address Book Bulk Upload Task Result
type AddressBookBulkUploadTaskResult struct {
	ExtensionId string                         `json:"extensionId"`
	Contact     AddressBookBulkContactResource `json:"contact"`
	Status      string                         `json:"status"`
	Errors      []ErrorEntity                  `json:"errors"`
}

// AuthProfileResource Auth Profile Resource
type AuthProfileResource struct {
	Uri         string                     `json:"uri"`
	Permissions []ActivePermissionResource `json:"permissions"`
}

// ActivePermissionResource Active Permission Resource
type ActivePermissionResource struct {
	Permission    PermissionIdResource `json:"permission"`
	EffectiveRole RoleIdResource       `json:"effectiveRole"`
	Scopes        []string             `json:"scopes"`
}

// PermissionIdResource Permission Id Resource
type PermissionIdResource struct {
	Uri            string `json:"uri"`
	Id             string `json:"id"`
	SiteCompatible string `json:"siteCompatible"`
	ReadOnly       bool   `json:"readOnly"`
	Assignable     bool   `json:"assignable"`
}

// RoleIdResource Role Id Resource
type RoleIdResource struct {
	Uri string `json:"uri"`
	Id  string `json:"id"`
}

// AuthProfileCheckResource Auth Profile Check Resource
type AuthProfileCheckResource struct {
	Uri        string                   `json:"uri"`
	Successful bool                     `json:"successful"`
	Details    ActivePermissionResource `json:"details"`
}

// GetRingOutStatusResponse Get Ring Out Status Response
type GetRingOutStatusResponse struct {
	Id     string            `json:"id"`
	Uri    string            `json:"uri"`
	Status RingOutStatusInfo `json:"status"`
}

// MakeRingOutCallerInfoRequestFrom Make Ring Out Caller Info Request From
type MakeRingOutCallerInfoRequestFrom struct {
	PhoneNumber        string `json:"phoneNumber"`
	ForwardingNumberId string `json:"forwardingNumberId"`
}

// MakeRingOutCallerInfoRequestTo Make Ring Out Caller Info Request To
type MakeRingOutCallerInfoRequestTo struct {
	PhoneNumber string `json:"phoneNumber"`
}

// MakeRingOutCallerIdInfo Make Ring Out Caller Id Info
type MakeRingOutCallerIdInfo struct {
	PhoneNumber string `json:"phoneNumber"`
}

// MakeRingOutCoutryInfo Make Ring Out Coutry Info
type MakeRingOutCoutryInfo struct {
	Id string `json:"id"`
}

// MakeRingOutRequest Make Ring Out Request
type MakeRingOutRequest struct {
	From       MakeRingOutCallerInfoRequestFrom `json:"from"`
	To         MakeRingOutCallerInfoRequestTo   `json:"to"`
	CallerId   MakeRingOutCallerIdInfo          `json:"callerId"`
	PlayPrompt bool                             `json:"playPrompt"`
	Country    MakeRingOutCoutryInfo            `json:"country"`
}

// RingOutStatusInfo Ring Out Status Info
type RingOutStatusInfo struct {
	CallStatus   string `json:"callStatus"`
	CallerStatus string `json:"callerStatus"`
	CalleeStatus string `json:"calleeStatus"`
}

// UserAddress User Address
type UserAddress struct {
	Country       string `json:"country"`
	Locality      string `json:"locality"`
	PostalCode    string `json:"postalCode"`
	Region        string `json:"region"`
	StreetAddress string `json:"streetAddress"`
	Type          string `json:"type"`
}

// AuthenticationScheme Authentication Scheme
type AuthenticationScheme struct {
	Description      string `json:"description"`
	DocumentationUri string `json:"documentationUri"`
	Name             string `json:"name"`
	SpecUri          string `json:"specUri"`
	Primary          bool   `json:"primary"`
}

// BulkSupported Bulk Supported
type BulkSupported struct {
	MaxOperations  int  `json:"maxOperations"`
	MaxPayloadSize int  `json:"maxPayloadSize"`
	Supported      bool `json:"supported"`
}

// Email Email
type Email struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// EnterpriseUser Enterprise User
type EnterpriseUser struct {
	Department string `json:"department"`
}

// ScimErrorResponse Scim Error Response
type ScimErrorResponse struct {
	Detail   string   `json:"detail"`
	Schemas  []string `json:"schemas"`
	ScimType string   `json:"scimType"`
	Status   string   `json:"status"`
}

// FilterSupported Filter Supported
type FilterSupported struct {
	MaxResults int  `json:"maxResults"`
	Supported  bool `json:"supported"`
}

// Meta Meta
type Meta struct {
	Created      string `json:"created"`
	LastModified string `json:"lastModified"`
	Location     string `json:"location"`
	ResourceType string `json:"resourceType"`
}

// Name Name
type Name struct {
	FamilyName string `json:"familyName"`
	GivenName  string `json:"givenName"`
}

// PatchOperation Patch Operation
type PatchOperation struct {
	Op    string `json:"op"`
	Path  string `json:"path"`
	Value string `json:"value"`
}

// PhoneNumber Phone Number
type PhoneNumber struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// Photo Photo
type Photo struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// SearchRequest Search Request
type SearchRequest struct {
	Count      int      `json:"count"`
	Filter     string   `json:"filter"`
	Schemas    []string `json:"schemas"`
	StartIndex int      `json:"startIndex"`
}

// ServiceProviderConfig Service Provider Config
type ServiceProviderConfig struct {
	AuthenticationSchemes []AuthenticationScheme `json:"authenticationSchemes"`
	Bulk                  BulkSupported          `json:"bulk"`
	ChangePassword        Supported              `json:"changePassword"`
	Etag                  Supported              `json:"etag"`
	Filter                FilterSupported        `json:"filter"`
	Patch                 Supported              `json:"patch"`
	Schemas               []string               `json:"schemas"`
	Sort                  Supported              `json:"sort"`
	XmlDataFormat         Supported              `json:"xmlDataFormat"`
}

// Supported Supported
type Supported struct {
	Supported bool `json:"supported"`
}

// CreateUser Create User
type CreateUser struct {
	Active                                              bool           `json:"active"`
	Addresses                                           []UserAddress  `json:"addresses"`
	Emails                                              []Email        `json:"emails"`
	ExternalId                                          string         `json:"externalId"`
	Name                                                Name           `json:"name"`
	PhoneNumbers                                        []PhoneNumber  `json:"phoneNumbers"`
	Photos                                              []Photo        `json:"photos"`
	Schemas                                             []string       `json:"schemas"`
	UrnIetfParamsScimSchemasExtensionEnterprise_2_0User EnterpriseUser `json:"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"`
	UserName                                            string         `json:"userName"`
}

// User User
type User struct {
	Active                                              bool           `json:"active"`
	Addresses                                           []UserAddress  `json:"addresses"`
	Emails                                              []Email        `json:"emails"`
	ExternalId                                          string         `json:"externalId"`
	Id                                                  string         `json:"id"`
	Name                                                Name           `json:"name"`
	PhoneNumbers                                        []PhoneNumber  `json:"phoneNumbers"`
	Photos                                              []Photo        `json:"photos"`
	Schemas                                             []string       `json:"schemas"`
	UrnIetfParamsScimSchemasExtensionEnterprise_2_0User EnterpriseUser `json:"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"`
	UserName                                            string         `json:"userName"`
}

// UserPatch User Patch
type UserPatch struct {
	Operations []PatchOperation `json:"Operations"`
	Schemas    []string         `json:"schemas"`
}

// UserResponse User Response
type UserResponse struct {
	Active                                              bool           `json:"active"`
	Addresses                                           []UserAddress  `json:"addresses"`
	Emails                                              []Email        `json:"emails"`
	ExternalId                                          string         `json:"externalId"`
	Id                                                  string         `json:"id"`
	Name                                                Name           `json:"name"`
	PhoneNumbers                                        []PhoneNumber  `json:"phoneNumbers"`
	Photos                                              []Photo        `json:"photos"`
	Schemas                                             []string       `json:"schemas"`
	UrnIetfParamsScimSchemasExtensionEnterprise_2_0User EnterpriseUser `json:"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"`
	UserName                                            string         `json:"userName"`
	Meta                                                Meta           `json:"meta"`
}

// UserSearchResponse User Search Response
type UserSearchResponse struct {
	Resources    []UserResponse `json:"Resources"`
	ItemsPerPage int            `json:"itemsPerPage"`
	Schemas      []string       `json:"schemas"`
	StartIndex   int            `json:"startIndex"`
	TotalResults int            `json:"totalResults"`
}

// BridgeTargetRequest Bridge Target Request
type BridgeTargetRequest struct {
	TelephonySessionId string `json:"telephonySessionId"`
	PartyId            string `json:"partyId"`
}

// PartySuperviseRequest Party Supervise Request
type PartySuperviseRequest struct {
	Mode               string `json:"mode"`
	SupervisorDeviceId string `json:"supervisorDeviceId"`
	AgentExtensionId   string `json:"agentExtensionId"`
}

// PartySuperviseResponse Party Supervise Response
type PartySuperviseResponse struct {
	From        PartyInfo      `json:"from"`
	To          PartyInfo      `json:"to"`
	Direction   string         `json:"direction"`
	Id          string         `json:"id"`
	AccountId   string         `json:"accountId"`
	ExtensionId string         `json:"extensionId"`
	Muted       bool           `json:"muted"`
	Owner       OwnerInfo      `json:"owner"`
	StandAlone  bool           `json:"standAlone"`
	Status      CallStatusInfo `json:"status"`
}

// CallSession Call Session
type CallSession struct {
	Session CallSessionObject `json:"session"`
}

// CallSessionObject Call Session Object
type CallSessionObject struct {
	Id             string      `json:"id"`
	Origin         OriginInfo  `json:"origin"`
	VoiceCallToken string      `json:"voiceCallToken"`
	Parties        []CallParty `json:"parties"`
	CreationTime   string      `json:"creationTime"`
}

// OriginInfo Origin Info
type OriginInfo struct {
	Type string `json:"type"`
}

// CallParty Call Party
type CallParty struct {
	Id             string          `json:"id"`
	Status         CallStatusInfo  `json:"status"`
	Muted          bool            `json:"muted"`
	StandAlone     bool            `json:"standAlone"`
	Park           ParkInfo        `json:"park"`
	From           PartyInfo       `json:"from"`
	To             PartyInfo       `json:"to"`
	Owner          OwnerInfo       `json:"owner"`
	Direction      string          `json:"direction"`
	ConferenceRole string          `json:"conferenceRole"`
	RingOutRole    string          `json:"ringOutRole"`
	RingMeRole     string          `json:"ringMeRole"`
	Recordings     []RecordingInfo `json:"recordings"`
}

// RecordingInfo Recording Info
type RecordingInfo struct {
	Id     string `json:"id"`
	Active bool   `json:"active"`
}

// ParkInfo Park Info
type ParkInfo struct {
	Id string `json:"id"`
}

// CallStatusInfo Call Status Info
type CallStatusInfo struct {
	Code        string   `json:"code"`
	PeerId      PeerInfo `json:"peerId"`
	Reason      string   `json:"reason"`
	Description string   `json:"description"`
}

// PeerInfo Peer Info
type PeerInfo struct {
	SessionId          string `json:"sessionId"`
	TelephonySessionId string `json:"telephonySessionId"`
	PartyId            string `json:"partyId"`
}

// PartyInfo Party Info
type PartyInfo struct {
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
	DeviceId    string `json:"deviceId"`
	ExtensionId string `json:"extensionId"`
}

// SupervisePartyTo Supervise Party To
type SupervisePartyTo struct {
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
	DeviceId    string `json:"deviceId"`
	ExtensionId string `json:"extensionId"`
}

// SupervisePartyFrom Supervise Party From
type SupervisePartyFrom struct {
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
	DeviceId    string `json:"deviceId"`
	ExtensionId string `json:"extensionId"`
}

// OwnerInfo Owner Info
type OwnerInfo struct {
	AccountId   string `json:"accountId"`
	ExtensionId string `json:"extensionId"`
}

// PartyUpdateRequest Party Update Request
type PartyUpdateRequest struct {
	Party PartyUpdateInfo `json:"party"`
}

// PartyUpdateInfo Party Update Info
type PartyUpdateInfo struct {
	Muted      bool `json:"muted"`
	StandAlone bool `json:"standAlone"`
}

// AddPartyRequest Add Party Request
type AddPartyRequest struct {
	TelephonySessionId string `json:"telephonySessionId"`
	PartyId            string `json:"partyId"`
}

// AnswerTarget Answer Target
type AnswerTarget struct {
	DeviceId string `json:"deviceId"`
}

// PickupTarget Pickup Target
type PickupTarget struct {
	DeviceId string `json:"deviceId"`
}

// ForwardTarget Forward Target
type ForwardTarget struct {
	PhoneNumber string `json:"phoneNumber"`
	Voicemail   string `json:"voicemail"`
}

// TransferTarget Transfer Target
type TransferTarget struct {
	PhoneNumber string `json:"phoneNumber"`
	Voicemail   string `json:"voicemail"`
	ParkOrbit   string `json:"parkOrbit"`
}

// CallPartyReply Call Party Reply
type CallPartyReply struct {
	ReplyWithText    string           `json:"replyWithText"`
	ReplyWithPattern ReplyWithPattern `json:"replyWithPattern"`
}

// ReplyWithPattern Reply With Pattern
type ReplyWithPattern struct {
	Pattern  string `json:"pattern"`
	Time     int    `json:"time"`
	TimeUnit string `json:"timeUnit"`
}

// CallPartyFlip Call Party Flip
type CallPartyFlip struct {
	CallFlipId string `json:"callFlipId"`
}

// CallRecordingUpdate Call Recording Update
type CallRecordingUpdate struct {
	Active bool `json:"active"`
}

// CallRecording Call Recording
type CallRecording struct {
	Id     string `json:"id"`
	Active bool   `json:"active"`
}

// MakeCallOutCallerInfoRequestFrom Make Call Out Caller Info Request From
type MakeCallOutCallerInfoRequestFrom struct {
	DeviceId string `json:"deviceId"`
}

// MakeCallOutCallerInfoRequestTo Make Call Out Caller Info Request To
type MakeCallOutCallerInfoRequestTo struct {
	PhoneNumber     string `json:"phoneNumber"`
	ExtensionNumber string `json:"extensionNumber"`
}

// MakeCallOutRequest Make Call Out Request
type MakeCallOutRequest struct {
	From MakeCallOutCallerInfoRequestFrom `json:"from"`
	To   MakeCallOutCallerInfoRequestTo   `json:"to"`
}

// SuperviseCallSessionRequest Supervise Call Session Request
type SuperviseCallSessionRequest struct {
	Mode                 string `json:"mode"`
	SupervisorDeviceId   string `json:"supervisorDeviceId"`
	AgentExtensionNumber string `json:"agentExtensionNumber"`
	AgentExtensionId     string `json:"agentExtensionId"`
}

// SuperviseCallSession Supervise Call Session
type SuperviseCallSession struct {
	From        SupervisePartyFrom `json:"from"`
	To          SupervisePartyTo   `json:"to"`
	Direction   string             `json:"direction"`
	Id          string             `json:"id"`
	AccountId   string             `json:"accountId"`
	ExtensionId string             `json:"extensionId"`
	Muted       bool               `json:"muted"`
	Owner       OwnerInfo          `json:"owner"`
	StandAlone  bool               `json:"standAlone"`
	Status      CallStatusInfo     `json:"status"`
}

// IgnoreRequestBody Ignore Request Body
type IgnoreRequestBody struct {
	DeviceId string `json:"deviceId"`
}

// ReplyParty Reply Party
type ReplyParty struct {
	Id         string         `json:"id"`
	Status     CallStatusInfo `json:"status"`
	Muted      bool           `json:"muted"`
	StandAlone bool           `json:"standAlone"`
	Park       ParkInfo       `json:"park"`
	From       PartyInfo      `json:"from"`
	To         PartyInfo      `json:"to"`
	Owner      OwnerInfo      `json:"owner"`
	Direction  string         `json:"direction"`
}

// AccountTelephonySessionsEvent Account Telephony Sessions Event
type AccountTelephonySessionsEvent struct {
	Uuid           string                     `json:"uuid"`
	Event          string                     `json:"@event"`
	Timestamp      string                     `json:"timestamp"`
	SubscriptionId string                     `json:"subscriptionId"`
	OwnerId        string                     `json:"ownerId"`
	Body           TelephonySessionsEventBody `json:"body"`
}

// ExtensionTelephonySessionsEvent Extension Telephony Sessions Event
type ExtensionTelephonySessionsEvent struct {
	Uuid           string                     `json:"uuid"`
	Event          string                     `json:"@event"`
	Timestamp      string                     `json:"timestamp"`
	SubscriptionId string                     `json:"subscriptionId"`
	OwnerId        string                     `json:"ownerId"`
	Body           TelephonySessionsEventBody `json:"body"`
}

// TelephonySessionsEventBody Telephony Sessions Event Body
type TelephonySessionsEventBody struct {
	Sequence           int                               `json:"sequence"`
	SessionId          string                            `json:"sessionId"`
	TelephonySessionId string                            `json:"telephonySessionId"`
	ServerId           string                            `json:"serverId"`
	EventTime          string                            `json:"eventTime"`
	Origin             OriginInfo                        `json:"origin"`
	Parties            []TelephonySessionsEventPartyInfo `json:"parties"`
}

// TelephonySessionsEventPartyInfo Telephony Sessions Event Party Info
type TelephonySessionsEventPartyInfo struct {
	AccountId   string                `json:"accountId"`
	ExtensionId string                `json:"extensionId"`
	Id          string                `json:"id"`
	Direction   string                `json:"direction"`
	To          CallPartyInfo         `json:"to"`
	From        CallPartyInfo         `json:"from"`
	Status      CallSessionStatusInfo `json:"status"`
}

// CallPartyInfo Call Party Info
type CallPartyInfo struct {
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
	ExtensionId string `json:"extensionId"`
}

// CallSessionStatusInfo Call Session Status Info
type CallSessionStatusInfo struct {
	Code             string           `json:"code"`
	Reason           string           `json:"reason"`
	ParkData         string           `json:"parkData"`
	PeerId           PeerInfo         `json:"peerId"`
	MobilePickupData MobilePickupData `json:"mobilePickupData"`
}

// MobilePickupData Mobile Pickup Data
type MobilePickupData struct {
	CcMailboxes []string `json:"ccMailboxes"`
	To          string   `json:"to"`
	Sid         string   `json:"sid"`
	Srvlvl      string   `json:"srvlvl"`
	SrvLvlExt   string   `json:"srvLvlExt"`
}

// MessageStoreCalleeInfoRequest Message Store Callee Info Request
type MessageStoreCalleeInfoRequest struct {
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
}

// CustomGreetingAnsweringRuleInfoRequest Custom Greeting Answering Rule Info Request
type CustomGreetingAnsweringRuleInfoRequest struct {
	Id string `json:"id"`
}

// CustomCompanyGreetingAnsweringRuleInfo Custom Company Greeting Answering Rule Info
type CustomCompanyGreetingAnsweringRuleInfo struct {
	Id string `json:"id"`
}

// AuthorizeRequest Authorize Request
type AuthorizeRequest struct {
	ResponseType        string `json:"response_type"`
	RedirectUri         string `json:"redirect_uri"`
	ClientId            string `json:"client_id"`
	State               string `json:"state"`
	BrandId             string `json:"brand_id"`
	Display             string `json:"display"`
	Prompt              string `json:"prompt"`
	LocaleId            string `json:"localeId"`
	UiLocales           string `json:"ui_locales"`
	UiOptions           string `json:"ui_options"`
	Scope               string `json:"scope"`
	AcceptLanguage      string `json:"accept_language"`
	Request             string `json:"request"`
	RequestUri          string `json:"request_uri"`
	Nonce               string `json:"nonce"`
	CodeChallenge       string `json:"code_challenge"`
	CodeChallengeMethod string `json:"code_challenge_method"`
	Discovery           bool   `json:"discovery"`
}

// RevokeTokenRequest Revoke Token Request
type RevokeTokenRequest struct {
	ClientAssertionType string `json:"client_assertion_type"`
	ClientAssertion     string `json:"client_assertion"`
	Token               string `json:"token"`
}

// GetTokenRequest Get Token Request
type GetTokenRequest struct {
	Username            string `json:"username"`
	Password            string `json:"password"`
	Extension           string `json:"extension"`
	GrantType           string `json:"grant_type"`
	Code                string `json:"code"`
	RedirectUri         string `json:"redirect_uri"`
	AccessTokenTtl      int    `json:"access_token_ttl"`
	RefreshTokenTtl     int    `json:"refresh_token_ttl"`
	Scope               string `json:"scope"`
	RefreshToken        string `json:"refresh_token"`
	EndpointId          string `json:"endpoint_id"`
	Pin                 string `json:"pin"`
	ClientId            string `json:"client_id"`
	AccountId           string `json:"account_id"`
	PartnerAccountId    string `json:"partner_account_id"`
	ClientAssertionType string `json:"client_assertion_type"`
	ClientAssertion     string `json:"client_assertion"`
	Assertion           string `json:"assertion"`
	BrandId             string `json:"brand_id"`
	CodeVerifier        string `json:"code_verifier"`
}

// CreateCompanyGreetingRequest Create Company Greeting Request
type CreateCompanyGreetingRequest struct {
	Type            string     `json:"type"`
	AnsweringRuleId string     `json:"answeringRuleId"`
	LanguageId      string     `json:"languageId"`
	Binary          Attachment `json:"binary"`
}

// CreateCustomUserGreetingRequest Create Custom User Greeting Request
type CreateCustomUserGreetingRequest struct {
	Type            string     `json:"type"`
	AnsweringRuleId string     `json:"answeringRuleId"`
	Binary          Attachment `json:"binary"`
}

// CreateIvrPromptRequest Create Ivr Prompt Request
type CreateIvrPromptRequest struct {
	Attachment Attachment `json:"attachment"`
	Name       string     `json:"name"`
}

// CreateFaxMessageRequest Create Fax Message Request
type CreateFaxMessageRequest struct {
	FaxResolution string                          `json:"faxResolution"`
	To            []MessageStoreCalleeInfoRequest `json:"to"`
	SendTime      string                          `json:"sendTime"`
	IsoCode       string                          `json:"isoCode"`
	CoverIndex    int                             `json:"coverIndex"`
	CoverPageText string                          `json:"coverPageText"`
	Attachments   []Attachment                    `json:"attachments"`
}

// UpdateUserProfileImageRequest Update User Profile Image Request
type UpdateUserProfileImageRequest struct {
	Image Attachment `json:"image"`
}

// CreateUserProfileImageRequest Create User Profile Image Request
type CreateUserProfileImageRequest struct {
	Image Attachment `json:"image"`
}

// ReadUserPresenceStatusParameters Read User Presence Status Parameters
type ReadUserPresenceStatusParameters struct {
	DetailedTelephonyState bool `json:"detailedTelephonyState"`
	SipData                bool `json:"sipData"`
}

// ReadAccountPresenceParameters Read Account Presence Parameters
type ReadAccountPresenceParameters struct {
	DetailedTelephonyState bool `json:"detailedTelephonyState"`
	SipData                bool `json:"sipData"`
	Page                   int  `json:"page"`
	PerPage                int  `json:"perPage"`
}

// ReadExtensionCallQueuePresenceParameters Read Extension Call Queue Presence Parameters
type ReadExtensionCallQueuePresenceParameters struct {
	EditableMemberStatus bool `json:"editableMemberStatus"`
}

// UpdateSubscriptionParameters Update Subscription Parameters
type UpdateSubscriptionParameters struct {
	Aggregated bool `json:"aggregated"`
}

// ListDataExportTasksParameters List Data Export Tasks Parameters
type ListDataExportTasksParameters struct {
	Status  string `json:"status"`
	Page    int    `json:"page"`
	PerPage int    `json:"perPage"`
}

// ListGlipChatsParameters List Glip Chats Parameters
type ListGlipChatsParameters struct {
	Type        []string `json:"type"`
	RecordCount int      `json:"recordCount"`
	PageToken   string   `json:"pageToken"`
}

// ListGlipConversationsParameters List Glip Conversations Parameters
type ListGlipConversationsParameters struct {
	RecordCount int    `json:"recordCount"`
	PageToken   string `json:"pageToken"`
}

// ListGlipTeamsParameters List Glip Teams Parameters
type ListGlipTeamsParameters struct {
	RecordCount int    `json:"recordCount"`
	PageToken   string `json:"pageToken"`
}

// ListRecentChatsParameters List Recent Chats Parameters
type ListRecentChatsParameters struct {
	Type        []string `json:"type"`
	RecordCount int      `json:"recordCount"`
}

// ListFavoriteChatsParameters List Favorite Chats Parameters
type ListFavoriteChatsParameters struct {
	RecordCount int `json:"recordCount"`
}

// ReadGlipPostsParameters Read Glip Posts Parameters
type ReadGlipPostsParameters struct {
	RecordCount int    `json:"recordCount"`
	PageToken   string `json:"pageToken"`
}

// ReadGlipEventsParameters Read Glip Events Parameters
type ReadGlipEventsParameters struct {
	RecordCount int    `json:"recordCount"`
	PageToken   string `json:"pageToken"`
}

// ListChatNotesParameters List Chat Notes Parameters
type ListChatNotesParameters struct {
	CreationTimeTo   string `json:"creationTimeTo"`
	CreationTimeFrom string `json:"creationTimeFrom"`
	CreatorId        string `json:"creatorId"`
	Status           string `json:"status"`
	PageToken        string `json:"pageToken"`
	RecordCount      int    `json:"recordCount"`
}

// ListChatTasksParameters List Chat Tasks Parameters
type ListChatTasksParameters struct {
	CreationTimeTo   string   `json:"creationTimeTo"`
	CreationTimeFrom string   `json:"creationTimeFrom"`
	CreatorId        []string `json:"creatorId"`
	Status           []string `json:"status"`
	AssignmentStatus string   `json:"assignmentStatus"`
	AssigneeId       []string `json:"assigneeId"`
	AssigneeStatus   string   `json:"assigneeStatus"`
	PageToken        string   `json:"pageToken"`
	RecordCount      int      `json:"recordCount"`
}

// ListContactsParameters List Contacts Parameters
type ListContactsParameters struct {
	StartsWith  string   `json:"startsWith"`
	SortBy      []string `json:"sortBy"`
	Page        int      `json:"page"`
	PerPage     int      `json:"perPage"`
	PhoneNumber []string `json:"phoneNumber"`
}

// CreateContactParameters Create Contact Parameters
type CreateContactParameters struct {
	DialingPlan string `json:"dialingPlan"`
}

// UpdateContactParameters Update Contact Parameters
type UpdateContactParameters struct {
	DialingPlan string `json:"dialingPlan"`
}

// SyncAddressBookParameters Sync Address Book Parameters
type SyncAddressBookParameters struct {
	SyncType  string `json:"syncType"`
	SyncToken string `json:"syncToken"`
	PerPage   int    `json:"perPage"`
	PageId    int    `json:"pageId"`
}

// ListDirectoryEntriesParameters List Directory Entries Parameters
type ListDirectoryEntriesParameters struct {
	ShowFederated bool   `json:"showFederated"`
	Type          string `json:"type"`
	Page          string `json:"page"`
	PerPage       int    `json:"perPage"`
	SiteId        string `json:"siteId"`
}

// ListBlockedAllowedNumbersParameters List Blocked Allowed Numbers Parameters
type ListBlockedAllowedNumbersParameters struct {
	Page    int    `json:"page"`
	PerPage int    `json:"perPage"`
	Status  string `json:"status"`
}

// ListForwardingNumbersParameters List Forwarding Numbers Parameters
type ListForwardingNumbersParameters struct {
	Page    int `json:"page"`
	PerPage int `json:"perPage"`
}

// ListAnsweringRulesParameters List Answering Rules Parameters
type ListAnsweringRulesParameters struct {
	View        string `json:"view"`
	EnabledOnly bool   `json:"enabledOnly"`
	Page        string `json:"page"`
	PerPage     string `json:"perPage"`
}

// ReadAnsweringRuleParameters Read Answering Rule Parameters
type ReadAnsweringRuleParameters struct {
	ShowInactiveNumbers bool `json:"showInactiveNumbers"`
}

// ListCompanyAnsweringRulesParameters List Company Answering Rules Parameters
type ListCompanyAnsweringRulesParameters struct {
	Page    int `json:"page"`
	PerPage int `json:"perPage"`
}

// ListStandardGreetingsParameters List Standard Greetings Parameters
type ListStandardGreetingsParameters struct {
	Page      int    `json:"page"`
	PerPage   int    `json:"perPage"`
	Type      string `json:"type"`
	UsageType string `json:"usageType"`
}

// CreateCustomUserGreetingParameters Create Custom User Greeting Parameters
type CreateCustomUserGreetingParameters struct {
	Apply bool `json:"apply"`
}

// ListCallRecordingCustomGreetingsParameters List Call Recording Custom Greetings Parameters
type ListCallRecordingCustomGreetingsParameters struct {
	Type string `json:"type"`
}

// ListA2PsmsParameters List A2Psms Parameters
type ListA2PsmsParameters struct {
	BatchId     string   `json:"batchId"`
	Direction   string   `json:"direction"`
	View        string   `json:"view"`
	DateFrom    string   `json:"dateFrom"`
	DateTo      string   `json:"dateTo"`
	PhoneNumber []string `json:"phoneNumber"`
	PageToken   string   `json:"pageToken"`
	PerPage     int      `json:"perPage"`
}

// ReadA2PsmsOptOutsParameters Read A2Psms Opt Outs Parameters
type ReadA2PsmsOptOutsParameters struct {
	From      string `json:"from"`
	To        string `json:"to"`
	PageToken string `json:"pageToken"`
	PerPage   int    `json:"perPage"`
}

// ReadUserCallLogParameters Read User Call Log Parameters
type ReadUserCallLogParameters struct {
	ExtensionNumber string   `json:"extensionNumber"`
	ShowBlocked     bool     `json:"showBlocked"`
	PhoneNumber     string   `json:"phoneNumber"`
	Direction       []string `json:"direction"`
	SessionId       string   `json:"sessionId"`
	Type            []string `json:"type"`
	Transport       []string `json:"transport"`
	View            string   `json:"view"`
	WithRecording   bool     `json:"withRecording"`
	RecordingType   string   `json:"recordingType"`
	DateTo          string   `json:"dateTo"`
	DateFrom        string   `json:"dateFrom"`
	Page            int      `json:"page"`
	PerPage         int      `json:"perPage"`
	ShowDeleted     bool     `json:"showDeleted"`
}

// DeleteUserCallLogParameters Delete User Call Log Parameters
type DeleteUserCallLogParameters struct {
	DateTo          string   `json:"dateTo"`
	PhoneNumber     string   `json:"phoneNumber"`
	ExtensionNumber string   `json:"extensionNumber"`
	Type            []string `json:"type"`
	Direction       []string `json:"direction"`
	DateFrom        string   `json:"dateFrom"`
}

// SyncUserCallLogParameters Sync User Call Log Parameters
type SyncUserCallLogParameters struct {
	SyncType    string   `json:"syncType"`
	SyncToken   string   `json:"syncToken"`
	DateFrom    string   `json:"dateFrom"`
	RecordCount int      `json:"recordCount"`
	StatusGroup []string `json:"statusGroup"`
	View        string   `json:"view"`
	ShowDeleted bool     `json:"showDeleted"`
}

// ReadUserCallRecordParameters Read User Call Record Parameters
type ReadUserCallRecordParameters struct {
	View string `json:"view"`
}

// ListExtensionActiveCallsParameters List Extension Active Calls Parameters
type ListExtensionActiveCallsParameters struct {
	Direction []string `json:"direction"`
	View      string   `json:"view"`
	Type      []string `json:"type"`
	Page      int      `json:"page"`
	PerPage   int      `json:"perPage"`
}

// ReadCompanyCallLogParameters Read Company Call Log Parameters
type ReadCompanyCallLogParameters struct {
	ExtensionNumber string   `json:"extensionNumber"`
	PhoneNumber     string   `json:"phoneNumber"`
	Direction       []string `json:"direction"`
	Type            []string `json:"type"`
	View            string   `json:"view"`
	WithRecording   bool     `json:"withRecording"`
	RecordingType   string   `json:"recordingType"`
	DateFrom        string   `json:"dateFrom"`
	DateTo          string   `json:"dateTo"`
	Page            int      `json:"page"`
	PerPage         int      `json:"perPage"`
	SessionId       string   `json:"sessionId"`
}

// SyncAccountCallLogParameters Sync Account Call Log Parameters
type SyncAccountCallLogParameters struct {
	SyncType    string   `json:"syncType"`
	SyncToken   string   `json:"syncToken"`
	DateFrom    string   `json:"dateFrom"`
	RecordCount int      `json:"recordCount"`
	StatusGroup []string `json:"statusGroup"`
	View        string   `json:"view"`
	ShowDeleted bool     `json:"showDeleted"`
}

// ReadCompanyCallRecordParameters Read Company Call Record Parameters
type ReadCompanyCallRecordParameters struct {
	View string `json:"view"`
}

// ListCompanyActiveCallsParameters List Company Active Calls Parameters
type ListCompanyActiveCallsParameters struct {
	Direction []string `json:"direction"`
	View      string   `json:"view"`
	Type      []string `json:"type"`
	Transport []string `json:"transport"`
	Page      int      `json:"page"`
	PerPage   int      `json:"perPage"`
}

// ReadDeviceParameters Read Device Parameters
type ReadDeviceParameters struct {
	SyncEmergencyAddress bool `json:"syncEmergencyAddress"`
}

// UpdateDeviceParameters Update Device Parameters
type UpdateDeviceParameters struct {
	Prestatement bool `json:"prestatement"`
}

// ListExtensionDevicesParameters List Extension Devices Parameters
type ListExtensionDevicesParameters struct {
	LinePooling string `json:"linePooling"`
	Feature     string `json:"feature"`
}

// ListAccountMeetingRecordingsParameters List Account Meeting Recordings Parameters
type ListAccountMeetingRecordingsParameters struct {
	MeetingId            string `json:"meetingId"`
	MeetingStartTimeFrom string `json:"meetingStartTimeFrom"`
	MeetingStartTimeTo   string `json:"meetingStartTimeTo"`
	Page                 int    `json:"page"`
	PerPage              int    `json:"perPage"`
}

// ListUserMeetingRecordingsParameters List User Meeting Recordings Parameters
type ListUserMeetingRecordingsParameters struct {
	MeetingId            string `json:"meetingId"`
	MeetingStartTimeFrom string `json:"meetingStartTimeFrom"`
	MeetingStartTimeTo   string `json:"meetingStartTimeTo"`
	Page                 int    `json:"page"`
	PerPage              int    `json:"perPage"`
}

// ListFaxCoverPagesParameters List Fax Cover Pages Parameters
type ListFaxCoverPagesParameters struct {
	Page    int `json:"page"`
	PerPage int `json:"perPage"`
}

// ListMessagesParameters List Messages Parameters
type ListMessagesParameters struct {
	Availability          []string `json:"availability"`
	ConversationId        int      `json:"conversationId"`
	DateFrom              string   `json:"dateFrom"`
	DateTo                string   `json:"dateTo"`
	Direction             []string `json:"direction"`
	DistinctConversations bool     `json:"distinctConversations"`
	MessageType           []string `json:"messageType"`
	ReadStatus            []string `json:"readStatus"`
	Page                  int      `json:"page"`
	PerPage               int      `json:"perPage"`
	PhoneNumber           string   `json:"phoneNumber"`
}

// DeleteMessageByFilterParameters Delete Message By Filter Parameters
type DeleteMessageByFilterParameters struct {
	ConversationId []string `json:"conversationId"`
	DateTo         string   `json:"dateTo"`
	Type           string   `json:"type"`
}

// UpdateMessageParameters Update Message Parameters
type UpdateMessageParameters struct {
	DateFrom string `json:"dateFrom"`
	Type     string `json:"type"`
}

// DeleteMessageParameters Delete Message Parameters
type DeleteMessageParameters struct {
	Purge          bool `json:"purge"`
	ConversationId int  `json:"conversationId"`
}

// ReadMessageContentParameters Read Message Content Parameters
type ReadMessageContentParameters struct {
	ContentDisposition string `json:"contentDisposition"`
}

// SyncMessagesParameters Sync Messages Parameters
type SyncMessagesParameters struct {
	ConversationId        int      `json:"conversationId"`
	DateFrom              string   `json:"dateFrom"`
	DateTo                string   `json:"dateTo"`
	Direction             []string `json:"direction"`
	DistinctConversations bool     `json:"distinctConversations"`
	MessageType           []string `json:"messageType"`
	RecordCount           int      `json:"recordCount"`
	SyncToken             string   `json:"syncToken"`
	SyncType              string   `json:"syncType"`
}

// ParsePhoneNumberParameters Parse Phone Number Parameters
type ParsePhoneNumberParameters struct {
	HomeCountry        string `json:"homeCountry"`
	NationalAsPriority bool   `json:"nationalAsPriority"`
}

// ListExtensionPhoneNumbersParameters List Extension Phone Numbers Parameters
type ListExtensionPhoneNumbersParameters struct {
	Status    string   `json:"status"`
	UsageType []string `json:"usageType"`
	Page      int      `json:"page"`
	PerPage   int      `json:"perPage"`
}

// DeleteExtensionParameters Delete Extension Parameters
type DeleteExtensionParameters struct {
	SavePhoneLines   bool `json:"savePhoneLines"`
	SavePhoneNumbers bool `json:"savePhoneNumbers"`
}

// ListExtensionGrantsParameters List Extension Grants Parameters
type ListExtensionGrantsParameters struct {
	ExtensionType string `json:"extensionType"`
	Page          string `json:"page"`
	PerPage       string `json:"perPage"`
}

// ListAutomaticLocationUpdatesUsersParameters List Automatic Location Updates Users Parameters
type ListAutomaticLocationUpdatesUsersParameters struct {
	Type           string `json:"type"`
	SearchString   string `json:"searchString"`
	Department     string `json:"department"`
	SiteId         string `json:"siteId"`
	FeatureEnabled bool   `json:"featureEnabled"`
	OrderBy        string `json:"orderBy"`
	PerPage        int    `json:"perPage"`
	Page           int    `json:"page"`
}

// ListWirelessPointsParameters List Wireless Points Parameters
type ListWirelessPointsParameters struct {
	SiteId       string `json:"siteId"`
	SearchString string `json:"searchString"`
	OrderBy      string `json:"orderBy"`
	PerPage      int    `json:"perPage"`
	Page         int    `json:"page"`
}

// ListDevicesAutomaticLocationUpdatesParameters List Devices Automatic Location Updates Parameters
type ListDevicesAutomaticLocationUpdatesParameters struct {
	SiteId         string `json:"siteId"`
	FeatureEnabled bool   `json:"featureEnabled"`
	Model          string `json:"model"`
	CompatibleOnly bool   `json:"compatibleOnly"`
	SearchString   string `json:"searchString"`
	OrderBy        string `json:"orderBy"`
	PerPage        int    `json:"perPage"`
	Page           int    `json:"page"`
}

// ListAccountSwitchesParameters List Account Switches Parameters
type ListAccountSwitchesParameters struct {
	SiteId       string `json:"siteId"`
	SearchString string `json:"searchString"`
	OrderBy      string `json:"orderBy"`
	PerPage      int    `json:"perPage"`
	Page         int    `json:"page"`
}

// ListEmergencyLocationsParameters List Emergency Locations Parameters
type ListEmergencyLocationsParameters struct {
	SearchString      string `json:"searchString"`
	SiteId            string `json:"siteId"`
	AddressStatus     string `json:"addressStatus"`
	UsageStatus       string `json:"usageStatus"`
	DomesticCountryId string `json:"domesticCountryId"`
	OrderBy           string `json:"orderBy"`
	PerPage           int    `json:"perPage"`
	Page              int    `json:"page"`
}

// DeleteEmergencyLocationParameters Delete Emergency Location Parameters
type DeleteEmergencyLocationParameters struct {
	NewLocationId string `json:"newLocationId"`
	ValidateOnly  bool   `json:"validateOnly"`
}

// ReadConferencingSettingsParameters Read Conferencing Settings Parameters
type ReadConferencingSettingsParameters struct {
	CountryId string `json:"countryId"`
}

// ListCountriesParameters List Countries Parameters
type ListCountriesParameters struct {
	LoginAllowed      bool `json:"loginAllowed"`
	SignupAllowed     bool `json:"signupAllowed"`
	NumberSelling     bool `json:"numberSelling"`
	Page              int  `json:"page"`
	PerPage           int  `json:"perPage"`
	FreeSoftphoneLine bool `json:"freeSoftphoneLine"`
}

// ListLocationsParameters List Locations Parameters
type ListLocationsParameters struct {
	OrderBy string `json:"orderBy"`
	Page    int    `json:"page"`
	PerPage int    `json:"perPage"`
	StateId string `json:"stateId"`
	WithNxx bool   `json:"withNxx"`
}

// ListStatesParameters List States Parameters
type ListStatesParameters struct {
	AllCountries     bool `json:"allCountries"`
	CountryId        int  `json:"countryId"`
	Page             int  `json:"page"`
	PerPage          int  `json:"perPage"`
	WithPhoneNumbers bool `json:"withPhoneNumbers"`
}

// ListTimezonesParameters List Timezones Parameters
type ListTimezonesParameters struct {
	Page    string `json:"page"`
	PerPage string `json:"perPage"`
}

// ReadTimezoneParameters Read Timezone Parameters
type ReadTimezoneParameters struct {
	Page    string `json:"page"`
	PerPage string `json:"perPage"`
}

// ListAccountPhoneNumbersParameters List Account Phone Numbers Parameters
type ListAccountPhoneNumbersParameters struct {
	Page      int      `json:"page"`
	PerPage   int      `json:"perPage"`
	UsageType []string `json:"usageType"`
	Status    string   `json:"status"`
}

// ListExtensionsParameters List Extensions Parameters
type ListExtensionsParameters struct {
	ExtensionNumber string   `json:"extensionNumber"`
	Email           string   `json:"email"`
	Page            int      `json:"page"`
	PerPage         int      `json:"perPage"`
	Status          []string `json:"status"`
	Type            []string `json:"type"`
}

// ListUserTemplatesParameters List User Templates Parameters
type ListUserTemplatesParameters struct {
	Type    string `json:"type"`
	Page    string `json:"page"`
	PerPage string `json:"perPage"`
}

// ListCallQueuesParameters List Call Queues Parameters
type ListCallQueuesParameters struct {
	Page              int    `json:"page"`
	PerPage           int    `json:"perPage"`
	MemberExtensionId string `json:"memberExtensionId"`
}

// ListCallQueueMembersParameters List Call Queue Members Parameters
type ListCallQueueMembersParameters struct {
	Page    int `json:"page"`
	PerPage int `json:"perPage"`
}

// ListPagingGroupUsersParameters List Paging Group Users Parameters
type ListPagingGroupUsersParameters struct {
	Page    int `json:"page"`
	PerPage int `json:"perPage"`
}

// ListPagingGroupDevicesParameters List Paging Group Devices Parameters
type ListPagingGroupDevicesParameters struct {
	Page    int `json:"page"`
	PerPage int `json:"perPage"`
}

// ListCallMonitoringGroupsParameters List Call Monitoring Groups Parameters
type ListCallMonitoringGroupsParameters struct {
	Page              int    `json:"page"`
	PerPage           int    `json:"perPage"`
	MemberExtensionId string `json:"memberExtensionId"`
}

// ReadUserFeaturesParameters Read User Features Parameters
type ReadUserFeaturesParameters struct {
	AvailableOnly bool     `json:"availableOnly"`
	FeatureId     []string `json:"featureId"`
}

// ListCallMonitoringGroupMembersParameters List Call Monitoring Group Members Parameters
type ListCallMonitoringGroupMembersParameters struct {
	Page    int `json:"page"`
	PerPage int `json:"perPage"`
}

// CheckUserPermissionParameters Check User Permission Parameters
type CheckUserPermissionParameters struct {
	PermissionId      string `json:"permissionId"`
	TargetExtensionId string `json:"targetExtensionId"`
}

// SearchViaGet2Parameters Search Via Get2Parameters
type SearchViaGet2Parameters struct {
	Filter     string `json:"filter"`
	StartIndex int    `json:"startIndex"`
	Count      int    `json:"count"`
}

// ReadCallSessionStatusParameters Read Call Session Status Parameters
type ReadCallSessionStatusParameters struct {
	Timestamp string `json:"timestamp"`
	Timeout   string `json:"timeout"`
}

// PauseResumeCallRecordingParameters Pause Resume Call Recording Parameters
type PauseResumeCallRecordingParameters struct {
	BrandId string `json:"brandId"`
}

// Attachment attachment
type Attachment struct {
	Filename    string
	Bytes       []byte
	ContentType string
}
