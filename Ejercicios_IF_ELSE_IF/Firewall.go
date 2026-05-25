package main

import "fmt"

func main() {
	tipotrafico := "SSH"
	puerto := 443
	esIpsospechosa := false

	if esIpsospechosa == true {
		fmt.Printf("Bloqueado, IP a lista negra")
	} else if tipotrafico == "SSH" && puerto != 22 {
		fmt.Printf("ALERTA, Trafico SSH por puerto no estandar")
	} else if tipotrafico == "HTTPS" && puerto == 443 {
		fmt.Printf("Permitido, conexión segura")
	} else {
		fmt.Printf("Revisar Tráfico bajo análisis")
	}
}
