package hw03frequencyanalysis

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func Top10(str string) ([]string, error) {
	hash := make(map[string]int)
	words := strings.Fields(str)

	if len(words) == 0 {
		return make([]string, 0), nil
	}

	for _, word := range words {
		hash[word]++
	}

	keys := make([]string, 0)
	for key := range hash {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i int, j int) bool {
		comparableI := hash[keys[i]]
		comparableJ := hash[keys[j]]
		if comparableI-comparableJ > 0 {
			return true
		} else if comparableI-comparableJ < 0 {
			return false
		} else {
			return keys[i] < keys[j]
		}
	})

	var err error
	defer func() {
		if val := recover(); val != nil {
			err = fmt.Errorf("panic occurs")
		}
	}()

	// if total number of words less than 10 then use it as index
	idx := int(math.Min(float64(len(hash)), 10))
	return keys[:idx], err
}
