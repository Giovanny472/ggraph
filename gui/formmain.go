package gui

import (
	"fmt"
	"strings"

	"github.com/Giovanny472/ggraph/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type formMain struct {

	// manager
	mng model.Manager

	// widgets
	txtMatrix *widget.Entry
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
	fmMain.Resize(fyne.NewSize(1020, 720))

	// lytTop
	//lytTop := container.NewGridWithColumns(1, widget.NewMultiLineEntry())
	fm.txtMatrix = widget.NewMultiLineEntry()
	//fm.txtMatrix.FocusGained()
	//!!интересно --> lytTop := container.NewVBox(txtMatrix)

	// lytCenter
	txtGraph := widget.NewMultiLineEntry()

	lytCenter := container.NewGridWithRows(2, fm.txtMatrix, txtGraph)

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

	//val := fm.txtMatrix.Text
	//enter := []byte("10")
	val := strings.Split(fm.txtMatrix.Text, " ")

	fmt.Println(val)

	for _, element := range val {

		abyte := []byte(element)
		if abyte == "10" {
			fmt.Println("enter")
			continue
		}

		fmt.Print(abyte, "--")
		//fmt.Print(len(abyte), "<>")
	}
	fmt.Println(" ")
	//var am model.AdjMatrix
	//am = append(am, )
	//fm.mng.Graph().SetSimpleAdjMatrix()
	//fm.mng.Graph().
}
