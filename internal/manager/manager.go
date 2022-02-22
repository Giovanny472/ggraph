package manager

import (
	"github.com/Giovanny472/ggraph/model"
)

type manager struct {
}

var amanager *manager

func NewConfig() model.Manager {
	if amanager == nil {
		amanager = new(manager)
	}
	return amanager
}

func (man *manager) Init() {

}

func (man *manager) Start() {
}
