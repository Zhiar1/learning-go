package main

import (
	"fmt"
	"time"
)

func main() {
	dttz := time.Now()
	fmt.Println("Current date, time, tz: ", dttz)

	dttzExplicit := time.Date(2023, time.December, 20, 6, 0, 0, 0, time.Local)
	fmt.Println("Declared date/time: ", dttzExplicit)
	fmt.Println("Formatted declared date/time: ", dttzExplicit.Format(time.ANSIC))

	parsedDt, _ := time.Parse(time.ANSIC, "Wed Dec 20 06:00:00 2023")
	fmt.Printf("The type of the variable is %T\n", parsedDt)
}
