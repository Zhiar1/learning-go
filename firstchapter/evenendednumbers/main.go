package main

import "fmt"

func main() {
	evenEndedNumberCount := 0
	for a := 1000; a <= 9999; a++ {
		for b := a; b <= 9999; b++ {
			numb := a * b
			str := fmt.Sprintf("%d", numb)
			if str[0] == str[len(str)-1] {
				evenEndedNumberCount++
			}
		}
	}
	println("Count of how many numbers have the same first and last digit: ", evenEndedNumberCount)
}
