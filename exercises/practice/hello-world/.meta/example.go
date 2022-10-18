package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now() //It will return time.Time object with current timestamp
	fmt.Printf("time.Time %s\n", t)

	tUnix := t.Unix()
	fmt.Printf("timeUnix: %d\n", tUnix)

	tUnixNano := t.UnixNano()
	fmt.Printf("timeUnixNano: %d\n", tUnixNano)
}
