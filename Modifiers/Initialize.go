package Modifiers

import (
    "time"
    "fmt"
)

// Samay object
type Samay struct {
    time.Time
}

// Pass object of type Samay
func Create(t time.Time) Samay {
    return Samay{t}
}

// Console printer
var p = fmt.Println