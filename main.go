package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	sec := now.Unix()
	nanosec := now.UnixNano()
	less_than_one_sec := nanosec - (sec * 1000 * 1000 * 1000)

	fmt.Printf("%d.%d\n", sec, less_than_one_sec)
}
