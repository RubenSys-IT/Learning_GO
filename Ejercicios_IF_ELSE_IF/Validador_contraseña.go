package main

import "fmt"

func main() {
	claveIngresada := "Admin123"
	if claveIngresada == "Admin123" {
		fmt.Printf("Acceso concedido al sistema")
	} else if claveIngresada == "" {
		fmt.Printf("Por favor, introduce una contraseña")
	} else {
		fmt.Printf("Contraseña incorrecta")
	}
}
