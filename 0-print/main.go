package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	p("Hello, we are Holberton School")
	time := time.Now()
	p("The date is", time)

	p("The year is", time.Year())
	p("The month is", time.Month())
	p("The day is", time.Day())
}
