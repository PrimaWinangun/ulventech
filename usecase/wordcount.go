package usecase

import (
	"fmt"
	"github.com/PrimaWinangun/ulventech/model"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

var basePath string = "upload"

type wcUsecase struct {
	wcRepo model.WordCountRepository
}

// NewWcUseCase
// Initiate the Word Count Use Case Layer
func NewWcUseCase(wcRepo model.WordCountRepository) model.WordCountUseCase {
	return wcUsecase{wcRepo: wcRepo}
}

// ProcessFile
// Function for upload the file and return top ten words along with the number of used in the file
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

	return word[:10], nil
}
