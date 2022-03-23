package ggraph

import (
	"strconv"

	"github.com/Giovanny472/ggraph/model"
)

func CreateGraphFromAdjMatrix(gp *ggraph) {

	// создание графа
	gp.directed.gvizgraph, _ = gp.directed.gviz.Graph()

	// количество вершин
	aCountNodes := len(*gp.directed.matAdj)

	// количество ребер
	aCountEdge := len(*gp.directed.matAdj)

	//if flag == model.TypeMatrixInc {}

	var alblCountEdge int = 0

	// cоздание набора вершин
	var listNodes model.ListNodes
	for idx := 0; idx < aCountNodes; idx++ {

		anode, _ := gp.directed.gvizgraph.CreateNode(model.PrefixVertex + strconv.Itoa(idx+1))
		//anode.SetShape(cgraph.StarShape)
		listNodes = append(listNodes, anode)
	}

	// определение набора ребер
	var alistEdges model.ListEdge

	for idxRow := 0; idxRow < aCountEdge; idxRow++ {

		for idxCol := 0; idxCol < aCountNodes; idxCol++ {

			// получаем значение
			aval := (*gp.directed.matAdj)[idxRow][idxCol]

			// проверка
			if aval == 0 {
				continue
			}

			// ребер
			var aEdge *model.GEdge = nil

			if aval == 1 {

				// cоздание
				aEdge = new(model.GEdge)

				// название ребра
				alblCountEdge++
				aEdge.EdgeName = model.PrefixEdge + strconv.Itoa(alblCountEdge)

				// start
				aEdge.NodeStart = listNodes[idxRow]

				// end
				aEdge.NodeEnd = listNodes[idxCol]

			}

			if aEdge != nil {
				if aEdge.EdgeName != "" &&
					aEdge.NodeStart != nil &&
					aEdge.NodeEnd != nil {

					alistEdges = append(alistEdges, aEdge)
				}
			}

		}
	}

	for _, anode := range alistEdges {
		aed, _ := gp.directed.gvizgraph.CreateEdge(anode.EdgeName, anode.NodeStart, anode.NodeEnd)
		//aed.SetArrowHead(cgraph.NoneArrow)
		//aed.SetColor("orange")
		//aed.SetStyle(cgraph.BoldEdgeStyle)
		aed.SetLabel(anode.EdgeName)
	}

}
