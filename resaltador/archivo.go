package resaltador

import (
	"fmt"
	"os"
	"strings"

	"resaltadorgo/lexer"
)

func ProcesarArchivo(ruta string) {
	contenido, err := os.ReadFile(ruta)
	if err != nil {
		fmt.Println("Error al leer archivo:", err)
		return
	}

	texto := string(contenido)
	lineas := strings.Split(texto, "\n")

	for _, linea := range lineas {
		linea = strings.TrimSpace(linea)
		if linea == "" {
			continue
		}
		tokens := strings.Fields(linea)
		for _, token := range tokens {
			fmt.Printf("%s → %s\n", token, lexer.Clasificar(token))
		}
	}
}
