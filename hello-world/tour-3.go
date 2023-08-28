package main

import (
	"fmt"
	"runtime"
	"time"
)

func switchCasesOS() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func dayGreet() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	case t.Hour() < 21:
		fmt.Println("Good evening.")
	default:
		fmt.Println("Good night.")
	}
}

func main() {
	switchCasesOS()
	dayGreet()
	defer fmt.Println(time.Now().Weekday(), time.Now().Weekday()+1, time.Saturday)
}
