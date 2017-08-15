package outlook

type Sensitivity int

const (
	SensitivityNormal Sensitivity = iota
	SensitivityPersonal
	SensitivityPrivate
	SensitivityConfidential
)

type FreeBusyStatus int

const (
	StatusFree FreeBusyStatus = iota
	StatusTentative
	StatusBusy
	StatusOof
	StatusWorkingElsewhere
	StatusUnknown FreeBusyStatus = -1
)

type EventType int

const (
	EventSingleInstance EventType = iota
	EventOccurrence
	EventException
	EventSeriesMaster
)

type Event struct {
	Attachments       []Attachment
	Attendees         []Attendee
	Body              ItemBody
	BodyPreview       string
	Calendar          Calendar
	Categories        []string
	ChangeKey         string
	Created           EventDateTimeOffset `json:"CreatedDateTime"`
	LastModified      EventDateTimeOffset `json:"LastModifiedDateTime"`
	End               EventDateTimeOffset
	EndTimeZone       string
	HasAttachments    bool
	ID                string `json:"Id"`
	Importance        Importance
	Instances         []Event
	ICalUID           string `json:"iCalUID"`
	IsAllDay          bool
	IsCancelled       bool
	IsOrganizer       bool
	Location          Location
	Organizer         Recipient
	Recurrence        PatternedRecurrence
	ResponseRequested bool
	ResponseStatus    ResponseStatus
	Sensitivity       Sensitivity
	SeriesMasterID    string `json:"SeriesMasterId"`
	ShowAs            FreeBusyStatus
	Start             EventDateTimeOffset
	StartTimeZone     string
	Subject           string
	Type              EventType
	WebLink           string
}
