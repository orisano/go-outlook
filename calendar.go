package outlook

type CalendarColor int

const (
	ColorLightBlue CalendarColor = iota
	ColorLightGreen
	ColorLightOrange
	ColorLightGray
	ColorLightYellow
	ColorLightTeal
	ColorLightPink
	ColorLightBrown
	ColorLightRed
	ColorMax
	ColorAuto CalendarColor = -1
)

type Calendar struct {
	Name         string
	ChangeKey    string
	Color        CalendarColor
	ID           string `json:"Id"`
	CalendarView []Event
	Events       []Event
}
