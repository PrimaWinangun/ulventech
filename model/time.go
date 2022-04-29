package model

import "github.com/gin-gonic/gin"

type ValidTime struct {
	Input []int `json:"input"`
}

type ValidTimeResponse struct {
	List  []string `json:"list"`
	Count int      `json:"count"`
}

type ValidTimeUseCase interface {
	ListAvailable(ctx *gin.Context, time ValidTime) (ValidTimeResponse, error)
}

type ValidTimeRepository interface {
	Permutation(t ValidTime) ([]string, error)
}
