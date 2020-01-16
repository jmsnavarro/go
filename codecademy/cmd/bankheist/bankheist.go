/*
Bank Heist
URL: https://www.codecademy.com/courses/learn-go/projects/bank-heist

We often include conditionals for when we need to account for different 
situations and outcomes. When it comes to coming up with a plan 
and executing it, there’s nothing more uncertain than a bank heist! 

(Watch any robbery themed movie if you need convincing).

In this project, we’re going to use conditionals to simulate this situation 
along with hi-jinks and hiccups that may pop up. Who knows? 

Maybe, just maybe, if all goes well, we’ll have a few more dollars after.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//initialize seed
	rand.Seed(time.Now().UnixNano())

	//variables
	isHeistOn := true
	eludedGuards :=  rand.Intn(100)
	
	//first condition: sneak past guards
	fmt.Println("")
	if eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
	} else {
		fmt.Println("Plan a better diguise next time?")
	}

	openedVault := rand.Intn(100)

	//show isHeistOn value
	fmt.Println(isHeistOn)

	//second condition: open the vault
	if isHeistOn && openedVault >= 70 {
		fmt.Println("Grab and GO!")
	} else if isHeistOn && openedVault < 70 {
		isHeistOn = false
		fmt.Println("Oh oh!")
	}

	leftSafely := rand.Intn(5)

	//third condition: leaving
	if isHeistOn {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("FAILED!")
		case 1:
			isHeistOn = false
			fmt.Println("Turns out vault doors don't open from the inside...")
		default:
			fmt.Println("Start the getaway car!")
		}
	}

	//wrapping up
	if isHeistOn {
		amtStolen := 10000 + rand.Intn(1000000)
		fmt.Printf("Amount stolen is $%d\n\n", amtStolen)
	}
}