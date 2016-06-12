package Modifiers

import (
	"time"
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
