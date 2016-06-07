package Modifiers

import (
	"time"
)

// Get time formatted to ANSIC
func (samay Samay) FormatANSIC() string {
	return Create(samay.Time).Format(time.ANSIC)
}

// Get time formatted to RFC822Z
func (samay Samay) FormatRFC822Z() string {
	return Create(samay.Time).Format(time.RFC822Z)
}

// Get time formatted to RFC822
func (samay Samay) FormatRFC822() string {
	return Create(samay.Time).Format(time.RFC822)
}
