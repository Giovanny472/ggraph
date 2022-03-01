package ggraph

import (
	"strconv"

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

	// создание графа
	gr.gvizgraph, _ = gr.gviz.Graph()
	/**
	defer func() {
		if err := gr.gvizgraph.Close(); err != nil {
			log.Fatal(err)
		}
		gr.gvizgraph = nil
	}()
	*/

	// количество вершин
	aCountNodes := len(*gr.incmat)

	// количество ребер
	aCountEdge := len((*gr.incmat)[0])

	// cоздание набора вершин
	var listNodes []*cgraph.Node
	for idx := 0; idx < aCountNodes; idx++ {

		anode, _ := gr.gvizgraph.CreateNode(model.PrefixVertex + strconv.Itoa(idx+1))
		listNodes = append(listNodes, anode)
	}

	// определение набора ребер
	var alistEdges model.ListEdge

	for idxCol := 0; idxCol < aCountEdge; idxCol++ {

		// ребер
		var aEdge *model.GEdge

		// cоздание
		aEdge = new(model.GEdge)

		for idxRow := 0; idxRow < aCountNodes; idxRow++ {

			// получаем значение
			aval := (*gr.incmat)[idxRow][idxCol]

			// проверка
			if aval == 0 {
				continue
			}

			if aval == -1 || aval == 1 {

				// название ребра
				aEdge.NodeName = model.PrefixEdge + strconv.Itoa(idxCol+1)

				// start
				if aval == -1 {
					aEdge.NodeStart = listNodes[idxRow]
				}

				// end
				if aval == 1 {
					aEdge.NodeEnd = listNodes[idxRow]
				}
			}
		}

		if aEdge != nil {
			if aEdge.NodeName != "" &&
				aEdge.NodeStart != nil &&
				aEdge.NodeEnd != nil {

				alistEdges = append(alistEdges, *aEdge)
			}
		}
		aEdge = nil

	}

	for _, anode := range alistEdges {
		aed, _ := gr.gvizgraph.CreateEdge(anode.NodeName, anode.NodeStart, anode.NodeEnd)
		aed.SetLabel(anode.NodeName)
	}
}

func (gr *ggraph) Save(pathFile string) {
	// создание
	gr.gviz.RenderFilename(gr.gvizgraph, graphviz.PNG, string(pathFile))
}
