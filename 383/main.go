package main

import (
	"fmt"
)

func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[byte]int)
	for ind := 0; ind < len(magazine); ind++ {
		fmt.Println(magazine[ind])
		m[magazine[ind]]++
	}
	for key, val := range m {
		fmt.Printf("key = %d, val = %d\n", key, val)
	}
	for ind := 0; ind < len(ransomNote); ind++ {
		if m[ransomNote[ind]] > 0 {
			fmt.Printf("char %d is entry\n", ransomNote[ind])
			m[ransomNote[ind]]--
		} else {
			return false
		}
	}
	return true
}
func main() {
	var ransomNote string = "bb"
	var magazine string = "aab"
	fmt.Printf("ransome Note = %s\nmagazine = %s\nconstructed = %v\n", ransomNote, magazine, canConstruct(ransomNote, magazine))
}
