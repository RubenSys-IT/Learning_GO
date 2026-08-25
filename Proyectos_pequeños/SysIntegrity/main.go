package main

import (
	"fmt"
	"sys/checker" // Ajusta según el nombre de tu módulo en go.mod
)

func main() {
	// Inicializamos el slice con structs reales
	peligroso := []checker.AuditTarget{
		{Path: "/etc/shadow", ExpectedPerm: "600", IsCritical: true},
		{Path: "/etc/passwd", ExpectedPerm: "644", IsCritical: true},
		{Path: "/tmp/script.sh", ExpectedPerm: "777", IsCritical: false},
	}

	// Iteramos pasando la dirección de memoria (&)
	for i := range peligroso {
		checker.AuditSystem(&peligroso[i])
	}

	// Imprimimos el resultado del auditor de integridad
	fmt.Println("=== INFORME DE INTEGRIDAD DEL SISTEMA ===")
	for _, p := range peligroso {
		fmt.Printf("Ruta: %-16s | Estado: %-7s | Crítico: %-5t | Riesgo: %s\n",
			p.Path, p.Status, p.IsCritical, p.RiskLevel)
	}
}
