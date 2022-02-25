package gui

import (
	"github.com/Giovanny472/ggraph/model"
)

type uigraph struct {
}

var uigrp *uigraph

func NewUIFactory() model.GUIFactory {

	if uigrp == nil {
		uigrp = new(uigraph)
		uigrp.init()
	}
	return uigrp
}

func (uigr *uigraph) init() {

}

func (uigr *uigraph) MakeMainWindow() model.UIFormMain {
	return NewFormMain()
}
