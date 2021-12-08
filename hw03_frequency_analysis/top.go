package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(text string) []string {
	wordsMap := getWordsCountMap(text)

	return getTopSortedWords(wordsMap, 10)
}

func getTopSortedWords(words map[string]int, top int) []string {
	type WordCounter struct {
		Word  string
		Count int
	}
	type WordCounterSlice []WordCounter

	wordCounterList := make(WordCounterSlice, len(words))
	i := 0
	for k, v := range words {
		wordCounterList[i] = WordCounter{k, v}
		i++
	}

	sort.Slice(wordCounterList, func(i, j int) bool {
		if wordCounterList[i].Count == wordCounterList[j].Count {
			return wordCounterList[i].Word < wordCounterList[j].Word
		}
		return wordCounterList[i].Count > wordCounterList[j].Count
	})

	var result []string
	for i := 0; i < len(wordCounterList); i++ {
		if i >= top {
			break
		}
		result = append(result, wordCounterList[i].Word)
	}

	return result
}

func getWordsCountMap(text string) map[string]int {
	words := make(map[string]int, 0)
	for _, s := range strings.Fields(text) {
		if _, ok := words[s]; ok {
			words[s] += 1
		} else {
			words[s] = 1
		}
	}

	return words
}
