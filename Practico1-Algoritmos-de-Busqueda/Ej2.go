package main

import (
	"fmt"
	"time"
)

/*
Dada una lista de números enteros, contar cuántos números pares contiene. Probar con diferentes
algoritmos, medir tiempos e indicar la complejidad algorítmica.
*/

func contarParesMod(datos []int) int {
	var contador int
	for _, v := range datos {
		if v%2 == 0 {
			contador++
		}
	}
	return contador
}

func contarParesBit(datos []int) int {
	var contador int
	for _, v := range datos {
		if v&1 == 0 {
			contador++
		}
	}
	return contador
}

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
