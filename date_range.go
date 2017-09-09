package main

type DateRange struct {
	StartDate *Date
	EndDate   *Date
}

func NewDateRange(start, end *Date) *DateRange {
	if start == nil || end == nil || !end.isAfter(start) {
		return nil
	}

	return &DateRange{
		StartDate: start,
		EndDate:   end,
	}
}

func (dr *DateRange) NumberOfDays() uint {

	return 0
}
