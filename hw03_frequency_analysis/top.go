package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(str string) []string {
	hash := make(map[string]int)
	words := strings.Fields(str)

	if len(words) == 0 {
		return make([]string, 0)
	}

	for _, word := range words {
		hash[word]++
	}

	keys := make([]string, 0, len(hash))
	for key := range hash {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i int, j int) bool {
		comparableI := hash[keys[i]]
		comparableJ := hash[keys[j]]

		if comparableI-comparableJ > 0 {
			return true
		}

		if comparableI-comparableJ < 0 {
			return false
		}

		return keys[i] < keys[j]
	})

	if len(hash) < 10 {
		return keys[:len(hash)]
	}

	return keys[:10]
}
