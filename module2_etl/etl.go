package main

import (
	"fmt"
	"strings"
)

func main() {
	oldScoring := map[int][]string{
		1:  {"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"},
		2:  {"D", "G"},
		3:  {"B", "C", "M", "P"},
		4:  {"F", "H", "V", "W", "Y"},
		5:  {"K"},
		8:  {"J", "X"},
		10: {"Q", "Z"},
	}
	newScoring := make(map[string]int)
	for score, letters := range oldScoring {
		for _, letter := range letters {
			character := strings.ToLower(letter)
			newScoring[character] = score
		}
	}
	for letter, score := range newScoring {
		fmt.Println(letter, " : ", score)
	}
}
