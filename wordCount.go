package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	result := make(map[string]int)
	for i := range words {
		result[words[i]] = result[words[i]] + 1;
	}
	return result
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    sentence, _ := reader.ReadString('\n')
    sentence = strings.TrimSuffix(sentence, "\n")

    /* in case the program is launched on linux, \r won't be present there, leaving the string unchanged */
    sentence = strings.TrimSuffix(sentence, "\r")
    fmt.Println(WordCount(sentence))
}