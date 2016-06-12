package Modifiers

import (
	"time"
)

/**
 * Subtract hours from the given time
 */
func (samay *Samay) SubtractHours(interval int) *Samay {
	return Create(samay.Add(time.Duration(NEGATE*interval*MINUTES_PER_HOUR) * time.Minute))
}

/**
 * Subtract days from the given time
 */
func (samay *Samay) SubtractDays(interval int) *Samay {
	return samay.SubtractHours(interval * HOURS_PER_DAY)
}

/**
 * Subtract weeks from the given time
 */
func (samay *Samay) SubtractWeeks(interval int) *Samay {
	return samay.SubtractDays(interval * DAYS_PER_WEEK)
}

/**
 * Subtract months from the given time
 */
func (samay *Samay) SubtractMonths(interval int) *Samay {
	return Create(samay.Time.AddDate(0, (NEGATE * interval), 0))
}

/**
 * Subtract years from the given time
 */
func (samay *Samay) SubtractYears(interval int) *Samay {
	return Create(samay.Time.AddDate((NEGATE * interval), 0, 0))
}
