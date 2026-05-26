// Analizador de Puertos y Protocolos
package main

import "fmt"

func main() {
	//Coger lo que introduce el usuario
	var puerto int
	fmt.Print("Puerto: ")
	fmt.Scan(&puerto)

	var origen string
	fmt.Print("Origen: ")
	fmt.Scan(&origen)

	//Primera capa

	if origen == "WAN" {
		fmt.Printf("Alerta: Conexión externa detectada. Analizando puerto...")
	} else {
		fmt.Printf("Conexión interna de confianza.\n")
	}
	switch puerto {
	case 21, 22, 23:
		fmt.Println("Tipo: Protocolo de administración o transferencia (Crítico)")
	case 80, 443:
		fmt.Println("Tipo: Tráfico Web Estándar")
	default:
		fmt.Println("Tipo: Puerto desconocido o no estándar")
	}
}
