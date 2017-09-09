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
	return y%4 == 0 && y%400 != 0
}

func inRange(start, end, val uint) bool {
	return val >= start && val <= end
}

func isValid(dayOfMonth, monthOfYear, year uint) bool {

	if year > 0 && inRange(1, 12, monthOfYear) && inRange(1, 31, dayOfMonth) {
		date := Date{0, monthOfYear, year}
		return dayOfMonth <= date.DaysInMonth()
	}

	return false
}

func (d *Date) Equals(d2 *Date) bool {
	return d.DayOfMonth == d2.DayOfMonth && d.Year == d2.Year && d.MonthOfYear == d2.MonthOfYear
}

func (d *Date) isAfter(d2 *Date) bool {
	if d.Year > d2.Year {
		return true
	} else if d.Year == d2.Year {
		if d.MonthOfYear > d2.MonthOfYear {
			return true
		} else if d.MonthOfYear == d2.MonthOfYear {
			return d.DayOfMonth > d2.DayOfMonth
		}
	}

	return false
}

func (d *Date) DaysInMonth() uint {
	var maxDay = daysInMonth[d.MonthOfYear]
	if d.MonthOfYear == february && isLeapYear(d.Year) {
		maxDay = 29
	}

	return maxDay
}

func (d *Date) RemainingDaysInMonth() uint {
	return d.DaysInMonth() - d.DayOfMonth + 1
}

func (d *Date) AddMonths(number uint) {
	extraYears := number / 12
	d.Year += extraYears

	extraMonths := number % 12
	d.MonthOfYear += extraMonths

	if d.MonthOfYear > 12 {
		d.MonthOfYear = d.MonthOfYear % 12
		d.Year++
	}
}

func (d *Date) AddMonth() {
	d.AddMonths(1)
}

func (d *Date) AddDay() {
	if d.DaysInMonth() >= d.DayOfMonth+1 {
		d.DayOfMonth++
	} else {
		d.AddMonths(1)
		d.DayOfMonth = 1
	}
}
