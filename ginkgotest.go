package ginkgotest

import (
	"fmt"
	"time"
)

func Hello() string {
	fmt.Printf("sleeping\n")
	time.Sleep(10 * time.Second)
	return "hello"
}
