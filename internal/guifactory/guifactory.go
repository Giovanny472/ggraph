package guifactory

import (
	"github.com/Giovanny472/ggraph/gui"
)

type uigraph struct {
}

var uigrp *uigraph

func NewUIFactory() gui.GUIFactory {

	if uigrp == nil {
		uigrp = new(uigraph)
		uigrp.init()
	}
	return uigrp
}

func (uigr *uigraph) init() {

}

func (uigr *uigraph) MakeMainWindow() gui.UIFormMain {
	return gui.NewFormMain()
}
