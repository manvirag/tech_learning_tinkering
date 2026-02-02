// Package golang_library provides time and random utilities for LLD, matching, competitive, and general use.
package golang_library

import (
	"fmt"
	"time"
)

// ============================================================================
// TIME - Quick reference for LLD, matching, competitive coding
// ============================================================================

// TimeDemo runs time-related examples.
func TimeDemo() {
	// ------------------------------------------------------------------------
	// 1. Current time
	// ------------------------------------------------------------------------
	now := time.Now()
	fmt.Println("Now:", now)

	// ------------------------------------------------------------------------
	// 2. Unix timestamp (epoch) - int64 seconds
	// ------------------------------------------------------------------------
	epochSec := time.Now().Unix()
	epochMilli := time.Now().UnixMilli()
	fmt.Printf("Epoch sec: %d, milli: %d\n", epochSec, epochMilli)

	// Timestamp from time
	t := time.Now()
	ts := t.Unix()
	fmt.Println("Timestamp:", ts)

	// Time from timestamp
	fromTs := time.Unix(epochSec, 0)
	fmt.Println("From timestamp:", fromTs)

	// ------------------------------------------------------------------------
	// 3. Duration - create and use
	// ------------------------------------------------------------------------
	d1 := 5 * time.Second
	d2 := 500 * time.Millisecond
	d3 := 2 * time.Minute
	d4 := 24 * time.Hour
	_ = d4
	fmt.Println("Duration 5s:", d1, "500ms:", d2, "2min:", d3)

	// ------------------------------------------------------------------------
	// 4. Add / subtract time
	// ------------------------------------------------------------------------
	future := now.Add(10 * time.Second)
	past := now.Add(-1 * time.Hour)
	_ = past
	fmt.Println("Now + 10s:", future)

	// ------------------------------------------------------------------------
	// 5. Time difference (duration between two times)
	// ------------------------------------------------------------------------
	start := time.Now()
	// ... do work ...
	elapsed := time.Since(start)
	fmt.Println("Elapsed:", elapsed)

	// Or: end - start
	end := time.Now()
	diff := end.Sub(start)
	fmt.Println("Diff:", diff)

	// As milliseconds / seconds (int64)
	ms := time.Since(start).Milliseconds()
	sec := time.Since(start).Seconds()
	fmt.Printf("Elapsed ms: %d, sec: %.2f\n", ms, sec)

	// ------------------------------------------------------------------------
	// 6. Compare times
	// ------------------------------------------------------------------------
	deadline := time.Now().Add(5 * time.Minute)
	if time.Now().Before(deadline) {
		fmt.Println("Before deadline")
	}
	if time.Now().After(deadline) {
		fmt.Println("After deadline")
	}
	t1 := time.Now()
	t2 := t1.Add(1 * time.Second)
	fmt.Println("t1.Before(t2):", t1.Before(t2))
	fmt.Println("t1.After(t2):", t1.After(t2))
	fmt.Println("t1.Equal(t2):", t1.Equal(t2))

	// ------------------------------------------------------------------------
	// 7. Sleep
	// ------------------------------------------------------------------------
	// time.Sleep(2 * time.Second)
	fmt.Println("Sleep: time.Sleep(2 * time.Second)")

	// ------------------------------------------------------------------------
	// 8. Format time (for logging, display)
	// ------------------------------------------------------------------------
	// Reference: Mon Jan 2 15:04:05 MST 2006 (01/02 03:04:05 PM '06 -0700)
	fmt.Println("RFC3339:", now.Format(time.RFC3339))
	fmt.Println("Custom:", now.Format("2006-01-02 15:04:05"))
	fmt.Println("Date only:", now.Format("2006-01-02"))
	fmt.Println("Time only:", now.Format("15:04:05"))

	// ------------------------------------------------------------------------
	// 9. Parse time string
	// ------------------------------------------------------------------------
	parsed, _ := time.Parse("2006-01-02", "2024-12-25")
	fmt.Println("Parsed:", parsed)
	parsed2, _ := time.Parse(time.RFC3339, "2024-12-25T10:30:00Z")
	fmt.Println("Parsed RFC3339:", parsed2)

	// ------------------------------------------------------------------------
	// 10. Timer (one-shot delay) - for timeouts
	// ------------------------------------------------------------------------
	timer := time.NewTimer(1 * time.Second)
	// <-timer.C  // block until fired
	timer.Stop() // cancel if not fired yet
	fmt.Println("Timer: time.NewTimer(d), <-timer.C, timer.Stop()")

	// ------------------------------------------------------------------------
	// 11. Ticker (repeated interval) - for periodic tasks
	// ------------------------------------------------------------------------
	ticker := time.NewTicker(1 * time.Second)
	// for t := range ticker.C { ... }
	ticker.Stop()
	fmt.Println("Ticker: time.NewTicker(d), range ticker.C, ticker.Stop()")

	// ------------------------------------------------------------------------
	// 12. Timeout pattern (select with time.After)
	// ------------------------------------------------------------------------
	// select {
	// case result := <-ch:
	//     return result
	// case <-time.After(5 * time.Second):
	//     return errors.New("timeout")
	// }
	fmt.Println("Timeout: select { case <-time.After(d): ... }")

	// ------------------------------------------------------------------------
	// 13. Truncate / Round
	// ------------------------------------------------------------------------
	trunc := now.Truncate(1 * time.Hour)
	rounded := now.Round(1 * time.Hour)
	fmt.Println("Truncate to hour:", trunc)
	fmt.Println("Round to hour:", rounded)

	// ------------------------------------------------------------------------
	// 14. Store date (year, month, day)
	// ------------------------------------------------------------------------
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("Date: %d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	// ------------------------------------------------------------------------
	// 15. Time of day (duration since midnight) - for "10:11 AM" style
	// ------------------------------------------------------------------------
	midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	timeOfDay := now.Sub(midnight)
	fmt.Println("Time of day (since midnight):", timeOfDay)
}
