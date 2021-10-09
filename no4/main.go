package main

import "fmt"

func areAnagram(wordA string, wordB string) bool {
	var anagramChecker = make(map[rune]int)

	if len(wordA) != len(wordB) {
		return false
	}

	runes1 := []rune(wordA)
	runes2 := []rune(wordB)
	for i := 0; i < len(runes1); i++ {
		anagramChecker[runes1[i]]++
		anagramChecker[runes2[i]]--
	}

	for k := range anagramChecker {
		if anagramChecker[k] != 0 {
			return false
		}
	}

	return true
}

func anagramWordsClassifier(words []string) [][]string {
	var strsByLen = make(map[int][]string)
	var result = make([][]string, 0)

	for i := 0; i < len(words); i++ {
		strsByLen[len(words[i])] = append(strsByLen[len(words[i])], words[i])
	}

	for k := range strsByLen {
		var strs = make(map[int]string)

		for i := 0; i < len(strsByLen[k]); i++ {
			strs[i] = strsByLen[k][i]
		}

		for len(strs) > 0 {
			var anagramChecker = make([]string, 0)
			strToCheck := ""

			for k2 := range strs {
				if strToCheck == "" {
					strToCheck = strs[k2]
					anagramChecker = append(anagramChecker, strToCheck)
					delete(strs, k2)
				} else {
					if areAnagram(strToCheck, strs[k2]) {
						anagramChecker = append(anagramChecker, strs[k2])
						delete(strs, k2)
					}
				}
			}

			strToCheck = ""
			result = append(result, anagramChecker)
		}

	}

	return result
}

func main() {
	var strs = []string{"makan", "mati", "itam", "set", "aut", "tau", "uut", "tua", "tes"}
	fmt.Println(anagramWordsClassifier(strs))
}
