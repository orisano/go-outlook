package outlook

import "time"

const (
	layout = `"` + time.RFC3339 + `"`
)

type DateTime struct {
	time.Time
}

func (d *DateTime) UnmarshalJSON(b []byte) error {
	t, err := time.Parse(layout, string(b))
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

func (d *DateTime) MarshalJSON() ([]byte, error) {
	s := d.Time.UTC().Format(layout)
	return []byte(s), nil
}

type DateTimeTimeZone struct {
	DateTime DateTime
	TimeZone string
}
