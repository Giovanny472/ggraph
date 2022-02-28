package ggraph

import (
	"fmt"
	"log"
	"strconv"

	"github.com/Giovanny472/ggraph/model"
	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

type ggraph struct {
}

var aggraph *ggraph

func NewGGraph() model.GGraph {
	if aggraph == nil {
		aggraph = new(ggraph)
	}
	return aggraph
}

func (gr *ggraph) SetSimpleAdjMatrix(adm *model.AdjMatrix) {

}

func (gr *ggraph) SetDirectedAdjMatrix(adm *model.AdjMatrix) {

	// создание графа
	agviz := graphviz.New()
	aggraph, err := agviz.Graph()
	if err != nil {
		log.Println("aggraph err: " + err.Error())
	}
	defer func() {
		if err := aggraph.Close(); err != nil {
			log.Fatal(err)
		}
		agviz.Close()
	}()

	// количество вершин
	aCountNodes := len(*adm)
	fmt.Println("size nodes--> ", aCountNodes)

	// количество ребер
	aCountEdge := len((*adm)[0])
	fmt.Println("size_edge --> ", aCountEdge)

	// cоздание набора вершин
	var listNodes []*cgraph.Node
	for idx := 0; idx < aCountNodes; idx++ {

		anode, _ := aggraph.CreateNode("N" + strconv.Itoa(idx+1))
		listNodes = append(listNodes, anode)
	}

	// определение набора ребер
	var alistEdges model.AdjListEdge

	for idxCol := 0; idxCol < aCountEdge; idxCol++ {

		// ребер
		var aEdge *model.AdjEdge

		// cоздание
		aEdge = new(model.AdjEdge)

		for idxRow := 0; idxRow < aCountNodes; idxRow++ {

			// получаем значение
			aval := (*adm)[idxRow][idxCol]
			fmt.Println("aval --> ", aval)

			// проверка
			if aval == 0 {
				continue
			}

			if aval == -1 || aval == 1 {

				// название ребра
				aEdge.NodeName = "e" + strconv.Itoa(idxCol+1)

				// start
				if aval == -1 {
					aEdge.NodeStart = listNodes[idxRow]
				}

				// end
				if aval == 1 {
					aEdge.NodeEnd = listNodes[idxRow]
				}

				fmt.Println("aEdge: ", aEdge)
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
		fmt.Println(alistEdges)
		fmt.Println("--------------")
	}

	for _, anode := range alistEdges {

		aed, _ := aggraph.CreateEdge(anode.NodeName, anode.NodeStart, anode.NodeEnd)
		aed.SetLabel(anode.NodeName)
	}
	// создание
	agviz.RenderFilename(aggraph, graphviz.PNG, "graph.png")
}

func (gr *ggraph) Save() {

}

func (gr *ggraph) Create() {

}
