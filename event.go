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
	Attachments                []Attachment
	Attendees                  []Attendee
	Body                       ItemBody
	BodyPreview                string
	Calendar                   Calendar
	Categories                 []string
	ChangeKey                  string
	Created                    DateTime `json:"CreatedDateTime"`
	LastModified               DateTime `json:"LastModifiedDateTime"`
	End                        DateTimeTimeZone
	HasAttachments             bool
	ID                         string `json:"Id"`
	Importance                 Importance
	Instances                  []Event
	ICalUID                    string `json:"iCalUID"`
	IsAllDay                   bool
	IsCancelled                bool
	IsOrganizer                bool
	IsReminderOn               bool
	Location                   Location
	OnlineMeetingURL           string `json:"OnlineMeetingUrl"`
	Organizer                  Recipient
	OriginalEndTimeZone        string
	OriginalStartTimeZone      string
	Recurrence                 PatternedRecurrence
	ReminderMinutesBeforeStart int
	ResponseRequested          bool
	ResponseStatus             ResponseStatus
	Sensitivity                Sensitivity
	SeriesMasterID             string `json:"SeriesMasterId"`
	ShowAs                     FreeBusyStatus
	Start                      DateTimeTimeZone
	Type                       EventType
	WebLink                    string
}
