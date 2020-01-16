# Cheatsheet


## Hello
```
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
```

## Variables and Types
```
package main

import "fmt"

func main() {
	var congrats string
  congrats = "Congratulations"
  congrats += "!!!"
  fmt.Println(congrats)
  
  var challenge string = "What else can you do?"
  fmt.Println(challenge)
  
  reminder := "Pratice is important!"
  fmt.Println(reminder)
}
```

## Conditions
```
package main

import "fmt"

func main() {
  if lessonLearned := true; lessonLearned {
    fmt.Println("Great job! You can continue on to the next exercise.")
  } else {
    fmt.Println("Practice makes perfect.")
  }
  // Create more conditions below!
  
}

```

## Functions
```
package main
import (
  "fmt"
  "time"
)

func joinTwoStrings(prefix string, suffix string) string {
  return prefix + suffix
}

func printOutDate() {
  t := time.Now()
  fmt.Println(t)
}

func waitForIt(message string) {
  defer fmt.Println("Done!")
  fmt.Println("Waiting")
  fmt.Println("Waiting")
  fmt.Println("Waiting")
  fmt.Println(message)
}

func main() {
  printOutDate()
  waitForIt(joinTwoStrings("Hi", " there"))
}
```

## Addresses and Pointers
```
package main

import "fmt"

func brainwash(saying string) {
	saying = "Beep Boop."
}

func main() {
	greeting := "Hello there!"
	
	brainwash(&greeting)
	
	fmt.Println("Greeting is now:", greeting)
}
```