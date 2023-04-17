package main

import (
	"fmt"
	"testing"
	"time"
)

func TestReverseRunes(t *testing.T) {
	startTime := time.Now()

	d := 100 * time.Microsecond
	fmt.Println(d) // Output: 100µs

	value := 100 // value is of type int

	d2 := time.Duration(value) * time.Millisecond
	fmt.Println(d2) // Output: 100ms

	ms := int64(d2 / time.Millisecond)
	fmt.Println("ms:", ms)                         // Output: ms: 100
	fmt.Println("ns:", int64(d2/time.Nanosecond))  // ns: 100000000
	fmt.Println("µs:", int64(d2/time.Microsecond)) // µs: 100000
	fmt.Println("ms:", int64(d2/time.Millisecond)) // ms: 100

	// End Time request
	endTime := time.Now()

	// execution time
	latencyTime := endTime.Sub(startTime)
	fmt.Println(latencyTime)
	fmt.Printf("latency : %s %v \n", latencyTime, latencyTime)
}
