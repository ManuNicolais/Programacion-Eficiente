/*
Dado un valor total y una lista de monedas, escribí una funcion que devuelva cuantas
combinaciones distintas de monedas suman ese total.
*/

package main

import (
	"fmt"
)

func combinacionesMonedas(monedas []int, indice, total int, memo map[string]int) int {
	// Si el total es 0, encontramos 1 forma exacta de dar el cambio.
	if total == 0 {
		return 1
	}
	// Si el total es negativo o ya no hay monedas, es un callejón sin salida (0 formas).
	if total < 0 || indice >= len(monedas) {
		return 0
	}

	clave := fmt.Sprintf("%d,%d", indice, total)
	if valor, existe := memo[clave]; existe {
		return valor
	}

	// Decisión:
	// 1. Tomar la moneda (mantenemos el índice para poder re-utilizarla)
	tomar := combinacionesMonedas(monedas, indice, total-monedas[indice], memo)

	// 2. No tomar la moneda (avanzamos al siguiente tipo de moneda)
	noTomar := combinacionesMonedas(monedas, indice+1, total, memo)

	// La cantidad total de formas es la suma de ambos caminos
	memo[clave] = tomar + noTomar
	fmt.Printf("Indice %d, total restante %d: usando moneda=%d, sin usarla=%d -> total=%d\n", indice, total, tomar, noTomar, memo[clave])
	return memo[clave]
}

func main() {
	monedas := []int{1, 2, 5}
	total := 5
	fmt.Printf("Monedas disponibles: %v\n", monedas)
	fmt.Printf("Total a formar: %d\n", total)
	fmt.Println("Se cuentan combinaciones sin importar el orden de las monedas.")
	fmt.Println("En cada paso se decide usar la moneda actual o pasar a la siguiente...")
	memo := make(map[string]int)

	formas := combinacionesMonedas(monedas, 0, total, memo)
	fmt.Printf("Estados guardados en memoria: %d\n", len(memo))
	fmt.Printf("Cantidad de combinaciones para formar %d: %d\n", total, formas)
}
