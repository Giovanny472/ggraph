package utilities

import (
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

func (ut *utilities) StrToAdjMatrix(data string) *model.AdjMatrix {

	/*
		listMatrix := strings.Split(fm.txtMatrix.Text, "\n")
		fmt.Println("listMatrix-->", listMatrix)

		var amatrix [][]string
		var row []string

		for idx := 0; idx < len(listMatrix); idx++ {

			listRow := strings.Split(listMatrix[idx], " ")

			row = nil

			for idxrow := 0; idxrow < len(listRow); idxrow++ {

				avalue := listRow[idxrow]
				avalue = strings.Trim(avalue, " ")
				if len(avalue) == 0 {
					continue
				}

				row = append(row, avalue)

			}

			if len(row) > 0 {
				amatrix = append(amatrix, row)
			}

		}

		fmt.Println("final:")
		fmt.Println(amatrix)
	*/

	return nil
}
