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

	// lytTop
	//lytTop := container.NewGridWithColumns(1, widget.NewMultiLineEntry())
	fm.txtMatrix = widget.NewMultiLineEntry()
	//fm.txtMatrix.FocusGained()
	//!!интересно --> lytTop := container.NewVBox(txtMatrix)

	// lytCenter
	fm.imgContainer = container.NewCenter()

	//txtGraph := widget.NewMultiLineEntry()
	lytCenter := container.NewGridWithColumns(2, fm.txtMatrix, fm.imgContainer)

	//lytColBottom
	btnExit := widget.NewButton("выход", fm.onClose)
	btnGraph := widget.NewButton("граф", fm.onGraph)
	lytbtn := container.NewGridWithColumns(2, btnExit, btnGraph)

	// layout
	baselyt := container.NewBorder(nil, lytbtn, nil, nil, lytCenter)

	fmMain.SetContent(baselyt)
}

func (fm *formMain) onClose() {
	fyne.CurrentApp().Quit()
}

func (fm *formMain) onGraph() {

	// получение матрицы
	adjMatrix, err := fm.mng.Utilities().StrToGMatrix(fm.txtMatrix.Text)
	if err != nil {
		log.Println("Error Ongraph: " + err.Error())
	}

	// настройка матрица
	fm.mng.Graph().Directed().SetIncidence(adjMatrix)

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
