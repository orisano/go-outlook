package outlook

type AttendeeType string

const (
	AttendeeRequired AttendeeType = "Required"
	AttendeeOptional AttendeeType = "Optional"
	AttendeeResource AttendeeType = "Resource"
)

type Attendee struct {
	Status ResponseStatus
	Type   AttendeeType
}
