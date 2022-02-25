package model

type AdjMatrix [][]int

type GGraph interface {
	SetSimpleAdjMatrix(adm *AdjMatrix)
	SetDirectedAdjMatrix(adm *AdjMatrix)
	Save()
}
