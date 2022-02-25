package gui

import (
	"fmt"

	"github.com/Giovanny472/ggraph/model"
)

type formMain struct {
}

var fmmain *formMain

func NewFormMain() model.UIFormMain {

	if fmmain == nil {
		fmmain = new(formMain)
	}
	return fmmain
}

func (fm *formMain) Init(manag model.Manager) {

}

func (fm *formMain) SetSimpleAdjMatrix(adm model.AdjMatrix) {

}

func (fm *formMain) Show() {

	// построим окно
	fm.buildform()

	fmt.Println("show form")
}

func (fmain *formMain) buildform() {

}
