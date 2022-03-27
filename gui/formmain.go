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

	// назначение выбраного radiogroup(тип графа)
	typeGraphRadioGp model.TypeGraph
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

	//cписок radiogroup - тип матрицы
	chkMatrices := widget.NewRadioGroup(*model.GetListRadioGp(), fm.onListTypeMatrix)
	//cписок radiogroup - тип графа
	chkTypeGraphs := widget.NewRadioGroup(*model.GetListTypeGraph(), fm.onListTypeGraphs)

	// кнопки
	btnCreateGraph := widget.NewButton(model.IdxCaptionBtnCreateGraph, fm.onCreateGraph)

	// layout
	lytbtn := container.NewGridWithColumns(3, chkMatrices, chkTypeGraphs, btnCreateGraph)

	// layout
	baselyt := container.NewBorder(nil, lytbtn, nil, nil, lytCenter)

	fmMain.SetContent(baselyt)
}

func (fm *formMain) onListTypeMatrix(changed string) {

	if changed == model.GetCaptionGpTypeMatrix(model.IdxCaptionChkMatrixAdj) {
		fm.typeMatrixRadioGp = model.TypeMatrixAdj
	} else if changed == model.GetCaptionGpTypeMatrix(model.IdxCaptionChkMatrixInd) {
		fm.typeMatrixRadioGp = model.TypeMatrixInc
	} else {
		fm.typeMatrixRadioGp = -1
	}
}

func (fm *formMain) onListTypeGraphs(changed string) {

	if changed == model.GetCaptionGpTypeGraph(model.IdxCaptionChkDiGraph) {
		fm.typeGraphRadioGp = model.TypeDiGraph
	} else if changed == model.GetCaptionGpTypeGraph(model.IdxCaptionChkNotDiGraph) {
		fm.typeGraphRadioGp = model.TypeNotDiGraph
	} else {
		fm.typeGraphRadioGp = -1
	}
}

func (fm *formMain) onCreateGraph() {

	err := fm.mng.GenerateGraph(fm.txtMatrix.Text, fm.typeGraphRadioGp, fm.typeMatrixRadioGp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fm.showGraph(string(model.GraphFileName))
}

func (fm *formMain) showGraph(fileNameGraph string) {

	// отображение графа
	img := canvas.NewImageFromFile(fileNameGraph)
	img.FillMode = canvas.ImageFillOriginal
	fm.imgContainer.Objects = nil
	fm.imgContainer.Add(img)
	fm.imgContainer.Refresh()
}
