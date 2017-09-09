package main

import (
	"errors"
	"fmt"
)

type DateRange struct {
	StartDate *Date
	EndDate   *Date
}

func NewDateRange(start, end *Date) (*DateRange, error) {
	if start == nil || end == nil {
		return nil, fmt.Errorf("nil values detected. start: %v, end %v", start, end)
	} else if !end.isAfter(start) {
		return nil, errors.New("End is not after start")
	}

	return &DateRange{
		StartDate: start,
		EndDate:   end,
	}, nil
}

func (dr *DateRange) NumberOfDays() uint {
	return 1
}
