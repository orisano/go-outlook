package outlook

type ResponseType string

const (
	ResponseNone                ResponseType = "None"
	ResponseOrganizer           ResponseType = "Organizer"
	ResponseTentativelyAccepted ResponseType = "TentativelyAccepted"
	ResponseAccepted            ResponseType = "Accepted"
	ResponseDeclined            ResponseType = "Declined"
	ResponseNotResponded        ResponseType = "NotResponded"
)

type ResponseStatus struct {
	Response ResponseType
	Time     DateTime
}
