package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	/*
		We are reading an input from the user, converting that input into float after trimming white space.
		If there is no errors during this process we will return the input back to the user if not return the error.
	*/
	fmt.Print("Enter a number: ")
	numInput, _ := reader.ReadString('\n')
	inputToFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Make sure to input only numerical values!")
	} else {
		fmt.Print("Entered value: ", inputToFloat)
	}
}
