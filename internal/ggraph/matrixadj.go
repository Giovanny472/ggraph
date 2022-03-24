package ggraph

import (
	"strconv"

	"github.com/Giovanny472/ggraph/model"
)

func CreateGraphFromAdjMatrix(gp *ggraph, typegp model.TypeGraph) {

	switch typegp {
	case model.TypeDiGraph:
		CreateDiGraphFromAdjMatrix(gp)

	case model.TypeNotDiGraph:
		CreateNotDiGraphFromAdjMatrix(gp)
	}
}

// ориентированный граф
func CreateDiGraphFromAdjMatrix(gp *ggraph) {

	// создание графа
	gp.gvizgraph, _ = gp.gviz.Graph()

	// количество вершин
	aCountNodes := len(*gp.matrix.Matrix())

	// количество ребер
	aCountEdge := len(*gp.matrix.Matrix())

	// для определения названия вершин
	var alblCountEdge int = 0

	// cоздание набора вершин
	var listNodes model.ListNodes
	for idx := 0; idx < aCountNodes; idx++ {

		anode, _ := gp.gvizgraph.CreateNode(model.PrefixVertex + strconv.Itoa(idx+1))
		//anode.SetShape(cgraph.StarShape)
		listNodes = append(listNodes, anode)
	}

	// определение набора ребер
	var alistEdges model.ListEdge

	for idxRow := 0; idxRow < aCountEdge; idxRow++ {

		for idxCol := 0; idxCol < aCountNodes; idxCol++ {

			// получаем значение
			aval := (*gp.matrix.Matrix())[idxRow][idxCol]

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
		aed, _ := gp.gvizgraph.CreateEdge(anode.EdgeName, anode.NodeStart, anode.NodeEnd)
		//aed.SetArrowHead(cgraph.NoneArrow)
		//aed.SetColor("orange")
		//aed.SetStyle(cgraph.BoldEdgeStyle)
		aed.SetLabel(anode.EdgeName)
	}

}

// неориентированный граф
func CreateNotDiGraphFromAdjMatrix(gp *ggraph) {
	/*
		// создание графа
		gp.notDirected.gvizgraph, _ = gp.notDirected.gviz.Graph()

		// количество вершин
		aCountNodes := len(*gp.notDirected.matAdj)

		// количество ребер
		aCountEdge := len(*gp.notDirected.matAdj)

		// для определения названия вершин
		var alblCountEdge int = 0

		// cоздание набора вершин
		var listNodes model.ListNodes

		for idx := 0; idx < aCountNodes; idx++ {

			anode, _ := gp.notDirected.gvizgraph.CreateNode(model.PrefixVertex + strconv.Itoa(idx+1))
			//anode.SetShape(cgraph.StarShape)
			listNodes = append(listNodes, anode)
		}

		fmt.Println("gp.notDirected-->", gp.notDirected)
		fmt.Println("listNodes-->", listNodes)
		fmt.Println("aCountNodes-->", aCountNodes)
		fmt.Println("aCountEdge-->", aCountEdge)

		// определение набора ребер
		var alistEdges model.ListEdge

		for idxRow := 0; idxRow < aCountEdge; idxRow++ {

			for idxCol := 0; idxCol < aCountNodes; idxCol++ {

				// получаем значение
				aval := (*gp.notDirected.matAdj)[idxRow][idxCol]

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
			aed, _ := gp.notDirected.gvizgraph.CreateEdge(anode.EdgeName, anode.NodeStart, anode.NodeEnd)
			aed.SetArrowHead(cgraph.NoneArrow)
			//aed.SetColor("orange")
			//aed.SetStyle(cgraph.BoldEdgeStyle)
			aed.SetLabel(anode.EdgeName)
		}

		fmt.Println("end...CreateNotDiGraphFromAdjMatrix..end")
	*/
}
