package main

import (
	"fmt"
	"secscan/analyzer"
)

func main() {
	targets := []analyzer.HeaderTarget{
		{
			Domain: "example.com",
			Headers: map[string]string{
				"Server": "nginx/1.18.0",
			},
		},
		{
			Domain: "secure.local",
			Headers: map[string]string{
				"X-Frame-Options":           "DENY",
				"Strict-Transport-Security": "max-age=31536000",
			},
		},
	}

	// 1. Evaluamos cada target pasando su dirección de memoria (&)
	for i := range targets {
		analyzer.EvaluateTarget(&targets[i])
	}

	// 2. Mostramos el reporte generado
	fmt.Println("=== REPORTE DE AUDITORÍA DE SEGURIDAD ===")
	for _, target := range targets {
		fmt.Printf("\nDominio: %s\n", target.Domain)
		fmt.Printf("Puntuación de Riesgo: %d\n", target.RiskScore)
		fmt.Println("Alertas encontradas:")
		for _, flag := range target.Flags {
			fmt.Printf("  [!] %s\n", flag)
		}
	}
}
