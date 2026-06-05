package resaltador

import "time"

type Token struct {
	Palabra   string
	Categoria string
}

type ResultadoArchivo struct {
	Nombre          string
	Tokens          int
	Tiempo          time.Duration
	Clasificaciones []Token
}
