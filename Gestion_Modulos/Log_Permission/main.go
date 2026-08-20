package main

import (
	"auditor/auditor" // Ajusta según el nombre en tu go.mod
	"fmt"
)

func main() {
	// Nota: Usamos valores enteros decimales directos (777, 666, 644)
	permissions := []auditor.FileTarget{
		{
			Path:       "/etc/shadow",
			Permission: 777,
		},
		{
			Path:       "/var/log/syslog",
			Permission: 666,
		},
		{
			Path:       "/app/config.json",
			Permission: 644,
		},
	}

	// 1. Iteramos por índice y pasamos la dirección de memoria (&)
	for i := range permissions {
		auditor.AnalyzerFile(&permissions[i])
	}

	// 2. Mostramos el reporte final
	fmt.Println("\n=== INFORME FINAL DE AUDITORÍA ===")
	for _, f := range permissions {
		fmt.Printf("Ruta: %-18s | Permiso: %d | Expuesto: %-5t | Riesgo: %s\n",
			f.Path, f.Permission, f.IsExposed, f.RiskLevel)
	}
}
