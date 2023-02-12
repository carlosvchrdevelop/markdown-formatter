package goutils

import "time"

// Timing counts the elapsed time during the execution of the callback
func Timing(f func()) float32 {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	return float32(elapsed.Microseconds()) / 1000
}
