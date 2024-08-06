package main

import (
	"fmt"
	"strings"
)

func wordFreqCount(word string) map[string]int {
	count := map[string]int{}
	compare := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	punctuations := ",.?/;:'!"
	result := strings.Split(word, " ")
	for _, part := range result {
		curr := make([]string, len(part))
		for _, char := range part {
			if strings.Contains(punctuations, string(char)) || !strings.Contains(compare, string(char)) {
				continue
			} else {
				curr = append(curr, string(char))
			}
		}
		count[strings.Join(curr, "")]++
	}
	return count
}

func palindromChecker(word string) bool {
	word = strings.Trim(word, " ")
	fmt.Println(word)
	curr := make([]string, 0, len(word))
	compare := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for _, itm := range word {
		if strings.Contains(compare, string(itm)) {
			curr = append(curr, string(itm))
		}
	}
	newWord := strings.Join(curr, "")
	newWord = strings.ToLower(newWord)
	fmt.Println(newWord)
	left, right := 0, len(newWord)-1
	for left < len(newWord) && right >= 0 {
		if newWord[left] != newWord[right] {
			return false
		}
		left++
		right--
	}
	return true
}
