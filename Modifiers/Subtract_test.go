package Modifiers

import (
	// "fmt"
	"testing"
	"time"
)

/**
 * Test SubtractHours
 * @param {testing} t *testing.T Testing
 */
func TestSubtractHours(t *testing.T) {
	layout := "Mon, 01/02/06, 03:04PM"
	one, _ := time.Parse(layout, "Thu, 05/19/11, 10:47PM")
	two, _ := time.Parse(layout, "Thu, 05/19/11, 08:47PM")

	start := Create(one).SubtractHours(2)

	if start.Time != two {
		t.Error("Assertion did not match ")
	}
}

/**
 * Test SubtractDays
 * @param {testing} t *testing.T Testing
 */
func TestSubtractDays(t *testing.T) {
	layout := "Mon, 01/02/06, 03:04PM"
	one, _ := time.Parse(layout, "Thu, 05/19/11, 10:47PM")
	two, _ := time.Parse(layout, "Thu, 05/17/11, 10:47PM")

	start := Create(one).SubtractDays(2)

	if start.Time != two {
		t.Error("Assertion did not match ")
	}
}

/**
 * Test SubtractWeeks
 * @param {testing} t *testing.T Testing
 */
func TestSubtractWeeks(t *testing.T) {
	layout := "Mon, 01/02/06, 03:04PM"
	one, _ := time.Parse(layout, "Thu, 05/19/11, 10:47PM")
	two, _ := time.Parse(layout, "Thu, 05/05/11, 10:47PM")

	start := Create(one).SubtractWeeks(2)

	if start.Time != two {
		t.Error("Assertion did not match ")
	}
}

/**
 * Test SubtractMonths
 * @param {testing} t *testing.T Testing
 */
func TestSubtractMonths(t *testing.T) {
	layout := "Mon, 01/02/06, 03:04PM"
	one, _ := time.Parse(layout, "Thu, 05/19/11, 10:47PM")
	two, _ := time.Parse(layout, "Thu, 03/19/11, 10:47PM")

	start := Create(one).SubtractMonths(2)

	if start.Time != two {
		t.Error("Assertion did not match ")
	}
}

/**
 * Test SubtractYears
 * @param {testing} t *testing.T Testing
 */
func TestSubtractYears(t *testing.T) {
	layout := "Mon, 01/02/06, 03:04PM"
	one, _ := time.Parse(layout, "Thu, 05/19/11, 10:47PM")
	two, _ := time.Parse(layout, "Thu, 05/19/09, 10:47PM")

	start := Create(one).SubtractYears(2)

	if start.Time != two {
		t.Error("Assertion did not match ")
	}
}
