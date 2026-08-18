package main

import (
	"fmt"
	"port/structure"
)

func main() {
	// Inicializamos el slice con los structs de prueba
	puertos := []structure.PortResult{
		{Port: 22, IsOpen: true},
		{Port: 80, IsOpen: false},
		{Port: 8080, IsOpen: true},
	}

	// Recorremos el slice pasando el PUNTERO (&) de cada elemento
	for i := range puertos {
		structure.AnalyzePort(&puertos[i])
	}

	// Imprimimos el resultado
	fmt.Println("=== RESULTADOS DEL ESCANEO ===")
	for _, p := range puertos {
		fmt.Printf("Puerto: %d | Abierto: %t | Servicio: %s | Riesgo: %s\n",
			p.Port, p.IsOpen, p.Service, p.Risk)
	}
}
