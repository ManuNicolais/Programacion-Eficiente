package main

import "fmt"

func busquedaLineal(lista []int, buscar int) int {
	for i, valor := range lista {
		if valor == buscar {
			return i
		}
	}
	return -1
}

func busquedaBinaria(lista []int, buscar int) int {
	inicio := 0
	fin := len(lista) - 1

	for inicio <= fin {
		medio := (inicio + fin) / 2
		if lista[medio] == buscar {
			return medio
		} else if lista[medio] < buscar {
			inicio = medio + 1
		} else {
			fin = medio - 1
		}
	}
	return -1
}

func main() {
	satelites := []int{101, 105, 109, 115, 120, 127, 132, 140, 148, 155, 160, 172, 180, 191}
	buscar := 155

	posicionLineal := busquedaLineal(satelites, buscar)
	posicionBinaria := busquedaBinaria(satelites, buscar)

	fmt.Println("Buscando:", buscar)
	fmt.Println("Busqueda lineal, posicion:", posicionLineal)
	fmt.Println("Busqueda binaria, posicion:", posicionBinaria)
}
