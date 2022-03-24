package ggraph

import (
	"github.com/Giovanny472/ggraph/model"
	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

type ggraph struct {

	// библиотека для рисования графов
	gviz      *graphviz.Graphviz
	gvizgraph *cgraph.Graph

	// тип графа
	tpgp model.TypeGraph

	// для назначени матрицы
	matrix model.IMatrix

	// матрица инцидентности
	//matIncidence *model.GMatrix

	// матрица смежности
	//matAdj *model.GMatrix

}

/*
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
*/
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

func (gp *ggraph) init() {

	// матрица (смежности или инцидентности)
	gp.matrix = NewMatrix()

	// cоздание экземпляра для рисования графов
	gp.gviz = graphviz.New()
}

// тип графов
func (gp *ggraph) SetType(tpg model.TypeGraph) {
	gp.tpgp = tpg
}

func (gp *ggraph) Create() {

	if gp.matrix.TypeMatrix() == model.TypeMatrixAdj {

		CreateGraphFromAdjMatrix(gp, gp.tpgp)

	} //else if gp.matrix.TypeMatrix() == model.TypeMatrixInc {
	//		CreateGraphFromIncidenceMatrix(gr)
	//}

}

//******************************************
//  MATRIX
//******************************************
func (gp *ggraph) Matrix() model.IMatrix {
	return gp.matrix
}

// тип матрицы
//func (gp *ggraph) SetTypeMatrix(tpm model.TypeMatrix) {
//	gp.matrix.SetTypeMatrix(tpm)
//}

// назначение матрицы
//func (gp *ggraph) SetMatrix(mat *model.GMatrix) {
//	gp.matrix.SetMatrix(mat)
//}

//******************************************
//  FILES
//******************************************
func (gp *ggraph) Save(pathFile string) {
	// создание
	gp.gviz.RenderFilename(gp.gvizgraph, graphviz.PNG, string(pathFile))
	//gr.directed.gviz.RenderFilename(gr.directed.gvizgraph, graphviz.PNG, string(pathFile))
}

/*
func (gr *ggraph) Directed() model.Matrix {
	return gr.directed
}

func (gr *ggraph) NoDirected() model.Matrix {
	return gr.notDirected
}

func (gr *ggraph) Create(tpMatrix model.TypeMatrix, tpGraph model.TypeGraph) {

	if tpMatrix == model.TypeMatrixAdj {
		CreateGraphFromAdjMatrix(gr, tpGraph)
	} else if tpMatrix == model.TypeMatrixInc {
		CreateGraphFromIncidenceMatrix(gr)
	}
}
*/

//func (gr *ggraph) Save(pathFile string) {
//	// создание
//	gr.directed.gviz.RenderFilename(gr.directed.gvizgraph, graphviz.PNG, string(pathFile))
//}

/*
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
*/
