package Rounding

import "time"

// Samay object
type Samay struct {
    time.Time
}

// Pass object of type Samay
func Create(t time.Time) *Samay {
    return &Samay{t}
}