package outlook

type RecurrenceRangeType int

const (
	RecurrenceRangeEndDate RecurrenceRangeType = iota
	RecurrenceRangeNoEnd
	RecurrenceRangeNumbered
)

type RecurrenceRange struct {
	Type                RecurrenceRangeType
	StartDate           DateTime
	EndDate             DateTime
	NumberOfOccurrences int
}
