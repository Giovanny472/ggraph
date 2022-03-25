package manager

import (
	"errors"

	"github.com/Giovanny472/ggraph/model"
)

type manager struct {
	graph model.GGraph
	util  model.Utilities
}

var amanager *manager

func NewManager(grph model.GGraph, ut model.Utilities) model.Manager {
	if amanager == nil {
		amanager = &manager{grph, ut}
	}
	amanager.init()
	return amanager
}

func (man *manager) init() {
}

func (man *manager) GenerateGraph(txtMatrix string, tpgp model.TypeGraph, tpm model.TypeMatrix) error {

	// получение матрицы
	aMatrix, err := man.util.StrToGMatrix(txtMatrix)
	if err != nil {
		return errors.New("Ошибка при создании матрицы : " + err.Error())
	}

	// настройка типа матрицы
	if tpm == -1 {
		return errors.New("тип матрицы не выбран ")
	}
	man.graph.Matrix().SetTypeMatrix(tpm)

	// настройка типа графа
	man.graph.SetType(tpgp)

	// назначение матрицы
	man.graph.Matrix().SetMatrix(aMatrix)

	// создание графа
	man.graph.Create()

	// cохранение графа
	man.graph.Save(model.GraphFileName)

	return nil
}
