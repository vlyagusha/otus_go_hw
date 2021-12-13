package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

func Top10(text string) []string {
	wordsMap := getWordsCountMap(text)

	return getTopSortedWords(wordsMap, 10)
}

var wordRegExpr = regexp.MustCompile(`[\wа-яёА-ЯЁ]+[-\wа-яёА-ЯЁ]*`)

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
	wordsCountMap := make(map[string]int)
	words := wordRegExpr.FindAllString(strings.ToLower(text), -1)
	for _, s := range words {
		wordsCountMap[s]++
	}

	return wordsCountMap
}
