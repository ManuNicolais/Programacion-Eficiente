package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generarUsuarios(n int) []string {
	usuarios := make([]string, n)
	for i := 0; i < n; i++ {
		usuarios[i] = fmt.Sprintf("USR%06d", i)
	}
	return usuarios
}

func generarIntentos(n int, maxUsuarios int) []string {
	intentos := make([]string, n)
	for i := 0; i < n; i++ {
		id := rand.Intn(maxUsuarios * 2)
		intentos[i] = fmt.Sprintf("USR%06d", id)
	}
	return intentos
}

func autorizadosConList(usuarios []string, intentos []string) []string {
	var autorizados []string
	for _, intento := range intentos {
		for _, usuario := range usuarios {
			if intento == usuario {
				autorizados = append(autorizados, intento)
				break
			}
		}
	}
	return autorizados
}

func autorizadosConSet(usuarios []string, intentos []string) []string {
	set := make(map[string]bool)
	for _, usuario := range usuarios {
		set[usuario] = true
	}

	var autorizados []string
	for _, intento := range intentos {
		if set[intento] {
			autorizados = append(autorizados, intento)
		}
	}
	return autorizados
}

func main() {
	usuarios := generarUsuarios(100000)
	intentos := generarIntentos(20000, 100000)

	inicio := time.Now()
	resultadoList := autorizadosConList(usuarios, intentos)
	tiempoList := time.Since(inicio)

	inicio = time.Now()
	resultadoSet := autorizadosConSet(usuarios, intentos)
	tiempoSet := time.Since(inicio)

	fmt.Println("Cantidad de usuarios:", len(usuarios))
	fmt.Println("Cantidad de intentos:", len(intentos))
	fmt.Println("Autorizados encontrados (list):", len(resultadoList))
	fmt.Println("Autorizados encontrados (set):", len(resultadoSet))
	fmt.Println("Tiempo con list:", tiempoList)
	fmt.Println("Tiempo con set:", tiempoSet)
}
