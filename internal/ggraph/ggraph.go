package ggraph

import (
	"github.com/Giovanny472/ggraph/model"
	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

type ggraph struct {

	// матрица инсидентности
	incmat *model.GMatrix

	// библиотека для рисования графов
	gviz      *graphviz.Graphviz
	gvizgraph *cgraph.Graph
}

var aggraph *ggraph

func NewGGraph() model.GGraph {
	if aggraph == nil {
		aggraph = new(ggraph)
	}
	aggraph.init()
	return aggraph
}

func (gr *ggraph) init() {

	// библиотека для рисования графов
	gr.gviz = graphviz.New()
}

func (gr *ggraph) SetIncidenceMatrix(mat *model.GMatrix) {

	gr.incmat = mat
}

func (gr *ggraph) Create() {

	if gr.incmat != nil {
		CreateGraphFromIncidenceMatrix(gr)
	}
}

func (gr *ggraph) Save(pathFile string) {
	// создание
	gr.gviz.RenderFilename(gr.gvizgraph, graphviz.PNG, string(pathFile))
}
