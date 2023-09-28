package bloomfilter

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/datadog/mmh3"
)

const ArraySize = 1024000 

var BitMap [ArraySize]int

var WordsSource = "/usr/share/dict/words"

func LoadWords() (words []string, err error) {
	fd, err := os.Open(WordsSource)
	if err != nil {
		return
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		word := scanner.Text()
		words = append(words, word)
	}
	return words, nil
}

func Hash(key []byte) (hash uint32) {
	hash = mmh3.Hash32(key)

	return
}

func LoadMap() {
	words, _ := LoadWords()
	for _, word := range words {
		idx := Hash([]byte(word)) % ArraySize
		BitMap[idx] = 1	
	}
}

func FindWords(phrase string) (exists bool) {
	words := strings.Split(phrase, " ")
	exists = true 
	for _, word := range words {
		idx := Hash([]byte(word)) % ArraySize
		value := BitMap[idx]
		fmt.Println(value, idx)
		if value == 0 {
			return false
		}
	}
	return
}
