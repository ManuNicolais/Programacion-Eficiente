/*
Dada una lista de enteros positivos y un numero objetivo, esribí una funcion que
determine si existe un subconjunto cuya suma sea exactamente el objetivo
*/

package main

import (
	"fmt"
)

func sumaSubconjunto(numeros []int, indice, objetivo int, memo map[string]bool) bool {
	// Casos base
	if objetivo == 0 {
		return true // Encontramos un subconjunto exacto
	}
	if indice >= len(numeros) || objetivo < 0 {
		return false // Nos pasamos o nos quedamos sin números
	}

	// Clave única basada en el índice actual y el objetivo restante
	clave := fmt.Sprintf("%d,%d", indice, objetivo)

	if valor, existe := memo[clave]; existe {
		return valor
	}

	// Tomar la decisión (¿Tomar o no tomar?)
	tomar := sumaSubconjunto(numeros, indice+1, objetivo-numeros[indice], memo)
	noTomar := sumaSubconjunto(numeros, indice+1, objetivo, memo)

	// Si cualquiera de las dos rutas es exitosa, guardamos true
	memo[clave] = tomar || noTomar
	fmt.Printf("Indice %d, objetivo restante %d: tomar=%t, no tomar=%t -> resultado=%t\n", indice, objetivo, tomar, noTomar, memo[clave])
	return memo[clave]
}

func main() {
	lista := []int{3, 34, 4, 12, 5, 2}
	objetivo := 9
	fmt.Printf("Numeros disponibles: %v\n", lista)
	fmt.Printf("Objetivo: %d\n", objetivo)
	fmt.Println("Para cada numero se prueban dos opciones: tomarlo o descartarlo.")
	fmt.Println("Calculando las posibilidades y guardando los estados repetidos...")
	memo := make(map[string]bool)

	existe := sumaSubconjunto(lista, 0, objetivo, memo)
	fmt.Printf("Estados guardados en memoria: %d\n", len(memo))
	fmt.Printf("¿Existe subconjunto que sume %d? %t\n", objetivo, existe)
}
