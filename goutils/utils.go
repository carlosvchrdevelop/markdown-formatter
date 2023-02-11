package goutils

import "time"

func Timing (f func()) float32 {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	return float32(elapsed.Microseconds())/1000
}