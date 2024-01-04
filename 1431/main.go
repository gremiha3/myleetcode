package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var out []bool
	var maxval int = 0
	for _, val := range candies {
		if val >= maxval {
			maxval = val
		}
	}
	for _, val := range candies {
		if val+extraCandies < maxval {
			out = append(out, false)
		} else {
			out = append(out, true)
		}
	}
	return out
}
func main() {
	c := []int{12, 1, 12}
	var ec int = 10
	fmt.Printf("candies = %v, extraCandies = %d\n", c, ec)
	fmt.Printf("array = %v\n", kidsWithCandies(c, ec))
}
