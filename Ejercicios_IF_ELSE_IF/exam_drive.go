package main

import "fmt"

func main() {
	// Usando la asignación corta que ya conoces para simplificar
	tieneTeorico := true
	tienePractico := true

	// Si tienePractico (es true) Y tieneTeorico (es true)
	if tienePractico && tieneTeorico {
		fmt.Println("¡Felicidades! Tienes tu carnet de conducir")
	} else {
		// Si cualquiera de los dos es false, entra aquí directo
		fmt.Println("Aún no cumples los requisitos. ¡Ánimos!")
	}
}
