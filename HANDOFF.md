# HANDOFF

## Convención de commits

```
dia<N>(persona-<letra>): descripción breve
```

- `dia1(persona-a): setup inicial, categorias lexicas y handoff`
- `dia2(persona-b): leer y clasificar archivo`
- `dia3(persona-c): version secuencial con medicion de tiempo`

---

## Día 1 — Persona A (Jp, 1 Junio)

### Lo que se hizo

- Proyecto inicializado (`go mod init resaltadorgo`)
- Carpeta `data/` con 3 archivos de prueba (`ejemplo1.txt`, `ejemplo2.txt`, `ejemplo3.txt`)
- Lexer implementado en `lexer/categorias.go`

### Categorías léxicas

| Categoría | Descripción | Acepta | No acepta |
|-----------|-------------|--------|-----------|
| `IDENTIFICADOR` | una letra seguida de letras o números | `abc`, `x1` | `1abc`, `123` |
| `RESERVADA_IF` | exactamente `if` | `if` | `IF`, `iff` |
| `RESERVADA_WHILE` | exactamente `while` | `while` | `While`, `whilee` |
| `LIT_NUMERO` | uno o más dígitos | `123`, `0` | `12a`, `abc` |
| `LIT_BOOLEANO` | exactamente `true` o `false` | `true`, `false` | `True`, `tru` |

### Funciones en `lexer/categorias.go`

- `esNumero`, `esIdentificador`, `esReservadaIF`, `esReservadaWHILE`, `esBooleano` — reciben un `string`, regresan `bool`
- `Clasificar(s string) string` — regresa el nombre de la categoría o `DESCONOCIDO`

### Cómo probar

Crea un `main.go` temporal en la raíz del proyecto:

```go
package main

import (
    "fmt"
    "resaltadorgo/lexer"
)

func main() {
    palabras := []string{"abc123", "if", "while", "true", "false", "123", "1abc"}
    for _, p := range palabras {
        fmt.Println(p, "→", lexer.Clasificar(p))
    }
}
```

```bash
go run main.go
```

Borra el `main.go` cuando termines de probar.

---