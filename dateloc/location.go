package dateloc

import "time"

var current *time.Location

func Init(location *string) error {
	var locLabel string

	if location == nil {
		locLabel = LocationSeoul
	} else {
		locLabel = *location
	}

	loc, err := time.LoadLocation(locLabel)
	if err != nil {
		return err
	}

	current = loc

	return nil
}

func Current() *time.Location {
	if current == nil {
		if err := Init(nil); err != nil {
			panic(err)
		}
	}

	return current
}
