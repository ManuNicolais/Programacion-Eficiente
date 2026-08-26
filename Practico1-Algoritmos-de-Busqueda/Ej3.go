package main

import (
	"fmt"
	"slices"
	"time"
)

/*
Desarrollar un algoritmo que determine si existen números repetidos dentro de una lista. No
utilizar estructuras como set o HashSet; comparar luego usando estas estructuras. Medir tiempos
e indicar la complejidad algorítmica.
*/

// Versión 1: fuerza bruta
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

// Versión 2: ordenamiento previo
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

// Versión 3: mediante tabla de dispersión
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
	datos[len(datos)-1] = datos[0]

	medirRepetidos("Fuerza bruta", datos, repetidosFuerzaBruta)
	medirRepetidos("Ordenamiento previo", datos, repetidosOrdenando)
	medirRepetidos("Mapa", datos, repetidosMapa)

	copia := slices.Clone(datos)
	slices.Sort(copia)
	fmt.Printf("Comparacion con slices.Sort: %t\n", repetidosOrdenando(datos) == slices.Contains(copia, datos[0]))
	fmt.Println("Complejidad: fuerza bruta O(n^2) y O(1) espacio; ordenamiento O(n log n) y O(n) espacio; mapa O(n) promedio y O(n) espacio.")
}
