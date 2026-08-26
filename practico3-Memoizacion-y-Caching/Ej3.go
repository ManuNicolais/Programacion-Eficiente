/*
Dado un string S y una lista de palabras validas, determina si S se puede
formar usando solo palabras de la lista, reutilizandolas si es necesario.
*/

package main

import (
	"fmt"
	"strings"
)

func sePuedeFormar(s string, diccionario []string, memo map[string]bool) bool {
	// Si el string quedó vacío, logramos formarlo completamente
	if s == "" {
		return true
	}

	// Verificamos en el caché si este string ya fue analizado
	if valor, existe := memo[s]; existe {
		return valor
	}

	// Probamos con cada palabra del diccionario
	for _, palabra := range diccionario {
		// Si el string actual arranca con la palabra de nuestro diccionario...
		if strings.HasPrefix(s, palabra) {
			// Cortamos el string y analizamos el resto
			restoDelString := s[len(palabra):]
			fmt.Printf("Se prueba la palabra '%s' en '%s' -> queda '%s'\n", palabra, s, restoDelString)
			if sePuedeFormar(restoDelString, diccionario, memo) {
				memo[s] = true
				fmt.Printf("La parte '%s' se pudo formar usando '%s'\n", s, palabra)
				return true
			}
		}
	}

	// Si probamos todas las palabras y ninguna funcionó
	memo[s] = false
	return false
}

func main() {
	s := "programacioneficiente"
	diccionario := []string{"programacion", "eficiente", "programa", "cion"}
	fmt.Printf("Texto a formar: '%s'\n", s)
	fmt.Printf("Diccionario: %v\n", diccionario)
	fmt.Println("Se buscan palabras que sean prefijo del texto restante.")
	fmt.Println("Cada prefijo encontrado se descarta y se analiza recursivamente el resto...")
	memo := make(map[string]bool)

	resultado := sePuedeFormar(s, diccionario, memo)
	fmt.Printf("Fragmentos evaluados y guardados en memoria: %d\n", len(memo))
	fmt.Printf("¿Se puede formar '%s'? %t\n", s, resultado)
}
