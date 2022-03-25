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
	} else if gp.matrix.TypeMatrix() == model.TypeMatrixInc {
		CreateGraphFromIncMatrix(gp, gp.tpgp)
	}
}

//******************************************
//  MATRIX
//******************************************
func (gp *ggraph) Matrix() model.IMatrix {
	return gp.matrix
}

//******************************************
//  FILES
//******************************************
func (gp *ggraph) Save(pathFile string) {
	// создание
	gp.gviz.RenderFilename(gp.gvizgraph, graphviz.PNG, string(pathFile))
}
