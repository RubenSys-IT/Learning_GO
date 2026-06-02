package main

import "fmt"

func main() {
	puertos := 0
	fmt.Print("Introduzca la cantidad de puertos abiertos detectados: ")
	fmt.Scan(&puertos)
	if puertos <= 0 {
		fmt.Print("NetGuard: Escaneo limpio. No hay puertos abiertos expuestos.")
		return
	} else {
		fmt.Println("NetGuard: Iniciando análisis de vulnerabilidades en los puertos...")
	}
	switch {
	case puertos >= 1 && puertos <= 5:
		fmt.Println("Riesgo: BAJO. Servicios estándar expuestos.")
		for auditados := 1; auditados <= 3; {
			fmt.Println("Verificando seguridad del puerto número", auditados)
			auditados++
		}

	case puertos >= 6 && puertos <= 12:
		fmt.Println("Riesgo: MEDIO. Demasiados servicios activos en el host.")
		for auditados := 1; auditados <= 3; {
			fmt.Println("Verificando seguridad del puerto número", auditados)
			auditados++
		}

	case puertos >= 12:
		fmt.Println("Riesgo: CRÍTICO. Posible vector de ataque masivo detectado.")
		puertosCerrados := 1
		for {
			fmt.Println("Cerrando a la fuerza el puerto crítico número..", puertosCerrados)
			if puertosCerrados == 4 {
				fmt.Println("Mitigación crítica completada. Host aislado provisionalmente.")
				break
			}
			puertosCerrados++
		}
	}
}
