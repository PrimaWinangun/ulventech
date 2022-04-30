package controller

import (
	"fmt"
	"github.com/PrimaWinangun/ulventech/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type vtDeliveryHttp struct {
	vtUseCase model.ValidTimeUseCase
}

type request struct {
	Input string `json:"input"`
}

func NewValidTime(rg *gin.RouterGroup, vtUseCase model.ValidTimeUseCase) {
	vt := vtDeliveryHttp{vtUseCase: vtUseCase}

	v1 := rg.Group("valid-time")
	v1.POST("process", vt.FindValidTime)
}

// FindValidTime godoc
// @Summary find all combination of valid time from slice of integer
// @Description return all available time from the combination of four value of slice of integer.
// @Tags root
// @Accept application/json
// @Produce json
// @Param payload body request true "string value with space between, ex: 1 2 3 4"
// @Success 200 {object} model.ValidTimeResponse
// @Router /api/v1/valid-time/process [post]
func (v vtDeliveryHttp) FindValidTime(c *gin.Context) {
	in := request{}
	err := c.ShouldBindJSON(&in)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
		return
	}

	val := strings.Split(in.Input, " ")
	if len(val) != 4 {
		c.String(http.StatusBadRequest, fmt.Sprintf("err: invalid input"))
		return
	}

	mvt := model.ValidTime{}
	for _, item := range val {
		intval, err := strconv.Atoi(item)
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("err: invalid input"))
			return
		}
		mvt.Input = append(mvt.Input, intval)
	}

	valid, err := v.vtUseCase.ListAvailable(c, mvt)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("err: error processing input"))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":     "Input Processed Successfully",
		"input":      in.Input,
		"valid time": valid.List,
		"return":     valid.Count,
	})
}
