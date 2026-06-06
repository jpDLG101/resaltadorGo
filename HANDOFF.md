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

## Día 5 — Persona B (RICKY, 5 Junio)

### Lo que se hizo

- Modifique el main.go, quite los `WaitGroup` xq ya no hacen falta, el loop de recepción espera hasta que ya esten todos los resultados y ya despues continua.
- Automaticamente se elimina sync de los imports porque ya no se usa.
- Se declaró el channel y trae los valores de `ResultadoArchivo` con length(archivos)
- Ya no se hace un append normal y cada goroutine se manda directamente a channel y ahi guarda los valores

### Cómo probar

```bash
go run -race . data/*.txt
```
### Resultado
que corra sin ningún WARNING: DATA RACE.

### Qué sigue (Día 6 — Persona C)

Assigment 8: permitir ejecutar en modo secuencial o concurrente, medir el tiempo de ambas versiones con el mismo conjunto de archivos, y construir una tabla comparativa

### Siento que esta la hagamos los 3, cuando acabes el A8 fabs, avisanos y hacemos esta juntos.
Assigment 9: escribir las conclusiones respondiendo las 4 preguntas del `README` (cuál fue más rápida, cuándo vale la pena la concurrencia, qué eliminó la race condition, qué pasaría con cientos de archivos)

---

## Día 6 — Persona C (Fabs, 6 Junio)

### Lo que se hizo

- A8: refactorizado `main.go` con dos funciones: `modoSecuencial` y `modoConcurrente`, ambas retornan resultados + tiempo de pared
- A8: nuevo flag `--modo=sec|con|ambos` (por defecto `ambos`)
- A8: cuando se usa `ambos`, corre los dos modos con los mismos archivos e imprime una tabla comparativa con tiempo por archivo y speedup total
- A9: conclusiones escritas abajo
- `go run -race . data/*.txt` → sin DATA RACE

### Cómo probar

```bash
# Modo ambos (default) — muestra tabla comparativa
go run . data/*.txt

# Solo secuencial
go run . --modo=sec data/*.txt

# Solo concurrente
go run . --modo=con data/*.txt

# Verificar sin race condition
go run -race . data/*.txt
```

Salida esperada (modo ambos):
```
--- Modo Secuencial ---
...
--- Modo Concurrente ---
...
=== TABLA COMPARATIVA ===
Archivo                        | Tokens |   Secuencial |  Concurrente
--------------------------------------------------------------------
data/ejemplo1.txt              |      5 |     23.083µs |     11.750µs
data/ejemplo2.txt              |      5 |      1.000µs |     10.583µs
data/ejemplo3.txt              |      5 |      1.167µs |      1.000µs
--------------------------------------------------------------------
TOTAL (tiempo de pared)        |        |   1.302291ms |    201.125µs
Speedup: 6.48x
```

---

## Assignment 9 — Conclusiones

**1. ¿Cuál versión fue más rápida?**
La versión concurrente fue más rápida en tiempo de pared total (~6x de speedup en nuestras pruebas). Cada goroutine procesa un archivo de forma independiente, y como las operaciones de I/O y clasificación no bloquean a las demás, el tiempo total se acerca al del archivo más lento en lugar de ser la suma de todos.

**2. ¿Cuándo vale la pena usar concurrencia?**
Vale la pena cuando hay múltiples tareas independientes que pueden ejecutarse en paralelo (como procesar archivos distintos) y el overhead de crear goroutines es pequeño frente al trabajo real. Con archivos muy pequeños o pocos archivos, el speedup puede ser marginal o incluso negativo por el costo de coordinar goroutines.

**3. ¿Qué eliminó la race condition?**
Sustituir el `append` directo a un slice compartido por un **channel buffered**. Con channels, cada goroutine envía su `ResultadoArchivo` de forma segura sin acceder a memoria compartida: el channel garantiza que solo un valor se escribe a la vez y `main` los recibe en orden sin conflictos.

**4. ¿Qué pasaría con cientos de archivos?**
La versión concurrente escalaría mejor: Go lanza una goroutine por archivo (ligeras, ~2KB de stack) y el scheduler las distribuye entre los núcleos disponibles. Con cientos de archivos, el tiempo total seguiría siendo cercano al del archivo más lento, mientras que la versión secuencial crecería linealmente. Sin embargo, con miles de archivos habría que considerar limitar la concurrencia (worker pool con semáforo) para no saturar el sistema de archivos.

