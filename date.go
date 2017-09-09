package main

type Date struct {
	DayOfMonth  uint
	MonthOfYear uint
	Year        uint
}

var daysInMonth = map[uint]uint{
	1:  31,
	2:  28,
	3:  31,
	4:  30,
	5:  31,
	6:  30,
	7:  31,
	8:  31,
	9:  30,
	10: 31,
	11: 30,
	12: 31,
}

const february = 2

func find(needle uint, collection []uint) bool {
	for _, v := range collection {
		if v == needle {
			return true
		}
	}

	return false
}

func NewDate(d, m, y uint) *Date {

	if !isValid(d, m, y) {
		return nil
	}

	return &Date{
		DayOfMonth:  d,
		MonthOfYear: m,
		Year:        y,
	}
}

func isLeapYear(y uint) bool {
	return false
}

func inRange(start, end, val uint) bool {
	return val >= start && val <= end
}

func isValid(dayOfMonth, monthOfYear, year uint) bool {

	if year > 0 && inRange(1, 12, monthOfYear) && inRange(1, 31, dayOfMonth) {
		var maxDay = daysInMonth[monthOfYear]
		if monthOfYear == february && isLeapYear(year) {
			maxDay = 29
		}

		return dayOfMonth <= maxDay
	}

	return false
}

func (d *Date) Equals(d2 *Date) bool {
	return d.DayOfMonth == d2.DayOfMonth && d.Year == d2.Year && d.MonthOfYear == d2.MonthOfYear
}

func (d *Date) isAfter(d2 *Date) bool {
	if d.Year > d2.Year {
		return true
	} else if d.Year == d2.Year { // TODO bug
		if d.MonthOfYear > d2.MonthOfYear {
			return true
		} else if d.MonthOfYear == d2.MonthOfYear {
			return d.DayOfMonth > d2.DayOfMonth
		}
	}

	return false
}
