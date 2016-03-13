package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	p("Hello, we are Holberton School")
	time := time.Now()
	p("the date is", time)

	p("the year is", time.Year())
	p("the month is", time.Month())
	p("the day is", time.Day())
}
