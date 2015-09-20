package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
	This function uses dynamic programming to find out if a particular word is a compound word or not
	It takes as parameters a word and a dictionary/map to find the longest word
*/
func isWordCompound(word string, dict map[string]bool) bool {
	bools := make([]bool, len(word)+1)
	for i := 1; i < len(word); i++ {
		if !bools[i] {
			if _, ok := dict[word[0:i]]; ok {
				bools[i] = true
			}
		}
		if bools[i] {
			for j := i + 1; j <= len(word); j++ {
				if !bools[j] {
					if _, ok := dict[word[i:j]]; ok {
						bools[j] = true
					}
				}
				if j == len(word) && bools[j] {
					return true
				}
			}
		}
	}
	return false
}

/*
	This function finds out the longest compund word by calling is compund word on every word in the list
	And then determines which one is the longest
*/
func longestCompound(dict map[string]bool) string {
	maxWordLength := 0
	longestWord := ""
	for word, _ := range dict {
		if isWordCompound(word, dict) {
			if len(word) > maxWordLength {
				maxWordLength = len(word)
				longestWord = word
			}
		}
	}
	return longestWord
}

/*
This  function uses a scanner to create a map of word from the file provided
*/
func makeDict(fname string) map[string]bool {
	file, err := os.Open(fname)
	if err != nil {
		fmt.Errorf("Opening the file had an error: %v\n", err)
	}
	dict := make(map[string]bool)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dict[scanner.Text()] = true
	}
	er := scanner.Err()
	if er != nil {
		fmt.Errorf("Error scanning the file: %v\n", er)
	}
	return dict
}

func main() {
	var fileName string
	if len(os.Args) == 1 {
		fileName = "word.list"
	} else if len(os.Args) == 2 {
		fileName = os.Args[1]
	}
	dict := makeDict(fileName)
	longestCompoundWord := longestCompound(dict)
	fmt.Printf("The longest compund word in the file is: %s\n", longestCompoundWord)
}
