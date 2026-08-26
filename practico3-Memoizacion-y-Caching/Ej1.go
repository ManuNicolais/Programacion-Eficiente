/*
Dada una cuadricula de tamaño m x n, escribí una funcion que devuelva
la cantidad de formas de llegar desde la esquina superior izquierda (0,0) hasta la
esquina inferior derecha (m-1, n-1), moviendo solo hacia la derecha o hacia abajo.
*/

package main

import (
	"fmt"
)

// resolverCaminos es nuestra función recursiva con memoización
func resolverCaminos(m, n, x, y int, memo map[string]int) int {
	// Si llegamos a la meta (m-1, n-1), hay 1 forma válida de haber llegado.
	if x == m-1 && y == n-1 {
		return 1
	}
	// Si nos salimos de los límites, no hay caminos válidos (0 formas).
	if x >= m || y >= n {
		return 0
	}

	// 1. Clave asociada a los argumentos: creamos un string "x,y"
	clave := fmt.Sprintf("%d,%d", x, y)

	// 2. Si ya lo calculamos antes, retornamos el resultado almacenado
	if valor, existe := memo[clave]; existe {
		return valor
	}

	// 3. Calculamos: ir a la derecha + ir hacia abajo
	caminosDerecha := resolverCaminos(m, n, x+1, y, memo)
	caminosAbajo := resolverCaminos(m, n, x, y+1, memo)

	// Guardamos en nuestro caché (memoización) antes de retornar
	memo[clave] = caminosDerecha + caminosAbajo
	fmt.Printf("Desde (%d,%d): derecha=%d, abajo=%d, total=%d\n", x, y, caminosDerecha, caminosAbajo, memo[clave])
	return memo[clave]
}

func main() {
	m, n := 3, 3 // Cuadrícula de 3x3
	fmt.Printf("Cuadricula: %dx%d\n", m, n)
	fmt.Println("Inicio: (0,0) | Meta: (m-1,n-1)")
	fmt.Println("En cada posicion se prueban dos movimientos: derecha y abajo.")
	fmt.Println("Calculando caminos y guardando los resultados ya obtenidos...")
	// Inicializamos el caché (un mapa de string a int)
	memo := make(map[string]int)

	resultado := resolverCaminos(m, n, 0, 0, memo)
	fmt.Printf("Estados guardados en memoria: %d\n", len(memo))
	fmt.Printf("Cantidad de formas de atravesar una cuadrícula %dx%d: %d\n", m, n, resultado)
}
