package main

import "math/rand"
import "fmt"

func main() {
	var y int
	var n int

	randomNumber := rand.Intn(100)
	p := fmt.Println
	if randomNumber > 50 {
		y = randomNumber
		p("my random number is", y, "and is greater than 50")
	} else {
		n = randomNumber
		p("my random number is", n, "and is not greater than 50")
	}

	school := "Holberton School"
	if school == "Holberton School" {
		p("I am a student of the Holberton School")
	} else {
		p("I am not a student of the Holberton School")
	}

	var beautifulWeather = true
	if beautifulWeather == true {
		p("It's a beautiful weather!")
	} else {
		p("It's not beautiful weather!")
	}

	var holbertonFounders = [3]string{"Rudy Rigot", "Sylvain Kalache", "Julien Barbier"}

	for i := 0; i < 3; i++ {
		p(holbertonFounders[i], "is a founder")
	}

}
