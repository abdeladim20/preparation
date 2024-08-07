package main

import "fmt"

func union(str1, str2 string) string {
	// Create a map to store the characters we've seen
	charMap := make(map[rune]bool)
	var result []rune

	// Helper function to process each string
	processString := func(s string) {
		for _, char := range s {
			if !charMap[char] {
				charMap[char] = true
				result = append(result, char)
			}
		}
	}

	// Process both strings
	processString(str1)
	processString(str2)

	return string(result)
}

func main() {
	s1 := "hello"
	s2 := "hello"

	result := union(s1, s2)
	fmt.Println(result)
}
