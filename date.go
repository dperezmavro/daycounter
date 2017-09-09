package main

type Date struct {
	DayOfMonth  uint
	MonthOfYear uint
	Year        uint
}

var thirtyDayMonths []uint = []uint{4, 6, 9, 11}
var thirtyOneDayMonths []uint = []uint{1, 3, 5, 7, 8, 10, 12}

const February = 2

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
		return &Date{}
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

func isValid(dayOfMonth, monthOfYear, year uint) bool {

	if year > 0 && monthOfYear >= 1 && monthOfYear <= 12 && dayOfMonth >= 1 && dayOfMonth <= 31 {
		var maxDay uint = 28
		if monthOfYear == February {
			if isLeapYear(year) {
				maxDay = 29
			}
		} else if find(monthOfYear, thirtyDayMonths) {
			maxDay = 30
		} else if find(monthOfYear, thirtyOneDayMonths) {
			maxDay = 31
		}

		return dayOfMonth <= maxDay
	}

	return false
}

func (d *Date) Equals(d2 *Date) bool {
	return d.DayOfMonth == d2.DayOfMonth && d.Year == d2.Year && d.MonthOfYear == d2.MonthOfYear
}
