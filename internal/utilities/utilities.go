package utilities

import (
	"strconv"
	"strings"

	"github.com/Giovanny472/ggraph/model"
)

type utilities struct {
}

var util *utilities

func NewUtilities() model.Utilities {

	if util == nil {
		util = new(utilities)
	}
	return util
}

func (ut *utilities) StrToAdjMatrix(data string) (*model.AdjMatrix, error) {

	listMatrix := strings.Split(data, "\n")

	var amatrix model.AdjMatrix
	var arow []int

	for idx := 0; idx < len(listMatrix); idx++ {

		// строка
		listRow := strings.Split(listMatrix[idx], " ")

		// инициализация
		arow = nil

		for idxrow := 0; idxrow < len(listRow); idxrow++ {

			astrval := listRow[idxrow]
			astrval = strings.Trim(astrval, " ")
			if len(astrval) == 0 {
				continue
			}

			aintval, err := strconv.Atoi(astrval)
			if err != nil {
				return nil, err
			}
			arow = append(arow, aintval)

		}

		if len(arow) > 0 {
			amatrix = append(amatrix, arow)
		}

	}

	return &amatrix, nil
}
