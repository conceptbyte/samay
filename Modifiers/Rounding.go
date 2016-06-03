package Rounding

import (
    "time"
    "fmt"
)

const NEGATE int = -1
const MINUTES_PER_HOUR int = 60
const MINUTES_PER_DAY int = 1439
const DAYS_PER_WEEK int = 7

// Get the start of hour
func (samay *Samay) StartOfHour() time.Time {
    return samay.Add(time.Duration(NEGATE * samay.Time.Minute()) * time.Minute)
}

// Get the end of hour
func (samay *Samay) EndOfHour() time.Time {
    return samay.Add(time.Duration(MINUTES_PER_HOUR - samay.Time.Minute()) * time.Minute)
}

// Get the start of day
func (samay *Samay) StartOfDay() time.Time {
    return samay.Add(time.Duration(NEGATE * samay.Time.Hour() * MINUTES_PER_HOUR + samay.Time.Minute()) * time.Minute)
}

// Get the end of day
func (samay *Samay) EndOfDay() time.Time {
    return samay.Add(time.Duration(MINUTES_PER_DAY - (samay.Time.Hour() * MINUTES_PER_HOUR + samay.Time.Minute())) * time.Minute)
}

// Get the start of week
func (samay *Samay) StartOfWeek() time.Time {
    day := int(samay.Weekday()) - 1
    return samay.Add(time.Duration(NEGATE * day * MINUTES_PER_DAY) * time.Minute)
}

// Get the end of week
func (samay *Samay) EndOfWeek() time.Time {
    day := DAYS_PER_WEEK - int(samay.Weekday())
    return samay.Add(time.Duration(day * MINUTES_PER_DAY) * time.Minute)
}

// Console printer
var p = fmt.Println