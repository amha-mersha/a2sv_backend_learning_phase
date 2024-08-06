package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func wordFreqCount(word string) map[string]int {
	count := map[string]int{}
	compare := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	punctuations := ",.?!'\""
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
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter a long sentence: ")
	input, _ := reader.ReadString('\n')
	fmt.Println(wordFreqCount(input))
}
