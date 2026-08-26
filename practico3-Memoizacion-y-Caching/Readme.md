# Practico 3 - Memoizacion y Caching

## Conceptos Teoricos

### Memoizacion
La **memoizacion** es una tecnica de optimizacion que consiste en almacenar los resultados de llamadas a funciones costosas para que, si se solicitan los mismos parametros nuevamente, se devuelva el resultado guardado en lugar de recalcularlo. Es una forma de **programacion dinamica top-down**.

### Caching
El **caching** es el acto de almacenar datos en una memoria temporal (cache) para reducir el tiempo de acceso futuro. La memoizacion es un tipo especifico de caching aplicado a funciones recursivas.

### Diferencia clave
- **Sin memoizacion**: una funcion recursiva puede re-calcular el mismo estado muchas veces, llevando a complejidad exponencial O(2^n).
- **Con memoizacion**: cada subproblema se resuelve una sola vez, reduciendo la complejidad a O(n * m) donde n y m son las dimensiones del espacio de estados.

---

## Ejercicio 1: Caminos en una Cuadricula

### Problema
Dada una cuadricula de tamanho m x n, encontrar la cantidad de formas de llegar desde la esquina superior izquierda (0,0) hasta la esquina inferior derecha (m-1, n-1), moviendose solo hacia la derecha o hacia abajo.

### Enfoque
Se utiliza recursion con memoizacion. En cada celda (x, y) se exploran dos movimientos: ir a la derecha (x+1, y) o ir hacia abajo (x, y+1). La clave del memo es el par de coordenadas `"x,y"`.

### Codigo fuente: `Ej1.go`

```go
func resolverCaminos(m, n, x, y int, memo map[string]int) int {
    if x == m-1 && y == n-1 {
        return 1
    }
    if x >= m || y >= n {
        return 0
    }
    clave := fmt.Sprintf("%d,%d", x, y)
    if valor, existe := memo[clave]; existe {
        return valor
    }
    caminosDerecha := resolverCaminos(m, n, x+1, y, memo)
    caminosAbajo := resolverCaminos(m, n, x, y+1, memo)
    memo[clave] = caminosDerecha + caminosAbajo
    return memo[clave]
}
```

### Salida en consola

```
Cuadricula: 3x3
Inicio: (0,0) | Meta: (m-1,n-1)
En cada posicion se prueban dos movimientos: derecha y abajo.
Calculando caminos y guardando los resultados ya obtenidos...
Desde (2,1): derecha=0, abajo=1, total=1
Desde (2,0): derecha=0, abajo=1, total=1
Desde (1,2): derecha=1, abajo=0, total=1
Desde (1,1): derecha=1, abajo=1, total=2
Desde (1,0): derecha=1, abajo=2, total=3
Desde (0,2): derecha=1, abajo=0, total=1
Desde (0,1): derecha=2, abajo=1, total=3
Desde (0,0): derecha=3, abajo=3, total=6
Estados guardados en memoria: 8
Cantidad de formas de atravesar una cuadricula 3x3: 6
```

### Analisis
Para una cuadricula 3x3 existen **6 caminos** validos. La memoizacion almacena **8 estados** (las 8 celdas intermedias). Sin memoizacion, la recursion exploraria muchos mas nodos redundantes.

### Complejidad
- **Sin memoizacion**: O(2^(m+n)) - exponencial
- **Con memoizacion**: O(m * n) - lineal en el numero de celdas

---

## Ejercicio 2: Problema del Subconjunto con Suma Objetivo

### Problema
Dada una lista de enteros positivos y un numero objetivo, determinar si existe un subconjunto cuya suma sea exactamente el objetivo.

### Enfoque
Para cada numero se consideran dos opciones: tomarlo (restandolo del objetivo) o descartarlo (avanzando al siguiente indice). La clave del memo es `"indice,objetivo"`.

### Codigo fuente: `Ej2.go`

```go
func sumaSubconjunto(numeros []int, indice, objetivo int, memo map[string]bool) bool {
    if objetivo == 0 {
        return true
    }
    if indice >= len(numeros) || objetivo < 0 {
        return false
    }
    clave := fmt.Sprintf("%d,%d", indice, objetivo)
    if valor, existe := memo[clave]; existe {
        return valor
    }
    tomar := sumaSubconjunto(numeros, indice+1, objetivo-numeros[indice], memo)
    noTomar := sumaSubconjunto(numeros, indice+1, objetivo, memo)
    memo[clave] = tomar || noTomar
    return memo[clave]
}
```

### Salida en consola

```
Numeros disponibles: [3 34 4 12 5 2]
Objetivo: 9
Para cada numero se prueban dos opciones: tomarlo o descartarlo.
Calculando las posibilidades y guardando los estados repetidos...
Indice 5, objetivo restante 2: tomar=true, no tomar=false -> resultado=true
Indice 4, objetivo restante 2: tomar=false, no tomar=true -> resultado=true
Indice 3, objetivo restante 2: tomar=false, no tomar=true -> resultado=true
Indice 5, objetivo restante 1: tomar=false, no tomar=false -> resultado=false
Indice 5, objetivo restante 6: tomar=false, no tomar=false -> resultado=false
Indice 4, objetivo restante 6: tomar=false, no tomar=false -> resultado=false
Indice 3, objetivo restante 6: tomar=false, no tomar=false -> resultado=false
Indice 2, objetivo restante 6: tomar=true, no tomar=false -> resultado=true
Indice 1, objetivo restante 6: tomar=false, no tomar=true -> resultado=true
Indice 5, objetivo restante 5: tomar=false, no tomar=false -> resultado=false
Indice 4, objetivo restante 5: tomar=true, no tomar=false -> resultado=true
Indice 3, objetivo restante 5: tomar=false, no tomar=true -> resultado=true
Indice 5, objetivo restante 4: tomar=false, no tomar=false -> resultado=false
Indice 5, objetivo restante 9: tomar=false, no tomar=false -> resultado=false
Indice 4, objetivo restante 9: tomar=false, no tomar=false -> resultado=false
Indice 3, objetivo restante 9: tomar=false, no tomar=false -> resultado=false
Indice 2, objetivo restante 9: tomar=true, no tomar=false -> resultado=true
Indice 1, objetivo restante 9: tomar=false, no tomar=true -> resultado=true
Indice 0, objetivo restante 9: tomar=true, no tomar=true -> resultado=true
Estados guardados en memoria: 19
¿Existe subconjunto que sume 9? true
```

### Analisis
Con la lista `[3, 34, 4, 12, 5, 2]` y objetivo 9, se encuentra que **si existe** un subconjunto (por ejemplo: 4 + 5 = 9). La memoizacion almacena **19 estados** unicos, evitando recalcular pares `(indice, objetivo)` que ya fueron evaluados.

### Complejidad
- **Sin memoizacion**: O(2^n) - exponencial en la cantidad de numeros
- **Con memoizacion**: O(n * objetivo) - pseudo-polinomica

---

## Ejercicio 3: Word Break - Formar un String con Palabras del Diccionario

### Problema
Dado un string S y una lista de palabras validas, determinar si S se puede formar usando solo palabras de la lista, reutilizandolas si es necesario.

### Enfoque
Se verifica si el string actual comienza con alguna palabra del diccionario. Si es asi, se "corta" esa palabra y se analiza el resto recursivamente. La clave del memo es el string restante.

### Codigo fuente: `Ej3.go`

```go
func sePuedeFormar(s string, diccionario []string, memo map[string]bool) bool {
    if s == "" {
        return true
    }
    if valor, existe := memo[s]; existe {
        return valor
    }
    for _, palabra := range diccionario {
        if strings.HasPrefix(s, palabra) {
            restoDelString := s[len(palabra):]
            if sePuedeFormar(restoDelString, diccionario, memo) {
                memo[s] = true
                return true
            }
        }
    }
    memo[s] = false
    return false
}
```

### Salida en consola

```
Texto a formar: 'programacioneficiente'
Diccionario: [programacion eficiente programa cion]
Se buscan palabras que sean prefijo del texto restante.
Cada prefijo encontrado se descarta y se analiza recursivamente el resto...
Se prueba la palabra 'programacion' en 'programacioneficiente' -> queda 'eficiente'
Se prueba la palabra 'eficiente' en 'eficiente' -> queda ''
La parte 'eficiente' se pudo formar usando 'eficiente'
La parte 'programacioneficiente' se pudo formar usando 'programacion'
Fragmentos evaluados y guardados en memoria: 2
¿Se puede formar 'programacioneficiente'? true
```

### Analisis
El string `"programacioneficiente"` se forma con `"programacion"` + `"eficiente"`. Solo se evaluaron **2 fragmentos** en el memo. Sin memoizacion, si hubiera prefijos compartidos entre ramas, se repetirian calculos.

### Complejidad
- **Sin memoizacion**: O(2^len(S)) en el peor caso
- **Con memoizacion**: O(len(S)^2) - cada substring unico se evalua una vez

---

## Ejercicio 4: Combinaciones de Monedas (Coin Change - Contar Formas)

### Problema
Dado un valor total y una lista de monedas, devolver cuantas combinaciones distintas de monedas suman ese total.

### Enfoque
Para cada tipo de moneda hay dos opciones: usarla (manteniendo el indice para poder reusarla) o pasar a la siguiente denominacion. La clave del memo es `"indice,total"`.

### Codigo fuente: `Ej4.go`

```go
func combinacionesMonedas(monedas []int, indice, total int, memo map[string]int) int {
    if total == 0 {
        return 1
    }
    if total < 0 || indice >= len(monedas) {
        return 0
    }
    clave := fmt.Sprintf("%d,%d", indice, total)
    if valor, existe := memo[clave]; existe {
        return valor
    }
    tomar := combinacionesMonedas(monedas, indice, total-monedas[indice], memo)
    noTomar := combinacionesMonedas(monedas, indice+1, total, memo)
    memo[clave] = tomar + noTomar
    return memo[clave]
}
```

### Salida en consola

```
Monedas disponibles: [1 2 5]
Total a formar: 5
Se cuentan combinaciones sin importar el orden de las monedas.
En cada paso se decide usar la moneda actual o pasar a la siguiente...
Indice 2, total restante 1: usando moneda=0, sin usarla=0 -> total=0
Indice 1, total restante 1: usando moneda=0, sin usarla=0 -> total=0
Indice 0, total restante 1: usando moneda=1, sin usarla=0 -> total=1
Indice 2, total restante 2: usando moneda=0, sin usarla=0 -> total=0
Indice 1, total restante 2: usando moneda=1, sin usarla=0 -> total=1
Indice 0, total restante 2: usando moneda=1, sin usarla=1 -> total=2
Indice 2, total restante 3: usando moneda=0, sin usarla=0 -> total=0
Indice 1, total restante 3: usando moneda=0, sin usarla=0 -> total=0
Indice 0, total restante 3: usando moneda=2, sin usarla=0 -> total=2
Indice 2, total restante 4: usando moneda=0, sin usarla=0 -> total=0
Indice 1, total restante 4: usando moneda=1, sin usarla=0 -> total=1
Indice 0, total restante 4: usando moneda=2, sin usarla=1 -> total=3
Indice 2, total restante 5: usando moneda=1, sin usarla=0 -> total=1
Indice 1, total restante 5: usando moneda=0, sin usarla=1 -> total=1
Indice 0, total restante 5: usando moneda=3, sin usarla=1 -> total=4
Estados guardados en memoria: 15
Cantidad de combinaciones para formar 5: 4
```

### Analisis
Con monedas `[1, 2, 5]` y total 5, existen **4 combinaciones**: {1+1+1+1+1}, {1+1+1+2}, {1+2+2}, {5}. La memoizacion almacena **15 estados**, uno por cada par unico de `(indice, total_restante)`.

### Complejidad
- **Sin memoizacion**: O(2^n) o peor, dependiendo de las monedas
- **Con memoizacion**: O(n * total) - pseudo-polinomica

---

## go.mod

```
module practico3

go 1.26.5
```

---

## Conclusion

Los cuatro ejercicios demuestran como la memoizacion transforma problemas con complejidad exponencial en problemas con complejidad polinomica o pseudo-polinomica. El patron comun es:

1. Identificar los **subproblemas** que se repiten en la recursion.
2. Definir una **clave** unica para cada estado (par de parametros).
3. **Almacenar** el resultado antes de retornar.
4. **Consultar** el cache antes de recalcular.

La memoizacion es fundamental en algoritmos de programacion dinamica y se aplica en problemas como caminos en grillas, subset sum, word break, coin change, y muchos otros.
