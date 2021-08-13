package calendar

import (
	"math/rand"
	"time"

	"github.com/golang/glog"
)

// Checks if specified Time is a weekday
func isWeekday(t time.Time) bool {
	switch t.Weekday() {
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		return true
	case time.Saturday, time.Sunday:
		return false
	}

	glog.Fatalf("Unrecognized day of the week: %s", t.Weekday().String())

	panic("Explicit Panic to avoid compiler error: missing return at end of function")
}

// Returns the next weekday in Location
func nextWeekday(loc *time.Location) time.Time {
	check := time.Now().In(loc)
	for {
		check = check.AddDate(0, 0, 1)
		if isWeekday(check) {
			return check
		}
	}
}

// NextRuntime calculates the next time the Scheduled should run
func NextRuntime(loc *time.Location, r int) time.Time {
	now := time.Now().In(loc)

	// Is today a weekday and are we still in time for it?
	if isWeekday(now) {
		runtimeToday := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), r, 0, 0, loc)
		if runtimeToday.After(now) {
			return runtimeToday
		}
	}

	// Missed the train for today. Schedule on next weekday
	nexWeekDay := nextWeekday(loc)
	year, month, day := nexWeekDay.Date()
	hours := nexWeekDay.Hour()
	minutes := nexWeekDay.Minute()
	return time.Date(year, month, day, hours, minutes, 0, 0, loc)
}

// RandomTimeInRange returns a random time within the range specified by RunInterval
func RandomTimeInRange(interval int, loc *time.Location) time.Time {
	// calculate the number of seconds in the range
	secondsInRange := interval * 60

	// calculate a random seconds-offset in range [0, secondsInRange)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randSecondsOffset := r.Intn(secondsInRange)
	offsetDuration := time.Duration(randSecondsOffset) * time.Second

	// Add the seconds offset to the start of the range to get a random
	// time within the range
	return time.Now().In(loc).Add(offsetDuration)
}
