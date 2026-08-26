package main

import (
	"fmt"
	"runtime"
	"time"
)

func cuadradosLista(n int) []int {
	lista := make([]int, n)
	for i := 0; i < n; i++ {
		lista[i] = i * i
	}
	return lista
}

func cuadradosGenerador(n int) <-chan int {
	canal := make(chan int)
	go func() {
		defer close(canal)
		for i := 0; i < n; i++ {
			canal <- i * i
		}
	}()
	return canal
}

func medirMemoria() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Alloc
}

func main() {
	n := 1000000

	runtime.GC()
	antes := medirMemoria()
	inicio := time.Now()
	lista := cuadradosLista(n)
	tiempoLista := time.Since(inicio)
	despues := medirMemoria()
	fmt.Println("Lista: primeros 5 valores:", lista[:5])
	fmt.Println("Lista: tiempo:", tiempoLista)
	fmt.Println("Lista: memoria usada (bytes):", despues-antes)

	fmt.Println("---")

	runtime.GC()
	antes = medirMemoria()
	inicio = time.Now()
	contador := 0
	var primeros []int
	for valor := range cuadradosGenerador(n) {
		contador++
		if contador <= 5 {
			primeros = append(primeros, valor)
		}
	}
	tiempoGenerador := time.Since(inicio)
	despues = medirMemoria()
	fmt.Println("Generador: primeros 5 valores:", primeros)
	fmt.Println("Generador: tiempo:", tiempoGenerador)
	fmt.Println("Generador: memoria usada (bytes):", despues-antes)
}
