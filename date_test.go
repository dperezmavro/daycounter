package main

import (
	"testing"
)

func makeDate(d, m, y uint) *Date {
	return &Date{
		DayOfMonth:  d,
		MonthOfYear: m,
		Year:        y,
	}
}

func TestDateInit(t *testing.T) {
	type testDate struct {
		want             *Date
		year, month, day uint
		shouldFail       bool
	}
	makeTestDate := func(da *Date, y uint, m uint, d uint) *testDate {
		return &testDate{
			want:  da,
			year:  y,
			month: m,
			day:   d,
		}
	}

	tests := []*testDate{
		// invalid dates by ranges
		makeTestDate(nil, 0, 0, 0),
		makeTestDate(nil, 0, 15, 0),
		makeTestDate(nil, 1, 10, 0),
		makeTestDate(nil, 1, 10, 33),
		makeTestDate(nil, 1990, 15, 18),

		//invalid dates by bad date
		makeTestDate(nil, 1990, 0, 0),
		makeTestDate(nil, 1992, 15, 0),
		makeTestDate(nil, 2015, 9, 31),
		makeTestDate(nil, 1992, 12, 40),
		makeTestDate(nil, 2015, 2, 30),
		makeTestDate(nil, 2015, 2, 29),
		makeTestDate(nil, 2015, 4, 31),

		//valid date
		makeTestDate(&Date{3, 10, 1}, 1, 10, 3),
		makeTestDate(&Date{29, 2, 2016}, 2016, 2, 29),
		makeTestDate(&Date{29, 4, 2016}, 2016, 4, 29),
	}

	for i, test := range tests {
		got := NewDate(test.day, test.month, test.year)
		if test.want == nil {
			if got != nil {
				t.Errorf(
					"TestDateInit: test(%d) failed: got value (%v) when expecting nil",
					i,
					got,
				)
			}
		} else if !test.want.Equals(got) {
			t.Errorf(
				"TestDateInit: test(%d) failed: want %v got %v",
				i,
				test.want,
				got,
			)
		}
	}
}

func TestEquals(t *testing.T) {
	type test struct {
		want   bool
		d1, d2 *Date
	}
	tests := []test{
		test{want: true, d1: makeDate(0, 0, 0), d2: makeDate(0, 0, 0)},
		test{want: true, d1: makeDate(3, 5, 1991), d2: makeDate(3, 5, 1991)},
		test{want: false, d1: makeDate(3, 5, 1991), d2: makeDate(3, 6, 1991)},
	}

	for i, test := range tests {
		got := test.d1.Equals(test.d2)
		if test.want != got {
			t.Errorf(
				"TestEquals: test(%d) failed: want %v got %v",
				i,
				test.want,
				got,
			)
		}
	}
}

func TestIsAfter(t *testing.T) {
	type test struct {
		want   bool
		d1, d2 *Date
	}
	tests := []test{
		test{want: false, d1: makeDate(0, 0, 0), d2: makeDate(0, 0, 0)},
		test{want: true, d1: makeDate(3, 5, 1991), d2: makeDate(3, 4, 1991)},
		test{want: false, d1: makeDate(3, 5, 1991), d2: makeDate(3, 4, 1993)},
		test{want: false, d1: makeDate(3, 5, 1991), d2: makeDate(3, 6, 1991)},
		test{want: false, d1: makeDate(3, 5, 1991), d2: makeDate(3, 5, 1991)},
		test{want: true, d1: makeDate(4, 5, 1991), d2: makeDate(3, 5, 1991)},
	}

	for i, test := range tests {
		got := test.d1.isAfter(test.d2)
		if test.want != got {
			t.Errorf(
				"TestIsAfter: test(%d) failed: want %v to be after %v",
				i,
				test.d1,
				test.d2,
			)
		}
	}
}

func TestRemainingDaysInMonth(t *testing.T) {
	type test struct {
		date          *Date
		daysRemaining uint
	}

	tests := []test{
		test{
			date:          makeDate(10, 10, 2015),
			daysRemaining: 22,
		},
		test{
			date:          makeDate(31, 12, 1800),
			daysRemaining: 1,
		},
		test{
			date:          makeDate(15, 2, 2016),
			daysRemaining: 15,
		},
	}

	for i, testCase := range tests {
		got := testCase.date.RemainingDaysInMonth()
		if got != testCase.daysRemaining {
			t.Errorf("TestRemainingDaysInMonth(%d): wanted %d got %d",
				i,
				testCase.daysRemaining,
				got,
			)
		}
	}
}

func TestGetDaysInMonth(t *testing.T) {
	type test struct {
		date        *Date
		daysInMonth uint
	}

	tests := []test{
		test{
			date:        makeDate(10, 10, 2015),
			daysInMonth: 31,
		},
		test{
			date:        makeDate(31, 12, 1800),
			daysInMonth: 31,
		},
		test{
			date:        makeDate(30, 6, 1997),
			daysInMonth: 30,
		},
		test{
			date:        makeDate(1, 2, 2016),
			daysInMonth: 29,
		},
	}

	for i, testCase := range tests {
		got := testCase.date.DaysInMonth()
		if got != testCase.daysInMonth {
			t.Errorf("TestGetDaysInMonth(%d): wanted %d got %d",
				i,
				testCase.daysInMonth,
				got,
			)
		}
	}
}

func TestGetAddMonths(t *testing.T) {
	type test struct {
		startDate, endDate *Date
		diff               uint
	}

	tests := []test{
		test{
			startDate: makeDate(10, 10, 2015),
			endDate:   makeDate(10, 12, 2015),
			diff:      2,
		},
		test{
			startDate: makeDate(10, 10, 2015),
			endDate:   makeDate(10, 2, 2016),
			diff:      4,
		},
		test{
			startDate: makeDate(10, 2, 2015),
			endDate:   makeDate(10, 2, 2017),
			diff:      24,
		},
		test{
			startDate: makeDate(10, 3, 2015),
			endDate:   makeDate(10, 1, 2017),
			diff:      22,
		},
		test{
			startDate: makeDate(10, 1, 2001),
			endDate:   makeDate(10, 5, 2012),
			diff:      16,
		},
	}

	for i, testCase := range tests {
		testCase.startDate.AddMonths(testCase.diff)
		if !testCase.startDate.Equals(testCase.endDate) {
			t.Errorf("TestGetAddMonths(%d): wanted %d got %d",
				i,
				testCase.endDate,
				testCase.startDate,
			)
		}
	}
}
