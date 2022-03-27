package model

const (
	PrefixVertex  string = "Вер-"
	PrefixEdge    string = "Реб-"
	GraphFileName string = "ggraph.png"
)

//******************************************
// Список названий radiogroup - матрица
//******************************************
type IdxCaptionRadioGpMatrix int

const (
	IdxCaptionChkMatrixAdj IdxCaptionRadioGpMatrix = 0
	IdxCaptionChkMatrixInd                         = 1
)

var listRadioGp []string = []string{"Матрица смежности", "Матрица инцидентноcти"}

//******************************************
// Список названий radiogroup - тип графа
//******************************************
type IdxCaptionRadioGpTypeGraph int

const (
	IdxCaptionChkNotDiGraph IdxCaptionRadioGpTypeGraph = 0
	IdxCaptionChkDiGraph                               = 1
)

var listChecksTypeGraph []string = []string{"Неориентированный граф", "Ориентированный граф"}

//******************************************
// Список названий radiogroup - тип графа
//******************************************
const IdxCaptionBtnCreateGraph string = "Граф"

//******************************************
// интерфейс модуля
//******************************************
type Utilities interface {
	StrToGMatrix(data string) (*GMatrix, error)
}

//****************************
// тип матрицы
//****************************

// получение caption для radiogroup матриц
func GetCaptionGpTypeMatrix(value IdxCaptionRadioGpMatrix) string {
	return listRadioGp[int(value)]
}

func GetListRadioGp() *[]string {
	return &listRadioGp
}

//****************************
// тип тип графа
//****************************

// получение caption для radiogroup матриц
func GetCaptionGpTypeGraph(value IdxCaptionRadioGpTypeGraph) string {
	return listChecksTypeGraph[int(value)]
}

// получение caption для типа граф
func GetListTypeGraph() *[]string {
	return &listChecksTypeGraph
}
