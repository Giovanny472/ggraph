build-ggraph:
	@go build -o .\bin\ggraph.exe .\cmd\main.go

run-ggraph: build-ggraph
	@.\bin\ggraph.exe

