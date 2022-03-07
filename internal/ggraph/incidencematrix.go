package ggraph

import (
	"strconv"

	"github.com/Giovanny472/ggraph/model"
)

func CreateGraphFromIncidenceMatrix(gp *ggraph) {

	// создание графа
	gp.gvizgraph, _ = gp.gviz.Graph()

	// количество вершин
	aCountNodes := len(*gp.incmat)

	// количество ребер
	aCountEdge := len((*gp.incmat)[0])

	// cоздание набора вершин
	var listNodes model.ListNodes
	for idx := 0; idx < aCountNodes; idx++ {

		anode, _ := gp.gvizgraph.CreateNode(model.PrefixVertex + strconv.Itoa(idx+1))
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
			aval := (*gp.incmat)[idxRow][idxCol]

			// проверка
			if aval == 0 {
				continue
			}

			if aval == -1 || aval == 1 {

				// название ребра
				aEdge.EdgeName = model.PrefixEdge + strconv.Itoa(idxCol+1)

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
			if aEdge.EdgeName != "" &&
				aEdge.NodeStart != nil &&
				aEdge.NodeEnd != nil {

				alistEdges = append(alistEdges, aEdge)
			}
		}
		aEdge = nil

	}

	for _, anode := range alistEdges {
		aed, _ := gp.gvizgraph.CreateEdge(anode.EdgeName, anode.NodeStart, anode.NodeEnd)
		aed.SetLabel(anode.EdgeName)
	}

}
