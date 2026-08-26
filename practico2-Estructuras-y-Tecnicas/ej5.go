package main

import "fmt"

func costoRecursivo(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return costoRecursivo(n-1) + costoRecursivo(n-2)
}

func costoMemo(n int, memo map[int]int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if valor, existe := memo[n]; existe {
		return valor
	}
	resultado := costoMemo(n-1, memo) + costoMemo(n-2, memo)
	memo[n] = resultado
	return resultado
}

func main() {
	n := 5

	resultadoRecursivo := costoRecursivo(n)
	memo := make(map[int]int)
	resultadoMemo := costoMemo(n, memo)

	fmt.Println("costo recursivo(", n, ") =", resultadoRecursivo)
	fmt.Println("costo memoizado(", n, ") =", resultadoMemo)
}
