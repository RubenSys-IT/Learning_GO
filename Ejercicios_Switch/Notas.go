//Este es un clasificador de Notas

package main

import "fmt"

func main() {
	nota := 5
	switch {
	case nota >= 9:
		fmt.Println("Sobresaliente")
	case nota >= 7:
		fmt.Println("Notable")
	case nota >= 5:
		fmt.Println("Aprobado")
	default:
		fmt.Println("Suspenso")
	}
}
