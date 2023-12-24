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

	fmt.Print("Enter the first number: ")
	value1, _ := reader.ReadString('\n')

	fmt.Print("Enter the second number: ")
	value2, _ := reader.ReadString('\n')

	fmt.Println(calculate(value1, value2))
}

func calculate(value1 string, value2 string) float64 {
	floatedStringV1, err := strconv.ParseFloat(strings.TrimSpace(value1), 64)
	floatedStringV2, err1 := strconv.ParseFloat(strings.TrimSpace(value2), 64)
	var result float64

	if err != nil {
		fmt.Println(err)
		panic("The first value gave an error, please make sure you input a correct numerical value!\n")
	} else if err1 != nil {
		fmt.Println(err1)
		panic("The second value gave an error, please make sure you input a correct numerical value!\n")
	} else {
		result = floatedStringV1 + floatedStringV2
	}
	return result
}
