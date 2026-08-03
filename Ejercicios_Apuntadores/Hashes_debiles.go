package main

import "fmt"

func EvaluarHash(algoritmo string, estado *string) {
	if algoritmo == "MD5" || algoritmo == "SHA1" {
		*estado = "VULNERABLE" // Modificamos directamente el valor original
	} else {
		*estado = "SEGURO" // Modificamos directamente el valor original
	}
}

func main() {
	algoritmos := []string{"MD5", "SHA256", "SHA1", "ARGON2"}

	for _, algo := range algoritmos {
		estado := "Pendiente"

		// Corregido el orden: primero el string ('algo'), luego el puntero ('&estado')
		EvaluarHash(algo, &estado)

		fmt.Printf("Algoritmo: %s | Estado: %s\n", algo, estado)
	}
}
