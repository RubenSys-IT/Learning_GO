package main

import "fmt"

func main() {
	ipsBloqueadas := map[string]string{
		"10.0.0.5":  "Intento SSH",
		"10.0.0.8":  "Escaneo Nmap",
		"10.0.0.12": "Falso positivo",
	}

	// Usamos %v para ver la lista completa inicial
	fmt.Printf("Lista inicial de IPs bloqueadas: %v\n", ipsBloqueadas)

	// Borramos el falso positivo
	delete(ipsBloqueadas, "10.0.0.12")

	// Comprobamos el resultado final
	fmt.Printf("IPs bloqueadas tras la limpieza: %v\n", ipsBloqueadas)
}
