package outlook

type ResponseType string

const (
	ResponseNone                ResponseType = "None"
	ResponseOrganizer                        = "Organizer"
	ResponseTentativelyAccepted              = "TentativelyAccepted"
	ResponseAccepted                         = "Accepted"
	ResponseDeclined                         = "Declined"
	ResponseNotResponded                     = "NotResponded"
)

type ResponseStatus struct {
	Response ResponseType
	Time     DateTimeOffset
}
