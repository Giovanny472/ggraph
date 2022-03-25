package ggraph

import (
	"fmt"
	"strconv"

	"github.com/Giovanny472/ggraph/model"
	"github.com/goccy/go-graphviz/cgraph"
)

func CreateGraphFromAdjMatrix(gp *ggraph, typegp model.TypeGraph) {

	switch typegp {
	case model.TypeNotDiGraph:
		CreateNotDiGraphFromAdjMatrix(gp)
	case model.TypeDiGraph:
		CreateDiGraphFromAdjMatrix(gp)
	}
}

// ориентированный граф
func CreateDiGraphFromAdjMatrix(gp *ggraph) {

	// создание графа
	gp.gvizgraph, _ = gp.gviz.Graph()

	// количество вершин
	aCountNodes := len(*gp.Matrix().Matrix())

	// количество ребер
	aCountEdge := len((*gp.Matrix().Matrix())[0])

	if aCountNodes != aCountEdge {
		fmt.Println("Размер матрицы неверен")
		return
	}

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
			aval := (*gp.Matrix().Matrix())[idxRow][idxCol]

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

	// создание графа
	gp.gvizgraph, _ = gp.gviz.Graph()

	// количество вершин
	aCountNodes := len(*gp.Matrix().Matrix())

	// количество ребер
	aCountEdge := len(*gp.Matrix().Matrix())

	// для определения названия вершин
	var alblCountEdge int = 0

	// cоздание набора вершин
	var listNodes model.ListNodes

	for idx := 0; idx < aCountNodes; idx++ {

		anode, _ := gp.gvizgraph.CreateNode(model.PrefixVertex + strconv.Itoa(idx+1))
		listNodes = append(listNodes, anode)
	}

	// определение набора ребер
	var alistEdges model.ListEdge

	for idxRow := 0; idxRow < aCountEdge; idxRow++ {

		for idxCol := 0; idxCol < aCountNodes; idxCol++ {

			// получаем значение
			aval := (*gp.Matrix().Matrix())[idxRow][idxCol]

			// проверка
			if aval == 0 {
				continue
			}

			// ребер
			var aEdge *model.GEdge = nil

			if aval == 1 {

				// cоздание
				aEdge = new(model.GEdge)

				// start
				aEdge.NodeStart = listNodes[idxRow]

				// end
				aEdge.NodeEnd = listNodes[idxCol]
			}

			if aEdge != nil {

				if aEdge.NodeStart != nil &&
					aEdge.NodeEnd != nil {

					var isitem bool = false
					for _, aitem := range alistEdges {
						if aitem.NodeStart == aEdge.NodeEnd && aitem.NodeEnd == aEdge.NodeStart {
							isitem = true
							break
						}
					}
					if !isitem {

						// название ребра
						alblCountEdge++
						aEdge.EdgeName = model.PrefixEdge + strconv.Itoa(alblCountEdge)

						alistEdges = append(alistEdges, aEdge)
					}
				}
			}

		}
	}

	for _, anode := range alistEdges {
		aed, _ := gp.gvizgraph.CreateEdge(anode.EdgeName, anode.NodeStart, anode.NodeEnd)
		aed.SetArrowHead(cgraph.NoneArrow)
		//aed.SetColor("orange")
		//aed.SetStyle(cgraph.BoldEdgeStyle)
		aed.SetLabel(anode.EdgeName)
	}

}
