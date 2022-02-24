package app

import (
	"github.com/Giovanny472/ggraph/model"

	"github.com/Giovanny472/ggraph/internal/config"
	"github.com/Giovanny472/ggraph/internal/ggraph"
	"github.com/Giovanny472/ggraph/internal/manager"
	"github.com/Giovanny472/ggraph/internal/uigraph"
)

type appGraph struct {
	config     model.Config
	manag      model.Manager
	guifactory model.UIFactory
}

var apg *appGraph

func NewAppGraph() Application {

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
	app.manag = manager.NewManager(ggraph.NewGGraph())

	// слой UI
	app.guifactory = uigraph.NewUIFactory()

}

func (app *appGraph) Config() {

}

func (app *appGraph) Start() {

	// главное окно
	fmMain := app.guifactory.MakeMainWindow()

	// инициализация параметров
	//fmMain.Init(app.manag)

	// отображение
	fmMain.Show()
}
