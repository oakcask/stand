package stand

import (
	"fmt"
	"time"
)

func ExampleSystemClock_Now() {
	now1 := time.Now()
	now2 := SystemClock.Now()

	if now2.Sub(now1) <= time.Duration(500)*time.Millisecond {
		fmt.Println("It's close enough to time.Now()")
	}

	// Output:
	// It's close enough to time.Now()
}

func ExampleNewFixed() {
	epoch, _ := time.Parse("2006-01-02T15:04:05Z", "1970-01-01T00:00:00Z")
	clock := NewFixed(epoch)

	fmt.Println(clock.Now().Format("2006-01-02T15:04:05Z"))

	// Output:
	// 1970-01-01T00:00:00Z
}

func ExampleAdvance() {
	epoch, _ := time.Parse("2006-01-02T15:04:05Z", "1970-01-01T00:00:00Z")
	fixedClock := NewFixed(epoch)

	clock := Advance(fixedClock, time.Duration(12)*time.Hour)
	fmt.Println(clock.Now().Format("2006-01-02T15:04:05Z"))

	// Output:
	// 1970-01-01T12:00:00Z
}

func ExampleRevert() {
	epoch, _ := time.Parse("2006-01-02T15:04:05Z", "1970-01-01T00:00:00Z")
	fixedClock := NewFixed(epoch)

	clock := Revert(fixedClock, time.Duration(12)*time.Hour)
	fmt.Println(clock.Now().Format("2006-01-02T15:04:05Z"))

	// Output:
	// 1969-12-31T12:00:00Z
}

func ExampleTravel() {
	epoch, _ := time.Parse("2006-01-02T15:04:05Z", "1970-01-01T00:00:00Z")

	clock := Travel(SystemClock, epoch)
	now := clock.Now()
	if now.Sub(epoch) <= time.Duration(500)*time.Millisecond {
		fmt.Println("It's close enough to epoch.")
	}

	// Output:
	// It's close enough to epoch.
}

func ExamplePause() {
	now0 := SystemClock.Now()

	clock := Pause(SystemClock)
	time.Sleep(time.Duration(100) * time.Millisecond)
	now := clock.Now()

	if now.Sub(now0) <= time.Duration(100)*time.Millisecond {
		fmt.Printf("It's close enough to time when paused.")
	}

	// Output:
	// It's close enough to time when paused.
}

func ExampleResume() {
	abs := func(a, b time.Duration) time.Duration {
		t := a - b
		if t >= 0 {
			return t
		}

		return -t
	}

	epoch, _ := time.Parse("2006-01-02T15:04:05Z", "1970-01-01T00:00:00Z")

	clock := Resume(NewFixed(epoch))
	t0 := clock.Now()
	now0 := SystemClock.Now()
	time.Sleep(time.Duration(100) * time.Millisecond)
	t1 := Resume(clock).Now()
	now1 := SystemClock.Now()

	delta := now1.Sub(now0)
	deltaT := t1.Sub(t0)
	d := time.Duration(50) * time.Millisecond
	if abs(delta, deltaT) <= d {
		fmt.Printf("Clock resumed.")
	}

	// Output:
	// Clock resumed.
}
