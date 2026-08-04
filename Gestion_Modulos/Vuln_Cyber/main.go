package main

import (
	"Vuln_Cyber/scanner"
	"fmt"
)

func main() {
	servicios := []scanner.Servicio{
		{Nombre: "FTP", Puerto: 21, Abierto: true},
		{Nombre: "SSH", Puerto: 22, Abierto: true},
		{Nombre: "HTTP", Puerto: 80, Abierto: true},
		{Nombre: "TELNET", Puerto: 23, Abierto: false},
	}

	totalAlertas := 0
	resultado := scanner.EvaluarServicios(servicios, &totalAlertas)

	fmt.Println("Resultados del escaneo:", resultado)
	fmt.Println("El total de alertas es:", totalAlertas)
}
