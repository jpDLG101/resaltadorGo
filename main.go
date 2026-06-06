package main

import (
	"fmt"
	"os"
	"resaltadorgo/resaltador"
	"strings"
	"time"
)

func modoSecuencial(archivos []string) ([]resaltador.ResultadoArchivo, time.Duration) {
	inicio := time.Now()
	var resultados []resaltador.ResultadoArchivo
	for _, ruta := range archivos {
		resultados = append(resultados, resaltador.ProcesarArchivo(ruta))
	}
	return resultados, time.Since(inicio)
}

func modoConcurrente(archivos []string) ([]resaltador.ResultadoArchivo, time.Duration) {
	inicio := time.Now()
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
	return resultados, time.Since(inicio)
}

func imprimirResultados(resultados []resaltador.ResultadoArchivo) {
	for _, r := range resultados {
		fmt.Printf("=== %s | Tokens: %d | Tiempo: %s ===\n", r.Nombre, r.Tokens, r.Tiempo)
		for _, t := range r.Clasificaciones {
			fmt.Printf("  %s → %s\n", t.Palabra, t.Categoria)
		}
	}
}

func imprimirTabla(resSec []resaltador.ResultadoArchivo, tSec time.Duration, resCon []resaltador.ResultadoArchivo, tCon time.Duration) {
	mapSec := make(map[string]resaltador.ResultadoArchivo)
	for _, r := range resSec {
		mapSec[r.Nombre] = r
	}

	sep := strings.Repeat("-", 68)
	fmt.Println("\n=== TABLA COMPARATIVA ===")
	fmt.Printf("%-30s | %6s | %12s | %12s\n", "Archivo", "Tokens", "Secuencial", "Concurrente")
	fmt.Println(sep)
	for _, rc := range resCon {
		rs := mapSec[rc.Nombre]
		fmt.Printf("%-30s | %6d | %12s | %12s\n", rc.Nombre, rc.Tokens, rs.Tiempo, rc.Tiempo)
	}
	fmt.Println(sep)
	fmt.Printf("%-30s | %6s | %12s | %12s\n", "TOTAL (tiempo de pared)", "", tSec, tCon)
	if tCon > 0 {
		fmt.Printf("Speedup: %.2fx\n", float64(tSec)/float64(tCon))
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run . [--modo=sec|con|ambos] <archivo.txt> [archivo2.txt ...]")
		return
	}

	modo := "ambos"
	var archivos []string
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "--modo=") {
			modo = strings.TrimPrefix(arg, "--modo=")
		} else {
			archivos = append(archivos, arg)
		}
	}

	if len(archivos) == 0 {
		fmt.Println("Error: no se proporcionaron archivos.")
		return
	}

	switch modo {
	case "sec":
		resultados, tiempo := modoSecuencial(archivos)
		imprimirResultados(resultados)
		fmt.Printf("\nTiempo total (secuencial): %s\n", tiempo)
	case "con":
		resultados, tiempo := modoConcurrente(archivos)
		imprimirResultados(resultados)
		fmt.Printf("\nTiempo total (concurrente): %s\n", tiempo)
	case "ambos":
		fmt.Println("--- Modo Secuencial ---")
		resSec, tSec := modoSecuencial(archivos)
		imprimirResultados(resSec)

		fmt.Println("\n--- Modo Concurrente ---")
		resCon, tCon := modoConcurrente(archivos)
		imprimirResultados(resCon)

		imprimirTabla(resSec, tSec, resCon, tCon)
	default:
		fmt.Printf("Modo desconocido: %s. Usa --modo=sec, --modo=con, o --modo=ambos.\n", modo)
	}
}
