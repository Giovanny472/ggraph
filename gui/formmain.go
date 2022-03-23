package gui

import (
	"log"

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

	// полученное caption из event radiogroup
	strEventRadioGp string
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
	fm.txtMatrix.SetText("")

	// lytCenter
	fm.imgContainer = container.NewCenter()

	// layout
	lytCenter := container.NewGridWithColumns(2, fm.txtMatrix, fm.imgContainer)

	//cписок radiogroup
	chkMatrices := widget.NewRadioGroup(*model.GetListRadioGp(), fm.onGrpBox)

	// кнопки
	btnGraph := widget.NewButton(model.GetCaptionButtons(model.IdxCaptionBtnGraph), fm.onGraph)
	btnDiGraph := widget.NewButton(model.GetCaptionButtons(model.IdxCaptionBtnDiGraph), fm.onDiGraph)

	// layout
	lytbuttons := container.NewGridWithRows(2, btnGraph, btnDiGraph)
	lytbtn := container.NewGridWithColumns(2, chkMatrices, lytbuttons)

	// layout
	baselyt := container.NewBorder(nil, lytbtn, nil, nil, lytCenter)

	fmMain.SetContent(baselyt)
}

func (fm *formMain) onClose() {
	fyne.CurrentApp().Quit()
}

func (fm *formMain) onGrpBox(changed string) {
	fm.strEventRadioGp = changed
}

func (fm *formMain) onDiGraph() {

	// получение матрицы
	aMatrix, err := fm.mng.Utilities().StrToGMatrix(fm.txtMatrix.Text)
	if err != nil {
		log.Println("Error onDiGraph: " + err.Error())
		return
	}
	// проверка выбранной матрицы
	if len(fm.strEventRadioGp) == 0 {
		log.Println("нет выбранной матрицы ")
		return
	}

	// настройка матрицы
	if fm.strEventRadioGp == model.GetCaptionChk(model.IdxCaptionChkMatrixAdj) {
		fm.mng.Graph().Directed().SetAdjacency(aMatrix)
	} else if fm.strEventRadioGp == model.GetCaptionChk(model.IdxCaptionChkMatrixInd) {
		fm.mng.Graph().Directed().SetIncidence(aMatrix)
	} else {
		log.Println("не найдена матрица")
		return
	}

	// создание графа
	fm.mng.Graph().Create()

	// cохранение графа
	fm.mng.Graph().Save(model.GraphFileName)

	// отображение графа
	img := canvas.NewImageFromFile(string(model.GraphFileName))
	img.FillMode = canvas.ImageFillOriginal
	fm.imgContainer.Objects = nil
	fm.imgContainer.Add(img)
	fm.imgContainer.Refresh()

}

func (fm *formMain) onGraph() {

	// получение матрицы
	aMatrix, err := fm.mng.Utilities().StrToGMatrix(fm.txtMatrix.Text)
	if err != nil {
		log.Println("Error onGraph: " + err.Error())
		return
	}
	// проверка выбранной матрицы
	if len(fm.strEventRadioGp) == 0 {
		log.Println("нет выбранной матрицы ")
		return
	}

	// настройка матрицы
	if fm.strEventRadioGp == model.GetCaptionChk(model.IdxCaptionChkMatrixAdj) {
		fm.mng.Graph().NoDirected().SetAdjacency(aMatrix)
	} else if fm.strEventRadioGp == model.GetCaptionChk(model.IdxCaptionChkMatrixInd) {
		fm.mng.Graph().NoDirected().SetIncidence(aMatrix)
	} else {
		log.Println("не найдена матрица")
		return
	}

	// создание графа
	fm.mng.Graph().Create()

	// cохранение графа
	fm.mng.Graph().Save(model.GraphFileName)

}
