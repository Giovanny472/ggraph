package model

type GUIFactory interface {
	// создание главного окна
	MakeMainWindow() UIFormMain
}
