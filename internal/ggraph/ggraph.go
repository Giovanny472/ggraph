package ggraph

import (
	"github.com/Giovanny472/ggraph/model"
)

type ggraph struct {
}

var aggraph *ggraph

func NewGGraph() model.GGraph {
	if aggraph == nil {
		aggraph = new(ggraph)
	}
	return aggraph
}

func (gr *ggraph) SetSimpleAdjMatrix(adm *model.AdjMatrix) {

}

func (gr *ggraph) SetDirectedAdjMatrix(adm *model.AdjMatrix) {

}

func (gr *ggraph) Save() {
}
