package Modifiers

import (
	Samay "../Modifiers"
	"testing"
	"time"
)

/**
 * Test StartOfHour
 * @param {testing} t *testing.T Testing
 */
func TestStartOfHour(t *testing.T) {
	layout := "Mon, 01/02/06, 03:04PM"
	one, _ := time.Parse(layout, "Thu, 05/19/11, 10:47PM")
	two, _ := time.Parse(layout, "Thu, 05/19/11, 10:00PM")

	start := Samay.Create(one).StartOfHour()

	if start.Time != two {
		t.Error("Assertion did not match")
	}
}

/**
 * Test EndOfHour
 * @param {testing} t *testing.T Testing
 */
func TestEndOfHour(t *testing.T) {
	layout := "Mon, 01/02/06, 03:04PM"
	one, _ := time.Parse(layout, "Thu, 05/19/11, 10:47PM")
	two, _ := time.Parse(layout, "Thu, 05/19/11, 11:00PM")

	start := Samay.Create(one).EndOfHour()

	if start.Time != two {
		t.Error("Assertion did not match")
	}
}

/**
 * Test StartOfDay
 * @param {testing} t *testing.T Testing
 */
func TestStartOfDay(t *testing.T) {
	layout := "Mon, 01/02/06, 03:04PM"
	one, _ := time.Parse(layout, "Thu, 05/19/11, 10:47PM")
	two, _ := time.Parse(layout, "Thu, 05/19/11, 00:00AM")

	start := Samay.Create(one).StartOfDay()

	if start.Time != two {
		t.Error("Assertion did not match")
	}
}

/**
 * Test EndOfDay
 * @param {testing} t *testing.T Testing
 */
func TestEndOfDay(t *testing.T) {
	layout := "Mon, 01/02/06, 03:04PM"
	one, _ := time.Parse(layout, "Thu, 05/19/11, 10:47PM")
	two, _ := time.Parse(layout, "Thu, 05/19/11, 11:59PM")

	start := Samay.Create(one).EndOfDay()

	if start.Time != two {
		t.Error("Assertion did not match")
	}
}

/**
 * Test StartOfWeek
 * @param {testing} t *testing.T Testing
 */
func TestStartOfWeek(t *testing.T) {
	layout := "Mon, 01/02/06, 03:04PM"
	one, _ := time.Parse(layout, "Thu, 05/19/11, 10:47PM")
	two, _ := time.Parse(layout, "Thu, 05/15/11, 00:00AM")

	start := Samay.Create(one).StartOfWeek()

	if start.Time != two {
		t.Error("Assertion did not match")
	}
}

/**
 * Test EndOfWeek
 * @param {testing} t *testing.T Testing
 */
func TestEndOfWeek(t *testing.T) {
	layout := "Mon, 01/02/06, 03:04PM"
	one, _ := time.Parse(layout, "Thu, 05/19/11, 10:47PM")
	two, _ := time.Parse(layout, "Thu, 05/22/11, 11:59PM")

	start := Samay.Create(one).EndOfWeek()

	if start.Time != two {
		t.Error("Assertion did not match")
	}
}
