package main

import "testing"

func TestNewDateRange(t *testing.T) {
	type test struct {
		dateRange  *DateRange
		start, end *Date
		isError    bool
	}

	makeTest := func(dr *DateRange, d1, d2 *Date, nilerr bool) test {
		return test{
			dateRange: dr,
			start:     d1,
			end:       d2,
			isError:   nilerr,
		}
	}

	tests := []test{
		makeTest(nil, nil, nil, true),
		makeTest(
			&DateRange{StartDate: makeDate(3, 10, 1991), EndDate: makeDate(3, 10, 1993)},
			makeDate(3, 10, 1991),
			makeDate(3, 10, 1993),
			false,
		),
		makeTest(
			nil,
			makeDate(3, 10, 1991),
			makeDate(3, 9, 1991),
			true,
		),
	}

	for i, testV := range tests {
		_, err := NewDateRange(testV.start, testV.end)
		if err == nil && testV.isError == true {
			t.Errorf("TestNewDateRange (%d) got nil error", i)
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
			&DateRange{
				StartDate: makeDate(3, 10, 1991),
				EndDate:   makeDate(3, 10, 1991)},
			1,
		),
		makeTest(
			&DateRange{
				StartDate: makeDate(3, 10, 1991),
				EndDate:   makeDate(5, 10, 1991)},
			3,
		),
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
