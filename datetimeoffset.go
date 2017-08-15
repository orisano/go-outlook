package outlook

import "time"

const (
	timeLayout      = time.RFC3339
	eventTimeLayout = "2015-02-25T21:34:00-08:00"
)

type DateTimeOffset struct {
	time.Time
}

type EventDateTimeOffset struct {
	time.Time
}

func (d *DateTimeOffset) UnmarshalJSON(b []byte) error {
	t, err := time.Parse(timeLayout, string(b))
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

func (d *DateTimeOffset) MarshalJSON() ([]byte, error) {
	s := d.UTC().Format(timeLayout)
	return []byte(s), nil
}

func (e *EventDateTimeOffset) UnmarshalJSON(b []byte) error {
	t, err := time.Parse(eventTimeLayout, string(b))
	if err != nil {
		return err
	}
	e.Time = t
	return nil
}

func (e *EventDateTimeOffset) MarshalJSON() ([]byte, error) {
	s := e.UTC().Format(eventTimeLayout)
	return []byte(s), nil
}
