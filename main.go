package main

import (
	"fmt"
	"time"
)

func main() {
	newYear := time.Date(2026, time.January, 1, 0, 0, 0, 0, time.Local)

	until := time.Until(newYear).Hours() / 24

	fmt.Println(int(until) + 1)
}
