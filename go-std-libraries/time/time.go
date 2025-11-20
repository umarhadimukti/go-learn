package time

import (
	"time"
	"fmt"
)

func expensiveCall() {
	time.Sleep(5 * time.Nanosecond)
}

func DemoTimePackage() {
	t0 := time.Now()
	expensiveCall()
	t1 := time.Now()
	fmt.Printf("Expensive call took %v\n", t1.Sub(t0))

	ut := time.Date(2025, time.October, 30, 12, 0, 0, 0, time.UTC)
	fmt.Println("Unix Time:", ut.Unix())
	nt := time.Unix(ut.Unix(), 0).UTC()
	fmt.Println("Converted Time:", nt)

	fmt.Println("Get hour:", ut.Hour())
	fmt.Println("Get weekday:", ut.Weekday())
	fmt.Println("Location:", ut.Location())

	lt := time.Now().Local()
	fmt.Println("Local Time:", lt)
	formattedLt := lt.Format("2006-01-02 15:04:05") // default format in Go
	fmt.Println("Formatted Local Time:", formattedLt)
}