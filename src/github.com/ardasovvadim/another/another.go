package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "abc ymn noa"
	//sentence := "abc ymn dd no"
	sequnces, counters := MakeSequences(sentence)
	maxSequnce := FindMaxSequnce(sequnces, counters)
	fmt.Println(maxSequnce)
}

func IndexOfWord(words []string, except []int) int {
	lastWord := words[except[len(except)-1]]
	lastLetter := lastWord[len(lastWord)-1]

	for i := 0; i < len(words); i++ {
		if Contains(except, i) {
			continue
		}

		if lastLetter == words[i][0] {
			return i
		}
	}

	return -1
}

func Contains(array []int, value int) bool {
	for i := range array {
		if array[i] == value {
			return true
		}
	}
	return false
}

func MakeSequences(sentence string) ([]string, []int) {
	var words []string
	words = strings.Split(sentence, " ")

	sequences := make([]string, 0, len(words))
	counters := make([]int, 0, len(words))
	for i := 0; i < len(words)-1; i++ {
		sequence := words[i]
		counter := 0
		exceptIndexes := []int{i}

		for true {
			index := IndexOfWord(words, exceptIndexes)
			if index != -1 {
				sequence += " " + words[index]
				counter++
				exceptIndexes = append(exceptIndexes, index)
			} else {
				for j := range words {
					if !Contains(exceptIndexes, j) {
						sequence += " " + words[j]
					}
				}
				break
			}
		}

		if counter > 0 {
			sequences = append(sequences, sequence)
			counters = append(counters, counter)
		}
	}
	return sequences, counters
}

func FindMaxSequnce(sequnces []string, counters []int) string {
	iMax := -1
	max := 0
	for i := range sequnces {
		if counters[i] > max {
			max = counters[i]
			iMax = i
		}
	}
	if iMax != -1 {
		return sequnces[iMax]
	} else {
		return ""
	}
}
