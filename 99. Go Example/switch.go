package main

import "time"

func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}

func isWeekend() bool {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		return true
	default:
		return false
	}
}

func main() {
	var value = unhex(byte('a'))
	println("'a':", value)

	println("Today:", time.Now().Weekday().String())
	if isWeekend() {
		println("Weekend")
	} else {
		println("Weekday")
	}
}
