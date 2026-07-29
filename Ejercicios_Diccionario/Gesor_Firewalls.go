package main

import "fmt"

func main() {
	firewall := map[string]bool{
		"SSH":    true,
		"HTTP":   true,
		"TELNET": false,
	}

	// 1. Consultar existencia
	fw := "FTP"
	permitido, ok := firewall[fw]
	if ok {
		fmt.Printf("Existe servicio %s | Permitido: %v\n", fw, permitido)
	} else {
		fmt.Printf("No se encontró el servicio %s\n", fw)
	}

	// 2. Borrar regla
	delete(firewall, "TELNET")

	// 3. Recorrer
	for servicio, estado := range firewall {
		fmt.Printf("Servicio: %s | Permitido: %v\n", servicio, estado)
	}
}
