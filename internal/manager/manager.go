package manager

import (
	"github.com/Giovanny472/ggraph/model"
)

type manager struct {
	graph model.GGraph
	util  model.Utilities
}

var amanager *manager

func NewManager(grph model.GGraph, ut model.Utilities) model.Manager {
	if amanager == nil {
		amanager = &manager{grph, ut}
	}
	return amanager
}

func (man *manager) Init() {
}

func (man *manager) Graph() model.GGraph {
	return man.graph
}

func (man *manager) Utilities() model.Utilities {
	return man.util
}
