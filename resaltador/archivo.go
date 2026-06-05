package resaltador

import (
	"fmt"
	"os"
	"resaltadorgo/lexer"
	"strings"
	"time"
)

func ProcesarArchivo(ruta string) ResultadoArchivo {
	contenido, err := os.ReadFile(ruta)
	if err != nil {
		fmt.Println("Error al leer archivo:", err)
		return ResultadoArchivo{}
	}

	texto := string(contenido)
	lineas := strings.Split(texto, "\n")

	conteo := 0
	var clasificaciones []Token
	inicio := time.Now()
	for _, linea := range lineas {
		linea = strings.TrimSpace(linea)
		if linea == "" {
			continue
		}
		for _, palabra := range strings.Fields(linea) {
			clasificaciones = append(clasificaciones, Token{palabra, lexer.Clasificar(palabra)})
			conteo++
		}
	}
	return ResultadoArchivo{ruta, conteo, time.Since(inicio), clasificaciones}
}
