package outlook

type RecurrenceRangeType int

const (
	RecurrenceRangeEndDate RecurrenceRangeType = iota
	RecurrenceRangeNoEnd
	RecurrenceRangeNumbered
)

type RecurrenceRange struct {
	Type                RecurrenceRangeType
	StartDate           DateTimeOffset
	EndDate             DateTimeOffset
	NumberOfOccurrences int
}
