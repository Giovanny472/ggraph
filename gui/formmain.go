package gui

import (
	"fmt"

	"github.com/Giovanny472/ggraph/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type formMain struct {

	// матрица
	adjmatrix *model.AdjMatrix

	// manager
	mng model.Manager
}

var fmmain *formMain

func NewFormMain() model.UIFormMain {

	if fmmain == nil {
		fmmain = new(formMain)
	}
	return fmmain
}

func (fm *formMain) Init(manag model.Manager) {
	fm.mng = manag
}

func (fm *formMain) SetSimpleAdjMatrix(adm *model.AdjMatrix) {
	fm.adjmatrix = adm
}

func (fm *formMain) Show() {

	// запуск UI
	appUI := app.New()
	formMain := appUI.NewWindow("Граф")

	// построим UI окно
	fm.buildform(formMain)

	// отображение окна
	formMain.Show()

	appUI.Run()
	fmt.Println("show form")
}

func (fm *formMain) buildform(fmMain fyne.Window) {

	// положение
	fmMain.CenterOnScreen()

	// размер
	fmMain.Resize(fyne.NewSize(1020, 720))

	// lytCenter
	lytTop := container.NewGridWithColumns(1, widget.NewMultiLineEntry())

	// lytCenter
	lytCenter := container.NewGridWithColumns(1, widget.NewMultiLineEntry())

	//lytColBottom
	lyBottom := container.NewGridWithColumns(2, widget.NewButton("выход", nil), widget.NewButton("граф", nil))

	// layout
	baselyt := container.NewGridWithRows(3,
		lytTop,
		lytCenter,
		lyBottom)

	fmMain.SetContent(baselyt)
}
