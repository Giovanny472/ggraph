package model

type GGraph interface {
	SetSimpleAdjMatrix()
	SetDirectedAdjMatrix()
	Show()
}
