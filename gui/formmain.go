package gui

import (
	"fmt"

	"github.com/Giovanny472/ggraph/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type formMain struct {

	// manager
	mng model.Manager

	// widgets
	txtMatrix *widget.Entry

	// контайнер
	imgContainer *fyne.Container

	// назначение выбраного radiogroup(матрицу)
	typeMatrixRadioGp model.TypeMatrix
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

func (fm *formMain) Show() {

	// запуск UI
	appUI := app.New()
	formMain := appUI.NewWindow("Граф")

	// построим UI окно
	fm.buildform(formMain)

	// отображение окна
	formMain.Show()

	appUI.Run()
}

func (fm *formMain) buildform(fmMain fyne.Window) {

	// положение
	fmMain.CenterOnScreen()

	// темная тема
	fyne.CurrentApp().Settings().SetTheme(theme.DarkTheme())

	// размер
	fmMain.Resize(fyne.NewSize(720, 480))

	// для ввохода матрицы
	fm.txtMatrix = widget.NewMultiLineEntry()
	fm.txtMatrix.SetPlaceHolder("Вводите матрицу здесь")
	fm.txtMatrix.SetText("")

	// lytCenter
	fm.imgContainer = container.NewCenter()

	// layout
	lytCenter := container.NewGridWithColumns(2, fm.txtMatrix, fm.imgContainer)

	//cписок radiogroup
	chkMatrices := widget.NewRadioGroup(*model.GetListRadioGp(), fm.onGrpBox)
	chkTypeGraphs := widget.NewRadioGroup(*model.GetListTypeGraph(), fm.onListTypeGraphs)
	//chkDiGraph := widget.NewRadioGroup(*model.GetCaptionButtons(model.IdxCaptionBtnDiGraph), fm.onDiGraph)

	// кнопки
	btnCreateGraph := widget.NewButton(model.IdxCaptionBtnCreateGraph, fm.onCreateGraph)

	// layout
	//lytbuttons := container.NewGridWithRows(2, chkTypeGraphs, btnCreateGraph)
	lytbtn := container.NewGridWithColumns(3, chkMatrices, chkTypeGraphs, btnCreateGraph)

	// layout
	baselyt := container.NewBorder(nil, lytbtn, nil, nil, lytCenter)

	fmMain.SetContent(baselyt)
}

func (fm *formMain) onClose() {
	fyne.CurrentApp().Quit()
}

func (fm *formMain) onGrpBox(changed string) {

	if changed == model.GetCaptionChk(model.IdxCaptionChkMatrixAdj) {
		fm.typeMatrixRadioGp = model.TypeMatrixAdj
	} else if changed == model.GetCaptionChk(model.IdxCaptionChkMatrixInd) {
		fm.typeMatrixRadioGp = model.TypeMatrixInc
	} else {
		fm.typeMatrixRadioGp = -1
	}
}

func (fm *formMain) onListTypeGraphs(changed string) {

}

func (fm *formMain) onDiGraph(changed string) {

	err := fm.mng.GenerateGraph(fm.txtMatrix.Text, model.TypeDiGraph, fm.typeMatrixRadioGp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fm.showGraph(string(model.GraphFileName))
}

func (fm *formMain) onNotDiGraph(changed string) {

	err := fm.mng.GenerateGraph(fm.txtMatrix.Text, model.TypeNotDiGraph, fm.typeMatrixRadioGp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fm.showGraph(string(model.GraphFileName))
}

func (fm *formMain) onCreateGraph() {

}

func (fm *formMain) showGraph(fileNameGraph string) {

	// отображение графа
	img := canvas.NewImageFromFile(fileNameGraph)
	img.FillMode = canvas.ImageFillOriginal
	fm.imgContainer.Objects = nil
	fm.imgContainer.Add(img)
	fm.imgContainer.Refresh()
}
