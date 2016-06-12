package Modifiers

import (
	"time"
)

const (
	NEGATE           int = -1
	MINUTES_PER_HOUR int = 60
	MINUTES_PER_DAY  int = 1439
	DAYS_PER_WEEK    int = 7
	HOURS_PER_DAY    int = 24
)

/**
 * Samay construct
 */
type Samay struct {
	time.Time
}

/**
 * Named constructor of Samay
 */
func Create(t time.Time) *Samay {
	return &Samay{t}
}
