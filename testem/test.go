package main

import "fmt"

func main() {

	p := fmt.Println
	var holbertonFounders = [3]string{"Rudy Rigot", "Sylvain Kalache", "Julien Barbier"}

	for i := 0; i < 3; i++ {
		p(holbertonFounders[i])
	}
}
