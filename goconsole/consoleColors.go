package goconsole

import "runtime"

// Reset to default color
var Reset = "\033[0m"

// Red color
var Red = "\033[31m"

// Green color
var Green = "\033[32m"

// Yellow color
var Yellow = "\033[33m"

// Blue color
var Blue = "\033[34m"

// Purple color
var Purple = "\033[35m"

// Cyan color
var Cyan = "\033[36m"

// Gray color
var Gray = "\033[37m"

// White color
var White = "\033[97m"

func init() {
	if runtime.GOOS == "windows" {
		Reset = ""
		Red = ""
		Green = ""
		Yellow = ""
		Blue = ""
		Purple = ""
		Cyan = ""
		Gray = ""
		White = ""
	}
}
