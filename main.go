package main

import (
	"fmt"
	"os"
	"resaltadorgo/resaltador"
	"sync"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run . <archivo.txt> [archivo2.txt ...]")
		return
	}

	resultados := []resaltador.ResultadoArchivo{}
	
	var wg sync.WaitGroup

	for _, ruta := range os.Args[1:] {
		wg.Add(1)
		go func (ruta string)  {
			defer wg.Done()
			r := resaltador.ProcesarArchivo(ruta)
			resultados = append(resultados, r)
		}(ruta)
	}
	wg.Wait()

	for _, r := range resultados {
		fmt.Printf("=== %s | Tokens: %d | Tiempo: %s ===\n", r.Nombre, r.Tokens, r.Tiempo)
		for _, t := range r.Clasificaciones {
			fmt.Printf("  %s → %s\n", t.Palabra, t.Categoria)
		}
	}
}