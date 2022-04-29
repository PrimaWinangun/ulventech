package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"ulventech/model"
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
