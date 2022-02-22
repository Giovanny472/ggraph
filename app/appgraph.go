package app

import "github.com/Giovanny472/ggraph/model"

type appGraph struct {
}

var apg *appGraph

func NewAppGraph() model.Application {

	if apg == nil {
		apg = new(appGraph)
	}
	return apg
}

func (app *appGraph) Config() {

}

func (app *appGraph) Start() {

}
