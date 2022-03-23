package model

const (
	PrefixVertex  string = "V"
	PrefixEdge    string = "e"
	GraphFileName string = "ggraph.png"
)

// Список названий radiogroup
type IdxCaptionRadioGp int

const (
	IdxCaptionChkMatrixAdj IdxCaptionRadioGp = 0
	IdxCaptionChkMatrixInd                   = 1
)

var listRadioGp []string = []string{"Матрица смежности", "Матрица инцидентноcти"}

// Список названий кнопок
type IdxCaptionBtn int

const (
	IdxCaptionBtnGraph   IdxCaptionBtn = 0
	IdxCaptionBtnDiGraph               = 1
)

var listButtons []string = []string{"Неориентированный граф", "Ориентированный граф"}

// интерфейс модуля
type Utilities interface {
	StrToGMatrix(data string) (*GMatrix, error)
}

// получение caption для radiogroup
func GetCaptionChk(value IdxCaptionRadioGp) string {
	return listRadioGp[int(value)]
}

func GetListRadioGp() *[]string {
	return &listRadioGp
}

// получение caption для кнопок
func GetCaptionButtons(value IdxCaptionBtn) string {
	return listButtons[int(value)]
}
