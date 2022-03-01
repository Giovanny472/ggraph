package model

import "github.com/goccy/go-graphviz/cgraph"

type GRow []int
type GMatrix []GRow

type GEdge struct {
	NodeName  string
	NodeStart *cgraph.Node
	NodeEnd   *cgraph.Node
}

type ListEdge []GEdge

type GGraph interface {
	SetIncidenceMatrix(mat *GMatrix)
	Create()
	Save(pathFile string)
}
