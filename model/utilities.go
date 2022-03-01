package model

const (
	PrefixVertex  string = "V"
	PrefixEdge    string = "e"
	GraphFileName string = "graph.png"
)

type Utilities interface {
	StrToGMatrix(data string) (*GMatrix, error)
}
