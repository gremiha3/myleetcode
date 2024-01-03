package main

import "fmt"

func numbersOfSteps(num int) int {
	var step int
	for num != 0 {
		if num&1 == 0 {
			num = num >> 1
		} else {
			num--
		}
		step++
	}
	return step
}
func main() {
	var num int = 123
	fmt.Printf("num = %d, steps = %d\n", num, numbersOfSteps(num))
}
