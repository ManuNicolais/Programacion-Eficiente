package main

import (
	"fmt"
	"slices"
	"time"
)

/*
Desarrollar un algoritmo que reciba una lista de números enteros y un número a buscar. El
algoritmo debe indicar si el número existe dentro de la lista. No utilizar funciones propias del
lenguaje como contains() o in; comparar luego usando estas funciones. Medir tiempos e indicar
la complejidad algorítmica.
*/

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
