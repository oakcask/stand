package stand

import (
	"time"
)

// Clock is a source of time.
type Clock interface {
	// Now returns the time be pointed by the clock when it called.
	Now() time.Time
}

type fixedClock struct {
	t time.Time
}

func (cl *fixedClock) Now() time.Time {
	if cl == nil {
		return time.Time{}
	}

	return cl.t
}

// NewFixed creates Clock instance where its Now method
// always return the given time.
func NewFixed(t time.Time) Clock {
	return &fixedClock{t}
}

type systemClock struct{}

// SystemClock is a clock which turns just as normal clock,
// and Now method will return the return value of time.Now.
var SystemClock *systemClock = &systemClock{}

func (cl *systemClock) Now() time.Time {
	return time.Now()
}

type laggedClock struct {
	source Clock
	delta  time.Duration
}

func (cl *laggedClock) Now() time.Time {
	return cl.source.Now().Add(cl.delta)
}

// Advance creates new clock which generates
// advanced time by the given duration from the given clock.
func Advance(cl Clock, d time.Duration) Clock {
	return &laggedClock{source: cl, delta: d}
}

// Revert is shorthand for Advance with negative duration,
// and creates reverted time by the given duration from the given clock.
func Revert(cl Clock, d time.Duration) Clock {
	return Advance(cl, -d)
}

// Travel creates new clock that generates time like
// travelled to the given time, which means the new clock has
// time difference from the given time to the time pointed by
// the given clock just now.
func Travel(cl Clock, t time.Time) Clock {
	return Advance(cl, t.Sub(cl.Now()))
}

// Pause creates new fixed clock which points
// the time be pointed by the given clock just now.
func Pause(cl Clock) Clock {
	return NewFixed(cl.Now())
}

// Resume creates new clock that proceeds just like SystemClock but
// it starts from the time pointed by the given clock.
func Resume(cl Clock) Clock {
	return Travel(SystemClock, cl.Now())
}
