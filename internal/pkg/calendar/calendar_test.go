package calendar

import (
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

func TestIsWeekDay(t *testing.T) {
	monday := time.Date(2018, 4, 16, 0, 0, 0, 0, time.UTC)

	assert.True(t, isWeekday(monday))
	assert.True(t, isWeekday(monday.Add(time.Hour*24)))
	assert.True(t, isWeekday(monday.Add(time.Hour*24*2)))
	assert.True(t, isWeekday(monday.Add(time.Hour*24*3)))
	assert.True(t, isWeekday(monday.Add(time.Hour*24*4)))

	assert.False(t, isWeekday(monday.Add(time.Hour*24*5)))
	assert.False(t, isWeekday(monday.Add(time.Hour*24*6)))
}

func TestNextWeekDay(t *testing.T) {
	var today, next time.Time
	defer monkey.Unpatch(time.Now)

	for i := 16; i < 23; i++ {
		today = time.Date(2018, 4, i, 0, 0, 0, 0, time.UTC)

		switch today.Weekday() {
		case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Sunday:
			next = today.AddDate(0, 0, 1)
		case time.Saturday:
			next = today.AddDate(0, 0, 2)
		case time.Friday:
			next = today.AddDate(0, 0, 3)
		}

		monkey.Patch(time.Now, func() time.Time {
			return today
		})

		assert.Equal(t, nextWeekday(time.UTC), next)
	}
}

func TestNextRuntime(t *testing.T) {

	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2018, 4, 16, 12, 0, 0, 0, time.UTC)
	})
	defer monkey.Unpatch(time.Now)

	monday := time.Date(2018, 4, 16, 12, 0, 0, 0, time.UTC)
	assert.Equalf(t, NextRuntime(time.UTC, 13), monday.Add(13*time.Minute), "Expected to be run today if today is a weekday and there is time for it")

	sunday := time.Date(2018, 4, 15, 0, 0, 0, 0, time.UTC)
	monkey.Patch(time.Now, func() time.Time {
		return sunday
	})

	assert.Equalf(t, NextRuntime(time.UTC, 10), sunday.Add(time.Hour*24), "Expected to be run next weekday if today is a weekend day")
}

func TestRandomTimeInRange(t *testing.T) {
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2018, 4, 16, 12, 0, 0, 0, time.UTC)
	})
	defer monkey.Unpatch(time.Now)

	randomTime := RandomTimeInRange(10, time.UTC)

	scheduledTime := func() (success bool) {
		if randomTime.Minute() >= 0 && randomTime.Minute() <= 10 {
			success = true
		}
		return
	}

	assert.Condition(t, scheduledTime)
}
