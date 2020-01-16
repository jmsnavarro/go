/*
Learn Go Functions: Interstellar Travel
URL: https://www.codecademy.com/courses/learn-go/projects/learn-go-functions-interstellar-travel

You’ll be writing Go functions to perform calculations and build out an 
interstellar travel agency!

As you work through the project, remember to check the output of 
your functions to ensure that you’re writing your functions correctly.
*/

package main

import "fmt"

// Create the function fuelGauge() here
func fuelGauge(fuel int) {
	fmt.Println("Fuel left:", fuel)
}

// Create the function calculateFuel() here
func calculateFuel(planet string) int {
	var fuel int

	switch planet {
	case "Venus":
		fuel = 300000
	case "Mercury":
		fuel = 500000
	case "Mars":
		fuel = 700000
	default:
		fuel = 0
	}

	return fuel
}

// Create the function greetPlanet() here
func greetPlanet(planet string) {
	fmt.Println("Welcome to " + planet + "!")
}

// Create the function cantFly() here
func cantFly() {
	fmt.Println("We do not have the available fuel to fly there.")
}

// Create the function flyToPlanet() here
func flyToPlanet(planet string, fuel int) int {
	var fuelRemaining, fuelCost int

	fuelRemaining = fuel
	fuelCost = calculateFuel(planet)

	if (fuelRemaining >= fuelCost) {
		greetPlanet(planet)
		//fuelRemaining -= fuelCost
		fuelRemaining = fuelRemaining - fuelCost
	}

	if (fuelCost > fuelRemaining) {
		cantFly()
	}

	return fuelRemaining
}

func main() {
	// Test your functions!

	// Create `planetChoice` and `fuel`
	var fuel int = 1000000
	var planetChoice = "Venus"
  
  	// And then liftoff!
	fuelGauge(flyToPlanet(planetChoice, fuel))
}