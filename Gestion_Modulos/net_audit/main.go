package main

import (
	"fmt"
	"fw/firewall"
)

func main() {
	// Creamos un slice con estructuras firewall.Regla
	misReglas := []firewall.Regla{
		{Puerto: 80, Activa: true},
		{Puerto: 22, Activa: true},
		{Puerto: 443, Activa: true},
	}

	totalCambios := 0

	fmt.Println("Reglas antes:", misReglas)

	// Pasamos la dirección de memoria del slice (&misReglas) y del contador (&totalCambios)
	firewall.DesactivarPuerto(&misReglas, 22, &totalCambios)

	fmt.Println("Reglas después:", misReglas)
	fmt.Println("Total de cambios realizados:", totalCambios)
}
