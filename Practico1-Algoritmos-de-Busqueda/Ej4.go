package main

import (
	"fmt"
	"slices"
	"time"
)

/*
Ordenar una lista de números utilizando Bubble Sort. No utilizar la función de ordenamiento
propia del lenguaje; comparar luego usando esta función. Medir tiempos e indicar la complejidad
algorítmica.
*/

// Bubble Sort clásico
func bubbleSort(datos []int) {
	for i := range datos {
		for j := 0; j < len(datos)-1-i; j++ {
			if datos[j+1] < datos[j] {
				datos[j+1], datos[j] = datos[j], datos[j+1]
			}
		}
	}
}

// Bubble Sort optimizado
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

// Insertion Sort
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
	fmt.Printf("slices.Sort            ordenado: %t, tiempo: %v\n", estaOrdenado(copia), time.Since(inicio))
	fmt.Println("Complejidad: Bubble clasico O(n^2) en mejor y peor caso; Bubble optimizado O(n) mejor y O(n^2) peor; Insertion O(n) mejor y O(n^2) peor; slices.Sort O(n log n).")
}
