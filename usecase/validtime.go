package usecase

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"ulventech/model"
)

type vTUseCase struct {
	vtRepo model.ValidTimeRepository
}

func NewVTUseCase(vtRepo model.ValidTimeRepository) model.ValidTimeUseCase {
	return &vTUseCase{vtRepo: vtRepo}
}

func (v vTUseCase) ListAvailable(ctx *gin.Context, mt model.ValidTime) (rv model.ValidTimeResponse, err error) {
	vt, err := v.vtRepo.Permutation(mt)
	if err != nil {
		return model.ValidTimeResponse{}, err
	}
	layout := fmt.Sprintf("15:04")
	for _, v := range vt {
		_, err := time.Parse(layout, v)
		if err == nil {
			rv.List = append(rv.List, v)
			rv.Count += 1
		}
	}

	return rv, nil
}
