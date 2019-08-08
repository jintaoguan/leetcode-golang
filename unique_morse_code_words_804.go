package main

import (
	"fmt"
)

func main() {
	//[".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."]
	words := []string{"gin", "zen", "gig", "msg"}
	fmt.Println(uniqueMorseRepresentations(words))
}

func uniqueMorseRepresentations(words []string) int {
	letters := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	lettersDict := generateLettersDict(letters)
	uniqueDict := make(map[string]bool)
	for _, word := range words {
		code := calculateMorseCode(lettersDict, word)
		if _, ok := uniqueDict[code]; !ok {
			uniqueDict[code] = true
		}
	}
	return len(uniqueDict)
}

func generateLettersDict(letterCodes []string) map[rune]string {
	ret := make(map[rune]string)
	ch := rune(97)
	for _, code := range letterCodes {
		ret[ch] = code
		ch++
	}
	return ret
}

func calculateMorseCode(dict map[rune]string, word string) string {
	ret := ""
	for _, ch := range word {
		if code, ok := dict[ch]; ok {
			ret += code
		}
	}
	return ret
}

