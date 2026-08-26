# Practico 1 - Programacion Eficiente

## Analisis de algoritmos de busqueda, conteo, deteccion y ordenamiento

---

## Indice

1. [Metodologia](#1-metodologia)
2. [Ejercicio 1: Busqueda Lineal](#2-ejercicio-1-busqueda-lineal)
3. [Ejercicio 2: Conteo de Numeros Pares](#3-ejercicio-2-conteo-de-numeros-pares)
4. [Ejercicio 3: Deteccion de Elementos Repetidos](#4-ejercicio-3-deteccion-de-elementos-repetidos)
5. [Ejercicio 4: Ordenamiento](#5-ejercicio-4-ordenamiento)
6. [Conclusiones](#6-conclusiones)
7. [Como ejecutar](#7-como-ejecutar)

---

## 1. Metodologia

### 1.1 Entorno de ejecucion

- **Sistema operativo:** Windows (PowerShell)
- **Lenguaje:** Go 1.26.5
- **go.mod:** modulo `practico1`

### 1.2 Medicion de tiempos

Se utiliza `time.Now()` y `time.Since()` para medir el tiempo de ejecucion de cada algoritmo. Para mediciones mas precisas, el informe original del PDF emplea `testing.Benchmark` con tecnicas de:
- Consumo del resultado (evitar eliminacion de codigo muerto por el compilador)
- Recoleccion de basura previa con `runtime.GC()`
- Restauracion de datos en algoritmos de ordenamiento

### 1.3 Sobre los slices en Go

Un slice en Go es un descriptor compuesto por un puntero al arreglo subyacente, la longitud y la capacidad. Pasarlo como argumento copia unicamente el descriptor (costo constante), pero las modificaciones son visibles para el invocador. Por eso los algoritmos de ordenamiento trabajan sobre copias explicitas.

---

## 2. Ejercicio 1: Busqueda Lineal

### 2.1 Consigna

Desarrollar un algoritmo que reciba una lista de numeros enteros y un numero a buscar. El algoritmo debe indicar si el numero existe dentro de la lista. No utilizar funciones propias del lenguaje como `contains()` o `in`; comparar luego usando estas funciones. Medir tiempos e indicar la complejidad algoritmica.

### 2.2 Codigo: `Ej1.go`

```go
package main

import (
    "fmt"
    "slices"
    "time"
)

func busquedaLineal(datos []int, valor int) bool {
    for _, v := range datos {
        if v == valor {
            return true
        }
    }
    return false
}

func main() {
    datos := make([]int, 100000)
    for i := range datos {
        datos[i] = i * 2
    }
    valor := 99998

    inicio := time.Now()
    resultadoLineal := busquedaLineal(datos, valor)
    tiempoLineal := time.Since(inicio)

    inicio = time.Now()
    resultadoContains := slices.Contains(datos, valor)
    tiempoContains := time.Since(inicio)

    fmt.Printf("Busqueda de %d: %t (lineal) / %t (slices.Contains)\n", valor, resultadoLineal, resultadoContains)
    fmt.Printf("Tiempo busqueda lineal: %v\n", tiempoLineal)
    fmt.Printf("Tiempo slices.Contains: %v\n", tiempoContains)
    fmt.Println("Complejidad: O(n) tiempo y O(1) espacio para ambos algoritmos.")
}
```

### 2.3 Explicacion del codigo

El algoritmo implementa **busqueda lineal** (sequential search):

1. Se genera una lista de 100,000 enteros donde cada valor es `i * 2` (valores pares: 0, 2, 4, ..., 199998).
2. Se busca el valor `99998` (el penultimo elemento).
3. Se implementa la busqueda manualmente con `range`, recorriendo cada elemento hasta encontrarlo o agotar la lista.
4. Se compara con `slices.Contains` de la biblioteca estandar de Go.

**Detalle del bucle `range`:** Go retorna en cada iteracion el indice y el valor. El indice se descarta con `_` porque no se necesita. Si el valor actual coincide con el buscado, se retorna `true` inmediatamente (terminacion temprana). Si se recorre toda la lista sin encontrarlo, se retorna `false`.

### 2.4 Salida en consola

```
Busqueda de 99998: true (lineal) / true (slices.Contains)
Tiempo busqueda lineal: 0s
Tiempo slices.Contains: 0s
Complejidad: O(n) tiempo y O(1) espacio para ambos algoritmos.
```

> **Nota:** Los tiempos aparecen como 0s porque la lista es muy chica para el-granularidad del reloj. El informe original del PDF usa `testing.Benchmark` con 10^3, 10^4 y 10^5 elementos para obtener mediciones significativas.

### 2.5 Analisis teorico

| Caso | Tiempo | Descripcion |
|------|--------|-------------|
| Mejor | O(1) | El valor esta en la primera posicion |
| Promedio | O(n) | Se recorre en promedio la mitad del arreglo |
| Peor | O(n) | El valor no existe: se recorre por completo |

**Espacio:** O(1) - no se usa memoria auxiliar.

**Observaciones del PDF:**
- La implementacion manual y `slices.Contains` son **equivalentes en rendimiento**, ya que ambas compilan al mismo bucle lineal.
- La busqueda manual con `range` es ~43% mas rapida que con indice explicito para n=10^5, porque el compilador puede omitir verificaciones de rango.
- El cociente peor-caso/promedio es ~2, porque buscar un valor ausente recorre todo el arreglo, mientras que uno presente recorre en promedio la mitad.

---

## 3. Ejercicio 2: Conteo de Numeros Pares

### 3.1 Consigna

Dada una lista de numeros enteros, contar cuantos numeros pares contiene. Probar con diferentes algoritmos, medir tiempos e indicar la complejidad algoritmica.

### 3.2 Codigo: `Ej2.go`

```go
package main

import (
    "fmt"
    "time"
)

// Version 1: usando operador modulo (%)
func contarParesMod(datos []int) int {
    var contador int
    for _, v := range datos {
        if v%2 == 0 {
            contador++
        }
    }
    return contador
}

// Version 2: usando mascara de bits (&1)
func contarParesBit(datos []int) int {
    var contador int
    for _, v := range datos {
        if v&1 == 0 {
            contador++
        }
    }
    return contador
}

// Version 3: sin bifurcacion (sin if)
func contarParesSinRama(datos []int) int {
    var contador int
    for _, v := range datos {
        contador += 1 - (v & 1)
    }
    return contador
}

func medirConteo(nombre string, datos []int, algoritmo func([]int) int) {
    inicio := time.Now()
    resultado := algoritmo(datos)
    fmt.Printf("%-22s resultado: %d, tiempo: %v\n", nombre, resultado, time.Since(inicio))
}

func main() {
    datos := make([]int, 1000000)
    for i := range datos {
        datos[i] = i - 500000
    }

    medirConteo("Modulo", datos, contarParesMod)
    medirConteo("Bitwise", datos, contarParesBit)
    medirConteo("Sin rama", datos, contarParesSinRama)
    fmt.Println("Complejidad: O(n) tiempo y O(1) espacio en los tres algoritmos.")
}
```

### 3.3 Explicacion del codigo

Se implementan **tres variantes** para resolver el mismo problema:

**Version 1 - Modulo (`v%2 == 0`):** Verifica si el residuo de dividir por 2 es cero. Es la forma mas directa e intuitiva. Un numero es par si y solo si `v % 2 == 0`.

**Version 2 - Mascara de bits (`v&1 == 0`):** En representacion binaria, la paridad depende exclusivamente del bit menos significativo. `v & 1` extrae ese bit: es 0 si el numero es par, 1 si es impar. Es conceptualmente equivalente al modulo.

**Version 3 - Sin bifurcacion:** Elimina el `if` usando aritmeteria: `1 - (v & 1)` devuelve 1 si el numero es par (bit 0) y 0 si es impar (bit 1). Se suma directamente al contador sin salto condicional.

**Generacion de datos:** Se crean 1,000,000 enteros desde -500,000 hasta 499,999, incluyendo negativos. Esto prueba que la mascara de bits funciona correctamente con negativos en complemento a dos, mientras que `v%2 == 1` fallaria (en Go, `-3 % 2 == -1`).

### 3.4 Salida en consola

```
Modulo                 resultado: 500000, tiempo: 2.0187ms
Bitwise                resultado: 500000, tiempo: 1.3107ms
Sin rama               resultado: 500000, tiempo: 1.4449ms
Complejidad: O(n) tiempo y O(1) espacio en los tres algoritmos.
```

### 3.5 Analisis teorico

**Complejidad de las tres variantes:** O(n) tiempo, O(1) espacio.

| Algoritmo | Mejor | Promedio | Peor |
|-----------|-------|----------|------|
| Modulo | O(n) | O(n) | O(n) |
| Bitwise | O(n) | O(n) | O(n) |
| Sin rama | O(n) | O(n) | O(n) |

> A diferencia de la busqueda, **no existen mejor ni peor caso**: determinar la cantidad de pares exige inspeccionar la totalidad de los elementos.

**Observaciones del PDF:**
- **Modulo y bitwise son indistinguibles:** El compilador de Go reconoce `% 2` con divisor constante potencia de dos y lo traduce automaticamente a operacion de bits. La optimizacion manual es innecesaria.
- **La version sin bifurcacion no mejoro el rendimiento:** El compilador ya generaba codigo sin saltos usando instrucciones de asignacion condicional.
- **`slices.IndexFunc` fue ~7 veces mas lenta:** No existe `CountFunc` en Go; la aproximacion con llamadas sucesivas a `IndexFunc` incurria en overhead de llamada a funcion por cada par encontrado.

---

## 4. Ejercicio 3: Deteccion de Elementos Repetidos

### 4.1 Consigna

Desarrollar un algoritmo que determine si existen numeros repetidos dentro de una lista. No utilizar estructuras como `set` o `HashSet`; comparar luego usando estas estructuras. Medir tiempos e indicar la complejidad algoritmica.

### 4.2 Codigo: `Ej3.go`

```go
package main

import (
    "fmt"
    "slices"
    "time"
)

// Version 1: fuerza bruta O(n^2)
func repetidosFuerzaBruta(datos []int) bool {
    for i := range datos {
        for j := i + 1; j < len(datos); j++ {
            if datos[j] == datos[i] {
                return true
            }
        }
    }
    return false
}

// Version 2: ordenamiento previo O(n log n)
func repetidosOrdenando(datos []int) bool {
    copia := make([]int, len(datos))
    copy(copia, datos)
    slices.Sort(copia)

    for i := 1; i < len(copia); i++ {
        if copia[i] == copia[i-1] {
            return true
        }
    }
    return false
}

// Version 3: tabla de dispersion (map) O(n)
func repetidosMapa(datos []int) bool {
    vistos := make(map[int]struct{}, len(datos))
    for _, v := range datos {
        if _, existe := vistos[v]; existe {
            return true
        }
        vistos[v] = struct{}{}
    }
    return false
}

func medirRepetidos(nombre string, datos []int, algoritmo func([]int) bool) {
    inicio := time.Now()
    resultado := algoritmo(datos)
    fmt.Printf("%-24s resultado: %t, tiempo: %v\n", nombre, resultado, time.Since(inicio))
}

func main() {
    datos := make([]int, 20000)
    for i := range datos {
        datos[i] = i
    }
    datos[len(datos)-1] = datos[0] // insertamos un duplicado

    medirRepetidos("Fuerza bruta", datos, repetidosFuerzaBruta)
    medirRepetidos("Ordenamiento previo", datos, repetidosOrdenando)
    medirRepetidos("Mapa", datos, repetidosMapa)

    copia := slices.Clone(datos)
    slices.Sort(copia)
    fmt.Printf("Comparacion con slices.Sort: %t\n",
        repetidosOrdenando(datos) == slices.Contains(copia, datos[0]))
    fmt.Println("Complejidad: fuerza bruta O(n^2) y O(1) espacio; " +
        "ordenamiento O(n log n) y O(n) espacio; mapa O(n) promedio y O(n) espacio.")
}
```

### 4.3 Explicacion del codigo

**Version 1 - Fuerza bruta O(n^2):**
Compara cada par de elementos. El bucle interno comienza en `i + 1` (no en 0) para: (a) no comparar un elemento consigo mismo, y (b) recorrer cada par exactamente una vez. Total de comparaciones: `n(n-1)/2`.

**Version 2 - Ordenamiento previo O(n log n):**
1. Copia el arreglo para no destruir los datos originales (costo O(n) en espacio).
2. Ordena con `slices.Sort` (O(n log n)).
3. Recorre una sola vez comparando elementos contiguos: despues de ordenar, los iguales quedan necesariamente juntos.

**Version 3 - Mapa (tabla de dispersion) O(n):**
Go no tiene tipo `set`; se usa `map[int]struct{}`. La estructura vacia `struct{}` no ocupa memoria, asi que solo se paga por las claves. Para cada elemento se verifica si ya existe en el mapa; si existe, hay duplicado. Si no, se agrega.

**Generacion de datos:** Se crean 20,000 enteros del 0 al 19,999 y luego se reemplaza el ultimo elemento por el primero (`datos[0] = 0`), garantizando exactamente un duplicado al final.

### 4.4 Salida en consola

```
Fuerza bruta             resultado: true, tiempo: 0s
Ordenamiento previo      resultado: true, tiempo: 0s
Mapa                     resultado: true, tiempo: 554.1µs
Comparacion con slices.Sort: true
Complejidad: fuerza bruta O(n^2) y O(1) espacio; ordenamiento O(n log n) y O(n) espacio; mapa O(n) promedio y O(n) espacio.
```

### 4.5 Analisis teorico

| Version | Peor caso | Mejor caso | Espacio |
|---------|-----------|------------|---------|
| Fuerza bruta | O(n^2) | O(1) | O(1) |
| Ordenar + vecinos | O(n log n) | O(n log n) | O(n) |
| Mapa | O(n) promedio | O(1) | O(n) |

**Observaciones del PDF:**
- **El contraste entre peor y mejor caso es enorme:** Con repetidos la fuerza bruta detecta la colision casi de inmediato (229 ns para n=10^3), siendo la mas rapida. Sin repetidos, tarda 542,378 us.
- **Un algoritmo O(n^2) puede vencer a uno O(n)** cuando la terminacion temprana es probable y el costo de preparacion del segundo es alto.
- **Reservar capacidad del mapa puede ser contraproducente:** Mejora el peor caso (evita rehashings) pero degrada el caso favorable por un factor de 29 cuando el algoritmo termina rapido.
- **El punto de cruze** entre ordenamiento y mapa se ubica entre 10^3 y 10^4 elementos. Para entradas chicas, el ordenamiento es mas rapido por sus mejores constantes.

---

## 5. Ejercicio 4: Ordenamiento

### 5.1 Consigna

Ordenar una lista de numeros utilizando Bubble Sort. No utilizar la funcion de ordenamiento propia del lenguaje; comparar luego usando esta funcion. Medir tiempos e indicar la complejidad algoritmica.

### 5.2 Codigo: `Ej4.go`

```go
package main

import (
    "fmt"
    "slices"
    "time"
)

// Bubble Sort clasico O(n^2)
func bubbleSort(datos []int) {
    for i := range datos {
        for j := 0; j < len(datos)-1-i; j++ {
            if datos[j+1] < datos[j] {
                datos[j+1], datos[j] = datos[j], datos[j+1]
            }
        }
    }
}

// Bubble Sort optimizado O(n) mejor caso, O(n^2) peor
func bubbleSortOptimizado(datos []int) {
    n := len(datos)
    for n > 1 {
        ultimoCambio := 0
        for j := 1; j < n; j++ {
            if datos[j] < datos[j-1] {
                datos[j], datos[j-1] = datos[j-1], datos[j]
                ultimoCambio = j
            }
        }
        n = ultimoCambio
    }
}

// Insertion Sort O(n) mejor caso, O(n^2) peor
func insertionSortCorrimiento(datos []int) {
    for i := 1; i < len(datos); i++ {
        actual := datos[i]
        j := i - 1
        for j >= 0 && datos[j] > actual {
            datos[j+1] = datos[j]
            j--
        }
        datos[j+1] = actual
    }
}

func estaOrdenado(datos []int) bool {
    for i := 1; i < len(datos); i++ {
        if datos[i] < datos[i-1] {
            return false
        }
    }
    return true
}

func medirOrdenamiento(nombre string, original []int, algoritmo func([]int)) {
    datos := slices.Clone(original)
    inicio := time.Now()
    algoritmo(datos)
    fmt.Printf("%-24s ordenado: %t, tiempo: %v\n", nombre, estaOrdenado(datos), time.Since(inicio))
}

func main() {
    datos := []int{42, 7, 19, 3, 88, 12, 1, 54, 27, 6}
    fmt.Println("Lista original:", datos)

    medirOrdenamiento("Bubble Sort", datos, bubbleSort)
    medirOrdenamiento("Bubble optimizado", datos, bubbleSortOptimizado)
    medirOrdenamiento("Insertion Sort", datos, insertionSortCorrimiento)

    copia := slices.Clone(datos)
    inicio := time.Now()
    slices.Sort(copia)
    fmt.Printf("slices.Sort            ordenado: %t, tiempo: %v\n",
        estaOrdenado(copia), time.Since(inicio))
    fmt.Println("Complejidad: Bubble clasico O(n^2) en mejor y peor caso; " +
        "Bubble optimizado O(n) mejor y O(n^2) peor; " +
        "Insertion O(n) mejor y O(n^2) peor; slices.Sort O(n log n).")
```

### 5.3 Explicacion del codigo

**Bubble Sort clasico:**
- En cada pasada del bucle interno, el mayor elemento restante "burbujea" hacia su posicion final.
- El limite del bucle interno se reduce en `i` porque las ultimas `i` posiciones ya estan ordenadas.
- Intercambio multiple: Go permite `a, b = b, sin variable temporal`.
- **Siempre ejecuta O(n^2) comparaciones**, sin importar el estado del arreglo.

**Bubble Sort optimizado:**
- Registra `ultimoCambio`: la ultima posicion donde hubo intercambio.
- Si una pasada completa no produce intercambios, `ultimoCambio` queda en 0 y el algoritmo termina.
- Si el arreglo ya esta ordenado, solo hace 1 pasada: mejor caso O(n).
- Si hay pocos elementos desordenados pero lejanos, la optimizacion no ayuda (el desorden no es local).

**Insertion Sort (con corrimiento):**
- Mantiene la invariante de que `[0, i-1]` esta ordenado.
- Toma el primer elemento de la zona desordenada y lo desplaza hacia atras hasta su posicion correcta.
- Usa `j >= 0 && datos[j] > actual` con evaluacion en cortocircuito para evitar acceso fuera de rango.
- La variante con corrimiento es ~15% mas rapida que con intercambios porque realiza menos escrituras en memoria.
- **Es adaptativo:** sobre datos casi ordenados es dramaticamente mas rapido.

### 5.4 Salida en consola

```
Lista original: [42 7 19 3 88 12 1 54 27 6]
Bubble Sort              ordenado: true, tiempo: 0s
Bubble optimizado        ordenado: true, tiempo: 0s
Insertion Sort           ordenado: true, tiempo: 0s
slices.Sort              ordenado: true, tiempo: 0s
Complejidad: Bubble clasico O(n^2) en mejor y peor caso; Bubble optimizado O(n) mejor y O(n^2) peor; Insertion O(n) mejor y O(n^2) peor; slices.Sort O(n log n).
```

### 5.5 Analisis teorico

| Algoritmo | Mejor | Promedio | Peor | Espacio | Estable |
|-----------|-------|----------|------|---------|---------|
| Bubble clasico | O(n^2) | O(n^2) | O(n^2) | O(1) | Si |
| Bubble optimizado | O(n) | O(n^2) | O(n^2) | O(1) | Si |
| Insertion Sort | O(n) | O(n^2) | O(n^2) | O(1) | Si |
| slices.Sort (pdqsort) | O(n) | O(n log n) | O(n log n) | O(log n) | No |

**Observaciones del PDF:**
- **Bubble clasico fue mas lento sobre datos aleatorios que inversos**, pese a realizar identico numero de comparaciones. El predictor de saltos del procesador predice perfectamente con datos inversos (la condicion siempre se cumple) pero falla con aleatorios.
- **Insertion sort fue 25 veces mas rapido** sobre datos casi ordenados que sobre aleatorios, ejecutando el mismo codigo.
- **`slices.Sort` fue ~31 veces mas rapido** que la mejor version manual para n=10^4, usando pdqsort (hibrido que detecta secuencias ordenadas).
- **`slices.Sort` no es estable:** no preserva el orden relativo de elementos equivalentes. Para enteros aislados esto es indistinguible, pero seria relevante con estructuras.
- **La variante con corrimiento supero consistentemente** a la de intercambios (2%-17%), demostrando que dos algoritmos de identica complejidad pueden tener rendimientos distintos.

---

## 6. Conclusiones

1. **La complejidad asintotica describe el crecimiento, no el rendimiento absoluto.** La fuerza bruta O(n^2) puede ser mas rapida que O(n) cuando la terminacion temprana es probable y n es chico.

2. **Las optimizaciones sobre constantes no alteran el orden, y pueden ser contraproducentes.** Reservar capacidad del mapa mejoro el peor caso pero degrado el caso favorable por un factor de 29.

3. **La forma de los datos puede importar tanto como el algoritmo.** Insertion sort fue 25x mas rapido sobre datos casi ordenados que sobre aleatorios.

4. **El hardware produce efectos que la teoria no contempla.** El predictor de saltos y la jerarquia de caches alteran el rendimiento real.

5. **El compilador ya realiza micro-optimizaciones.** La sustitucion de `% 2` por `& 1` no produjo efecto. Optimizar manualmente lo que el compilador resuelve agrega complejidad sin contrapartida.

6. **Medir es imprescindible.** Varias hipotesis planteadas antes de las mediciones resultaron incorrectas.

---

## 7. Como ejecutar

### Requisitos previos

- Tener [Go](https://go.dev/dl/) instalado (version 1.26.5 o superior).
- Estar en la carpeta del proyecto.

### Paso a paso

```powershell
# Navegar a la carpeta del practico
cd "C:\Users\manue\OneDrive\Documentos\Programacion\PEF\practico1"

# Ejecutar cada ejercicio individualmente
go run Ej1.go    # Busqueda lineal
go run Ej2.go    # Conteo de pares
go run Ej3.go    # Deteccion de repetidos
go run Ej4.go    # Ordenamiento
```

### Verificacion estatica

```powershell
# Compilar todos los archivos sin generar ejecutable (verificar errores)
go build ./...

# Analisis estatico del codigo
go vet ./...
```

### Estructura del proyecto

```
practico1/
  |- go.mod                        (modulo Go: practico1, Go 1.26.5)
  |- Ej1.go                        (Busqueda lineal)
  |- Ej2.go                        (Conteo de numeros pares)
  |- Ej3.go                        (Deteccion de repetidos)
  |- Ej4.go                        (Ordenamiento: Bubble, Insertion)
  |- PracticoPEF1_...pdf           (Informe completo del practico)
  |- README.md                     (Este archivo)
```

### Nota sobre los ejecutables

Cada archivo `EjN.go` es un paquete `main` independiente. Go permite ejecutar archivos sueltos con `go run`, pero no se pueden compilar todos juntos en un solo modulo porque cada uno tiene su propia funcion `main`. Para ejecutar todos de una vez:

```powershell
# Ejecutar cada uno secuencialmente
go run Ej1.go; go run Ej2.go; go run Ej3.go; go run Ej4.go
```

---

## go.mod

```
module practico1

go 1.26.5
```
