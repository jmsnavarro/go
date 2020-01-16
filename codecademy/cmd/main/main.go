package main

import (
	"fmt"
	t "time"
)

func main() {
	/*
  		  __      _
		o'')}____//
		`_/      )
		(_(_/-(_/

		Today is 2020-01-15 08:09:35.4944677 +0800 +08 m=+0.000330701
		Total price is: 780.068

		Ahoy! Land Lovers! 6969
	*/

	// print a dog
	fmt.Println("")
	fmt.Println("  __      _")
	fmt.Println("o'')}____//")
	fmt.Println("`_/      ) ")
	fmt.Println("(_(_/-(_/  ")
	fmt.Println("")

	var currentDateTime = t.Now()

	// show something like nothing
	fmt.Println("Today is" , currentDateTime)
}