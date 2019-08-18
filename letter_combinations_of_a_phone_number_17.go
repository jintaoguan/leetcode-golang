package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("23"))
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	dict := map[int][]rune{
		2: {'a', 'b', 'c'},
		3: {'d', 'e', 'f'},
		4: {'g', 'h', 'i'},
		5: {'j', 'k', 'l'},
		6: {'m', 'n', 'o'},
		7: {'p', 'q', 'r', 's'},
		8: {'t', 'u', 'v'},
		9: {'w', 'x', 'y', 'z'},
	}
	result := make([]string, 0)
	current := make([]rune, len(digits))
	letterCombinationsHelper(digits, 0, dict, current, &result)
	return result
}

func letterCombinationsHelper(digits string, index int, dict map[int][]rune, current []rune, result *[]string) {
	if index >= len(digits) {
		*result = append(*result, string(current))
		return
	}
	digit := int(digits[index] - '0')
	candidates := dict[digit]
	for _, candidate := range candidates {
		current[index] = candidate
		letterCombinationsHelper(digits, index+1, dict, current, result)
	}
}

