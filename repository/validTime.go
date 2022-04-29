package repository

import (
	"fmt"
	"ulventech/model"
)

type vtRepository struct {
}

func NewVtRepository() model.ValidTimeRepository {
	return &vtRepository{}
}

func found(needle string, hay []string) bool {
	for _, v := range hay {
		if needle == v {
			return true
		}
	}
	return false
}

func (v vtRepository) Permutation(t model.ValidTime) (str []string, err error) {
	var rc func([]int, int)
	rc = func(x []int, y int) {
		if y == len(x) {
			strTime := fmt.Sprintf("%v%v:%v%v", x[0], x[1], x[2], x[3])
			if f := found(strTime, str); !f {
				str = append(str, strTime)
			}
		} else {
			for z := 0; z < len(t.Input); z++ {
				x[y], x[z] = x[z], x[y]
				rc(x, y+1)
				x[y], x[z] = x[z], x[y]
			}
		}
	}

	rc(t.Input, 0)
	return
}
