package app

import (
	"github.com/Giovanny472/ggraph/model"

	"github.com/Giovanny472/ggraph/internal/config"
	"github.com/Giovanny472/ggraph/internal/ggraph"
	"github.com/Giovanny472/ggraph/internal/manager"
)

type appGraph struct {
	config model.Config
	manag  model.Manager
	ggraph model.GGraph
}

var apg *appGraph

func NewAppGraph() model.Application {

	if apg == nil {
		apg = new(appGraph)
	}
	apg.init()
	return apg
}

func (app *appGraph) init() {

	// слой данных
	app.config = config.NewConfig()

	// слой логики
	app.manag = manager.NewManager()
	app.ggraph = ggraph.NewGGraph()

	// слой ui
}

func (app *appGraph) Config() {

}

func (app *appGraph) Start() {

}
