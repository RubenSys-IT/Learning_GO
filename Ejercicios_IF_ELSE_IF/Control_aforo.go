package main

import "fmt"

func main() {
	invitadosActuales := 800
	capacidadMaxima := 100
	if invitadosActuales >= capacidadMaxima {
		fmt.Printf("Lo sentimos, aforo completo")
	} else if invitadosActuales > 90 {
		fmt.Printf("Quedan pocas entradas")
	} else {
		fmt.Printf("Entrada disponibles")
	}
}
