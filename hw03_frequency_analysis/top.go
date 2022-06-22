package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var separator = regexp.MustCompile(`[.,!:?;"'\s]+-*\s*`)

type wordFrequency struct {
	len  int
	word string
}

func Top10(text string) []string {
	res := []string{}

	if len(text) > 0 {
		wordsList := separator.Split(text, -1)
		wordFrequencies := calculateWords(wordsList)
		topWordFrequencies := getTop10(wordFrequencies)
		setTopWords(topWordFrequencies, &res)
	}

	return res
}

func calculateWords(wordsList []string) []wordFrequency {
	wordCountMap := map[string]int{}
	for _, word := range wordsList {
		wordCountMap[strings.ToLower(word)]++
	}

	wordFrequencies := make([]wordFrequency, 0, len(wordCountMap))
	for word, count := range wordCountMap {
		wordFrequencies = append(wordFrequencies, wordFrequency{count, word})
	}

	return wordFrequencies
}

func getTop10(wordFrequencies []wordFrequency) []wordFrequency {
	sort.Slice(wordFrequencies, func(left, right int) bool {
		if wordFrequencies[left].len > wordFrequencies[right].len {
			return true
		}
		if wordFrequencies[left].len < wordFrequencies[right].len {
			return false
		}
		return wordFrequencies[left].word < wordFrequencies[right].word
	})

	rightBorder := 10
	if len(wordFrequencies) < 10 {
		rightBorder = len(wordFrequencies)
	}

	return wordFrequencies[:rightBorder]
}

func setTopWords(wordFrequencies []wordFrequency, res *[]string) {
	for _, v := range wordFrequencies {
		*res = append(*res, v.word)
	}
}
