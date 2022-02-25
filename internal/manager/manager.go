package manager

import (
	"github.com/Giovanny472/ggraph/model"
)

type manager struct {
	graph model.GGraph
}

var amanager *manager

func NewManager(grph model.GGraph) model.Manager {
	if amanager == nil {
		amanager = &manager{grph}
	}
	return amanager
}

func (man *manager) Init() {

}

func (man *manager) Graph() model.GGraph {
	//	man.graph.SetSimpleAdjMatrix()
	return nil
}
