package model

type Manager interface {

	// генерируем граф
	GenerateGraph(txtMatrix string, tpgp TypeGraph, tpm TypeMatrix) error
}
