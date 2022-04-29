package model

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type WordCount struct {
	Word  string `json:"word"`
	Count int    `json:"count"`
}

type WordCountRepository interface {
	CountWord(filename string) ([]WordCount, error)
}

type WordCountUseCase interface {
	ProcessFile(file *multipart.FileHeader, ctx *gin.Context) ([]WordCount, error)
}
