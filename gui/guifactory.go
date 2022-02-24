package gui

type GUIFactory interface {
	// создание главного окна
	MakeMainWindow() UIFormMain
}
