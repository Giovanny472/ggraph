package model

import "github.com/goccy/go-graphviz/cgraph"

type AdjRow []int
type AdjMatrix []AdjRow

type AdjEdge struct {
	NodeName  string
	NodeStart *cgraph.Node
	NodeEnd   *cgraph.Node
}

type AdjListEdge []AdjEdge

type GGraph interface {
	SetSimpleAdjMatrix(adm *AdjMatrix)
	SetIncidenceAdjMatrix(adm *AdjMatrix)
	Create()
	Save()
}
