package main

import "fmt"

func recuperar(p *int) {
	*p = *p + 2
}

func main() {
	// Usar [] sin tamaño lo convierte en un slice (más flexible)
	notas := []int{3, 8, 2, 9, 5}

	for i := 0; i < len(notas); i++ {
		if notas[i] >= 5 {
			fmt.Printf("Estudiante aprobado con %d\n", notas[i])
		} else {
			recuperar(&notas[i])
			// Usamos Printf para que el %d se sustituya por el valor
			fmt.Printf("Nota subida a %d\n", notas[i])
		}
	}
	fmt.Println("Notas finales:", notas)
}
