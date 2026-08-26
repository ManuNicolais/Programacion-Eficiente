package main

import (
	"fmt"
	"time"
)

func generarClientes(n int) []string {
	clientes := make([]string, n)
	for i := 0; i < n; i++ {
		clientes[i] = fmt.Sprintf("CLI%06d", i)
	}
	return clientes
}

func generarCompradores(n int, maxClientes int) []string {
	compradores := make([]string, n)
	for i := 0; i < n; i++ {
		id := (i * 3) % (maxClientes * 2)
		compradores[i] = fmt.Sprintf("CLI%06d", id)
	}
	return compradores
}

func compradoresConList(clientes []string, compradores []string) []string {
	var resultado []string
	for _, cliente := range clientes {
		for _, comprador := range compradores {
			if cliente == comprador {
				resultado = append(resultado, cliente)
				break
			}
		}
	}
	return resultado
}

func compradoresConSet(clientes []string, compradores []string) []string {
	set := make(map[string]bool)
	for _, comprador := range compradores {
		set[comprador] = true
	}

	var resultado []string
	for _, cliente := range clientes {
		if set[cliente] {
			resultado = append(resultado, cliente)
		}
	}
	return resultado
}

func main() {
	tamanios := []int{10000, 100000, 1000000}

	for _, n := range tamanios {
		clientes := generarClientes(n)
		compradores := generarCompradores(n, n)

		fmt.Println("Tamanio:", n)

		if n <= 100000 {
			inicio := time.Now()
			resultadoList := compradoresConList(clientes, compradores)
			tiempoList := time.Since(inicio)
			fmt.Println("Compradores encontrados (list):", len(resultadoList))
			fmt.Println("Tiempo con list:", tiempoList)
		} else {
			fmt.Println("Version list omitida en este tamanio: tarda demasiado (O(n*m))")
		}

		inicio := time.Now()
		resultadoSet := compradoresConSet(clientes, compradores)
		tiempoSet := time.Since(inicio)
		fmt.Println("Compradores encontrados (set):", len(resultadoSet))
		fmt.Println("Tiempo con set:", tiempoSet)
		fmt.Println("---")
	}
}
