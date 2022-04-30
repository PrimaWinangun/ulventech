package repository

import (
	"github.com/PrimaWinangun/ulventech/model"
	"os"
	"regexp"
	"sort"
)

type wcRepository struct {
}

// NewWcRepository
// Initiate the Word Count Repository Layer
func NewWcRepository() model.WordCountRepository {
	return wcRepository{}
}

// CountWord
// Function for counting all the word inside the files and return top ten words along with number
// used in file in the descending order
func (w wcRepository) CountWord(filename string) ([]model.WordCount, error) {
	reg := regexp.MustCompile("[a-zA-Z']+")
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	text := string(file)
	matches := reg.FindAllString(text, -1)

	words := make(map[string]int)
	for _, match := range matches {
		words[match]++
	}

	var wc []model.WordCount
	for x, v := range words {
		wc = append(wc, model.WordCount{Word: x, Count: v})
	}

	sort.Slice(wc, func(i, j int) bool {
		return wc[i].Count > wc[j].Count
	})

	return wc, nil
}
