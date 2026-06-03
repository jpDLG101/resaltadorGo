package main

import (
	"fmt"
	"os"
	"time"

	"resaltadorgo/resaltador"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run . <archivo.txt> [archivo2.txt ...]")
		return
	}

	inicio := time.Now()

	for _, ruta := range os.Args[1:] {
		fmt.Printf("=== %s ===\n", ruta)
		resaltador.ProcesarArchivo(ruta)
	}

	fmt.Printf("\nTiempo total: %s\n", time.Since(inicio))
}
