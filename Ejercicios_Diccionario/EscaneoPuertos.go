package main

import "fmt"

func main() {
	puertos := map[string]string{
		"22":    "SSH",
		"80":    "HTTP",
		"443":   "HTTPS",
		"51820": "WireGuard",
	}

	// Usamos %v para ver la lista completa inicial
	fmt.Printf("Lista inicial de puertos escaneados: %v\n", puertos)

	// Recorremos con bucle for
	for peri, puertos := range puertos {
		fmt.Printf("Puerto %s/TCP abierto -> Servicio: %s\n", peri, puertos)
	}

}
