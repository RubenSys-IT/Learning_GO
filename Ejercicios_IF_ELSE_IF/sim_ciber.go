package main

import "fmt"

func main() {
	tipoEvento := "ATTACK"

	if tipoEvento == "ATTACK" {
		fmt.Printf("BLOQUEO DE IP INMEDIATAMENTE")
	} else if tipoEvento == "LOGIN" {
		fmt.Printf("Usuario autenticado correctamente")
	} else {
		fmt.Printf("Evento registrado en el historial")
	}
}
