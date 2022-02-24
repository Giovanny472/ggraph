package uigraph

import (
	"github.com/Giovanny472/ggraph/gui"
	"github.com/Giovanny472/ggraph/model"
)

type uigraph struct {
}

var uigrp *uigraph

func NewUIFactory() model.UIFactory {

	if uigrp == nil {
		uigrp = new(uigraph)
		uigrp.init()
	}
	return uigrp
}

func (uigr *uigraph) init() {

}

func (uigr *uigraph) MakeMainWindow() gui.GUI {
	return gui.NewFormMain()
}
