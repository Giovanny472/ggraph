package ggraph

import (
	"fmt"
	"log"

	"github.com/Giovanny472/ggraph/model"
	"github.com/goccy/go-graphviz"
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
	gviz := graphviz.New()
	graph01, err := gviz.Graph()
	if err != nil {
		log.Println("graph01 err: " + err.Error())
	}
	defer func() {
		if err := graph01.Close(); err != nil {
			log.Fatal(err)
		}
		gviz.Close()
	}()

	// количество вершин
	aCountNodes := len(*adm)
	fmt.Println("size nodes--> ", aCountNodes)

	// количество ребер
	aCountEdge := len((*adm)[0])
	fmt.Println("size_edge --> ", aCountEdge)

}

func (gr *ggraph) Save() {

}

func (gr *ggraph) Create() {

}
