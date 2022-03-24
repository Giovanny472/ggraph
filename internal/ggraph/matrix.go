package ggraph

import "github.com/Giovanny472/ggraph/model"

type matrix struct {

	// тип матрицы
	tpm model.TypeMatrix

	// матрица инцидентности
	mat *model.GMatrix
}

func NewMatrix() model.IMatrix {
	return &matrix{}
}

// тип матрицы
func (mt *matrix) SetTypeMatrix(tpm model.TypeMatrix) {
	mt.tpm = tpm
}

// возвращение типа матрицы
func (mt *matrix) TypeMatrix() model.TypeMatrix {
	return mt.tpm
}

// назначение матрицы
func (mt *matrix) SetMatrix(mat *model.GMatrix) {
	mt.mat = mat
}

func (mt *matrix) Matrix() *model.GMatrix {
	return mt.mat
}
