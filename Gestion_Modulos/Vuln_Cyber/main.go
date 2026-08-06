package main

import (
	"Vuln_Cyber/scanner"
	"fmt"
)

func main() {
	servicios := []scanner.Servicio{
		{Puerto: 21, Activo: true},
		{Puerto: 22, Activo: true},
		{Puerto: 80, Activo: true},
		{Puerto: 23, Activo: false},
	}

	totalAlertas := 0
	resultado := scanner.firewall(servicios, &totalAlertas)

	fmt.Println("Resultados del escaneo:", resultado)
	fmt.Println("El total de alertas es:", totalAlertas)
}
