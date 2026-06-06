package main

import (
	"fmt"
	"os"
	"resaltadorgo/resaltador"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run . <archivo.txt> [archivo2.txt ...]")
		return
	}

	archivos := os.Args[1:]
	ch := make(chan resaltador.ResultadoArchivo, len(archivos))

	for _, ruta := range archivos {
		go func(ruta string) {
			ch <- resaltador.ProcesarArchivo(ruta)
		}(ruta)
	}

	var resultados []resaltador.ResultadoArchivo
	for range archivos {
		resultados = append(resultados, <-ch)
	}

	for _, r := range resultados {
		fmt.Printf("=== %s | Tokens: %d | Tiempo: %s ===\n", r.Nombre, r.Tokens, r.Tiempo)
		for _, t := range r.Clasificaciones {
			fmt.Printf("  %s → %s\n", t.Palabra, t.Categoria)
		}
	}
}
