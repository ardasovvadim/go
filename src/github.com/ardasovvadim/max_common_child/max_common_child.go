package max_common_child

import "errors"

func GetMaxCommonChild(child, parent string) (string, error) {
	words := make([]string, 0)
	for i := range child {
		word := ""
		index := -1
		for {
			index++
			position := i + len(word)
			if position >= len(child) {
				break
			}
			index = IndexNextSymbol(parent, child[i+len(word)], index)
			if index != -1 {
				word += string(parent[index])
			} else {
				break
			}
		}
		if len(word) != 0 {
			words = append(words, word)
		}
	}
	if len(words) == 1 {
		return "", errors.New("Error")
	}
	maxWord := ""
	for _, word := range words {
		if len(maxWord) < len(word) {
			maxWord = word
		}
	}
	return maxWord, nil
}

func IndexNextSymbol(str string, findSymbol uint8, index int) int {
	for i := index; i < len(str); i++ {
		if str[i] == findSymbol {
			return i
		}
	}
	return -1
}
