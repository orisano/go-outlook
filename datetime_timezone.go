package outlook

import "time"

type DateTime struct {
	time.Time
}

func (d *DateTime) UnmarshalJSON(b []byte) error {
	time.Parse()
}

type DateTimeTimeZone struct {
}
