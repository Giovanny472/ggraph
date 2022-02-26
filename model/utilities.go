package model

type Utilities interface {
	StrToAdjMatrix(data string) (*AdjMatrix, error)
}
