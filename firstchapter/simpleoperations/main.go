package main

import (
	"fmt"
	"reflect"
)

func main() {
	var str string = "Go!"
	fmt.Printf("The variable's type is: %T\n", str)

	str1 := "Go!!"
	fmt.Println("The variable's type is:", reflect.TypeOf(str))

	if reflect.TypeOf(str) == reflect.TypeOf(str1) {
		fmt.Printf("the explicitly declared string variable is the same type as the implicitly declared string variable \n")
	} else {
		fmt.Println("they are not the same")
	}
}
