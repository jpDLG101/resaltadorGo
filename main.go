package main

import (
	"fmt"
	"os"

	"resaltadorgo/resaltador"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run . <archivo.txt>")
		return
	}

	resaltador.ProcesarArchivo(os.Args[1])
}
