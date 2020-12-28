package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	sec := now.Unix()
	nanosec := now.UnixNano()
	lessThanOneSec := nanosec - (sec * 1000 * 1000 * 1000)

	fmt.Printf("%d.%09d\n", sec, lessThanOneSec)
}
