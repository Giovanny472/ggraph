package ggraph

import (
	"strconv"

	"github.com/Giovanny472/ggraph/model"
	"github.com/goccy/go-graphviz/cgraph"
)

func CreateGraphFromIncMatrix(gp *ggraph, typegp model.TypeGraph) {

	switch typegp {
	case model.TypeNotDiGraph:
		CreateNotDiGraphFromIncMatrix(gp)
	case model.TypeDiGraph:
		CreateDiGraphFromIncMatrix(gp)
	}
}

func CreateNotDiGraphFromIncMatrix(gp *ggraph) {

	// создание графа
	gp.gvizgraph, _ = gp.gviz.Graph()
	//defer gp.gvizgraph.Close()

	// количество вершин
	aCountNodes := len(*gp.Matrix().Matrix())

	// количество ребер
	aCountEdge := len((*gp.Matrix().Matrix())[0])

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
		aEdge.NodeStart = nil
		aEdge.NodeEnd = nil

		for idxRow := 0; idxRow < aCountNodes; idxRow++ {

			// получаем значение
			aval := (*gp.Matrix().Matrix())[idxRow][idxCol]

			// проверка
			if aval == 0 {
				continue
			}

			// start - end - петля
			if aval == 1 {

				if aEdge.NodeStart == nil {
					aEdge.NodeStart = listNodes[idxRow]
				} else {
					aEdge.NodeEnd = listNodes[idxRow]
				}

				// название ребра
				if aEdge.NodeStart != nil && aEdge.NodeEnd != nil {
					aEdge.EdgeName = model.PrefixEdge + strconv.Itoa(idxCol+1)
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
		aed.SetArrowHead(cgraph.NoneArrow)
		//aed.SetColor("orange")
		//aed.SetStyle(cgraph.BoldEdgeStyle)
		aed.SetLabel(anode.EdgeName)
	}

}

func CreateDiGraphFromIncMatrix(gp *ggraph) {

	// создание графа
	gp.gvizgraph, _ = gp.gviz.Graph()
	//defer gp.gvizgraph.Close()

	// количество вершин
	aCountNodes := len(*gp.Matrix().Matrix())

	// количество ребер
	aCountEdge := len((*gp.Matrix().Matrix())[0])

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
			aval := (*gp.Matrix().Matrix())[idxRow][idxCol]

			// проверка
			if aval == 0 {
				continue
			}

			if aval == -1 || aval == 1 || aval == 2 {

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

				// start - end - петля
				if aval == 2 {
					aEdge.NodeStart = listNodes[idxRow]
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
		//aed.SetArrowHead(cgraph.NoneArrow)
		//aed.SetColor("orange")
		//aed.SetStyle(cgraph.BoldEdgeStyle)
		aed.SetLabel(anode.EdgeName)
	}

}
