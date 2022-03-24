package model

import "github.com/goccy/go-graphviz/cgraph"

type GRow []int
type GMatrix []GRow

type GEdge struct {
	EdgeName  string
	NodeStart *cgraph.Node
	NodeEnd   *cgraph.Node
}

type ListNodes []*cgraph.Node
type ListEdge []*GEdge

//******************************************
// флаг тип матрицы
//******************************************
type TypeMatrix int

const (
	TypeMatrixAdj TypeMatrix = 0
	TypeMatrixInc            = 1
)

//******************************************
// определение операцией с матрицей
//******************************************

// представление графов
type IMatrix interface {

	// тип матрицы
	SetTypeMatrix(TypeMatrix)
	TypeMatrix() TypeMatrix

	// назначение матрицы
	SetMatrix(mat *GMatrix)
	Matrix() *GMatrix
}

//******************************************
// файл
//******************************************

// хранение графа
type File interface {
	Save(pathFile string)
}

//******************************************
// флаг - тип графа
//******************************************
type TypeGraph int

const (
	TypeDiGraph    TypeGraph = 0
	TypeNotDiGraph           = 1
)

//******************************************
// определение графа
//******************************************

type GGraph interface {

	// работа с файлами
	File

	// тип графов
	SetType(TypeGraph)

	// Матрица
	Matrix() IMatrix

	// cоздание графа
	Create()
}
