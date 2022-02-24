package gui

import (
	"fmt"
)

type formMain struct {
}

var fmmain *formMain

func NewFormMain() GUI {

	if fmmain == nil {
		fmmain = new(formMain)
	}
	return fmmain
}

//func (fm *formMain) Init(manag model.Manager) {

//}

func (fm *formMain) Show() {

	// построим окно
	fm.buildform()

	fmt.Println("show form")
}

func (fmain *formMain) buildform() {

}
