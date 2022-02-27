package model

type AdjRow []int
type AdjMatrix []AdjRow

type GGraph interface {
	SetSimpleAdjMatrix(adm *AdjMatrix)
	SetDirectedAdjMatrix(adm *AdjMatrix)
	Create()
	Save()
}
