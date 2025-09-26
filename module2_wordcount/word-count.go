package main

import (
	"fmt"
	"regexp"
	"strings"
)

type Frequency map[string]int

func wordcount(text string) Frequency {
	freq := make(Frequency)
	text = strings.ToLower(text)
	re := regexp.MustCompile(`[a-z0-9]+(?:'[a-z0-9]+)?`)
	words := re.FindAllString(text, -1)
	for _, w := range words {
		freq[w]++
	}
	return freq
}
func main() {
	text := "hello hi how are you hi hello"
	result := wordcount(text)
	for word, count := range result {
		fmt.Println(word, " : ", count)
	}

}
