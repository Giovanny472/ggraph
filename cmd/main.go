package main

import "github.com/Giovanny472/ggraph/app"

func main() {

	// создание апп
	app := app.NewAppGraph()

	// настройка апп
	app.Config()

	// запуска апп
	app.Start()

}
