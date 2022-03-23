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

// представление графов
type Matrix interface {

	// назначение матрицы смежности
	SetAdjacency(mat *GMatrix)

	// назначение матрицы инцидентности
	SetIncidence(mat *GMatrix)
}

// хранение графа
type File interface {
	Save(pathFile string)
}

type GGraph interface {

	// тип графов
	// Ориентированный граф
	Directed() Matrix
	// Неориентированный граф
	NoDirected() Matrix

	// работа с файлами
	File

	// cоздание графа
	Create(tpMatrix TypeMatrix)
}
