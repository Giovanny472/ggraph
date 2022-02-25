package model

type AdjMatrix [][]int

type UIFormMain interface {
	UIBase
	Init(manag Manager)
	SetSimpleAdjMatrix(adm AdjMatrix)
}
