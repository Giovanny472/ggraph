package ggraph

import (
	"github.com/Giovanny472/ggraph/model"
	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

type dataGGraph struct {

	// матрица инцидентности
	matIncidence *model.GMatrix

	// матрица смежности
	matAdj *model.GMatrix

	// библиотека для рисования графов
	gviz      *graphviz.Graphviz
	gvizgraph *cgraph.Graph
}

// Ориентированный граф
type dataDirGGraph struct {
	*dataGGraph
}

// Неориентированный граф
type dataNotDirGGraph struct {
	*dataGGraph
}

type ggraph struct {

	// Неориентированный граф
	directed *dataDirGGraph

	// Oриентированный граф
	notDirected *dataNotDirGGraph
}

var (
	aggraph *ggraph
)

//******************************************
//  GGRAPH
//******************************************

func NewGGraph() model.GGraph {
	if aggraph == nil {
		aggraph = new(ggraph)
	}
	aggraph.init()
	return aggraph
}

func (gr *ggraph) init() {

	// инициализация
	// библиотеки для рисования графов:

	// ориентированный граф
	gr.directed = newDataDirGGraph()

	// неориентированный граф
	gr.notDirected = newDataNotDirGGraph()

}

func (gr *ggraph) Directed() model.Matrix {
	return gr.directed
}

func (gr *ggraph) NoDirected() model.Matrix {
	return gr.notDirected
}

func (gr *ggraph) Create(tpMatrix model.TypeMatrix) {

	if tpMatrix == model.TypeMatrixAdj {
		CreateGraphFromAdjMatrix(gr)
	} else if tpMatrix == model.TypeMatrixInc {
		CreateGraphFromIncidenceMatrix(gr)
	}
}

func (gr *ggraph) Save(pathFile string) {
	// создание
	gr.directed.gviz.RenderFilename(gr.directed.gvizgraph, graphviz.PNG, string(pathFile))
}

//******************************************
//  DATADIRGGRAPH
//******************************************
func newDataDirGGraph() *dataDirGGraph {

	return &dataDirGGraph{
		&dataGGraph{gviz: graphviz.New()},
	}
}

// назначение матрицы смежности
func (dirg *dataDirGGraph) SetAdjacency(mat *model.GMatrix) {
	dirg.dataGGraph.matAdj = mat
}

// назначение матрицы инцидентности
func (dirg *dataDirGGraph) SetIncidence(mat *model.GMatrix) {
	dirg.dataGGraph.matIncidence = mat
}

//******************************************
//  DATANOTDIRGGRAPH
//******************************************
func newDataNotDirGGraph() *dataNotDirGGraph {

	return &dataNotDirGGraph{
		&dataGGraph{gviz: graphviz.New()},
	}
}

// назначение матрицы смежности
func (dirng *dataNotDirGGraph) SetAdjacency(mat *model.GMatrix) {
	dirng.dataGGraph.matAdj = mat
}

// назначение матрицы инцидентности
func (dirng *dataNotDirGGraph) SetIncidence(mat *model.GMatrix) {
	dirng.dataGGraph.matIncidence = mat
}
