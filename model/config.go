package model

type Config interface {
	Load()
	Save()
}
