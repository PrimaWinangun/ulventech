package controller

import (
	"fmt"
	"github.com/PrimaWinangun/ulventech/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type wcDeliveryHttp struct {
	wcUsecase model.WordCountUseCase
}

func NewWordCount(rg *gin.RouterGroup, wcUsecase model.WordCountUseCase) {
	wd := wcDeliveryHttp{wcUsecase: wcUsecase}

	v1 := rg.Group("/word-count")
	v1.POST("upload", wd.CountWord)
}

// CountWord godoc
// @Summary upload file via rest api http & return top ten most used words with the number of used
// @Description return top ten most used words along with how many times they occur in the text.
// @Tags root
// @Accept */*
// @Param payload path string true "file type with words contained"
// @Produce json
// @Success 200 {array} model.WordCount
// @Router /api/v1/word-count/upload [post]
func (wd wcDeliveryHttp) CountWord(c *gin.Context) {
	file, err := c.FormFile("file")

	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
		return
	}

	processedWord, err := wd.wcUsecase.ProcessFile(file, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Failed To Process File",
			"name":   file.Filename,
			"error":  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "File Processed Successfully",
			"name":   file.Filename,
			"data":   processedWord,
		})
	}
}
