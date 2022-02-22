package config

import "github.com/Giovanny472/ggraph/model"

type config struct {
}

var aconfig *config

func NewConfig() model.Config {
	if aconfig == nil {
		aconfig = new(config)
	}
	return aconfig
}

func (conf *config) Load() {

}

func (conf *config) Save() {

}
