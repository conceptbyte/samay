package Modifiers

import (
	"time"
)

const NEGATE int = -1
const MINUTES_PER_HOUR int = 60
const MINUTES_PER_DAY int = 1439
const DAYS_PER_WEEK int = 7

/**
 * Get the start of hour
 * For a given time t, it return the first minute of the hour.
 */
func (samay Samay) StartOfHour() Samay {
	return Create(samay.Add(time.Duration(NEGATE*samay.Time.Minute()) * time.Minute))
}

/**
 * Get the end of hour
 * For a given time t, it returns the last minute of the hour.
 */
func (samay Samay) EndOfHour() Samay {
	return Create(samay.Add(time.Duration(MINUTES_PER_HOUR-samay.Time.Minute()) * time.Minute))
}

/**
 * Get the start of day
 * For a given time t, return the first minute of the day.
 */
func (samay Samay) StartOfDay() Samay {
	d := time.Duration(NEGATE*samay.Time.Hour()*MINUTES_PER_HOUR+samay.Time.Minute()) * time.Minute
	return Create(samay.Add(d))
}

/**
 * Get the end of day
 * For a given time t, return the last minute of the day.
 */
func (samay Samay) EndOfDay() Samay {
	d := time.Duration(MINUTES_PER_DAY-(samay.Time.Hour()*MINUTES_PER_HOUR+samay.Time.Minute())) * time.Minute
	return Create(samay.Add(d))
}

/**
 * Get the start of week
 * For a given time t, return the first minute of the day.
 */
func (samay Samay) StartOfWeek() Samay {
	day := int(samay.Weekday()) - 1
	return Create(samay.Time.Add(time.Duration(NEGATE*day*MINUTES_PER_DAY) * time.Minute)).StartOfDay()
}

// Get the end of week
func (samay Samay) EndOfWeek() Samay {
	day := DAYS_PER_WEEK - int(samay.Weekday())
	return Create(samay.Add(time.Duration(day*MINUTES_PER_DAY) * time.Minute)).EndOfDay()
}
