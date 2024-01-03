package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var out = []bool{
		true,
		true,
		false,
	}

	return out
}
func main() {
	c := []int{2, 3, 5, 1, 3}
	ec := 3
	fmt.Printf("candies = %v, extraCandies = %d\n", c, ec)
	fmt.Printf("array = %v", kidsWithCandies(c, ec))
}
