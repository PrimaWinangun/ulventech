package repository

import (
	"os"
	"regexp"
	"sort"
	"ulventech/model"
)

type wcRepository struct {
}

func NewWcRepository() model.WordCountRepository {
	return wcRepository{}
}

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
