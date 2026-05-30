# 🎨 Resaltador Léxico Concurrente en Go: de un archivo a muchos

> Un Guided Project para construir, paso a paso, una herramienta en **Go** que lee
> varios archivos, reconoce categorías léxicas con tus propias reglas y resalta
> cada token. Primero lo harás **secuencial**; después lo harás **concurrente** con
> goroutines y channels, y compararás cuál es más rápido.

Este proyecto es la **continuación directa** del motor de expresiones regulares con
S-expressions que construiste en Python (`previous-lexer-project`). Allí aprendiste a
reconocer categorías léxicas (`IDENTIFICADOR`, `LIT_NUMERO`, `RESERVADA_IF`, …)
evaluando cadenas carácter por carácter con las ideas `letra`, `numero`, `seq` y
`star`. Aquí llevarás **esas mismas categorías** a Go y, sobre ellas, construirás una
aplicación que procesa **muchos archivos a la vez**.

---

## 🎯 Objetivos de aprendizaje

Al terminar serás capaz de:

- Leer y procesar el contenido de **múltiples archivos** en Go.
- Reutilizar las **categorías léxicas** del proyecto anterior como reconocedores en Go.
- Construir una versión **secuencial** funcional de un resaltador léxico.
- Evolucionar la solución a una versión **concurrente** con goroutines, `WaitGroup` y channels.
- Diseñar la concurrencia **evitando race conditions** y verificarlo con `go run -race`.
- **Medir y comparar** tiempos entre ambas versiones y redactar **conclusiones** propias.

---

## 🧠 Lo que ya sabes (y vas a reutilizar)

Durante el curso ya trabajaste con:

- Variables, tipos, `:=`, ciclos `for` y `range`.
- Funciones con parámetros y valores de retorno.
- Manejo de errores con el tipo `error` (`if err != nil { ... }`).
- Creación de archivos (`os.Create`, `defer file.Close()`).
- El paquete `strings`.
- Medición de tiempo con `time.Now()` y `time.Since()`.
- Goroutines (`go func(){ ... }()`).
- `sync.WaitGroup` (`Add`, `Done`, `Wait`).
- Channels (`make(chan ...)`, `<-`).
- Slices y `append`.

> 💡 También viste, **a propósito**, ejemplos donde las goroutines producían
> resultados raros porque compartían una variable (`counter`, `Tasks`). Ese problema
> se llama **race condition** y en este proyecto aprenderás a evitarlo por diseño.

### Conceptos nuevos que introduciremos (poco a poco)

No te preocupes si no los conoces todavía; cada uno aparece justo cuando lo necesitas:

- **Leer** el contenido de un archivo (no solo crearlo).
- **Argumentos de línea de comandos** (`os.Args`).
- **`struct`s** mínimos para agrupar un resultado.
- **Channels que transportan structs** (no solo `int` o `string`).
- La bandera **`-race`** para detectar condiciones de carrera.

---

## 🗂️ Estructura del proyecto

Trabajarás dentro de la carpeta `resaltadorGo/`. La estructura objetivo es:

```
resaltadorGo/
├── README.md            # Este documento
├── go.mod               # Necesario porque el código está dividido en paquetes (lexer, resaltador); sin él Go no resuelve las importaciones internas
├── main.go              # Punto de entrada (orquesta secuencial / concurrente)
├── lexer/
│   └── categorias.go    # Reconocedores: IDENTIFICADOR, LIT_NUMERO, ...
├── resaltador/
│   ├── archivo.go       # Leer y clasificar un archivo
│   └── resultado.go     # struct ResultadoArchivo / Token
├── data/                # Archivos de prueba (.txt con cadenas, uno por línea)
│   ├── ejemplo1.txt
│   ├── ejemplo2.txt
│   └── ...
└── HANDOFF.md           # Bitácora de relevos del equipo
```

> No tienes que crear todas las carpetas el primer día. Irán apareciendo conforme
> avances en los assignments.

---

## 👥 Plan de equipo (3 personas)

Este proyecto está pensado para **3 personas** trabajando **una por día**, de modo
que cada quien retoma donde lo dejó la persona anterior. Empiezan el **lunes 1 de
junio** y entregan el **domingo 7 de junio**.

| Día | Fecha | Responsable | Assignments |
|-----|-------|-------------|-------------|
| 1 | Lun 1 jun | **Persona A** | A0 · A1 · A2 |
| 2 | Mar 2 jun | **Persona B** | A3 |
| 3 | Mié 3 jun | **Persona C** | A4 |
| 4 | Jue 4 jun | **Persona A** | A5 · A6 |
| 5 | Vie 5 jun | **Persona B** | A7 |
| 6 | Sáb 6 jun | **Persona C** | A8 · A9 |
| 7 | Dom 7 jun | **Equipo** | Entrega: revisión conjunta y buffer |

### Reglas de relevo (handoff)

Para que el relevo funcione, **al terminar tu día**:

1. Asegúrate de que `go build ./...` **compila sin errores**.
2. Anota en `HANDOFF.md`: qué hiciste, cómo se ejecuta y **qué sigue**.
3. (Recomendado) Haz `git init` el primer día y un **commit por día/persona**.

> 💡 Cada persona toca el proyecto **dos veces** y su segundo turno continúa
> naturalmente lo que ya conoce. Si un día queda incompleto, la persona siguiente lo
> termina antes de empezar lo suyo.

---

## ▶️ Cómo ejecutar y verificar

```bash
# Compilar todo el proyecto
go build ./...

# Ejecutar pasando archivos
go run . data/ejemplo1.txt data/ejemplo2.txt

# (Desde el día 4) Detectar condiciones de carrera
go run -race . data/*.txt
```

> La bandera `-race` hace que Go vigile si dos goroutines acceden a la misma
> variable al mismo tiempo. Si imprime un `WARNING: DATA RACE`, tienes una carrera
> que resolver.

---

# 🧩 Assignments

Cada assignment tiene un **objetivo**, un poco de **contexto**, las **tareas** a
realizar, algunas **pistas** y un **criterio de listo** (cómo saber que terminaste).
Construye tú la solución: aquí encontrarás guía, no el código resuelto.

---

## 🟢 Día 1 — Persona A

### Assignment 0 — Preparación del proyecto

**Objetivo:** dejar el esqueleto listo para que el equipo trabaje sin fricción.

**Tareas**

1. Inicializa un módulo de Go dentro de `resaltadorGo/`:
   ```bash
   go mod init resaltadorgo
   ```
2. Crea la carpeta `data/` con **al menos 3 archivos** `.txt`. Cada archivo contiene
   cadenas, **una por línea**, al estilo del `strings.txt` del proyecto anterior
   (`abc123`, `if`, `while`, `true`, `123`, `1abc`, …).
3. Crea `HANDOFF.md` con una primera entrada (fecha, tu nombre, "proyecto inicializado").
4. (Recomendado) `git init` y primer commit.

**Pista:** reutiliza y amplía las cadenas de prueba del proyecto en Python; mientras
más variadas, mejor podrás probar tus categorías.

**Listo cuando:** `go version` corre, existe `go.mod`, y `data/` tiene varios `.txt`.

---

### Assignment 1 — Recordando el motor anterior

**Objetivo:** fijar la **continuidad** con tu motor de Python antes de escribir Go.

**Contexto.** En `previous-lexer-project` definiste estas categorías:

| Categoría | Qué acepta |
|-----------|-----------|
| `IDENTIFICADOR` | una letra seguida de letras o números |
| `LIT_NUMERO` | uno o más dígitos |
| `RESERVADA_IF` | exactamente `if` |
| `RESERVADA_WHILE` | exactamente `while` |
| `LIT_BOOLEANO` | `true` o `false` |

Y las construías con las operaciones `letra`, `numero`, `seq` (secuencia),
`or` (alternativa) y `star` (cero o más repeticiones).

**Tareas**

1. En `HANDOFF.md` (o en un comentario), escribe con tus palabras qué significa cada
   categoría y da **2 ejemplos** que **sí** acepta y **2** que **no**.
2. Decide cómo vas a representar estas categorías en Go (lo implementarás en A2).

**Pista:** en Go no tienes el motor de S-expressions. No lo necesitas: vas a portar la
**semántica** (la regla de cada categoría), no el parser. Piensa cada categoría como
una pregunta de sí/no sobre una cadena.

**Listo cuando:** puedes explicar, para cualquier cadena, a qué categoría pertenece.

---

### Assignment 2 — Reconocedores de categorías en Go

**Objetivo:** implementar funciones que clasifiquen una cadena.

**Contexto.** Cada categoría será una **función reconocedora** que recibe una cadena y
responde si pertenece o no. Una firma natural es:

```go
// Devuelve true si toda la cadena s corresponde a la categoría.
func esNumero(s string) bool
func esIdentificador(s string) bool
// ... y así para las demás
```

**Tareas**

1. Crea `lexer/categorias.go` (paquete `lexer`).
2. Implementa una función reconocedora por cada categoría del A1.
3. Crea una función que reciba una cadena y devuelva **el nombre de la categoría** a la
   que pertenece (o algo como `"DESCONOCIDO"` si no encaja en ninguna).

**Pistas**

- Para recorrer los caracteres de una cadena usa `for i, r := range s { ... }`, donde
  `r` es de tipo `rune` (un carácter).
- Para clasificar caracteres tienes el paquete estándar `unicode`:
  `unicode.IsLetter(r)` y `unicode.IsDigit(r)`. Son el equivalente de `isalpha()` e
  `isdigit()` de Python.
- Las palabras reservadas (`if`, `while`) y los booleanos son comparaciones directas de
  cadenas (`s == "if"`).
- Define un **orden de prioridad** para clasificar: por ejemplo, `if` debe reconocerse
  como `RESERVADA_IF` antes que como `IDENTIFICADOR`.

**Listo cuando:** dada una cadena suelta, tu código imprime su categoría correctamente
para todos los ejemplos del A1.

> 📌 **Handoff del Día 1:** `go build ./...` compila y el paquete `lexer` clasifica
> cadenas sueltas. Deja anotado en `HANDOFF.md` cómo probar una clasificación.

---

## 🔵 Día 2 — Persona B

### Assignment 3 — Leer y clasificar un archivo

**Objetivo:** procesar **un** archivo completo y mostrar el resaltado.

**Contexto: nuevos conceptos.**

- **Leer un archivo.** Hasta ahora solo creabas archivos. Para leer todo el contenido:
  ```go
  contenido, err := os.ReadFile(ruta)   // contenido es []byte
  if err != nil { /* maneja el error */ }
  texto := string(contenido)
  ```
- **Argumentos de línea de comandos.** En Python usabas `sys.argv`. En Go es
  `os.Args`, un slice de strings donde `os.Args[0]` es el nombre del programa y a
  partir de `os.Args[1]` vienen los argumentos.

**Tareas**

1. Crea `resaltador/archivo.go` (paquete `resaltador`).
2. Lee la ruta de **un** archivo desde `os.Args` y carga su contenido.
3. Separa el contenido en **tokens** (por líneas y/o por espacios) y usa tu paquete
   `lexer` para asignar una categoría a cada token.
4. Imprime el **resaltado**: cada token junto a su categoría, por ejemplo
   `abc123 → IDENTIFICADOR`.

**Pistas**

- Para separar por líneas: `strings.Split(texto, "\n")`. Para separar por espacios
  dentro de una línea: `strings.Fields(linea)`.
- Recuerda **limpiar** espacios sobrantes con `strings.TrimSpace` e ignorar líneas
  vacías (lo hacías en Python con `leer_strings`).
- Maneja el caso de que **no** se reciba ningún archivo: imprime un mensaje de uso y
  termina.

**Listo cuando:** `go run . data/ejemplo1.txt` imprime cada token con su categoría.

> 📌 **Handoff del Día 2:** un archivo se lee y se clasifica completo. Anota el comando
> exacto de prueba en `HANDOFF.md`.

---

## 🟣 Día 3 — Persona C

### Assignment 4 — Procesar varios archivos (versión secuencial)

**Objetivo:** procesar **una lista** de archivos, uno tras otro, y medir el tiempo.

**Contexto.** Esta es la **versión secuencial**: el programa toma todos los archivos y
los procesa en orden, uno completamente antes de empezar el siguiente. Es la base
contra la que después compararás la versión concurrente.

**Tareas**

1. Lee **todos** los archivos recibidos en `os.Args[1:]`.
2. Recórrelos con un `for` y clasifica cada uno reutilizando el A3.
3. Mide el **tiempo total** del procesamiento con `time.Now()` y `time.Since()`.
4. Imprime, al final, el tiempo total.

**Pistas**

- Ya viste el patrón de medición:
  ```go
  inicio := time.Now()
  // ... trabajo ...
  fmt.Println("Tiempo total:", time.Since(inicio))
  ```
- Puedes pasar muchos archivos con un patrón del shell: `go run . data/*.txt`.
- Para que la diferencia de tiempo se note más adelante, conviene tener **varios**
  archivos y/o archivos **grandes**. Puedes duplicar líneas en tus `.txt` de prueba.

**Listo cuando:** `go run . data/*.txt` procesa todos los archivos en serie e imprime
el tiempo total.

> 📌 **Handoff del Día 3:** versión secuencial completa y midiendo tiempo. Anota cuánto
> tardó con tus archivos de prueba.

---

## 🟢 Día 4 — Persona A

### Assignment 5 — Agrupar resultados con structs

**Objetivo:** introducir un `struct` para representar el resultado de cada archivo.

**Contexto: nuevo concepto — `struct`.** Un `struct` agrupa varios datos
relacionados bajo un mismo tipo. Por ejemplo:

```go
type ResultadoArchivo struct {
    Nombre  string
    Tokens  int           // o lo que decidas guardar
    Tiempo  time.Duration
}
```

Lo necesitas porque, en la versión concurrente, cada goroutine deberá devolver **su
resultado completo** (de qué archivo, cuántos tokens, cuánto tardó) en un solo valor.

**Tareas**

1. Crea `resaltador/resultado.go` con un `struct` que represente el resultado de
   procesar un archivo (decide qué campos te sirven).
2. Refactoriza la versión secuencial para que procesar un archivo **devuelva** un valor
   de ese `struct` en lugar de solo imprimir.
3. En `main`, recolecta los resultados en un slice y muéstralos al final.

**Pistas**

- Una función puede devolver un struct: `func procesar(ruta string) ResultadoArchivo`.
- Puedes ir guardando resultados con `append(resultados, r)`.
- Aún **no** introduzcas concurrencia; primero deja esta versión secuencial ordenada
  alrededor del struct.

**Listo cuando:** la versión secuencial sigue funcionando, pero ahora cada archivo
produce un `ResultadoArchivo`.

---

### Assignment 6 — Versión concurrente: una goroutine por archivo

**Objetivo:** lanzar el procesamiento de cada archivo en su propia goroutine y
**observar** el problema de las race conditions.

**Contexto.** Ya sabes lanzar goroutines y esperar con `sync.WaitGroup`. La idea: por
cada archivo, lanza `go func(){ ... }()` que lo procese; usa un `WaitGroup` para que
`main` no termine antes de tiempo (recuerda el error de medir el tiempo **antes** de
que las goroutines terminaran).

**Tareas**

1. Crea una versión concurrente donde cada archivo se procese en una goroutine.
2. Usa `WaitGroup` (`Add`, `Done`, `Wait`) para esperar a que todas terminen.
3. Intenta recolectar los resultados como lo hiciste en A5 (por ejemplo, haciendo
   `append` a un slice compartido **desde dentro** de cada goroutine).
4. Ejecuta con detección de carreras:
   ```bash
   go run -race . data/*.txt
   ```

**Pistas**

- Pasa el dato de cada iteración **como argumento** a la goroutine, igual que hiciste
  en clase: `go func(ruta string){ ... }(archivo)`. Evita capturar la variable del
  `for` directamente.
- Es muy probable que `-race` te reporte un **DATA RACE** al hacer `append` al mismo
  slice desde varias goroutines. **Eso es esperado** y es el punto de aprendizaje de
  hoy.

**Tareas de cierre**

5. Anota en `HANDOFF.md` **qué carrera detectó `-race`** y por qué crees que ocurre.
   La Persona B la resolverá mañana.

**Listo cuando:** existe una versión concurrente que corre con `WaitGroup`, y tienes
documentada la race condition detectada por `-race`.

> 📌 **Handoff del Día 4:** struct de resultado + versión concurrente que **expone** la
> carrera. No intentes resolverla hoy: ese es el trabajo del Día 5.

---

## 🔵 Día 5 — Persona B

### Assignment 7 — Recolectar resultados por channel (sin carreras)

**Objetivo:** eliminar la race condition usando un **channel** para recolectar los
resultados, sin compartir variables mutables.

**Contexto: por qué ocurría la carrera.** Cuando varias goroutines hacen `append` al
mismo slice (o escriben la misma variable) sin coordinación, pueden pisarse entre sí.
La forma idiomática en Go de evitarlo es **no compartir memoria**: que cada goroutine
**envíe** su resultado por un channel y que `main` los **reciba** y los junte. Es el
patrón **fan-in** que ya viste con channels de `int`.

**Tareas**

1. Crea un channel que transporte tu struct de resultado:
   ```go
   resultados := make(chan ResultadoArchivo)
   ```
2. Cada goroutine, al terminar de procesar su archivo, **envía** su resultado por el
   channel (`resultados <- r`) en lugar de hacer `append` a un slice compartido.
3. En `main`, **recibe** del channel tantos resultados como archivos haya y arma el
   slice final ahí (en una sola goroutine, sin carrera).
4. Verifica:
   ```bash
   go run -race . data/*.txt
   ```
   Ahora **no** debe aparecer ningún `DATA RACE`.

**Pistas**

- Necesitas saber **cuántos** resultados esperar (igual al número de archivos) para
  dejar de recibir. Recuerda el patrón de `ex_8`, donde recibías hasta llegar a un
  conteo.
- Cuidado con bloqueos: un channel **sin búfer** bloquea al emisor hasta que alguien
  recibe. Asegúrate de que `main` esté recibiendo mientras las goroutines envían.
- Cuando recolectas resultados solo en `main`, ya **no necesitas** que las goroutines
  toquen ninguna variable compartida: ahí está la clave para que desaparezca la
  carrera.

**Listo cuando:** `go run -race .` corre **sin advertencias** y los resultados son
correctos y completos.

> 📌 **Handoff del Día 5:** versión concurrente **libre de carreras**. Anota en
> `HANDOFF.md` que `-race` quedó limpio.

---

## 🟣 Día 6 — Persona C

### Assignment 8 — Medir y comparar

**Objetivo:** medir correctamente ambas versiones y compararlas.

**Contexto.** Para comparar de forma justa, mide el **tiempo total** de cada versión
sobre **el mismo conjunto** de archivos. Recuerda la lección de clase: en la versión
concurrente, mide **después** de `wg.Wait()` (o después de recibir todos los resultados
del channel), nunca antes.

**Tareas**

1. Permite ejecutar el programa en modo **secuencial** o **concurrente** (por ejemplo,
   con un argumento como `seq` / `conc`, o ejecutando ambas y mostrando las dos cifras).
2. Ejecuta ambas versiones con el mismo conjunto de archivos y registra los tiempos.
3. Construye una pequeña tabla comparativa con tus mediciones.

**Pistas**

- Para que la diferencia sea visible, usa **varios** archivos y/o archivos **grandes**.
  Con muy poco trabajo, la versión concurrente puede incluso parecer más lenta por el
  costo de crear goroutines: eso también es un hallazgo interesante.
- Corre cada medición **un par de veces**: los tiempos varían entre ejecuciones.

**Listo cuando:** tienes una tabla con tiempos de la versión secuencial vs concurrente.

---

### Assignment 9 — Conclusiones y reflexión

**Objetivo:** cerrar el proyecto con conclusiones basadas en **tus** resultados.

**Tareas**

Responde, en una sección de conclusiones (en `HANDOFF.md` o en un archivo aparte):

1. ¿Cuál versión fue más rápida con tus archivos? ¿Por qué crees que fue así?
2. ¿En qué situaciones la concurrencia **ayuda** y en cuáles **no vale la pena**?
3. ¿Qué cambio concreto (channel + recolectar en `main`) **eliminó** la race condition,
   y por qué?
4. ¿Qué pasaría si tuvieras **cientos** de archivos? ¿Y si fueran muy pequeños?

**Tareas de cierre**

5. Repasa el `HANDOFF.md` completo y deja el proyecto listo para la entrega.

**Listo cuando:** las conclusiones están escritas y respaldadas por las mediciones del
A8.

> 📌 **Handoff del Día 6:** proyecto completo, medido y con conclusiones.

---

## 🏁 Día 7 — Entrega (equipo)

Revisión final conjunta. Antes de entregar, verifiquen juntos:

- [ ] `go build ./...` compila sin errores.
- [ ] `go run . data/*.txt` produce el resaltado esperado (versión secuencial).
- [ ] `go run -race . data/*.txt` corre **sin** `DATA RACE` (versión concurrente).
- [ ] Existe la tabla comparativa de tiempos.
- [ ] Están escritas las conclusiones.
- [ ] `HANDOFF.md` refleja el trabajo de los 6 días.

> Este día es también **buffer**: si algún día quedó incompleto, aquí se termina.

---

## 📌 Reglas del proyecto (recordatorio)

- Reutiliza **las mismas categorías léxicas** del proyecto anterior: la continuidad
  debe notarse.
- Primero **secuencial y funcionando**; la concurrencia llega después.
- Evita las race conditions **por diseño** (no compartir memoria; comunicar por
  channels), y verifícalo con `-race`.
- Construye tú la solución: este README te da contexto, ejemplos pequeños y pistas, no
  el código resuelto.

¡Éxito! 🚀
