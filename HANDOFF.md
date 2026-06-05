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

## Día 2 — Persona B (Ricky, 2 Junio)

### Lo que se hizo

- Creado `resaltador/archivo.go` (paquete `resaltador`) con la función `ProcesarArchivo`
- Creado `main.go` como punto de entrada: lee `os.Args[1]` y llama a `ProcesarArchivo`
- El programa lee un archivo, separa por líneas y espacios, y clasifica cada token con `lexer.Clasificar`
- `go build ./...` compila sin errores

### Cómo probar

```bash
go run . data/ejemplo1.txt
```

Salida esperada:
```
abc123 → IDENTIFICADOR
if → RESERVADA_IF
while → RESERVADA_WHILE
true → BOOLEANO
false → BOOLEANO
```

### Qué sigue (Fabs, Junio 3)

- Assignment 4: procesar **todos** los archivos de `os.Args[1:]` en un `for`, uno tras otro (versión secuencial)
- Medir el tiempo total con `time.Now()` / `time.Since()` e imprimirlo al final
- Probar con: `go run . data/*.txt`

---

## Día 3 — Persona C (Fabs, 2 Junio)

### Lo que se hizo

- Modificado `main.go` para procesar **todos** los archivos recibidos en `os.Args[1:]`
- Se itera con un `for range` llamando `resaltador.ProcesarArchivo` por cada ruta
- Se mide el tiempo total con `time.Now()` antes del loop y `time.Since()` al final
- `go build ./...` compila sin errores

### Cómo probar

```bash
go run . data/ejemplo1.txt data/ejemplo2.txt data/ejemplo3.txt
# o con wildcard:
go run . data/*.txt
```

Salida esperada (al final):
```
Tiempo total: 507.625µs
```

### Qué sigue (Persona A, Junio 4)

- Assignment 5: crear `resaltador/resultado.go` con un `struct ResultadoArchivo`
- Refactorizar para que `ProcesarArchivo` devuelva el struct en lugar de solo imprimir
- Recolectar resultados en un slice en `main`
- Assignment 6: versión concurrente con goroutines + WaitGroup, exponer la race condition con `go run -race . data/*.txt`

---

## Día 4 — Persona A (Jp, 4 Junio)

### Lo que se hizo

- A5: agregado `struct Token` y campo `Clasificaciones []Token` a `resaltador/resultado.go`
- A5: refactorizado `ProcesarArchivo` en `resaltador/archivo.go` para devolver `ResultadoArchivo` con nombre, conteo de tokens, tiempo y clasificaciones
- A5: `main.go` ahora recolecta resultados en una lista y los imprime al final
- A6: versión concurrente en `main.go` usando goroutines + `sync.WaitGroup`
- A6: detectada race condition con `go run -race . data/*.txt`

### Race condition detectada

Varias goroutines hacen `append` a la misma variable global `resultados` al mismo tiempo. Como todas la comparten sin coordinación, pueden pisarse entre sí al escribir — una goroutine puede sobreescribir lo que otra acaba de agregar o corromper la estructura interna.

### Cómo probar

```bash
go run -race . data/*.txt
```

### Qué sigue (Día 5 — Persona B)

- Assignment 7: eliminar la race condition usando un **channel** para recolectar resultados
- Cada goroutine envía su `ResultadoArchivo` por el channel en lugar de hacer `append`
- `main` recibe del channel y arma el slice (sin compartir variables)
- Verificar con `go run -race . data/*.txt` que no haya DATA RACE

---