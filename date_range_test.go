package main

import "testing"

func TestNewDateRange(t *testing.T) {
	type test struct {
		dateRange  *DateRange
		start, end *Date
	}

	makeTest := func(dr *DateRange, d1, d2 *Date) test {
		return test{
			dateRange: dr,
			start:     d1,
			end:       d2,
		}
	}

	tests := []test{
		makeTest(nil, nil, nil),
		makeTest(
			&DateRange{StartDate: makeDate(3, 10, 1991), EndDate: makeDate(3, 10, 1993)},
			makeDate(3, 10, 1991),
			makeDate(3, 10, 1993),
		),
		makeTest(
			nil,
			makeDate(3, 10, 1991),
			makeDate(3, 9, 1991),
		),
	}

	for i, testV := range tests {
		got := NewDateRange(testV.start, testV.end)
		if got == nil && testV.dateRange != nil {
			t.Errorf(
				"TestNewDateRange (%d) got nil when wanting %v",
				i,
				testV.dateRange,
			)
		}

	}
}

func TestNumberOfDays(t *testing.T) {
	type test struct {
		dateRange *DateRange
		want      uint
	}

	makeTest := func(dr *DateRange, days uint) test {
		return test{
			dateRange: dr,
			want:      days,
		}
	}

	tests := []test{
		makeTest(
			&DateRange{StartDate: makeDate(3, 10, 1991), EndDate: makeDate(3, 10, 1991)},
			1,
		),
		makeTest(
			&DateRange{StartDate: makeDate(3, 10, 1991), EndDate: makeDate(5, 10, 1991)},
			3,
		),
		// makeTest(
		// 	nil,
		// 	makeDate(3, 10, 1991),
		// 	makeDate(3, 9, 1991),
		// ),
	}

	for i, testV := range tests {
		got := testV.dateRange.NumberOfDays()
		if got != testV.want {
			t.Errorf(
				"TestNumberOfDays (%d) got %d NoD when wanting %v",
				i,
				got,
				testV.want,
			)
		}

	}
}
