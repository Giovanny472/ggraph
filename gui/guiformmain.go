package gui

import "github.com/Giovanny472/ggraph/model"

type AdjMatrix [][]int

type UIFormMain interface {
	model.UIBase
	Init(manag model.Manager)
	SetSimpleAdjMatrix(adm AdjMatrix)
}
