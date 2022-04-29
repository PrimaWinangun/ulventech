package usecase

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"ulventech/model"
)

var basePath string = "upload"

type wcUsecase struct {
	wcRepo model.WordCountRepository
}

func NewWcUsecase(wcRepo model.WordCountRepository) model.WordCountUseCase {
	return wcUsecase{wcRepo: wcRepo}
}

func (wc wcUsecase) ProcessFile(file *multipart.FileHeader, ctx *gin.Context) ([]model.WordCount, error) {
	dest := fmt.Sprintf("%v/%v", basePath, file.Filename)
	err := ctx.SaveUploadedFile(file, dest)
	if err != nil {
		return nil, err
	}

	word, err := wc.wcRepo.CountWord(dest)
	if err != nil {
		return nil, err
	}

	return word, nil
}
