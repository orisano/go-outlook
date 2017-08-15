package outlook

import "time"

type RecurrencePatternType int

const (
	RecurrencePatternDaily RecurrencePatternType = iota
	RecurrencePatternWeekly
	RecurrencePatternAbsoluteMonthly
	RecurrencePatternRelativeMonthly
	RecurrencePatternAbsoluteYearly
	RecurrencePatternRelativeYearly
)

type WeekIndex int

const (
	WeekFirst WeekIndex = iota
	WeekSecond
	WeekThird
	WeekFourth
	WeekLast
)

type RecurrencePattern struct {
	Type           RecurrencePatternType
	Interval       int
	DayOfMonth     int
	Month          int
	DaysOfWeek     []time.Weekday
	FirstDayOfWeek time.Weekday
	Index          WeekIndex
}
