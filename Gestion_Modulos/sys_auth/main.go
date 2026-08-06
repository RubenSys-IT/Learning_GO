package main

import (
	"fmt"
	"tokens/session" // Asegúrate de que "tokens" coincida con tu go.mod
)

func main() {
	tokensActivos := map[string]bool{
		"token_ftp_21":    true,
		"token_ssh_22":    true,
		"token_http_80":   true,
		"token_telnet_23": false,
	}

	totaltokens := 0

	fmt.Println("Sesiones antes de purgar:", tokensActivos)

	// Pasamos el mapa y la dirección de memoria de totaltokens
	session.PurgarExpiradas(tokensActivos, &totaltokens)

	fmt.Println("Sesiones activas tras la purga:", tokensActivos)
	fmt.Println("Total de sesiones eliminadas:", totaltokens)
}
