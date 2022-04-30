package usecase

import (
	"fmt"
	"github.com/PrimaWinangun/ulventech/model"
	"github.com/gin-gonic/gin"
	"time"
)

type vTUseCase struct {
	vtRepo model.ValidTimeRepository
}

// NewVTUseCase
// Initiate the Valid Time Use Case layer
func NewVTUseCase(vtRepo model.ValidTimeRepository) model.ValidTimeUseCase {
	return &vTUseCase{vtRepo: vtRepo}
}

// ListAvailable
// Function for check the content of the list if it is valid time
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
