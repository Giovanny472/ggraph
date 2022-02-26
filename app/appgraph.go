package app

import (
	"github.com/Giovanny472/ggraph/gui"
	"github.com/Giovanny472/ggraph/internal/config"
	"github.com/Giovanny472/ggraph/internal/ggraph"
	"github.com/Giovanny472/ggraph/internal/manager"
	"github.com/Giovanny472/ggraph/internal/utilities"
	"github.com/Giovanny472/ggraph/model"
)

type appGraph struct {
	config  model.Config
	manag   model.Manager
	guifact model.GUIFactory
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
	app.manag = manager.NewManager(ggraph.NewGGraph(), utilities.NewUtilities())

	// слой UI
	app.guifact = gui.NewUIFactory()

}

func (app *appGraph) Config() {
	// пока не реализовано
}

func (app *appGraph) Start() {

	// главное окно
	fmMain := app.guifact.MakeMainWindow()

	// инициализация параметров
	fmMain.Init(app.manag)

	// отображение
	fmMain.Show()
}
