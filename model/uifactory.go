package model

import "github.com/Giovanny472/ggraph/gui"

type UIFactory interface {
	// создание главного окна
	MakeMainWindow() gui.GUI
}
