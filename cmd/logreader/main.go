package main

import (
	"fmt"
	"logreader/internal/app"
)

func main() {
	if err := runApp(); err != nil {
		fmt.Println("Erro ao executar a aplicação:", err)
		return
	}
}

func runApp() error {
	return app.Run()
}
