package auditor

import "fmt"

// FileTarget exportado (Campos en mayúscula)
type FileTarget struct {
	Path       string
	Permission int
	IsExposed  bool
	RiskLevel  string
}

// Array estático de permisos críticos
var PermissionCriticos = [2]int{777, 666}

// Función que modifica el struct a través del puntero
func AnalyzerFile(target *FileTarget) {
	// 1. Creación de la vista / sub-slice dentro de la función
	vistaPermisos := PermissionCriticos[0:1] // Contiene únicamente [777]

	coincide := false

	// 2. Comprobamos si el permiso del archivo está en la vista rápida
	for _, permiso := range vistaPermisos {
		if target.Permission == permiso {
			coincide = true
			break
		}
	}

	// 3. Evaluamos el nivel de riesgo según los resultados
	switch {
	case coincide:
		target.IsExposed = true
		target.RiskLevel = "CRITICAL"
		fmt.Println("[!] Permiso crítico detectado en:", target.Path)

	case target.Permission == 666: // Uso de == en lugar de ===
		target.IsExposed = true
		target.RiskLevel = "MEDIUM"

	default:
		target.IsExposed = false
		target.RiskLevel = "LOW"
	}
}
