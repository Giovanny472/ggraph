package ggraph

import (
	"github.com/Giovanny472/ggraph/model"
)

type ggraph struct {
}

var aggraph *ggraph

func NewConfig() model.GGraph {
	if aggraph == nil {
		aggraph = new(ggraph)
	}
	return aggraph
}

func (gr *ggraph) SetSimpleAdjMatrix() {

}

func (gr *ggraph) SetDirectedAdjMatrix() {

}

func (gr *ggraph) Show() {
}
